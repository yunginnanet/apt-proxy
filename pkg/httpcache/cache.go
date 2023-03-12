package httpcache

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"hash/fnv"
	"io"
	"log"
	"net/http"
	"net/textproto"
	"os"
	pathutil "path"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/afero"
)

const (
	headerPrefix = "header/"
	bodyPrefix   = "body/"
	formatPrefix = "v1/"
)

// ErrNotFoundInCache is returned when a resource doesn't exist
var ErrNotFoundInCache = errors.New("not found in cache")

type Cache interface {
	Header(key string) (Header, error)
	Store(res *Resource, keys ...string) error
	Retrieve(key string) (*Resource, error)
	Invalidate(keys ...string)
	Freshen(res *Resource, keys ...string) error
}

// cache provides a storage mechanism for cached Resources
type cache struct {
	fs     afero.Fs
	chroot string
	stale  map[string]time.Time
}

func (c *cache) Open(path string) (afero.File, error) {
	return c.fs.Open(path)
}

func (c *cache) OpenFile(path string, flag int, perm os.FileMode) (afero.File, error) {
	return c.fs.OpenFile(path, flag, perm)
}

func (c *cache) Lstat(path string) (os.FileInfo, error) {
	return c.fs.Stat(path) // FIXME
}

func (c *cache) Stat(path string) (os.FileInfo, error) {
	return c.fs.Stat(path)
}

func (c *cache) ReadDir(path string) ([]os.FileInfo, error) {
	dir, err := c.fs.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = dir.Close()
	}()
	var res []os.FileInfo
	res, err = dir.Readdir(-1)
	return res, err
}

func (c *cache) Mkdir(path string, perm os.FileMode) error {
	return c.fs.Mkdir(path, perm)
}

func (c *cache) Remove(path string) error {
	return c.fs.Remove(path)
}

func (c *cache) String() string {
	return fmt.Sprintf("Cache %s", "afero")
}

var _ Cache = (*cache)(nil)

type Header struct {
	http.Header
	StatusCode int
}

// NewVFSCache returns a cache backend off the provided VFS
func NewVFSCache(fs afero.Fs) Cache {
	return &cache{fs: fs, stale: map[string]time.Time{}}
}

// NewMemoryCache returns an ephemeral cache in memory
func NewMemoryCache() Cache {
	return NewVFSCache(afero.NewMemMapFs())
}

// NewDiskCache returns a disk-backed cache
func NewDiskCache(dir string) (Cache, error) {
	if err := os.MkdirAll(dir, 0750); err != nil {
		return nil, err
	}
	fs := &cache{fs: afero.NewBasePathFs(afero.NewOsFs(), dir), chroot: dir, stale: map[string]time.Time{}}
	return NewVFSCache(fs.fs), nil
}

func (c *cache) vfsWrite(path string, r io.Reader) error {
	if err := c.fs.MkdirAll(pathutil.Dir(path), 0700); err != nil {
		return err
	}
	f, err := c.fs.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := io.Copy(f, r); err != nil {
		return err
	}
	return nil
}

// Retrieve the Status and Headers for a given key path
func (c *cache) Header(key string) (Header, error) {
	path := headerPrefix + formatPrefix + hashKey(key)
	f, err := c.fs.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Header{}, ErrNotFoundInCache
		}
		return Header{}, err
	}

	return readHeaders(bufio.NewReader(f))
}

// Store a resource against a number of keys
func (c *cache) Store(res *Resource, keys ...string) error {
	var buf = &bytes.Buffer{}

	if length, err := strconv.ParseInt(res.Header().Get("Content-Length"), 10, 64); err == nil {
		if _, err = io.CopyN(buf, res, length); err != nil {
			return err
		}
	} else if _, err = io.Copy(buf, res); err != nil {
		return err
	}

	for _, key := range keys {
		delete(c.stale, key)

		if err := c.storeBody(buf, key); err != nil {
			return err
		}

		if err := c.storeHeader(res.Status(), res.Header(), key); err != nil {
			return err
		}
	}

	return nil
}

func (c *cache) storeBody(r io.Reader, key string) error {
	if err := c.vfsWrite(bodyPrefix+formatPrefix+hashKey(key), r); err != nil {
		return err
	}
	return nil
}

func (c *cache) storeHeader(code int, h http.Header, key string) error {
	hb := &bytes.Buffer{}
	hb.Write([]byte(fmt.Sprintf("HTTP/1.1 %d %s\r\n", code, http.StatusText(code))))
	if err := headersToWriter(h, hb); err != nil {
		return err
	}
	if err := c.vfsWrite(headerPrefix+formatPrefix+hashKey(key), bytes.NewReader(hb.Bytes())); err != nil {
		return err
	}
	return nil
}

// Retrieve returns a cached Resource for the given key
func (c *cache) Retrieve(key string) (*Resource, error) {
	f, err := c.fs.Open(bodyPrefix + formatPrefix + hashKey(key))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNotFoundInCache
		}
		return nil, err
	}
	h, err := c.Header(key)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNotFoundInCache
		}
		return nil, err
	}
	res := NewResource(h.StatusCode, f, h.Header)
	if staleTime, exists := c.stale[key]; exists {
		if !res.DateAfter(staleTime) {
			log.Printf("stale marker of %s found", staleTime)
			res.MarkStale()
		}
	}
	return res, nil
}

func (c *cache) Invalidate(keys ...string) {
	log.Printf("invalidating %q", keys)
	for _, key := range keys {
		c.stale[key] = Clock()
	}
}

func (c *cache) Freshen(res *Resource, keys ...string) error {
	for _, key := range keys {
		if h, err := c.Header(key); err == nil {
			if h.StatusCode == res.Status() && headersEqual(h.Header, res.Header()) {
				debugf("freshening key %s", key)
				if err := c.storeHeader(h.StatusCode, res.Header(), key); err != nil {
					return err
				}
			} else {
				debugf("freshen failed, invalidating %s", key)
				c.Invalidate(key)
			}
		}
	}
	return nil
}

func hashKey(key string) string {
	h := fnv.New64a()
	_, err := h.Write([]byte(key))
	if err != nil {
		return "unable-to-calculate"
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func readHeaders(r *bufio.Reader) (Header, error) {
	tp := textproto.NewReader(r)
	line, err := tp.ReadLine()
	if err != nil {
		return Header{}, err
	}

	f := strings.SplitN(line, " ", 3)
	if len(f) < 2 {
		return Header{}, fmt.Errorf("malformed HTTP response: %s", line)
	}
	statusCode, err := strconv.Atoi(f[1])
	if err != nil {
		return Header{}, fmt.Errorf("malformed HTTP status code: %s", f[1])
	}

	mimeHeader, err := tp.ReadMIMEHeader()
	if err != nil {
		return Header{}, err
	}
	return Header{StatusCode: statusCode, Header: http.Header(mimeHeader)}, nil
}

func headersToWriter(h http.Header, w io.Writer) error {
	if err := h.Write(w); err != nil {
		return err
	}
	// ReadMIMEHeader expects a trailing newline
	_, err := w.Write([]byte("\r\n"))
	return err
}

package server

import (
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/soulteary/apt-proxy/pkgs/system"
)

const (
	INTERNAL_PAGE_HOME string = "/"
)

const (
	TYPE_NOT_FOUND int = 0
	TYPE_HOME          = 1
)

func isInternalUrls(url string) bool {
	if strings.Contains(url, "ubuntu") || strings.Contains(url, "debian") {
		return false
	}
	return url == INTERNAL_PAGE_HOME
}

func getInternalResType(url string) int {
	if url == INTERNAL_PAGE_HOME {
		return TYPE_HOME
	}

	return TYPE_NOT_FOUND
}

func renderInternalUrls(url string, rw *http.ResponseWriter) {
	types := getInternalResType(url)
	if types == TYPE_NOT_FOUND {
		return
	}

	if types == TYPE_HOME {
		cacheSize, _ := system.DirSize("./.aptcache")
		files, _ := ioutil.ReadDir("./.aptcache/header/v1")
		available, _ := system.DiskAvailable()

		memoryUsage, goroutine := system.GetMemoryUsageAndGoroutine()

		io.WriteString(*rw, getBaseTemplate(
			system.HumanFileSize(cacheSize),
			strconv.Itoa(len(files)),
			system.HumanFileSize(available),
			system.HumanFileSize(memoryUsage),
			goroutine,
		))
		return
	}

}

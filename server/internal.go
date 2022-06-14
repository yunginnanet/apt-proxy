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
	INTERNAL_PAGE_PING        = "/_/ping/"
)

const (
	TYPE_NOT_FOUND int = 0
	TYPE_HOME          = 1
	TYPE_PING          = 2
)

func isInternalUrls(url string) bool {
	if strings.Contains(url, "ubuntu") || strings.Contains(url, "debian") {
		return false
	}
	return url == INTERNAL_PAGE_HOME || url == INTERNAL_PAGE_PING
}

func getInternalResType(url string) int {
	if url == INTERNAL_PAGE_HOME {
		return TYPE_HOME
	}

	if url == INTERNAL_PAGE_PING {
		return TYPE_PING
	}

	return TYPE_NOT_FOUND
}

func renderInternalUrls(url string, rw *http.ResponseWriter) {
	types := getInternalResType(url)
	if types == TYPE_NOT_FOUND {
		return
	} else if types == TYPE_HOME {
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
	} else if types == TYPE_PING {
		io.WriteString(*rw, "pong")
	}
}

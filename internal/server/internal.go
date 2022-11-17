package server

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/soulteary/apt-proxy/pkg/system"
)

const (
	INTERNAL_PAGE_HOME string = "/"
	INTERNAL_PAGE_PING string = "/_/ping/"
)

const (
	TYPE_NOT_FOUND int = 0
	TYPE_HOME      int = 1
	TYPE_PING      int = 2
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

const LABEL_NO_VALID_VALUE = "N/A"
const CACHE_META_DIR = "./.aptcache/header/v1"

func renderInternalUrls(url string, rw *http.ResponseWriter) {
	types := getInternalResType(url)
	if types == TYPE_NOT_FOUND {
		return
	} else if types == TYPE_HOME {

		cacheSizeLabel := LABEL_NO_VALID_VALUE
		cacheSize, err := system.DirSize("./.aptcache")
		if err == nil {
			cacheSizeLabel = system.ByteCountDecimal(cacheSize)
		}

		filesNumberLabel := LABEL_NO_VALID_VALUE
		if _, err := os.Stat(CACHE_META_DIR); !os.IsNotExist(err) {
			files, err := os.ReadDir(CACHE_META_DIR)
			if err == nil {
				filesNumberLabel = strconv.Itoa(len(files))
			}
		}

		diskAvailableLabel := LABEL_NO_VALID_VALUE
		available, err := system.DiskAvailable()
		if err == nil {
			diskAvailableLabel = system.ByteCountDecimal(available)
		}

		memoryUsageLabel := LABEL_NO_VALID_VALUE
		memoryUsage, goroutine := system.GetMemoryUsageAndGoroutine()
		memoryUsageLabel = system.ByteCountDecimal(memoryUsage)

		_, err = io.WriteString(*rw, getBaseTemplate(
			cacheSizeLabel,
			filesNumberLabel,
			diskAvailableLabel,
			memoryUsageLabel,
			goroutine,
		))
		if err != nil {
			log.Println("Failed to render page")
		}
	} else if types == TYPE_PING {
		_, err := io.WriteString(*rw, "pong")
		if err != nil {
			log.Println("Response ping failed")
		}
	}
}

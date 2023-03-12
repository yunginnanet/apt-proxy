package server

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/soulteary/apt-proxy/pkg/system"
)

const (
	InternalPageHome string = "/"
	InternalPagePing string = "/_/ping/"
)

const (
	TypeNotFound int = 0
	TypeHome     int = 1
	TypePing     int = 2
)

func IsInternalUrls(url string) bool {
	u := strings.ToLower(url)
	return !(strings.Contains(u, "/ubuntu") || strings.Contains(u, "/debian") || strings.Contains(u, "/centos") || strings.Contains(u, "/alpine"))
}

func GetInternalResType(url string) int {
	if url == InternalPageHome {
		return TypeHome
	}

	if url == InternalPagePing {
		return TypePing
	}

	return TypeNotFound
}

// TODO: use configuration
const CacheMetaDir = "./.aptcache/header/v1"
const LabelNoValidValue = "N/A"

func RenderInternalUrls(url string) (string, int) {
	switch GetInternalResType(url) {
	case TypeHome:
		cacheSizeLabel := LabelNoValidValue
		// TODO: use configuration
		cacheSize, err := system.DirSize("./.aptcache")
		if err == nil {
			cacheSizeLabel = system.ByteCountDecimal(cacheSize)
			// } else {
			// return "Get Cache Size Failed", http.StatusBadGateway
		}

		filesNumberLabel := LabelNoValidValue
		if _, err := os.Stat(CacheMetaDir); !os.IsNotExist(err) {
			files, err := os.ReadDir(CacheMetaDir)
			if err == nil {
				filesNumberLabel = strconv.Itoa(len(files))
				// } else {
				// return "Get Cache Meta Dir Failed", http.StatusBadGateway
			}
			// } else {
			// return "Get Cache Meta Failed", http.StatusBadGateway
		}

		diskAvailableLabel := LabelNoValidValue
		available, err := system.DiskAvailable()
		if err == nil {
			diskAvailableLabel = system.ByteCountDecimal(available)
			// } else {
			// return "Get Disk Available Failed", http.StatusBadGateway
		}

		memoryUsageLabel := LabelNoValidValue
		memoryUsage, goroutine := system.GetMemoryUsageAndGoroutine()
		memoryUsageLabel = system.ByteCountDecimal(memoryUsage)

		return GetBaseTemplate(cacheSizeLabel, filesNumberLabel, diskAvailableLabel, memoryUsageLabel, goroutine), 200
	case TypePing:
		return "pong", http.StatusOK
	}
	return "Not Found", http.StatusNotFound
}

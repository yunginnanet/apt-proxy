package server

import (
	"io"
	"net/http"
)

const (
	INTERNAL_PAGE_HOME string = "/"
	INTERNAL_PAGE_HELP        = "/_/help"
)

const (
	TYPE_NOT_FOUND int = 0
	TYPE_HOME          = 1
	TYPE_HELP          = 2
)

func isInternalUrls(url string) bool {
	return url == INTERNAL_PAGE_HOME ||
		url == INTERNAL_PAGE_HELP
}

func getInternalResType(url string) int {
	if url == INTERNAL_PAGE_HOME {
		return TYPE_HOME
	}

	if url == INTERNAL_PAGE_HELP {
		return TYPE_HELP
	}
	return TYPE_NOT_FOUND
}

// TODO Home, Help, Stats
func renderInternalUrls(url string, rw *http.ResponseWriter) {
	if getInternalResType(url) != 0 {
		io.WriteString(*rw, "Hello, APT Proxy!\n")
	}
}

package state

import (
	"net/url"
)

var UBUNTU_MIRROR *url.URL
var DEBIAN_MIRROR *url.URL
var PROXY_MODE = 0

func SetProxyMode(mode int) {
	PROXY_MODE = mode
}

func GetProxyMode() int {
	return PROXY_MODE
}

func SetUbuntuMirror(mirror string) {
	url, err := url.Parse(mirror)
	if err != nil || mirror == "" {
		UBUNTU_MIRROR = nil
	}
	UBUNTU_MIRROR = url
}

func GetUbuntuMirror() *url.URL {
	return UBUNTU_MIRROR
}

func ClearUbuntuMirror() {
	UBUNTU_MIRROR = nil
}

func SetDebianMirror(mirror string) {
	url, err := url.Parse(mirror)
	if err != nil || mirror == "" {
		DEBIAN_MIRROR = nil
	}
	DEBIAN_MIRROR = url
}

func GetDebianMirror() *url.URL {
	return DEBIAN_MIRROR
}

func ClearDebianMirror() {
	DEBIAN_MIRROR = nil
}

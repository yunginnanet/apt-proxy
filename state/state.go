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

func SetUbuntuMirror(input string) {
	mirror := ""
	alias := getUbuntuMirrorByAliases(input)
	if alias == "" {
		mirror = input
	} else {
		mirror = alias
	}

	url, err := url.Parse(mirror)
	if err != nil || mirror == "" {
		UBUNTU_MIRROR = nil
		return
	}
	UBUNTU_MIRROR = url
}

func GetUbuntuMirror() *url.URL {
	return UBUNTU_MIRROR
}

func ResetUbuntuMirror() {
	UBUNTU_MIRROR = nil
}

func SetDebianMirror(input string) {
	mirror := ""
	alias := getDebianMirrorByAliases(input)
	if alias == "" {
		mirror = input
	} else {
		mirror = alias
	}

	url, err := url.Parse(mirror)
	if err != nil || mirror == "" {
		DEBIAN_MIRROR = nil
		return
	}
	DEBIAN_MIRROR = url
}

func GetDebianMirror() *url.URL {
	return DEBIAN_MIRROR
}

func ResetDebianMirror() {
	DEBIAN_MIRROR = nil
}

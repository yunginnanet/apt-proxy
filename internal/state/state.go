package state

import (
	"net/url"

	"github.com/soulteary/apt-proxy/internal/mirrors"
)

var UBUNTU_MIRROR *url.URL
var DEBIAN_MIRROR *url.URL
var CENTOS_MIRROR *url.URL
var PROXY_MODE = 0

func SetProxyMode(mode int) {
	PROXY_MODE = mode
}

func GetProxyMode() int {
	return PROXY_MODE
}

func SetUbuntuMirror(input string) {
	if input == "" {
		UBUNTU_MIRROR = nil
		return
	}

	mirror := input
	alias := mirrors.GetUbuntuMirrorByAliases(input)
	if alias != "" {
		mirror = alias
	}

	url, err := url.Parse(mirror)
	if err != nil {
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
	if input == "" {
		DEBIAN_MIRROR = nil
		return
	}

	mirror := input
	alias := mirrors.GetDebianMirrorByAliases(input)
	if alias != "" {
		mirror = alias
	}

	url, err := url.Parse(mirror)
	if err != nil {
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

func SetCentOSMirror(input string) {
	if input == "" {
		CENTOS_MIRROR = nil
		return
	}

	mirror := input
	alias := mirrors.GetCentOSMirrorByAliases(input)
	if alias != "" {
		mirror = alias
	}

	url, err := url.Parse(mirror)
	if err != nil {
		CENTOS_MIRROR = nil
		return
	}
	CENTOS_MIRROR = url
}

func GetCentOSMirror() *url.URL {
	return CENTOS_MIRROR
}

func ResetCentOSMirror() {
	CENTOS_MIRROR = nil
}

package state

var UBUNTU_MIRROR = ""
var DEBIAN_MIRROR = ""
var PROXY_MODE = 0

func SetProxyMode(mode int) {
	PROXY_MODE = mode
}

func GetProxyMode() int {
	return PROXY_MODE
}

func SetUbuntuMirror(mirror string) {
	UBUNTU_MIRROR = mirror
}

func GetUbuntuMirror() string {
	return UBUNTU_MIRROR
}

func SetDebianMirror(mirror string) {
	DEBIAN_MIRROR = mirror
}

func GetDebianMirror() string {
	return DEBIAN_MIRROR
}

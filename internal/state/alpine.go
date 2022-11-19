package state

import (
	"net/url"

	Define "github.com/soulteary/apt-proxy/internal/define"
	Mirrors "github.com/soulteary/apt-proxy/internal/mirrors"
)

var ALPINE_MIRROR *url.URL

func SetAlpineMirror(input string) {
	if input == "" {
		CENTOS_MIRROR = nil
		return
	}

	mirror := input
	alias := Mirrors.GetMirrorURLByAliases(Define.TYPE_LINUX_DISTROS_ALPINE, input)
	if alias != "" {
		mirror = alias
	}

	url, err := url.Parse(mirror)
	if err != nil {
		ALPINE_MIRROR = nil
		return
	}
	CENTOS_MIRROR = url
}

func GetAlpineMirror() *url.URL {
	return CENTOS_MIRROR
}

func ResetAlpineMirror() {
	CENTOS_MIRROR = nil
}

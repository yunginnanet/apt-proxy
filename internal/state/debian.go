package state

import (
	"net/url"

	Define "github.com/soulteary/apt-proxy/internal/define"
	Mirrors "github.com/soulteary/apt-proxy/internal/mirrors"
)

var DebianMirror *url.URL

func SetDebianMirror(input string) {
	if input == "" {
		DebianMirror = nil
		return
	}

	mirror := input
	alias := Mirrors.GetMirrorURLByAliases(Define.TypeLinuxDistrosDebian, input)
	if alias != "" {
		mirror = alias
	}

	url, err := url.Parse(mirror)
	if err != nil {
		DebianMirror = nil
		return
	}
	DebianMirror = url
}

func GetDebianMirror() *url.URL {
	return DebianMirror
}

func ResetDebianMirror() {
	DebianMirror = nil
}

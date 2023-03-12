package state

import (
	"net/url"

	Define "github.com/soulteary/apt-proxy/internal/define"
	Mirrors "github.com/soulteary/apt-proxy/internal/mirrors"
)

var UbuntuMirror *url.URL

func SetUbuntuMirror(input string) {
	if input == "" {
		UbuntuMirror = nil
		return
	}

	mirror := input
	alias := Mirrors.GetMirrorURLByAliases(Define.TypeLinuxDistrosUbuntu, input)
	if alias != "" {
		mirror = alias
	}

	url, err := url.Parse(mirror)
	if err != nil {
		UbuntuMirror = nil
		return
	}
	UbuntuMirror = url
}

func GetUbuntuMirror() *url.URL {
	return UbuntuMirror
}

func ResetUbuntuMirror() {
	UbuntuMirror = nil
}

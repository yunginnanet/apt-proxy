package state

import (
	"net/url"

	Define "github.com/soulteary/apt-proxy/internal/define"
	Mirrors "github.com/soulteary/apt-proxy/internal/mirrors"
)

var AlpineMirror *url.URL

func SetAlpineMirror(input string) {
	if input == "" {
		AlpineMirror = nil
		return
	}

	mirror := input
	alias := Mirrors.GetMirrorURLByAliases(Define.TypeLinuxDistrosAlpine, input)
	if alias != "" {
		mirror = alias
	}

	url, err := url.Parse(mirror)
	if err != nil {
		AlpineMirror = nil
		return
	}
	AlpineMirror = url
}

func GetAlpineMirror() *url.URL {
	return AlpineMirror
}

func ResetAlpineMirror() {
	AlpineMirror = nil
}

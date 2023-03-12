package state

import (
	"net/url"

	Define "github.com/soulteary/apt-proxy/internal/define"
	Mirrors "github.com/soulteary/apt-proxy/internal/mirrors"
)

var CentosMirror *url.URL

func SetCentOSMirror(input string) {
	if input == "" {
		CentosMirror = nil
		return
	}

	mirror := input
	alias := Mirrors.GetMirrorURLByAliases(Define.TypeLinuxDistrosCentos, input)
	if alias != "" {
		mirror = alias
	}

	url, err := url.Parse(mirror)
	if err != nil {
		CentosMirror = nil
		return
	}
	CentosMirror = url
}

func GetCentOSMirror() *url.URL {
	return CentosMirror
}

func ResetCentOSMirror() {
	CentosMirror = nil
}

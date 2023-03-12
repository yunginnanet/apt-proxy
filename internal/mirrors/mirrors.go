package mirrors

import (
	"regexp"
	"strings"

	Define "github.com/soulteary/apt-proxy/internal/define"
)

func GenerateMirrorListByPredefined(osType int) (mirrors []string) {
	var src []Define.UrlWithAlias
	switch osType {
	case Define.TypeLinuxAllDistros:
		src = append(src, Define.UbuntuDefaults...)
		src = append(src, Define.BuildinDebianMirrors...)
		src = append(src, Define.BuildinCentosMirrors...)
		src = append(src, Define.BuildinAlpineMirrors...)
	case Define.TypeLinuxDistrosUbuntu:
		src = Define.UbuntuDefaults
	case Define.TypeLinuxDistrosDebian:
		src = Define.BuildinDebianMirrors
	case Define.TypeLinuxDistrosCentos:
		src = Define.BuildinCentosMirrors
	case Define.TypeLinuxDistrosAlpine:
		src = Define.BuildinAlpineMirrors
	}

	for _, mirror := range src {
		mirrors = append(mirrors, mirror.URL)
	}
	return mirrors
}

var BuildinUbuntuMirrors = GenerateMirrorListByPredefined(Define.TypeLinuxDistrosUbuntu)
var BuildinDebianMirrors = GenerateMirrorListByPredefined(Define.TypeLinuxDistrosDebian)
var BuildinCentosMirrors = GenerateMirrorListByPredefined(Define.TypeLinuxDistrosCentos)
var BuildinAlpineMirrors = GenerateMirrorListByPredefined(Define.TypeLinuxDistrosAlpine)

func GetGeoMirrorUrlsByMode(mode int) (mirrors []string) {
	if mode == Define.TypeLinuxDistrosUbuntu {
		ubuntuMirrorsOnline, err := GetUbuntuMirrorUrlsByGeo()
		if err != nil {
			return BuildinUbuntuMirrors
		}
		return ubuntuMirrorsOnline
	}

	if mode == Define.TypeLinuxDistrosDebian {
		return BuildinDebianMirrors
	}

	if mode == Define.TypeLinuxDistrosCentos {
		return BuildinCentosMirrors
	}

	if mode == Define.TypeLinuxDistrosAlpine {
		return BuildinAlpineMirrors
	}

	mirrors = append(mirrors, BuildinUbuntuMirrors...)
	mirrors = append(mirrors, BuildinDebianMirrors...)
	mirrors = append(mirrors, BuildinCentosMirrors...)
	mirrors = append(mirrors, BuildinAlpineMirrors...)
	return mirrors
}

func GetFullMirrorURL(mirror Define.UrlWithAlias) string {
	if mirror.Http {
		if strings.HasPrefix(mirror.URL, "http://") {
			return mirror.URL
		}
		return "http://" + mirror.URL
	}
	if mirror.Https {
		if strings.HasPrefix(mirror.URL, "https://") {
			return mirror.URL
		}
		return "https://" + mirror.URL
	}
	return "https://" + mirror.URL
}

func GetMirrorURLByAliases(osType int, alias string) string {
	switch osType {
	case Define.TypeLinuxDistrosUbuntu:
		for _, mirror := range Define.UbuntuDefaults {
			if mirror.Alias == alias {
				return GetFullMirrorURL(mirror)
			}
		}
	case Define.TypeLinuxDistrosDebian:
		for _, mirror := range Define.BuildinDebianMirrors {
			if mirror.Alias == alias {
				return GetFullMirrorURL(mirror)
			}
		}
	case Define.TypeLinuxDistrosCentos:
		for _, mirror := range Define.BuildinCentosMirrors {
			if mirror.Alias == alias {
				return GetFullMirrorURL(mirror)
			}
		}
	case Define.TypeLinuxDistrosAlpine:
		for _, mirror := range Define.BuildinAlpineMirrors {
			if mirror.Alias == alias {
				return GetFullMirrorURL(mirror)
			}
		}
	}
	return ""
}

func GetPredefinedConfiguration(proxyMode int) (string, *regexp.Regexp) {
	switch proxyMode {
	case Define.TypeLinuxDistrosUbuntu:
		return Define.UbuntuBenchmakrUrl, Define.UbuntuHostPattern
	case Define.TypeLinuxDistrosDebian:
		return Define.DebianBenchmakrUrl, Define.DebianHostPattern
	case Define.TypeLinuxDistrosCentos:
		return Define.CentosBenchmakrUrl, Define.CentosHostPattern
	case Define.TypeLinuxDistrosAlpine:
		return Define.AlpineBenchmakrUrl, Define.AlpineHostPattern
	}
	return "", nil
}

package mirrors

import (
	"bufio"
	"net/http"
	"regexp"
)

type buildin_custom_mirror struct {
	url   string
	alias string
}

// TODO: combine
func GetUbuntuMirrorByAliases(alias string) string {
	for _, mirror := range BUILDIN_CUSTOM_UBUNTU_MIRRORS {
		if mirror.alias == alias {
			return mirror.url
		}
	}
	return ""
}

// TODO: combine
func GetDebianMirrorByAliases(alias string) string {
	for _, mirror := range BUILDIN_CUSTOM_DEBIAN_MIRRORS {
		if mirror.alias == alias {
			return mirror.url
		}
	}
	return ""
}

// TODO: combine
func GetCentOSMirrorByAliases(alias string) string {
	for _, mirror := range BUILDIN_CUSTOM_CENTOS_MIRRORS {
		if mirror.alias == alias {
			return mirror.url
		}
	}
	return ""
}

func GetGeoMirrorUrlsByMode(mode int) (mirrors []string) {
	if mode == TYPE_LINUX_DISTROS_UBUNTU {
		ubuntuMirrorsOnline, err := getUbuntuMirrorUrlsByGeo()
		if err != nil {
			return BUILDIN_OFFICAL_UBUNTU_MIRRORS
		}
		return ubuntuMirrorsOnline
	}

	if mode == TYPE_LINUX_DISTROS_DEBIAN {
		return BUILDIN_OFFICAL_DEBIAN_MIRRORS
	}

	mirrors = append(mirrors, BUILDIN_OFFICAL_UBUNTU_MIRRORS...)
	mirrors = append(mirrors, BUILDIN_OFFICAL_DEBIAN_MIRRORS...)
	return mirrors
}

func getUbuntuMirrorUrlsByGeo() (mirrors []string, err error) {
	response, err := http.Get(UBUNTU_GEO_MIRROR_API)
	if err != nil {
		return mirrors, err
	}

	defer response.Body.Close()
	scanner := bufio.NewScanner(response.Body)

	for scanner.Scan() {
		mirrors = append(mirrors, scanner.Text())
	}

	return mirrors, scanner.Err()
}

func GetPredefinedConfiguration(proxyMode int) (string, *regexp.Regexp) {
	if proxyMode == TYPE_LINUX_DISTROS_UBUNTU {
		return UBUNTU_BENCHMAKR_URL, UBUNTU_HOST_PATTERN
	} else {
		return DEBIAN_BENCHMAKR_URL, DEBIAN_HOST_PATTERN
	}
}

package mirrors

import (
	"bufio"
	"net/http"
	"regexp"

	Define "github.com/soulteary/apt-proxy/internal/define"
)

func GetMirrorURLByAliases(osType int, alias string) string {
	switch osType {
	case Define.TYPE_LINUX_DISTROS_UBUNTU:
		for _, mirror := range Define.BUILDIN_UBUNTU_MIRRORS {
			if mirror.Alias == alias {
				return mirror.URL
			}
		}
	case Define.TYPE_LINUX_DISTROS_DEBIAN:
		for _, mirror := range Define.BUILDIN_DEBIAN_MIRRORS {
			if mirror.Alias == alias {
				return mirror.URL
			}
		}
	case Define.TYPE_LINUX_DISTROS_CENTOS:
		for _, mirror := range Define.BUILDIN_CENTOS_MIRRORS {
			if mirror.Alias == alias {
				return mirror.URL
			}
		}
	}
	return ""
}

func GetGeoMirrorUrlsByMode(mode int) (mirrors []string) {
	if mode == Define.TYPE_LINUX_DISTROS_UBUNTU {
		ubuntuMirrorsOnline, err := getUbuntuMirrorUrlsByGeo()
		if err != nil {
			return BUILDIN_OFFICAL_UBUNTU_MIRRORS
		}
		return ubuntuMirrorsOnline
	}

	if mode == Define.TYPE_LINUX_DISTROS_DEBIAN {
		return BUILDIN_OFFICAL_DEBIAN_MIRRORS
	}

	if mode == Define.TYPE_LINUX_DISTROS_CENTOS {
		return BUILDIN_OFFICAL_CENTOS_MIRRORS
	}

	mirrors = append(mirrors, BUILDIN_OFFICAL_UBUNTU_MIRRORS...)
	mirrors = append(mirrors, BUILDIN_OFFICAL_DEBIAN_MIRRORS...)
	mirrors = append(mirrors, BUILDIN_OFFICAL_CENTOS_MIRRORS...)
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
	switch proxyMode {
	case Define.TYPE_LINUX_DISTROS_UBUNTU:
		return UBUNTU_BENCHMAKR_URL, UBUNTU_HOST_PATTERN
	case Define.TYPE_LINUX_DISTROS_DEBIAN:
		return DEBIAN_BENCHMAKR_URL, DEBIAN_HOST_PATTERN
	case Define.TYPE_LINUX_DISTROS_CENTOS:
		return CENTOS_BENCHMAKR_URL, CENTOS_HOST_PATTERN
	}
	return "", nil
}

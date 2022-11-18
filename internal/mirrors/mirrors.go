package mirrors

import (
	"bufio"
	"net/http"
	"regexp"

	Define "github.com/soulteary/apt-proxy/internal/define"
)

// TODO:  1. Redefine the official definition list and the three-party list
// TODO:  2. Automatically generate lists with Alias
// TODO:  3. Automatically generate rewriter rules
func GenerateMirrorListByPredefined(osType int) (mirrors []string) {
	var src []Define.UrlWithAlias
	switch osType {
	case Define.TYPE_LINUX_DISTROS_UBUNTU:
		src = Define.BUILDIN_UBUNTU_MIRRORS
	case Define.TYPE_LINUX_DISTROS_DEBIAN:
		src = Define.BUILDIN_DEBIAN_MIRRORS
	case Define.TYPE_LINUX_DISTROS_CENTOS:
		src = Define.BUILDIN_CENTOS_MIRRORS
	}
	for _, mirror := range src {
		mirrors = append(mirrors, mirror.URL)
	}
	return mirrors
}

var BUILDIN_OFFICAL_UBUNTU_MIRRORS = GenerateMirrorListByPredefined(Define.TYPE_LINUX_DISTROS_UBUNTU)
var BUILDIN_OFFICAL_CENTOS_MIRRORS = GenerateMirrorListByPredefined(Define.TYPE_LINUX_DISTROS_CENTOS)
var BUILDIN_OFFICAL_DEBIAN_MIRRORS = GenerateMirrorListByPredefined(Define.TYPE_LINUX_DISTROS_DEBIAN)

var DEBIAN_DEFAULT_CACHE_RULES = []Define.Rule{
	{Pattern: regexp.MustCompile(`deb$`), CacheControl: `max-age=100000`, Rewrite: true, OS: Define.TYPE_LINUX_DISTROS_DEBIAN},
	{Pattern: regexp.MustCompile(`udeb$`), CacheControl: `max-age=100000`, Rewrite: true, OS: Define.TYPE_LINUX_DISTROS_DEBIAN},
	{Pattern: regexp.MustCompile(`DiffIndex$`), CacheControl: `max-age=3600`, Rewrite: true, OS: Define.TYPE_LINUX_DISTROS_DEBIAN},
	{Pattern: regexp.MustCompile(`PackagesIndex$`), CacheControl: `max-age=3600`, Rewrite: true, OS: Define.TYPE_LINUX_DISTROS_DEBIAN},
	{Pattern: regexp.MustCompile(`Packages\.(bz2|gz|lzma)$`), CacheControl: `max-age=3600`, Rewrite: true, OS: Define.TYPE_LINUX_DISTROS_DEBIAN},
	{Pattern: regexp.MustCompile(`SourcesIndex$`), CacheControl: `max-age=3600`, Rewrite: true, OS: Define.TYPE_LINUX_DISTROS_DEBIAN},
	{Pattern: regexp.MustCompile(`Sources\.(bz2|gz|lzma)$`), CacheControl: `max-age=3600`, Rewrite: true, OS: Define.TYPE_LINUX_DISTROS_DEBIAN},
	{Pattern: regexp.MustCompile(`Release(\.gpg)?$`), CacheControl: `max-age=3600`, Rewrite: true, OS: Define.TYPE_LINUX_DISTROS_DEBIAN},
	{Pattern: regexp.MustCompile(`Translation-(en|fr)\.(gz|bz2|bzip2|lzma)$`), CacheControl: `max-age=3600`, Rewrite: true, OS: Define.TYPE_LINUX_DISTROS_DEBIAN},
	// Add file file hash
	{Pattern: regexp.MustCompile(`/by-hash/`), CacheControl: `max-age=3600`, Rewrite: true, OS: Define.TYPE_LINUX_DISTROS_DEBIAN},
}

// Ubuntu
var UBUNTU_DEFAULT_CACHE_RULES = []Define.Rule{
	{Pattern: regexp.MustCompile(`deb$`), CacheControl: `max-age=100000`, Rewrite: true, OS: Define.TYPE_LINUX_DISTROS_UBUNTU},
	{Pattern: regexp.MustCompile(`udeb$`), CacheControl: `max-age=100000`, Rewrite: true, OS: Define.TYPE_LINUX_DISTROS_UBUNTU},
	{Pattern: regexp.MustCompile(`DiffIndex$`), CacheControl: `max-age=3600`, Rewrite: true, OS: Define.TYPE_LINUX_DISTROS_UBUNTU},
	{Pattern: regexp.MustCompile(`PackagesIndex$`), CacheControl: `max-age=3600`, Rewrite: true, OS: Define.TYPE_LINUX_DISTROS_UBUNTU},
	{Pattern: regexp.MustCompile(`Packages\.(bz2|gz|lzma)$`), CacheControl: `max-age=3600`, Rewrite: true, OS: Define.TYPE_LINUX_DISTROS_UBUNTU},
	{Pattern: regexp.MustCompile(`SourcesIndex$`), CacheControl: `max-age=3600`, Rewrite: true, OS: Define.TYPE_LINUX_DISTROS_UBUNTU},
	{Pattern: regexp.MustCompile(`Sources\.(bz2|gz|lzma)$`), CacheControl: `max-age=3600`, Rewrite: true, OS: Define.TYPE_LINUX_DISTROS_UBUNTU},
	{Pattern: regexp.MustCompile(`Release(\.gpg)?$`), CacheControl: `max-age=3600`, Rewrite: true, OS: Define.TYPE_LINUX_DISTROS_UBUNTU},
	{Pattern: regexp.MustCompile(`Translation-(en|fr)\.(gz|bz2|bzip2|lzma)$`), CacheControl: `max-age=3600`, Rewrite: true, OS: Define.TYPE_LINUX_DISTROS_UBUNTU},
	// Add file file hash
	{Pattern: regexp.MustCompile(`/by-hash/`), CacheControl: `max-age=3600`, Rewrite: true, OS: Define.TYPE_LINUX_DISTROS_UBUNTU},
}

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
	response, err := http.Get(Define.UBUNTU_GEO_MIRROR_API)
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
		return Define.UBUNTU_BENCHMAKR_URL, Define.UBUNTU_HOST_PATTERN
	case Define.TYPE_LINUX_DISTROS_DEBIAN:
		return Define.DEBIAN_BENCHMAKR_URL, Define.DEBIAN_HOST_PATTERN
	case Define.TYPE_LINUX_DISTROS_CENTOS:
		return Define.CENTOS_BENCHMAKR_URL, Define.CENTOS_HOST_PATTERN
	}
	return "", nil
}

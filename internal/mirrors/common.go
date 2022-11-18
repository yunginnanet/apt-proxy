package mirrors

import (
	"regexp"

	Define "github.com/soulteary/apt-proxy/internal/define"
)

var BUILDIN_OFFICAL_UBUNTU_MIRRORS = (func() (mirrors []string) {
	for _, mirror := range Define.BUILDIN_UBUNTU_MIRRORS {
		mirrors = append(mirrors, mirror.URL)
	}
	return mirrors
})()

// CENTOS
var BUILDIN_OFFICAL_CENTOS_MIRRORS = (func() (mirrors []string) {
	for _, mirror := range Define.BUILDIN_CENTOS_MIRRORS {
		mirrors = append(mirrors, mirror.URL)
	}
	return mirrors
})()

var CENTOS_HOST_PATTERN = regexp.MustCompile(`https?://.+/centos/(.+)$`)

// DEBIAN
const (
	DEBIAN_BENCHMAKR_URL = "dists/bullseye/main/binary-amd64/Release"
)

var BUILDIN_OFFICAL_DEBIAN_MIRRORS = (func() (mirrors []string) {
	for _, mirror := range Define.BUILDIN_DEBIAN_MIRRORS {
		mirrors = append(mirrors, mirror.URL)
	}
	return mirrors
})()

var DEBIAN_HOST_PATTERN = regexp.MustCompile(
	`https?://(deb|security|snapshot).debian.org/debian/(.+)$`,
)

const CENTOS_BENCHMAKR_URL = "TIME"

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
const (
	UBUNTU_GEO_MIRROR_API = "http://mirrors.ubuntu.com/mirrors.txt"
	UBUNTU_BENCHMAKR_URL  = "dists/jammy/main/binary-amd64/Release"
)

var UBUNTU_HOST_PATTERN = regexp.MustCompile(
	`https?://(\w{2}\.)?(security|archive).ubuntu.com/ubuntu/(.+)$`,
)

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

package linux

import "regexp"

const (
	LINUX_ALL_DISTROS    string = "all"
	LINUX_DISTROS_UBUNTU        = "ubuntu"
	LINUX_DISTROS_DEBIAN        = "debian"
)

type Rule struct {
	Pattern      *regexp.Regexp
	CacheControl string
	Rewrite      bool
}

const (
	BENCHMARK_MAX_TIMEOUT    = 15 // 15 seconds, detect resource timeout
	BENCHMARK_MAX_TRIES      = 3  // times, maximum number of attempts
	BENCHMARK_DETECT_TIMEOUT = 3  // 3 seconds, for select fast mirror
)

// DEBIAN
const (
	DEBIAN_BENCHMAKR_URL = "dists/bullseye/main/binary-amd64/Release"
)

var BUILDIN_DEBIAN_MIRRORS = []string{
	"http://ftp.cn.debian.org/debian/",
	"http://mirror.bjtu.edu.cn/debian/",
	"http://mirror.lzu.edu.cn/debian/",
	"http://mirror.nju.edu.cn/debian/",
	"http://mirrors.163.com/debian/",
	"http://mirrors.bfsu.edu.cn/debian/",
	"http://mirrors.hit.edu.cn/debian/",
	"http://mirrors.huaweicloud.com/debian/",
	"http://mirrors.neusoft.edu.cn/debian/",
	"http://mirrors.tuna.tsinghua.edu.cn/debian/",
	"http://mirrors.ustc.edu.cn/debian/",
}

var DEBIAN_HOST_PATTERN = regexp.MustCompile(
	`https?://(deb|security|snapshot).debian.org/debian/(.+)$`,
)

var DEBIAN_DEFAULT_CACHE_RULES = []Rule{
	{Pattern: regexp.MustCompile(`deb$`), CacheControl: `max-age=100000`, Rewrite: true},
	{Pattern: regexp.MustCompile(`udeb$`), CacheControl: `max-age=100000`, Rewrite: true},
	{Pattern: regexp.MustCompile(`DiffIndex$`), CacheControl: `max-age=3600`, Rewrite: true},
	{Pattern: regexp.MustCompile(`PackagesIndex$`), CacheControl: `max-age=3600`, Rewrite: true},
	{Pattern: regexp.MustCompile(`Packages\.(bz2|gz|lzma)$`), CacheControl: `max-age=3600`, Rewrite: true},
	{Pattern: regexp.MustCompile(`SourcesIndex$`), CacheControl: `max-age=3600`, Rewrite: true},
	{Pattern: regexp.MustCompile(`Sources\.(bz2|gz|lzma)$`), CacheControl: `max-age=3600`, Rewrite: true},
	{Pattern: regexp.MustCompile(`Release(\.gpg)?$`), CacheControl: `max-age=3600`, Rewrite: true},
	{Pattern: regexp.MustCompile(`Translation-(en|fr)\.(gz|bz2|bzip2|lzma)$`), CacheControl: `max-age=3600`, Rewrite: true},
	// Add file file hash
	{Pattern: regexp.MustCompile(`/by-hash/`), CacheControl: `max-age=3600`, Rewrite: true},
}

// Ubuntu
const (
	UBUNTU_GEO_MIRROR_API = "http://mirrors.ubuntu.com/mirrors.txt"
	UBUNTU_BENCHMAKR_URL  = "dists/jammy/main/binary-amd64/Release"
)

var BUILDIN_UBUNTU_MIRRORS = []string{
	"http://mirrors.aliyun.com/ubuntu/",
	"http://mirrors.huaweicloud.com/repository/ubuntu/",
	"http://mirror.dlut.edu.cn/ubuntu/",
	"http://mirrors.dgut.edu.cn/ubuntu/",
	"http://mirrors.njupt.edu.cn/ubuntu/",
	"https://mirrors.hit.edu.cn/ubuntu/",
	"http://mirrors.yun-idc.com/ubuntu/",
	"http://ftp.sjtu.edu.cn/ubuntu/",
	"https://mirror.nju.edu.cn/ubuntu/",
	"https://mirrors.bupt.edu.cn/ubuntu/",
	"http://mirrors.skyshe.cn/ubuntu/",
	"https://repo.huaweicloud.com/ubuntu/",
	"http://mirror.lzu.edu.cn/ubuntu/",
	"http://mirrors.cn99.com/ubuntu/",
	"http://mirrors.cqu.edu.cn/ubuntu/",
	"https://mirrors.cloud.tencent.com/ubuntu/",
	"https://mirrors.ustc.edu.cn/ubuntu/",
	"https://mirror.bjtu.edu.cn/ubuntu/",
	"http://mirrors.sohu.com/ubuntu/",
	"http://archive.ubuntu.com/ubuntu/",
}

var UBUNTU_HOST_PATTERN = regexp.MustCompile(
	`https?://(security|archive).ubuntu.com/ubuntu/(.+)$`,
)

var UBUNTU_DEFAULT_CACHE_RULES = []Rule{
	{Pattern: regexp.MustCompile(`deb$`), CacheControl: `max-age=100000`, Rewrite: true},
	{Pattern: regexp.MustCompile(`udeb$`), CacheControl: `max-age=100000`, Rewrite: true},
	{Pattern: regexp.MustCompile(`DiffIndex$`), CacheControl: `max-age=3600`, Rewrite: true},
	{Pattern: regexp.MustCompile(`PackagesIndex$`), CacheControl: `max-age=3600`, Rewrite: true},
	{Pattern: regexp.MustCompile(`Packages\.(bz2|gz|lzma)$`), CacheControl: `max-age=3600`, Rewrite: true},
	{Pattern: regexp.MustCompile(`SourcesIndex$`), CacheControl: `max-age=3600`, Rewrite: true},
	{Pattern: regexp.MustCompile(`Sources\.(bz2|gz|lzma)$`), CacheControl: `max-age=3600`, Rewrite: true},
	{Pattern: regexp.MustCompile(`Release(\.gpg)?$`), CacheControl: `max-age=3600`, Rewrite: true},
	{Pattern: regexp.MustCompile(`Translation-(en|fr)\.(gz|bz2|bzip2|lzma)$`), CacheControl: `max-age=3600`, Rewrite: true},
	// Add file file hash
	{Pattern: regexp.MustCompile(`/by-hash/`), CacheControl: `max-age=3600`, Rewrite: true},
}

package mirrors

import (
	"fmt"
	"regexp"

	Define "github.com/soulteary/apt-proxy/internal/define"
)

type Rule struct {
	OS           int
	Pattern      *regexp.Regexp
	CacheControl string
	Rewrite      bool
}

type buildin_custom_mirror struct {
	url   string
	alias string
}

var BUILDIN_CUSTOM_CENTOS_MIRRORS = []buildin_custom_mirror{
	{url: "https://mirrors.tuna.tsinghua.edu.cn/centos/", alias: "cn:tsinghua"},
	{url: "http://mirrors.aliyun.com/centos/", alias: "cn:aliyun"},
	{url: "https://mirrors.bfsu.edu.cn/centos/", alias: "cn:bfsu"},
	{url: "https://mirrors.cqu.edu.cn/CentOS/", alias: "cn:cqu"},
	{url: "http://mirrors.neusoft.edu.cn/centos/", alias: "cn:neusoft"},
	{url: "https://mirror.nju.edu.cn/centos/", alias: "cn:nju"},
	{url: "https://mirrors.huaweicloud.com/centos/", alias: "cn:huaweicloud"},
	{url: "http://mirror.lzu.edu.cn/centos/", alias: "cn:lzu"},
	{url: "https://mirrors.njupt.edu.cn/centos/", alias: "cn:njupt"},
	{url: "http://mirrors.163.com/centos/", alias: "cn:163"},
	{url: "https://mirrors.bupt.edu.cn/centos/", alias: "cn:bupt"},
	{url: "https://ftp.sjtu.edu.cn/centos/", alias: "cn:sjtu"},
	{url: "https://mirrors.ustc.edu.cn/centos/", alias: "cn:ustc"},
}

var BUILDIN_CUSTOM_DEBIAN_MIRRORS = []buildin_custom_mirror{
	{url: "https://mirrors.tuna.tsinghua.edu.cn/debian/", alias: "cn:tsinghua"},
	{url: "http://mirrors.ustc.edu.cn/debian/", alias: "cn:ustc"},
	{url: "https://mirrors.163.com/debian/", alias: "cn:163"},
	{url: "https://mirrors.aliyun.com/debian/", alias: "cn:aliyun"},
	{url: "https://repo.huaweicloud.com/debian/", alias: "cn:huawei"},
	{url: "https://mirrors.cloud.tencent.com/debian/", alias: "cn:tencent"},
	{url: "http://ftp.cn.debian.org/debian/", alias: "cn:debian"},
	{url: "http://mirror.bjtu.edu.cn/debian/", alias: "cn:bjtu"},
	{url: "http://mirror.lzu.edu.cn/debian/", alias: "cn:lzu"},
	{url: "http://mirror.nju.edu.cn/debian/", alias: "cn:nju"},
	{url: "http://mirrors.bfsu.edu.cn/debian/", alias: "cn:bfsu"},
	{url: "http://mirrors.hit.edu.cn/debian/", alias: "cn:hit"},
	{url: "http://mirrors.neusoft.edu.cn/debian/", alias: "cn:neusoft"},
}

var BUILDIN_CUSTOM_UBUNTU_MIRRORS = []buildin_custom_mirror{
	{url: "https://mirrors.tuna.tsinghua.edu.cn/ubuntu/", alias: "cn:tsinghua"},
	{url: "http://mirrors.ustc.edu.cn/ubuntu/", alias: "cn:ustc"},
	{url: "https://mirrors.163.com/ubuntu/", alias: "cn:163"},
	{url: "https://mirrors.aliyun.com/ubuntu/", alias: "cn:aliyun"},
	// "http://mirrors.huaweicloud.com/repository/ubuntu/"
	{url: "https://repo.huaweicloud.com/ubuntu/", alias: "cn:huawei"},
	{url: "https://mirrors.cloud.tencent.com/ubuntu/", alias: "cn:tencent"},
	{url: "http://mirror.dlut.edu.cn/ubuntu/", alias: "cn:dlut"},
	{url: "http://mirrors.dgut.edu.cn/ubuntu/", alias: "cn:dgut"},
	{url: "http://mirrors.njupt.edu.cn/ubuntu/", alias: "cn:njupt"},
	{url: "https://mirrors.hit.edu.cn/ubuntu/", alias: "cn:hit"},
	{url: "http://mirrors.yun-idc.com/ubuntu/", alias: "cn:yun-idc"},
	{url: "http://ftp.sjtu.edu.cn/ubuntu/", alias: "cn:sjtu"},
	{url: "https://mirror.nju.edu.cn/ubuntu/", alias: "cn:nju"},
	{url: "https://mirrors.bupt.edu.cn/ubuntu/", alias: "cn:bupt"},
	{url: "http://mirrors.skyshe.cn/ubuntu/", alias: "cn:skyshe"},
	{url: "http://mirror.lzu.edu.cn/ubuntu/", alias: "cn:lzu"},
	{url: "http://mirrors.cn99.com/ubuntu/", alias: "cn:cn99"},
	{url: "http://mirrors.cqu.edu.cn/ubuntu/", alias: "cn:cqu"},
	{url: "https://mirror.bjtu.edu.cn/ubuntu/", alias: "cn:bjtu"},
	{url: "http://mirrors.sohu.com/ubuntu/", alias: "cn:sohu"},
	{url: "http://cn.archive.ubuntu.com/ubuntu/", alias: "cn:ubuntu"},
}

var BUILDIN_OFFICAL_UBUNTU_MIRRORS = (func() (mirrors []string) {
	for _, mirror := range BUILDIN_CUSTOM_UBUNTU_MIRRORS {
		mirrors = append(mirrors, mirror.url)
	}
	return mirrors
})()

// CENTOS
var BUILDIN_OFFICAL_CENTOS_MIRRORS = (func() (mirrors []string) {
	for _, mirror := range BUILDIN_CUSTOM_CENTOS_MIRRORS {
		mirrors = append(mirrors, mirror.url)
	}
	return mirrors
})()

var CENTOS_HOST_PATTERN = regexp.MustCompile(`https?://.+/centos/(.+)$`)

// DEBIAN
const (
	DEBIAN_BENCHMAKR_URL = "dists/bullseye/main/binary-amd64/Release"
)

var BUILDIN_OFFICAL_DEBIAN_MIRRORS = (func() (mirrors []string) {
	for _, mirror := range BUILDIN_CUSTOM_DEBIAN_MIRRORS {
		mirrors = append(mirrors, mirror.url)
	}
	return mirrors
})()

var DEBIAN_HOST_PATTERN = regexp.MustCompile(
	`https?://(deb|security|snapshot).debian.org/debian/(.+)$`,
)

const CENTOS_BENCHMAKR_URL = "TIME"

var DEBIAN_DEFAULT_CACHE_RULES = []Rule{
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

var UBUNTU_DEFAULT_CACHE_RULES = []Rule{
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

func (r *Rule) String() string {
	return fmt.Sprintf("%s Cache-Control=%s Rewrite=%#v",
		r.Pattern.String(), r.CacheControl, r.Rewrite)
}

package define

import "regexp"

const (
	DEBIAN_BENCHMAKR_URL = "dists/bullseye/main/binary-amd64/Release"
)

var DEBIAN_HOST_PATTERN = regexp.MustCompile(
	`https?://(deb|security|snapshot).debian.org/debian/(.+)$`,
)

// https://www.debian.org/mirror/list 2022.11.19
// Sites that contain protocol headers, restrict access to resources using that protocol
var DEBIAN_OFFICAL_MIRRORS = []string{
	"http://ftp.cn.debian.org/debian/",
	"mirror.bjtu.edu.cn/debian/",
	"mirrors.163.com/debian/",
	"mirrors.bfsu.edu.cn/debian/",
	"mirrors.huaweicloud.com/debian/",
	"http://mirrors.neusoft.edu.cn/debian/",
	"mirrors.tuna.tsinghua.edu.cn/debian/",
	"mirrors.ustc.edu.cn/debian/",
}

var DEBIAN_CUSTOM_MIRRORS = []string{
	"repo.huaweicloud.com/debian/",
	"mirrors.cloud.tencent.com/debian/",
	"mirrors.hit.edu.cn/debian/",
	"mirrors.aliyun.com/debian/",
	"mirror.lzu.edu.cn/debian/",
	"mirror.nju.edu.cn/debian/",
}

var BUILDIN_DEBIAN_MIRRORS = GenerateBuildInList(DEBIAN_OFFICAL_MIRRORS, DEBIAN_CUSTOM_MIRRORS)

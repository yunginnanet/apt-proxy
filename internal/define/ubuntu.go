package define

import (
	"regexp"
)

const (
	UBUNTU_GEO_MIRROR_API = "http://mirrors.ubuntu.com/mirrors.txt"
	UBUNTU_BENCHMAKR_URL  = "dists/jammy/main/binary-amd64/Release"
)

var UBUNTU_HOST_PATTERN = regexp.MustCompile(
	`https?://(\w{2}\.)?(security|archive).ubuntu.com/ubuntu/(.+)$`,
)

// http://mirrors.ubuntu.com/mirrors.txt 2022.11.19
// Sites that contain protocol headers, restrict access to resources using that protocol
var UBUNTU_OFFICAL_MIRRORS = []string{
	"mirrors.cn99.com/ubuntu/",
	"mirrors.tuna.tsinghua.edu.cn/ubuntu/",
	"mirrors.cnnic.cn/ubuntu/",
	"mirror.bjtu.edu.cn/ubuntu/",
	"mirrors.cqu.edu.cn/ubuntu/",
	"http://mirrors.skyshe.cn/ubuntu/",
	// duplicate "mirrors.tuna.tsinghua.edu.cn/ubuntu-ports/",
	"mirrors.yun-idc.com/ubuntu/",
	"http://mirror.dlut.edu.cn/ubuntu/",
	"mirrors.xjtu.edu.cn/ubuntu/",
	"mirrors.huaweicloud.com/repository/ubuntu/",
	"mirrors.bupt.edu.cn/ubuntu/",
	"mirrors.hit.edu.cn/ubuntu/",
	// duplicate "repo.huaweicloud.com/ubuntu/",
	"http://mirrors.sohu.com/ubuntu/",
	"mirror.nju.edu.cn/ubuntu/",
	"mirrors.bfsu.edu.cn/ubuntu/",
	"mirror.lzu.edu.cn/ubuntu/",
	"mirrors.aliyun.com/ubuntu/",
	"ftp.sjtu.edu.cn/ubuntu/",
	"mirrors.njupt.edu.cn/ubuntu/",
	"mirrors.cloud.tencent.com/ubuntu/",
	"http://mirrors.dgut.edu.cn/ubuntu/",
	"mirrors.ustc.edu.cn/ubuntu/",
	"mirrors.sdu.edu.cn/ubuntu/",
	"http://cn.archive.ubuntu.com/ubuntu/",
}

var UBUNTU_CUSTOM_MIRRORS = []string{
	"mirrors.163.com/ubuntu/",
}

var BUILDIN_UBUNTU_MIRRORS = GenerateBuildInList(UBUNTU_OFFICAL_MIRRORS, UBUNTU_CUSTOM_MIRRORS)

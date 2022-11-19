package define

import "regexp"

var CENTOS_HOST_PATTERN = regexp.MustCompile(`https?://.+/centos/(.+)$`)

const CENTOS_BENCHMAKR_URL = "TIME"

// https://www.centos.org/download/mirrors/ 2022.11.19
// Sites that contain protocol headers, restrict access to resources using that protocol
var CENTOS_OFFICAL_MIRRORS = []string{
	"mirrors.bfsu.edu.cn/centos/",
	"mirrors.cqu.edu.cn/CentOS/",
	"http://mirrors.neusoft.edu.cn/centos/",
	"mirrors.nju.edu.cn/centos/",
	"mirrors.huaweicloud.com/centos/",
	"mirror.lzu.edu.cn/centos/",
	"http://mirrors.njupt.edu.cn/centos/",
	"mirrors.163.com/centos/",
	"mirrors.bupt.edu.cn/centos/",
	"ftp.sjtu.edu.cn/centos/",
	"mirrors.tuna.tsinghua.edu.cn/centos/",
	"mirrors.ustc.edu.cn/centos/",
}

var CENTOS_CUSTOM_MIRRORS = []string{
	"http://mirrors.aliyun.com/centos/",
}

var BUILDIN_CENTOS_MIRRORS = GenerateBuildInList(CENTOS_OFFICAL_MIRRORS, CENTOS_CUSTOM_MIRRORS)

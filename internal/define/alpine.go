package define

import "regexp"

var ALPINE_HOST_PATTERN = regexp.MustCompile(`https?://.+/alpine/(.+)$`)

const ALPINE_BENCHMAKR_URL = "MIRRORS.txt"

// https://mirrors.alpinelinux.org/ 2022.11.19
// Sites that contain protocol headers, restrict access to resources using that protocol
var ALPINE_OFFICAL_MIRRORS = []string{
	"mirrors.tuna.tsinghua.edu.cn/alpine/",
	"mirrors.ustc.edu.cn/alpine/",
	"mirrors.nju.edu.cn/alpine/",
	// offline "mirror.lzu.edu.cn/alpine/",
	"mirrors.sjtug.sjtu.edu.cn/alpine/",
	"mirrors.aliyun.com/alpine/",
	// not vaild "mirrors.bfsu.edu.cn/alpine",
	// offline "mirrors.neusoft.edu.cn/alpine/",
}

var ALPINE_CUSTOM_MIRRORS = []string{}

var BUILDIN_ALPINE_MIRRORS = GenerateBuildInList(ALPINE_OFFICAL_MIRRORS, ALPINE_CUSTOM_MIRRORS)

package define

import "regexp"

var ALPINE_HOST_PATTERN = regexp.MustCompile(`https?://.+/alpine/(.+)$`)

const ALPINE_BENCHMAKR_URL = "MIRRORS.txt"

// https://mirrors.alpinelinux.org/ 2022.11.19
var BUILDIN_ALPINE_MIRRORS = []UrlWithAlias{
	{URL: "mirrors.tuna.tsinghua.edu.cn/alpine/", Alias: "cn:tsinghua", Http: true, Https: true, Official: true, Bandwidth: 10000},
	{URL: "mirrors.ustc.edu.cn/alpine/", Alias: "cn:ustc", Http: true, Https: true, Official: true, Bandwidth: 700},
	{URL: "mirrors.nju.edu.cn/alpine/", Alias: "cn:nju", Http: true, Https: true, Official: true, Bandwidth: 10000},
	{URL: "mirror.lzu.edu.cn/alpine/", Alias: "cn:lzu", Http: true, Https: true, Official: true, Bandwidth: 100},
	{URL: "mirrors.sjtug.sjtu.edu.cn/alpine/", Alias: "cn:sjtug", Http: true, Https: true, Official: true, Bandwidth: 500},
	{URL: "mirrors.aliyun.com/alpine/", Alias: "cn:aliyun", Http: true, Https: true, Official: true, Bandwidth: 10000},
	{URL: "mirrors.bfsu.edu.cn/centos/", Alias: "cn:bfsu", Http: true, Https: true, Official: true, Bandwidth: 1000},
	{URL: "mirrors.neusoft.edu.cn/centos/", Alias: "cn:neusoft", Http: true, Https: true, Official: true, Bandwidth: 1000},
}

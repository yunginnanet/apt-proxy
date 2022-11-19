package define

import "regexp"

const (
	DEBIAN_BENCHMAKR_URL = "dists/bullseye/main/binary-amd64/Release"
)

var DEBIAN_HOST_PATTERN = regexp.MustCompile(
	`https?://(deb|security|snapshot).debian.org/debian/(.+)$`,
)

// https://www.debian.org/mirror/list 2022.11.19
var BUILDIN_DEBIAN_MIRRORS = []UrlWithAlias{
	{URL: "https://mirrors.tuna.tsinghua.edu.cn/debian/", Alias: "cn:tsinghua", Http: true, Https: true, Official: true},
	{URL: "http://mirrors.ustc.edu.cn/debian/", Alias: "cn:ustc", Http: true, Https: true, Official: true},
	{URL: "https://mirrors.163.com/debian/", Alias: "cn:163", Http: true, Https: true, Official: true},
	{URL: "https://repo.huaweicloud.com/debian/", Alias: "cn:huawei", Http: true, Https: true, Official: true},
	{URL: "https://mirrors.cloud.tencent.com/debian/", Alias: "cn:tencent"},
	{URL: "http://ftp.cn.debian.org/debian/", Alias: "cn:debian", Http: true, Https: true, Official: true},
	{URL: "http://mirror.bjtu.edu.cn/debian/", Alias: "cn:bjtu", Http: true, Https: true, Official: true},
	{URL: "http://mirrors.bfsu.edu.cn/debian/", Alias: "cn:bfsu", Http: true, Https: true, Official: true},
	{URL: "http://mirrors.neusoft.edu.cn/debian/", Alias: "cn:neusoft", Http: true, Https: true, Official: true},
	{URL: "http://mirrors.hit.edu.cn/debian/", Alias: "cn:hit", Http: true, Https: true, Official: false},
	{URL: "https://mirrors.aliyun.com/debian/", Alias: "cn:aliyun", Http: true, Https: true, Official: false},
	{URL: "http://mirror.lzu.edu.cn/debian/", Alias: "cn:lzu", Http: true, Https: true, Official: false},
	{URL: "http://mirror.nju.edu.cn/debian/", Alias: "cn:nju", Http: true, Https: true, Official: false},
}

package define

import "regexp"

const (
	DEBIAN_BENCHMAKR_URL = "dists/bullseye/main/binary-amd64/Release"
)

var DEBIAN_HOST_PATTERN = regexp.MustCompile(
	`https?://(deb|security|snapshot).debian.org/debian/(.+)$`,
)

var BUILDIN_DEBIAN_MIRRORS = []UrlWithAlias{
	{URL: "https://mirrors.tuna.tsinghua.edu.cn/debian/", Alias: "cn:tsinghua"},
	{URL: "http://mirrors.ustc.edu.cn/debian/", Alias: "cn:ustc"},
	{URL: "https://mirrors.163.com/debian/", Alias: "cn:163"},
	{URL: "https://mirrors.aliyun.com/debian/", Alias: "cn:aliyun"},
	{URL: "https://repo.huaweicloud.com/debian/", Alias: "cn:huawei"},
	{URL: "https://mirrors.cloud.tencent.com/debian/", Alias: "cn:tencent"},
	{URL: "http://ftp.cn.debian.org/debian/", Alias: "cn:debian"},
	{URL: "http://mirror.bjtu.edu.cn/debian/", Alias: "cn:bjtu"},
	{URL: "http://mirror.lzu.edu.cn/debian/", Alias: "cn:lzu"},
	{URL: "http://mirror.nju.edu.cn/debian/", Alias: "cn:nju"},
	{URL: "http://mirrors.bfsu.edu.cn/debian/", Alias: "cn:bfsu"},
	{URL: "http://mirrors.hit.edu.cn/debian/", Alias: "cn:hit"},
	{URL: "http://mirrors.neusoft.edu.cn/debian/", Alias: "cn:neusoft"},
}

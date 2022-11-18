package define

import "regexp"

const (
	UBUNTU_GEO_MIRROR_API = "http://mirrors.ubuntu.com/mirrors.txt"
	UBUNTU_BENCHMAKR_URL  = "dists/jammy/main/binary-amd64/Release"
)

var UBUNTU_HOST_PATTERN = regexp.MustCompile(
	`https?://(\w{2}\.)?(security|archive).ubuntu.com/ubuntu/(.+)$`,
)

var BUILDIN_UBUNTU_MIRRORS = []UrlWithAlias{
	{URL: "https://mirrors.tuna.tsinghua.edu.cn/ubuntu/", Alias: "cn:tsinghua"},
	{URL: "http://mirrors.ustc.edu.cn/ubuntu/", Alias: "cn:ustc"},
	{URL: "https://mirrors.163.com/ubuntu/", Alias: "cn:163"},
	{URL: "https://mirrors.aliyun.com/ubuntu/", Alias: "cn:aliyun"},
	// "http://mirrors.huaweicloud.com/repository/ubuntu/"
	{URL: "https://repo.huaweicloud.com/ubuntu/", Alias: "cn:huawei"},
	{URL: "https://mirrors.cloud.tencent.com/ubuntu/", Alias: "cn:tencent"},
	{URL: "http://mirror.dlut.edu.cn/ubuntu/", Alias: "cn:dlut"},
	{URL: "http://mirrors.dgut.edu.cn/ubuntu/", Alias: "cn:dgut"},
	{URL: "http://mirrors.njupt.edu.cn/ubuntu/", Alias: "cn:njupt"},
	{URL: "https://mirrors.hit.edu.cn/ubuntu/", Alias: "cn:hit"},
	{URL: "http://mirrors.yun-idc.com/ubuntu/", Alias: "cn:yun-idc"},
	{URL: "http://ftp.sjtu.edu.cn/ubuntu/", Alias: "cn:sjtu"},
	{URL: "https://mirror.nju.edu.cn/ubuntu/", Alias: "cn:nju"},
	{URL: "https://mirrors.bupt.edu.cn/ubuntu/", Alias: "cn:bupt"},
	{URL: "http://mirrors.skyshe.cn/ubuntu/", Alias: "cn:skyshe"},
	{URL: "http://mirror.lzu.edu.cn/ubuntu/", Alias: "cn:lzu"},
	{URL: "http://mirrors.cn99.com/ubuntu/", Alias: "cn:cn99"},
	{URL: "http://mirrors.cqu.edu.cn/ubuntu/", Alias: "cn:cqu"},
	{URL: "https://mirror.bjtu.edu.cn/ubuntu/", Alias: "cn:bjtu"},
	{URL: "http://mirrors.sohu.com/ubuntu/", Alias: "cn:sohu"},
	{URL: "http://cn.archive.ubuntu.com/ubuntu/", Alias: "cn:ubuntu"},
}

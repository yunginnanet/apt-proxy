package define

import "regexp"

var CENTOS_HOST_PATTERN = regexp.MustCompile(`https?://.+/centos/(.+)$`)

const CENTOS_BENCHMAKR_URL = "TIME"

// https://www.centos.org/download/mirrors/
var BUILDIN_CENTOS_MIRRORS = []UrlWithAlias{
	{URL: "https://mirrors.tuna.tsinghua.edu.cn/centos/", Alias: "cn:tsinghua", Http: true, Https: true, Official: true},
	{URL: "http://mirrors.aliyun.com/centos/", Alias: "cn:aliyun", Http: true, Https: true, Official: true},
	{URL: "https://mirrors.bfsu.edu.cn/centos/", Alias: "cn:bfsu", Http: true, Https: true, Official: true},
	{URL: "https://mirrors.cqu.edu.cn/CentOS/", Alias: "cn:cqu", Http: true, Https: true, Official: true},
	{URL: "https://mirror.nju.edu.cn/centos/", Alias: "cn:nju", Http: true, Https: true, Official: true},
	{URL: "http://mirror.lzu.edu.cn/centos/", Alias: "cn:lzu", Http: true, Https: true, Official: true},
	{URL: "https://mirrors.njupt.edu.cn/centos/", Alias: "cn:njupt", Http: true, Https: true, Official: true},
	{URL: "http://mirrors.163.com/centos/", Alias: "cn:163", Http: true, Https: true, Official: true},
	{URL: "https://mirrors.bupt.edu.cn/centos/", Alias: "cn:bupt", Http: true, Https: true, Official: true},
	{URL: "https://ftp.sjtu.edu.cn/centos/", Alias: "cn:sjtu", Http: true, Https: true, Official: true},
	{URL: "https://mirrors.ustc.edu.cn/centos/", Alias: "cn:ustc", Http: true, Https: true, Official: true},
	// TODO: valid?
	// {URL: "http://mirrors.neusoft.edu.cn/centos/", Alias: "cn:neusoft"},
	{URL: "https://mirrors.huaweicloud.com/centos/", Alias: "cn:huaweicloud", Http: true, Https: true, Official: false},
}

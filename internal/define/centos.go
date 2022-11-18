package define

import "regexp"

var CENTOS_HOST_PATTERN = regexp.MustCompile(`https?://.+/centos/(.+)$`)

const CENTOS_BENCHMAKR_URL = "TIME"

var BUILDIN_CENTOS_MIRRORS = []UrlWithAlias{
	{URL: "https://mirrors.tuna.tsinghua.edu.cn/centos/", Alias: "cn:tsinghua"},
	{URL: "http://mirrors.aliyun.com/centos/", Alias: "cn:aliyun"},
	{URL: "https://mirrors.bfsu.edu.cn/centos/", Alias: "cn:bfsu"},
	{URL: "https://mirrors.cqu.edu.cn/CentOS/", Alias: "cn:cqu"},
	{URL: "http://mirrors.neusoft.edu.cn/centos/", Alias: "cn:neusoft"},
	{URL: "https://mirror.nju.edu.cn/centos/", Alias: "cn:nju"},
	{URL: "https://mirrors.huaweicloud.com/centos/", Alias: "cn:huaweicloud"},
	{URL: "http://mirror.lzu.edu.cn/centos/", Alias: "cn:lzu"},
	{URL: "https://mirrors.njupt.edu.cn/centos/", Alias: "cn:njupt"},
	{URL: "http://mirrors.163.com/centos/", Alias: "cn:163"},
	{URL: "https://mirrors.bupt.edu.cn/centos/", Alias: "cn:bupt"},
	{URL: "https://ftp.sjtu.edu.cn/centos/", Alias: "cn:sjtu"},
	{URL: "https://mirrors.ustc.edu.cn/centos/", Alias: "cn:ustc"},
}

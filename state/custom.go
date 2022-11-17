package state

type buildin_custom_mirror struct {
	url   string
	alias string
}

var BUILDIN_CUSTOM_UBUNTU_MIRRORS = []buildin_custom_mirror{
	{url: "https://mirrors.tuna.tsinghua.edu.cn/ubuntu/", alias: "cn:tsinghua"},
	{url: "http://mirrors.ustc.edu.cn/ubuntu/", alias: "cn:ustc"},
	{url: "https://mirrors.163.com/ubuntu/", alias: "cn:163"},
	{url: "https://mirrors.aliyun.com/ubuntu/", alias: "cn:aliyun"},
	{url: "https://repo.huaweicloud.com/ubuntu/", alias: "cn:huawei"},
	{url: "https://mirrors.cloud.tencent.com/ubuntu/", alias: "cn:tencent"},
}

var BUILDIN_CUSTOM_DEBIAN_MIRRORS = []buildin_custom_mirror{
	{url: "https://mirrors.tuna.tsinghua.edu.cn/debian/", alias: "cn:tsinghua"},
	{url: "http://mirrors.ustc.edu.cn/debian/", alias: "cn:ustc"},
	{url: "https://mirrors.163.com/debian/", alias: "cn:163"},
	{url: "https://mirrors.aliyun.com/debian/", alias: "cn:aliyun"},
	{url: "https://repo.huaweicloud.com/debian/", alias: "cn:huawei"},
	{url: "https://mirrors.cloud.tencent.com/debian/", alias: "cn:tencent"},
}

var BUILDIN_CUSTOM_CENTOS_MIRRORS = []buildin_custom_mirror{
	{url: "http://mirrors.aliyun.com/centos/", alias: "cn:aliyun"},
	{url: "https://mirrors.bfsu.edu.cn/centos/", alias: "cn:bfsu"},
	{url: "https://mirrors.cqu.edu.cn/CentOS/", alias: "cn:cqu"},
	{url: "http://mirrors.neusoft.edu.cn/centos/", alias: "cn:neusoft"},
	{url: "https://mirror.nju.edu.cn/centos/", alias: "cn:nju"},
	{url: "https://mirrors.huaweicloud.com/centos/", alias: "cn:huaweicloud"},
	{url: "http://mirror.lzu.edu.cn/centos/", alias: "cn:lzu"},
	{url: "https://mirrors.njupt.edu.cn/centos/", alias: "cn:njupt"},
	{url: "http://mirrors.163.com/centos/", alias: "cn:163"},
	{url: "https://mirrors.bupt.edu.cn/centos/", alias: "cn:bupt"},
	{url: "https://ftp.sjtu.edu.cn/centos/", alias: "cn:sjtu"},
	{url: "https://mirrors.tuna.tsinghua.edu.cn/centos/", alias: "cn:tsinghua"},
	{url: "https://mirrors.ustc.edu.cn/centos/", alias: "cn:ustc"},
}

func getUbuntuMirrorByAliases(alias string) string {
	for _, mirror := range BUILDIN_CUSTOM_UBUNTU_MIRRORS {
		if mirror.alias == alias {
			return mirror.url
		}
	}
	return ""
}

func getDebianMirrorByAliases(alias string) string {
	for _, mirror := range BUILDIN_CUSTOM_DEBIAN_MIRRORS {
		if mirror.alias == alias {
			return mirror.url
		}
	}
	return ""
}

func getCentOSMirrorByAliases(alias string) string {
	for _, mirror := range BUILDIN_CUSTOM_CENTOS_MIRRORS {
		if mirror.alias == alias {
			return mirror.url
		}
	}
	return ""
}

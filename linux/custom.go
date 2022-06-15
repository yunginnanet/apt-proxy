package linux

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

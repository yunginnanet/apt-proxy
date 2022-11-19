package mirrors

import (
	"strings"
	"testing"

	Define "github.com/soulteary/apt-proxy/internal/define"
)

func TestGetUbuntuMirrorByAliases(t *testing.T) {
	alias := GetMirrorURLByAliases(Define.TYPE_LINUX_DISTROS_UBUNTU, "cn:tsinghua")
	if !strings.Contains(alias, "mirrors.tuna.tsinghua.edu.cn/ubuntu/") {
		t.Fatal("Test Get Mirror By Custom Name Failed")
	}

	alias = GetMirrorURLByAliases(Define.TYPE_LINUX_DISTROS_UBUNTU, "cn:not-found")
	if alias != "" {
		t.Fatal("Test Get Mirror By Custom Name Failed")
	}
}

func TestGetDebianMirrorByAliases(t *testing.T) {
	alias := GetMirrorURLByAliases(Define.TYPE_LINUX_DISTROS_DEBIAN, "cn:tsinghua")
	if !strings.Contains(alias, "mirrors.tuna.tsinghua.edu.cn/debian/") {
		t.Fatal("Test Get Mirror By Custom Name Failed")
	}

	alias = GetMirrorURLByAliases(Define.TYPE_LINUX_DISTROS_DEBIAN, "cn:not-found")
	if alias != "" {
		t.Fatal("Test Get Mirror By Custom Name Failed")
	}
}

func TestGetCentOSMirrorByAliases(t *testing.T) {
	alias := GetMirrorURLByAliases(Define.TYPE_LINUX_DISTROS_CENTOS, "cn:tsinghua")
	if !strings.Contains(alias, "mirrors.tuna.tsinghua.edu.cn/centos/") {
		t.Fatal("Test Get Mirror By Custom Name Failed")
	}

	alias = GetMirrorURLByAliases(Define.TYPE_LINUX_DISTROS_CENTOS, "cn:not-found")
	if alias != "" {
		t.Fatal("Test Get Mirror By Custom Name Failed")
	}
}

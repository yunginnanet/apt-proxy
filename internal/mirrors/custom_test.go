package mirrors

import (
	"testing"

	Define "github.com/soulteary/apt-proxy/internal/define"
)

func TestGetUbuntuMirrorByAliases(t *testing.T) {
	alias := GetMirrorURLByAliases(Define.TYPE_LINUX_DISTROS_UBUNTU, "cn:tsinghua")
	if alias != Define.BUILDIN_UBUNTU_MIRRORS[0].URL {
		t.Fatal("Test Get Mirror By Custom Name Failed")
	}

	alias = GetMirrorURLByAliases(Define.TYPE_LINUX_DISTROS_UBUNTU, "cn:not-found")
	if alias != "" {
		t.Fatal("Test Get Mirror By Custom Name Failed")
	}
}

func TestGetDebianMirrorByAliases(t *testing.T) {
	alias := GetMirrorURLByAliases(Define.TYPE_LINUX_DISTROS_DEBIAN, "cn:tsinghua")
	if alias != Define.BUILDIN_DEBIAN_MIRRORS[0].URL {
		t.Fatal("Test Get Mirror By Custom Name Failed")
	}

	alias = GetMirrorURLByAliases(Define.TYPE_LINUX_DISTROS_DEBIAN, "cn:not-found")
	if alias != "" {
		t.Fatal("Test Get Mirror By Custom Name Failed")
	}
}

func TestGetCentOSMirrorByAliases(t *testing.T) {
	alias := GetMirrorURLByAliases(Define.TYPE_LINUX_DISTROS_CENTOS, "cn:tsinghua")
	if alias != Define.BUILDIN_CENTOS_MIRRORS[0].URL {
		t.Fatal("Test Get Mirror By Custom Name Failed", alias)
	}

	alias = GetMirrorURLByAliases(Define.TYPE_LINUX_DISTROS_CENTOS, "cn:not-found")
	if alias != "" {
		t.Fatal("Test Get Mirror By Custom Name Failed")
	}
}

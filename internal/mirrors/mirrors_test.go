package mirrors

import (
	"strings"
	"testing"

	Define "github.com/soulteary/apt-proxy/internal/define"
)

func TestGetUbuntuMirrorByAliases(t *testing.T) {
	alias := GetMirrorURLByAliases(Define.TypeLinuxDistrosUbuntu, "cn:siena")
	if !strings.Contains(alias, "mirror.siena.edu/ubuntu/") {
		t.Fatal("Test Get Mirror By Custom Name Failed")
	}

	alias = GetMirrorURLByAliases(Define.TypeLinuxDistrosUbuntu, "cn:not-found")
	if alias != "" {
		t.Fatal("Test Get Mirror By Custom Name Failed")
	}
}

func TestGetDebianMirrorByAliases(t *testing.T) {
	alias := GetMirrorURLByAliases(Define.TypeLinuxDistrosDebian, "cn:csail")
	if !strings.Contains(alias, "debian.csail.mit.edu/debian/") {
		t.Fatal("Test Get Mirror By Custom Name Failed")
	}

	alias = GetMirrorURLByAliases(Define.TypeLinuxDistrosDebian, "cn:not-found")
	if alias != "" {
		t.Fatal("Test Get Mirror By Custom Name Failed")
	}
}

func TestGetCentOSMirrorByAliases(t *testing.T) {
	alias := GetMirrorURLByAliases(Define.TypeLinuxDistrosCentos, "cn:tsinghua")
	if !strings.Contains(alias, "mirrors.tuna.tsinghua.edu.cn/centos/") {
		t.Fatal("Test Get Mirror By Custom Name Failed")
	}

	alias = GetMirrorURLByAliases(Define.TypeLinuxDistrosCentos, "cn:not-found")
	if alias != "" {
		t.Fatal("Test Get Mirror By Custom Name Failed")
	}
}

func TestGetMirrorUrlsByGeo(t *testing.T) {
	mirrors := GetGeoMirrorUrlsByMode(Define.TypeLinuxAllDistros)
	if len(mirrors) == 0 {
		t.Fatal("No mirrors found")
	}

	mirrors = GetGeoMirrorUrlsByMode(Define.TypeLinuxDistrosDebian)
	if len(mirrors) != len(BuildinDebianMirrors) {
		t.Fatal("Get mirrors error")
	}

	mirrors = GetGeoMirrorUrlsByMode(Define.TypeLinuxDistrosUbuntu)
	if len(mirrors) == 0 {
		t.Fatal("No mirrors found")
	}
}

func TestGetPredefinedConfiguration(t *testing.T) {
	res, pattern := GetPredefinedConfiguration(Define.TypeLinuxDistrosUbuntu)
	if res != Define.UbuntuBenchmakrUrl {
		t.Fatal("Failed to get resource link")
	}
	if !pattern.MatchString("/ubuntu/InRelease") {
		t.Fatal("Failed to verify domain name rules")
	}
	if !pattern.MatchString("/ubuntu/InRelease") {
		t.Fatal("Failed to verify domain name rules")
	}

	res, pattern = GetPredefinedConfiguration(Define.TypeLinuxDistrosDebian)
	if res != Define.DebianBenchmakrUrl {
		t.Fatal("Failed to get resource link")
	}
	if !pattern.MatchString("/debian/InRelease") {
		t.Fatal("Failed to verify domain name rules")
	}

	res, pattern = GetPredefinedConfiguration(Define.TypeLinuxDistrosCentos)
	if res != Define.CentosBenchmakrUrl {
		t.Fatal("Failed to get resource link")
	}
	if !pattern.MatchString("/centos/test/repomd.xml") {
		t.Fatal("Failed to verify domain name rules")
	}

	res, pattern = GetPredefinedConfiguration(Define.TypeLinuxDistrosAlpine)
	if res != Define.AlpineBenchmakrUrl {
		t.Fatal("Failed to get resource link")
	}
	if !pattern.MatchString("/alpine/test/APKINDEX.tar.gz") {
		t.Fatal("Failed to verify domain name rules")
	}
}

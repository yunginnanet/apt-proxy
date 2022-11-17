package mirrors

import "testing"

func TestGetUbuntuMirrorByAliases(t *testing.T) {
	alias := GetUbuntuMirrorByAliases("cn:tsinghua")
	if alias != BUILDIN_CUSTOM_UBUNTU_MIRRORS[0].url {
		t.Fatal("Test Get Mirror By Custom Name Failed")
	}

	alias = GetUbuntuMirrorByAliases("cn:not-found")
	if alias != "" {
		t.Fatal("Test Get Mirror By Custom Name Failed")
	}
}

func TestGetDebianMirrorByAliases(t *testing.T) {
	alias := GetDebianMirrorByAliases("cn:tsinghua")
	if alias != BUILDIN_CUSTOM_DEBIAN_MIRRORS[0].url {
		t.Fatal("Test Get Mirror By Custom Name Failed")
	}

	alias = GetDebianMirrorByAliases("cn:not-found")
	if alias != "" {
		t.Fatal("Test Get Mirror By Custom Name Failed")
	}
}

func TestGetCentOSMirrorByAliases(t *testing.T) {
	alias := GetDebianMirrorByAliases("cn:tsinghua")
	if alias != BUILDIN_CUSTOM_DEBIAN_MIRRORS[0].url {
		t.Fatal("Test Get Mirror By Custom Name Failed")
	}

	alias = GetDebianMirrorByAliases("cn:not-found")
	if alias != "" {
		t.Fatal("Test Get Mirror By Custom Name Failed")
	}
}

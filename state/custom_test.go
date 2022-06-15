package state

import "testing"

func TestGetUbuntuMirrorByAliases(t *testing.T) {
	alias := getUbuntuMirrorByAliases("cn:tsinghua")
	if alias != BUILDIN_CUSTOM_UBUNTU_MIRRORS[0].url {
		t.Fatal("Test Get Mirror By Custom Name Failed")
	}

	alias = getUbuntuMirrorByAliases("cn:not-found")
	if alias != "" {
		t.Fatal("Test Get Mirror By Custom Name Failed")
	}
}

func TestGetDebianMirrorByAliases(t *testing.T) {
	alias := getDebianMirrorByAliases("cn:tsinghua")
	if alias != BUILDIN_CUSTOM_DEBIAN_MIRRORS[0].url {
		t.Fatal("Test Get Mirror By Custom Name Failed")
	}

	alias = getDebianMirrorByAliases("cn:not-found")
	if alias != "" {
		t.Fatal("Test Get Mirror By Custom Name Failed")
	}
}

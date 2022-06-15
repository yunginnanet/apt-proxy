package linux

import (
	"testing"
)

func TestGetGeoMirrors(t *testing.T) {
	mirrors, err := getUbuntuMirrorUrlsByGeo()
	if err != nil {
		t.Fatal(err)
	}

	if len(mirrors) == 0 {
		t.Fatal("No mirrors found")
	}
}

func TestGetMirrorUrlsByGeo(t *testing.T) {
	mirrors := getGeoMirrorUrlsByMode(TYPE_LINUX_ALL_DISTROS)
	if len(mirrors) == 0 {
		t.Fatal("No mirrors found")
	}

	mirrors = getGeoMirrorUrlsByMode(TYPE_LINUX_DISTROS_DEBIAN)
	if len(mirrors) != len(BUILDIN_OFFICAL_DEBIAN_MIRRORS) {
		t.Fatal("Get mirrors error")
	}

	mirrors = getGeoMirrorUrlsByMode(TYPE_LINUX_DISTROS_UBUNTU)
	if len(mirrors) == 0 {
		t.Fatal("No mirrors found")
	}
}

func TestGetPredefinedConfiguration(t *testing.T) {
	res, pattern := getPredefinedConfiguration(TYPE_LINUX_DISTROS_UBUNTU)
	if res != UBUNTU_BENCHMAKR_URL {
		t.Fatal("Failed to get resource link")
	}
	if !pattern.MatchString("http://archive.ubuntu.com/ubuntu/InRelease") {
		t.Fatal("Failed to verify domain name rules")
	}

	res, pattern = getPredefinedConfiguration(TYPE_LINUX_DISTROS_DEBIAN)
	if res != DEBIAN_BENCHMAKR_URL {
		t.Fatal("Failed to get resource link")
	}
	if !pattern.MatchString("http://deb.debian.org/debian/InRelease") {
		t.Fatal("Failed to verify domain name rules")
	}
}

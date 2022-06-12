package linux

import (
	"testing"
)

func TestGetGeoMirrors(t *testing.T) {
	mirrors, err := getGeoMirrors(UBUNTU_MIRROR_URLS)
	if err != nil {
		t.Fatal(err)
	}

	if len(mirrors.URLs) == 0 {
		t.Fatal("No mirrors found")
	}
}

func TestGetPredefinedConfiguration(t *testing.T) {
	url, res, pattern := getPredefinedConfiguration(LINUX_DISTROS_UBUNTU)
	if url != UBUNTU_MIRROR_URLS || res != UBUNTU_BENCHMAKR_URL {
		t.Fatal("Failed to get resource link")
	}
	if !pattern.MatchString("http://archive.ubuntu.com/ubuntu/InRelease") {
		t.Fatal("Failed to verify domain name rules")
	}

	url, res, pattern = getPredefinedConfiguration(LINUX_DISTROS_DEBIAN)
	if url != "" || res != DEBIAN_BENCHMAKR_URL {
		t.Fatal("Failed to get resource link")
	}
	if !pattern.MatchString("http://deb.debian.org/debian/InRelease") {
		t.Fatal("Failed to verify domain name rules")
	}

}

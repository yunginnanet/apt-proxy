package linux_test

import (
	"testing"

	"github.com/soulteary/apt-proxy/linux"
)

func TestGetGeoMirrors(t *testing.T) {
	mirrors, err := linux.GetGeoMirrors(linux.UBUNTU_MIRROR_URLS)
	if err != nil {
		t.Fatal(err)
	}

	if len(mirrors.URLs) == 0 {
		t.Fatal("No mirrors found")
	}
}

func TestGetMirrorsUrlAndBenchmarkUrl(t *testing.T) {
	url, res, pattern := linux.GetPredefinedConfiguration(linux.UBUNTU)
	if url != linux.UBUNTU_MIRROR_URLS || res != linux.UBUNTU_BENCHMAKR_URL {
		t.Fatal("Failed to get resource link")
	}

	if !pattern.MatchString("http://archive.ubuntu.com/ubuntu/InRelease") {
		t.Fatal("Failed to verify domain name rules")
	}
}

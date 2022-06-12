package linux

import (
	"log"
	"testing"
)

func TestBenchmark(t *testing.T) {
	_, err := resourceBenchmark(UBUNTU_MIRROR_URLS, "", benchmarkTimes)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMirrorsBenchmark(t *testing.T) {
	mirrors, err := getGeoMirrors(UBUNTU_MIRROR_URLS)
	if err != nil {
		t.Fatal(err)
	}

	mirror, err := fastest(mirrors, UBUNTU_BENCHMAKR_URL)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("Fastest mirror is %s", mirror)
}

package linux

import (
	"log"
	"testing"
)

func TestResourceBenchmark(t *testing.T) {
	_, err := benchmark(UBUNTU_MIRROR_URLS, "", BENCHMARK_MAX_TRIES)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMirrorsBenchmark(t *testing.T) {
	mirrors, err := getGeoMirrors(UBUNTU_MIRROR_URLS)
	if err != nil {
		t.Fatal(err)
	}

	mirror, err := getTheFastestMirror(mirrors, UBUNTU_BENCHMAKR_URL)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("Fastest mirror is %s", mirror)
}

package linux_test

import (
	"log"
	"testing"

	"github.com/soulteary/apt-proxy/linux"
)

func TestBenchmark(t *testing.T) {
	_, err := linux.Benchmark(linux.UBUNTU_MIRROR_URLS, "", linux.BenchmarkTimes)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMirrorsBenchmark(t *testing.T) {
	mirrors, err := linux.GetGeoMirrors(linux.UBUNTU_MIRROR_URLS)
	if err != nil {
		t.Fatal(err)
	}

	fastest, err := linux.Fastest(mirrors, linux.UBUNTU_BENCHMAKR_URL)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("Fastest mirror is %s", fastest)
}

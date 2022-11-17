package benchmarks_test

import (
	"log"
	"testing"

	Benchmark "github.com/soulteary/apt-proxy/internal/benchmark"
	Mirrors "github.com/soulteary/apt-proxy/internal/mirrors"
)

func TestResourceBenchmark(t *testing.T) {
	const resourcePath = ""
	_, err := Benchmark.Benchmark(Mirrors.UBUNTU_GEO_MIRROR_API, resourcePath, Benchmark.BENCHMARK_MAX_TRIES)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMirrorsBenchmark(t *testing.T) {
	mirrors := Mirrors.GetGeoMirrorUrlsByMode(Mirrors.TYPE_LINUX_DISTROS_UBUNTU)
	mirror, err := Benchmark.GetTheFastestMirror(mirrors, Mirrors.UBUNTU_BENCHMAKR_URL)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("Fastest mirror is %s", mirror)
}

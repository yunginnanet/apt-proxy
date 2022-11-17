package mirrors

import (
	"log"
	"testing"
)

func TestResourceBenchmark(t *testing.T) {
	const resourcePath = ""
	_, err := benchmark(UBUNTU_GEO_MIRROR_API, resourcePath, BENCHMARK_MAX_TRIES)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMirrorsBenchmark(t *testing.T) {
	mirrors := GetGeoMirrorUrlsByMode(TYPE_LINUX_DISTROS_UBUNTU)
	mirror, err := GetTheFastestMirror(mirrors, UBUNTU_BENCHMAKR_URL)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("Fastest mirror is %s", mirror)
}

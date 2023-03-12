package benchmarks_test

import (
	"errors"
	"testing"
	"time"

	Benchmark "github.com/soulteary/apt-proxy/internal/benchmark"
	Define "github.com/soulteary/apt-proxy/internal/define"
	Mirrors "github.com/soulteary/apt-proxy/internal/mirrors"
)

func TestBenchmark(t *testing.T) {
	const resourcePath = ""
	_, err := Benchmark.Benchmark(Define.UbuntuGeoMirrorApi, resourcePath, Benchmark.BenchmarkMaxTries)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetTheFastestMirror(t *testing.T) {
	mirrors := Mirrors.GetGeoMirrorUrlsByMode(Define.TypeLinuxDistrosUbuntu)
	_, err := Benchmark.GetTheFastestMirror(mirrors, Define.UbuntuBenchmakrUrl)
	if err != nil {
		t.Fatal(err)
	}

	// test empty list
	var emptyList []string
	_, err = Benchmark.GetTheFastestMirror(emptyList, "")
	if err == nil {
		t.Fatal("not cache an empty error")
	}
}

func TestReadResults(t *testing.T) {
	ch := make(chan Benchmark.Result)
	for i := 0; i < 5; i++ {
		go func(t int) {
			ch <- Benchmark.Result{"code", time.Duration(t)}
		}(i)
	}
	results, err := Benchmark.ReadResults(ch, 2)

	if len(results) == 0 {
		t.Fatal(errors.New("No results found: " + err.Error()))
	} else if err != nil {
		t.Fatalf("Error benchmarking mirrors: %s", err.Error())
	}
}

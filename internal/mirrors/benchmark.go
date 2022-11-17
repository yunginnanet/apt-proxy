package mirrors

import (
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

func benchmark(base string, query string, times int) (time.Duration, error) {
	var sum int64

	timeout := time.Duration(BENCHMARK_MAX_TIMEOUT * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	for i := 0; i < times; i++ {
		timer := time.Now()
		response, err := client.Get(base + query)
		if err != nil {
			return timeout, err
		}

		defer response.Body.Close()
		_, err = io.ReadAll(response.Body)
		if err != nil {
			return timeout, err
		}

		sum = sum + int64(time.Since(timer))
	}

	return time.Duration(sum / int64(times)), nil
}

type benchmarkResult struct {
	URL      string
	Duration time.Duration
}

func GetTheFastestMirror(mirrors []string, testUrl string) (string, error) {
	ch := make(chan benchmarkResult)
	log.Printf("Start benchmarking mirrors")
	// kick off all benchmarks in parallel
	for _, url := range mirrors {
		go func(u string) {
			duration, err := benchmark(u, testUrl, BENCHMARK_MAX_TRIES)
			if err == nil {
				ch <- benchmarkResult{u, duration}
			}
		}(url)
	}

	maxTries := len(mirrors)
	if maxTries > 3 {
		maxTries = 3
	}

	// wait for the fastest results to come back
	results, err := readResults(ch, maxTries)
	log.Printf("Finished benchmarking mirrors")
	if len(results) == 0 {
		return "", errors.New("No results found: " + err.Error())
	} else if err != nil {
		log.Printf("Error benchmarking mirrors: %s", err.Error())
	}

	return results[0].URL, nil
}

func readResults(ch <-chan benchmarkResult, size int) (br []benchmarkResult, err error) {
	for {
		select {
		case r := <-ch:
			br = append(br, r)
			if len(br) >= size {
				return br, nil
			}
		case <-time.After(BENCHMARK_DETECT_TIMEOUT * time.Second):
			return br, errors.New("timed out waiting for results")
		}
	}
}

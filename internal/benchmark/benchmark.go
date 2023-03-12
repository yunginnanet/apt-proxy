package benchmarks

import (
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	BenchmarkMaxTimeout    = 15 // 15 seconds, detect resource timeout
	BenchmarkMaxTries      = 3  // times, maximum number of attempts
	BenchmarkDetectTimeout = 3  // 3 seconds, for select fast mirror
)

type Result struct {
	URL      string
	Duration time.Duration
}

func Benchmark(base string, query string, times int) (time.Duration, error) {
	var sum int64

	timeout := time.Duration(BenchmarkMaxTimeout * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	for i := 0; i < times; i++ {
		timer := time.Now()
		response, err := client.Get(base + query)
		if err != nil {
			return timeout, err
		}

		if response.StatusCode == http.StatusNotFound {
			log.Printf("Benchmark: %s%s not found", base, query)
			break
		}

		_, err = io.ReadAll(response.Body)
		if err != nil {
			if response.Body != nil {
				_ = response.Body.Close()
			}
			return timeout, err
		}
		_ = response.Body.Close()
		sum = sum + int64(time.Since(timer))
	}

	return time.Duration(sum / int64(times)), nil
}

func GetTheFastestMirror(mirrors []string, testUrl string) (string, error) {
	ch := make(chan Result)
	log.Printf("Start benchmarking mirrors")
	// kick off all benchmarks in parallel
	for _, url := range mirrors {
		go func(u string) {
			duration, err := Benchmark(u, testUrl, BenchmarkMaxTries)
			if err == nil {
				ch <- Result{u, duration}
			}
		}(url)
	}

	maxTries := len(mirrors)
	if maxTries > 3 {
		maxTries = 3
	}

	// wait for the fastest results to come back
	results, err := ReadResults(ch, maxTries)
	log.Printf("Finished benchmarking mirrors")
	if len(results) == 0 {
		return "", errors.Join(errors.New("no results found"), err)
	} else if err != nil {
		log.Printf("Error benchmarking mirrors: %s", err.Error())
	}

	return results[0].URL, nil
}

func ReadResults(ch <-chan Result, size int) (br []Result, err error) {
	for {
		select {
		case r := <-ch:
			br = append(br, r)
			if len(br) >= size {
				return br, nil
			}
		case <-time.After(BenchmarkDetectTimeout * time.Second):
			return br, errors.New("timed out waiting for results")
		}
	}
}

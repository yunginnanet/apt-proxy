package cli

import (
	"log"
	"net/http"

	"github.com/soulteary/apt-proxy/pkgs/httpcache"
	"github.com/soulteary/apt-proxy/pkgs/httplog"
	"github.com/soulteary/apt-proxy/proxy"
)

type AppFlags struct {
	Debug    bool
	Version  string
	CacheDir string
	Mirror   string
	Types    string
	Listen   string
}

func Parse(appFlags AppFlags) {
	log.Printf("running apt-proxy %s", appFlags.Version)

	if appFlags.Debug {
		log.Printf("enable debug: true")
		httpcache.DebugLogging = true
	}

	cache, err := httpcache.NewDiskCache(appFlags.CacheDir)
	if err != nil {
		log.Fatal(err)
	}

	ap := proxy.NewAptProxyFromDefaults(appFlags.Mirror, appFlags.Types)
	ap.Handler = httpcache.NewHandler(cache, ap.Handler)

	logger := httplog.NewResponseLogger(ap.Handler)
	logger.DumpRequests = appFlags.Debug
	logger.DumpResponses = appFlags.Debug
	logger.DumpErrors = appFlags.Debug
	ap.Handler = logger

	log.Printf("proxy listening on %s", appFlags.Listen)
	log.Fatal(http.ListenAndServe(appFlags.Listen, ap))
}

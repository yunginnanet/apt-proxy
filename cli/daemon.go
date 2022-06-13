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
	Mode     string
	Listen   string
}

func initStore(appFlags AppFlags) (cache httpcache.Cache, err error) {
	cache, err = httpcache.NewDiskCache(appFlags.CacheDir)
	if err != nil {
		return cache, err
	}
	return cache, nil
}

func initProxy(appFlags AppFlags, cache httpcache.Cache) (ap *proxy.AptProxy) {
	ap = proxy.CreateAptProxyRouter(appFlags.Mirror, appFlags.Mode)
	ap.Handler = httpcache.NewHandler(cache, ap.Handler)
	return ap
}

func initLogger(appFlags AppFlags, ap *proxy.AptProxy) {
	if appFlags.Debug {
		log.Printf("enable debug: true")
		httpcache.DebugLogging = true
	}
	logger := httplog.NewResponseLogger(ap.Handler)
	logger.DumpRequests = appFlags.Debug
	logger.DumpResponses = appFlags.Debug
	logger.DumpErrors = appFlags.Debug
	ap.Handler = logger
}

func StartServer(appFlags *AppFlags, ap *proxy.AptProxy) {
	log.Printf("proxy listening on %s", appFlags.Listen)
	log.Fatal(http.ListenAndServe(appFlags.Listen, ap))
}

func Daemon(appFlags *AppFlags) {
	log.Printf("running apt-proxy %s", appFlags.Version)

	cache, err := initStore(*appFlags)
	if err != nil {
		log.Fatal(err)
	}

	ap := initProxy(*appFlags, cache)

	initLogger(*appFlags, ap)

	StartServer(&*appFlags, ap)
}

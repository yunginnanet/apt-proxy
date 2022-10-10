package cli

import (
	"log"
	"net/http"
	"time"

	"github.com/soulteary/apt-proxy/pkg/httpcache"
	"github.com/soulteary/apt-proxy/pkg/httplog"
	"github.com/soulteary/apt-proxy/server"
)

type AppFlags struct {
	Debug    bool
	Version  string
	CacheDir string
	Mode     int
	Listen   string

	// mirror
	Ubuntu string
	Debian string
}

func initStore(appFlags AppFlags) (cache httpcache.Cache, err error) {
	cache, err = httpcache.NewDiskCache(appFlags.CacheDir)
	if err != nil {
		return cache, err
	}
	return cache, nil
}

func initProxy(appFlags AppFlags, cache httpcache.Cache) (ap *server.AptProxy) {
	ap = server.CreateAptProxyRouter()
	ap.Handler = httpcache.NewHandler(cache, ap.Handler)
	return ap
}

func initLogger(appFlags AppFlags, ap *server.AptProxy) {
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

func StartServer(appFlags *AppFlags, ap *server.AptProxy) {
	log.Printf("proxy listening on %s", appFlags.Listen)

	server := &http.Server{
		Addr:              appFlags.Listen,
		Handler:           ap,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       5 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
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

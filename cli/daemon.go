package cli

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/soulteary/apt-proxy/internal/server"
	"github.com/soulteary/apt-proxy/pkg/httpcache"
	"github.com/soulteary/apt-proxy/pkg/httplog"
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
	CentOS string
	Alpine string
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
	// graceful shutdown
	// https://github.com/soulteary/flare/blob/main/cmd/daemon.go
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:              appFlags.Listen,
		Handler:           ap,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       5 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("program startup error: %s\n", err)
		}
	}()
	log.Println("Program has been started 🚀")

	<-ctx.Done()

	stop()
	log.Println("The program is closing, to end immediately press CTRL+C")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("program force close: ", err)
	}
}

func Daemon(appFlags *AppFlags) {
	log.Printf("running apt-proxy %s", appFlags.Version)

	cache, err := initStore(*appFlags)
	if err != nil {
		log.Fatal(err)
	}

	ap := initProxy(*appFlags, cache)

	initLogger(*appFlags, ap)

	StartServer(appFlags, ap)
}

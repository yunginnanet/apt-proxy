package main

import (
	"flag"

	"github.com/soulteary/apt-proxy/cli"
	"github.com/soulteary/apt-proxy/linux"
)

const (
	DEFAULT_HOST      = "0.0.0.0"
	DEFAULT_PORT      = "3142"
	DEFAULT_CACHE_DIR = "./.aptcache"
	DEFAULT_MIRROR    = "" // "https://mirrors.tuna.tsinghua.edu.cn/ubuntu/"
	DEFAULT_TYPE      = linux.UBUNTU
	DEFAULT_DEBUG     = false
)

var version string

func parseFlags() (appFlags cli.AppFlags) {
	var (
		host  string
		port  string
		types string
	)
	flag.StringVar(&host, "host", DEFAULT_HOST, "the host to bind to")
	flag.StringVar(&port, "port", DEFAULT_PORT, "the port to bind to")
	flag.StringVar(&types, "type", DEFAULT_TYPE, "select the type of system to cache: ubuntu/debian")
	flag.BoolVar(&appFlags.Debug, "debug", DEFAULT_DEBUG, "whether to output debugging logging")
	flag.StringVar(&appFlags.CacheDir, "cachedir", DEFAULT_CACHE_DIR, "the dir to store cache data in")
	flag.StringVar(&appFlags.Mirror, "mirror", DEFAULT_MIRROR, "the mirror for fetching packages")
	flag.Parse()

	if types != linux.UBUNTU && types != linux.DEBIAN {
		types = linux.UBUNTU
	}

	appFlags.Types = types
	appFlags.Listen = host + ":" + port
	appFlags.Version = version

	return appFlags
}

func main() {
	cli.Parse(parseFlags())
}

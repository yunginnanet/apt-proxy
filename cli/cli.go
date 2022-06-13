package cli

import (
	"flag"

	"github.com/soulteary/apt-proxy/linux"
)

const (
	DEFAULT_HOST      = "0.0.0.0"
	DEFAULT_PORT      = "3142"
	DEFAULT_CACHE_DIR = "./.aptcache"
	DEFAULT_MIRROR    = "" // "https://mirrors.tuna.tsinghua.edu.cn/ubuntu/"
	DEFAULT_TYPE      = linux.LINUX_ALL_DISTROS
	DEFAULT_DEBUG     = false
)

var Version string

func getProxyMode(mode string) string {
	if mode != linux.LINUX_ALL_DISTROS &&
		mode != linux.LINUX_DISTROS_DEBIAN &&
		mode != linux.LINUX_DISTROS_UBUNTU {
		return linux.LINUX_ALL_DISTROS
	}
	return mode
}

func ParseFlags() (appFlags AppFlags) {
	var (
		host string
		port string
		mode string
	)
	flag.StringVar(&host, "host", DEFAULT_HOST, "the host to bind to")
	flag.StringVar(&port, "port", DEFAULT_PORT, "the port to bind to")
	flag.StringVar(&mode, "mode", DEFAULT_TYPE, "select the mode of system to cache: `all` / `ubuntu` / `debian`")
	flag.BoolVar(&appFlags.Debug, "debug", DEFAULT_DEBUG, "whether to output debugging logging")
	flag.StringVar(&appFlags.CacheDir, "cachedir", DEFAULT_CACHE_DIR, "the dir to store cache data in")
	flag.StringVar(&appFlags.Mirror, "mirror", DEFAULT_MIRROR, "the mirror for fetching packages")
	flag.Parse()

	appFlags.Mode = getProxyMode(mode)
	appFlags.Listen = host + ":" + port
	appFlags.Version = Version

	return appFlags
}

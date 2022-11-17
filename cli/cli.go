package cli

import (
	"flag"

	Mirrors "github.com/soulteary/apt-proxy/internal/mirrors"
	"github.com/soulteary/apt-proxy/state"
)

const (
	DEFAULT_HOST          = "0.0.0.0"
	DEFAULT_PORT          = "3142"
	DEFAULT_CACHE_DIR     = "./.aptcache"
	DEFAULT_UBUNTU_MIRROR = "" // "https://mirrors.tuna.tsinghua.edu.cn/ubuntu/"
	DEFAULT_DEBIAN_MIRROR = "" // "https://mirrors.tuna.tsinghua.edu.cn/debian/"
	DEFAULT_MODE_NAME     = Mirrors.LINUX_ALL_DISTROS
	DEFAULT_DEBUG         = false
)

var Version string

func getProxyMode(mode string) int {
	if mode == Mirrors.LINUX_DISTROS_UBUNTU {
		return Mirrors.TYPE_LINUX_DISTROS_UBUNTU
	}

	if mode == Mirrors.LINUX_DISTROS_DEBIAN {
		return Mirrors.TYPE_LINUX_DISTROS_DEBIAN
	}

	return Mirrors.TYPE_LINUX_ALL_DISTROS
}

func ParseFlags() (appFlags AppFlags) {
	var (
		host     string
		port     string
		userMode string
	)
	flag.StringVar(&host, "host", DEFAULT_HOST, "the host to bind to")
	flag.StringVar(&port, "port", DEFAULT_PORT, "the port to bind to")
	flag.StringVar(&userMode, "mode", DEFAULT_MODE_NAME, "select the mode of system to cache: `all` / `ubuntu` / `debian`")
	flag.BoolVar(&appFlags.Debug, "debug", DEFAULT_DEBUG, "whether to output debugging logging")
	flag.StringVar(&appFlags.CacheDir, "cachedir", DEFAULT_CACHE_DIR, "the dir to store cache data in")
	flag.StringVar(&appFlags.Ubuntu, "ubuntu", DEFAULT_UBUNTU_MIRROR, "the ubuntu mirror for fetching packages")
	flag.StringVar(&appFlags.Debian, "debian", DEFAULT_DEBIAN_MIRROR, "the debian mirror for fetching packages")
	flag.Parse()

	mode := getProxyMode(userMode)

	appFlags.Mode = mode
	appFlags.Listen = host + ":" + port
	appFlags.Version = Version

	state.SetProxyMode(mode)
	state.SetDebianMirror(appFlags.Debian)
	state.SetUbuntuMirror(appFlags.Ubuntu)

	return appFlags
}

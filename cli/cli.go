package cli

import (
	"flag"

	Define "github.com/soulteary/apt-proxy/internal/define"
	State "github.com/soulteary/apt-proxy/internal/state"
)

const (
	DefaultHost         = "0.0.0.0"
	DefaultPort         = "3142"
	DefaultCacheDir     = "./.aptcache"
	DefaultUbuntuMirror = "" // "https://mirrors.tuna.tsinghua.edu.cn/ubuntu/"
	DefaultDebianMirror = "" // "https://mirrors.tuna.tsinghua.edu.cn/debian/"
	DefaultCentosMirror = "" // "https://mirrors.tuna.tsinghua.edu.cn/centos/"
	DefaultAlpineMirror = "" // "https://mirrors.tuna.tsinghua.edu.cn/alpine/"
	DefaultModeName     = Define.LinuxAllDistros
	DefaultDebug        = false
)

var Version string

func getProxyMode(mode string) int {
	if mode == Define.LinuxDistrosUbuntu {
		return Define.TypeLinuxDistrosUbuntu
	}

	if mode == Define.LinuxDistrosDebian {
		return Define.TypeLinuxDistrosDebian
	}

	if mode == Define.LinuxDistrosCentos {
		return Define.TypeLinuxDistrosCentos
	}

	if mode == Define.LinuxDistrosAlpine {
		return Define.TypeLinuxDistrosAlpine
	}

	return Define.TypeLinuxAllDistros
}

func ParseFlags() (appFlags AppFlags) {
	var (
		host     string
		port     string
		userMode string
	)
	flag.StringVar(&host, "host", DefaultHost, "the host to bind to")
	flag.StringVar(&port, "port", DefaultPort, "the port to bind to")
	flag.StringVar(&userMode, "mode", DefaultModeName, "select the mode of system to cache: `all` / `ubuntu` / `debian` / `centos` / `alpine`")
	flag.BoolVar(&appFlags.Debug, "debug", DefaultDebug, "whether to output debugging logging")
	flag.StringVar(&appFlags.CacheDir, "cachedir", DefaultCacheDir, "the dir to store cache data in")
	flag.StringVar(&appFlags.Ubuntu, "ubuntu", DefaultUbuntuMirror, "the ubuntu mirror for fetching packages")
	flag.StringVar(&appFlags.Debian, "debian", DefaultDebianMirror, "the debian mirror for fetching packages")
	flag.StringVar(&appFlags.CentOS, "centos", DefaultCentosMirror, "the centos mirror for fetching packages")
	flag.StringVar(&appFlags.Alpine, "alpine", DefaultAlpineMirror, "the alpine mirror for fetching packages")
	flag.Parse()

	mode := getProxyMode(userMode)

	appFlags.Mode = mode
	appFlags.Listen = host + ":" + port
	appFlags.Version = Version

	State.SetProxyMode(mode)
	State.SetUbuntuMirror(appFlags.Ubuntu)
	State.SetDebianMirror(appFlags.Debian)
	State.SetCentOSMirror(appFlags.CentOS)
	State.SetAlpineMirror(appFlags.Alpine)

	return appFlags
}

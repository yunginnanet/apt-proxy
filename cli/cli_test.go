package cli

import (
	"os"
	"testing"

	Define "github.com/soulteary/apt-proxy/internal/define"
)

func TestGetProxyMode(t *testing.T) {
	if getProxyMode("not-support-os") != Define.TypeLinuxAllDistros {
		t.Fatal("Incorrect return default value")
	}

	if getProxyMode(Define.LinuxDistrosUbuntu) != Define.TypeLinuxDistrosUbuntu {
		t.Fatal("Incorrect return value")
	}

	if getProxyMode(Define.LinuxDistrosDebian) != Define.TypeLinuxDistrosDebian {
		t.Fatal("Incorrect return value")
	}

	if getProxyMode(Define.LinuxDistrosCentos) != Define.TypeLinuxDistrosCentos {
		t.Fatal("Incorrect return value")
	}

	if getProxyMode(Define.LinuxDistrosAlpine) != Define.TypeLinuxDistrosAlpine {
		t.Fatal("Incorrect return value")
	}
}

func TestParseFlagsAndDaemonInit(t *testing.T) {
	os.Args = append(os.Args, "--mode=not-support-os")
	flags := ParseFlags()

	if flags.Debug != DefaultDebug {
		t.Fatal("Default option `Debug` value mismatch")
	}

	if flags.Listen != (DefaultHost + ":" + DefaultPort) {
		t.Fatal("Default option `Listen` value mismatch")
	}

	if flags.Mode != getProxyMode(DefaultModeName) {
		t.Fatal("Default option `Mode` value mismatch")
	}

	if flags.Ubuntu != DefaultUbuntuMirror {
		t.Fatal("Default option `Ubuntu` value mismatch")
	}

	if flags.Debian != DefaultDebianMirror {
		t.Fatal("Default option `Debian` value mismatch")
	}

	if flags.CacheDir != DefaultCacheDir {
		t.Fatal("Default option `CacheDir` value mismatch")
	}

	cache, err := initStore(flags)
	if err != nil {
		t.Fatal("Init Store Failed")
	}

	ap := initProxy(flags, cache)
	initLogger(flags, ap)
}

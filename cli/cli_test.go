package cli

import (
	"os"
	"testing"

	Define "github.com/soulteary/apt-proxy/internal/define"
)

func TestGetProxyMode(t *testing.T) {
	if getProxyMode("not-support-os") != Define.TYPE_LINUX_ALL_DISTROS {
		t.Fatal("Incorrect return default value")
	}

	if getProxyMode(Define.LINUX_DISTROS_UBUNTU) != Define.TYPE_LINUX_DISTROS_UBUNTU {
		t.Fatal("Incorrect return value")
	}

	if getProxyMode(Define.LINUX_DISTROS_DEBIAN) != Define.TYPE_LINUX_DISTROS_DEBIAN {
		t.Fatal("Incorrect return value")
	}

	if getProxyMode(Define.LINUX_DISTROS_CENTOS) != Define.TYPE_LINUX_DISTROS_CENTOS {
		t.Fatal("Incorrect return value")
	}

	if getProxyMode(Define.LINUX_DISTROS_ALPINE) != Define.TYPE_LINUX_DISTROS_ALPINE {
		t.Fatal("Incorrect return value")
	}
}

func TestParseFlagsAndDaemonInit(t *testing.T) {
	os.Args = append(os.Args, "--mode=not-support-os")
	flags := ParseFlags()

	if flags.Debug != DEFAULT_DEBUG {
		t.Fatal("Default option `Debug` value mismatch")
	}

	if flags.Listen != (DEFAULT_HOST + ":" + DEFAULT_PORT) {
		t.Fatal("Default option `Listen` value mismatch")
	}

	if flags.Mode != getProxyMode(DEFAULT_MODE_NAME) {
		t.Fatal("Default option `Mode` value mismatch")
	}

	if flags.Ubuntu != DEFAULT_UBUNTU_MIRROR {
		t.Fatal("Default option `Ubuntu` value mismatch")
	}

	if flags.Debian != DEFAULT_DEBIAN_MIRROR {
		t.Fatal("Default option `Debian` value mismatch")
	}

	if flags.CacheDir != DEFAULT_CACHE_DIR {
		t.Fatal("Default option `CacheDir` value mismatch")
	}

	cache, err := initStore(flags)
	if err != nil {
		t.Fatal("Init Store Failed")
	}

	ap := initProxy(flags, cache)
	initLogger(flags, ap)
}

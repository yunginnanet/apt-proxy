package cli

import (
	"os"
	"testing"
)

func TestParseFlags(t *testing.T) {
	os.Args = append(os.Args, "--type=not-support-os")
	flags := ParseFlags()

	if flags.Debug != DEFAULT_DEBUG {
		t.Fatal("Default option `Debug` value mismatch")
	}

	if flags.Listen != (DEFAULT_HOST + ":" + DEFAULT_PORT) {
		t.Fatal("Default option `Listen` value mismatch")
	}

	if flags.Types != DEFAULT_TYPE {
		t.Fatal("Default option `Types` value mismatch")
	}

	if flags.Mirror != DEFAULT_MIRROR {
		t.Fatal("Default option `Mirror` value mismatch")
	}

	if flags.CacheDir != DEFAULT_CACHE_DIR {
		t.Fatal("Default option `CacheDir` value mismatch")
	}
}

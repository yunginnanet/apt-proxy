package cli

import (
	"testing"
)

func TestInit(t *testing.T) {
	flags := ParseFlags()

	cache, err := initStore(flags)
	if err != nil {
		t.Fatal("Init Store Failed")
	}

	ap := initProxy(flags, cache)
	initLogger(flags, ap)
}

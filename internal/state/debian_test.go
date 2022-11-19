package state_test

import (
	"strings"
	"testing"

	State "github.com/soulteary/apt-proxy/internal/state"
)

func TestGetAndSetDebianMirror(t *testing.T) {
	State.SetDebianMirror("https://mirrors.tuna.tsinghua.edu.cn/debian/")
	mirror := State.GetDebianMirror()
	if !strings.Contains(mirror.Path, "debian") {
		t.Fatal("Test Set/Get Debian Mirror Value Faild")
	}

	State.SetDebianMirror("")
	mirror = State.GetDebianMirror()
	if mirror != nil {
		t.Fatal("Test Set/Get Debian Mirror to Null Faild")
	}

	State.ResetDebianMirror()
	mirror = State.GetDebianMirror()
	if mirror != nil {
		t.Fatal("Test Clear Debian Mirror Faild")
	}
}

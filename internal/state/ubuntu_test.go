package state_test

import (
	"strings"
	"testing"

	State "github.com/soulteary/apt-proxy/internal/state"
)

func TestGetAndSetUbuntuMirror(t *testing.T) {
	State.SetUbuntuMirror("https://mirrors.tuna.tsinghua.edu.cn/ubuntu/")
	mirror := State.GetUbuntuMirror()
	if !strings.Contains(mirror.Path, "ubuntu") {
		t.Fatal("Test Set/Get Ubuntu Mirror Value Faild")
	}

	State.SetUbuntuMirror("")
	mirror = State.GetUbuntuMirror()
	if mirror != nil {
		t.Fatal("Test Set/Get Ubuntu Mirror to Null Faild")
	}

	State.ResetUbuntuMirror()
	mirror = State.GetUbuntuMirror()
	if mirror != nil {
		t.Fatal("Test Clear Ubuntu Mirror Faild")
	}
}

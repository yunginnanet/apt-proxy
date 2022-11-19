package state_test

import (
	"strings"
	"testing"

	State "github.com/soulteary/apt-proxy/internal/state"
)

func TestGetAndSetAlpineMirror(t *testing.T) {
	State.SetAlpineMirror("https://mirrors.tuna.tsinghua.edu.cn/alpine/")
	mirror := State.GetAlpineMirror()
	if !strings.Contains(strings.ToLower(mirror.Path), "alpine") {
		t.Fatal("Test Set/Get Alpine Mirror Value Faild")
	}

	State.SetAlpineMirror("")
	mirror = State.GetAlpineMirror()
	if mirror != nil {
		t.Fatal("Test Set/Get Alpine Mirror to Null Faild")
	}

	State.ResetAlpineMirror()
	mirror = State.GetAlpineMirror()
	if mirror != nil {
		t.Fatal("Test Clear Alpine Mirror Faild")
	}

	State.SetAlpineMirror("cn:tsinghua")
	mirror = State.GetAlpineMirror()
	if !strings.Contains(strings.ToLower(mirror.Path), "alpine") {
		t.Fatal("Test Set/Get Alpine Mirror Value Faild")
	}

	State.SetAlpineMirror("!#$%(not://abc")
	mirror = State.GetAlpineMirror()
	if mirror != nil {
		t.Fatal("Test Set/Get Alpine Mirror Value Faild")
	}
}

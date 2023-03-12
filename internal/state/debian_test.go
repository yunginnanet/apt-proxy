package state_test

import (
	"strings"
	"testing"

	State "github.com/soulteary/apt-proxy/internal/state"
)

func TestGetAndSetDebianMirror(t *testing.T) {
	State.SetDebianMirror("mirrors.163.com/debian/")
	mirror := State.GetDebianMirror()
	if !strings.Contains(mirror.Path, "debian") {
		t.Fatal("Test Set/Get Debian Mirror Value Faild")
	}
	if !strings.Contains(mirror.Path, "mirrors.163.com") {
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

	State.SetDebianMirror("cn:163")
	mirror = State.GetDebianMirror()
	if !strings.Contains(strings.ToLower(mirror.Path), "debian") {
		t.Fatal("Test Set/Get Debian Mirror Value Faild")
	}

	State.SetDebianMirror("!#$%(not://abc")
	mirror = State.GetDebianMirror()
	if mirror != nil {
		t.Fatal("Test Set/Get Debian Mirror Value Faild")
	}
}

package state_test

import (
	"strings"
	"testing"

	State "github.com/soulteary/apt-proxy/internal/state"
)

func TestGetAndSetCentOSMirror(t *testing.T) {
	State.SetCentOSMirror("https://mirrors.tuna.tsinghua.edu.cn/centos/")
	mirror := State.GetCentOSMirror()
	if !strings.Contains(strings.ToLower(mirror.Path), "centos") {
		t.Fatal("Test Set/Get CentOS Mirror Value Faild")
	}

	State.SetCentOSMirror("")
	mirror = State.GetCentOSMirror()
	if mirror != nil {
		t.Fatal("Test Set/Get CentOS Mirror to Null Faild")
	}

	State.ResetCentOSMirror()
	mirror = State.GetCentOSMirror()
	if mirror != nil {
		t.Fatal("Test Clear CentOS Mirror Faild")
	}

	State.SetCentOSMirror("cn:tsinghua")
	mirror = State.GetCentOSMirror()
	if !strings.Contains(strings.ToLower(mirror.Path), "centos") {
		t.Fatal("Test Set/Get CentOS Mirror Value Faild")
	}

	State.SetCentOSMirror("!#$%(not://abc")
	mirror = State.GetCentOSMirror()
	if mirror != nil {
		t.Fatal("Test Set/Get CentOS Mirror Value Faild")
	}
}

package state

import (
	"strings"
	"testing"
)

func TestSetProxyMode(t *testing.T) {
	// 1 => linux.TYPE_LINUX_DISTROS_UBUNTU
	SetProxyMode(1)
	if GetProxyMode() != 1 {
		t.Fatal("Test Set/Get ProxyMode Faild")
	}
}

func TestGetAndSetUbuntuMirror(t *testing.T) {
	SetUbuntuMirror("https://mirrors.tuna.tsinghua.edu.cn/ubuntu/")
	mirror := GetUbuntuMirror()
	if !strings.Contains(mirror.Path, "ubuntu") {
		t.Fatal("Test Set/Get Ubuntu Mirror Value Faild")
	}

	SetUbuntuMirror("")
	mirror = GetUbuntuMirror()
	if mirror != nil {
		t.Fatal("Test Set/Get Ubuntu Mirror to Null Faild")
	}

	ResetUbuntuMirror()
	mirror = GetUbuntuMirror()
	if mirror != nil {
		t.Fatal("Test Clear Ubuntu Mirror Faild")
	}
}

func TestGetAndSetDebianMirror(t *testing.T) {
	SetDebianMirror("https://mirrors.tuna.tsinghua.edu.cn/debian/")
	mirror := GetDebianMirror()
	if !strings.Contains(mirror.Path, "debian") {
		t.Fatal("Test Set/Get Debian Mirror Value Faild")
	}

	SetDebianMirror("")
	mirror = GetDebianMirror()
	if mirror != nil {
		t.Fatal("Test Set/Get Debian Mirror to Null Faild")
	}

	ResetDebianMirror()
	mirror = GetDebianMirror()
	if mirror != nil {
		t.Fatal("Test Clear Debian Mirror Faild")
	}
}

func TestGetAndSetCentOSMirror(t *testing.T) {
	SetCentOSMirror("https://mirrors.tuna.tsinghua.edu.cn/centos/")
	mirror := GetCentOSMirror()
	if !strings.Contains(strings.ToLower(mirror.Path), "centos") {
		t.Fatal("Test Set/Get CentOS Mirror Value Faild")
	}

	SetCentOSMirror("")
	mirror = GetCentOSMirror()
	if mirror != nil {
		t.Fatal("Test Set/Get CentOS Mirror to Null Faild")
	}

	ResetCentOSMirror()
	mirror = GetCentOSMirror()
	if mirror != nil {
		t.Fatal("Test Clear CentOS Mirror Faild")
	}
}

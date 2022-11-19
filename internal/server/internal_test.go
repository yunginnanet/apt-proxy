package server_test

import (
	"testing"

	Server "github.com/soulteary/apt-proxy/internal/server"
)

func TestIsInternalUrls(t *testing.T) {
	if Server.IsInternalUrls("mirrors.tuna.tsinghua.edu.cn/ubuntu/") {
		t.Fatal("test internal url failed")
	}
	if !Server.IsInternalUrls(Server.INTERNAL_PAGE_HOME) {
		t.Fatal("test internal url failed")
	}
	if !Server.IsInternalUrls(Server.INTERNAL_PAGE_PING) {
		t.Fatal("test internal url failed")
	}
}

func TestGetInternalResType(t *testing.T) {
	if Server.GetInternalResType(Server.INTERNAL_PAGE_HOME) != Server.TYPE_HOME {
		t.Fatal("test get internal res type failed")
	}

	if Server.GetInternalResType(Server.INTERNAL_PAGE_PING) != Server.TYPE_PING {
		t.Fatal("test get internal res type failed")
	}

	if Server.GetInternalResType("/url-not-found") != Server.TYPE_NOT_FOUND {
		t.Fatal("test get internal res type failed")
	}
}

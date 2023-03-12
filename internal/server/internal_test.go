package server_test

import (
	"fmt"
	"net/http"
	"testing"

	Server "github.com/soulteary/apt-proxy/internal/server"
)

func TestIsInternalUrls(t *testing.T) {
	if Server.IsInternalUrls("mirrors.tuna.tsinghua.edu.cn/ubuntu/") {
		t.Fatal("test internal url failed")
	}
	if !Server.IsInternalUrls(Server.InternalPageHome) {
		t.Fatal("test internal url failed")
	}
	if !Server.IsInternalUrls(Server.InternalPagePing) {
		t.Fatal("test internal url failed")
	}
}

func TestGetInternalResType(t *testing.T) {
	if Server.GetInternalResType(Server.InternalPageHome) != Server.TypeHome {
		t.Fatal("test get internal res type failed")
	}

	if Server.GetInternalResType(Server.InternalPagePing) != Server.TypePing {
		t.Fatal("test get internal res type failed")
	}

	if Server.GetInternalResType("/url-not-found") != Server.TypeNotFound {
		t.Fatal("test get internal res type failed")
	}
}

func TestRenderInternalUrls(t *testing.T) {
	res, code := Server.RenderInternalUrls(Server.InternalPagePing)
	if code != http.StatusOK {
		t.Fatal("test render internal urls failed")
	}
	if res != "pong" {
		t.Fatal("test render internal urls failed")
	}

	_, code = Server.RenderInternalUrls("/url-not-exists")
	if code != http.StatusNotFound {
		t.Fatal("test render internal urls failed")
	}

	res, code = Server.RenderInternalUrls(Server.InternalPageHome)
	fmt.Println(res)
	if !(code == http.StatusOK || code == http.StatusBadGateway) {
		t.Fatal("test render internal urls failed")
	}
	if res == "" {
		t.Fatal("test render internal urls failed")
	}
}

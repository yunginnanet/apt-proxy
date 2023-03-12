package state

import (
	"testing"

	Define "github.com/soulteary/apt-proxy/internal/define"
)

func TestSetProxyMode(t *testing.T) {
	SetProxyMode(Define.TypeLinuxAllDistros)
	if GetProxyMode() != Define.TypeLinuxAllDistros {
		t.Fatal("Test Set/Get ProxyMode Faild")
	}
}

package state

var PROXY_MODE = 0

func SetProxyMode(mode int) {
	PROXY_MODE = mode
}

func GetProxyMode() int {
	return PROXY_MODE
}

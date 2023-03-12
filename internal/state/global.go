package state

var ProxyMode = 0

func SetProxyMode(mode int) {
	ProxyMode = mode
}

func GetProxyMode() int {
	return ProxyMode
}

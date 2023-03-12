package mirrors

import (
	"bufio"
	"net/http"

	Define "github.com/soulteary/apt-proxy/internal/define"
)

func GetUbuntuMirrorUrlsByGeo() (mirrors []string, err error) {
	response, err := http.Get(Define.UbuntuGeoMirrorApi)
	if err != nil {
		return mirrors, err
	}
	defer response.Body.Close()

	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		mirrors = append(mirrors, scanner.Text())
	}
	return mirrors, scanner.Err()
}

package define

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	LinuxAllDistros    string = "all"
	LinuxDistrosUbuntu string = "ubuntu"
	LinuxDistrosDebian string = "debian"
	LinuxDistrosCentos string = "centos"
	LinuxDistrosAlpine string = "alpine"
)

const (
	TypeLinuxAllDistros    int = 0
	TypeLinuxDistrosUbuntu int = 1
	TypeLinuxDistrosDebian int = 2
	TypeLinuxDistrosCentos int = 3
	TypeLinuxDistrosAlpine int = 4
)

type Rule struct {
	OS           int
	Pattern      *regexp.Regexp
	CacheControl string
	Rewrite      bool
}

func (r *Rule) String() string {
	return fmt.Sprintf("%s Cache-Control=%s Rewrite=%#v",
		r.Pattern.String(), r.CacheControl, r.Rewrite)
}

type UrlWithAlias struct {
	URL       string
	Alias     string
	Http      bool
	Https     bool
	Official  bool
	Bandwidth int64
}

func GenerateAliasFromURL(url string) string {
	pureHost := regexp.MustCompile(`^https?://|\/.*`).ReplaceAllString(url, "")
	tldRemoved := regexp.MustCompile(`\.edu\.cn$|\.edu$|\.mit\.edu$|\.cn$|\.com$|\.net$|\.net.cn$|\.org$|\.org\.cn$`).ReplaceAllString(pureHost, "")
	group := strings.Split(tldRemoved, ".")
	println("cn:" + group[len(group)-1])
	return "cn:" + group[len(group)-1]
}

func GenerateBuildInMirorItem(url string, official bool) UrlWithAlias {
	var mirror UrlWithAlias
	mirror.Official = official
	mirror.Alias = GenerateAliasFromURL(url)

	if strings.HasPrefix(url, "http://") {
		mirror.Http = true
		mirror.Https = false
	} else if strings.HasPrefix(url, "https://") {
		mirror.Http = false
		mirror.Https = true
	}
	mirror.URL = url
	// TODO
	mirror.Bandwidth = 0
	return mirror
}

func GenerateBuildInList(officialList []string, customList []string) (mirrors []UrlWithAlias) {
	for _, url := range officialList {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			mirror := GenerateBuildInMirorItem("http://"+url, true)
			mirrors = append(mirrors, mirror)
			mirror = GenerateBuildInMirorItem("https://"+url, true)
			mirrors = append(mirrors, mirror)
		} else {
			mirror := GenerateBuildInMirorItem(url, true)
			mirrors = append(mirrors, mirror)
		}
	}

	for _, url := range customList {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			mirror := GenerateBuildInMirorItem("http://"+url, false)
			mirrors = append(mirrors, mirror)
			mirror = GenerateBuildInMirorItem("https://"+url, false)
			mirrors = append(mirrors, mirror)
		} else {
			mirror := GenerateBuildInMirorItem(url, false)
			mirrors = append(mirrors, mirror)
		}
	}

	return mirrors
}

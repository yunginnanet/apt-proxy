package define

import (
	"fmt"
	"regexp"
)

const (
	LINUX_ALL_DISTROS    string = "all"
	LINUX_DISTROS_UBUNTU string = "ubuntu"
	LINUX_DISTROS_DEBIAN string = "debian"
	LINUX_DISTROS_CENTOS string = "centos"
)

const (
	TYPE_LINUX_ALL_DISTROS    int = 0
	TYPE_LINUX_DISTROS_UBUNTU int = 1
	TYPE_LINUX_DISTROS_DEBIAN int = 2
	TYPE_LINUX_DISTROS_CENTOS int = 3
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
	URL   string
	Alias string
}

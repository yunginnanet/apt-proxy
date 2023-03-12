package define

import "regexp"

var AlpineHostPattern = regexp.MustCompile(`/alpine/(.+)$`)

const AlpineBenchmakrUrl = "MIRRORS.txt"

// AlpineOfficialMirrors contains official mirrors of alpine
var AlpineOfficialMirrors = []string{
	"https://mirror.math.princeton.edu/pub/alpinelinux/",
	"https://mirrors.ocf.berkeley.edu/alpine/",
	"http://mirror.fcix.net/alpine/",
}

var AlpineCustomMirrors []string

var BuildinAlpineMirrors = GenerateBuildInList(AlpineOfficialMirrors, AlpineCustomMirrors)

var AlpineDefaultCacheRules = []Rule{
	{Pattern: regexp.MustCompile(`APKINDEX.tar.gz$`), CacheControl: `max-age=3600`, Rewrite: true, OS: TypeLinuxDistrosAlpine},
	{Pattern: regexp.MustCompile(`tar.gz$`), CacheControl: `max-age=3600`, Rewrite: true, OS: TypeLinuxDistrosAlpine},
	{Pattern: regexp.MustCompile(`apk$`), CacheControl: `max-age=3600`, Rewrite: true, OS: TypeLinuxDistrosAlpine},
	{Pattern: regexp.MustCompile(`.*`), CacheControl: `max-age=100000`, Rewrite: true, OS: TypeLinuxDistrosAlpine},
}

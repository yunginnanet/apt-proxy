package define

import (
	"regexp"
)

const (
	UbuntuGeoMirrorApi = "http://mirrors.ubuntu.com/mirrors.txt"
	UbuntuBenchmakrUrl = "dists/jammy/main/binary-amd64/Release"
)

var UbuntuHostPattern = regexp.MustCompile(`/ubuntu/(.+)$`)

var UbuntuOfficialMirrors = []string{
	// ubuntu mirrors (not debian!), based in the united states, west is better
	"mirror.math.princeton.edu/pub/ubuntu/",
	"mirrors.ocf.berkeley.edu/ubuntu/",
	"mirror.fcix.net/ubuntu/",
	"mirror.csclub.uwaterloo.ca/ubuntu/",
	"mirrors.vcea.wsu.edu/ubuntu/",
	"mirror.siena.edu/ubuntu/",
	"mirror.umd.edu/ubuntu/",
	"mirror.cs.uchicago.edu/ubuntu/",
	"mirror.csclub.uwaterloo.ca/ubuntu/",
	"mirror.cse.ohio-state.edu/ubuntu/",
}

var UbuntuCustomMirrors = []string{
	"mirrors.163.com/ubuntu/",
}

var UbuntuDefaults = GenerateBuildInList(UbuntuOfficialMirrors, UbuntuCustomMirrors)

var UbuntuDefaultCacheRules = []Rule{
	{Pattern: regexp.MustCompile(`deb$`), CacheControl: `max-age=100000`, Rewrite: true, OS: TypeLinuxDistrosUbuntu},
	{Pattern: regexp.MustCompile(`udeb$`), CacheControl: `max-age=100000`, Rewrite: true, OS: TypeLinuxDistrosUbuntu},
	{Pattern: regexp.MustCompile(`InRelease$`), CacheControl: `max-age=3600`, Rewrite: true, OS: TypeLinuxDistrosUbuntu},
	{Pattern: regexp.MustCompile(`DiffIndex$`), CacheControl: `max-age=3600`, Rewrite: true, OS: TypeLinuxDistrosUbuntu},
	{Pattern: regexp.MustCompile(`PackagesIndex$`), CacheControl: `max-age=3600`, Rewrite: true, OS: TypeLinuxDistrosUbuntu},
	{Pattern: regexp.MustCompile(`Packages\.(bz2|gz|lzma)$`), CacheControl: `max-age=3600`, Rewrite: true, OS: TypeLinuxDistrosUbuntu},
	{Pattern: regexp.MustCompile(`SourcesIndex$`), CacheControl: `max-age=3600`, Rewrite: true, OS: TypeLinuxDistrosUbuntu},
	{Pattern: regexp.MustCompile(`Sources\.(bz2|gz|lzma)$`), CacheControl: `max-age=3600`, Rewrite: true, OS: TypeLinuxDistrosUbuntu},
	{Pattern: regexp.MustCompile(`Release(\.gpg)?$`), CacheControl: `max-age=3600`, Rewrite: true, OS: TypeLinuxDistrosUbuntu},
	{Pattern: regexp.MustCompile(`Translation-(en|fr)\.(gz|bz2|bzip2|lzma)$`), CacheControl: `max-age=3600`, Rewrite: true, OS: TypeLinuxDistrosUbuntu},
	// Add file file hash
	{Pattern: regexp.MustCompile(`\/by-hash\/`), CacheControl: `max-age=3600`, Rewrite: true, OS: TypeLinuxDistrosUbuntu},
}

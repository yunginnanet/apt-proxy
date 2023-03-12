package define

import "regexp"

const (
	DebianBenchmakrUrl = "dists/bullseye/main/binary-amd64/Release"
)

var DebianHostPattern = regexp.MustCompile(`/debian/(.+)$`)

// DebianOfficialMirrors contains official mirrors of debian
var DebianOfficialMirrors = []string{
	"ftp.us.debian.org/debian/",
	"atl.mirrors.clouvider.net/debian/",
	"mirrors.163.com/debian/",
	"debian.csail.mit.edu/debian/",
	"mirrors.vcea.wsu.edu/debian/",
}

var DebianCustomMirrors = []string{
	"nyc.mirrors.clouvider.net/debian/",
}

var BuildinDebianMirrors = GenerateBuildInList(DebianOfficialMirrors, DebianCustomMirrors)

var DebianDefaultCacheRules = []Rule{
	{Pattern: regexp.MustCompile(`deb$`), CacheControl: `max-age=100000`, Rewrite: true, OS: TypeLinuxDistrosDebian},
	{Pattern: regexp.MustCompile(`udeb$`), CacheControl: `max-age=100000`, Rewrite: true, OS: TypeLinuxDistrosDebian},
	{Pattern: regexp.MustCompile(`InRelease$`), CacheControl: `max-age=3600`, Rewrite: true, OS: TypeLinuxDistrosDebian},
	{Pattern: regexp.MustCompile(`DiffIndex$`), CacheControl: `max-age=3600`, Rewrite: true, OS: TypeLinuxDistrosDebian},
	{Pattern: regexp.MustCompile(`PackagesIndex$`), CacheControl: `max-age=3600`, Rewrite: true, OS: TypeLinuxDistrosDebian},
	{Pattern: regexp.MustCompile(`Packages\.(bz2|gz|lzma)$`), CacheControl: `max-age=3600`, Rewrite: true, OS: TypeLinuxDistrosDebian},
	{Pattern: regexp.MustCompile(`SourcesIndex$`), CacheControl: `max-age=3600`, Rewrite: true, OS: TypeLinuxDistrosDebian},
	{Pattern: regexp.MustCompile(`Sources\.(bz2|gz|lzma)$`), CacheControl: `max-age=3600`, Rewrite: true, OS: TypeLinuxDistrosDebian},
	{Pattern: regexp.MustCompile(`Release(\.gpg)?$`), CacheControl: `max-age=3600`, Rewrite: true, OS: TypeLinuxDistrosDebian},
	{Pattern: regexp.MustCompile(`Translation-(en|fr)\.(gz|bz2|bzip2|lzma)$`), CacheControl: `max-age=3600`, Rewrite: true, OS: TypeLinuxDistrosDebian},
	// Add file file hash
	{Pattern: regexp.MustCompile(`/by-hash/`), CacheControl: `max-age=3600`, Rewrite: true, OS: TypeLinuxDistrosDebian},
}

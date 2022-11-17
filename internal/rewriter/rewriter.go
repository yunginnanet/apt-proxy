package rewriter

import (
	"log"
	"net/http"
	"net/url"
	"regexp"

	Mirrors "github.com/soulteary/apt-proxy/internal/mirrors"
	"github.com/soulteary/apt-proxy/state"
)

type URLRewriters struct {
	ubuntu *URLRewriter
	debian *URLRewriter
}

type URLRewriter struct {
	mirror  *url.URL
	pattern *regexp.Regexp
}

func GetRewriteRulesByMode(mode int) (rules []Mirrors.Rule) {
	if mode == Mirrors.TYPE_LINUX_DISTROS_UBUNTU {
		return Mirrors.UBUNTU_DEFAULT_CACHE_RULES
	}
	if mode == Mirrors.TYPE_LINUX_DISTROS_DEBIAN {
		return Mirrors.DEBIAN_DEFAULT_CACHE_RULES
	}

	rules = append(rules, Mirrors.UBUNTU_DEFAULT_CACHE_RULES...)
	rules = append(rules, Mirrors.DEBIAN_DEFAULT_CACHE_RULES...)
	return rules
}

func getRewriterForDebian() *URLRewriter {
	u := &URLRewriter{}
	debianMirror := state.GetDebianMirror()
	benchmarkUrl, pattern := Mirrors.GetPredefinedConfiguration(Mirrors.TYPE_LINUX_DISTROS_DEBIAN)
	u.pattern = pattern

	if debianMirror != nil {
		log.Printf("using specify debian mirror %s", debianMirror)
		u.mirror = debianMirror
		return u
	}

	mirrors := Mirrors.GetGeoMirrorUrlsByMode(Mirrors.TYPE_LINUX_DISTROS_DEBIAN)
	fastest, err := Mirrors.GetTheFastestMirror(mirrors, benchmarkUrl)
	if err != nil {
		log.Println("Error finding fastest mirror", err)
	}

	if mirror, err := url.Parse(fastest); err == nil {
		log.Printf("using fastest mirror %s", fastest)
		u.mirror = mirror
	}

	return u
}

func getRewriterForUbuntu() *URLRewriter {
	u := &URLRewriter{}
	ubuntuMirror := state.GetUbuntuMirror()
	benchmarkUrl, pattern := Mirrors.GetPredefinedConfiguration(Mirrors.TYPE_LINUX_DISTROS_UBUNTU)
	u.pattern = pattern

	if ubuntuMirror != nil {
		log.Printf("using specify ubuntu mirror %s", ubuntuMirror)
		u.mirror = ubuntuMirror
		return u
	}

	mirrors := Mirrors.GetGeoMirrorUrlsByMode(Mirrors.TYPE_LINUX_DISTROS_UBUNTU)
	fastest, err := Mirrors.GetTheFastestMirror(mirrors, benchmarkUrl)
	if err != nil {
		log.Println("Error finding fastest mirror", err)
	}

	if mirror, err := url.Parse(fastest); err == nil {
		log.Printf("using fastest mirror %s", fastest)
		u.mirror = mirror
	}

	return u
}

func CreateNewRewriters(mode int) *URLRewriters {
	rewriters := &URLRewriters{}

	if mode == Mirrors.TYPE_LINUX_DISTROS_DEBIAN {
		rewriters.debian = getRewriterForDebian()
		return rewriters
	}

	if mode == Mirrors.TYPE_LINUX_DISTROS_UBUNTU {
		rewriters.ubuntu = getRewriterForUbuntu()
		return rewriters
	}

	rewriters.debian = getRewriterForDebian()
	rewriters.ubuntu = getRewriterForUbuntu()
	return rewriters
}

func RewriteRequestByMode(r *http.Request, rewriters *URLRewriters, mode int) {
	uri := r.URL.String()
	var rewriter *URLRewriter
	if mode == Mirrors.TYPE_LINUX_DISTROS_UBUNTU {
		rewriter = rewriters.ubuntu
	} else {
		// mode == TYPE_LINUX_DISTROS_DEBIAN
		rewriter = rewriters.debian
	}

	if rewriter.mirror != nil && rewriter.pattern.MatchString(uri) {
		r.Header.Add("Content-Location", uri)
		m := rewriter.pattern.FindAllStringSubmatch(uri, -1)
		// Fix the problem of double escaping of symbols
		unescapedQuery, err := url.PathUnescape(m[0][3])
		if err != nil {
			unescapedQuery = m[0][3]
		}
		r.URL.Host = rewriter.mirror.Host
		r.URL.Path = rewriter.mirror.Path + unescapedQuery
	}
}

func MatchingRule(subject string, rules []Mirrors.Rule) (*Mirrors.Rule, bool) {
	for _, rule := range rules {
		if rule.Pattern.MatchString(subject) {
			return &rule, true
		}
	}
	return nil, false
}

package rewriter

import (
	"log"
	"net/http"
	"net/url"
	"regexp"

	Benchmark "github.com/soulteary/apt-proxy/internal/benchmark"
	Define "github.com/soulteary/apt-proxy/internal/define"
	Mirrors "github.com/soulteary/apt-proxy/internal/mirrors"
	State "github.com/soulteary/apt-proxy/internal/state"
)

type URLRewriters struct {
	ubuntu *URLRewriter
	debian *URLRewriter
	centos *URLRewriter
	alpine *URLRewriter
}

type URLRewriter struct {
	mirror  *url.URL
	pattern *regexp.Regexp
}

func GetRewriteRulesByMode(mode int) (rules []Define.Rule) {
	if mode == Define.TypeLinuxDistrosUbuntu {
		return Define.UbuntuDefaultCacheRules
	}
	if mode == Define.TypeLinuxDistrosDebian {
		return Define.DebianDefaultCacheRules
	}
	if mode == Define.TypeLinuxDistrosCentos {
		return Define.CentosDefaultCacheRules
	}
	if mode == Define.TypeLinuxDistrosAlpine {
		return Define.AlpineDefaultCacheRules
	}

	rules = append(rules, Define.UbuntuDefaultCacheRules...)
	rules = append(rules, Define.DebianDefaultCacheRules...)
	rules = append(rules, Define.CentosDefaultCacheRules...)
	rules = append(rules, Define.AlpineDefaultCacheRules...)
	return rules
}

func getRewriterForAlpine() *URLRewriter {
	u := &URLRewriter{}
	mirror := State.GetAlpineMirror()
	benchmarkUrl, pattern := Mirrors.GetPredefinedConfiguration(Define.TypeLinuxDistrosAlpine)
	u.pattern = pattern

	if mirror != nil {
		log.Printf("using specify Alpine mirror %s", mirror)
		u.mirror = mirror
		return u
	}

	mirrors := Mirrors.GetGeoMirrorUrlsByMode(Define.TypeLinuxDistrosAlpine)
	fastest, err := Benchmark.GetTheFastestMirror(mirrors, benchmarkUrl)
	if err != nil {
		log.Println("Error finding fastest Alpine mirror", err)
	}

	if mirror, err := url.Parse(fastest); err == nil {
		log.Printf("using fastest Alpine mirror %s", fastest)
		u.mirror = mirror
	}

	return u
}

func getRewriterForCentOS() *URLRewriter {
	u := &URLRewriter{}
	mirror := State.GetCentOSMirror()
	benchmarkUrl, pattern := Mirrors.GetPredefinedConfiguration(Define.TypeLinuxDistrosCentos)
	u.pattern = pattern

	if mirror != nil {
		log.Printf("using specify CentOS mirror %s", mirror)
		u.mirror = mirror
		return u
	}

	mirrors := Mirrors.GetGeoMirrorUrlsByMode(Define.TypeLinuxDistrosCentos)
	fastest, err := Benchmark.GetTheFastestMirror(mirrors, benchmarkUrl)
	if err != nil {
		log.Println("Error finding fastest CentOS mirror", err)
	}

	if mirror, err := url.Parse(fastest); err == nil {
		log.Printf("using fastest CentOS mirror %s", fastest)
		u.mirror = mirror
	}

	return u
}

func getRewriterForDebian() *URLRewriter {
	u := &URLRewriter{}
	mirror := State.GetDebianMirror()
	benchmarkUrl, pattern := Mirrors.GetPredefinedConfiguration(Define.TypeLinuxDistrosDebian)
	u.pattern = pattern

	if mirror != nil {
		log.Printf("using specify Debian mirror %s", mirror)
		u.mirror = mirror
		return u
	}

	mirrors := Mirrors.GetGeoMirrorUrlsByMode(Define.TypeLinuxDistrosDebian)
	fastest, err := Benchmark.GetTheFastestMirror(mirrors, benchmarkUrl)
	if err != nil {
		log.Println("Error finding fastest Debian mirror", err)
	}

	if mirror, err := url.Parse(fastest); err == nil {
		log.Printf("using fastest Debian mirror %s", fastest)
		u.mirror = mirror
	}

	return u
}

func getRewriterForUbuntu() *URLRewriter {
	u := &URLRewriter{}
	mirror := State.GetUbuntuMirror()
	benchmarkUrl, pattern := Mirrors.GetPredefinedConfiguration(Define.TypeLinuxDistrosUbuntu)
	u.pattern = pattern

	if mirror != nil {
		log.Printf("using Ubuntu mirror %s", mirror)
		u.mirror = mirror
		return u
	}

	mirrors := Mirrors.GetGeoMirrorUrlsByMode(Define.TypeLinuxDistrosUbuntu)
	fastest, err := Benchmark.GetTheFastestMirror(mirrors, benchmarkUrl)
	if err != nil {
		log.Println("Error finding fastest Ubuntu mirror", err)
	}

	if mirror, err := url.Parse(fastest); err == nil {
		log.Printf("using fastest Ubuntu mirror %s", fastest)
		u.mirror = mirror
	}

	return u
}

func CreateNewRewriters(mode int) *URLRewriters {
	rewriters := &URLRewriters{}

	if mode == Define.TypeLinuxDistrosUbuntu {
		rewriters.ubuntu = getRewriterForUbuntu()
		return rewriters
	}

	if mode == Define.TypeLinuxDistrosDebian {
		rewriters.debian = getRewriterForDebian()
		return rewriters
	}

	if mode == Define.TypeLinuxDistrosCentos {
		rewriters.centos = getRewriterForCentOS()
		return rewriters
	}

	if mode == Define.TypeLinuxDistrosAlpine {
		rewriters.alpine = getRewriterForAlpine()
		return rewriters
	}

	rewriters.ubuntu = getRewriterForUbuntu()
	rewriters.debian = getRewriterForDebian()
	rewriters.centos = getRewriterForCentOS()
	rewriters.alpine = getRewriterForAlpine()
	return rewriters
}

func RewriteRequestByMode(r *http.Request, rewriters *URLRewriters, mode int) {
	uri := r.URL.String()

	var rewriter *URLRewriter
	switch mode {
	case Define.TypeLinuxDistrosUbuntu:
		rewriter = rewriters.ubuntu
	case Define.TypeLinuxDistrosDebian:
		rewriter = rewriters.debian
	case Define.TypeLinuxDistrosCentos:
		rewriter = rewriters.centos
	case Define.TypeLinuxDistrosAlpine:
		rewriter = rewriters.alpine
	}

	if rewriter.mirror != nil && rewriter.pattern.MatchString(uri) {
		r.Header.Add("Content-Location", uri)
		m := rewriter.pattern.FindAllStringSubmatch(uri, -1)
		// Fix the problem of double escaping of symbols
		queryRaw := m[0][len(m[0])-1]
		unescapedQuery, err := url.PathUnescape(queryRaw)
		if err != nil {
			unescapedQuery = queryRaw
		}
		r.URL.Scheme = rewriter.mirror.Scheme
		r.URL.Host = rewriter.mirror.Host
		r.URL.Path = rewriter.mirror.Path + unescapedQuery
	}
}

func MatchingRule(subject string, rules []Define.Rule) (*Define.Rule, bool) {
	for _, rule := range rules {
		if rule.Pattern.MatchString(subject) {
			return &rule, true
		}
	}
	return nil, false
}

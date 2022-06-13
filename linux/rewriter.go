package linux

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

type URLRewriter struct {
	mirror  *url.URL
	pattern *regexp.Regexp
}

func NewRewriter(mirrorUrl string, osType string) *URLRewriter {
	u := &URLRewriter{}

	// user specify mirror
	// TODO Both systems are supported
	if len(mirrorUrl) > 0 {
		mirror, err := url.Parse(mirrorUrl)
		if err == nil {
			log.Printf("using specify mirror %s", mirrorUrl)
			u.mirror = mirror
			_, pattern := getPredefinedConfiguration(osType)
			u.pattern = pattern
			return u
		}
	}

	benchmarkUrl, pattern := getPredefinedConfiguration(osType)
	u.pattern = pattern
	mirrors := getGeoMirrorUrlsByMode(osType)
	mirrorUrl, err := getTheFastestMirror(mirrors, benchmarkUrl)
	if err != nil {
		log.Println("Error finding fastest mirror", err)
	}

	if mirror, err := url.Parse(mirrorUrl); err == nil {
		log.Printf("using fastest mirror %s", mirrorUrl)
		u.mirror = mirror
	}

	return u
}

func Rewrite(r *http.Request, rewriter *URLRewriter) {
	uri := r.URL.String()
	if rewriter.mirror != nil && rewriter.pattern.MatchString(uri) {
		r.Header.Add("Content-Location", uri)
		m := rewriter.pattern.FindAllStringSubmatch(uri, -1)
		// Fix the problem of double escaping of symbols
		unescapedQuery, err := url.PathUnescape(m[0][2])
		if err != nil {
			unescapedQuery = m[0][2]
		}
		r.URL.Host = rewriter.mirror.Host
		r.URL.Path = rewriter.mirror.Path + unescapedQuery
	}
}

func MatchingRule(subject string, rules []Rule) (*Rule, bool) {
	for _, rule := range rules {
		if rule.Pattern.MatchString(subject) {
			return &rule, true
		}
	}
	return nil, false
}

func (r *Rule) String() string {
	return fmt.Sprintf("%s Cache-Control=%s Rewrite=%#v",
		r.Pattern.String(), r.CacheControl, r.Rewrite)
}

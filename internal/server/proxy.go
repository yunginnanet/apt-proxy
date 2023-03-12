package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	Define "github.com/soulteary/apt-proxy/internal/define"
	Rewriter "github.com/soulteary/apt-proxy/internal/rewriter"
	State "github.com/soulteary/apt-proxy/internal/state"
)

var rewriter *Rewriter.URLRewriters

var defaultTransport http.RoundTripper = &http.Transport{
	Proxy:                 http.ProxyFromEnvironment,
	ResponseHeaderTimeout: time.Second * 45,
	DisableKeepAlives:     true,
}

type AptProxy struct {
	Handler http.Handler
	Rules   []Define.Rule
}

func CreateAptProxyRouter() *AptProxy {
	mode := State.GetProxyMode()
	rewriter = Rewriter.CreateNewRewriters(mode)

	return &AptProxy{
		Rules: Rewriter.GetRewriteRulesByMode(mode),
		Handler: &httputil.ReverseProxy{
			Director:  func(r *http.Request) {},
			Transport: defaultTransport,
		},
	}
}

func (ap *AptProxy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var rule *Define.Rule
	if IsInternalUrls(r.URL.Path) {
		rule = nil
		tpl, status := RenderInternalUrls(r.URL.Path)

		rw.WriteHeader(status)
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, err := io.WriteString(rw, tpl)
		if err != nil {
			fmt.Println("render internal urls error")
		}
	} else {
		if Define.UbuntuHostPattern.MatchString(r.URL.Path) {
			rule, match := Rewriter.MatchingRule(r.URL.Path, Define.UbuntuDefaultCacheRules)
			if match {
				r.Header.Del("Cache-Control")
				if rule.Rewrite {
					before := r.URL.String()
					Rewriter.RewriteRequestByMode(r, rewriter, rule.OS)
					log.Printf("rewrote %q to %q", before, r.URL.String())
					r.Host = r.URL.Host
				}
			}
		}

		if Define.DebianHostPattern.MatchString(r.URL.Path) {
			rule, match := Rewriter.MatchingRule(r.URL.Path, Define.DebianDefaultCacheRules)
			if match {
				r.Header.Del("Cache-Control")
				if rule.Rewrite {
					before := r.URL.String()
					Rewriter.RewriteRequestByMode(r, rewriter, rule.OS)
					log.Printf("rewrote %q to %q", before, r.URL.String())
					r.Host = r.URL.Host
				}
			}
		}

		if Define.CentosHostPattern.MatchString(r.URL.Path) {
			rule, match := Rewriter.MatchingRule(r.URL.Path, Define.CentosDefaultCacheRules)
			if match {
				r.Header.Del("Cache-Control")
				if rule.Rewrite {
					before := r.URL.String()
					Rewriter.RewriteRequestByMode(r, rewriter, rule.OS)
					log.Printf("rewrote %q to %q", before, r.URL.String())
					r.Host = r.URL.Host
				}
			}
		}

		if Define.AlpineHostPattern.MatchString(r.URL.Path) {
			rule, match := Rewriter.MatchingRule(r.URL.Path, Define.AlpineDefaultCacheRules)
			if match {
				r.Header.Del("Cache-Control")
				if rule.Rewrite {
					before := r.URL.String()
					Rewriter.RewriteRequestByMode(r, rewriter, rule.OS)
					log.Printf("rewrote %q to %q", before, r.URL.String())
					r.Host = r.URL.Host
				}
			}
		}
	}
	ap.Handler.ServeHTTP(&responseWriter{rw, rule}, r)
}

type responseWriter struct {
	http.ResponseWriter
	rule *Define.Rule
}

func (rw *responseWriter) WriteHeader(status int) {
	if rw.rule != nil && rw.rule.CacheControl != "" &&
		(status == http.StatusOK || status == http.StatusNotFound) {
		rw.Header().Set("Cache-Control", rw.rule.CacheControl)
	}
	rw.ResponseWriter.WriteHeader(status)
}

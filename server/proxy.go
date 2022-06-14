package server

import (
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/soulteary/apt-proxy/linux"
	"github.com/soulteary/apt-proxy/state"
)

var rewriter *linux.URLRewriters

var defaultTransport http.RoundTripper = &http.Transport{
	Proxy:                 http.ProxyFromEnvironment,
	ResponseHeaderTimeout: time.Second * 45,
	DisableKeepAlives:     true,
}

type AptProxy struct {
	Handler http.Handler
	Rules   []linux.Rule
}

func CreateAptProxyRouter() *AptProxy {
	mode := state.GetProxyMode()
	rewriter = linux.CreateNewRewriters(mode)

	return &AptProxy{
		Rules: linux.GetRewriteRulesByMode(mode),
		Handler: &httputil.ReverseProxy{
			Director:  func(r *http.Request) {},
			Transport: defaultTransport,
		},
	}
}

func (ap *AptProxy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var rule *linux.Rule
	if isInternalUrls(r.URL.Path) {
		rule = nil
		renderInternalUrls(r.URL.Path, &rw)
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	} else {
		rule, match := linux.MatchingRule(r.URL.Path, ap.Rules)
		if match {
			r.Header.Del("Cache-Control")
			if rule.Rewrite {
				before := r.URL.String()
				linux.RewriteRequestByMode(r, rewriter, rule.OS)
				log.Printf("rewrote %q to %q", before, r.URL.String())
				r.Host = r.URL.Host
			}
		}
	}
	ap.Handler.ServeHTTP(&responseWriter{rw, rule}, r)
}

type responseWriter struct {
	http.ResponseWriter
	rule *linux.Rule
}

func (rw *responseWriter) WriteHeader(status int) {
	if rw.rule != nil && rw.rule.CacheControl != "" &&
		(status == http.StatusOK || status == http.StatusNotFound) {
		rw.Header().Set("Cache-Control", rw.rule.CacheControl)
	}
	rw.ResponseWriter.WriteHeader(status)
}

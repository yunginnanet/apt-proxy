package linux

import (
	"strings"
	"testing"
	"time"
)

func TestNewRewriter(t *testing.T) {
	ap := *NewRewriter("", LINUX_DISTROS_UBUNTU)

	time.Sleep((BENCHMARK_DETECT_TIMEOUT + 1) * time.Second)

	if len(ap.mirror.Path) == 0 {
		t.Fatal("No mirrors found")
	}
	if !strings.Contains(ap.mirror.Path, "ubuntu") {
		t.Fatal("mirror path incorrect")
	}
}

func TestNewRewriterWithSpecifyMirror(t *testing.T) {
	ap := *NewRewriter("https://mirrors.tuna.tsinghua.edu.cn/ubuntu/", LINUX_DISTROS_UBUNTU)
	if ap.mirror.Host != "mirrors.tuna.tsinghua.edu.cn" {
		t.Fatal("Mirror host incorrect")
	}
	if !strings.Contains(ap.mirror.Path, "ubuntu") {
		t.Fatal("mirror path incorrect")
	}
}

func TestMatchingRule(t *testing.T) {
	_, match := MatchingRule("http://archive.ubuntu.com/ubuntu/InRelease", UBUNTU_DEFAULT_CACHE_RULES)
	if !match {
		t.Fatal("test ap rules faild")
	}
}

func TestRuleToString(t *testing.T) {
	rule, _ := MatchingRule("http://archive.ubuntu.com/ubuntu/InRelease", UBUNTU_DEFAULT_CACHE_RULES)
	if rule.String() != "Release(\\.gpg)?$ Cache-Control=max-age=3600 Rewrite=true" {
		t.Fatal("test rules toString faild")
	}
}

package linux

import (
	"strings"
	"testing"
	"time"

	"github.com/soulteary/apt-proxy/state"
)

func TestGetRewriteRulesByMode(t *testing.T) {
	rules := GetRewriteRulesByMode(TYPE_LINUX_DISTROS_UBUNTU)
	if rules[0].Pattern != UBUNTU_DEFAULT_CACHE_RULES[0].Pattern {
		t.Fatal("Pattern Not Match")
	}

	rules = GetRewriteRulesByMode(TYPE_LINUX_DISTROS_DEBIAN)
	if rules[0].Pattern != DEBIAN_DEFAULT_CACHE_RULES[0].Pattern {
		t.Fatal("Pattern Not Match")
	}

	rules = GetRewriteRulesByMode(TYPE_LINUX_ALL_DISTROS)
	if len(rules) != (len(DEBIAN_DEFAULT_CACHE_RULES) + len(UBUNTU_DEFAULT_CACHE_RULES)) {
		t.Fatal("Pattern Length Not Match")
	}
}

func TestGetRewriterForDebian(t *testing.T) {
	debian := getRewriterForDebian()
	if !strings.Contains(debian.mirror.Path, "debian") {
		t.Fatal("mirror path incorrect")
	}
}

func TestGetRewriterForUbuntu(t *testing.T) {
	ubuntu := getRewriterForUbuntu()
	if !strings.Contains(ubuntu.mirror.Path, "ubuntu") {
		t.Fatal("mirror path incorrect")
	}
}

func TestCreateNewRewritersForUbuntu(t *testing.T) {
	ap := *CreateNewRewriters(TYPE_LINUX_DISTROS_UBUNTU)
	time.Sleep((BENCHMARK_DETECT_TIMEOUT / 2) * time.Second)

	if len(ap.ubuntu.mirror.Path) == 0 {
		t.Fatal("No mirrors found")
	}

	if !strings.Contains(ap.ubuntu.mirror.Path, "ubuntu") {
		t.Fatal("mirror path incorrect")
	}
}

func TestCreateNewRewritersForDebian(t *testing.T) {
	ap := *CreateNewRewriters(TYPE_LINUX_DISTROS_DEBIAN)
	time.Sleep((BENCHMARK_DETECT_TIMEOUT / 2) * time.Second)

	if len(ap.debian.mirror.Path) == 0 {
		t.Fatal("No mirrors found")
	}

	if !strings.Contains(ap.debian.mirror.Path, "debian") {
		t.Fatal("mirror path incorrect")
	}
}

func TestCreateNewRewritersForAll(t *testing.T) {
	ap := *CreateNewRewriters(TYPE_LINUX_ALL_DISTROS)
	time.Sleep((BENCHMARK_DETECT_TIMEOUT / 2) * time.Second)

	if len(ap.debian.mirror.Path) == 0 || len(ap.ubuntu.mirror.Host) == 0 {
		t.Fatal("No mirrors found")
	}

	if !strings.Contains(ap.debian.mirror.Path, "debian") || !strings.Contains(ap.ubuntu.mirror.Path, "ubuntu") {
		t.Fatal("mirror path incorrect")
	}
}

func TestCreateNewRewritersWithSpecifyMirror(t *testing.T) {
	state.SetUbuntuMirror("https://mirrors.tuna.tsinghua.edu.cn/ubuntu/")

	ap := *CreateNewRewriters(TYPE_LINUX_DISTROS_UBUNTU)
	if ap.ubuntu.mirror.Host != "mirrors.tuna.tsinghua.edu.cn" {
		t.Fatal("Mirror host incorrect")
	}

	if !strings.Contains(ap.ubuntu.mirror.Path, "ubuntu") {
		t.Fatal("mirror path incorrect")
	}

	state.ClearUbuntuMirror()
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

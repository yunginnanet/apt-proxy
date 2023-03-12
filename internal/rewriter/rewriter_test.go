package rewriter

import (
	"fmt"
	"strings"
	"testing"
	"time"

	Benchmark "github.com/soulteary/apt-proxy/internal/benchmark"
	Define "github.com/soulteary/apt-proxy/internal/define"
	State "github.com/soulteary/apt-proxy/internal/state"
)

func TestGetRewriteRulesByMode(t *testing.T) {
	rules := GetRewriteRulesByMode(Define.TypeLinuxDistrosUbuntu)
	if rules[0].OS != Define.TypeLinuxDistrosUbuntu {
		t.Fatal("Pattern Not Match")
	}

	rules = GetRewriteRulesByMode(Define.TypeLinuxDistrosDebian)
	if rules[0].OS != Define.TypeLinuxDistrosDebian {
		t.Fatal("Pattern Not Match")
	}

	rules = GetRewriteRulesByMode(Define.TypeLinuxDistrosCentos)
	if rules[0].OS != Define.TypeLinuxDistrosCentos {
		t.Fatal("Pattern Not Match")
	}

	rules = GetRewriteRulesByMode(Define.TypeLinuxDistrosAlpine)
	if rules[0].OS != Define.TypeLinuxDistrosAlpine {
		t.Fatal("Pattern Not Match")
	}

	rules = GetRewriteRulesByMode(Define.TypeLinuxAllDistros)
	if rules[0].OS != Define.TypeLinuxDistrosUbuntu {
		t.Fatal("Pattern Not Match")
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
	if !strings.Contains(strings.ToLower(ubuntu.mirror.Path), "ubuntu") {
		t.Fatal("mirror path incorrect")
	}
}

func TestCreateNewRewritersForUbuntu(t *testing.T) {
	ap := *CreateNewRewriters(Define.TypeLinuxDistrosUbuntu)
	time.Sleep((Benchmark.BenchmarkDetectTimeout / 2) * time.Second)

	if len(ap.ubuntu.mirror.Path) == 0 {
		t.Fatal("No mirrors found")
	}

	if !strings.Contains(ap.ubuntu.mirror.Path, "ubuntu") {
		t.Fatal("mirror path incorrect")
	}
}

func TestCreateNewRewritersForDebian(t *testing.T) {
	ap := *CreateNewRewriters(Define.TypeLinuxDistrosDebian)
	time.Sleep((Benchmark.BenchmarkDetectTimeout / 2) * time.Second)

	if len(ap.debian.mirror.Path) == 0 {
		t.Fatal("No mirrors found")
	}

	if !strings.Contains(ap.debian.mirror.Path, "debian") {
		t.Fatal("mirror path incorrect")
	}
}

func TestCreateNewRewritersForAll(t *testing.T) {
	ap := *CreateNewRewriters(Define.TypeLinuxAllDistros)
	time.Sleep((Benchmark.BenchmarkDetectTimeout / 2) * time.Second)

	if len(ap.debian.mirror.Path) == 0 || len(ap.ubuntu.mirror.Host) == 0 {
		t.Fatal("No mirrors found")
	}

	if !strings.Contains(ap.debian.mirror.Path, "debian") || !strings.Contains(ap.ubuntu.mirror.Path, "ubuntu") {
		t.Fatal("mirror path incorrect")
	}
}

func TestCreateNewRewritersWithSpecifyMirror(t *testing.T) {
	State.SetUbuntuMirror("https://mirrors.cn99.com/ubuntu/")

	ap := *CreateNewRewriters(Define.TypeLinuxDistrosUbuntu)
	/*if ap.ubuntu.mirror.Host != "mirrors.cn99.com" {
		t.Fatal("Mirror host incorrect")
	}*/

	if !strings.Contains(ap.ubuntu.mirror.Path, "ubuntu") {
		t.Fatal("mirror path incorrect, should contain ubuntu")
	}

	State.ResetUbuntuMirror()
}

func TestMatchingRule(t *testing.T) {
	_, match := MatchingRule("http://archive.ubuntu.com/ubuntu/InRelease", Define.UbuntuDefaultCacheRules)
	if !match {
		t.Fatal("test ap rules faild")
	}
	_, match = MatchingRule("http://cn.archive.ubuntu.com/ubuntu/InRelease", Define.UbuntuDefaultCacheRules)
	if !match {
		t.Fatal("test ap rules faild")
	}
}

func TestRuleToString(t *testing.T) {
	rule, _ := MatchingRule("http://archive.ubuntu.com/ubuntu/InRelease", Define.UbuntuDefaultCacheRules)
	fmt.Println(rule.String())
	if rule.String() != "InRelease$ Cache-Control=max-age=3600 Rewrite=true" {
		t.Fatal("test rules toString faild")
	}
}

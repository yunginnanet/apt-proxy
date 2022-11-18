package define_test

import (
	"fmt"
	"regexp"
	"testing"

	Define "github.com/soulteary/apt-proxy/internal/define"
)

func TestRuleToString(t *testing.T) {
	r := Define.Rule{
		Pattern:      regexp.MustCompile(`a$`),
		CacheControl: "1",
		Rewrite:      true,
	}

	expect := fmt.Sprintf("%s Cache-Control=%s Rewrite=%#v", r.Pattern.String(), r.CacheControl, r.Rewrite)
	if expect != r.String() {
		t.Fatal("parse rule to string failed")
	}
}

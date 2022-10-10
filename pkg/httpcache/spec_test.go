package httpcache_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/soulteary/apt-proxy/pkg/httpcache"
	"github.com/soulteary/apt-proxy/pkg/httplog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testSetup() (*client, *upstreamServer) {
	upstream := &upstreamServer{
		Body:    []byte("llamas"),
		asserts: []func(r *http.Request){},
		Now:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		Header:  http.Header{},
	}

	httpcache.Clock = func() time.Time {
		return upstream.Now
	}

	cacheHandler := httpcache.NewHandler(
		httpcache.NewMemoryCache(),
		upstream,
	)

	var handler http.Handler = cacheHandler

	if testing.Verbose() {
		rlogger := httplog.NewResponseLogger(cacheHandler)
		rlogger.DumpRequests = true
		rlogger.DumpResponses = true
		handler = rlogger
		httpcache.DebugLogging = true
	} else {
		log.SetOutput(ioutil.Discard)
	}

	return &client{handler, cacheHandler}, upstream
}

func TestSpecResponseCacheControl(t *testing.T) {
	var cases = []struct {
		cacheControl   string
		cacheStatus    string
		requests       int
		secondsElapsed time.Duration
		shared         bool
	}{
		{cacheControl: "", requests: 2},
		{cacheControl: "no-cache", requests: 2, cacheStatus: "SKIP"},
		{cacheControl: "no-store", requests: 2, cacheStatus: "SKIP"},
		{cacheControl: "max-age=0, no-cache", requests: 2, cacheStatus: "SKIP"},
		{cacheControl: "max-age=0", requests: 2, cacheStatus: "SKIP"},
		{cacheControl: "s-maxage=0", requests: 2, cacheStatus: "SKIP", shared: true},
		{cacheControl: "s-maxage=60", requests: 2, cacheStatus: "HIT", shared: true},
		{cacheControl: "s-maxage=60", requests: 2, secondsElapsed: 65, shared: true},
		{cacheControl: "max-age=60", requests: 1, cacheStatus: "HIT"},
		{cacheControl: "max-age=60", requests: 1, secondsElapsed: 35, cacheStatus: "HIT"},
		{cacheControl: "max-age=60", requests: 2, secondsElapsed: 65},
		{cacheControl: "max-age=60, must-revalidate", requests: 2, cacheStatus: "HIT"},
		{cacheControl: "max-age=60, proxy-revalidate", requests: 1, cacheStatus: "HIT"},
		{cacheControl: "max-age=60, proxy-revalidate", requests: 2, cacheStatus: "HIT", shared: true},
		{cacheControl: "private, max-age=60", requests: 1, cacheStatus: "HIT"},
		{cacheControl: "private, max-age=60", requests: 2, cacheStatus: "SKIP", shared: true},
	}

	for idx, c := range cases {
		client, upstream := testSetup()
		upstream.CacheControl = c.cacheControl
		client.cacheHandler.Shared = c.shared

		if http.StatusOK != client.get("/").Code {
			t.Fatalf("HTTP status code: %d not equal Status OK", client.get("/").Code)
		}
		upstream.timeTravel(time.Second * time.Duration(c.secondsElapsed))

		r := client.get("/")
		if http.StatusOK != r.statusCode {
			t.Fatalf("HTTP status code: %d not equal Status OK", r.statusCode)
		}

		require.Equal(t, c.requests, upstream.requests,
			fmt.Sprintf("case #%d failed, %+v", idx+1, c))

		if c.cacheStatus != "" {
			require.Equal(t, c.cacheStatus, r.cacheStatus,
				fmt.Sprintf("case #%d failed, %+v", idx+1, c))
		}
	}
}

func TestSpecResponseCacheControlWithPrivateHeaders(t *testing.T) {
	client, upstream := testSetup()
	client.cacheHandler.Shared = false
	upstream.CacheControl = `max-age=10, private=X-Llamas, private=Set-Cookie"`
	upstream.Header.Add("X-Llamas", "fully")
	upstream.Header.Add("Set-Cookie", "llamas=true")
	if http.StatusOK != client.get("/r1").Code {
		t.Fatalf("HTTP status code: %d not equal Status OK", client.get("/r1").Code)
	}

	r1 := client.get("/r1")
	if http.StatusOK != r1.statusCode {
		t.Fatalf("HTTP status code: %d not equal Status OK", r1.statusCode)
	}
	if strings.Compare("HIT", r1.cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", r1.cacheStatus)
	}
	if strings.Compare("fully", r1.HeaderMap.Get("X-Llamas")) != 0 {
		t.Fatalf("Cache status: %s not equal", r1.HeaderMap.Get("X-Llamas"))
	}
	if strings.Compare("llamas=true", r1.HeaderMap.Get("Set-Cookie")) != 0 {
		t.Fatalf("Cache status: %s not equal", r1.HeaderMap.Get("Set-Cookie"))
	}
	if upstream.requests != 1 {
		t.Fatalf("Unexpected requests: %d", upstream.requests)
	}

	client.cacheHandler.Shared = true
	if http.StatusOK != client.get("/r2").Code {
		t.Fatalf("HTTP status code: %d not equal Status OK", client.get("/r2").Code)
	}

	r2 := client.get("/r2")
	if http.StatusOK != r1.statusCode {
		t.Fatalf("HTTP status code: %d not equal Status OK", r1.statusCode)
	}
	if strings.Compare("HIT", r2.cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", r2.cacheStatus)
	}
	if r2.HeaderMap.Get("X-Llamas") != "" || r2.HeaderMap.Get("Set-Cookie") != "" {
		t.Fatalf("X-Llamas, Set-Cookie should not be empty")
	}
	if upstream.requests != 2 {
		t.Fatalf("Unexpected requests: %d", upstream.requests)
	}
}

func TestSpecResponseCacheControlWithAuthorizationHeaders(t *testing.T) {
	client, upstream := testSetup()
	client.cacheHandler.Shared = true
	upstream.CacheControl = `max-age=10`
	upstream.Header.Add("Authorization", "fully")
	if http.StatusOK != client.get("/r1").Code {
		t.Fatalf("HTTP status code: %d not equal Status OK", client.get("/r1").Code)
	}

	r1 := client.get("/r1")
	if http.StatusOK != r1.statusCode {
		t.Fatalf("HTTP status code: %d not equal Status OK", r1.statusCode)
	}
	if strings.Compare("SKIP", r1.cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", r1.cacheStatus)
	}
	if strings.Compare("fully", r1.HeaderMap.Get("Authorization")) != 0 {
		t.Fatalf("Cache status: %s not equal", r1.HeaderMap.Get("Authorization"))
	}
	if upstream.requests != 2 {
		t.Fatalf("Unexpected requests: %d", upstream.requests)
	}

	client.cacheHandler.Shared = false
	if http.StatusOK != client.get("/r2").Code {
		t.Fatalf("HTTP status code: %d not equal Status OK", client.get("/r2").Code)
	}

	r3 := client.get("/r2")
	if http.StatusOK != r3.statusCode {
		t.Fatalf("HTTP status code: %d not equal Status OK", r3.statusCode)
	}
	if strings.Compare("HIT", r3.cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", r3.cacheStatus)
	}
	if strings.Compare("fully", r3.HeaderMap.Get("Authorization")) != 0 {
		t.Fatalf("Cache status: %s not equal", r3.HeaderMap.Get("Authorization"))
	}
	if upstream.requests != 3 {
		t.Fatalf("Unexpected requests: %d", upstream.requests)
	}
}

func TestSpecRequestCacheControl(t *testing.T) {
	var cases = []struct {
		cacheControl   string
		cacheStatus    string
		requests       int
		secondsElapsed time.Duration
	}{
		{cacheControl: "", requests: 1},
		{cacheControl: "no-cache", requests: 2},
		{cacheControl: "no-store", requests: 2},
		{cacheControl: "max-age=0", requests: 2},
		{cacheControl: "max-stale", requests: 1, secondsElapsed: 65},
		{cacheControl: "max-stale=0", requests: 2, secondsElapsed: 65},
		{cacheControl: "max-stale=60", requests: 1, secondsElapsed: 65},
		{cacheControl: "max-stale=60", requests: 1, secondsElapsed: 65},
		{cacheControl: "max-age=30", requests: 2, secondsElapsed: 40},
		{cacheControl: "min-fresh=5", requests: 1},
		{cacheControl: "min-fresh=120", requests: 2},
	}

	for idx, c := range cases {
		client, upstream := testSetup()
		upstream.CacheControl = "max-age=60"

		if http.StatusOK != client.get("/").Code {
			t.Fatalf("HTTP status code: %d not equal Status OK", client.get("/").Code)
		}
		upstream.timeTravel(time.Second * time.Duration(c.secondsElapsed))

		r := client.get("/", "Cache-Control: "+c.cacheControl)
		if http.StatusOK != r.statusCode {
			t.Fatalf("HTTP status code: %d not equal Status OK", r.statusCode)
		}
		if upstream.requests != c.requests {
			t.Fatalf("case #%d failed, %+v", idx+1, c)
		}
	}
}

func TestSpecRequestCacheControlWithOnlyIfCached(t *testing.T) {
	client, upstream := testSetup()
	upstream.CacheControl = "max-age=10"

	if http.StatusOK != client.get("/").Code {
		t.Fatalf("HTTP status code: %d not equal Status OK", client.get("/").Code)
	}
	if http.StatusOK != client.get("/").Code {
		t.Fatalf("HTTP status code: %d not equal Status OK", client.get("/").Code)
	}

	upstream.timeTravel(time.Second * 20)
	if http.StatusGatewayTimeout != client.get("/", "Cache-Control: only-if-cached").Code {
		t.Fatalf("HTTP status code: %d not equal StatusGatewayTimeout", client.get("/", "Cache-Control: only-if-cached").Code)
	}

	if upstream.requests != 1 {
		t.Fatalf("Unexpected requests: %d", upstream.requests)
	}
}

func TestSpecCachingStatusCodes(t *testing.T) {
	client, upstream := testSetup()
	upstream.StatusCode = http.StatusNotFound
	upstream.CacheControl = "public, max-age=60"

	r1 := client.get("/r1")
	if http.StatusNotFound != r1.statusCode {
		t.Fatalf("HTTP status code: %d not equal Status NotFound", r1.statusCode)
	}
	if strings.Compare("MISS", r1.cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", r1.cacheStatus)
	}
	if strings.Compare(string(upstream.Body), string(r1.body)) != 0 {
		t.Fatalf("body: %s not equal", string(r1.body))
	}

	upstream.timeTravel(time.Second * 10)
	r2 := client.get("/r1")
	if http.StatusNotFound != r2.statusCode {
		t.Fatalf("HTTP status code: %d not equal Status NotFound", r2.statusCode)
	}
	if strings.Compare("HIT", r2.cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", r2.cacheStatus)
	}
	if strings.Compare(string(upstream.Body), string(r2.body)) != 0 {
		t.Fatalf("body: %s not equal", string(r2.body))
	}
	if time.Second*10 != r2.age {
		t.Fatalf("age: %d not equal", r2.age)
	}

	upstream.StatusCode = http.StatusPaymentRequired
	r3 := client.get("/r2")
	if http.StatusPaymentRequired != r3.statusCode {
		t.Fatalf("HTTP status code: %d not equal StatusPaymentRequired", r3.statusCode)
	}
	if strings.Compare("SKIP", r3.cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", r3.cacheStatus)
	}
}

func TestSpecConditionalCaching(t *testing.T) {
	client, upstream := testSetup()
	upstream.Etag = `"llamas"`

	r1 := client.get("/")
	if strings.Compare("MISS", r1.cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", r1.cacheStatus)
	}
	if strings.Compare(string(upstream.Body), string(r1.body)) != 0 {
		t.Fatalf("body: %s not equal", string(r1.body))
	}

	r2 := client.get("/", `If-None-Match: "llamas"`)
	if http.StatusNotModified != r2.Code {
		t.Fatalf("HTTP status code: %d not equal StatusNotModified", r2.Code)
	}
	if string(r2.body) != "" {
		t.Fatal("body should not be empty")
	}
	if strings.Compare("HIT", r2.cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", r2.cacheStatus)
	}
}

func TestSpecRangeRequests(t *testing.T) {
	client, upstream := testSetup()

	r1 := client.get("/", "Range: bytes=0-3")
	if http.StatusPartialContent != r1.Code {
		t.Fatalf("HTTP status code: %d not equal StatusPartialContent", r1.Code)
	}
	if strings.Compare("SKIP", r1.cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", r1.cacheStatus)
	}
	if strings.Compare(string(upstream.Body[0:4]), string(r1.body)) != 0 {
		t.Fatalf("body: %s not equal", string(r1.body))
	}
}

func TestSpecHeuristicCaching(t *testing.T) {
	client, upstream := testSetup()
	upstream.LastModified = upstream.Now.AddDate(-1, 0, 0)
	if strings.Compare("MISS", client.get("/").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.get("/").cacheStatus)
	}

	upstream.timeTravel(time.Hour * 48)
	r2 := client.get("/")
	if strings.Compare("HIT", r2.cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", r2.cacheStatus)
	}
	assert.Equal(t, []string{"113 - \"Heuristic Expiration\""}, r2.Header()["Warning"])
	assert.Equal(t, 1, upstream.requests, "The second request shouldn't validate")
}

func TestSpecCacheControlTrumpsExpires(t *testing.T) {
	client, upstream := testSetup()
	upstream.LastModified = upstream.Now.AddDate(-1, 0, 0)
	upstream.CacheControl = "max-age=2"
	if strings.Compare("MISS", client.get("/").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.get("/").cacheStatus)
	}
	if strings.Compare("HIT", client.get("/").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.get("/").cacheStatus)
	}
	if upstream.requests != 1 {
		t.Fatalf("Unexpected requests: %d", upstream.requests)
	}

	upstream.timeTravel(time.Hour * 48)
	if strings.Compare("HIT", client.get("/").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.get("/").cacheStatus)
	}
	if upstream.requests != 2 {
		t.Fatalf("Unexpected requests: %d", upstream.requests)
	}
}

func TestSpecNotCachedWithoutValidatorOrExpiration(t *testing.T) {
	client, upstream := testSetup()
	upstream.LastModified = time.Time{}
	upstream.Etag = ""

	if strings.Compare("SKIP", client.get("/").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.get("/").cacheStatus)
	}
	if strings.Compare("SKIP", client.get("/").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.get("/").cacheStatus)
	}
	if upstream.requests != 2 {
		t.Fatalf("Unexpected requests: %d", upstream.requests)
	}
}

func TestSpecNoCachingForInvalidExpires(t *testing.T) {
	client, upstream := testSetup()
	upstream.LastModified = time.Time{}
	upstream.Header.Set("Expires", "-1")

	if strings.Compare("SKIP", client.get("/").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.get("/").cacheStatus)
	}
}

func TestSpecRequestsWithoutHostHeader(t *testing.T) {
	client, _ := testSetup()

	r := newRequest("GET", "http://example.org")
	r.Header.Del("Host")
	r.Host = ""

	resp := client.do(r)
	if http.StatusBadRequest != resp.Code {
		t.Fatal("Requests without a Host header should result in a 400")
	}
}

func TestSpecCacheControlMaxStale(t *testing.T) {
	client, upstream := testSetup()
	upstream.CacheControl = "max-age=60"
	if strings.Compare("MISS", client.get("/").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.get("/").cacheStatus)
	}
	upstream.timeTravel(time.Second * 90)
	upstream.Body = []byte("brand new content")
	r2 := client.get("/", "Cache-Control: max-stale=3600")
	if strings.Compare("HIT", r2.cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", r2.cacheStatus)
	}
	if time.Second*90 != r2.age {
		t.Fatalf("age: %d not equal", r2.age)
	}

	upstream.timeTravel(time.Second * 90)
	r3 := client.get("/")
	if strings.Compare("MISS", r3.cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", r3.cacheStatus)
	}
	if time.Duration(0) != r3.age {
		t.Fatalf("age: %d not equal", r3.age)
	}
}

func TestSpecValidatingStaleResponsesUnchanged(t *testing.T) {
	client, upstream := testSetup()
	upstream.CacheControl = "max-age=60"
	upstream.Etag = "llamas1"
	if strings.Compare("MISS", client.get("/").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.get("/").cacheStatus)
	}
	upstream.timeTravel(time.Second * 90)
	upstream.Header.Add("X-New-Header", "1")

	r2 := client.get("/")
	if http.StatusOK != r2.Code {
		t.Fatalf("HTTP status code: %d not equal StatusOK", r2.Code)
	}
	if strings.Compare(string(upstream.Body), string(r2.body)) != 0 {
		t.Fatalf("body: %s not equal", string(r2.body))
	}
	if strings.Compare("HIT", r2.cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", r2.cacheStatus)
	}
}

func TestSpecValidatingStaleResponsesWithNewContent(t *testing.T) {
	client, upstream := testSetup()
	upstream.CacheControl = "max-age=60"
	if strings.Compare("MISS", client.get("/").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.get("/").cacheStatus)
	}
	upstream.timeTravel(time.Second * 90)
	upstream.Body = []byte("brand new content")

	r2 := client.get("/")
	if http.StatusOK != r2.Code {
		t.Fatalf("HTTP status code: %d not equal StatusOK", r2.Code)
	}
	if strings.Compare("MISS", r2.cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", r2.cacheStatus)
	}
	if strings.Compare("brand new content", string(r2.body)) != 0 {
		t.Fatalf("body: %s not equal", string(r2.body))
	}
	if time.Duration(0) != r2.age {
		t.Fatalf("age: %d not equal", r2.age)
	}
}

func TestSpecValidatingStaleResponsesWithNewEtag(t *testing.T) {
	client, upstream := testSetup()
	upstream.CacheControl = "max-age=60"
	upstream.Etag = "llamas1"

	if strings.Compare("MISS", client.get("/").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.get("/").cacheStatus)
	}

	upstream.timeTravel(time.Second * 90)
	upstream.Etag = "llamas2"

	r2 := client.get("/")
	if http.StatusOK != r2.Code {
		t.Fatalf("HTTP status code: %d not equal StatusOK", r2.Code)
	}
	if strings.Compare("MISS", r2.cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", r2.cacheStatus)
	}
}

func TestSpecVaryHeader(t *testing.T) {
	client, upstream := testSetup()
	upstream.CacheControl = "max-age=60"
	upstream.Vary = "Accept-Language"
	upstream.Etag = "llamas"

	if strings.Compare("MISS", client.get("/", "Accept-Language: en").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.get("/", "Accept-Language: en").cacheStatus)
	}
	if strings.Compare("HIT", client.get("/", "Accept-Language: en").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.get("/", "Accept-Language: en").cacheStatus)
	}
	if strings.Compare("MISS", client.get("/", "Accept-Language: de").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.get("/", "Accept-Language: de").cacheStatus)
	}
	if strings.Compare("HIT", client.get("/", "Accept-Language: de").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.get("/", "Accept-Language: de").cacheStatus)
	}
}

func TestSpecHeadersPropagated(t *testing.T) {
	client, upstream := testSetup()
	upstream.CacheControl = "max-age=60"
	upstream.Header.Add("X-Llamas", "1")
	upstream.Header.Add("X-Llamas", "3")
	upstream.Header.Add("X-Llamas", "2")
	if strings.Compare("MISS", client.get("/").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.get("/").cacheStatus)
	}
	r2 := client.get("/")
	if strings.Compare("HIT", r2.cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", r2.cacheStatus)
	}
	assert.Equal(t, []string{"1", "3", "2"}, r2.Header()["X-Llamas"])
}

func TestSpecAgeHeaderFromUpstream(t *testing.T) {
	client, upstream := testSetup()
	upstream.CacheControl = "max-age=86400"
	upstream.Header.Set("Age", "3600") //1hr
	if time.Hour != client.get("/").age {
		t.Fatalf("age: %d not equal", client.get("/").age)
	}

	upstream.timeTravel(time.Hour * 2)
	if time.Hour*3 != client.get("/").age {
		t.Fatalf("age: %d not equal", client.get("/").age)
	}
}

func TestSpecAgeHeaderWithResponseDelay(t *testing.T) {
	client, upstream := testSetup()
	upstream.CacheControl = "max-age=86400"
	upstream.Header.Set("Age", "3600") //1hr
	upstream.ResponseDuration = time.Second * 2
	if time.Second*3602 != client.get("/").age {
		t.Fatalf("age: %d not equal", client.get("/").age)
	}

	upstream.timeTravel(time.Second * 60)
	if time.Second*3662 != client.get("/").age {
		t.Fatalf("age: %d not equal", client.get("/").age)
	}
	if upstream.requests != 1 {
		t.Fatalf("Unexpected requests: %d", upstream.requests)
	}
}

func TestSpecAgeHeaderGeneratedWhereNoneExists(t *testing.T) {
	client, upstream := testSetup()
	upstream.CacheControl = "max-age=86400"
	upstream.ResponseDuration = time.Second * 2
	if time.Second*2 != client.get("/").age {
		t.Fatalf("age: %d not equal", client.get("/").age)
	}

	upstream.timeTravel(time.Second * 60)
	if time.Second*62 != client.get("/").age {
		t.Fatalf("age: %d not equal", client.get("/").age)
	}
	if upstream.requests != 1 {
		t.Fatalf("Unexpected requests: %d", upstream.requests)
	}
}

func TestSpecWarningForOldContent(t *testing.T) {
	client, upstream := testSetup()
	upstream.LastModified = upstream.Now.AddDate(-1, 0, 0)
	if strings.Compare("MISS", client.get("/").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.get("/").cacheStatus)
	}

	upstream.timeTravel(time.Hour * 48)
	r2 := client.get("/")
	if strings.Compare("HIT", r2.cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", r2.cacheStatus)
	}
	assert.Equal(t, []string{"113 - \"Heuristic Expiration\""}, r2.Header()["Warning"])
}

func TestSpecHeadCanBeServedFromCacheOnlyWithExplicitFreshness(t *testing.T) {
	client, upstream := testSetup()
	upstream.CacheControl = "max-age=3600"
	if strings.Compare("MISS", client.get("/explicit").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.get("/explicit").cacheStatus)
	}
	if strings.Compare("HIT", client.head("/explicit").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.head("/explicit").cacheStatus)
	}
	if strings.Compare("HIT", client.head("/explicit").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.head("/explicit").cacheStatus)
	}
	upstream.CacheControl = ""
	if strings.Compare("SKIP", client.get("/implicit").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.get("/implicit").cacheStatus)
	}
	if strings.Compare("SKIP", client.head("/implicit").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.head("/implicit").cacheStatus)
	}
	if strings.Compare("SKIP", client.head("/implicit").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.head("/implicit").cacheStatus)
	}
}

func TestSpecInvalidatingGetWithHeadRequest(t *testing.T) {
	client, upstream := testSetup()
	upstream.CacheControl = "max-age=3600"
	if strings.Compare("MISS", client.get("/explicit").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.get("/explicit").cacheStatus)
	}
	upstream.Body = []byte("brand new content")
	if strings.Compare("SKIP", client.head("/explicit", "Cache-Control: max-age=0").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.head("/explicit", "Cache-Control: max-age=0").cacheStatus)
	}
	if strings.Compare("MISS", client.get("/explicit").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.get("/explicit").cacheStatus)
	}
}

func TestSpecFresheningGetWithHeadRequest(t *testing.T) {
	client, upstream := testSetup()
	upstream.CacheControl = "max-age=3600"
	if strings.Compare("MISS", client.get("/explicit").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.get("/explicit").cacheStatus)
	}

	upstream.timeTravel(time.Second * 10)
	if time.Second*10 != client.get("/explicit").age {
		t.Fatalf("age: %d not equal", client.get("/explicit").age)
	}

	upstream.Header.Add("X-Llamas", "llamas")
	if strings.Compare("SKIP", client.head("/explicit", "Cache-Control: max-age=0").cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", client.head("/explicit", "Cache-Control: max-age=0").cacheStatus)
	}

	refreshed := client.get("/explicit")
	if strings.Compare("HIT", refreshed.cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", refreshed.cacheStatus)
	}
	if time.Duration(0) != refreshed.age {
		t.Fatalf("refreshed age: %d not equal", refreshed.age)
	}
	if strings.Compare("llamas", refreshed.header.Get("X-Llamas")) != 0 {
		t.Fatalf("Cache status: %s not equal", refreshed.header.Get("X-Llamas"))
	}
}

func TestSpecContentHeaderInRequestRespected(t *testing.T) {
	client, upstream := testSetup()
	upstream.CacheControl = "max-age=3600"

	r1 := client.get("/llamas/rock")
	if strings.Compare("MISS", r1.cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", r1.cacheStatus)
	}
	if strings.Compare(string(upstream.Body), string(r1.body)) != 0 {
		t.Fatalf("Cache body: %s not equal", string(r1.body))
	}

	r2 := client.get("/another/llamas", "Content-Location: /llamas/rock")
	if strings.Compare("HIT", r2.cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", r2.cacheStatus)
	}
	if strings.Compare(string(upstream.Body), string(r2.body)) != 0 {
		t.Fatalf("Cache body: %s not equal", string(r2.body))
	}
}

func TestSpecMultipleCacheControlHeaders(t *testing.T) {
	client, upstream := testSetup()
	upstream.Header.Add("Cache-Control", "max-age=60, max-stale=10")
	upstream.Header.Add("Cache-Control", "no-cache")

	r1 := client.get("/")
	if strings.Compare("SKIP", r1.cacheStatus) != 0 {
		t.Fatalf("Cache status: %s not equal", r1.cacheStatus)
	}
}

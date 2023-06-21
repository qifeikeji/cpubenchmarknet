package cpubenchmarknet

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/perimeterx/marshmallow"
)

func TestExtractCookie(t *testing.T) {
	server404 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		r.Header.Get("") // removes unused parameter warning
	}))
	defer server404.Close()

	for _, url := range []string{
		server404.URL,
		"a" + server404.URL,
	} {
		res, err := extractCookie(url, "")
		if res != "" {
			t.Errorf("URL: %s | Response should be empty.", url)
		}
		if err == nil {
			t.Errorf("URL: %s | Error should not be nil", url)
		}
	}
}

func TestGetJSONString(t *testing.T) {
	type testCase struct {
		response    string
		expected    string
		description string
	}
	testCases := []testCase{
		{"[", "", "Invalid JSON"},
		{"", "", "Invalid URL"},
		{"[]", "", "Content Length mismatch; unexpected EOF"},
	}
	for _, tc := range testCases {
		server200 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if tc.description == "Content Length mismatch; unexpected EOF" {
				w.Header().Add("Content-Length", "42")
			}
			w.Write([]byte(tc.response))
			r.Header.Get("") // removes unused parameter warning
		}))
		defer server200.Close()
		url := server200.URL
		if tc.description == "Invalid URL" {
			url = "a" + url
		}
		req, _ := http.NewRequest("GET", url, nil)
		res, _ := getJSONString(req)
		if res != tc.expected {
			t.Errorf("Test: %q | Output %q not equal to expected %q", tc.description,
				res, tc.expected)
		}
	}
}

func TestGetCPUMegaList(t *testing.T) {
	CPUMegaList, err := GetCPUMegaList()
	if err != nil {
		t.Fatalf("Error downloading CPUMegaList, got: %v", err)
	}

	res, err := marshmallow.Unmarshal([]byte(CPUMegaList), &struct{}{})
	if err != nil {
		t.Fatalf("Error unmarshalling CPUMegaList, got: %v", err)
	}

	numKeys := len(res)
	if numKeys != 1 {
		t.Fatalf("Expected numKeys == 1, got %d", numKeys)
	}
	entries, ok := res["data"]
	if !ok {
		t.Fatal("Failed to access value of key 'data'")
	}
	numEntries := len(entries.([]interface{}))
	minimumNumEntries := 4800
	if numEntries < minimumNumEntries {
		t.Fatalf("Expected numEntries >= %d, got %d", minimumNumEntries, numEntries)
	}
}

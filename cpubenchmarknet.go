// Package cpubenchmarknet is an unofficial library
// for downloading the CPU Mega List dataset by PassMark Software.
package cpubenchmarknet

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// extractPHPSESSID visits cpubenchmark.net, extracts
// value of `PHPSESSID` cookie, and returns it.
func extractPHPSESSID() (string, error) {
	cookieName := "PHPSESSID"
	resp, err := http.Get("https://cpubenchmark.net/CPU_mega_page.html")
	if err != nil {
		return "", err
	}
	for _, cookie := range resp.Cookies() {
		if cookie.Name == cookieName {
			return cookie.Value, nil
		}
	}
	return "", fmt.Errorf("%s cookie not found", cookieName)
}

// GetCPUMegaList downloads CPU Mega List from cpubenchmark.net and returns its contents.
func GetCPUMegaList() (string, error) {
	PHPSESSID, err := extractPHPSESSID()
	if err != nil {
		return "", err
	}

	timeNowMilli := time.Now().UnixMilli()
	req, err := http.NewRequest("GET", fmt.Sprintf("https://cpubenchmark.net/data/?_=%d", timeNowMilli), nil)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.AddCookie(&http.Cookie{
		Name:  "PHPSESSID",
		Value: PHPSESSID,
	})

	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if !json.Valid(b) {
		return "", errors.New("HTTP Response is not valid JSON")
	}
	return string(b), nil
}

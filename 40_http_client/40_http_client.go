// 40_http_client.go
// Topic: net/http client — Get, Post, custom client
//
// QUICK REQUESTS
//   resp, err := http.Get(url)
//   resp, err := http.Post(url, "application/json", body)
//   resp, err := http.PostForm(url, url.Values{"k": {"v"}})
//   defer resp.Body.Close()
//   data, _ := io.ReadAll(resp.Body)
//
// CUSTOM REQUEST
//   req, _ := http.NewRequestWithContext(ctx, "PUT", url, body)
//   req.Header.Set("Authorization", "Bearer x")
//   resp, _ := http.DefaultClient.Do(req)
//
// CUSTOM CLIENT (recommended for production)
//   client := &http.Client{
//       Timeout: 10 * time.Second,
//       Transport: &http.Transport{
//           MaxIdleConns:        100,
//           IdleConnTimeout:     90 * time.Second,
//           TLSHandshakeTimeout: 5 * time.Second,
//       },
//   }
//
// RESPONSE
//   resp.StatusCode, resp.Status, resp.Header
//   resp.Body  — io.ReadCloser, MUST close.
//
// JSON
//   var v T; json.NewDecoder(resp.Body).Decode(&v)
//
// PITFALLS
//   - Always close Body, even on errors (when err == nil).
//   - http.DefaultClient has NO timeout.
//   - For repeated requests, reuse a single Client.
//
// Run: go run 40_http_client.go    (needs internet)

package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func main() {
	client := &http.Client{Timeout: 5 * time.Second}

	// Simple GET
	resp, err := client.Get("https://httpbin.org/get?x=1")
	if err != nil {
		fmt.Println("get err:", err)
		return
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println("status:", resp.Status, "len:", len(body))

	// POST JSON
	payload := []byte(`{"name":"Ada"}`)
	resp, err = client.Post("https://httpbin.org/post", "application/json", bytes.NewReader(payload))
	if err == nil {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
		fmt.Println("post status:", resp.Status)
	}

	// PostForm
	resp, err = client.PostForm("https://httpbin.org/post", url.Values{"a": {"1"}, "b": {"2"}})
	if err == nil {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}

	// Custom request with header + context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, "GET", "https://httpbin.org/headers", nil)
	req.Header.Set("X-Custom", "go-guide")
	resp, err = client.Do(req)
	if err == nil {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Println("headers resp:", string(body)[:min(120, len(body))])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

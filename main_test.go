package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudingcity/ratelimit-server/pkg/ratelimit"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	limiter = ratelimit.New(ratelimit.Config{
		Max:     3,
		Timeout: 1,
	})

	req := httptest.NewRequest("GET", "http://example.com", nil)
	w := httptest.NewRecorder()

	index(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "3", resp.Header.Get("X-RateLimit-Limit"))
	assert.Equal(t, "2", resp.Header.Get("X-RateLimit-Remaining"))
	assert.Equal(t, "1", resp.Header.Get("X-RateLimit-Reset"))
	assert.Equal(t, "requests: 1\n", string(body))

	index(w, req)
	index(w, req)

	w = httptest.NewRecorder()
	index(w, req)
	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	assert.Equal(t, http.StatusTooManyRequests, resp.StatusCode)
	assert.Equal(t, "1", resp.Header.Get("Retry-After"))
	assert.Equal(t, "Error\n", string(body))
}

func TestRealIP(t *testing.T) {
	t.Run("X-Forwarded-For exists", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://example.com", nil)
		req.Header.Set("X-Forwarded-For", "9.9.9.9, 0.0.0.0, 0.0.0.0")
		want := "9.9.9.9"
		got := realIP(req)
		assert.Equal(t, want, got)
	})

	t.Run("Remote Addr", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://example.com", nil)
		req.RemoteAddr = "10.10.10.10:8080"
		want := "10.10.10.10"
		got := realIP(req)
		assert.Equal(t, want, got)
	})
}

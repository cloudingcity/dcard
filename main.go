package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/cloudingcity/dcard/pkg/ratelimit"
)

const (
	defaultPort        = "8080"
	defaultRateLimit   = 60
	defaultRateTimeout = 60
)

var limiter *ratelimit.Limiter

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}
	limit, _ := strconv.Atoi(os.Getenv("RATE_LIMIT"))
	if limit == 0 {
		limit = defaultRateLimit
	}
	timeout, _ := strconv.Atoi(os.Getenv("RATE_TIMEOUT"))
	if timeout == 0 {
		timeout = defaultRateTimeout
	}

	limiter = ratelimit.New(ratelimit.Config{
		Max:     limit,
		Timeout: timeout,
	})

	http.HandleFunc("/", index)

	log.Printf("Started listening on %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalln(err)
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	key := realIP(req)
	hit, err := limiter.Hit(key)
	if err != nil {
		w.Header().Set("Retry-After", strconv.Itoa(hit.ResetTime))
		w.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprintln(w, "Error")
		return
	}

	w.Header().Set("X-RateLimit-Limit", strconv.Itoa(limiter.Max))
	w.Header().Set("X-RateLimit-Remaining", strconv.Itoa(hit.Remaining))
	w.Header().Set("X-RateLimit-Reset", strconv.Itoa(hit.ResetTime))
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "requests: "+strconv.Itoa(hit.Count))
}

func realIP(req *http.Request) string {
	xForwardedFor := req.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(req.RemoteAddr)); err == nil {
		return ip
	}
	return "127.0.0.1"
}

package handlers

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/go-chi/chi/v5"
)

var allowedOrigins = []string{
	"http://localhost:5173",
	"http://localhost:8081",
	"http://localhost:80",
	"http://localhost",
}

func isOriginAllowed(origin string) bool {
	for _, allowed := range allowedOrigins {
		if origin == allowed {
			return true
		}
	}
	return false
}

func addCORSHeaders(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	if origin != "" && isOriginAllowed(origin) {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
	}
}

func createProxyWithCORS(targetURL string) *httputil.ReverseProxy {
	target, err := url.Parse(targetURL)
	if err != nil {
		panic("failed to parse service URL: " + err.Error())
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	// Store the original director
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		// Preserve the original request's path
		req.URL.Path = strings.TrimPrefix(req.URL.Path, "/users")
		req.URL.Path = strings.TrimPrefix(req.URL.Path, "/movies")
		req.URL.Path = strings.TrimPrefix(req.URL.Path, "/reviews")
	}

	// Add CORS headers to the response
	proxy.ModifyResponse = func(resp *http.Response) error {
		origin := resp.Request.Header.Get("Origin")
		if origin != "" && isOriginAllowed(origin) {
			resp.Header.Set("Access-Control-Allow-Origin", origin)
			resp.Header.Set("Access-Control-Allow-Credentials", "true")
			resp.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			resp.Header.Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		}
		return nil
	}

	return proxy
}

func MountProxies(r chi.Router) {
	userProxy := createProxyWithCORS("http://user-service:10002")
	r.Mount("/users", userProxy)

	movieProxy := createProxyWithCORS("http://movie-service:10002")
	r.Mount("/movies", movieProxy)

	reviewProxy := createProxyWithCORS("http://review-service:10002")
	r.Mount("/reviews", reviewProxy)
}

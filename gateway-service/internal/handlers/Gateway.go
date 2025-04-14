package handlers

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/go-chi/chi/v5"
)

func MountProxies(r chi.Router) {
	userServiceURL, _ := url.Parse("http://user-service:10003")
	userProxy := httputil.NewSingleHostReverseProxy(userServiceURL)
	r.Mount("/api", http.StripPrefix("/api", userProxy))

	movieServiceURL, _ := url.Parse("http://movie-service:10002")
	movieProxy := httputil.NewSingleHostReverseProxy(movieServiceURL)
	r.Mount("/movies", movieProxy)
}

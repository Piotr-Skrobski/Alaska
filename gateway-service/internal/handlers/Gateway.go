package handlers

import (
	"net/http/httputil"
	"net/url"

	"github.com/go-chi/chi/v5"
)

func MountProxies(r chi.Router) {
	userServiceURL, _ := url.Parse("http://user-service:10002")
	userProxy := httputil.NewSingleHostReverseProxy(userServiceURL)
	r.Mount("/users", userProxy)

	movieServiceURL, _ := url.Parse("http://movie-service:10002")
	movieProxy := httputil.NewSingleHostReverseProxy(movieServiceURL)
	r.Mount("/movies", movieProxy)

	reviewServiceURL, _ := url.Parse("http://review-service:10002")
	reviewProxy := httputil.NewSingleHostReverseProxy(reviewServiceURL)
	r.Mount("/reviews", reviewProxy)
}

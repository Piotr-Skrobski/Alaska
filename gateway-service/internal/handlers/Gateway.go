package handlers

import (
	"net/http/httputil"
	"net/url"

	"github.com/go-chi/chi/v5"
)

func MountProxies(r chi.Router) {
	userServiceURL, err := url.Parse("http://user-service:10002")
	if err != nil {
		panic("failed to parse user service URL: " + err.Error())
	}
	userProxy := httputil.NewSingleHostReverseProxy(userServiceURL)
	r.Mount("/users", userProxy)

	movieServiceURL, err := url.Parse("http://movie-service:10002")
	if err != nil {
		panic("failed to parse movie service URL: " + err.Error())
	}
	movieProxy := httputil.NewSingleHostReverseProxy(movieServiceURL)
	r.Mount("/movies", movieProxy)

	reviewServiceURL, err := url.Parse("http://review-service:10002")
	if err != nil {
		panic("failed to parse review service URL: " + err.Error())
	}
	reviewProxy := httputil.NewSingleHostReverseProxy(reviewServiceURL)
	r.Mount("/reviews", reviewProxy)
}

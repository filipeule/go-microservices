package main

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestRouteExists(t *testing.T) {
	testApp := Config{}

	testRoutes := testApp.routes()
	chiRoutes := testRoutes.(chi.Router)

	routes := []string{
		"/authenticate",
	}

	for _, route := range routes {
		routeExist(t, chiRoutes, route)
	}
}

func routeExist(t *testing.T, routes chi.Router, route string) {
	var found bool

	_ = chi.Walk(routes, func(method, foundRoute string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if route == foundRoute {
			found = true
		}
		return nil
	})

	if !found {
		t.Errorf("did not found %s in registered routes", route)
	}
}

package routes

import (
	"net/http"

	"github.com/ermusthofa/randomname/backend/api/middleware"
	"github.com/gorilla/mux"
)

type Route struct {
	URI     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

func Load() []Route {
	routes := randomNameRoutes
	return routes
}

func SetupRoutes(r *mux.Router) *mux.Router {

	v1 := r.PathPrefix("/v1").Subrouter()
	for _, route := range Load() {
		v1.Path(route.URI).HandlerFunc(route.Handler).Methods(route.Method)
	}

	return r

}

func SetupRoutesWithMiddleware(r *mux.Router) *mux.Router {

	v1 := r.PathPrefix("/v1").Subrouter()
	for _, route := range Load() {
		v1.Path(route.URI).HandlerFunc(
			middleware.MiddlewareJSON(route.Handler),
		).Methods(route.Method)
	}

	return r

}

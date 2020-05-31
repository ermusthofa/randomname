package router

import (
	"github.com/ermusthofa/randomname/backend/api/router/routes"
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutesWithMiddleware(r)
}

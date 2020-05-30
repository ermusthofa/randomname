package router

import (
	"github.com/ermusthofa/triviadate/backend/api/router/routes"
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutesWithMiddleware(r)
}

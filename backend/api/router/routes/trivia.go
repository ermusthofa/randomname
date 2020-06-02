package routes

import (
	"net/http"

	"github.com/ermusthofa/randomname/backend/api/controller"
)

var triviaRoutes = []Route{
	Route{
		URI:     "/trivia",
		Method:  http.MethodGet,
		Handler: controller.GetTrivia,
	},
}

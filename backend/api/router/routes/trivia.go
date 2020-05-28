package routes

import (
	"net/http"

	"github.com/ermusthofa/triviadate/backend/api/controller"
)

var triviaRoutes = []Route{
	Route{
		URI:     "/trivia/{month}/{date}",
		Method:  http.MethodGet,
		Handler: controller.GetTrivia,
	},
}

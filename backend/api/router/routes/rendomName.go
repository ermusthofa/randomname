package routes

import (
	"net/http"

	"github.com/ermusthofa/triviadate/backend/api/controller"
)

var randomNameRoutes = []Route{
	Route{
		URI:     "/randomname",
		Method:  http.MethodGet,
		Handler: controller.GetRandomName,
	},
}

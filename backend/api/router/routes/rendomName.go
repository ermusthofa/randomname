package routes

import (
	"net/http"

	"github.com/ermusthofa/randomname/backend/api/controller"
)

var randomNameRoutes = []Route{
	Route{
		URI:     "/randomname",
		Method:  http.MethodGet,
		Handler: controller.GetRandomName,
	},
	Route{
		URI:     "/healthz",
		Method:  http.MethodGet,
		Handler: controller.GetHealthz,
	},
}

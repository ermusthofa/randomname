package routes

import (
	"net/http"

	"github.com/ermusthofa/randomname/backend/api/controller"
)

var healthcheck = Route{
	URI:     "/healthz",
	Method:  http.MethodGet,
	Handler: controller.GetHealthz,
}

package controller

import (
	"net/http"

	"github.com/ermusthofa/randomname/backend/api/response"
)

func GetHealthz(w http.ResponseWriter, r *http.Request) {

	backendStatus := struct {
		Status string `json:"status"`
	}{
		Status: "healthy",
	}

	response.JSON(w, http.StatusOK, backendStatus)

}

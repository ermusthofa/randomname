package api

import (
	"net/http"

	"github.com/ermusthofa/triviadate/backend/api/router"
	"github.com/prometheus/common/log"
)

func Run() {
	r := router.New()
	log.Fatal(http.ListenAndServe(":8080", r))
}

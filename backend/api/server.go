package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"

	"github.com/ermusthofa/triviadate/backend/api/model"
	"github.com/ermusthofa/triviadate/backend/api/router"
	"github.com/prometheus/common/log"
)

func Run(runtimeConfiguration model.Config) {

	listenAddress := runtimeConfiguration.Server.ListenAddress

	fmt.Printf("\n\tListening [::]:%d\n", listenAddress)

	r := router.New()
	log.Fatal(http.ListenAndServe(":"+fmt.Sprint(listenAddress), handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET"}),
		handlers.AllowedOrigins([]string{"*"}),
	)(r)))
}

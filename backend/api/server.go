package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"

	"github.com/ermusthofa/randomname/backend/api/model"
	"github.com/ermusthofa/randomname/backend/api/router"
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

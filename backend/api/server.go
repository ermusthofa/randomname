package api

import (
	"fmt"
	"net/http"

	"github.com/ermusthofa/triviadate/backend/api/model"
	"github.com/ermusthofa/triviadate/backend/api/router"
	"github.com/prometheus/common/log"
)

func Run(runtimeConfiguration model.Config) {

	listenAddress := runtimeConfiguration.Server.ListenAddress

	fmt.Printf("\n\tListening [::]:%d\n", listenAddress)

	r := router.New()
	log.Fatal(http.ListenAndServe(":"+fmt.Sprint(listenAddress), r))
}

package main

import (
	"github.com/ermusthofa/triviadate/backend/api"
	"github.com/ermusthofa/triviadate/backend/api/config"
	"github.com/ermusthofa/triviadate/backend/api/model"
)

func main() {

	var runtimeConfiguration model.Config
	config.LoadConfiguration(&runtimeConfiguration)

	api.Run(runtimeConfiguration)
}

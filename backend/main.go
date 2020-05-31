package main

import (
	"github.com/ermusthofa/randomname/backend/api"
	"github.com/ermusthofa/randomname/backend/api/config"
	"github.com/ermusthofa/randomname/backend/api/model"
)

func main() {

	var runtimeConfiguration model.Config
	config.LoadConfiguration(&runtimeConfiguration)

	api.Run(runtimeConfiguration)
}

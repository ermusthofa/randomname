package config

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/kelseyhightower/envconfig"
	yaml "gopkg.in/yaml.v2"

	"github.com/ermusthofa/triviadate/backend/api/model"
)

// Read yaml configuration file, if there are any environment variable
// replace the value from environment variable

func readConfigFile(cfg *model.Config) {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	f, err := os.Open(path.Join(basepath, "../../config.yaml"))
	if err != nil {
		failedToLoad(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		failedToLoad(err)
	}
}

func readEnvironmentVariable(cfg *model.Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		failedToLoad(err)
	}
}

func LoadConfiguration(cfg *model.Config) {
	readConfigFile(cfg)
	readEnvironmentVariable(cfg)
}

func failedToLoad(err error) {
	log.Fatalln(err)
}

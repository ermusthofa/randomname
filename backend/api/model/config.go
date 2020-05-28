package model

type Config struct {
	Server struct {
		ListenAddress uint16 `yaml:"listenAddress" envconfig:"SERVER_LISTEN_ADDRESS"`
	} `yaml:"server"`
}

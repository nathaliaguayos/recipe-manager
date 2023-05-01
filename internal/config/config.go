package config

import (
	"github.com/kelseyhightower/envconfig"
)

const EnvPrefix = "IN" //as inside nutrition

type Config struct {
	LogLevel    string `required:"true" split_words:"true" default:"info"`
	Port        uint   `required:"true" default:"80"`
	Host        string `default:"0.0.0.0"`
	ServiceName string `required:"true" split_words:"true" default:"recipe-manager"`
}

func Get() (*Config, error) {
	cfg := Config{}
	err := envconfig.Process(EnvPrefix, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

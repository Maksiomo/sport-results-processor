package config

import (
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/adapter/pgclient"
	"sport-results-pocessor/internal/common/service/client"
	"sport-results-pocessor/internal/common/service/server"
)

type Config struct {
	Logger *logger.Config `yaml:"logger"`
	DB *pgclient.Config `yaml:"db"`

	Client client.Config `yaml:"client"`
	Server server.Config `yaml:"server"`
}

var cfg = &Config{}

func NewConfig() (*Config, error) {
	return cfg, populateConfig(cfg)
}

func populateConfig(cfg any, opts ...Option) error {
	o := &configOptions{}
	for _, opt := range opts {
		opt.apply(o)
	}

	provider, err := newProvider(o)
	if err != nil {
		return err
	}

	return provider.Populate(cfg)
}

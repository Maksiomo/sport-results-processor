package config

import (
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/internal/common/adapter/pgclient"
)

type Config struct {
	Logger *logger.Config `yaml:"logger"`
	DB *pgclient.Config `yaml:"db"`
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

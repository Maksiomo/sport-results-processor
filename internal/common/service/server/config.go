package server

import "sport-results-pocessor/internal/common/service/server/registry"

type Config struct {
	Registry *registry.Config `yaml:"registry"`
}

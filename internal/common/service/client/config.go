package client

import "sport-results-pocessor/internal/common/service/client/blockchain"

type Config struct {
	Blockchain *blockchain.Config `yaml:"blockchain"`
}

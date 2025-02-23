package controller

import "github.com/evill/learning-raft/internal/core/node"

type Config struct {
	Api   ApiConfig
	Node  node.Config
	Peers node.PeersConfig
}

type ApiConfig struct {
	InternalPort uint16
}

func (c *Config) Validate() error {
	if err := c.Node.Validate(); err != nil {
		return err
	}

	return nil
}

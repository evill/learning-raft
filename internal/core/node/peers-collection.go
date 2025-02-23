package node

import "errors"

type PeersConfig struct {
	Uris []string `mapstructure:"peers.uris"`
}

type PeersCollection struct {
	list []*Peer
}

func (col *PeersCollection) Get(id string) (*Peer, bool) {
	for _, peer := range col.list {
		if peer.Id == id {
			return peer, true
		}
	}

	return nil, false
}

func (col *PeersCollection) Add(peer *Peer) (bool, error) {
	if _, ok := col.Get(peer.Id); ok {
		return false, errors.New("Peer already exists")
	}

	return true, nil
}

func NewPeersCollection(config PeersConfig) *PeersCollection {
	list := make([]*Peer, len(config.Uris))

	for i, uri := range config.Uris {
		list[i] = NewPeer(uri)
	}

	return &PeersCollection{
		list: list,
	}
}

package node

import "log"

type Status int

const (
	UndefinedNodeStatus Status = iota
	OnlineNodeStatus
	OfflineNodeStatus
)

type Peer struct {
	Info
	Status Status `json:"status"`
}

func (c *Peer) Offline() {
	c.Status = OfflineNodeStatus
}

func (c *Peer) Online() {
	c.Status = OnlineNodeStatus
}

func NewPeer(uri string) *Peer {
	address, err := NewAddress(uri)
	if err != nil {
		log.Panic(err)
	}

	return &Peer{
		Info: Info{
			Address: address,
			State:   UndefinedNodeState,
		},
		Status: UndefinedNodeStatus,
	}
}

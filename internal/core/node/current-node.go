package node

import (
	"errors"
	"log"
	"sync"
)

type CurrentNode struct {
	Info
	Peers    *PeersCollection
	mutex    sync.RWMutex
	votedFor string
}

func (node *CurrentNode) GetTerm() Term {
	node.mutex.RLock()
	defer node.mutex.RUnlock()
	return node.Info.GetTerm()
}

func (node *CurrentNode) SetTerm(value Term) {
	node.mutex.Lock()
	defer node.mutex.Unlock()
	node.Info.SetTerm(value)
}

func (node *CurrentNode) VoteFor(candidateId string) {
	node.mutex.Lock()
	// TODO: Validate the peer in list
	node.votedFor = candidateId
	defer node.mutex.Unlock()
}

func (node *CurrentNode) HasVoted() bool {
	return node.votedFor != ""
}

func (node *CurrentNode) HasVotedFor(candidateId string) bool {
	return node.votedFor != "" && node.votedFor == candidateId
}

type Config struct {
	InternalUri string `mapstructure:"node.internal.uri"`
}

func (config Config) Validate() error {
	if config.InternalUri == "" {
		return errors.New("Internal URI of node is not configured")
	}

	return nil
}

func NewCurrentNode(nodeCfg Config, peersCfg PeersConfig) *CurrentNode {
	if err := nodeCfg.Validate(); err != nil {
		log.Panic(err)
	}

	address, err := NewAddress(nodeCfg.InternalUri)
	if err != nil {
		log.Panic(err)
	}

	return &CurrentNode{
		Info: Info{
			Address: address,
			State:   FollowerNode,
		},
		Peers: NewPeersCollection(peersCfg),
	}
}

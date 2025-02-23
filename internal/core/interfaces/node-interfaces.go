package interfaces

import "github.com/evill/learning-raft/internal/core/node"

type INodeOperations interface {
	GetNodeInfo() node.Info
}

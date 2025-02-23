package oplog

import (
	"github.com/evill/learning-raft/internal/core/node"
	"sync"
)

type Operation struct {
	index LogIndex
	term  node.Term
	data  any
}

type OpLog struct {
	node        *node.CurrentNode
	ops         []Operation
	commitIndex LogIndex
	lastApplied LogIndex
	lastIndex   LogIndex
	mutex       sync.RWMutex
}

func (log *OpLog) Append(data any) {
	log.mutex.Lock()
	defer log.mutex.Unlock()
	log.ops = append(log.ops, Operation{
		index: log.lastIndex + 1,
	})
}

func (log *OpLog) LastKey() (index LogIndex, term node.Term) {
	log.mutex.RLock()
	defer log.mutex.RUnlock()

	opCount := len(log.ops)

	if opCount == 0 {
		return 0, 0
	}

	lastOp := log.ops[opCount-1]
	return lastOp.index, lastOp.term
}

func NewOpLog(node *node.CurrentNode) *OpLog {
	return &OpLog{
		commitIndex: 0,
		lastApplied: 0,
		lastIndex:   0,
	}
}

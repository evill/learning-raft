package controller

import (
	"github.com/evill/learning-raft/internal/core/interfaces"
	"github.com/evill/learning-raft/internal/core/node"
	"github.com/evill/learning-raft/internal/core/oplog"
	"github.com/evill/learning-raft/internal/transports/http/inbound"
	"log"
	"sync"
)

type IPeerOperations interface {
	GetNodeInfo(node *node.Peer) *node.Peer
}

type NodeController struct {
	config       Config
	node         *node.CurrentNode
	peersService IPeerOperations
	apiServer    interfaces.IApiServer
	oplog        *oplog.OpLog
}

func (nc *NodeController) Start() {
	log.Println("Starting Raft node controller...")
	var wg sync.WaitGroup

	// Gather info about peers

	wg.Add(1)
	go func() {
		port := nc.config.Api.InternalPort

		nc.apiServer.Start(port)
		wg.Done()
	}()

	log.Println("Raft node controller started.")
	wg.Wait()
}

func (nc *NodeController) Stop() *NodeController {
	log.Println("Stopping Raft node controller...")
	nc.apiServer.Stop()
	log.Println("Raft node controller stopped.") // TODO: Code won't complete
	return nc
}

func (nc *NodeController) GetNodeInfo() node.Info {
	return nc.node.Info.Clone()
}

func (nc *NodeController) AppendEntries() {}

// RequestVote
/**
* @param term candidate’s term
* @param candidateId candidate requesting vote
* @param lastLogIndex index of candidate’s last log entry
* @param lastLogTerm term of candidate’s last log entry
 */
func (nc *NodeController) RequestVote(
	term node.Term,
	candidateId string,
	lastLogIndex oplog.LogIndex,
	lastLogTerm node.Term,
) (latestTerm node.Term, voteGranted bool) {
	if nodeTerm := nc.node.Info.GetTerm(); term < nodeTerm {
		return nodeTerm, false
	}

	// TODO: Double check why we reject vite fir candidate which already was voted.
	// Maybe this is second round of voting with greater term
	voteAvailable := !nc.node.HasVoted() || nc.node.HasVotedFor(candidateId)
	nodeLastLogIndex, nodeLastLogTerm := nc.oplog.LastKey()
	candidateLogEligible := lastLogTerm > nodeLastLogTerm ||
		(lastLogTerm == nodeLastLogTerm && lastLogIndex >= nodeLastLogIndex)

	//If votedFor is null or candidateId, and candidate’s log is at
	//least as up-to-date as receiver’s log, grant vote
	if voteAvailable && candidateLogEligible {
		nc.node.VoteFor(candidateId)
		// TODO: Set peer to candidate by id. Should i set peer to candidate only in case vote is given? Or it should be set to candidate anyway?
		// Is there are any case when I should use information about peer is candidate. Or list of peers and their statuses is synthetic?

		// Need to update own term?
		return term, true
	}

	return term, false
}

func NewNodeController(config Config) *NodeController {
	if err := config.Validate(); err != nil {
		log.Panic(err)
	}
	node := node.NewCurrentNode(config.Node, config.Peers)
	// Read node info from disc and CLI arguments. Possible in specific method on start
	oplog := oplog.NewOpLog(node)
	nc := &NodeController{node: node, oplog: oplog, config: config}
	apiServer := inbound.NewApiService(nc)
	nc.apiServer = apiServer

	return nc
}

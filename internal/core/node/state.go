package node

type State int8

const (
	UndefinedNodeState State = iota
	FollowerNode
	CandidateNode
	LeaderNode
)

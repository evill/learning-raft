package node

type Term int

type Info struct {
	// A unique identifier
	Id    string `json:"id"`
	State State  `json:"state"`
	// A term number
	term    Term     `json:"term"`
	Address *Address `json:"address"`
}

func (node *Info) GetTerm() Term {
	return node.term
}

func (node *Info) SetTerm(value Term) {
	node.term = value
}

func (node *Info) Clone() Info {
	return Info{
		Id:      node.Id,
		State:   node.State,
		term:    node.term,
		Address: (*node.Address).Clone(),
	}
}

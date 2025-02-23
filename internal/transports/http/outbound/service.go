package outbound

import (
	"encoding/json"
	"github.com/evill/learning-raft/internal/core/node"
	"log"
	"net/http"
	"net/url"
)

func createRequestUrl(peer *node.Peer, path string) *url.URL {
	u := peer.Address.Internal.Uri.JoinPath(path)
	// fmt.Sprintf("%s://%s/node", u.Scheme, u.Host)
	return u
}

type Service struct {
}

func (s *Service) GetInfo(peer *node.Peer) (*node.Peer, error) {
	nodeURL := createRequestUrl(peer, "/nodes/current")

	resp, err := http.Get(nodeURL.String())
	if err != nil {
		log.Println("Error calling endpoint:", err)
		return peer, err // TODO: Provide special offline error
	}
	defer resp.Body.Close()

	var info node.Info
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		log.Println("Error decoding JSON:", err)
		return peer, err // TODO: Provide special broken body error
	}

	return &node.Peer{
		Info:   info,
		Status: node.OnlineNodeStatus,
	}, nil
}

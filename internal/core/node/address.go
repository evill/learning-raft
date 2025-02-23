package node

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
)

type IfaceAddress struct {
	Uri *url.URL `json:"uri,string"`
}

func (a *IfaceAddress) GetHostAndPort() (string, uint16) {
	pieces := strings.Split(a.Uri.Host, ":")
	port, err := strconv.ParseUint(pieces[1], 10, 16)
	if err != nil {
		return pieces[0], 0
	}
	return pieces[0], uint16(port)
}

// MarshalJSON converts URL to string for JSON
func (a *IfaceAddress) MarshalJSON() ([]byte, error) {
	tempStruct := struct {
		Uri string `json:"uri"`
	}{
		Uri: a.Uri.String(),
	}

	return json.Marshal(tempStruct)
}

// UnmarshalJSON converts string back to URL
func (a *IfaceAddress) UnmarshalJSON(data []byte) error {
	var urlStr string
	if err := json.Unmarshal(data, &urlStr); err != nil {
		return err
	}
	if urlStr == "" {
		a.Uri = nil
		return nil
	}
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}
	a.Uri = parsedURL
	return nil
}

type Address struct {
	Internal *IfaceAddress `json:"internal"`
}

func (a *Address) Clone() *Address {
	return &Address{
		Internal: &IfaceAddress{
			Uri: a.Internal.Uri,
		},
	}
}

func NewAddress(uri string) (*Address, error) {
	connectionString, err := url.Parse(uri)
	if err != nil {
		return &Address{}, err
	}

	connectionString.User = nil // Not used now

	address := &Address{
		Internal: &IfaceAddress{
			Uri: connectionString,
		},
	}

	return address, nil
}

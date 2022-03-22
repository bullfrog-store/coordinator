package coordinator

import (
	"errors"
)

type Peer interface {
	Propose(m *Modify) error
	Apply() (<-chan Modify, error)
}

type CommandType uint8

const (
	PutCommand CommandType = iota
	DelCommand
)

// Modify is an encapsulation of the content of a write request
type Modify struct {
	CmdType CommandType
	Key     []byte
	Value   []byte
}

// peer is the default implementation of the Peer interface,
// and it cannot expose to user.
type peer struct {
	coor Coordinator
}

func StartPeer(c *Config) *peer {
	p := &peer{}
	p.coor = newCoor(c)
	// TODO: create a zk client and run
	return p
}

func (p *peer) Propose(m *Modify) error {
	//key, value := m.Key, m.Value
	// TODO: get version in zk to judge if our proposal is stale
	switch m.CmdType {
	case PutCommand:
		// TODO: Send data to zk
	case DelCommand:
		// TODO: Send data to zk
	default:
		return errors.New("CommandTypeError: We support only Put command and Delete command")
	}
	return nil
}

func (p *peer) Apply() (<-chan Modify, error) {
	return nil, nil
}

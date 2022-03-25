package coordinator

import (
	"errors"
)

var (
	ErrCommandType = errors.New("commandTypeError: we support only Put command and Delete command")
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

func StartPeer(c *Config) (*peer, error) {
	p := &peer{}
	p.coor = newCoor(c)
	// create a zk client and run
	err := p.coor.connect()
	return p, err
}

func (p *peer) Propose(m *Modify) error {
	switch m.CmdType {
	case PutCommand:
		return p.putToZk(m.Key, m.Value)
	case DelCommand:
		return p.delToZk(m.Key)
	default:
		return ErrCommandType
	}
	return nil
}

func (p *peer) putToZk(key []byte, val []byte) error {
	return p.coor.putToZk(key, val)
}

func (p *peer) delToZk(key []byte) error {
	return p.coor.delToZk(key)
}

func (p *peer) Apply() (<-chan Modify, error) {
	return nil, nil
}

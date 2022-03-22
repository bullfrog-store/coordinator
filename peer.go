package coordinator

import (
	"errors"
	"fmt"
	"github.com/go-zookeeper/zk"
	"time"
)

var (
	ErrCommandType       = errors.New("commandTypeError: we support only Put command and Delete command")
	ErrDelNullPath       = errors.New("path does not exist")
	ZkConnTimeOut        = time.Second * 5
	ZkKeyPrefix          = ""
	ZkAllAcl             = zk.WorldACL(zk.PermAll)
	ZkFlags        int32 = 0
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
	coor   Coordinator
	zkConn *zk.Conn
}

func StartPeer(c *Config) (*peer, error) {
	p := &peer{}
	p.coor = newCoor(c)
	// create a zk client and run
	// TODO: implement function callback
	eventCallback := zk.WithEventCallback(callback)
	var err error
	p.zkConn, _, err = zk.Connect(c.ZKAddrs, ZkConnTimeOut, eventCallback)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// TODO:Handling various callbacks
func callback(event zk.Event) {
	if event.Type == zk.EventSession {
		return
	}
	// TODO: refining this function
}

func (p *peer) Propose(m *Modify) error {
	//key, value := m.Key, m.Value
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
	path := fmt.Sprintf("%s/%s", ZkKeyPrefix, string(key))
	// check if path exists
	exist, stat, _, err := p.zkConn.ExistsW(path)
	if err != nil {
		return err
	}
	// if it does not exist, then tell zk to create this path
	if !exist {
		_, err = p.zkConn.Create(path, val, ZkFlags, ZkAllAcl)
		return err
	}
	// if it exists,then send data to zk with the version
	_, err = p.zkConn.Set(path, val, stat.Version)
	return err
}

func (p *peer) delToZk(key []byte) error {
	path := fmt.Sprintf("%s/%s", ZkKeyPrefix, string(key))
	// check if path exists
	exist, stat, _, err := p.zkConn.ExistsW(path)
	if err != nil {
		return err
	}
	if !exist {
		return ErrDelNullPath
	}
	// if path exists,then delete the data related to key from zk
	err = p.zkConn.Delete(path, stat.Version)
	return err
}

func (p *peer) Apply() (<-chan Modify, error) {
	return nil, nil
}

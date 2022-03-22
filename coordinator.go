package coordinator

import (
	"errors"
	"fmt"
	"github.com/go-zookeeper/zk"
	"time"
)

var (
	ErrDelNullPath       = errors.New("path does not exist")
	ZkConnTimeOut        = time.Second * 5
	ZkKeyPrefix          = ""
	ZkAllAcl             = zk.WorldACL(zk.PermAll)
	ZkFlags        int32 = 0
)

type Coordinator interface {
	putToZk(key []byte, val []byte) error
	delToZk(key []byte) error
	connect() error
}

// coor is the default implementation of the Coordinator interface,
// and it cannot expose to user.
type coor struct {
	id      string
	zkAddrs []string
	zkConn  *zk.Conn
}

func newCoor(c *Config) *coor {
	coor := &coor{
		id:      c.ID,
		zkAddrs: c.ZKAddrs,
	}
	// TODO(qyl): finish coor struct
	return coor
}

func (c *coor) connect() error {
	// TODO: implement function callback
	eventCallback := zk.WithEventCallback(callback)
	var err error
	c.zkConn, _, err = zk.Connect(c.zkAddrs, ZkConnTimeOut, eventCallback)
	return err
}

// TODO:Handling various callbacks
func callback(event zk.Event) {
	if event.Type == zk.EventSession {
		return
	}
	// TODO: refining this function
}

func (c *coor) putToZk(key []byte, val []byte) error {
	path := fmt.Sprintf("%s/%s", ZkKeyPrefix, string(key))
	// check if path exists
	exist, stat, _, err := c.zkConn.ExistsW(path)
	if err != nil {
		return err
	}
	// if it does not exist, then tell zk to create this path
	if !exist {
		_, err = c.zkConn.Create(path, val, ZkFlags, ZkAllAcl)
		return err
	}
	// if it exists,then send data to zk with the version
	_, err = c.zkConn.Set(path, val, stat.Version)
	return err
}

func (c *coor) delToZk(key []byte) error {
	path := fmt.Sprintf("%s/%s", ZkKeyPrefix, string(key))
	// check if path exists
	exist, stat, _, err := c.zkConn.ExistsW(path)
	if err != nil {
		return err
	}
	if !exist {
		return ErrDelNullPath
	}
	// if path exists,then delete the data related to key from zk
	err = c.zkConn.Delete(path, stat.Version)
	return err
}

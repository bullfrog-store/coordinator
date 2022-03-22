package coordinator

type Coordinator interface {
}

// coor is the default implementation of the Coordinator interface,
// and it cannot expose to user.
type coor struct {
	id      string
	peers   []string
	zkAddrs []string
}

func newCoor(c *Config) *coor {
	coor := &coor{
		id:      c.ID,
		peers:   c.Peers,
		zkAddrs: c.ZKAddrs,
	}
	// TODO(qyl): finish coor struct
	return coor
}

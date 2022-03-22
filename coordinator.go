package coordinator

type Coordinator interface {
}

// coor is the default implementation of the Coordinator interface,
// and it cannot expose to user.
type coor struct {
	id      string
	zkAddrs []string
}

func newCoor(c *Config) *coor {
	coor := &coor{
		id:      c.ID,
		zkAddrs: c.ZKAddrs,
	}
	// TODO(qyl): finish coor struct
	return coor
}

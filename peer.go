package coordinator

type Peer interface {
}

func StartPeer(c *Config) Peer {
	p := &peer{}
	p.coor = newCoor(c)
	// TODO(qyl): finish peer struct
	return p
}

// peer is the default implementation of the Peer interface,
// and it cannot expose to user.
type peer struct {
	coor Coordinator
}

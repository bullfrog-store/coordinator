package coordinator

type Config struct {
	// ID is the IP of current peer.
	ID string
	// Peers are the IPs of the cluster peers except me.
	Peers []string
	// ZKAddrs are the IPs of the ZK cluster.
	ZKAddrs []string
}

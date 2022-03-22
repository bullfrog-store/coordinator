package coordinator

type Config struct {
	// ID is the IP of current peer.
	ID string
	// ZKAddrs are the IPs of the ZK cluster.
	ZKAddrs []string
}

package config

type P2P struct {
	Name string `json:"Name""`

	Sig string `json:"Sig"`

	// use for sign data
	PrivateKey string `json:"PrivateKey"`
	// use for NodeID
	PublicKey string `json:"PublicKey"`

	// `MaxPeers` is the maximum number of peers that can be connected.
	MaxPeers uint32 `json:"MaxPeers"`

	// `MaxPassivePeersRatio` is the ratio of MaxPeers that initiate an active connection to this node.
	// the actual value is `MaxPeers / MaxPassivePeersRatio`
	MaxPassivePeersRatio uint32 `json:"MaxPassivePeersRatio"`

	// `MaxPendingPeers` is the maximum number of peers that wait to connect.
	MaxPendingPeers uint32 `json:"MaxPendingPeers"`

	BootNodes []string `json:"BootNodes"`

	Addr string `json:"Addr"`

	Datadir string `json:"Datadir"`

	NetID uint `json:"NetID"`
}

func MergeP2PConfig(cfg *P2P) P2P {
	p2p := GlobalConfig.P2P

	if cfg == nil {
		return p2p
	}

	if cfg.Name != "" {
		p2p.Name = cfg.Name
	}

	if cfg.Sig != "" {
		p2p.Sig = cfg.Sig
	}

	if cfg.PrivateKey != "" {
		p2p.PrivateKey = cfg.PrivateKey
	}

	if cfg.PublicKey != "" {
		p2p.PublicKey = cfg.PublicKey
	}

	if cfg.MaxPeers != 0 {
		p2p.MaxPeers = cfg.MaxPeers
	}

	if cfg.MaxPassivePeersRatio != 0 {
		p2p.MaxPassivePeersRatio = cfg.MaxPassivePeersRatio
	}

	if cfg.MaxPendingPeers != 0 {
		p2p.MaxPendingPeers = cfg.MaxPendingPeers
	}

	if cfg.BootNodes != nil {
		p2p.BootNodes = cfg.BootNodes
	}

	if cfg.Addr != "" {
		p2p.Addr = cfg.Addr
	}

	if cfg.Datadir != "" {
		p2p.Datadir = cfg.Datadir
	}

	if cfg.NetID != 0 {
		p2p.NetID = cfg.NetID
	}

	return p2p
}

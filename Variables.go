package assignment03IBC

import (
	"net"
	"sync"

	a2 "github.com/Qureshi-DH/assignment02IBC"
)

// Quorum stores the peer limit
var Quorum int

const (
	bufferSize = 4096
)

var peerList map[string]net.Conn = make(map[string]net.Conn)
var listeningAddresses []string
var mutex sync.Mutex
var chainHead *a2.Block

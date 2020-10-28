package assignment03IBC

import (
	"encoding/gob"
	"log"
	"net"
	"sync"

	a2 "github.com/Qureshi-DH/assignment02IBC"
)

var peerList map[string]net.Conn = make(map[string]net.Conn)
var listeningAddresses []string
var mutex sync.Mutex
var chainHead *a2.Block

func handleConnection(conn net.Conn, name string) {
	if name == "satoshi" {

		receivedAddress := ReadString(conn)

		mutex.Lock()

		listeningAddresses = append(listeningAddresses, receivedAddress)
		peerList[conn.RemoteAddr().String()] = conn
		chainHead = a2.InsertBlock("", "", "Satoshi", 0, chainHead)

		mutex.Unlock()

	} else if name == "others" {
		message := ReadString(conn)
		log.Println(message)
	}
}

// StartListening initializes a server
func StartListening(address string, name string) {
	ln, err := net.Listen("tcp", address)

	if err != nil {
		log.Println("Couldn't start the listening server")
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Couldn't accept the connection request from " + conn.LocalAddr().String())
		}
		go handleConnection(conn, name)
	}
}

// SendChainandConnInfo send in the information to the peer
func SendChainandConnInfo() {

	peers := getKeys(&peerList)

	for i := 0; i < len(peers); i++ {

		conn := peerList[peers[i]]
		conn.Write([]byte(listeningAddresses[(i+1)%Quorum]))

		gobEncoder := gob.NewEncoder(conn)

		if err := gobEncoder.Encode(chainHead); err != nil {
			log.Println("Failed to encode blockchain")
		}
	}
}

// WaitForQuorum ...
func WaitForQuorum() {
	for len(listeningAddresses) < Quorum {

	}
}

// ReadString reads a string from a connection
func ReadString(conn net.Conn) string {
	recvdSlice := make([]byte, bufferSize)
	n, _ := conn.Read(recvdSlice)
	return string(recvdSlice[:n])
}

// WriteString writes a string to a connection
func WriteString(conn net.Conn, text string) {
	conn.Write([]byte(text))
}

// ReceiveChain is used to deserialize chain from a network connection
func ReceiveChain(conn net.Conn) a2.Block {

	chainHead := a2.Block{}

	gobDecoder := gob.NewDecoder(conn)

	if err := gobDecoder.Decode(&chainHead); err != nil {
		log.Println("Failed to decode blockchain")
	}

	return chainHead
}

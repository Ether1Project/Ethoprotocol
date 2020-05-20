package main

import (
	"context"
	"fmt"
	"github.com/ipfs/go-ipfs-api"
	"log"
	"math/rand"
	"strings"
	"time"
)

var StalePeerArrayFullAddress []string
var StalePeerArray []string
var BootnodePeers []string
var PeerCount = int(0)

// PeerManagement is main entry point for all peer management
func PeerManagement() {
	for {
		RandomPeerSelection()
		randomWait := rand.Intn(1200)
		time.Sleep(time.Duration(randomWait) * time.Second)
	}
}

// ConnectToBootnodePeers is used as network entry point using dynamic bootnodes on ETHO network
func ConnectToBootnodePeers() {
	sh := shell.NewShell("localhost:" + apiPort)
	var localBootnodePeers = BootnodePeers
	if len(BootnodePeers) > 0 {
		for k := range localBootnodePeers {
			err := sh.SwarmConnect(context.Background(), localBootnodePeers[k])
			if err != nil && verboseFlag {
				log.Println("Error.... Failure To Connect To ethoFS Peer: ", err)
			} else if verboseFlag {
				fmt.Println("Swarm Connection Successful - ethoFS Peer List Updated")
			}
		}
	}
}

// RandomPeerSelection is used to keep network connections random
func RandomPeerSelection() {
	sh := shell.NewShell("localhost:" + apiPort)
	var swarmConnectionCount = 0
	for k := range NodeHealthConsensus {
		ID := strings.Split(NodeHealthConsensus[k].NodeID, "/")
		if len(ID) == 7 && (BlockHeight-NodeHealthConsensus[k].LatestBlock) < 50 {
			err := sh.SwarmConnect(context.Background(), NodeHealthConsensus[k].NodeID)
			if err != nil && verboseFlag {
				fmt.Println("Error.... Failure To Connect To ethoFS Peer: ", err)
			} else {
				fmt.Println("Swarm Connection Successful - ethoFS Peer Randomization Complete")
				swarmConnectionCount++
				if swarmConnectionCount >= 10 {
					return
				}
			}
		}
	}
}

// ConnectToPeers initiates a peer connection to a group of peers
func ConnectToPeers(PeerArray []string) {
	time.Sleep(120 * time.Second)
	for {
		sh := shell.NewShell("localhost:" + apiPort)
		var localPeerArray = PeerArray
		if len(localPeerArray) > 0 {
			for k := range localPeerArray {
				err := sh.SwarmConnect(context.Background(), localPeerArray[k])
				if err != nil && verboseFlag {
					log.Println("Error.... Failure To Connect To Peer (PeerArray): ", err)
				} else if verboseFlag {
					fmt.Println("Swarm Connection Successful (PeerArray)")
				}
			}
		}
		randomWait := rand.Intn(30)
		time.Sleep(time.Duration(randomWait) * time.Minute)
	}
}

// SingleConnectToBootnodePeers is used to find network connections when needed
func SingleConnectToBootnodePeers() {
	sh := shell.NewShell("localhost:" + apiPort)
	var localBootnodePeers = BootnodePeers
	if len(localBootnodePeers) > 0 {
		for k := range localBootnodePeers {
			err := sh.SwarmConnect(context.Background(), localBootnodePeers[k])
			if err != nil && verboseFlag {
				log.Println("Error.... Failure To Connect To ethoFS Peer: ", err)
			} else if verboseFlag {
				fmt.Println("Swarm Connection Successful - ethoFS Peer List Updated")
			}
		}
	}
}

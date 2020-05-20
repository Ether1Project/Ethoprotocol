package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/ipfs/go-ipfs-api"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"log"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type NodeHealth struct {
	NodeID                       string
	GNEligible                   string
	MNEligible                   string
	SNEligible                   string
	LatestBlock                  int
	ActiveHistory                int
	LastHistoryUpdateBlockHeight int
	MaxStorage                   int
	UsedStorage                  int
	PositiveCount                int
	NegativeCount                int
}

// BroadcastHealthConsensus uses ipfs pubsub to communicate node health
func BroadcastHealthConsensus() {
	time.Sleep(30 * time.Second)
	for {
		sh := shell.NewShell("localhost:" + apiPort)
		var buffer bytes.Buffer
		for k := range NodeHealthConsensus {
			buffer.WriteString(NodeHealthConsensus[k].NodeID + ":" + NodeHealthConsensus[k].GNEligible + ":" + NodeHealthConsensus[k].MNEligible + ":" + NodeHealthConsensus[k].SNEligible + ":" + strconv.Itoa(NodeHealthConsensus[k].LatestBlock) + ":" + strconv.Itoa(NodeHealthConsensus[k].ActiveHistory) + ":" + strconv.Itoa(NodeHealthConsensus[k].LastHistoryUpdateBlockHeight) + ":" + strconv.Itoa(NodeHealthConsensus[k].MaxStorage) + ":" + strconv.Itoa(NodeHealthConsensus[k].UsedStorage) + ",")
		}
		HealthConsensusString := buffer.String()
		LastHealthMessage = HealthConsensusString
		if 2 > PeerCount {
			GetBootnodeContractValues()
			SingleConnectToBootnodePeers()
		}
		fmt.Println("Synced ethoFS Peers: " + strconv.Itoa(len(NodeHealthConsensus)))
		HealthConsensusString = strings.TrimSuffix(HealthConsensusString, ",")
		EncryptedVerificationNodeID := selfNodeID
		EncryptedHealthConsensusString := (EncryptedVerificationNodeID) + ";" + (HealthConsensusString)
		swarmPeers, _ := sh.SwarmPeers(context.Background())
		peerInfo := swarmPeers.Peers
		for p := range peerInfo {
			var HealthConsensusBroadcastChannel = peerInfo[p].Peer + "_alpha11"
			sh.PubSubPublish(HealthConsensusBroadcastChannel, EncryptedHealthConsensusString)
		}
		randomWait := rand.Intn(30)
		time.Sleep(time.Duration(randomWait) * time.Second)
	}
}

func HealthConsensusSubscribe() {
	for {
		sh := shell.NewShell("localhost:" + apiPort)
		sub, err := sh.PubSubSubscribe(selfNodeIDHashOnly + "_alpha11")
		if err != nil {
			log.Println("Error.... Health Consensus Broadcast Channel Subscription Failure (In-Bound Messaging): ", err)
		} else {
			r, _ := sub.Next()
			HealthConsensusMessage = ByteSliceToString(r.Data)
			sub.Cancel()
		}
		time.Sleep(30 * time.Second)
	}
}

func GetHealthConsensus() {
	for {
		if len(HealthConsensusMessage) > 0 {
			message := HealthConsensusMessage
			EncryptedMessageArray := strings.Split(message, ";")
			DecryptedMessage := EncryptedMessageArray[1]
			message = strings.TrimSuffix(DecryptedMessage, ",")
			messageArray := strings.Split(DecryptedMessage, ",")
			if len(messageArray) > 0 {
				CheckHealthConsensus(messageArray)
			}
		}
		time.Sleep(30 * time.Second)
	}
}

func CheckHealthConsensus(messageArray []string) {
	if 0 >= len(messageArray) {
		return
	}
	var PeerArray []string
	var PeerArrayHashOnly []string
	var UpdatedMaxStorage = int(0)
	var UpdatedUsedStorage = int(0)
	var UpdatedPeerCount = int(0)
	for k := range messageArray {
		elementArray := strings.Split(messageArray[k], ":")
		if len(elementArray) == 9 {
			elementBlockHeight, _ := strconv.Atoi(elementArray[4])
			elementActivity, _ := strconv.Atoi(elementArray[5])
			elementLastHistoryUpdateBlockHeight, _ := strconv.Atoi(elementArray[6])
			elementMaxStorage, _ := strconv.Atoi(elementArray[7])
			elementUsedStorage, _ := strconv.Atoi(elementArray[8])

			if elementActivity >= 95 {
				UpdatedPeerCount++
			}
			UpdatedMaxStorage += elementMaxStorage
			UpdatedUsedStorage += elementUsedStorage

			if mapValue, ok := NodeHealthConsensus[elementArray[0]]; ok {
				//ADD STALE NODE FOR CONNECTION TRY
				if (BlockHeight - elementBlockHeight) > 50 {
					NodeIDString := strings.Split(mapValue.NodeID, "/")
					PeerArrayHashOnly = append(PeerArrayHashOnly, NodeIDString[6])
					PeerArray = append(PeerArray, mapValue.NodeID)
				}
				//UPDATE GENERICS
				if elementBlockHeight > mapValue.LatestBlock {
					mapValue.GNEligible = elementArray[1]
					mapValue.MNEligible = elementArray[2]
					mapValue.SNEligible = elementArray[3]
					mapValue.LatestBlock = elementBlockHeight
					mapValue.UsedStorage = elementUsedStorage
					mapValue.MaxStorage = elementMaxStorage
				}
				//CHECK FOR OUT OF SYNC ACTIVITY
				if mapValue.ActiveHistory != elementActivity {
					if mapValue.NegativeCount > mapValue.PositiveCount && mapValue.NegativeCount > (PeerCount/2) {
						mapValue.ActiveHistory = elementActivity
						mapValue.LastHistoryUpdateBlockHeight = elementLastHistoryUpdateBlockHeight
						mapValue.GNEligible = elementArray[1]
						mapValue.MNEligible = elementArray[2]
						mapValue.SNEligible = elementArray[3]
						mapValue.LatestBlock = elementBlockHeight
						mapValue.UsedStorage = elementUsedStorage
						mapValue.MaxStorage = elementMaxStorage
						mapValue.NegativeCount = 0
						mapValue.PositiveCount = 0
					} else {
						mapValue.NegativeCount++
					}
				} else if mapValue.ActiveHistory == -1 {
					mapValue.ActiveHistory = 100
					mapValue.LastHistoryUpdateBlockHeight = BlockHeight
					mapValue.NegativeCount = 0
					mapValue.PositiveCount = 0
				} else if mapValue.NegativeCount != 0 {
					if mapValue.PositiveCount > mapValue.NegativeCount && mapValue.PositiveCount > (PeerCount/2) {
						mapValue.PositiveCount = 0
						mapValue.NegativeCount = 0
					} else {
						mapValue.PositiveCount++
					}
				} else if (200 < (BlockHeight - mapValue.LastHistoryUpdateBlockHeight)) && (150 >= (BlockHeight - mapValue.LatestBlock)) {
					if mapValue.NodeID != selfNodeID && mapValue.ActiveHistory < 100 {
						mapValue.ActiveHistory += 3
						mapValue.LastHistoryUpdateBlockHeight = BlockHeight
					}
					if mapValue.NodeID != selfNodeID && mapValue.ActiveHistory >= 100 {
						mapValue.ActiveHistory = 100
						mapValue.LastHistoryUpdateBlockHeight = BlockHeight
					}
				} else if (200 < (BlockHeight - mapValue.LastHistoryUpdateBlockHeight)) && (150 < (BlockHeight - mapValue.LatestBlock)) {
					if mapValue.NodeID != selfNodeID {
						mapValue.ActiveHistory -= 3
						mapValue.LastHistoryUpdateBlockHeight = BlockHeight
					}
				}
				//FINALIZE AND UPDATE ALL DATA
				if mapValue.NodeID != selfNodeID && mapValue.ActiveHistory < -1 {
					delete(NodeHealthConsensus, elementArray[0])
					fmt.Println("Stale ethoFS Node Removed From Peer Consensus: " + elementArray[0])
				} else {
					NodeHealthConsensus[elementArray[0]] = mapValue
				}
			} else if (elementActivity == -1 || elementActivity >= 3) && mapValue.NodeID != selfNodeID {
				newHealthConsensusStruct := NodeHealth{elementArray[0], elementArray[1], elementArray[2], elementArray[3], elementBlockHeight, elementActivity, elementLastHistoryUpdateBlockHeight, elementMaxStorage, elementUsedStorage, 0, 0}
				NodeHealthConsensus[elementArray[0]] = newHealthConsensusStruct
				SplitString := strings.Split(elementArray[0], "/")
				fmt.Println("ethoFS Peer Found: " + SplitString[6])
			}
		} //END OF LENGTH CHECK
	} //END OF OUTER LOOP
	StalePeerArray = PeerArrayHashOnly
	StalePeerArrayFullAddress = PeerArray
	PeerCount = UpdatedPeerCount
	if UpdatedMaxStorage > MaxStorage {
		MaxStorage = UpdatedMaxStorage
	}
	if UpdatedUsedStorage > UsedStorage {
		UsedStorage = UpdatedUsedStorage
	}
}

func UpdateSelfHealthConsensus() {
	if mapValue, ok := NodeHealthConsensus[selfNodeID]; ok {
		mapValue.GNEligible = gnEligibility
		mapValue.MNEligible = mnEligibility
		mapValue.SNEligible = snEligibility
		mapValue.LatestBlock = BlockHeight
		mapValue.MaxStorage = selfMaxStorage
		mapValue.UsedStorage = selfUsedStorage
		NodeHealthConsensus[selfNodeID] = mapValue
		NodeHealthConsensus[selfNodeID].PrintNodeHealth()
	} else {
		newHealthConsensusStruct := NodeHealth{selfNodeID, gnEligibility, mnEligibility, snEligibility, BlockHeight, -1, -1, selfMaxStorage, selfUsedStorage, 0, 0}
		NodeHealthConsensus[selfNodeID] = newHealthConsensusStruct
		NodeHealthConsensus[selfNodeID].PrintNodeHealth()
	}
}

func (n NodeHealth) PrintNodeHealth() {
	ActiveHistoryString := strconv.Itoa(n.ActiveHistory)
	ID := strings.Split(n.NodeID, "/")
	if ActiveHistoryString == "-1" {
		ActiveHistoryString = "N/A"
		fmt.Println("Node ID: " + ID[6] + "  Node Health: " + ActiveHistoryString)
	} else {
		fmt.Println("Node ID: " + ID[6] + "  Node Health: " + ActiveHistoryString + "%")
	}

	red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)

	v, _ := mem.VirtualMemory()
	path := "/"
	if runtime.GOOS == "windows" {
		path = "C:"
	}
	d, _ := disk.Usage(path)

	if masterNodeFlag {
		if (float64(v.Total) / float64(GB)) < float64(1.5) {
			boldRed.Println("Masternode does not have enough RAM")
		}
		if (float64(d.Total) / float64(GB)) < float64(38.00) {
			boldRed.Println("Masternode does not have enough Disk space")
		}
	} else if serviceNodeFlag {
		if (float64(v.Total) / float64(GB)) < float64(0.7) {
			boldRed.Println("Service Node does not have enough RAM")
		}
		if (float64(d.Total) / float64(GB)) < float64(18.00) {
			boldRed.Println("Service Node does not have enough Disk Space")
		}
	} else if gatewayNodeFlag {
		if (float64(v.Total) / float64(GB)) < float64(3.0) {
			boldRed.Println("Gateway Node does not have enough RAM")
		}
		if (float64(d.Total) / float64(GB)) < float64(70.00) {
			boldRed.Println("Gateway Node does not have enough Disk Space")
		}
	}
}

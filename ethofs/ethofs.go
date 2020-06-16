package ethofs

import (
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/ipfs/go-ipfs/core"
)

var selfNodeID string
var repFactor int
var BlockHeight = int(0)

var testHash = "QmdP3gTCyZwR4F8Kf5qFH6JovbXVXhLM7XiCRqnsTY5dHG"
var Node *core.IpfsNode

func InitializeEthofs(nodeType string) {
	log.Info("Starting ethoFS node initialization", "type", nodeType)

	initializeEthofsNode()
	selfPins = make(map[string]string)
	PinCounts = make(map[string]int)
	ErroredPinsMap = make(map[string]string)
	NodePinningConsensus = make(map[string]NodePins)
	AssignNodeID()                    //SET NODE ID FOR BROADCAST

	go NewBlock()
}

func NewBlock() {
	for {
		time.Sleep(30 * time.Second)
		//UpdateContractBootnodeValues() //GET BOOTNODE PEERS AND OTHER CONTRACT VALUES
		//UpdateContractPinValues()      //GET PIN CONTRACT VALUES
		providerCount := FindProvs(Node, testHash)
		log.Info("ethoFS provider search successful", "count", providerCount, "hash", testHash)
	}
}

// Find and store ipfs node id
func AssignNodeID() {
	/*var ipfsId string
	for {
		id, err := GetIpfsId()
		if err == nil {
			if verboseFlag {
				fmt.Println("ethoFS ID Found: " + id)
			}
			ipfsId = id
			break
		} else {
			fmt.Println("Error: Unable to Connect to ethoFS - Waiting To Retry")
			time.Sleep(10 * time.Second)
		}
	}
	if ipfsId == "" {
		log.Fatal("Error: Unable To Obtain ethoFS Node ID - Restart Node")
	}
	selfNodeIDHashOnly = ipfsId
	consensus := externalip.DefaultConsensus(nil, nil)
	ip, err := consensus.ExternalIP()
	if err == nil {
		ipAddressString := ip.String()
		selfNodeID = "/ip4/" + ipAddressString + "/tcp/" + swarmPort + "/ipfs/" + selfNodeIDHashOnly
	} else {
		log.Fatal("Error: Unable to Get External IP Address - Exiting Program\n")
	}
	go HealthConsensusSubscribe()*/
}

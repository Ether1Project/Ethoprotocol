package ethofs

import (
	"github.com/ethereum/go-ethereum/log"
)

var selfNodeID string
var repFactor int
var BlockHeight = int(0)

func InitializeEthofs(nodeType string) {
	initializeEthofsNode()
	selfPins = make(map[string]string)
	PinCounts = make(map[string]int)
	ErroredPinsMap = make(map[string]string)
	NodePinningConsensus = make(map[string]NodePins)

	log.Info("ethoFS Enabled Node Deployment Initiated", "type", nodeType)

	AssignNodeID()                    //SET NODE ID FOR BROADCAST
}

func newBlock() {

	UpdateContractBootnodeValues() //GET BOOTNODE PEERS AND OTHER CONTRACT VALUES
	UpdateContractPinValues()      //GET PIN CONTRACT VALUES

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

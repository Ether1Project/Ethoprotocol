package ethofs

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/glendc/go-external-ip"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"
)

// ethoFS Logo
const ethoFSLogo = `
|
|         __  __          ___________
|   ___  / /_/ /_  ____  / ____/ ___/
|  / _ \/ __/ __ \/ __ \/ /_   \__ \
| /  __/ /_/ / / / /_/ / __/  ___/ /
| \___/\__/_/ /_/\____/_/    /____/
|                 Powered By Ether-1
|
|
`

var selfNodeID string
var repFactor int
var BlockHeight = int(0)

func initializeEthofs() {
	//OUTPUT LOGO
	fmt.Println(ethoFSLogo)
	initializeEthofsNode()
	selfPins = make(map[string]string)
	PinCounts = make(map[string]int)
	ErroredPinsMap = make(map[string]string)
	NodePinningConsensus = make(map[string]NodePins)

	AssignNodeID()                    //SET NODE ID FOR BROADCAST
	if gatewayNodeFlag || masterNodeFlag {
		//ScanPinList()
	}
}

func newBlock() {

	UpdateContractBootnodeValues() //GET BOOTNODE PEERS AND OTHER CONTRACT VALUES
	UpdateContractPinValues()      //GET PIN CONTRACT VALUES

	for {
		info := color.New(color.FgBlue, color.Bold).SprintFunc()

		latestBlock := GetLatestEther1Block()
		tempBlockHeight, err := strconv.Atoi(latestBlock)
		if err != nil {
			log.Println("Error... Retrieving Ether-1 Network Block Height (During Conversion)", err)
		} else {
			BlockHeight = tempBlockHeight
		}
		fmt.Println("Current ETHO Block Height:", info(latestBlock))

		gnEligibility = "No"
		mnEligibility = "No"
		snEligibility = "No"
		v, _ := mem.VirtualMemory()
		path := "/"
		if runtime.GOOS == "windows" {
			path = "C:"
		}
		d, _ := disk.Usage(path)
		if (float64(v.Total)/float64(GB)) > float64(3.00) && (float64(d.Total)/float64(GB)) > float64(70.00) && gatewayTest == "Yes" {
			gnEligibility = "Yes"
			mnEligibility = "Yes"
			snEligibility = "Yes"
		} else if (float64(v.Total)/float64(GB)) > float64(1.5) && (float64(d.Total)/float64(GB)) > float64(38.00) {
			gnEligibility = "No"
			mnEligibility = "Yes"
			snEligibility = "Yes"
		} else if (float64(v.Total)/float64(GB)) > float64(0.75) && (float64(d.Total)/float64(GB)) > float64(18.00) {
			gnEligibility = "No"
			mnEligibility = "No"
			snEligibility = "Yes"
		}

		UpdateSelfHealthConsensus()

		HealthCheckMessage = "ETHOBlockHeight:" + latestBlock + ",MNEligibility:" + mnEligibility + ",SNEligibility:" + snEligibility
		randomWait := rand.Intn(10)
		time.Sleep(time.Duration(randomWait) * time.Second)
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

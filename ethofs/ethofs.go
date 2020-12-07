package ethofs

import (
	//"fmt"
	"math/rand"
	"os"
	"runtime"
	//"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"

	icore "github.com/ipfs/interface-go-ipfs-core"

	"github.com/ipfs/go-ipfs/core"
	//cid "github.com/ipfs/go-cidutil"
)

var localPinMapping map[string]string
var selfNodeID string
var repFactor = uint64(10)
var BlockHeight = int(0)
var Ipfs icore.CoreAPI
var Node *core.IpfsNode
var contractControllerAddress = common.HexToAddress("0xc38B47169950D8A28bC77a6Fa7467464f25ADAFc")
var mainChannelString = "ethoFSPinningChannel_alpha11"
var defaultDataDir string
var ipcLocation string
var isInitialized = false

func IsInitialized() bool {
	return isInitialized
}

func InitializeEthofs(initFlag bool, configFlag bool, nodeType string, blockCommunication chan *types.Block) {

	checkResources(nodeType)

	localPinMapping = make(map[string]string)

	// initalize default locations
	defaultDataDir = node.DefaultDataDir()
	if runtime.GOOS == "linux" {
		ipcLocation = defaultDataDir + "/geth.ipc"
	} else if runtime.GOOS == "windows" {
		ipcLocation = defaultDataDir + "\\geth.ipc"
	} else if runtime.GOOS == "darwin" {
		ipcLocation = defaultDataDir + "/geth.ipc"
	}

	if initFlag {

 		log.Info("Starting ethoFS repo initialization")
		err := initializeEthofsNodeRepo(nodeType)
		if err == nil {
	 		log.Info("ethoFS repo initialization successful")
		} else {
	 		log.Warn("ethoFS repo initialization failed")
		}
		os.Exit(0)

	} else if configFlag {

 		log.Info("Starting ethoFS repo/node configuration")
		err := initializeEthofsNodeConfig(nodeType)
		if err == nil {
	 		log.Info("ethoFS configuration successful")
		} else {
	 		log.Warn("ethoFS configuration failed")
		}
		os.Exit(0)

	} else {
 		log.Info("Starting ethoFS node initialization", "type", nodeType)
		Ipfs, Node = initializeEthofsNode(nodeType)

		go func() {
			err := updatePinContractValues()
			if err != nil {
				log.Debug("ethoFS - error updating pin contract values")
			} else {
				log.Debug("ethoFS - pin contract value update successful")
			}

			updateLocalPinMapping(Ipfs)
		}()
		// Initialize block listener
		go BlockListener(blockCommunication)
	}
}

//func NewBlock(block *types.Block) {
func BlockListener(blockCommunication chan *types.Block) {

	returnFlag := false

	for {
		select {
		case block := <-blockCommunication:
			log.Info("ethoFS - new block received for processing", "number", block.Header().Number.Int64(), "txs", len(block.Transactions()))
			if len(block.Transactions()) > 0 {
				go CheckForUploads(block.Transactions())
			}
			go func() {

				err := updatePinContractValues()
				if err != nil {
					log.Debug("ethoFS - error updating pin contract values")
				} else {
					log.Debug("ethoFS - pin contract value update successful")
				}

				randomBlockSelector := rand.Intn(100)

				if randomBlockSelector > 95 {
					// Initiate garbage collection randomly roughly every 20 blocks
					go gc(Node)
				}
				if randomBlockSelector <= 5 && returnFlag == false {
					returnFlag = true
					// Update local pin tracking/mapping
					_, returnFlag = updateLocalPinMapping(Ipfs)
				}
			}()
		}
	}
}

func CheckForUploads(transactions types.Transactions) {
	for _, transaction := range transactions {
		recipient := transaction.To()
		if recipient == nil {
			continue
		} else if *recipient == contractControllerAddress {
			go func() {
				log.Info("ethoFS - new upload transaction detected", "hash", transaction.Hash())
				cids := scanForCids(transaction.Data())
				for _, pin := range cids {
					log.Debug("ethoFS - immediate pin request detail", "hash", pin)
					pinned := pinSearch(pin, localPinMapping)
					if !pinned {
						log.Debug("ethoFS - error while searching for pin", "error", "Pin not found")
						continue
					} else {
						log.Debug("ethoFS - data is pinned to local node", "hash", pin)
					}

					providerCount, err := FindProvs(Node, pin)
					if err != nil {
						log.Debug("ethoFS - provider search error", "error", err)
						continue
					}

					if !pinned && providerCount < (repFactor/uint64(2)) {
						// Pin data due to insufficient existing providers
						addedPin, err := pinAdd(Ipfs, pin)
						if err != nil {
							log.Debug("ethoFS - error adding pin", "hash", pin, "error", err)
							continue
						} else {
							log.Debug("ethoFS - pin added successfully", "hash", addedPin)
						}
					} else if pinned && providerCount > (repFactor+(repFactor/uint64(2))) {
						// Pin data due to insufficient existing providers
						removedPin, err := pinRemove(Ipfs, pin)
						if err != nil {
							log.Debug("ethoFS - pin removal error", "hash", pin, "error", err)
							continue
						} else {
							log.Debug("ethoFS - pin removal successful", "hash", removedPin)
						}
					}
				}
			}()
		}
	}
}

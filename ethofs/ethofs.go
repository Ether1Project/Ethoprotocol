package ethofs

import (
	//"fmt"
	"math/rand"
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

//var ContractPinTrackingMap map[string][]string
//var MasterPinArray []string
var localPinMapping map[string]string
var selfNodeID string
var repFactor = uint64(10)
var BlockHeight = int(0)
var Ipfs icore.CoreAPI
var Node *core.IpfsNode
var contractControllerAddress = common.HexToAddress("0xc38B47169950D8A28bC77a6Fa7467464f25ADAFc")
var mainChannelString = "ethoFSPinningChannel_alpha11"
var defaultDataDir string
//var DefaultDataDir = "/home/ether1node/.ether1"
var ipcLocation string

//var testHash = "QmdP3gTCyZwR4F8Kf5qFH6JovbXVXhLM7XiCRqnsTY5dHG"

func InitializeEthofs(nodeType string, blockCommunication chan *types.Block) {

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

	log.Info("Starting ethoFS node initialization", "type", nodeType)
	Ipfs, Node = initializeEthofsNode(nodeType)

	go func() {
		err := updatePinContractValues()
		if err != nil {
			log.Debug("ethoFS - error updating pin contract values")
		} else {
			log.Debug("ethoFS - pin contract value update successful")
		}
	}()
	// Initialize block listener
	go BlockListener(blockCommunication)
}

//func NewBlock(block *types.Block) {
func BlockListener(blockCommunication chan *types.Block) {
	for {
        	select {
        		case block := <-blockCommunication:
				log.Info("ethoFS - new block received for processing", "number", block.Header().Number.Int64(), "txs", len(block.Transactions()))
				if len(block.Transactions()) > 0 {
					go CheckForUploads(block.Transactions())
				}
				go func() {

					updateLocalPinMapping(Ipfs)

					err := updatePinContractValues()
					if err != nil {
						log.Debug("ethoFS - error updating pin contract values")
					} else {
						log.Debug("ethoFS - pin contract value update successful")
					}

					if rand.Intn(100) > 95 {
						// Initiate garbage collection randomly roughly every 20 blocks
						go gc(Node)
					}
				}()
	        }
    	}
}

func CheckForUploads(transactions types.Transactions) {
	for _, transaction := range transactions {
		recipient := transaction.To()
		if *recipient == contractControllerAddress {
			go func() {
				log.Info("ethoFS - new upload transaction detected", "hash", transaction.Hash())
				cids := scanForCids(transaction.Data())
				for _, pin := range cids {
					log.Debug("ethoFS - immediate pin request detail", "hash", pin)
					pinned := pinSearch(pin)
                        		if !pinned {
                                		log.Debug("ethoFS - pin search error", "error", "pin not found")
                                		continue
                        		} else {
                                		log.Debug("ethoFS - data is pinned to local node", "hash", pin)
                        		}

					providerCount, err := FindProvs(Node, pin)
					if err != nil {
						log.Warn("ethoFS - provider search error", "error", err)
						continue
					}

                        		if !pinned && providerCount < (repFactor / uint64(2))  {
                                		// Pin data due to insufficient existing providers
                                		addedPin, err := pinAdd(Ipfs, pin)
                                		if err != nil {
                                        		log.Debug("ethoFS - pin add error", "hash", pin, "error", err)
                                        		continue
                                		} else {
                                        		log.Debug("ethoFS - pin add successful", "hash", addedPin)
                                		}
                        		} else if pinned && providerCount > (repFactor + (repFactor / uint64(2)))  {
                                		// Pin data due to insufficient existing providers
                                		removedPin, err := pinRemove(Ipfs, pin)
                                		if err != nil {
                                        		log.Debug("ethoFS - pin remove error", "hash", pin, "error", err)
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

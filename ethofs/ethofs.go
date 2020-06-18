package ethofs

import (
	//"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	icore "github.com/ipfs/interface-go-ipfs-core"

	"github.com/ipfs/go-ipfs/core"
)

var selfNodeID string
var repFactor = 10
var BlockHeight = int(0)
var Ipfs icore.CoreAPI
var Node *core.IpfsNode
var contractControllerAddress = common.HexToAddress("0xc38B47169950D8A28bC77a6Fa7467464f25ADAFc")
var DefaultDataDir = "/home/ether1node/.ether1"

var testHash = "QmdP3gTCyZwR4F8Kf5qFH6JovbXVXhLM7XiCRqnsTY5dHG"

func InitializeEthofs(nodeType string) {
	log.Info("Starting ethoFS node initialization", "type", nodeType)
	Ipfs, Node = initializeEthofsNode()
}

func NewBlock(block *types.Block) {
	log.Info("ethoFS - new block received for processing", "number", block.Header().Number.Int64(), "txs", len(block.Transactions()))
	if len(block.Transactions()) > 0 {
		for _, transaction := range block.Transactions() {
			recipient := transaction.To()
			if *recipient == contractControllerAddress {
				log.Info("ethoFS - new upload transaction detected", "hash", transaction.Hash())
				txDataString := Between(string(transaction.Data()), "ethoFSPinningChannel_alpha11:", ",ethoFSPinningChannel_alpha11:")
				immediatePins := strings.Split(txDataString, ",")
				if len(immediatePins) > 0 {
				log.Info("ethoFS - immediate pin request detected", "pins", len(immediatePins))
					for _, pin := range immediatePins {
						pin = testHash
						log.Info("ethoFS - immediate pin request detail", "hash", pin)
						go func() {
							//if !(FindProvs(Node, pin)) {
								// Pin data due to insufficient existing providers
								addedPin, err := pinAdd(Ipfs, pin)
								if err != nil {
									log.Error("ethoFS - pin add error", "hash", pin, "error", err)
								} else {
									log.Info("ethoFS - pin add successful", "hash", addedPin)
								}
								_, err = pinSearch(Ipfs, pin)
								if err != nil {
									log.Error("ethoFS - pin search error", "error", err)
								} else {
									log.Info("ethoFS - pin search successful")
								}
							//}
						}()
					}
				}
			}
		}
	}
	//PinAnalysis()
}

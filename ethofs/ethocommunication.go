package ethofs

import (
	//	"./NodesStorage"
	//	"./PinStorage"
	"bytes"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ipfs/go-ipfs-api"
	"github.com/mitchellh/go-homedir"
	"log"
	"math/big"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var location string

// GetLatestEther1Block utilizes local geth.ipc to obtain latest block information
func GetLatestEther1Block() string {
	if runtime.GOOS == "linux" {
		if ethoLocation != "" {
			location = ethoLocation
		} else {
			location = "/.ether1/geth.ipc"
		}
	} else if runtime.GOOS == "windows" {
		location = "\\AppData\\Roaming\\Ether1\\geth.ipc"
	} else if runtime.GOOS == "darwin" {
		location = "/Library/Ether1/geth.ipc"
	}
	homedirectory, _ := homedir.Dir()
	c, _ := ethclient.Dial(homedirectory + location)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	if c != nil && ctx != nil {
		latestBlock, err := c.BlockByNumber(ctx, nil)
		if err != nil {
			log.Println("Ether-1 Network Connection Error: ", err)
			c.Close()
			return "0"
		}
		c.Close()
		return latestBlock.Number().String()
	}
	return "0"
}

// GetBootnodeContractValues utilizes local geth.ipc fo etho chain communication to obtain smart contract values
func GetBootnodeContractValues() {
	if runtime.GOOS == "linux" {
		if ethoLocation != "" {
			location = ethoLocation
		} else {
			location = "/.ether1/geth.ipc"
		}
	} else if runtime.GOOS == "windows" {
		location = "\\AppData\\Roaming\\Ether1\\geth.ipc"
	} else if runtime.GOOS == "darwin" {
		location = "/Library/Ether1/geth.ipc"
	}
	homedirectory, _ := homedir.Dir()
	c, _ := ethclient.Dial(homedirectory + location)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	if c != nil && ctx != nil {
		address := common.HexToAddress("0xd5Cc0D9031bD7CA18A320A777A9aCCFbEdFA8709")
		contract, err := NewNodesStorage(address, c)
		if err != nil {
			log.Println("Ether-1 Network Connection Error: ", err)
		} else {
			NodeCount, err := contract.NodeID(nil)
			if err != nil {
				log.Println("Ether-1 Network Connection Error: (NodeCount) ", err)
			} else {
				for j := uint32(1); j <= NodeCount; j++ {
					BootnodeArray, err := contract.Nodes(nil, j)
					if err != nil {
						log.Println("Ether-1 Network Connection Error: (Bootnodes) ", err)
					} else {
						BootnodePeers = append(BootnodePeers, BootnodeArray)
					}
				}
			}
			c.Close()
		}
	}
}

// GetPinContractValues utilizes local geth.ipc to obtain smart contract values
func GetPinContractValues(ContractPinTrackingMap map[string][]string) map[string][]string {
	if runtime.GOOS == "linux" {
		if ethoLocation != "" {
			location = ethoLocation
		} else {
			location = "/.ether1/geth.ipc"
		}
	} else if runtime.GOOS == "windows" {
		location = "\\AppData\\Roaming\\Ether1\\geth.ipc"
	} else if runtime.GOOS == "darwin" {
		location = "/Library/Ether1/geth.ipc"
	}
	InternalContractPinTrackingMap := make(map[string][]string)
	homedirectory, _ := homedir.Dir()
	c, _ := ethclient.Dial(homedirectory + location)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	if c != nil && ctx != nil {
		address := common.HexToAddress("0xD3b80c611999D46895109d75322494F7A49D742F")
		contract, err := NewPinStorage(address, c)
		if err != nil {
			log.Println("Ether-1 Network Connection Error: ", err)
		} else {
			ReplicationFactor, err := contract.ReplicationFactor(nil)
			if err != nil {
				log.Println("Ether-1 Network Connection Error: (ReplicationFactor) ", err)
			} else {
				repFactor = int(ReplicationFactor)
			}
			PinCount, err := contract.PinCount(nil)
			if err != nil {
				log.Println("Ether-1 Network Connection Error: (PinCount) ", err)
			} else {
				var ContractPinArray []string
				for j := uint64(0); j < uint64(PinCount); j++ {
					i := new(big.Int).SetUint64(j)
					ContractPin, err := contract.Pins(nil, i)
					if err != nil {
						log.Println("Ether-1 Network Connection Error: (ContractPin) ", err)
					} else {
						if _, ok := ContractPinTrackingMap[ContractPin]; ok {
							var tempPinArray []string
							for l := uint64(0); uint64(len(ContractPinTrackingMap[ContractPin])) > l; l++ {
								ContractPinArray = append(ContractPinArray, ContractPinTrackingMap[ContractPin][l])
								tempPinArray = append(tempPinArray, ContractPinTrackingMap[ContractPin][l])
							}
							InternalContractPinTrackingMap[ContractPin] = tempPinArray
						} else {
							ContractPinArray = append(ContractPinArray, ContractPin)
							//NOW CAT FOR SERIALISED PIN LIST IN IPFS
							catsh := shell.NewShell("localhost:" + apiPort)
							catsh.SetTimeout(1 * time.Second)
							resp, err := catsh.Cat(ContractPin)
							if err != nil {
							} else {
								buf := new(bytes.Buffer)
								buf.ReadFrom(resp)
								catString := buf.String()
								IPFSPinListArray := strings.Split(catString, ":")
								var tempPinArray []string
								if len(IPFSPinListArray) > 1 {
									if IPFSPinListArray[0] == mainChannelString {
										for k := uint64(1); k < uint64(len(IPFSPinListArray)); k++ {
											tempPinArray = append(tempPinArray, IPFSPinListArray[k])
											ContractPinArray = append(ContractPinArray, IPFSPinListArray[k])
										}
									}
								}
								InternalContractPinTrackingMap[ContractPin] = tempPinArray
							}
						}
					}
				}
				MasterPinArray = ContractPinArray
			}
		}
		c.Close()
	}
	return InternalContractPinTrackingMap
}

// InitialContractValues initiates local mapping with values stored in smart contract
func InitialContractValues() {
	fmt.Println("Waiting For Ether-1 Node To Finish Sync...")
	time.Sleep(30 * time.Second)
	for BlockHeight == 0 {
		latestBlock := GetLatestEther1Block()
		if latestBlock != "0" {
			tempBlockHeight, err := strconv.Atoi(latestBlock)
			if err != nil {
				log.Println("Ether-1 Network Sync Error ", err)
			} else {
				BlockHeight = tempBlockHeight
			}
		} else {
			fmt.Println("Waiting For Ether-1 Node To Finish Sync...")
		}
		time.Sleep(30 * time.Second)
	}
}

// UpdateContractBootnodeValues updates local mapping of smart contract bootnodes
func UpdateContractBootnodeValues() {
	GetBootnodeContractValues()
}

// UpdateContractPinValues updates local mapping of smart contract pins being stored
func UpdateContractPinValues() {
	ContractPinTrackingMap := make(map[string][]string)
	ContractPinTrackingMap = GetPinContractValues(ContractPinTrackingMap)
}

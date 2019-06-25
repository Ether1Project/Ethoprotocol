// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package nodeprotocol

import (
 	"fmt"
        "strconv"
        "crypto/ecdsa"
        "net/url"
        "net"
        "sync"
        "errors"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/p2p/enode"
        "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/enr"
	"github.com/ethereum/go-ethereum/params"
//	"github.com/ethereum/go-ethereum/core/nodeprotocolmessaging"
)

var activeNode *node.Node
var mux = sync.RWMutex{}
var protocolInitiationFlag bool
var nodeProtocolData map[string]map[uint64]NodeData

func ActiveNode() *node.Node {
	return activeNode
}

func SetActiveNode(stack *node.Node) {
        initializeProtocolHeadTracker()
	activeNode = stack
        SetupNodeProtocolMapping()
}

func initializeProtocolHeadTracker() {
        /*protocolHeadTracker := nodeprotocolmessaging.GetChan()

        done := make(chan uint64)

        go nodeprotocolmessaging.Background(protocolHeadTracker, done)*/
}

// CheckNodeProtocolStatus determines if node protocol data has been initiated
func CheckNodeProtocolStatus() bool {
        if len(nodeProtocolData) > 0 {
                return true
        }
        return false

}

// SetupNodeProtocolMap initiates node protocol mapping
func SetupNodeProtocolMapping() {
        if !protocolInitiationFlag {
                protocolInitiationFlag = true

                log.Info("Initializing Node Protocol Data Mapping")
                nodeProtocolData = make(map[string]map[uint64]NodeData)
                for _, nodeType := range params.NodeTypes {
                        InitiateNodeProtocolStateAccess(nodeType.Name)
                        nodeData, err := GetFullNodeDataState(nodeType.Name)
                        if err == nil {
                                log.Info("Node Protocol Data Mapping Initialized", "Type", nodeType.Name)
                                nodeProtocolData[nodeType.Name] = nodeData
                        } else {
                                log.Error("Unable To Initialize Node Protocol Data Mapping", "Type", nodeType.Name)
                        }
                }
        }
}

// CheckNodeStatus checks to see if specified node has been validated
func CheckNodeStatus(nodeType string, nodeId common.Hash, nodeIp common.Hash, blockHash common.Hash, blockNumber uint64) bool {
        if len(nodeProtocolData) == 0 {
                SetupNodeProtocolMapping()
        }

        if data, ok := nodeProtocolData[nodeType][blockNumber]; ok {
                if nodeId == common.BytesToHash([]byte(data.Id)) && nodeIp == common.BytesToHash([]byte(data.Ip)) {
                        log.Info("Node ID Found in Node Protocol Data", "Validated", "True", "Type", nodeType, "ID", nodeId, "IP", nodeIp)
                        return true
                }
                log.Warn("Node ID Not Found in Node Protocol Data", "Validated", "False", "Type", nodeType, "ID Needing Verification", nodeId, "Hash", blockHash, "IP", nodeIp, "Saved ID", common.BytesToHash([]byte(data.Id)), "Saved Hash", data.Hash, "Saved IP", common.BytesToHash([]byte(data.Ip)))
        }

        return false
}

// CheckUpToDate checks to see if blockHash has been recorded in mapping
func CheckUpToDate(nodeType string, blockHash common.Hash, blockNumber uint64) bool {
        if len(nodeProtocolData) == 0 {
                SetupNodeProtocolMapping()
        }

        if _, ok :=nodeProtocolData[nodeType][blockNumber]; ok {
                return true
        }
        return false
}

// GetNodeProtocolData returns the nodeid at specified blockHash of specific node type
func GetNodeProtocolData(nodeType string, blockHash common.Hash, blockNumber uint64) (string, string, error) {
        if len(nodeProtocolData) == 0 {
                SetupNodeProtocolMapping()
        }

        if data, ok := nodeProtocolData[nodeType][blockNumber]; ok {
                return data.Id, data.Ip, nil
        }
        return "", "", errors.New("Node Protocol Data Not Found")
}

// UpdateNodeProtocolData updates protocol mapping data for verified nodes
func UpdateNodeProtocolData(nodeType string, nodeId string, nodeIp string, peerId string, peerCount int, blockHash common.Hash, blockNumber uint64, syncing bool) {
        mux.Lock()
        defer mux.Unlock()

        log.Trace("Attempting To Update Node Protocol Data", "Type", nodeType, "ID", nodeId, "IP", nodeIp, "Number", blockNumber, "Hash", blockHash)

        nodeData := NodeData{Id: nodeId, Ip: nodeIp, Hash: blockHash, Number: blockNumber}
        nodeProtocolData[nodeType][blockNumber] = nodeData

        log.Trace("Node Protocol Data Updated - Initial Data Added", "Type", nodeType, "ID", nodeId, "IP", nodeIp, "Hash", blockHash)
}

func SaveNodeProtocolDataMapping(){
        if len(nodeProtocolData) == 0 {
                SetupNodeProtocolMapping()
        }

        for _, nodeType := range params.NodeTypes {
                SaveNodeDataStateGroup(nodeType.Name, nodeProtocolData[nodeType.Name])
        }
}

// SyncNodeProtocolDataGroup adds a slice of NodeData to state is consenus is reached
func SyncNodeProtocolDataGroup(nodeType string, nodeData map[uint64]NodeData, peerId string, peerCount int) {
        if len(nodeProtocolData) == 0 {
                SetupNodeProtocolMapping()
        }

        largestBlockNumber := uint64(0)
        smallestBlockNumber := uint64(99999999)
        for blockNumber, _ := range nodeData {
                if blockNumber > largestBlockNumber {
                        largestBlockNumber = blockNumber
                } else if blockNumber < smallestBlockNumber {
                        smallestBlockNumber = blockNumber
                }
        }

        mux.Lock()
        defer mux.Unlock()
        for number, data := range nodeData {
                nodeProtocolData[nodeType][number] = data
        }

        if len(nodeData) > 0 {
                log.Info("Importing Node Protocol Data", "Entries", len(nodeData), "Blocks", strconv.FormatUint(smallestBlockNumber, 10) + "->" + strconv.FormatUint(largestBlockNumber, 10))
        }
        /*if largestBlockNumber >= nodeprotocolmessaging.GetNodeProtocolHead() {
                log.Info("Updating Head Block", "Largest", largestBlockNumber)
                nodeprotocolmessaging.Set(largestBlockNumber)
        }*/
}

//func GetNodeProtocolDataGroup(nodeType string, startBlock uint64, endBlock uint64) (map[uint64]NodeData, error) {
func GetNodeProtocolDataGroup(nodeType string, startBlock uint64, endBlock uint64) ([]string, []string, []string, []string, error) {
        var hashes []string
        var nodes  []string
        var ips  []string
        var numbers []string
        //dataMap, err := GetFullNodeDataState(nodeType)
        //if err == nil {
        dataMap := nodeProtocolData[nodeType]
        for i := startBlock; i <= endBlock; i++ {
                if data, ok := dataMap[i]; ok {
                        hashes  = append(hashes, data.Hash.String())
                        nodes   = append(nodes, data.Id)
                        ips     = append(ips, data.Ip)
                        numbers = append(numbers, strconv.FormatUint(i, 10))
                }
        }
        return hashes, nodes, numbers, ips, nil
}
// GetNodeId return enodeid in string format from *enode.Node
func GetNodeId(n *enode.Node) string {
        var (
                scheme enr.ID
                nodeid string
                key    ecdsa.PublicKey
        )
        n.Load(&scheme)
        n.Load((*enode.Secp256k1)(&key))
        nid := n.ID()
        switch {
        case scheme == "v4" || key != ecdsa.PublicKey{}:
                nodeid = fmt.Sprintf("%x", crypto.FromECDSAPub(&key)[1:])
        default:
                nodeid = fmt.Sprintf("%s.%x", scheme, nid[:])
        }
        u := url.URL{Scheme: "enode"}
        if n.Incomplete() {
                u.Host = nodeid
        } else {
                addr := net.TCPAddr{IP: n.IP(), Port: n.TCP()}
                u.User = url.User(nodeid)
                u.Host = addr.String()
                if n.UDP() != n.TCP() {
                        u.RawQuery = "discport=" + strconv.Itoa(n.UDP())
                }
        }
        return u.User.String()
}

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

package nodeprotocolmessaging

import (
        "sync"

        "github.com/ethereum/go-ethereum/common"
        "github.com/ethereum/go-ethereum/core/state"
        "github.com/ethereum/go-ethereum/core/types"
)

var pm Manager
var peerSet PeerSet
var bc Blockchain
var SyncWg sync.WaitGroup

type Manager interface {
        AsyncGetNodeProtocolData(data []string)
        AsyncSendNodeProtocolData(data []string)
        AsyncGetNodeProtocolSyncData(data []string)
        AsyncGetNodeProtocolPeerVerification(data []string)
}

type PeerSet interface {
        Len() int
        String() []string
        Ips() map[string]string
}

type Blockchain interface {
        StateAt(hash common.Hash) (*state.StateDB, error)
        Rollback(chain []common.Hash) 
        GetBlockByNumber(number uint64) *types.Block
        CurrentBlock() *types.Block
        GetBlockByHash(hash common.Hash) *types.Block
}

func SetBlockchain(blockchain Blockchain) {
        bc = blockchain
}

func GetStateAt(hash common.Hash) (*state.StateDB, error) {
        return bc.StateAt(hash)
}

func GetBlockByHash(hash common.Hash) *types.Block {
        return bc.GetBlockByHash(hash)
}

func SetProtocolManager(manager Manager) {
        pm = manager
}

func SetPeerSet(ps PeerSet) {
       peerSet = ps
}

func RollBackChain(count uint64) {
        var chain []common.Hash
        currentBlockNumber := bc.CurrentBlock().Header().Number.Uint64()
        for i := uint64(0); i < count; i++ {
                hash := bc.GetBlockByNumber(currentBlockNumber - i).Hash()
                chain = append(chain, hash)
        }
        bc.Rollback(chain)
}

func CheckPeerSet(id string, ip string) bool {
        ipMap := peerSet.Ips()
        for _, peerId := range peerSet.String() {
                if peerIp, ok := ipMap[peerId]; ok {
                        // Return true if peer is found
                        if id == peerId && ip == peerIp {
                                return true
                        }
                }
        }
        return false
}

func GetPeerCount() int {
        return peerSet.Len()
}

func RequestNodeProtocolData(data []string) {
        pm.AsyncGetNodeProtocolData(data)
}

func SendNodeProtocolData(data []string) {
        pm.AsyncSendNodeProtocolData(data)
}

func RequestNodeProtocolSyncData(data []string) {
        pm.AsyncGetNodeProtocolSyncData(data)
}

func RequestNodeProtocolPeerVerification(data []string) {
        pm.AsyncGetNodeProtocolPeerVerification(data)
}

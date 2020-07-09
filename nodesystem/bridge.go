// Copyright 2015 The go-ethereum Authors
// Copyright 2019 The Etho.Black Development Team
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

package nodesystem

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

var pm Manager
var peerSet PeerSet
var bc Blockchain
var SyncWg sync.WaitGroup
var privateAdminApi PrivateAdminAPI
var publicEthereumApi PublicEthereumAPI

type Manager interface {
	SyncStatus() bool
	AsyncGetNodeProtocolValidations(state *state.StateDB, id string, hash common.Hash, number uint64)
}

type PeerSet interface {
	Len() int
	String() []string
	Ips() map[string]string
}

type PrivateAdminAPI interface {
	AddPeer(url string) (bool, error)
}

type PublicEthereumAPI interface {
	Syncing() (interface {}, error)
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

func SetPrivateAdminApi(api PrivateAdminAPI) {
	privateAdminApi = api
}

func SetPublicEthereumApi(api PublicEthereumAPI) {
	publicEthereumApi = api
}

func Syncing() bool {
	data, err := publicEthereumApi.Syncing()
	if err == nil {
		if data == false {
			log.Debug("blockchain sync status", "syncing", "false")
			return false
		} else {
			log.Debug("blockchain sync status", "syncing", "true")
			return true
		}
	}
	return true
}

func GetStateAt(hash common.Hash) (*state.StateDB, error) {
	return bc.StateAt(hash)
}

func GetBlockByHash(hash common.Hash) *types.Block {
	return bc.GetBlockByHash(hash)
}

func GetBlockByNumber(number uint64) *types.Block {
	return bc.GetBlockByNumber(number)
}

func SetProtocolManager(manager Manager) {
	pm = manager
}

func SetPeerSet(ps PeerSet) {
	peerSet = ps
}

func RequestNodeProtocolValidations(state *state.StateDB, id string, hash common.Hash, number uint64) {
	pm.AsyncGetNodeProtocolValidations(state, id, hash, number)
}

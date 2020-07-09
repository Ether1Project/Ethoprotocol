// Copyright 2015 The go-ethereum Authors
// Copyright 2020 The Etho.Black Development Team
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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/node"
)

var activeNode *node.Node
var protocolInitiationFlag bool
var chainDB ethdb.Database
var SyncStartingBlock uint64
var HoldBlockNumber string
var HoldBlockCount int64

type NodeData struct {
	Hash   common.Hash `json:"hash"`
	Id     string      `json:"id"`
	Number uint64      `json:"number"`
	Ip     string      `json:"ip"`
}

func ActiveNode() *node.Node {
	return activeNode
}

func SetChainDB(db ethdb.Database) {
	chainDB = db
}

func SetActiveNode(stack *node.Node) {
	activeNode = stack
}

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
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/log"
)

// Check validity of previously recorded node address
func ValidateNodeAddress(state *state.StateDB, chain consensus.ChainReader, parent *types.Block, verifiedNode common.Address, contractAddress common.Address) bool {
        previousBlock := chain.GetBlock(parent.Header().ParentHash, parent.Header().Number.Uint64()-1)
        nodeIndex := new(big.Int).Mod(previousBlock.Hash().Big(), big.NewInt(GetNodeCount(state, contractAddress))).Int64()

        nodeAddressString := GetNodeKey(state, nodeIndex, contractAddress)

        if common.HexToAddress(nodeAddressString) == verifiedNode {
        	log.Info("Node Address Validation Successful", "Node Address", verifiedNode)
                return true
        }
       	log.Warn("Node Address Validation Failed", "Node Address", verifiedNode)
        return true
}

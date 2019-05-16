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
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/params"
)

func GetNodeCount(state *state.StateDB) (int64) {
        // Get storage state form db using index
        nodeCount := state.GetState(params.NodeProtocolContract, common.HexToHash("2"))

        return nodeCount.Big().Int64()
}

func GetNodeKey(state *state.StateDB, nodeIndex int64) (string) {
	solcIndex := int64(1)

        hash := sha3.NewKeccak256()
	var buf []byte

        // Index is the contract variable index based on solc storage state standards
        index := abi.U256(big.NewInt(solcIndex))

        // Key is the mapping key to lookup
        key := abi.U256(big.NewInt(int64(nodeIndex)))

        // Prepare the Keccak256 seed
        location := append(key, index...)

	hash.Write(location)
	buf = hash.Sum(nil)
        storageLocation := common.BytesToHash(buf)

        // Get storage state form db using the hashed data
        response := state.GetState(params.NodeProtocolContract, storageLocation)

        // Assemble the strings
        nodeAddressString := response

        return nodeAddressString.String()
}

func GetNodeData(state *state.StateDB, nodeAddress string) (string, string) {
	solcIndex := int64(0)

        hash := sha3.NewKeccak256()
	var buf []byte

        // Left-fill with zeros to meet abi packing standards
        prepend := make([]byte, 12)

        // Index is the contract variable index based on solc storage state standards
        index := abi.U256(big.NewInt(solcIndex))

        // Key is the mapping key to lookup
        key := common.HexToAddress(nodeAddress).Bytes()[:]

        // Prepare the Keccak256 seed
        location := append(prepend, key...)
        location = append(location, index...)

	hash.Write(location)
	buf = hash.Sum(nil)
        storageLocation := common.BytesToHash(buf)

	nodeIdLocation := common.BigToHash(new(big.Int).Add(storageLocation.Big(), big.NewInt(2))).Bytes()

        // Get offsets for long enodeid string
        hash = sha3.NewKeccak256()
	var buf1 []byte
	hash.Write(nodeIdLocation)
	buf1 = hash.Sum(nil)
        finalNodeIdLocation1 := common.BytesToHash(buf1)
	finalNodeIdLocation2 := common.BigToHash(new(big.Int).Add(finalNodeIdLocation1.Big(), big.NewInt(1)))
	finalNodeIdLocation3 := common.BigToHash(new(big.Int).Add(finalNodeIdLocation1.Big(), big.NewInt(2)))
	finalNodeIdLocation4 := common.BigToHash(new(big.Int).Add(finalNodeIdLocation1.Big(), big.NewInt(3)))

	nodeAddressLocation := common.BigToHash(new(big.Int).Add(storageLocation.Big(), big.NewInt(1)))
	nodeIpLocation := common.BigToHash(new(big.Int).Add(storageLocation.Big(), big.NewInt(3)))
	nodePortLocation := common.BigToHash(new(big.Int).Add(storageLocation.Big(), big.NewInt(4)))

        // Get storage state form db using the hashed data
        responseNodeId1 := state.GetState(params.NodeProtocolContract, finalNodeIdLocation1)
        responseNodeId2 := state.GetState(params.NodeProtocolContract, finalNodeIdLocation2)
        responseNodeId3 := state.GetState(params.NodeProtocolContract, finalNodeIdLocation3)
        responseNodeId4 := state.GetState(params.NodeProtocolContract, finalNodeIdLocation4)
        responseNodeIp := state.GetState(params.NodeProtocolContract, nodeIpLocation)
	responseNodePort := state.GetState(params.NodeProtocolContract, nodePortLocation)
	responseNodeAddress := state.GetState(params.NodeProtocolContract, nodeAddressLocation)

        // Assemble the strings
	nodeIdString := "enode://" + stripCtlAndExtFromBytes(string(responseNodeId1.Bytes())) + stripCtlAndExtFromBytes(string(responseNodeId2.Bytes())) + stripCtlAndExtFromBytes(string(responseNodeId3.Bytes())) + stripCtlAndExtFromBytes(string(responseNodeId4.Bytes())) + "@" + stripCtlAndExtFromBytes(string(responseNodeIp.Bytes())) + ":" + stripCtlAndExtFromBytes(string(responseNodePort.Bytes()))
        nodeAddressString := common.BytesToAddress(responseNodeAddress.Bytes()).String()

        return nodeIdString, nodeAddressString
}

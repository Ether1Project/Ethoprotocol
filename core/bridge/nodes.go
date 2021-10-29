// Copyright 2014 The go-ethereum Authors
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

// Package types contains data types related to Ethereum consensus.
package types

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/big"
	"reflect"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
)

var (
	ConfigurationAddress  = common.HexToAddress("0xB1C7E6384288dc3F677Def5B2aa2A5807e31a42D")
)

// DecodeBridgeConfig is entry point for cross-chain collateralized node activation
func DecodeBridgeConfig(state *state.StateDB, index int64) *big.Int {
	addressIndex := int64(3)
	valueIndex := int64(2)

	hash := sha3.NewKeccak256()
	var buf []byte

	// Left-fill with zeros to meet abi packing standards
	prepend := make([]byte, 12)

	// Index is the contract variable index based on solc storage state standards
	varIndex := abi.U256(big.NewInt(addressIndex))

	// Key is the mapping key to lookup
	key := index

	// Prepare the Keccak256 seed
	location := append(prepend, key...)
	location = append(location, varIndex...)

	hash.Write(location)
	buf = hash.Sum(nil)
	storageLocation := common.BytesToHash(buf)

	lookupLocation := common.BigToHash(new(big.Int).Add(storageLocation.Big(), big.NewInt(2))).Bytes()

	// Get offsets for a long enodeid string
	hash = sha3.NewKeccak256()
	var buf1 []byte
	hash.Write(lookupLocation)
	buf1 = hash.Sum(nil)
	finalLocation := common.BytesToHash(buf1)

	nodeAddressLocation := common.BigToHash(new(big.Int).Add(storageLocation.Big(), big.NewInt(1)))
	nodeIpLocation := common.BigToHash(new(big.Int).Add(storageLocation.Big(), big.NewInt(3)))

	// Get storage state from the db using the hashed data
	responseAddress := state.GetState(ConfigurationAddress, storageLocation)

        // Work in progress
	return nil
}


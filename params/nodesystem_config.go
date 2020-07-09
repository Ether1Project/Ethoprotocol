// Copyright 2016 The go-ethereum Authors
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

package params

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

//var NodeIdArray []string
//var NodeIpArray []string

var NodeSystemContractAddress = common.HexToAddress("0xdF1Fa89137948338BF3bC96c340b0Ee97C144b23")
var TotalRewardShares = big.NewInt(13)

// Main node type configuration for node-protocol
var NodeTypes = []NodeType{
	NodeType{
		Name:                        "Service Node",
		RequiredCollateral:          new(big.Int).Mul(big.NewInt(5000), big.NewInt(1e+18)),
		BlockReward:                 big.NewInt(0),  //Initialize to 0 since we are using monetary policy config during consensus
		RewardSplit:                 big.NewInt(1),
		RemainderAddress:            common.HexToAddress("0x1000000000000000000000000000000000000001"),
		TxAddress:     	             common.HexToAddress("0x1000000000000000000000000000000000001000"),
		CountCachingAddresses:       []common.Address{common.HexToAddress("0x1000000000000000000000000000000000000014"), common.HexToAddress("0x1000000000000000000000000000000000000013"), common.HexToAddress("0x1000000000000000000000000000000000000012"), common.HexToAddress("0x1000000000000000000000000000000000000011"), common.HexToAddress("0x1000000000000000000000000000000000000010")},
		NodeIdCachingAddresses:      []common.Address{common.HexToAddress("0x1000000000000000000000000000000000000114"), common.HexToAddress("0x1000000000000000000000000000000000000113"), common.HexToAddress("0x1000000000000000000000000000000000000112"), common.HexToAddress("0x1000000000000000000000000000000000000111"), common.HexToAddress("0x1000000000000000000000000000000000000110")},
		NodeAddressCachingAddresses: []common.Address{common.HexToAddress("0x1000000000000000000000000000000000000214"), common.HexToAddress("0x1000000000000000000000000000000000000213"), common.HexToAddress("0x1000000000000000000000000000000000000212"), common.HexToAddress("0x1000000000000000000000000000000000000211"), common.HexToAddress("0x1000000000000000000000000000000000000210")},
		NodeIpCachingAddresses:      []common.Address{common.HexToAddress("0x1000000000000000000000000000000000000314"), common.HexToAddress("0x1000000000000000000000000000000000000313"), common.HexToAddress("0x1000000000000000000000000000000000000312"), common.HexToAddress("0x1000000000000000000000000000000000000311"), common.HexToAddress("0x1000000000000000000000000000000000000310")},
	},
	NodeType{
		Name:                        "Masternode",
		RequiredCollateral:          new(big.Int).Mul(big.NewInt(15000), big.NewInt(1e+18)),
		BlockReward:                 big.NewInt(0),  //Initialize to 0 since we are using monetary policy config during consensus
		RewardSplit:                 big.NewInt(4),
		RemainderAddress:            common.HexToAddress("0x1000000000000000000000000000000000000002"),
		TxAddress:     	             common.HexToAddress("0x1000000000000000000000000000000000002000"),
		CountCachingAddresses:       []common.Address{common.HexToAddress("0x1000000000000000000000000000000000000024"), common.HexToAddress("0x1000000000000000000000000000000000000023"), common.HexToAddress("0x1000000000000000000000000000000000000022"), common.HexToAddress("0x1000000000000000000000000000000000000021"), common.HexToAddress("0x1000000000000000000000000000000000000020")},
		NodeIdCachingAddresses:      []common.Address{common.HexToAddress("0x1000000000000000000000000000000000000124"), common.HexToAddress("0x1000000000000000000000000000000000000123"), common.HexToAddress("0x1000000000000000000000000000000000000122"), common.HexToAddress("0x1000000000000000000000000000000000000121"), common.HexToAddress("0x1000000000000000000000000000000000000120")},
		NodeAddressCachingAddresses: []common.Address{common.HexToAddress("0x1000000000000000000000000000000000000224"), common.HexToAddress("0x1000000000000000000000000000000000000223"), common.HexToAddress("0x1000000000000000000000000000000000000222"), common.HexToAddress("0x1000000000000000000000000000000000000221"), common.HexToAddress("0x1000000000000000000000000000000000000220")},
		NodeIpCachingAddresses:      []common.Address{common.HexToAddress("0x1000000000000000000000000000000000000324"), common.HexToAddress("0x1000000000000000000000000000000000000323"), common.HexToAddress("0x1000000000000000000000000000000000000322"), common.HexToAddress("0x1000000000000000000000000000000000000321"), common.HexToAddress("0x1000000000000000000000000000000000000320")},
	},
	NodeType{
		Name:                        "Gateway Node",
		RequiredCollateral:          new(big.Int).Mul(big.NewInt(30000), big.NewInt(1e+18)),
		BlockReward:                 big.NewInt(0),  //Initialize to 0 since we are using monetary policy config during consensus
		RewardSplit:                 big.NewInt(8),
		RemainderAddress:            common.HexToAddress("0x1000000000000000000000000000000000000003"),
		TxAddress:     	             common.HexToAddress("0x1000000000000000000000000000000000003000"),
		CountCachingAddresses:       []common.Address{common.HexToAddress("0x1000000000000000000000000000000000000034"), common.HexToAddress("0x1000000000000000000000000000000000000033"), common.HexToAddress("0x1000000000000000000000000000000000000032"), common.HexToAddress("0x1000000000000000000000000000000000000031"), common.HexToAddress("0x1000000000000000000000000000000000000030")},
		NodeIdCachingAddresses:      []common.Address{common.HexToAddress("0x1000000000000000000000000000000000000134"), common.HexToAddress("0x1000000000000000000000000000000000000133"), common.HexToAddress("0x1000000000000000000000000000000000000132"), common.HexToAddress("0x1000000000000000000000000000000000000131"), common.HexToAddress("0x1000000000000000000000000000000000000130")},
		NodeAddressCachingAddresses: []common.Address{common.HexToAddress("0x1000000000000000000000000000000000000234"), common.HexToAddress("0x1000000000000000000000000000000000000233"), common.HexToAddress("0x1000000000000000000000000000000000000232"), common.HexToAddress("0x1000000000000000000000000000000000000231"), common.HexToAddress("0x1000000000000000000000000000000000000230")},
		NodeIpCachingAddresses:      []common.Address{common.HexToAddress("0x1000000000000000000000000000000000000334"), common.HexToAddress("0x1000000000000000000000000000000000000333"), common.HexToAddress("0x1000000000000000000000000000000000000332"), common.HexToAddress("0x1000000000000000000000000000000000000331"), common.HexToAddress("0x1000000000000000000000000000000000000330")},
	},
}

// NodeType is the main node system configuration for various
// node types to be deployed on the network
type NodeType struct {
	Name                        string
	RequiredCollateral          *big.Int
	BlockReward                 *big.Int
	RewardSplit                 *big.Int
	RemainderAddress            common.Address
	TxAddress                   common.Address
	CountCachingAddresses       []common.Address
	NodeIdCachingAddresses      []common.Address
	NodeAddressCachingAddresses []common.Address
	NodeIpCachingAddresses      []common.Address
}

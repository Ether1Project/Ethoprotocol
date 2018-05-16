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

package params

// Ether-1 Bootnodes
var MainnetBootnodes = []string{
	// Ether-1 Bootnodes
	"enode://69d27ffbb1366b16f51e8d6df6b7de39e3f6346a143bfbfb8589842a3c30f6a8c6cf707235a0704d5bfe0986a0744311e97521aacc40b850e8ee1352f3f9893b@144.202.6.22:30301",      //USA
        "enode://b24d11b71d7bb2510bec40b923f6ae49da052d10841f9d18028dc40e439009bd9f865bef7e67143526a9b04c428d4fc0e3f2454ec7e27341686d9b611b98d334@108.160.142.71:30301",    //JAPAN
	"enode://2ab3c27920b9effe185faea7260852a603e906427547b5382e276e682bcb04381d8575b63b41a244a6841fa19aab74469722183eb49811f9a64d3d46bf3ea87c@45.76.148.77:30301",      //SINGAPORE
        "enode://4086a974cb98b9d47dc9823c42a2c4df5446684a4566cf52159bf49d849708cd66f733c6d5fca4b89a72173fe53c048e8aec3961d3772bc6038535bda538778f@45.76.117.164:30301",     //AUSTALIA
	}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	}

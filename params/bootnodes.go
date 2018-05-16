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
	"enode://58b166ee3a71bd44bb4b95ba1f906e139eb72e076626e440abbcda778ba18709d8bba890790f7876ae08acbd537a6ceca9682a815bf30d872892f7db1e247ab5@144.202.6.22:30301",      //USA
        "enode://3ddab6414cb534443f3166cb8c589287487e1ee0d8c97750bca34dc3c784dc8cf0300c76bbc023aa87bd76caeaef2e917a927c8c8ef8a51d4445b877e2bac991@108.160.142.71:30301",    //JAPAN
	"enode://6579977c0452a4c2fc56789e495a9c74354b7ce9b537eaa855b4555d12de8208b4a47b4130e6ee867dfbd294c6105c5e40c07b1e1d8591653ebe0aa7d229d5f2@45.76.148.77:30301",      //SINGAPORE
        "enode://a9c054ea0a2abb4d5d09b16415c1dd0efb2d3738ecc29e27ff9a081114426802de34cccaf2b47dd7100e740ce8f82223325626423ad04a53a461345eadd95a0a@45.76.117.164:30301",     //AUSTALIA
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

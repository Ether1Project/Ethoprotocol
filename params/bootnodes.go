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
	"enode://959caa70cbb602b0abe14b97f47f3c43813d1dbe8b54dd48882f22338b248752e8509c0e23b480481c550cca44c05e8df999680a8edb07d1f73ab0bf96575050@45.77.184.141:30301",     //WEST COAST - USA
	"enode://74f65548bbee33122a9fce4165e43e9aa21919e99de6bf204e7829eb1491115659a56f8d39e7460937a6c16311616a499c1a5ae8c7b2f8aa71e5a7701c07b54b@144.202.6.22:30301",      //EAST COAST - USA
        "enode://8a78da78bd8e347e8941c490b849168a12b67ed542141d05ec826902531c337d241c73b3f74e3ce0c984ae5fa1a0a6229f709dc63463dd81b2eb4910f7bda348@108.160.142.71:30301",    //JAPAN
	"enode://e646471531c627e0b74a8ff15967c7d3f982c99883098d4b6921e3cc184232549d159dd682a08d59b58fdf864aa7966163aa0f6d661a9e7ae2d3bcb8157f51a1@45.76.148.77:30301",      //SINGAPORE
        "enode://d6859aed0e935fe1c79c2b914b1b3981ff44d2a350997ff5a92fe57b996d01d8b7afe74237d39274937045a41f97b4fb71d0051f046e5541819c51dec1fe8907@45.76.117.164:30301",     //AUSTALIA
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

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
	//"enode://cc3ef2e41bb6a88e38786d8c807fe3560c71f4af4164e12d927cf9ba67df5e78d45bd64292cb88e2007fd165b9202e4f46fffeae439065a3fbf6388e4b3331d4@104.238.183.9:30303",     //MIAMI-USA
        //"enode://4949ce1aa4d7232569c836c06eca5324a16c46c1214b7032493c8ca75b38bc3f3ff2ce7b396bdbbd03ebe835cb02a6390cc6c9055abb3099db9d758bdf70448b@199.247.21.164:30303",    //GERMANY
        //"enode://759d2f7fe430876aa69f7e5a8585f0cf49b9d809b133c404b14e01e7016c685eae660480d5200063d92dc218494aae7ac13c137f4bda6453a8f41f70d6e273e9@45.32.190.70:30303",      //AUSTALIA
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

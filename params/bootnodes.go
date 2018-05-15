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
	"enode://978111e37e0227f768f95081ee78861f0d44bcf8c0a01f75105e9c82c7d1420b32ed2f98571f2497797e40e818c9ce09cdb62c03b75fe13a87f32c3d873b02bf@45.77.184.141:30301",     //WEST COAST - USA
	"enode://7c39901095e21c7a73208a33fedf3750e9dc7411bd610ac63a23e59c5f65be3770f3528eff97407acf4255a512ae6f6080b7052d22981c18ac17270669d718c4@144.202.6.22:30301",      //EAST COAST - USA
        "enode://d10b16f1b64a100247a8008140a984f629f5d08a45c39b9043e56f96804b579ea59713dad2ba0a3bed19ba885bf760fbdc4e6ce7bed62d018b2e57d7b04a57dd@108.160.142.71:30301",    //JAPAN
	"enode://98394c2274cfcd67ce71e0a8f8e2a02fb0a4641b77d10c603a2a40be8ee32275ab69086c3f9ed374afc96740e4b3f9300decc5e3b59260f4fdea7439cada726a@45.76.148.77:30301",      //SINGAPORE
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

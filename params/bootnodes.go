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
        // Xerom Bootnodes
        "enode://22b31e7fd3cf44072387d7e3765bfac545b9e623746e3ca19233b010997d6074a36e0851d8fe6fda98964315c81d445e5a2ba077960d9950446c514d56e61224@149.28.49.244:30305",
        "enode://3480c503c27de9220b9440eb1f4f92d7e16e396f223ace88661bf319d8871f179b0214d8121bbb229985090c1d2eaf7b046329dd27f605851b2860b7661922ca@45.32.170.196:30305",
        "enode://2c2fc4d4eb62081096aad210daaaba59f6aa84dc1408351f05d35143c6b08a8bae8e5283fc59756e0dca99b33e541d4307ea41f295ad5b486942e4b7ee74cbf9@45.32.152.162:30305",
        "enode://9e3db882078176028076cd4510c299e112069f969c7fa5c662517eb6cf119f54e74ab0348b110d7f3d82b40fe3a9c855f5e5c11673088378d3e4f56813dbca4b@140.82.50.249:30305",
        "enode://2b58a5e8fd16050d3de67edc081ccc3a95f41c502c34679064514cbe87ef71afea21a694f8e0a2cb4fca3a6b76c6702881e601a909e9d085f9050b17699dafa5@45.32.221.141:30305",
        "enode://d3b89d32e89febbd0c28f58c1b5c72fd770343739cec9b9c525b0cc7f2383d68720fe752f06259d157d3b4489ca6214085b0e337512b4c0a53f93e5020d89a21@139.180.160.239:30305",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
	// Upstrem bootnodes
	"enode://011f758e6552d105183b1761c5e2dea0111bc20fd5f6422bc7f91e0fabbec9a6595caf6239b37feb773dddd3f87240d99d859431891e4a642cf2a0a9e6cbb98a@51.141.78.53:30303",
	"enode://176b9417f511d05b6b2cf3e34b756cf0a7096b3094572a8f6ef4cdcb9d1f9d00683bf0f83347eebdf3b81c3521c2332086d9592802230bf528eaf606a1d9677b@13.93.54.137:30303",
	"enode://46add44b9f13965f7b9875ac6b85f016f341012d84f975377573800a863526f4da19ae2c620ec73d11591fa9510e992ecc03ad0751f53cc02f7c7ed6d55c7291@94.237.54.114:30313",
	"enode://c1f8b7c2ac4453271fa07d8e9ecf9a2e8285aa0bd0c07df0131f47153306b0736fd3db8924e7a9bf0bed6b1d8d4f87362a71b033dc7c64547728d953e43e59b2@52.64.155.147:30303",
	"enode://f4a9c6ee28586009fb5a96c8af13a58ed6d8315a9eee4772212c1d4d9cebe5a8b8a78ea4434f318726317d04a3f531a1ef0420cf9752605a562cfe858c46e263@213.186.16.82:30303",

	// Ethereum Foundation bootnode
	"enode://573b6607cd59f241e30e4c4943fd50e99e2b6f42f9bd5ca111659d309c06741247f4f1e93843ad3e8c8c18b6e2d94c161b7ef67479b3938780a97134b618b5ce@52.56.136.200:30303",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	}

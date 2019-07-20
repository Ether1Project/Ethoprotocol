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
	"enode://b59510af66c4ace3ba489e135d39d1c96ab549261922684efa41121612433feb64e63173c53a45dd5d8da2826152b08a6587992af66c53fe7cdd269c54286d70@157.230.144.98:30305",
	"enode://335a900592ae34ec50cc13a5434f8737dd5e29172f7bcfab071b35414cd48de4d1a51a49d177f85f4ff85f7edcb8ec137bee739afb6fc6053bad10478b8f3da9@206.189.30.232:30305",
	"enode://a557fe5cbad84b33a3fe35d46b9153b0514e7b6c533cf9ee9adf53f144799243a2100ad44bc8c12dd185549bdebb446a28b62dc390ba1ed20681cdbb97324cfc@95.179.163.200:30305",
	"enode://908d14e9791f57eb5db026c30777f8bb8be90f11b2a17643d8ba2c6d449fd13a84a11d0bce1d458b1834624a5f6b3af486bbf297cdffbb3697817e7c64cd5867@142.93.151.20:30305",
	"enode://07bd358da373fd4d982517ad6879503bdc1885092830ee799609f442ebc0b08a99896eaf073ec7b14741545839f8a9e0f776812f1dd4b243cc152108090ef0f3@45.32.197.67:30303",
  "enode://b66778fe345930e14e5d7684bc2ba59cc1f4fa9ec6dad33d2f008d89d5091582f11d04d896986a81b3bb2ce907a1f562813c6c0740e14082f1360dca9d8ac2fb@140.82.57.96:30303",
	"enode://9cb32f5393dac554a60d5e4a062bb9dec5fcfbdfe730fe6047dc987720e456f409a1c92f525631f15bafe6760a4ce683073ae2d9d5b24cc46aa5b164627caf4d@199.247.21.43:30303",
	"enode://d8e3ffee500298645ee198a4e2cf4eb04a4e82b81b3a7627d17f134c03cf5f5d3c20bf73b1b4967a0b13f1b22f973837991063c75a6bb23f2ddce9c8489fbdf6@45.77.184.141:30303",
	"enode://1ed58103ebf40db5079e9100114de7520185bbfedaf88250e44e9e75066e26f8fc3ea5233a58249e276c8509517dcd856f52889ccadb1a5a172458edefef02c8@104.238.188.219:30303",
	"enode://8dd418438c7a7da7e9fa33e9129e1361d029e935623fe000af71b9c4169c2e8c9e6f315ac34029bfbf4e12076210cb2cd4cbe8362a76d4d127f5c59c58905085@45.63.99.111:30303",
	"enode://cc67fe6390132583f2079f3d34e84124b619c7a4456437c55c13030fc09422bbdba1fddd3ebab71d514b103a43594504cf6c377c87465821913151afa0c4cd5b@149.28.164.205:30303",
	"enode://5ed9aa0bdbf3e2bf39f33307003f63ff5cba72f9a6888e27c2af3659272819b52fc85011210c1511ce7705d1606dd1170b28470387f2da3d3d678f5ea93ec19e@167.86.95.197:30305",
	"enode://458735831e90cfcdcd782b798cafb46d3a4d41969dd190f5bacc751a35f5bc26a9411f92e2e2a4d8536d34d64e4937c878f29080f64bb52afdb280392188161e@173.212.202.40:30305",
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
	// Upstream bootnodes
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

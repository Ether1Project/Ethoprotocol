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
        "enode://6055f924546a4f81b2d10d2e55df54f31fb5d2286dff11e0e7c108c10e796c7809e644e2797bfc31b0212cd2e9cda9bd966ebfad9bc75215552a32fdef3bfbcb@45.76.117.164:30301",     //AUSTALIA
	"enode://07bd358da373fd4d982517ad6879503bdc1885092830ee799609f442ebc0b08a99896eaf073ec7b14741545839f8a9e0f776812f1dd4b243cc152108090ef0f3@45.32.197.67:30303",
        "enode://b66778fe345930e14e5d7684bc2ba59cc1f4fa9ec6dad33d2f008d89d5091582f11d04d896986a81b3bb2ce907a1f562813c6c0740e14082f1360dca9d8ac2fb@140.82.57.96:30303",
	"enode://9cb32f5393dac554a60d5e4a062bb9dec5fcfbdfe730fe6047dc987720e456f409a1c92f525631f15bafe6760a4ce683073ae2d9d5b24cc46aa5b164627caf4d@199.247.21.43:30303",
	"enode://d8e3ffee500298645ee198a4e2cf4eb04a4e82b81b3a7627d17f134c03cf5f5d3c20bf73b1b4967a0b13f1b22f973837991063c75a6bb23f2ddce9c8489fbdf6@45.77.184.141:30303",
	"enode://1ed58103ebf40db5079e9100114de7520185bbfedaf88250e44e9e75066e26f8fc3ea5233a58249e276c8509517dcd856f52889ccadb1a5a172458edefef02c8@104.238.188.219:30303",
	"enode://8dd418438c7a7da7e9fa33e9129e1361d029e935623fe000af71b9c4169c2e8c9e6f315ac34029bfbf4e12076210cb2cd4cbe8362a76d4d127f5c59c58905085@45.63.99.111:30303",
	"enode://0b37be00538e8afa3d7db3150b8ac2aa6bbfd2e355e3b6609b31050cd7acac09f730b412fed6f853695424852a58d609ddf7bc206b3ebcd27001cc6c299ab89b@144.202.68.146:30303",
	"enode://f82195f83c29c94b32ed772a0d068766e8830268ebaf8a19dcaa1b323f9fa17dab6ff41368d9e768b4e171aa1285f44780747f8ef695beeebde99b411ad60607@108.61.91.160:30303",
	"enode://4df8875eaedc39e4d46d918ddf3d32b4ae44ebf3229e4b0c5dba5468433d7e8a57d50fbb38cf49d039084a28f9d56ab4cea33219edb2cf2cc917169aa9f33e63@95.179.129.96:30303",
	"enode://cc67fe6390132583f2079f3d34e84124b619c7a4456437c55c13030fc09422bbdba1fddd3ebab71d514b103a43594504cf6c377c87465821913151afa0c4cd5b@149.28.164.205:30303"
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

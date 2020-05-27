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

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://335a900592ae34ec50cc13a5434f8737dd5e29172f7bcfab071b35414cd48de4d1a51a49d177f85f4ff85f7edcb8ec137bee739afb6fc6053bad10478b8f3da9@206.189.30.232:30305",
	"enode://a557fe5cbad84b33a3fe35d46b9153b0514e7b6c533cf9ee9adf53f144799243a2100ad44bc8c12dd185549bdebb446a28b62dc390ba1ed20681cdbb97324cfc@95.179.163.200:30305",
	"enode://908d14e9791f57eb5db026c30777f8bb8be90f11b2a17643d8ba2c6d449fd13a84a11d0bce1d458b1834624a5f6b3af486bbf297cdffbb3697817e7c64cd5867@142.93.151.20:30305",
	"enode://07bd358da373fd4d982517ad6879503bdc1885092830ee799609f442ebc0b08a99896eaf073ec7b14741545839f8a9e0f776812f1dd4b243cc152108090ef0f3@45.32.197.67:30303",
	"enode://b66778fe345930e14e5d7684bc2ba59cc1f4fa9ec6dad33d2f008d89d5091582f11d04d896986a81b3bb2ce907a1f562813c6c0740e14082f1360dca9d8ac2fb@140.82.57.96:30303",
	"enode://9cb32f5393dac554a60d5e4a062bb9dec5fcfbdfe730fe6047dc987720e456f409a1c92f525631f15bafe6760a4ce683073ae2d9d5b24cc46aa5b164627caf4d@199.247.21.43:30303",
	"enode://d8e3ffee500298645ee198a4e2cf4eb04a4e82b81b3a7627d17f134c03cf5f5d3c20bf73b1b4967a0b13f1b22f973837991063c75a6bb23f2ddce9c8489fbdf6@45.77.184.141:30303",
	"enode://cc67fe6390132583f2079f3d34e84124b619c7a4456437c55c13030fc09422bbdba1fddd3ebab71d514b103a43594504cf6c377c87465821913151afa0c4cd5b@149.28.164.205:30303",
	"enode://5ed9aa0bdbf3e2bf39f33307003f63ff5cba72f9a6888e27c2af3659272819b52fc85011210c1511ce7705d1606dd1170b28470387f2da3d3d678f5ea93ec19e@167.86.95.197:30305",
	"enode://458735831e90cfcdcd782b798cafb46d3a4d41969dd190f5bacc751a35f5bc26a9411f92e2e2a4d8536d34d64e4937c878f29080f64bb52afdb280392188161e@173.212.202.40:30305",
	"enode://6d0f3acdc177ff57c3ec1795690f37a2d3520cd0766fd5b4d360e9cae906172df5714691ea2db5d571922481871980d14370ac5bf36201329d417810c5bc179d@142.44.246.43:30305",
	"enode://3422f8bc295521007dc0e1595d9f6726c433f0dd9dfdfce901fc20c65a5e7bd610a6c2ccad93354bb803c6000ae0de84704a3111289aaf5293cc02618d4bbca3@51.79.70.144:30305",
	"enode://6349cac48451fd38cb9b9a19bac2e5f6aaa626e215079eb0d9225689749b11950e05e175db382598eea13e98d123e0581838ce1d4575800b17801360273b59ce@51.77.150.202:30305",
	"enode://38db07788e8ad06dacb42125de2aaa89cab0ca965578ccc9711b4fea2a635ed8d2c948758b6df5d313bb8f831d1fab8e686e03aecfceb15729b93bed9f3bcc53@51.75.163.238:30305",
	"enode://8ee9dde384b316c5baa9616fe3da3f0e8f6091fd36141dd4175e88b720b884e6ef219c7e54bb6fee8e7d395315de609213bc412dbb5886b0bea1a018f7832ef0@51.38.131.241:30305",
	"enode://03944d60046265f36aa333373e36604570029dc0dc9518d4226ba2037ae33cc2c5dd6940ee22c3ce85ad8a3c5791f81b73530dbe77aacd22d9e25593c4a354c8@140.82.54.221:30305",
	"enode://03944d60046265f36aa333373e36604570029dc0dc9518d4226ba2037ae33cc2c5dd6940ee22c3ce85ad8a3c5791f81b73530dbe77aacd22d9e25593c4a354c8@149.28.167.176:30305",
	"enode://44a4cc298c2877201339d27a84b10eb67959bc089f04e98a240428a80598880501cb6fcbb43cea0157efeaa719450edcdff2233d3614ff72319510d49db6ade8@144.91.86.189:30305",
	// Cluster Node 1
	"enode://6d0f3acdc177ff57c3ec1795690f37a2d3520cd0766fd5b4d360e9cae906172df5714691ea2db5d571922481871980d14370ac5bf36201329d417810c5bc179d@142.44.246.43:30305",
	// Cluster Node 2
	"enode://3422f8bc295521007dc0e1595d9f6726c433f0dd9dfdfce901fc20c65a5e7bd610a6c2ccad93354bb803c6000ae0de84704a3111289aaf5293cc02618d4bbca3@51.79.70.144:30305",
	// Cluster Node 3
	"enode://6349cac48451fd38cb9b9a19bac2e5f6aaa626e215079eb0d9225689749b11950e05e175db382598eea13e98d123e0581838ce1d4575800b17801360273b59ce@51.77.150.202:30305",
	// Cluster Node 4
	"enode://6d52b88413881718cc2bad61a28f81d84d43140196cacb590c0a59d6a5aa31d1b33cee84b430aa28ecb69339f24dff6c11fffa23335ede60f1174f679025ea3c@164.68.107.82:30305",
	// Cluster Node 5
	"enode://66c70a5c9bd2418cf8a559073774a9ddb3160aef5f4cd5936243acc873cbe7e77679675eacb8d268aab643644566c3ad34512d831e17163cf3c4cb232466be14@164.68.98.94:30305",
	// Cluster Node 6
	"enode://544435e729af7dbc26f0343ccc128ebcca5b1f685de1e1ecd612a3f7778456e61304ea34190e7386d6204e14b904aa9a0ef153729f7a27ff50136aa8474c0ed9@164.68.108.54:30305",
	// Pistol's Bootnode
	"enode://91b89285c5f313b9badfb41fe309ee9b0d307a1976b687bd79dea6b344033bb4a690a0371640f79f1709f5199a492428318a16da854039b4a9fbafb3bd10e051@207.180.197.210:303",
	// Crypto_Saiyan's Bootnode
	"enode://6d9ed0b577af7162694fdc38098ebf037ac5984eb051f144ee23e15686c6396b6a1a1b43c8a105479de6742af088dbf3b543bc56a782b427cfb6cc0b60fede8b@144.91.117.137:30305",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}

const dnsPrefix = ""

// These DNS names provide bootstrap connectivity for public testnets and the mainnet.
// See https://github.com/ethereum/discv4-dns-lists for more information.
var KnownDNSNetworks = map[common.Hash]string{
	MainnetGenesisHash: dnsPrefix + "",
	TestnetGenesisHash: dnsPrefix + "",
	RinkebyGenesisHash: dnsPrefix + "",
	GoerliGenesisHash:  dnsPrefix + "",
}

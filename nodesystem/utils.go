// Copyright 2015 The go-ethereum Authors
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

package nodesystem

import (
	"fmt"
	"math/big"
	"crypto/ecdsa"
	"net"
	"net/url"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/enr"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/log"
)

var NodeFlag bool

func SetProtocolFlag(active bool) {
	NodeFlag = active
}

func CheckActiveNode(totalNodeCount uint64, hash common.Hash, blockNumber int64) {
	if NodeFlag {
		log.Info("node system is active", "active", "true")
		log.Info("validating node system reward candidates", "number", blockNumber, "hash", hash, "registered nodes", totalNodeCount)
	}
}

// GetNodePublicKey returns public key in a string format from *enode.Node
func GetNodePublicKey(n *enode.Node) string {
        var (
                scheme enr.ID
                nodeid string
                key    ecdsa.PublicKey
        )
        n.Load(&scheme)
        n.Load((*enode.Secp256k1)(&key))
        nid := n.ID()
        switch {
        case scheme == "v4" || key != ecdsa.PublicKey{}:
                nodeid = fmt.Sprintf("%x", crypto.FromECDSAPub(&key)[1:])
        default:
                nodeid = fmt.Sprintf("%s.%x", scheme, nid[:])
        }
        u := url.URL{Scheme: "enode"}
        if n.Incomplete() {
                u.Host = nodeid
        } else {
                addr := net.TCPAddr{IP: n.IP(), Port: n.TCP()}
                u.User = url.User(nodeid)
                u.Host = addr.String()
                if n.UDP() != n.TCP() {
                        u.RawQuery = "discport=" + strconv.Itoa(n.UDP())
                }
        }
        return u.User.String()
}

// GetNodeEnodeId returns public key in a string format from *enode.Node
func GetNodeEnodeId(n *enode.Node) string {
        var (
                scheme enr.ID
                nodeid string
                key    ecdsa.PublicKey
        )
        n.Load(&scheme)
        n.Load((*enode.Secp256k1)(&key))
        nid := n.ID()
        switch {
        case scheme == "v4" || key != ecdsa.PublicKey{}:
                nodeid = fmt.Sprintf("%x", crypto.FromECDSAPub(&key)[1:])
        default:
                nodeid = fmt.Sprintf("%s.%x", scheme, nid[:])
        }
        u := url.URL{Scheme: "enode"}
        if n.Incomplete() {
                u.Host = nodeid
        } else {
                addr := net.TCPAddr{IP: n.IP(), Port: n.UDP()}
                u.User = url.User(nodeid)
                u.Host = addr.String()
                if n.UDP() != n.TCP() {
                        u.RawQuery = "discport=" + strconv.Itoa(n.UDP())
                }
        }
        return u.User.String() + ":" + n.IP().String()
}

// GetNodePrivateKey returns private key in a ecdsa.PrivateKey format from *enode.Node
func GetNodePrivateKey(srv *p2p.Server) *ecdsa.PrivateKey {
        return srv.PrivateKey
}

// Clean up data
func stripCtlAndExtFromBytes(str string) string {
	b := make([]byte, len(str))
	var bl int
	for i := 0; i < len(str); i++ {
		c := str[i]
		if c >= 32 && c < 127 {
			b[bl] = c
			bl++
		}
	}
	return string(b[:bl])
}

// Calculate and return node remainder balance for payment share
func GetNodeRemainder(state *state.StateDB, nodeCount uint64, remainderAddress common.Address) *big.Int {

	remainderBalance := state.GetBalance(remainderAddress)

	if remainderBalance.Cmp(big.NewInt(0)) > 0 && nodeCount > 0 {

		// Disburse remainder funds over extended period using a full days block count as divisor
		var remainderPayment *big.Int
		remainderPayment = new(big.Int).Div(remainderBalance, big.NewInt(6646))

		return remainderPayment
	}
	return big.NewInt(0)
}

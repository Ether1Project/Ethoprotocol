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

package nodeprotocol

import (
	"math/big"

        "github.com/ethereum/go-ethereum/core/state"
        "github.com/ethereum/go-ethereum/common"
)

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

// Calculate and return node remainder balance payment share
func GetNodeRemainder(state *state.StateDB, nodeCount int64) *big.Int {

        remainderBalance := state.GetBalance(common.HexToAddress("0x0000000000000000000000000000000000000001"))

        if remainderBalance.Cmp(big.NewInt(0)) > 0 {

                var remainderPayment *big.Int
                remainderPayment = new(big.Int).Div(remainderBalance, big.NewInt(nodeCount))

                return remainderPayment
        }
        return big.NewInt(0)
}

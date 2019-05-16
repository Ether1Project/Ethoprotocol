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
	"fmt"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
)

var activeNode *node.Node

func ActiveNode() *node.Node {
        return activeNode
}

func SetActiveNode(stack *node.Node) {
        activeNode = stack
}

func ConfirmNodeActivity(nodeID string) (bool, error) {
	log.Info("Attempting to Contact Node", "NodeID", nodeID)
	destinationNode, err := enode.ParseV4(nodeID)
	if err != nil {
	        return true, fmt.Errorf("Node ID Format Error: %v", err)
	}
        // Try to contact node
	response := p2p.Resolve(activeNode.Server(), destinationNode)
	if response {
		log.Info("Node Contacted Successfully", "Node", "Verified")
	        return true, nil
	}
	return false, fmt.Errorf("Node Communication Error: %v", err)
}

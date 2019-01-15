// Copyright 2018 The go-ethereum Authors
// Copyright 2018 The Pirl Team <dev@pirl.io>
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

// Package core implements the Penalty System proposed by Pirl Team
package core

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/params"
)

var syncing bool

const (
	description         = "Penalty System (based on the PirlGuard by Pirl Team)"
	delayedBlockInfoLen = 3
	delayedBlockWarnLen = 15
)

var (
	blockDelayedMeter = metrics.NewRegisteredMeter("chain/block/delayed", nil)
	blockPenaltyMeter = metrics.NewRegisteredMeter("chain/block/penalty", nil)
)

// CheckDelayedChain will check possible 51% attack.
// Penalty System penalize newly inserted blocks.
// The amount of penalty depends on the amount of blocks mined by the malicious miner privatly.
func (bc *BlockChain) CheckDelayedChain(blocks types.Blocks, logonly, reverse bool) error {
	current := bc.CurrentBlock().NumberU64()

	if !syncing {
		head := rawdb.ReadHeadBlockHash(bc.db)
		if head == (common.Hash{}) {
			// Corrupt or empty database.
			return nil
		}
		// Get current block
		currentBlock := bc.GetBlockByHash(head)
		if currentBlock == nil {
			// Corrupt or empty database.
			return nil
		}
		// Get current fast block
		currentFastBlock := bc.CurrentFastBlock()

		// Setup sync status
		syncing = currentFastBlock.NumberU64() == currentBlock.NumberU64()
		log.Info("sync status", "sync", syncing)
	}

	var penalty uint64
	tipidx := 0
	if reverse {
		tipidx = len(blocks) - 1
	}
	if syncing && current > uint64(params.PenaltySystemBlock) && current > blocks[tipidx].NumberU64() && (current - blocks[tipidx].NumberU64()) > delayedBlockInfoLen {
		delayed, score := bc.penaltyForBlocks(blocks)
		logFn := log.Info
		if delayed > delayedBlockWarnLen {
			logFn = log.Warn
		}
		blockDelayedMeter.Mark(int64(delayed))
		blockPenaltyMeter.Mark(int64(score))
		penalty = score

		logFn("Checking the legitimity of the chain", "delayed chain length", delayed, "penalty", penalty, "description", description)
	} else {
		return nil
	}

	if !logonly && penalty >= params.DelayedBlockLength*(params.DelayedBlockLength+1)/2 {
		context := []interface{}{
			"penalty", penalty,
		}
		log.Error("Malicious Chain! We should reject it", context...)
		bc.setBadHash(blocks[tipidx], params.DelayedBlockLength)
		return ErrDelayTooHigh
	}

	return nil
}

func (bc *BlockChain) penaltyForBlocks(blocks types.Blocks) (uint64, uint64) {
	var sum, penalty, n uint64
	current := bc.CurrentBlock().NumberU64()
	sum = 0
	for _, b := range blocks {
		if current >= b.NumberU64() {
			penalty = current - b.NumberU64()
			n++;
		} else {
			penalty = 0
		}
		sum += penalty
		context := []interface{}{
			"head", current, "number", b.NumberU64(), "hash", b.Hash() , "penalty", penalty, "sum", sum,
		}

		log.Warn("Penalty check", context...)
	}
	return n, sum
}

func (bc *BlockChain) setBadHash(block *types.Block, minPenalty uint64) {
	current := bc.CurrentBlock().NumberU64()
	if current >= block.NumberU64() {
		penalty := current - block.NumberU64()
		if penalty >= minPenalty {
			BadHashes[block.Header().Hash()] = true
			log.Error("New Bad Hash", "block", block.NumberU64(), "hash", block.Header().Hash(), "penalty", penalty)
		}
	}
}

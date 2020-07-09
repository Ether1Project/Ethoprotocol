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
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"
	"sync"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	//"github.com/ethereum/go-ethereum/nodesystem"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

var validationMap map[common.Hash][][]byte
var mux = &sync.Mutex{}

type NodeValidation struct {
	Id          []byte            `json:"id"`
	Validations [][]byte          `json:"validations"`
}

//func CheckValidNodeTx(address common.Address, input []byte) bool {
func CheckValidNodeTx(tx *types.Transaction) bool {

	address := *tx.To()
	input := tx.Data()

	if address != params.NodeSystemContractAddress {
		return false
	}

	abi, err := abi.JSON(strings.NewReader(NodeSystemABI))
	if err != nil {
		log.Error("invalid node tx detected", "error", err)
		return false
	}

	// Decode tx input method signature
	decodedSig, err := hex.DecodeString(hex.EncodeToString(input)[2:10])
	if err != nil {
		log.Error("invalid node tx detected", "error", err)
		return false
	}

	// Recover method from signature and ABI
	method, err := abi.MethodById(decodedSig)
	if err != nil {
		log.Error("invalid node tx detected", "error", err)
		return false
	}

	// Decode tx input payload
	decodedData, err := hex.DecodeString(hex.EncodeToString(input)[10:])
	if err != nil {
		log.Error("invalid node tx detected", "error", err)
		return false
	}

	type FunctionInputs struct {
		Validations [][]byte
		Id []byte
		Hash common.Hash
	}

	var data FunctionInputs

	err = method.Inputs.Unpack(&data, decodedData)
	if err != nil {
		log.Error("invalid node tx detected", "error", err)
		return false
	}

	for _, sig := range data.Validations {
		if !ValidateNodeSignatureByHash(data.Id, sig, data.Hash) {
			return false
		}
	}
	return true
}


// SignNodeValidation is used to respond to a peer/next node's validation request
// A signed validation using enode private key signals an unequivocal validation of activity
func SignNodeValidation(privateKey *ecdsa.PrivateKey, data []byte, blockHash common.Hash) []byte {
	data = append(data, ":"...)
	data = append(data, blockHash.String()...)
	hash := crypto.Keccak256(data)
        signedValidation, err := crypto.Sign(hash, privateKey)
        if err != nil {
		log.Error("sign node validation error", "error", err)
        }
	return signedValidation
}

// ValidateNodeSignatureByHash is used to verify validation signatures of a block when only a block hash
// is known
func ValidateNodeSignatureByHash(nodeId []byte, signedValidation []byte, hash common.Hash) bool {
	recoveredPub, err := crypto.Ecrecover(crypto.Keccak256(nodeId), signedValidation)
	if err != nil {
		log.Error("node signature validation by hash error", "error", err)
	}
	recoveredId, _ := crypto.UnmarshalPubkey(recoveredPub)
	recoveredIdString := fmt.Sprintf("%x", crypto.FromECDSAPub(recoveredId)[1:])

	state, err := GetStateAt(hash)
	if err != nil {
		log.Error("error retrieving state db", "hash", hash)
		return false
	}

	collateralizedPeerGroup := GetCollateralizedHashedGroup(state, hash)

	if _, ok := collateralizedPeerGroup[common.BytesToHash([]byte(recoveredIdString))]; ok {
		log.Info("node signature validation", "valid", "true", "author", recoveredIdString)
		return true
	}
	log.Warn("node signature validation", "valid", "false", "author", recoveredIdString)
	return false
}

// ValidateNodeSignature is used to verify validation signatures when a node validation tx
// is recevied to decentrally validate a nodes activity
func ValidateNodeSignature(nodeId []byte, signedValidation []byte, validationId []byte, blockHash common.Hash) bool {
	nodeId = append(nodeId, ":"...)
	nodeId = append(nodeId, blockHash.String()...)
	recoveredPub, err := crypto.Ecrecover(crypto.Keccak256(nodeId), signedValidation)
	if err != nil {
		log.Error("node signature validation error", "error", err)
	}
	recoveredId, _ := crypto.UnmarshalPubkey(recoveredPub)
	recoveredIdString := fmt.Sprintf("%x", crypto.FromECDSAPub(recoveredId)[1:])

	if common.BytesToHash(validationId) == common.BytesToHash([]byte(recoveredIdString)) {
		log.Info("node signature validation", "valid", "true", "author", recoveredIdString)
		return true
	}
	log.Warn("node signature validation", "valid", "false", "author", recoveredIdString)
	return false
}

func AddValidationSignature(hash common.Hash, signedValidation []byte) {
	mux.Lock()
	defer mux.Unlock()
	if len(validationMap) == 0 {
		validationMap = make(map[common.Hash][][]byte)
	}
	if validations, ok := validationMap[hash]; ok {
		validations = append(validations, signedValidation)
		if len(validations) >= params.MinNodeValidations {
			nodeValidations := NodeValidation{Id: []byte(GetNodePublicKey(ActiveNode().Server().Self())), Validations: validations}
			SendSignedNodeTx(GetNodePrivateKey(ActiveNode().Server()), nodeValidations, hash)
			delete(validationMap, hash)
		} else {
			validationMap[hash] = validations
		}
	} else {
		var validations [][]byte
		validations = append(validations, signedValidation)
		validationMap[hash] = validations
	}
}

func SendSignedNodeTx(privateKey *ecdsa.PrivateKey, validations NodeValidation, hash common.Hash) {
	client, err := ethclient.Dial(ActiveNode().IPCEndpoint())
	if err != nil {
		log.Error("node validation tx error", "error", err)
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error("node validation tx error", "error", "cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return
	}

	from := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), from)
	if err != nil {
		log.Error("node validation tx error", "error", err)
		return
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error("node validation tx error", "error", err)
		return
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	instance, err := NewNodeSystem(params.NodeSystemContractAddress, client)
	if err != nil {
		log.Error("node validation tx error", "error", err)
		return
	}

	tx, err := instance.NodeCheckIn(auth, validations.Validations, validations.Id, hash.Bytes())
	if err != nil {
		log.Error("node validation tx error", "error", err)
		return
	}
	log.Info("node validation tx sent", "hash", tx.Hash().Hex())
}

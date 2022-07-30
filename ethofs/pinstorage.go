// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethofs

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PinStorageABI is the input ABI used to generate the binding from.
const PinStorageABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"changeOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"set\",\"type\":\"uint32\"}],\"name\":\"SetReplicationFactor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pin\",\"type\":\"string\"}],\"name\":\"PinRemove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deleteContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pin\",\"type\":\"string\"},{\"name\":\"size\",\"type\":\"uint32\"}],\"name\":\"PinAdd\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PinCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ReplicationFactor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Pins\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"pin\",\"type\":\"string\"}],\"name\":\"GetPinSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// PinStorage is an auto generated Go binding around an Ethereum contract.
type PinStorage struct {
	PinStorageCaller     // Read-only binding to the contract
	PinStorageTransactor // Write-only binding to the contract
	PinStorageFilterer   // Log filterer for contract events
}

// PinStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type PinStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PinStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PinStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PinStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PinStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PinStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PinStorageSession struct {
	Contract     *PinStorage       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PinStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PinStorageCallerSession struct {
	Contract *PinStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PinStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PinStorageTransactorSession struct {
	Contract     *PinStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PinStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type PinStorageRaw struct {
	Contract *PinStorage // Generic contract binding to access the raw methods on
}

// PinStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PinStorageCallerRaw struct {
	Contract *PinStorageCaller // Generic read-only contract binding to access the raw methods on
}

// PinStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PinStorageTransactorRaw struct {
	Contract *PinStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPinStorage creates a new instance of PinStorage, bound to a specific deployed contract.
func NewPinStorage(address common.Address, backend bind.ContractBackend) (*PinStorage, error) {
	contract, err := bindPinStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PinStorage{PinStorageCaller: PinStorageCaller{contract: contract}, PinStorageTransactor: PinStorageTransactor{contract: contract}, PinStorageFilterer: PinStorageFilterer{contract: contract}}, nil
}

// NewPinStorageCaller creates a new read-only instance of PinStorage, bound to a specific deployed contract.
func NewPinStorageCaller(address common.Address, caller bind.ContractCaller) (*PinStorageCaller, error) {
	contract, err := bindPinStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PinStorageCaller{contract: contract}, nil
}

// NewPinStorageTransactor creates a new write-only instance of PinStorage, bound to a specific deployed contract.
func NewPinStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*PinStorageTransactor, error) {
	contract, err := bindPinStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PinStorageTransactor{contract: contract}, nil
}

// NewPinStorageFilterer creates a new log filterer instance of PinStorage, bound to a specific deployed contract.
func NewPinStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*PinStorageFilterer, error) {
	contract, err := bindPinStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PinStorageFilterer{contract: contract}, nil
}

// bindPinStorage binds a generic wrapper to an already deployed contract.
func bindPinStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PinStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PinStorage *PinStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PinStorage.Contract.PinStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PinStorage *PinStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PinStorage.Contract.PinStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PinStorage *PinStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PinStorage.Contract.PinStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PinStorage *PinStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PinStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PinStorage *PinStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PinStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PinStorage *PinStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PinStorage.Contract.contract.Transact(opts, method, params...)
}

// GetPinSize is a free data retrieval call binding the contract method 0xfde1505e.
//
// Solidity: function GetPinSize(string pin) view returns(uint32)
func (_PinStorage *PinStorageCaller) GetPinSize(opts *bind.CallOpts, pin string) (uint32, error) {
	var out []interface{}
	err := _PinStorage.contract.Call(opts, &out, "GetPinSize", pin)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetPinSize is a free data retrieval call binding the contract method 0xfde1505e.
//
// Solidity: function GetPinSize(string pin) view returns(uint32)
func (_PinStorage *PinStorageSession) GetPinSize(pin string) (uint32, error) {
	return _PinStorage.Contract.GetPinSize(&_PinStorage.CallOpts, pin)
}

// GetPinSize is a free data retrieval call binding the contract method 0xfde1505e.
//
// Solidity: function GetPinSize(string pin) view returns(uint32)
func (_PinStorage *PinStorageCallerSession) GetPinSize(pin string) (uint32, error) {
	return _PinStorage.Contract.GetPinSize(&_PinStorage.CallOpts, pin)
}

// PinCount is a free data retrieval call binding the contract method 0x92927c60.
//
// Solidity: function PinCount() view returns(uint32)
func (_PinStorage *PinStorageCaller) PinCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PinStorage.contract.Call(opts, &out, "PinCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// PinCount is a free data retrieval call binding the contract method 0x92927c60.
//
// Solidity: function PinCount() view returns(uint32)
func (_PinStorage *PinStorageSession) PinCount() (uint32, error) {
	return _PinStorage.Contract.PinCount(&_PinStorage.CallOpts)
}

// PinCount is a free data retrieval call binding the contract method 0x92927c60.
//
// Solidity: function PinCount() view returns(uint32)
func (_PinStorage *PinStorageCallerSession) PinCount() (uint32, error) {
	return _PinStorage.Contract.PinCount(&_PinStorage.CallOpts)
}

// Pins is a free data retrieval call binding the contract method 0xe4416802.
//
// Solidity: function Pins(uint256 ) view returns(string)
func (_PinStorage *PinStorageCaller) Pins(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _PinStorage.contract.Call(opts, &out, "Pins", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Pins is a free data retrieval call binding the contract method 0xe4416802.
//
// Solidity: function Pins(uint256 ) view returns(string)
func (_PinStorage *PinStorageSession) Pins(arg0 *big.Int) (string, error) {
	return _PinStorage.Contract.Pins(&_PinStorage.CallOpts, arg0)
}

// Pins is a free data retrieval call binding the contract method 0xe4416802.
//
// Solidity: function Pins(uint256 ) view returns(string)
func (_PinStorage *PinStorageCallerSession) Pins(arg0 *big.Int) (string, error) {
	return _PinStorage.Contract.Pins(&_PinStorage.CallOpts, arg0)
}

// ReplicationFactor is a free data retrieval call binding the contract method 0xdb837e73.
//
// Solidity: function ReplicationFactor() view returns(uint32)
func (_PinStorage *PinStorageCaller) ReplicationFactor(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PinStorage.contract.Call(opts, &out, "ReplicationFactor")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ReplicationFactor is a free data retrieval call binding the contract method 0xdb837e73.
//
// Solidity: function ReplicationFactor() view returns(uint32)
func (_PinStorage *PinStorageSession) ReplicationFactor() (uint32, error) {
	return _PinStorage.Contract.ReplicationFactor(&_PinStorage.CallOpts)
}

// ReplicationFactor is a free data retrieval call binding the contract method 0xdb837e73.
//
// Solidity: function ReplicationFactor() view returns(uint32)
func (_PinStorage *PinStorageCallerSession) ReplicationFactor() (uint32, error) {
	return _PinStorage.Contract.ReplicationFactor(&_PinStorage.CallOpts)
}

// PinAdd is a paid mutator transaction binding the contract method 0x8d036731.
//
// Solidity: function PinAdd(string pin, uint32 size) returns()
func (_PinStorage *PinStorageTransactor) PinAdd(opts *bind.TransactOpts, pin string, size uint32) (*types.Transaction, error) {
	return _PinStorage.contract.Transact(opts, "PinAdd", pin, size)
}

// PinAdd is a paid mutator transaction binding the contract method 0x8d036731.
//
// Solidity: function PinAdd(string pin, uint32 size) returns()
func (_PinStorage *PinStorageSession) PinAdd(pin string, size uint32) (*types.Transaction, error) {
	return _PinStorage.Contract.PinAdd(&_PinStorage.TransactOpts, pin, size)
}

// PinAdd is a paid mutator transaction binding the contract method 0x8d036731.
//
// Solidity: function PinAdd(string pin, uint32 size) returns()
func (_PinStorage *PinStorageTransactorSession) PinAdd(pin string, size uint32) (*types.Transaction, error) {
	return _PinStorage.Contract.PinAdd(&_PinStorage.TransactOpts, pin, size)
}

// PinRemove is a paid mutator transaction binding the contract method 0x3f0854a7.
//
// Solidity: function PinRemove(string pin) returns()
func (_PinStorage *PinStorageTransactor) PinRemove(opts *bind.TransactOpts, pin string) (*types.Transaction, error) {
	return _PinStorage.contract.Transact(opts, "PinRemove", pin)
}

// PinRemove is a paid mutator transaction binding the contract method 0x3f0854a7.
//
// Solidity: function PinRemove(string pin) returns()
func (_PinStorage *PinStorageSession) PinRemove(pin string) (*types.Transaction, error) {
	return _PinStorage.Contract.PinRemove(&_PinStorage.TransactOpts, pin)
}

// PinRemove is a paid mutator transaction binding the contract method 0x3f0854a7.
//
// Solidity: function PinRemove(string pin) returns()
func (_PinStorage *PinStorageTransactorSession) PinRemove(pin string) (*types.Transaction, error) {
	return _PinStorage.Contract.PinRemove(&_PinStorage.TransactOpts, pin)
}

// SetReplicationFactor is a paid mutator transaction binding the contract method 0x1986a58c.
//
// Solidity: function SetReplicationFactor(uint32 set) returns()
func (_PinStorage *PinStorageTransactor) SetReplicationFactor(opts *bind.TransactOpts, set uint32) (*types.Transaction, error) {
	return _PinStorage.contract.Transact(opts, "SetReplicationFactor", set)
}

// SetReplicationFactor is a paid mutator transaction binding the contract method 0x1986a58c.
//
// Solidity: function SetReplicationFactor(uint32 set) returns()
func (_PinStorage *PinStorageSession) SetReplicationFactor(set uint32) (*types.Transaction, error) {
	return _PinStorage.Contract.SetReplicationFactor(&_PinStorage.TransactOpts, set)
}

// SetReplicationFactor is a paid mutator transaction binding the contract method 0x1986a58c.
//
// Solidity: function SetReplicationFactor(uint32 set) returns()
func (_PinStorage *PinStorageTransactorSession) SetReplicationFactor(set uint32) (*types.Transaction, error) {
	return _PinStorage.Contract.SetReplicationFactor(&_PinStorage.TransactOpts, set)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address newOperator) returns()
func (_PinStorage *PinStorageTransactor) ChangeOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _PinStorage.contract.Transact(opts, "changeOperator", newOperator)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address newOperator) returns()
func (_PinStorage *PinStorageSession) ChangeOperator(newOperator common.Address) (*types.Transaction, error) {
	return _PinStorage.Contract.ChangeOperator(&_PinStorage.TransactOpts, newOperator)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address newOperator) returns()
func (_PinStorage *PinStorageTransactorSession) ChangeOperator(newOperator common.Address) (*types.Transaction, error) {
	return _PinStorage.Contract.ChangeOperator(&_PinStorage.TransactOpts, newOperator)
}

// DeleteContract is a paid mutator transaction binding the contract method 0x5a58cd4c.
//
// Solidity: function deleteContract() returns()
func (_PinStorage *PinStorageTransactor) DeleteContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PinStorage.contract.Transact(opts, "deleteContract")
}

// DeleteContract is a paid mutator transaction binding the contract method 0x5a58cd4c.
//
// Solidity: function deleteContract() returns()
func (_PinStorage *PinStorageSession) DeleteContract() (*types.Transaction, error) {
	return _PinStorage.Contract.DeleteContract(&_PinStorage.TransactOpts)
}

// DeleteContract is a paid mutator transaction binding the contract method 0x5a58cd4c.
//
// Solidity: function deleteContract() returns()
func (_PinStorage *PinStorageTransactorSession) DeleteContract() (*types.Transaction, error) {
	return _PinStorage.Contract.DeleteContract(&_PinStorage.TransactOpts)
}

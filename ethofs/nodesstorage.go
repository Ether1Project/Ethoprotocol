// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// NodesStorageABI is the input ABI used to generate the binding from.
const NodesStorageABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"changeOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\"}],\"name\":\"nodeRemove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"set\",\"type\":\"uint32\"}],\"name\":\"SetSN\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SNcount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MNcount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nodeID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deleteContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GNcount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"multihash\",\"type\":\"string\"}],\"name\":\"nodeAdd\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"set\",\"type\":\"uint32\"}],\"name\":\"SetMN\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"nodes\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"set\",\"type\":\"uint32\"}],\"name\":\"SetGN\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// NodesStorageBin is the compiled bytecode used for deploying new contracts.
const NodesStorageBin = `608060405234801561001057600080fd5b5033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008060006101000a81548163ffffffff021916908363ffffffff16021790555060008060046101000a81548163ffffffff021916908363ffffffff16021790555060008060086101000a81548163ffffffff021916908363ffffffff160217905550600080600c6101000a81548163ffffffff021916908363ffffffff160217905550610d75806101266000396000f3006080604052600436106100ba576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306394c9b146100bf5780630b05e3ac1461010257806310897c7b1461013557806314ad4061146101685780633440b0381461019f5780633dff2228146101d65780635a58cd4c1461020d57806387ac7a6014610224578063b68869601461025b578063bcaadfe2146102d6578063c87b5ba614610309578063feac17a1146103b5575b600080fd5b3480156100cb57600080fd5b50610100600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506103e8565b005b34801561010e57600080fd5b50610133600480360381019080803563ffffffff169060200190929190505050610488565b005b34801561014157600080fd5b50610166600480360381019080803563ffffffff169060200190929190505050610671565b005b34801561017457600080fd5b5061017d61074b565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b3480156101ab57600080fd5b506101b4610761565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b3480156101e257600080fd5b506101eb610776565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b34801561021957600080fd5b5061022261078c565b005b34801561023057600080fd5b50610239610823565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b34801561026757600080fd5b506102d4600480360381019080803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610839565b005b3480156102e257600080fd5b50610307600480360381019080803563ffffffff169060200190929190505050610972565b005b34801561031557600080fd5b5061033a600480360381019080803563ffffffff169060200190929190505050610a4b565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561037a57808201518184015260208101905061035f565b50505050905090810190601f1680156103a75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156103c157600080fd5b506103e6600480360381019080803563ffffffff169060200190929190505050610afb565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561044457600080fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141580156105365750600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b1561054057600080fd5b600360008363ffffffff1663ffffffff168152602001908152602001600020600061056b9190610bd5565b8190505b6000600c9054906101000a900463ffffffff1663ffffffff168163ffffffff16101561063157600360006001830163ffffffff1663ffffffff168152602001908152602001600020600360008363ffffffff1663ffffffff16815260200190815260200160002090805460018160011615610100020316600290046105f5929190610c1d565b50600360006001830163ffffffff1663ffffffff16815260200190815260200160002060006106249190610bd5565b808060010191505061056f565b6000600c81819054906101000a900463ffffffff16809291906001900391906101000a81548163ffffffff021916908363ffffffff160217905550505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415801561071d5750600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b1561072757600080fd5b80600060046101000a81548163ffffffff021916908363ffffffff16021790555050565b600060049054906101000a900463ffffffff1681565b6000809054906101000a900463ffffffff1681565b6000600c9054906101000a900463ffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156107e857600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b600060089054906101000a900463ffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141580156108e55750600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b156108ef57600080fd5b6000600c81819054906101000a900463ffffffff168092919060010191906101000a81548163ffffffff021916908363ffffffff16021790555050806003600080600c9054906101000a900463ffffffff1663ffffffff1663ffffffff168152602001908152602001600020908051906020019061096e929190610ca4565b5050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614158015610a1e5750600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b15610a2857600080fd5b806000806101000a81548163ffffffff021916908363ffffffff16021790555050565b60036020528060005260406000206000915090508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610af35780601f10610ac857610100808354040283529160200191610af3565b820191906000526020600020905b815481529060010190602001808311610ad657829003601f168201915b505050505081565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614158015610ba75750600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b15610bb157600080fd5b80600060086101000a81548163ffffffff021916908363ffffffff16021790555050565b50805460018160011615610100020316600290046000825580601f10610bfb5750610c1a565b601f016020900490600052602060002090810190610c199190610d24565b5b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610c565780548555610c93565b82800160010185558215610c9357600052602060002091601f016020900482015b82811115610c92578254825591600101919060010190610c77565b5b509050610ca09190610d24565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610ce557805160ff1916838001178555610d13565b82800160010185558215610d13579182015b82811115610d12578251825591602001919060010190610cf7565b5b509050610d209190610d24565b5090565b610d4691905b80821115610d42576000816000905550600101610d2a565b5090565b905600a165627a7a72305820dc7ff4db562623a17cce08244a1f2aa3ee7ac57f0d4d453d97e2760357515a4d0029`

// DeployNodesStorage deploys a new Ethereum contract, binding an instance of NodesStorage to it.
func DeployNodesStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NodesStorage, error) {
	parsed, err := abi.JSON(strings.NewReader(NodesStorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NodesStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NodesStorage{NodesStorageCaller: NodesStorageCaller{contract: contract}, NodesStorageTransactor: NodesStorageTransactor{contract: contract}, NodesStorageFilterer: NodesStorageFilterer{contract: contract}}, nil
}

// NodesStorage is an auto generated Go binding around an Ethereum contract.
type NodesStorage struct {
	NodesStorageCaller     // Read-only binding to the contract
	NodesStorageTransactor // Write-only binding to the contract
	NodesStorageFilterer   // Log filterer for contract events
}

// NodesStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodesStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodesStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodesStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodesStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodesStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodesStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodesStorageSession struct {
	Contract     *NodesStorage     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodesStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodesStorageCallerSession struct {
	Contract *NodesStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// NodesStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodesStorageTransactorSession struct {
	Contract     *NodesStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NodesStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodesStorageRaw struct {
	Contract *NodesStorage // Generic contract binding to access the raw methods on
}

// NodesStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodesStorageCallerRaw struct {
	Contract *NodesStorageCaller // Generic read-only contract binding to access the raw methods on
}

// NodesStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodesStorageTransactorRaw struct {
	Contract *NodesStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodesStorage creates a new instance of NodesStorage, bound to a specific deployed contract.
func NewNodesStorage(address common.Address, backend bind.ContractBackend) (*NodesStorage, error) {
	contract, err := bindNodesStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodesStorage{NodesStorageCaller: NodesStorageCaller{contract: contract}, NodesStorageTransactor: NodesStorageTransactor{contract: contract}, NodesStorageFilterer: NodesStorageFilterer{contract: contract}}, nil
}

// NewNodesStorageCaller creates a new read-only instance of NodesStorage, bound to a specific deployed contract.
func NewNodesStorageCaller(address common.Address, caller bind.ContractCaller) (*NodesStorageCaller, error) {
	contract, err := bindNodesStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodesStorageCaller{contract: contract}, nil
}

// NewNodesStorageTransactor creates a new write-only instance of NodesStorage, bound to a specific deployed contract.
func NewNodesStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*NodesStorageTransactor, error) {
	contract, err := bindNodesStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodesStorageTransactor{contract: contract}, nil
}

// NewNodesStorageFilterer creates a new log filterer instance of NodesStorage, bound to a specific deployed contract.
func NewNodesStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*NodesStorageFilterer, error) {
	contract, err := bindNodesStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodesStorageFilterer{contract: contract}, nil
}

// bindNodesStorage binds a generic wrapper to an already deployed contract.
func bindNodesStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodesStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodesStorage *NodesStorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodesStorage.Contract.NodesStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodesStorage *NodesStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodesStorage.Contract.NodesStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodesStorage *NodesStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodesStorage.Contract.NodesStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodesStorage *NodesStorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodesStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodesStorage *NodesStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodesStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodesStorage *NodesStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodesStorage.Contract.contract.Transact(opts, method, params...)
}

// GNcount is a free data retrieval call binding the contract method 0x87ac7a60.
//
// Solidity: function GNcount() constant returns(uint32)
func (_NodesStorage *NodesStorageCaller) GNcount(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _NodesStorage.contract.Call(opts, out, "GNcount")
	return *ret0, err
}

// GNcount is a free data retrieval call binding the contract method 0x87ac7a60.
//
// Solidity: function GNcount() constant returns(uint32)
func (_NodesStorage *NodesStorageSession) GNcount() (uint32, error) {
	return _NodesStorage.Contract.GNcount(&_NodesStorage.CallOpts)
}

// GNcount is a free data retrieval call binding the contract method 0x87ac7a60.
//
// Solidity: function GNcount() constant returns(uint32)
func (_NodesStorage *NodesStorageCallerSession) GNcount() (uint32, error) {
	return _NodesStorage.Contract.GNcount(&_NodesStorage.CallOpts)
}

// MNcount is a free data retrieval call binding the contract method 0x3440b038.
//
// Solidity: function MNcount() constant returns(uint32)
func (_NodesStorage *NodesStorageCaller) MNcount(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _NodesStorage.contract.Call(opts, out, "MNcount")
	return *ret0, err
}

// MNcount is a free data retrieval call binding the contract method 0x3440b038.
//
// Solidity: function MNcount() constant returns(uint32)
func (_NodesStorage *NodesStorageSession) MNcount() (uint32, error) {
	return _NodesStorage.Contract.MNcount(&_NodesStorage.CallOpts)
}

// MNcount is a free data retrieval call binding the contract method 0x3440b038.
//
// Solidity: function MNcount() constant returns(uint32)
func (_NodesStorage *NodesStorageCallerSession) MNcount() (uint32, error) {
	return _NodesStorage.Contract.MNcount(&_NodesStorage.CallOpts)
}

// SNcount is a free data retrieval call binding the contract method 0x14ad4061.
//
// Solidity: function SNcount() constant returns(uint32)
func (_NodesStorage *NodesStorageCaller) SNcount(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _NodesStorage.contract.Call(opts, out, "SNcount")
	return *ret0, err
}

// SNcount is a free data retrieval call binding the contract method 0x14ad4061.
//
// Solidity: function SNcount() constant returns(uint32)
func (_NodesStorage *NodesStorageSession) SNcount() (uint32, error) {
	return _NodesStorage.Contract.SNcount(&_NodesStorage.CallOpts)
}

// SNcount is a free data retrieval call binding the contract method 0x14ad4061.
//
// Solidity: function SNcount() constant returns(uint32)
func (_NodesStorage *NodesStorageCallerSession) SNcount() (uint32, error) {
	return _NodesStorage.Contract.SNcount(&_NodesStorage.CallOpts)
}

// NodeID is a free data retrieval call binding the contract method 0x3dff2228.
//
// Solidity: function nodeID() constant returns(uint32)
func (_NodesStorage *NodesStorageCaller) NodeID(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _NodesStorage.contract.Call(opts, out, "nodeID")
	return *ret0, err
}

// NodeID is a free data retrieval call binding the contract method 0x3dff2228.
//
// Solidity: function nodeID() constant returns(uint32)
func (_NodesStorage *NodesStorageSession) NodeID() (uint32, error) {
	return _NodesStorage.Contract.NodeID(&_NodesStorage.CallOpts)
}

// NodeID is a free data retrieval call binding the contract method 0x3dff2228.
//
// Solidity: function nodeID() constant returns(uint32)
func (_NodesStorage *NodesStorageCallerSession) NodeID() (uint32, error) {
	return _NodesStorage.Contract.NodeID(&_NodesStorage.CallOpts)
}

// Nodes is a free data retrieval call binding the contract method 0xc87b5ba6.
//
// Solidity: function nodes( uint32) constant returns(string)
func (_NodesStorage *NodesStorageCaller) Nodes(opts *bind.CallOpts, arg0 uint32) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _NodesStorage.contract.Call(opts, out, "nodes", arg0)
	return *ret0, err
}

// Nodes is a free data retrieval call binding the contract method 0xc87b5ba6.
//
// Solidity: function nodes( uint32) constant returns(string)
func (_NodesStorage *NodesStorageSession) Nodes(arg0 uint32) (string, error) {
	return _NodesStorage.Contract.Nodes(&_NodesStorage.CallOpts, arg0)
}

// Nodes is a free data retrieval call binding the contract method 0xc87b5ba6.
//
// Solidity: function nodes( uint32) constant returns(string)
func (_NodesStorage *NodesStorageCallerSession) Nodes(arg0 uint32) (string, error) {
	return _NodesStorage.Contract.Nodes(&_NodesStorage.CallOpts, arg0)
}

// SetGN is a paid mutator transaction binding the contract method 0xfeac17a1.
//
// Solidity: function SetGN(set uint32) returns()
func (_NodesStorage *NodesStorageTransactor) SetGN(opts *bind.TransactOpts, set uint32) (*types.Transaction, error) {
	return _NodesStorage.contract.Transact(opts, "SetGN", set)
}

// SetGN is a paid mutator transaction binding the contract method 0xfeac17a1.
//
// Solidity: function SetGN(set uint32) returns()
func (_NodesStorage *NodesStorageSession) SetGN(set uint32) (*types.Transaction, error) {
	return _NodesStorage.Contract.SetGN(&_NodesStorage.TransactOpts, set)
}

// SetGN is a paid mutator transaction binding the contract method 0xfeac17a1.
//
// Solidity: function SetGN(set uint32) returns()
func (_NodesStorage *NodesStorageTransactorSession) SetGN(set uint32) (*types.Transaction, error) {
	return _NodesStorage.Contract.SetGN(&_NodesStorage.TransactOpts, set)
}

// SetMN is a paid mutator transaction binding the contract method 0xbcaadfe2.
//
// Solidity: function SetMN(set uint32) returns()
func (_NodesStorage *NodesStorageTransactor) SetMN(opts *bind.TransactOpts, set uint32) (*types.Transaction, error) {
	return _NodesStorage.contract.Transact(opts, "SetMN", set)
}

// SetMN is a paid mutator transaction binding the contract method 0xbcaadfe2.
//
// Solidity: function SetMN(set uint32) returns()
func (_NodesStorage *NodesStorageSession) SetMN(set uint32) (*types.Transaction, error) {
	return _NodesStorage.Contract.SetMN(&_NodesStorage.TransactOpts, set)
}

// SetMN is a paid mutator transaction binding the contract method 0xbcaadfe2.
//
// Solidity: function SetMN(set uint32) returns()
func (_NodesStorage *NodesStorageTransactorSession) SetMN(set uint32) (*types.Transaction, error) {
	return _NodesStorage.Contract.SetMN(&_NodesStorage.TransactOpts, set)
}

// SetSN is a paid mutator transaction binding the contract method 0x10897c7b.
//
// Solidity: function SetSN(set uint32) returns()
func (_NodesStorage *NodesStorageTransactor) SetSN(opts *bind.TransactOpts, set uint32) (*types.Transaction, error) {
	return _NodesStorage.contract.Transact(opts, "SetSN", set)
}

// SetSN is a paid mutator transaction binding the contract method 0x10897c7b.
//
// Solidity: function SetSN(set uint32) returns()
func (_NodesStorage *NodesStorageSession) SetSN(set uint32) (*types.Transaction, error) {
	return _NodesStorage.Contract.SetSN(&_NodesStorage.TransactOpts, set)
}

// SetSN is a paid mutator transaction binding the contract method 0x10897c7b.
//
// Solidity: function SetSN(set uint32) returns()
func (_NodesStorage *NodesStorageTransactorSession) SetSN(set uint32) (*types.Transaction, error) {
	return _NodesStorage.Contract.SetSN(&_NodesStorage.TransactOpts, set)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(newOperator address) returns()
func (_NodesStorage *NodesStorageTransactor) ChangeOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _NodesStorage.contract.Transact(opts, "changeOperator", newOperator)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(newOperator address) returns()
func (_NodesStorage *NodesStorageSession) ChangeOperator(newOperator common.Address) (*types.Transaction, error) {
	return _NodesStorage.Contract.ChangeOperator(&_NodesStorage.TransactOpts, newOperator)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(newOperator address) returns()
func (_NodesStorage *NodesStorageTransactorSession) ChangeOperator(newOperator common.Address) (*types.Transaction, error) {
	return _NodesStorage.Contract.ChangeOperator(&_NodesStorage.TransactOpts, newOperator)
}

// DeleteContract is a paid mutator transaction binding the contract method 0x5a58cd4c.
//
// Solidity: function deleteContract() returns()
func (_NodesStorage *NodesStorageTransactor) DeleteContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodesStorage.contract.Transact(opts, "deleteContract")
}

// DeleteContract is a paid mutator transaction binding the contract method 0x5a58cd4c.
//
// Solidity: function deleteContract() returns()
func (_NodesStorage *NodesStorageSession) DeleteContract() (*types.Transaction, error) {
	return _NodesStorage.Contract.DeleteContract(&_NodesStorage.TransactOpts)
}

// DeleteContract is a paid mutator transaction binding the contract method 0x5a58cd4c.
//
// Solidity: function deleteContract() returns()
func (_NodesStorage *NodesStorageTransactorSession) DeleteContract() (*types.Transaction, error) {
	return _NodesStorage.Contract.DeleteContract(&_NodesStorage.TransactOpts)
}

// NodeAdd is a paid mutator transaction binding the contract method 0xb6886960.
//
// Solidity: function nodeAdd(multihash string) returns()
func (_NodesStorage *NodesStorageTransactor) NodeAdd(opts *bind.TransactOpts, multihash string) (*types.Transaction, error) {
	return _NodesStorage.contract.Transact(opts, "nodeAdd", multihash)
}

// NodeAdd is a paid mutator transaction binding the contract method 0xb6886960.
//
// Solidity: function nodeAdd(multihash string) returns()
func (_NodesStorage *NodesStorageSession) NodeAdd(multihash string) (*types.Transaction, error) {
	return _NodesStorage.Contract.NodeAdd(&_NodesStorage.TransactOpts, multihash)
}

// NodeAdd is a paid mutator transaction binding the contract method 0xb6886960.
//
// Solidity: function nodeAdd(multihash string) returns()
func (_NodesStorage *NodesStorageTransactorSession) NodeAdd(multihash string) (*types.Transaction, error) {
	return _NodesStorage.Contract.NodeAdd(&_NodesStorage.TransactOpts, multihash)
}

// NodeRemove is a paid mutator transaction binding the contract method 0x0b05e3ac.
//
// Solidity: function nodeRemove(id uint32) returns()
func (_NodesStorage *NodesStorageTransactor) NodeRemove(opts *bind.TransactOpts, id uint32) (*types.Transaction, error) {
	return _NodesStorage.contract.Transact(opts, "nodeRemove", id)
}

// NodeRemove is a paid mutator transaction binding the contract method 0x0b05e3ac.
//
// Solidity: function nodeRemove(id uint32) returns()
func (_NodesStorage *NodesStorageSession) NodeRemove(id uint32) (*types.Transaction, error) {
	return _NodesStorage.Contract.NodeRemove(&_NodesStorage.TransactOpts, id)
}

// NodeRemove is a paid mutator transaction binding the contract method 0x0b05e3ac.
//
// Solidity: function nodeRemove(id uint32) returns()
func (_NodesStorage *NodesStorageTransactorSession) NodeRemove(id uint32) (*types.Transaction, error) {
	return _NodesStorage.Contract.NodeRemove(&_NodesStorage.TransactOpts, id)
}

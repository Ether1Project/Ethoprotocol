// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethofs

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// PinStorageABI is the input ABI used to generate the binding from.
const PinStorageABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"changeOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"set\",\"type\":\"uint32\"}],\"name\":\"SetReplicationFactor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pin\",\"type\":\"string\"}],\"name\":\"PinRemove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deleteContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pin\",\"type\":\"string\"},{\"name\":\"size\",\"type\":\"uint32\"}],\"name\":\"PinAdd\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PinCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ReplicationFactor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Pins\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"pin\",\"type\":\"string\"}],\"name\":\"GetPinSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// PinStorageBin is the compiled bytecode used for deploying new contracts.
const PinStorageBin = `0x608060405234801561001057600080fd5b506000805460018054600160a060020a0319163390811790915567ffffffffffffffff19680100000000000000009091027fffffffff0000000000000000000000000000000000000000ffffffffffffffff90921691909117169055610a6d8061007b6000396000f3006080604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306394c9b811461009d5780631986a58c146100c05780633f0854a7146100de5780635a58cd4c1461013d5780638d0367311461015257806392927c60146101b9578063db837e73146101e7578063e4416802146101fc578063fde1505e14610289575b600080fd5b3480156100a957600080fd5b506100be600160a060020a03600435166102e8565b005b3480156100cc57600080fd5b506100be63ffffffff6004351661033a565b3480156100ea57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100be9436949293602493928401919081908401838280828437600092019190915250929550610392945050505050565b34801561014957600080fd5b506100be61047b565b34801561015e57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100be94369492936024939284019190819084018382808284376000920191909152509295505050903563ffffffff1691506104b89050565b3480156101c557600080fd5b506101ce6105db565b6040805163ffffffff9092168252519081900360200190f35b3480156101f357600080fd5b506101ce6105ef565b34801561020857600080fd5b506102146004356105fb565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561024e578181015183820152602001610236565b50505050905090810190601f16801561027b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561029557600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526101ce94369492936024939284019190819084018382808284376000920191909152509295506106a2945050505050565b600054680100000000000000009004600160a060020a0316331461030b57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054680100000000000000009004600160a060020a0316331480159061036c5750600154600160a060020a03163314155b1561037657600080fd5b6000805463ffffffff191663ffffffff92909216919091179055565b600054680100000000000000009004600160a060020a031633148015906103c45750600154600160a060020a03163314155b156103ce57600080fd5b6103d781610710565b6002816040518082805190602001908083835b602083106104095780518252601f1990920191602091820191016103ea565b51815160001960209485036101000a810191821691199290921617909152939091019586526040519586900301909420805463ffffffff191690556000805464010000000080820463ffffffff90811690940190931690920267ffffffff000000001990921691909117905550505050565b600054680100000000000000009004600160a060020a0316331461049e57600080fd5b600054680100000000000000009004600160a060020a0316ff5b600054680100000000000000009004600160a060020a031633148015906104ea5750600154600160a060020a03163314155b156104f457600080fd5b60038054600181018083556000929092528351610538917fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0190602086019061089e565b5050806002836040518082805190602001908083835b6020831061056d5780518252601f19909201916020918201910161054e565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220805463ffffffff191663ffffffff94851617905550506000805467ffffffff00000000198116640100000000918290048416600101909316029190911790555050565b600054640100000000900463ffffffff1681565b60005463ffffffff1681565b600380548290811061060957fe5b600091825260209182902001805460408051601f600260001961010060018716150201909416939093049283018590048502810185019091528181529350909183018282801561069a5780601f1061066f5761010080835404028352916020019161069a565b820191906000526020600020905b81548152906001019060200180831161067d57829003601f168201915b505050505081565b60006002826040518082805190602001908083835b602083106106d65780518252601f1990920191602091820191016106b7565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205463ffffffff16949350505050565b600061071b8261072a565b905061072681610821565b5050565b6000805b826040518082805190602001908083835b6020831061075e5780518252601f19909201916020918201910161073f565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060038281548110151561079a57fe5b9060005260206000200160405180828054600181600116156101000203166002900480156107ff5780601f106107dd5761010080835404028352918201916107ff565b820191906000526020600020905b8154815290600101906020018083116107eb575b5050915050604051809103902014151561081b5760010161072e565b92915050565b6003546000190181101561088b57600380546001830190811061084057fe5b9060005260206000200160038281548110151561085957fe5b90600052602060002001908054600181600116156101000203166002900461088292919061091c565b50600101610821565b6003805490610726906000198301610991565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106108df57805160ff191683800117855561090c565b8280016001018555821561090c579182015b8281111561090c5782518255916020019190600101906108f1565b506109189291506109ba565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610955578054855561090c565b8280016001018555821561090c57600052602060002091601f016020900482015b8281111561090c578254825591600101919060010190610976565b8154818355818111156109b5576000838152602090206109b59181019083016109d7565b505050565b6109d491905b8082111561091857600081556001016109c0565b90565b6109d491905b808211156109185760006109f182826109fa565b506001016109dd565b50805460018160011615610100020316600290046000825580601f10610a205750610a3e565b601f016020900490600052602060002090810190610a3e91906109ba565b505600a165627a7a72305820e7f0ef6da9592cf23c9d04a374beb1a7db555267fd01ce08e60827477debfdcc0029`

// DeployPinStorage deploys a new Ethereum contract, binding an instance of PinStorage to it.
func DeployPinStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PinStorage, error) {
	parsed, err := abi.JSON(strings.NewReader(PinStorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PinStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PinStorage{PinStorageCaller: PinStorageCaller{contract: contract}, PinStorageTransactor: PinStorageTransactor{contract: contract}, PinStorageFilterer: PinStorageFilterer{contract: contract}}, nil
}

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
func (_PinStorage *PinStorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_PinStorage *PinStorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
// Solidity: function GetPinSize(pin string) constant returns(uint32)
func (_PinStorage *PinStorageCaller) GetPinSize(opts *bind.CallOpts, pin string) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _PinStorage.contract.Call(opts, out, "GetPinSize", pin)
	return *ret0, err
}

// GetPinSize is a free data retrieval call binding the contract method 0xfde1505e.
//
// Solidity: function GetPinSize(pin string) constant returns(uint32)
func (_PinStorage *PinStorageSession) GetPinSize(pin string) (uint32, error) {
	return _PinStorage.Contract.GetPinSize(&_PinStorage.CallOpts, pin)
}

// GetPinSize is a free data retrieval call binding the contract method 0xfde1505e.
//
// Solidity: function GetPinSize(pin string) constant returns(uint32)
func (_PinStorage *PinStorageCallerSession) GetPinSize(pin string) (uint32, error) {
	return _PinStorage.Contract.GetPinSize(&_PinStorage.CallOpts, pin)
}

// PinCount is a free data retrieval call binding the contract method 0x92927c60.
//
// Solidity: function PinCount() constant returns(uint32)
func (_PinStorage *PinStorageCaller) PinCount(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _PinStorage.contract.Call(opts, out, "PinCount")
	return *ret0, err
}

// PinCount is a free data retrieval call binding the contract method 0x92927c60.
//
// Solidity: function PinCount() constant returns(uint32)
func (_PinStorage *PinStorageSession) PinCount() (uint32, error) {
	return _PinStorage.Contract.PinCount(&_PinStorage.CallOpts)
}

// PinCount is a free data retrieval call binding the contract method 0x92927c60.
//
// Solidity: function PinCount() constant returns(uint32)
func (_PinStorage *PinStorageCallerSession) PinCount() (uint32, error) {
	return _PinStorage.Contract.PinCount(&_PinStorage.CallOpts)
}

// Pins is a free data retrieval call binding the contract method 0xe4416802.
//
// Solidity: function Pins( uint256) constant returns(string)
func (_PinStorage *PinStorageCaller) Pins(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PinStorage.contract.Call(opts, out, "Pins", arg0)
	return *ret0, err
}

// Pins is a free data retrieval call binding the contract method 0xe4416802.
//
// Solidity: function Pins( uint256) constant returns(string)
func (_PinStorage *PinStorageSession) Pins(arg0 *big.Int) (string, error) {
	return _PinStorage.Contract.Pins(&_PinStorage.CallOpts, arg0)
}

// Pins is a free data retrieval call binding the contract method 0xe4416802.
//
// Solidity: function Pins( uint256) constant returns(string)
func (_PinStorage *PinStorageCallerSession) Pins(arg0 *big.Int) (string, error) {
	return _PinStorage.Contract.Pins(&_PinStorage.CallOpts, arg0)
}

// ReplicationFactor is a free data retrieval call binding the contract method 0xdb837e73.
//
// Solidity: function ReplicationFactor() constant returns(uint32)
func (_PinStorage *PinStorageCaller) ReplicationFactor(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _PinStorage.contract.Call(opts, out, "ReplicationFactor")
	return *ret0, err
}

// ReplicationFactor is a free data retrieval call binding the contract method 0xdb837e73.
//
// Solidity: function ReplicationFactor() constant returns(uint32)
func (_PinStorage *PinStorageSession) ReplicationFactor() (uint32, error) {
	return _PinStorage.Contract.ReplicationFactor(&_PinStorage.CallOpts)
}

// ReplicationFactor is a free data retrieval call binding the contract method 0xdb837e73.
//
// Solidity: function ReplicationFactor() constant returns(uint32)
func (_PinStorage *PinStorageCallerSession) ReplicationFactor() (uint32, error) {
	return _PinStorage.Contract.ReplicationFactor(&_PinStorage.CallOpts)
}

// PinAdd is a paid mutator transaction binding the contract method 0x8d036731.
//
// Solidity: function PinAdd(pin string, size uint32) returns()
func (_PinStorage *PinStorageTransactor) PinAdd(opts *bind.TransactOpts, pin string, size uint32) (*types.Transaction, error) {
	return _PinStorage.contract.Transact(opts, "PinAdd", pin, size)
}

// PinAdd is a paid mutator transaction binding the contract method 0x8d036731.
//
// Solidity: function PinAdd(pin string, size uint32) returns()
func (_PinStorage *PinStorageSession) PinAdd(pin string, size uint32) (*types.Transaction, error) {
	return _PinStorage.Contract.PinAdd(&_PinStorage.TransactOpts, pin, size)
}

// PinAdd is a paid mutator transaction binding the contract method 0x8d036731.
//
// Solidity: function PinAdd(pin string, size uint32) returns()
func (_PinStorage *PinStorageTransactorSession) PinAdd(pin string, size uint32) (*types.Transaction, error) {
	return _PinStorage.Contract.PinAdd(&_PinStorage.TransactOpts, pin, size)
}

// PinRemove is a paid mutator transaction binding the contract method 0x3f0854a7.
//
// Solidity: function PinRemove(pin string) returns()
func (_PinStorage *PinStorageTransactor) PinRemove(opts *bind.TransactOpts, pin string) (*types.Transaction, error) {
	return _PinStorage.contract.Transact(opts, "PinRemove", pin)
}

// PinRemove is a paid mutator transaction binding the contract method 0x3f0854a7.
//
// Solidity: function PinRemove(pin string) returns()
func (_PinStorage *PinStorageSession) PinRemove(pin string) (*types.Transaction, error) {
	return _PinStorage.Contract.PinRemove(&_PinStorage.TransactOpts, pin)
}

// PinRemove is a paid mutator transaction binding the contract method 0x3f0854a7.
//
// Solidity: function PinRemove(pin string) returns()
func (_PinStorage *PinStorageTransactorSession) PinRemove(pin string) (*types.Transaction, error) {
	return _PinStorage.Contract.PinRemove(&_PinStorage.TransactOpts, pin)
}

// SetReplicationFactor is a paid mutator transaction binding the contract method 0x1986a58c.
//
// Solidity: function SetReplicationFactor(set uint32) returns()
func (_PinStorage *PinStorageTransactor) SetReplicationFactor(opts *bind.TransactOpts, set uint32) (*types.Transaction, error) {
	return _PinStorage.contract.Transact(opts, "SetReplicationFactor", set)
}

// SetReplicationFactor is a paid mutator transaction binding the contract method 0x1986a58c.
//
// Solidity: function SetReplicationFactor(set uint32) returns()
func (_PinStorage *PinStorageSession) SetReplicationFactor(set uint32) (*types.Transaction, error) {
	return _PinStorage.Contract.SetReplicationFactor(&_PinStorage.TransactOpts, set)
}

// SetReplicationFactor is a paid mutator transaction binding the contract method 0x1986a58c.
//
// Solidity: function SetReplicationFactor(set uint32) returns()
func (_PinStorage *PinStorageTransactorSession) SetReplicationFactor(set uint32) (*types.Transaction, error) {
	return _PinStorage.Contract.SetReplicationFactor(&_PinStorage.TransactOpts, set)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(newOperator address) returns()
func (_PinStorage *PinStorageTransactor) ChangeOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _PinStorage.contract.Transact(opts, "changeOperator", newOperator)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(newOperator address) returns()
func (_PinStorage *PinStorageSession) ChangeOperator(newOperator common.Address) (*types.Transaction, error) {
	return _PinStorage.Contract.ChangeOperator(&_PinStorage.TransactOpts, newOperator)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(newOperator address) returns()
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

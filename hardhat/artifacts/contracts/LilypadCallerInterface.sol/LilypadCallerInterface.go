// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package LilypadCallerInterface

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// LilypadCallerInterfaceMetaData contains all meta data concerning the LilypadCallerInterface contract.
var LilypadCallerInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_jobId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_ipfsResult\",\"type\":\"string\"}],\"name\":\"lilypadReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LilypadCallerInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use LilypadCallerInterfaceMetaData.ABI instead.
var LilypadCallerInterfaceABI = LilypadCallerInterfaceMetaData.ABI

// LilypadCallerInterface is an auto generated Go binding around an Ethereum contract.
type LilypadCallerInterface struct {
	LilypadCallerInterfaceCaller     // Read-only binding to the contract
	LilypadCallerInterfaceTransactor // Write-only binding to the contract
	LilypadCallerInterfaceFilterer   // Log filterer for contract events
}

// LilypadCallerInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type LilypadCallerInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadCallerInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LilypadCallerInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadCallerInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LilypadCallerInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadCallerInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LilypadCallerInterfaceSession struct {
	Contract     *LilypadCallerInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LilypadCallerInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LilypadCallerInterfaceCallerSession struct {
	Contract *LilypadCallerInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// LilypadCallerInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LilypadCallerInterfaceTransactorSession struct {
	Contract     *LilypadCallerInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// LilypadCallerInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type LilypadCallerInterfaceRaw struct {
	Contract *LilypadCallerInterface // Generic contract binding to access the raw methods on
}

// LilypadCallerInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LilypadCallerInterfaceCallerRaw struct {
	Contract *LilypadCallerInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// LilypadCallerInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LilypadCallerInterfaceTransactorRaw struct {
	Contract *LilypadCallerInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLilypadCallerInterface creates a new instance of LilypadCallerInterface, bound to a specific deployed contract.
func NewLilypadCallerInterface(address common.Address, backend bind.ContractBackend) (*LilypadCallerInterface, error) {
	contract, err := bindLilypadCallerInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LilypadCallerInterface{LilypadCallerInterfaceCaller: LilypadCallerInterfaceCaller{contract: contract}, LilypadCallerInterfaceTransactor: LilypadCallerInterfaceTransactor{contract: contract}, LilypadCallerInterfaceFilterer: LilypadCallerInterfaceFilterer{contract: contract}}, nil
}

// NewLilypadCallerInterfaceCaller creates a new read-only instance of LilypadCallerInterface, bound to a specific deployed contract.
func NewLilypadCallerInterfaceCaller(address common.Address, caller bind.ContractCaller) (*LilypadCallerInterfaceCaller, error) {
	contract, err := bindLilypadCallerInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadCallerInterfaceCaller{contract: contract}, nil
}

// NewLilypadCallerInterfaceTransactor creates a new write-only instance of LilypadCallerInterface, bound to a specific deployed contract.
func NewLilypadCallerInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*LilypadCallerInterfaceTransactor, error) {
	contract, err := bindLilypadCallerInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadCallerInterfaceTransactor{contract: contract}, nil
}

// NewLilypadCallerInterfaceFilterer creates a new log filterer instance of LilypadCallerInterface, bound to a specific deployed contract.
func NewLilypadCallerInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*LilypadCallerInterfaceFilterer, error) {
	contract, err := bindLilypadCallerInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LilypadCallerInterfaceFilterer{contract: contract}, nil
}

// bindLilypadCallerInterface binds a generic wrapper to an already deployed contract.
func bindLilypadCallerInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LilypadCallerInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadCallerInterface *LilypadCallerInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadCallerInterface.Contract.LilypadCallerInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadCallerInterface *LilypadCallerInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadCallerInterface.Contract.LilypadCallerInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadCallerInterface *LilypadCallerInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadCallerInterface.Contract.LilypadCallerInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadCallerInterface *LilypadCallerInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadCallerInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadCallerInterface *LilypadCallerInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadCallerInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadCallerInterface *LilypadCallerInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadCallerInterface.Contract.contract.Transact(opts, method, params...)
}

// LilypadReceiver is a paid mutator transaction binding the contract method 0x35c764e3.
//
// Solidity: function lilypadReceiver(address _from, uint256 _jobId, string _ipfsResult) returns()
func (_LilypadCallerInterface *LilypadCallerInterfaceTransactor) LilypadReceiver(opts *bind.TransactOpts, _from common.Address, _jobId *big.Int, _ipfsResult string) (*types.Transaction, error) {
	return _LilypadCallerInterface.contract.Transact(opts, "lilypadReceiver", _from, _jobId, _ipfsResult)
}

// LilypadReceiver is a paid mutator transaction binding the contract method 0x35c764e3.
//
// Solidity: function lilypadReceiver(address _from, uint256 _jobId, string _ipfsResult) returns()
func (_LilypadCallerInterface *LilypadCallerInterfaceSession) LilypadReceiver(_from common.Address, _jobId *big.Int, _ipfsResult string) (*types.Transaction, error) {
	return _LilypadCallerInterface.Contract.LilypadReceiver(&_LilypadCallerInterface.TransactOpts, _from, _jobId, _ipfsResult)
}

// LilypadReceiver is a paid mutator transaction binding the contract method 0x35c764e3.
//
// Solidity: function lilypadReceiver(address _from, uint256 _jobId, string _ipfsResult) returns()
func (_LilypadCallerInterface *LilypadCallerInterfaceTransactorSession) LilypadReceiver(_from common.Address, _jobId *big.Int, _ipfsResult string) (*types.Transaction, error) {
	return _LilypadCallerInterface.Contract.LilypadReceiver(&_LilypadCallerInterface.TransactOpts, _from, _jobId, _ipfsResult)
}


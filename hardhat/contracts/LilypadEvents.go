// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"jobName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"params\",\"type\":\"string\"}],\"name\":\"NewBacalhauJobSubmitted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"jobID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ipfsID\",\"type\":\"string\"}],\"name\":\"returnResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_jobName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"params\",\"type\":\"string\"}],\"name\":\"runBacalhauJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610447806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063390109341461003b578063fe41ff3b14610057575b600080fd5b61005560048036038101906100509190610280565b610073565b005b610071600480360381019061006c9190610280565b610078565b005b505050565b60405161008490610362565b60405180910390207fbd4c84f97376a916a5ac1d6ba01328a813fc80455c92026c217af8c04f8e3bb1336040516100bb91906103e3565b60405180910390a2505050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610107826100dc565b9050919050565b610117816100fc565b811461012257600080fd5b50565b6000813590506101348161010e565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61018d82610144565b810181811067ffffffffffffffff821117156101ac576101ab610155565b5b80604052505050565b60006101bf6100c8565b90506101cb8282610184565b919050565b600067ffffffffffffffff8211156101eb576101ea610155565b5b6101f482610144565b9050602081019050919050565b82818337600083830152505050565b600061022361021e846101d0565b6101b5565b90508281526020810184848401111561023f5761023e61013f565b5b61024a848285610201565b509392505050565b600082601f8301126102675761026661013a565b5b8135610277848260208601610210565b91505092915050565b600080600060608486031215610299576102986100d2565b5b60006102a786828701610125565b935050602084013567ffffffffffffffff8111156102c8576102c76100d7565b5b6102d486828701610252565b925050604084013567ffffffffffffffff8111156102f5576102f46100d7565b5b61030186828701610252565b9150509250925092565b600081905092915050565b7f537461626c65446966667573696f6e4750550000000000000000000000000000600082015250565b600061034c60128361030b565b915061035782610316565b601282019050919050565b600061036d8261033f565b9150819050919050565b610380816100fc565b82525050565b600082825260208201905092915050565b7f5261696e626f7720556e69636f726e0000000000000000000000000000000000600082015250565b60006103cd600f83610386565b91506103d882610397565b602082019050919050565b60006040820190506103f86000830184610377565b8181036020830152610409816103c0565b90509291505056fea264697066735822122028585cd089d84cc2d1e3f55eb8f4051accbe5335c803bb750f571b918ab9963164736f6c63430008110033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// ReturnResults is a paid mutator transaction binding the contract method 0x39010934.
//
// Solidity: function returnResults(address to, string jobID, string ipfsID) returns()
func (_Contracts *ContractsTransactor) ReturnResults(opts *bind.TransactOpts, to common.Address, jobID string, ipfsID string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "returnResults", to, jobID, ipfsID)
}

// ReturnResults is a paid mutator transaction binding the contract method 0x39010934.
//
// Solidity: function returnResults(address to, string jobID, string ipfsID) returns()
func (_Contracts *ContractsSession) ReturnResults(to common.Address, jobID string, ipfsID string) (*types.Transaction, error) {
	return _Contracts.Contract.ReturnResults(&_Contracts.TransactOpts, to, jobID, ipfsID)
}

// ReturnResults is a paid mutator transaction binding the contract method 0x39010934.
//
// Solidity: function returnResults(address to, string jobID, string ipfsID) returns()
func (_Contracts *ContractsTransactorSession) ReturnResults(to common.Address, jobID string, ipfsID string) (*types.Transaction, error) {
	return _Contracts.Contract.ReturnResults(&_Contracts.TransactOpts, to, jobID, ipfsID)
}

// RunBacalhauJob is a paid mutator transaction binding the contract method 0xfe41ff3b.
//
// Solidity: function runBacalhauJob(address from, string _jobName, string params) returns()
func (_Contracts *ContractsTransactor) RunBacalhauJob(opts *bind.TransactOpts, from common.Address, _jobName string, params string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "runBacalhauJob", from, _jobName, params)
}

// RunBacalhauJob is a paid mutator transaction binding the contract method 0xfe41ff3b.
//
// Solidity: function runBacalhauJob(address from, string _jobName, string params) returns()
func (_Contracts *ContractsSession) RunBacalhauJob(from common.Address, _jobName string, params string) (*types.Transaction, error) {
	return _Contracts.Contract.RunBacalhauJob(&_Contracts.TransactOpts, from, _jobName, params)
}

// RunBacalhauJob is a paid mutator transaction binding the contract method 0xfe41ff3b.
//
// Solidity: function runBacalhauJob(address from, string _jobName, string params) returns()
func (_Contracts *ContractsTransactorSession) RunBacalhauJob(from common.Address, _jobName string, params string) (*types.Transaction, error) {
	return _Contracts.Contract.RunBacalhauJob(&_Contracts.TransactOpts, from, _jobName, params)
}

// ContractsNewBacalhauJobSubmittedIterator is returned from FilterNewBacalhauJobSubmitted and is used to iterate over the raw logs and unpacked data for NewBacalhauJobSubmitted events raised by the Contracts contract.
type ContractsNewBacalhauJobSubmittedIterator struct {
	Event *ContractsNewBacalhauJobSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsNewBacalhauJobSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsNewBacalhauJobSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsNewBacalhauJobSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsNewBacalhauJobSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsNewBacalhauJobSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsNewBacalhauJobSubmitted represents a NewBacalhauJobSubmitted event raised by the Contracts contract.
type ContractsNewBacalhauJobSubmitted struct {
	From    common.Address
	JobName common.Hash
	Params  string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewBacalhauJobSubmitted is a free log retrieval operation binding the contract event 0xbd4c84f97376a916a5ac1d6ba01328a813fc80455c92026c217af8c04f8e3bb1.
//
// Solidity: event NewBacalhauJobSubmitted(address from, string indexed jobName, string params)
func (_Contracts *ContractsFilterer) FilterNewBacalhauJobSubmitted(opts *bind.FilterOpts, jobName []string) (*ContractsNewBacalhauJobSubmittedIterator, error) {

	var jobNameRule []interface{}
	for _, jobNameItem := range jobName {
		jobNameRule = append(jobNameRule, jobNameItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "NewBacalhauJobSubmitted", jobNameRule)
	if err != nil {
		return nil, err
	}
	return &ContractsNewBacalhauJobSubmittedIterator{contract: _Contracts.contract, event: "NewBacalhauJobSubmitted", logs: logs, sub: sub}, nil
}

// WatchNewBacalhauJobSubmitted is a free log subscription operation binding the contract event 0xbd4c84f97376a916a5ac1d6ba01328a813fc80455c92026c217af8c04f8e3bb1.
//
// Solidity: event NewBacalhauJobSubmitted(address from, string indexed jobName, string params)
func (_Contracts *ContractsFilterer) WatchNewBacalhauJobSubmitted(opts *bind.WatchOpts, sink chan<- *ContractsNewBacalhauJobSubmitted, jobName []string) (event.Subscription, error) {

	var jobNameRule []interface{}
	for _, jobNameItem := range jobName {
		jobNameRule = append(jobNameRule, jobNameItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "NewBacalhauJobSubmitted", jobNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsNewBacalhauJobSubmitted)
				if err := _Contracts.contract.UnpackLog(event, "NewBacalhauJobSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewBacalhauJobSubmitted is a log parse operation binding the contract event 0xbd4c84f97376a916a5ac1d6ba01328a813fc80455c92026c217af8c04f8e3bb1.
//
// Solidity: event NewBacalhauJobSubmitted(address from, string indexed jobName, string params)
func (_Contracts *ContractsFilterer) ParseNewBacalhauJobSubmitted(log types.Log) (*ContractsNewBacalhauJobSubmitted, error) {
	event := new(ContractsNewBacalhauJobSubmitted)
	if err := _Contracts.contract.UnpackLog(event, "NewBacalhauJobSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}


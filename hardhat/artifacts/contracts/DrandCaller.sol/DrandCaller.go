// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DrandCaller

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

// DrandCallerMetaData contains all meta data concerning the DrandCaller contract.
var DrandCallerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"RandomNumberGenerated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximum\",\"type\":\"uint256\"}],\"name\":\"getRandomNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_jobId\",\"type\":\"uint256\"},{\"internalType\":\"enumLilypadResultType\",\"name\":\"_resultType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_result\",\"type\":\"string\"}],\"name\":\"lilypadReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001270380380620012708339818101604052810190620000379190620000e8565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200011a565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000b08262000083565b9050919050565b620000c281620000a3565b8114620000ce57600080fd5b50565b600081519050620000e281620000b7565b92915050565b6000602082840312156200010157620001006200007e565b5b60006200011184828501620000d1565b91505092915050565b611146806200012a6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063511deb471461003b5780636e68fc0a14610057575b600080fd5b610055600480360381019061005091906109fa565b610073565b005b610071600480360381019061006c9190610a82565b610145565b005b6001600381111561008757610086610ac2565b5b83600381111561009a57610099610ac2565b5b146100a457600080fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16146100fc57600080fd5b7f252851c72bba973a3eff4c55d976631b74c5e463b736ab3ba7ea37cbe40bbc19846101288484610252565b604051610136929190610b00565b60405180910390a15050505050565b6000610150836105e8565b9050600061015d836105e8565b9050600060405180610100016040528060e0815260200161103160e09139838360405180610100016040528060c58152602001610f6c60c591396040516020016101aa9493929190610bc0565b604051602081830303815290604052905060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d992cac8308360016040518463ffffffff1660e01b815260040161021993929190610cbf565b600060405180830381600087803b15801561023357600080fd5b505af1158015610247573d6000803e3d6000fd5b505050505050505050565b6000366000848491509150604082829050111561026e57600080fd5b6000828290501161027e57600080fd5b6000805b838390508160ff1610156105db57600482901b915083838260ff168181106102ad576102ac610cfd565b5b9050013560f81c60f81b60f81c60ff16826102c89190610d5b565b91507f410000000000000000000000000000000000000000000000000000000000000084848360ff1681811061030157610300610cfd565b5b9050013560f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161015801561039c57507f460000000000000000000000000000000000000000000000000000000000000084848360ff1681811061036d5761036c610cfd565b5b9050013560f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191611155b156103b5576037826103ae9190610d8f565b91506105c8565b7f610000000000000000000000000000000000000000000000000000000000000084848360ff168181106103ec576103eb610cfd565b5b9050013560f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161015801561048757507f660000000000000000000000000000000000000000000000000000000000000084848360ff1681811061045857610457610cfd565b5b9050013560f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191611155b156104a0576057826104999190610d8f565b91506105c7565b7f300000000000000000000000000000000000000000000000000000000000000084848360ff168181106104d7576104d6610cfd565b5b9050013560f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161015801561057257507f390000000000000000000000000000000000000000000000000000000000000084848360ff1681811061054357610542610cfd565b5b9050013560f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191611155b1561058b576030826105849190610d8f565b91506105c6565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105bd90610e0f565b60405180910390fd5b5b5b80806105d390610e3c565b915050610282565b5080935050505092915050565b60606105ff8260016105f985610606565b01610696565b9050919050565b600080600090506000608084901c111561062857608083901c92506010810190505b6000604084901c111561064357604083901c92506008810190505b6000602084901c111561065e57602083901c92506004810190505b6000601084901c111561067957601083901c92506002810190505b6000600884901c111561068d576001810190505b80915050919050565b6060600060028360026106a99190610e65565b6106b39190610d5b565b67ffffffffffffffff8111156106cc576106cb610ea7565b5b6040519080825280601f01601f1916602001820160405280156106fe5781602001600182028036833780820191505090505b5090507f30000000000000000000000000000000000000000000000000000000000000008160008151811061073657610735610cfd565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f78000000000000000000000000000000000000000000000000000000000000008160018151811061079a57610799610cfd565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600060018460026107da9190610e65565b6107e49190610d5b565b90505b6001811115610884577f3031323334353637383961626364656600000000000000000000000000000000600f86166010811061082657610825610cfd565b5b1a60f81b82828151811061083d5761083c610cfd565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c94508061087d90610ed6565b90506107e7565b50600084146108c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108bf90610f4b565b60405180910390fd5b8091505092915050565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610907826108dc565b9050919050565b610917816108fc565b811461092257600080fd5b50565b6000813590506109348161090e565b92915050565b6000819050919050565b61094d8161093a565b811461095857600080fd5b50565b60008135905061096a81610944565b92915050565b6004811061097d57600080fd5b50565b60008135905061098f81610970565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f8401126109ba576109b9610995565b5b8235905067ffffffffffffffff8111156109d7576109d661099a565b5b6020830191508360018202830111156109f3576109f261099f565b5b9250929050565b600080600080600060808688031215610a1657610a156108d2565b5b6000610a2488828901610925565b9550506020610a358882890161095b565b9450506040610a4688828901610980565b935050606086013567ffffffffffffffff811115610a6757610a666108d7565b5b610a73888289016109a4565b92509250509295509295909350565b60008060408385031215610a9957610a986108d2565b5b6000610aa78582860161095b565b9250506020610ab88582860161095b565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b610afa8161093a565b82525050565b6000604082019050610b156000830185610af1565b610b226020830184610af1565b9392505050565b600081519050919050565b600081905092915050565b60005b83811015610b5d578082015181840152602081019050610b42565b60008484015250505050565b6000610b7482610b29565b610b7e8185610b34565b9350610b8e818560208601610b3f565b80840191505092915050565b7f222c220000000000000000000000000000000000000000000000000000000000815250565b6000610bcc8287610b69565b9150610bd88286610b69565b9150610be382610b9a565b600382019150610bf38285610b69565b9150610bff8284610b69565b915081905095945050505050565b610c16816108fc565b82525050565b600082825260208201905092915050565b6000601f19601f8301169050919050565b6000610c4982610b29565b610c538185610c1c565b9350610c63818560208601610b3f565b610c6c81610c2d565b840191505092915050565b60048110610c8857610c87610ac2565b5b50565b6000819050610c9982610c77565b919050565b6000610ca982610c8b565b9050919050565b610cb981610c9e565b82525050565b6000606082019050610cd46000830186610c0d565b8181036020830152610ce68185610c3e565b9050610cf56040830184610cb0565b949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610d668261093a565b9150610d718361093a565b9250828201905080821115610d8957610d88610d2c565b5b92915050565b6000610d9a8261093a565b9150610da58361093a565b9250828203905081811115610dbd57610dbc610d2c565b5b92915050565b7f696e76616c69642063686172616374657220696e2068657820737472696e6700600082015250565b6000610df9601f83610c1c565b9150610e0482610dc3565b602082019050919050565b60006020820190508181036000830152610e2881610dec565b9050919050565b600060ff82169050919050565b6000610e4782610e2f565b915060ff8203610e5a57610e59610d2c565b5b600182019050919050565b6000610e708261093a565b9150610e7b8361093a565b9250828202610e898161093a565b91508282048414831517610ea057610e9f610d2c565b5b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000610ee18261093a565b915060008203610ef457610ef3610d2c565b5b600182039050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b6000610f35602083610c1c565b9150610f4082610eff565b602082019050919050565b60006020820190508181036000830152610f6481610f28565b905091905056fe225d7d2c22696e70757473223a5b7b2253746f72616765536f75726365223a2255524c446f776e6c6f6164222c2255524c223a2268747470733a2f2f6170692e6472616e642e73682f383939306537613961616564326666656437336462643730393231323364366632383939333035343064373635313333363232356463313732653531623263652f7075626c69632f6c6174657374222c2270617468223a222f696e70757473227d5d2c224465616c223a7b22436f6e63757272656e6379223a317d7d7b22456e67696e65223a22446f636b6572222c225665726966696572223a224e6f6f70222c225075626c6973686572223a2245737475617279222c22446f636b6572223a7b22496d616765223a22676863722e696f2f626163616c6861752d70726f6a6563742f6472616e642d6578616d706c653a6c6174657374407368613235363a61306330353030363935313236386532336462343638623032383735643961343338386239623063313036373738303162376230376131636534646433653830222c22456e747279706f696e74223a5b222f62696e2f72616e64222c22a26469706673582212202ce3907f5fbf894f1d3c4f91a21c69f43708592b8bb94c7a80bbb70242d2db4c64736f6c63430008110033",
}

// DrandCallerABI is the input ABI used to generate the binding from.
// Deprecated: Use DrandCallerMetaData.ABI instead.
var DrandCallerABI = DrandCallerMetaData.ABI

// DrandCallerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DrandCallerMetaData.Bin instead.
var DrandCallerBin = DrandCallerMetaData.Bin

// DeployDrandCaller deploys a new Ethereum contract, binding an instance of DrandCaller to it.
func DeployDrandCaller(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeContract common.Address) (common.Address, *types.Transaction, *DrandCaller, error) {
	parsed, err := DrandCallerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DrandCallerBin), backend, bridgeContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DrandCaller{DrandCallerCaller: DrandCallerCaller{contract: contract}, DrandCallerTransactor: DrandCallerTransactor{contract: contract}, DrandCallerFilterer: DrandCallerFilterer{contract: contract}}, nil
}

// DrandCaller is an auto generated Go binding around an Ethereum contract.
type DrandCaller struct {
	DrandCallerCaller     // Read-only binding to the contract
	DrandCallerTransactor // Write-only binding to the contract
	DrandCallerFilterer   // Log filterer for contract events
}

// DrandCallerCaller is an auto generated read-only Go binding around an Ethereum contract.
type DrandCallerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DrandCallerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DrandCallerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DrandCallerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DrandCallerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DrandCallerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DrandCallerSession struct {
	Contract     *DrandCaller      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DrandCallerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DrandCallerCallerSession struct {
	Contract *DrandCallerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DrandCallerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DrandCallerTransactorSession struct {
	Contract     *DrandCallerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DrandCallerRaw is an auto generated low-level Go binding around an Ethereum contract.
type DrandCallerRaw struct {
	Contract *DrandCaller // Generic contract binding to access the raw methods on
}

// DrandCallerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DrandCallerCallerRaw struct {
	Contract *DrandCallerCaller // Generic read-only contract binding to access the raw methods on
}

// DrandCallerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DrandCallerTransactorRaw struct {
	Contract *DrandCallerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDrandCaller creates a new instance of DrandCaller, bound to a specific deployed contract.
func NewDrandCaller(address common.Address, backend bind.ContractBackend) (*DrandCaller, error) {
	contract, err := bindDrandCaller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DrandCaller{DrandCallerCaller: DrandCallerCaller{contract: contract}, DrandCallerTransactor: DrandCallerTransactor{contract: contract}, DrandCallerFilterer: DrandCallerFilterer{contract: contract}}, nil
}

// NewDrandCallerCaller creates a new read-only instance of DrandCaller, bound to a specific deployed contract.
func NewDrandCallerCaller(address common.Address, caller bind.ContractCaller) (*DrandCallerCaller, error) {
	contract, err := bindDrandCaller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DrandCallerCaller{contract: contract}, nil
}

// NewDrandCallerTransactor creates a new write-only instance of DrandCaller, bound to a specific deployed contract.
func NewDrandCallerTransactor(address common.Address, transactor bind.ContractTransactor) (*DrandCallerTransactor, error) {
	contract, err := bindDrandCaller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DrandCallerTransactor{contract: contract}, nil
}

// NewDrandCallerFilterer creates a new log filterer instance of DrandCaller, bound to a specific deployed contract.
func NewDrandCallerFilterer(address common.Address, filterer bind.ContractFilterer) (*DrandCallerFilterer, error) {
	contract, err := bindDrandCaller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DrandCallerFilterer{contract: contract}, nil
}

// bindDrandCaller binds a generic wrapper to an already deployed contract.
func bindDrandCaller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DrandCallerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DrandCaller *DrandCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DrandCaller.Contract.DrandCallerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DrandCaller *DrandCallerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DrandCaller.Contract.DrandCallerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DrandCaller *DrandCallerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DrandCaller.Contract.DrandCallerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DrandCaller *DrandCallerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DrandCaller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DrandCaller *DrandCallerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DrandCaller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DrandCaller *DrandCallerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DrandCaller.Contract.contract.Transact(opts, method, params...)
}

// GetRandomNumber is a paid mutator transaction binding the contract method 0x6e68fc0a.
//
// Solidity: function getRandomNumber(uint256 minimum, uint256 maximum) returns()
func (_DrandCaller *DrandCallerTransactor) GetRandomNumber(opts *bind.TransactOpts, minimum *big.Int, maximum *big.Int) (*types.Transaction, error) {
	return _DrandCaller.contract.Transact(opts, "getRandomNumber", minimum, maximum)
}

// GetRandomNumber is a paid mutator transaction binding the contract method 0x6e68fc0a.
//
// Solidity: function getRandomNumber(uint256 minimum, uint256 maximum) returns()
func (_DrandCaller *DrandCallerSession) GetRandomNumber(minimum *big.Int, maximum *big.Int) (*types.Transaction, error) {
	return _DrandCaller.Contract.GetRandomNumber(&_DrandCaller.TransactOpts, minimum, maximum)
}

// GetRandomNumber is a paid mutator transaction binding the contract method 0x6e68fc0a.
//
// Solidity: function getRandomNumber(uint256 minimum, uint256 maximum) returns()
func (_DrandCaller *DrandCallerTransactorSession) GetRandomNumber(minimum *big.Int, maximum *big.Int) (*types.Transaction, error) {
	return _DrandCaller.Contract.GetRandomNumber(&_DrandCaller.TransactOpts, minimum, maximum)
}

// LilypadReceiver is a paid mutator transaction binding the contract method 0x511deb47.
//
// Solidity: function lilypadReceiver(address _from, uint256 _jobId, uint8 _resultType, string _result) returns()
func (_DrandCaller *DrandCallerTransactor) LilypadReceiver(opts *bind.TransactOpts, _from common.Address, _jobId *big.Int, _resultType uint8, _result string) (*types.Transaction, error) {
	return _DrandCaller.contract.Transact(opts, "lilypadReceiver", _from, _jobId, _resultType, _result)
}

// LilypadReceiver is a paid mutator transaction binding the contract method 0x511deb47.
//
// Solidity: function lilypadReceiver(address _from, uint256 _jobId, uint8 _resultType, string _result) returns()
func (_DrandCaller *DrandCallerSession) LilypadReceiver(_from common.Address, _jobId *big.Int, _resultType uint8, _result string) (*types.Transaction, error) {
	return _DrandCaller.Contract.LilypadReceiver(&_DrandCaller.TransactOpts, _from, _jobId, _resultType, _result)
}

// LilypadReceiver is a paid mutator transaction binding the contract method 0x511deb47.
//
// Solidity: function lilypadReceiver(address _from, uint256 _jobId, uint8 _resultType, string _result) returns()
func (_DrandCaller *DrandCallerTransactorSession) LilypadReceiver(_from common.Address, _jobId *big.Int, _resultType uint8, _result string) (*types.Transaction, error) {
	return _DrandCaller.Contract.LilypadReceiver(&_DrandCaller.TransactOpts, _from, _jobId, _resultType, _result)
}

// DrandCallerRandomNumberGeneratedIterator is returned from FilterRandomNumberGenerated and is used to iterate over the raw logs and unpacked data for RandomNumberGenerated events raised by the DrandCaller contract.
type DrandCallerRandomNumberGeneratedIterator struct {
	Event *DrandCallerRandomNumberGenerated // Event containing the contract specifics and raw log

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
func (it *DrandCallerRandomNumberGeneratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DrandCallerRandomNumberGenerated)
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
		it.Event = new(DrandCallerRandomNumberGenerated)
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
func (it *DrandCallerRandomNumberGeneratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DrandCallerRandomNumberGeneratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DrandCallerRandomNumberGenerated represents a RandomNumberGenerated event raised by the DrandCaller contract.
type DrandCallerRandomNumberGenerated struct {
	Id     *big.Int
	Number *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRandomNumberGenerated is a free log retrieval operation binding the contract event 0x252851c72bba973a3eff4c55d976631b74c5e463b736ab3ba7ea37cbe40bbc19.
//
// Solidity: event RandomNumberGenerated(uint256 id, uint256 number)
func (_DrandCaller *DrandCallerFilterer) FilterRandomNumberGenerated(opts *bind.FilterOpts) (*DrandCallerRandomNumberGeneratedIterator, error) {

	logs, sub, err := _DrandCaller.contract.FilterLogs(opts, "RandomNumberGenerated")
	if err != nil {
		return nil, err
	}
	return &DrandCallerRandomNumberGeneratedIterator{contract: _DrandCaller.contract, event: "RandomNumberGenerated", logs: logs, sub: sub}, nil
}

// WatchRandomNumberGenerated is a free log subscription operation binding the contract event 0x252851c72bba973a3eff4c55d976631b74c5e463b736ab3ba7ea37cbe40bbc19.
//
// Solidity: event RandomNumberGenerated(uint256 id, uint256 number)
func (_DrandCaller *DrandCallerFilterer) WatchRandomNumberGenerated(opts *bind.WatchOpts, sink chan<- *DrandCallerRandomNumberGenerated) (event.Subscription, error) {

	logs, sub, err := _DrandCaller.contract.WatchLogs(opts, "RandomNumberGenerated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DrandCallerRandomNumberGenerated)
				if err := _DrandCaller.contract.UnpackLog(event, "RandomNumberGenerated", log); err != nil {
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

// ParseRandomNumberGenerated is a log parse operation binding the contract event 0x252851c72bba973a3eff4c55d976631b74c5e463b736ab3ba7ea37cbe40bbc19.
//
// Solidity: event RandomNumberGenerated(uint256 id, uint256 number)
func (_DrandCaller *DrandCallerFilterer) ParseRandomNumberGenerated(log types.Log) (*DrandCallerRandomNumberGenerated, error) {
	event := new(DrandCallerRandomNumberGenerated)
	if err := _DrandCaller.contract.UnpackLog(event, "RandomNumberGenerated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}


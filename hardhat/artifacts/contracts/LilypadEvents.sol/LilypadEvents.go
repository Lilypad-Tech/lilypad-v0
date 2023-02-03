// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package LilypadEvents

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

// LilypadEventsBacalhauJob is an auto generated low-level Go binding around an user-defined struct.
type LilypadEventsBacalhauJob struct {
	Requestor  common.Address
	JobId      *big.Int
	JobName    string
	IPFSresult string
}

// LilypadEventsMetaData contains all meta data concerning the LilypadEvents contract.
var LilypadEventsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requestorContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"jobName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"results\",\"type\":\"string\"}],\"name\":\"BacalhauJobResultsReturned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requestorContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"jobName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"params\",\"type\":\"string\"}],\"name\":\"NewBacalhauJobSubmitted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bacalhauJobHistory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jobName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"IPFSresult\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchAllJobs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jobName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"IPFSresult\",\"type\":\"string\"}],\"internalType\":\"structLilypadEvents.BacalhauJob[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_requestor\",\"type\":\"address\"}],\"name\":\"fetchJobsByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jobName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"IPFSresult\",\"type\":\"string\"}],\"internalType\":\"structLilypadEvents.BacalhauJob[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_jobId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_jobName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ipfsResult\",\"type\":\"string\"}],\"name\":\"returnBacalhauResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_jobName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_params\",\"type\":\"string\"}],\"name\":\"runBacalhauJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506200005d6040518060400160405280602081526020017f4465706c6f79696e67204c696c797061644576656e747320636f6e74726163748152506200006360201b6200094f1760201c565b620001ed565b62000103816040516024016200007a9190620001c9565b6040516020818303038152906040527f41304fac000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506200010660201b60201c565b50565b60008151905060006a636f6e736f6c652e6c6f679050602083016000808483855afa5050505050565b600081519050919050565b600082825260208201905092915050565b60005b838110156200016b5780820151818401526020810190506200014e565b60008484015250505050565b6000601f19601f8301169050919050565b600062000195826200012f565b620001a181856200013a565b9350620001b38185602086016200014b565b620001be8162000177565b840191505092915050565b60006020820190508181036000830152620001e5818462000188565b905092915050565b61151080620001fd6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063399fad371461005c57806372f8c7d51461008c578063807d6dcf146100a8578063ce4af1d7146100db578063fe41ff3b146100f9575b600080fd5b61007660048036038101906100719190610b40565b610115565b6040516100839190610d51565b60405180910390f35b6100a660048036038101906100a19190610ed4565b610337565b005b6100c260048036038101906100bd9190610f73565b610587565b6040516100d29493929190611008565b60405180910390f35b6100e36106f7565b6040516100f09190610d51565b60405180910390f35b610113600480360381019061010e919061105b565b6108da565b005b6060600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b8282101561032c57838290600052602060002090600402016040518060800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820154815260200160028201805461020990611115565b80601f016020809104026020016040519081016040528092919081815260200182805461023590611115565b80156102825780601f1061025757610100808354040283529160200191610282565b820191906000526020600020905b81548152906001019060200180831161026557829003601f168201915b5050505050815260200160038201805461029b90611115565b80601f01602080910402602001604051908101604052809291908181526020018280546102c790611115565b80156103145780601f106102e957610100808354040283529160200191610314565b820191906000526020600020905b8154815290600101906020018083116102f757829003601f168201915b50505050508152505081526020019060010190610176565b505050509050919050565b6103756040518060400160405280600881526020017f646f20737475666600000000000000000000000000000000000000000000000081525061094f565b600060405180608001604052808673ffffffffffffffffffffffffffffffffffffffff168152602001858152602001848152602001838152509050600181908060018154018082558091505060019003906000526020600020906004020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201908161044291906112f2565b50606082015181600301908161045891906112f2565b505050600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081908060018154018082558091505060019003906000526020600020906004020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201908161052a91906112f2565b50606082015181600301908161054091906112f2565b5050507fbd272723382793172868359a6bd8728230f58381cdbc51499a89c6e63ba0a07f858585856040516105789493929190611008565b60405180910390a15050505050565b6001818154811061059757600080fd5b90600052602060002090600402016000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020180546105e690611115565b80601f016020809104026020016040519081016040528092919081815260200182805461061290611115565b801561065f5780601f106106345761010080835404028352916020019161065f565b820191906000526020600020905b81548152906001019060200180831161064257829003601f168201915b50505050509080600301805461067490611115565b80601f01602080910402602001604051908101604052809291908181526020018280546106a090611115565b80156106ed5780601f106106c2576101008083540402835291602001916106ed565b820191906000526020600020905b8154815290600101906020018083116106d057829003601f168201915b5050505050905084565b60606001805480602002602001604051908101604052809291908181526020016000905b828210156108d157838290600052602060002090600402016040518060800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820180546107ae90611115565b80601f01602080910402602001604051908101604052809291908181526020018280546107da90611115565b80156108275780601f106107fc57610100808354040283529160200191610827565b820191906000526020600020905b81548152906001019060200180831161080a57829003601f168201915b5050505050815260200160038201805461084090611115565b80601f016020809104026020016040519081016040528092919081815260200182805461086c90611115565b80156108b95780601f1061088e576101008083540402835291602001916108b9565b820191906000526020600020905b81548152906001019060200180831161089c57829003601f168201915b5050505050815250508152602001906001019061071b565b50505050905090565b6108e3336109e8565b60006108ef6000610a81565b90503373ffffffffffffffffffffffffffffffffffffffff167f46e1125364fc85bb83a2688b3a3a653a703b28f056096ae099d104e66ab6ab0982604051610937919061145c565b60405180910390a26109496000610a8f565b50505050565b6109e581604051602401610963919061149d565b6040516020818303038152906040527f41304fac000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610aa5565b50565b610a7e816040516024016109fc91906114bf565b6040516020818303038152906040527f2c2ecbc2000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610aa5565b50565b600081600001549050919050565b6001816000016000828254019250508190555050565b60008151905060006a636f6e736f6c652e6c6f679050602083016000808483855afa5050505050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610b0d82610ae2565b9050919050565b610b1d81610b02565b8114610b2857600080fd5b50565b600081359050610b3a81610b14565b92915050565b600060208284031215610b5657610b55610ad8565b5b6000610b6484828501610b2b565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610ba281610b02565b82525050565b6000819050919050565b610bbb81610ba8565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610bfb578082015181840152602081019050610be0565b60008484015250505050565b6000601f19601f8301169050919050565b6000610c2382610bc1565b610c2d8185610bcc565b9350610c3d818560208601610bdd565b610c4681610c07565b840191505092915050565b6000608083016000830151610c696000860182610b99565b506020830151610c7c6020860182610bb2565b5060408301518482036040860152610c948282610c18565b91505060608301518482036060860152610cae8282610c18565b9150508091505092915050565b6000610cc78383610c51565b905092915050565b6000602082019050919050565b6000610ce782610b6d565b610cf18185610b78565b935083602082028501610d0385610b89565b8060005b85811015610d3f5784840389528151610d208582610cbb565b9450610d2b83610ccf565b925060208a01995050600181019050610d07565b50829750879550505050505092915050565b60006020820190508181036000830152610d6b8184610cdc565b905092915050565b610d7c81610ba8565b8114610d8757600080fd5b50565b600081359050610d9981610d73565b92915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610de182610c07565b810181811067ffffffffffffffff82111715610e0057610dff610da9565b5b80604052505050565b6000610e13610ace565b9050610e1f8282610dd8565b919050565b600067ffffffffffffffff821115610e3f57610e3e610da9565b5b610e4882610c07565b9050602081019050919050565b82818337600083830152505050565b6000610e77610e7284610e24565b610e09565b905082815260208101848484011115610e9357610e92610da4565b5b610e9e848285610e55565b509392505050565b600082601f830112610ebb57610eba610d9f565b5b8135610ecb848260208601610e64565b91505092915050565b60008060008060808587031215610eee57610eed610ad8565b5b6000610efc87828801610b2b565b9450506020610f0d87828801610d8a565b935050604085013567ffffffffffffffff811115610f2e57610f2d610add565b5b610f3a87828801610ea6565b925050606085013567ffffffffffffffff811115610f5b57610f5a610add565b5b610f6787828801610ea6565b91505092959194509250565b600060208284031215610f8957610f88610ad8565b5b6000610f9784828501610d8a565b91505092915050565b610fa981610b02565b82525050565b610fb881610ba8565b82525050565b600082825260208201905092915050565b6000610fda82610bc1565b610fe48185610fbe565b9350610ff4818560208601610bdd565b610ffd81610c07565b840191505092915050565b600060808201905061101d6000830187610fa0565b61102a6020830186610faf565b818103604083015261103c8185610fcf565b905081810360608301526110508184610fcf565b905095945050505050565b60008060006060848603121561107457611073610ad8565b5b600061108286828701610b2b565b935050602084013567ffffffffffffffff8111156110a3576110a2610add565b5b6110af86828701610ea6565b925050604084013567ffffffffffffffff8111156110d0576110cf610add565b5b6110dc86828701610ea6565b9150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061112d57607f821691505b6020821081036111405761113f6110e6565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026111a87fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261116b565b6111b2868361116b565b95508019841693508086168417925050509392505050565b6000819050919050565b60006111ef6111ea6111e584610ba8565b6111ca565b610ba8565b9050919050565b6000819050919050565b611209836111d4565b61121d611215826111f6565b848454611178565b825550505050565b600090565b611232611225565b61123d818484611200565b505050565b5b818110156112615761125660008261122a565b600181019050611243565b5050565b601f8211156112a65761127781611146565b6112808461115b565b8101602085101561128f578190505b6112a361129b8561115b565b830182611242565b50505b505050565b600082821c905092915050565b60006112c9600019846008026112ab565b1980831691505092915050565b60006112e283836112b8565b9150826002028217905092915050565b6112fb82610bc1565b67ffffffffffffffff81111561131457611313610da9565b5b61131e8254611115565b611329828285611265565b600060209050601f83116001811461135c576000841561134a578287015190505b61135485826112d6565b8655506113bc565b601f19841661136a86611146565b60005b828110156113925784890151825560018201915060208501945060208101905061136d565b868310156113af57848901516113ab601f8916826112b8565b8355505b6001600288020188555050505b505050505050565b7f537461626c65446966667573696f6e4750550000000000000000000000000000600082015250565b60006113fa601283610fbe565b9150611405826113c4565b602082019050919050565b7f7b2270726f6d7074223a225261696e626f77556e69636f726e7d000000000000600082015250565b6000611446601a83610fbe565b915061145182611410565b602082019050919050565b60006060820190506114716000830184610faf565b8181036020830152611482816113ed565b9050818103604083015261149581611439565b905092915050565b600060208201905081810360008301526114b78184610fcf565b905092915050565b60006020820190506114d46000830184610fa0565b9291505056fea26469706673582212209d2cbfc3e1a675ab545a2043ceab837b650b35a630b7fbf17c395c7512f4411964736f6c63430008110033",
}

// LilypadEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use LilypadEventsMetaData.ABI instead.
var LilypadEventsABI = LilypadEventsMetaData.ABI

// LilypadEventsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LilypadEventsMetaData.Bin instead.
var LilypadEventsBin = LilypadEventsMetaData.Bin

// DeployLilypadEvents deploys a new Ethereum contract, binding an instance of LilypadEvents to it.
func DeployLilypadEvents(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LilypadEvents, error) {
	parsed, err := LilypadEventsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LilypadEventsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LilypadEvents{LilypadEventsCaller: LilypadEventsCaller{contract: contract}, LilypadEventsTransactor: LilypadEventsTransactor{contract: contract}, LilypadEventsFilterer: LilypadEventsFilterer{contract: contract}}, nil
}

// LilypadEvents is an auto generated Go binding around an Ethereum contract.
type LilypadEvents struct {
	LilypadEventsCaller     // Read-only binding to the contract
	LilypadEventsTransactor // Write-only binding to the contract
	LilypadEventsFilterer   // Log filterer for contract events
}

// LilypadEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type LilypadEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LilypadEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LilypadEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LilypadEventsSession struct {
	Contract     *LilypadEvents    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LilypadEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LilypadEventsCallerSession struct {
	Contract *LilypadEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// LilypadEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LilypadEventsTransactorSession struct {
	Contract     *LilypadEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// LilypadEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type LilypadEventsRaw struct {
	Contract *LilypadEvents // Generic contract binding to access the raw methods on
}

// LilypadEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LilypadEventsCallerRaw struct {
	Contract *LilypadEventsCaller // Generic read-only contract binding to access the raw methods on
}

// LilypadEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LilypadEventsTransactorRaw struct {
	Contract *LilypadEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLilypadEvents creates a new instance of LilypadEvents, bound to a specific deployed contract.
func NewLilypadEvents(address common.Address, backend bind.ContractBackend) (*LilypadEvents, error) {
	contract, err := bindLilypadEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LilypadEvents{LilypadEventsCaller: LilypadEventsCaller{contract: contract}, LilypadEventsTransactor: LilypadEventsTransactor{contract: contract}, LilypadEventsFilterer: LilypadEventsFilterer{contract: contract}}, nil
}

// NewLilypadEventsCaller creates a new read-only instance of LilypadEvents, bound to a specific deployed contract.
func NewLilypadEventsCaller(address common.Address, caller bind.ContractCaller) (*LilypadEventsCaller, error) {
	contract, err := bindLilypadEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadEventsCaller{contract: contract}, nil
}

// NewLilypadEventsTransactor creates a new write-only instance of LilypadEvents, bound to a specific deployed contract.
func NewLilypadEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*LilypadEventsTransactor, error) {
	contract, err := bindLilypadEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadEventsTransactor{contract: contract}, nil
}

// NewLilypadEventsFilterer creates a new log filterer instance of LilypadEvents, bound to a specific deployed contract.
func NewLilypadEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*LilypadEventsFilterer, error) {
	contract, err := bindLilypadEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LilypadEventsFilterer{contract: contract}, nil
}

// bindLilypadEvents binds a generic wrapper to an already deployed contract.
func bindLilypadEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LilypadEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadEvents *LilypadEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadEvents.Contract.LilypadEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadEvents *LilypadEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadEvents.Contract.LilypadEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadEvents *LilypadEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadEvents.Contract.LilypadEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadEvents *LilypadEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadEvents *LilypadEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadEvents *LilypadEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadEvents.Contract.contract.Transact(opts, method, params...)
}

// BacalhauJobHistory is a free data retrieval call binding the contract method 0x807d6dcf.
//
// Solidity: function bacalhauJobHistory(uint256 ) view returns(address requestor, uint256 jobId, string jobName, string IPFSresult)
func (_LilypadEvents *LilypadEventsCaller) BacalhauJobHistory(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Requestor  common.Address
	JobId      *big.Int
	JobName    string
	IPFSresult string
}, error) {
	var out []interface{}
	err := _LilypadEvents.contract.Call(opts, &out, "bacalhauJobHistory", arg0)

	outstruct := new(struct {
		Requestor  common.Address
		JobId      *big.Int
		JobName    string
		IPFSresult string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Requestor = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.JobId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.JobName = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.IPFSresult = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// BacalhauJobHistory is a free data retrieval call binding the contract method 0x807d6dcf.
//
// Solidity: function bacalhauJobHistory(uint256 ) view returns(address requestor, uint256 jobId, string jobName, string IPFSresult)
func (_LilypadEvents *LilypadEventsSession) BacalhauJobHistory(arg0 *big.Int) (struct {
	Requestor  common.Address
	JobId      *big.Int
	JobName    string
	IPFSresult string
}, error) {
	return _LilypadEvents.Contract.BacalhauJobHistory(&_LilypadEvents.CallOpts, arg0)
}

// BacalhauJobHistory is a free data retrieval call binding the contract method 0x807d6dcf.
//
// Solidity: function bacalhauJobHistory(uint256 ) view returns(address requestor, uint256 jobId, string jobName, string IPFSresult)
func (_LilypadEvents *LilypadEventsCallerSession) BacalhauJobHistory(arg0 *big.Int) (struct {
	Requestor  common.Address
	JobId      *big.Int
	JobName    string
	IPFSresult string
}, error) {
	return _LilypadEvents.Contract.BacalhauJobHistory(&_LilypadEvents.CallOpts, arg0)
}

// FetchAllJobs is a free data retrieval call binding the contract method 0xce4af1d7.
//
// Solidity: function fetchAllJobs() view returns((address,uint256,string,string)[])
func (_LilypadEvents *LilypadEventsCaller) FetchAllJobs(opts *bind.CallOpts) ([]LilypadEventsBacalhauJob, error) {
	var out []interface{}
	err := _LilypadEvents.contract.Call(opts, &out, "fetchAllJobs")

	if err != nil {
		return *new([]LilypadEventsBacalhauJob), err
	}

	out0 := *abi.ConvertType(out[0], new([]LilypadEventsBacalhauJob)).(*[]LilypadEventsBacalhauJob)

	return out0, err

}

// FetchAllJobs is a free data retrieval call binding the contract method 0xce4af1d7.
//
// Solidity: function fetchAllJobs() view returns((address,uint256,string,string)[])
func (_LilypadEvents *LilypadEventsSession) FetchAllJobs() ([]LilypadEventsBacalhauJob, error) {
	return _LilypadEvents.Contract.FetchAllJobs(&_LilypadEvents.CallOpts)
}

// FetchAllJobs is a free data retrieval call binding the contract method 0xce4af1d7.
//
// Solidity: function fetchAllJobs() view returns((address,uint256,string,string)[])
func (_LilypadEvents *LilypadEventsCallerSession) FetchAllJobs() ([]LilypadEventsBacalhauJob, error) {
	return _LilypadEvents.Contract.FetchAllJobs(&_LilypadEvents.CallOpts)
}

// FetchJobsByAddress is a free data retrieval call binding the contract method 0x399fad37.
//
// Solidity: function fetchJobsByAddress(address _requestor) view returns((address,uint256,string,string)[])
func (_LilypadEvents *LilypadEventsCaller) FetchJobsByAddress(opts *bind.CallOpts, _requestor common.Address) ([]LilypadEventsBacalhauJob, error) {
	var out []interface{}
	err := _LilypadEvents.contract.Call(opts, &out, "fetchJobsByAddress", _requestor)

	if err != nil {
		return *new([]LilypadEventsBacalhauJob), err
	}

	out0 := *abi.ConvertType(out[0], new([]LilypadEventsBacalhauJob)).(*[]LilypadEventsBacalhauJob)

	return out0, err

}

// FetchJobsByAddress is a free data retrieval call binding the contract method 0x399fad37.
//
// Solidity: function fetchJobsByAddress(address _requestor) view returns((address,uint256,string,string)[])
func (_LilypadEvents *LilypadEventsSession) FetchJobsByAddress(_requestor common.Address) ([]LilypadEventsBacalhauJob, error) {
	return _LilypadEvents.Contract.FetchJobsByAddress(&_LilypadEvents.CallOpts, _requestor)
}

// FetchJobsByAddress is a free data retrieval call binding the contract method 0x399fad37.
//
// Solidity: function fetchJobsByAddress(address _requestor) view returns((address,uint256,string,string)[])
func (_LilypadEvents *LilypadEventsCallerSession) FetchJobsByAddress(_requestor common.Address) ([]LilypadEventsBacalhauJob, error) {
	return _LilypadEvents.Contract.FetchJobsByAddress(&_LilypadEvents.CallOpts, _requestor)
}

// ReturnBacalhauResults is a paid mutator transaction binding the contract method 0x72f8c7d5.
//
// Solidity: function returnBacalhauResults(address _to, uint256 _jobId, string _jobName, string _ipfsResult) returns()
func (_LilypadEvents *LilypadEventsTransactor) ReturnBacalhauResults(opts *bind.TransactOpts, _to common.Address, _jobId *big.Int, _jobName string, _ipfsResult string) (*types.Transaction, error) {
	return _LilypadEvents.contract.Transact(opts, "returnBacalhauResults", _to, _jobId, _jobName, _ipfsResult)
}

// ReturnBacalhauResults is a paid mutator transaction binding the contract method 0x72f8c7d5.
//
// Solidity: function returnBacalhauResults(address _to, uint256 _jobId, string _jobName, string _ipfsResult) returns()
func (_LilypadEvents *LilypadEventsSession) ReturnBacalhauResults(_to common.Address, _jobId *big.Int, _jobName string, _ipfsResult string) (*types.Transaction, error) {
	return _LilypadEvents.Contract.ReturnBacalhauResults(&_LilypadEvents.TransactOpts, _to, _jobId, _jobName, _ipfsResult)
}

// ReturnBacalhauResults is a paid mutator transaction binding the contract method 0x72f8c7d5.
//
// Solidity: function returnBacalhauResults(address _to, uint256 _jobId, string _jobName, string _ipfsResult) returns()
func (_LilypadEvents *LilypadEventsTransactorSession) ReturnBacalhauResults(_to common.Address, _jobId *big.Int, _jobName string, _ipfsResult string) (*types.Transaction, error) {
	return _LilypadEvents.Contract.ReturnBacalhauResults(&_LilypadEvents.TransactOpts, _to, _jobId, _jobName, _ipfsResult)
}

// RunBacalhauJob is a paid mutator transaction binding the contract method 0xfe41ff3b.
//
// Solidity: function runBacalhauJob(address _from, string _jobName, string _params) returns()
func (_LilypadEvents *LilypadEventsTransactor) RunBacalhauJob(opts *bind.TransactOpts, _from common.Address, _jobName string, _params string) (*types.Transaction, error) {
	return _LilypadEvents.contract.Transact(opts, "runBacalhauJob", _from, _jobName, _params)
}

// RunBacalhauJob is a paid mutator transaction binding the contract method 0xfe41ff3b.
//
// Solidity: function runBacalhauJob(address _from, string _jobName, string _params) returns()
func (_LilypadEvents *LilypadEventsSession) RunBacalhauJob(_from common.Address, _jobName string, _params string) (*types.Transaction, error) {
	return _LilypadEvents.Contract.RunBacalhauJob(&_LilypadEvents.TransactOpts, _from, _jobName, _params)
}

// RunBacalhauJob is a paid mutator transaction binding the contract method 0xfe41ff3b.
//
// Solidity: function runBacalhauJob(address _from, string _jobName, string _params) returns()
func (_LilypadEvents *LilypadEventsTransactorSession) RunBacalhauJob(_from common.Address, _jobName string, _params string) (*types.Transaction, error) {
	return _LilypadEvents.Contract.RunBacalhauJob(&_LilypadEvents.TransactOpts, _from, _jobName, _params)
}

// LilypadEventsBacalhauJobResultsReturnedIterator is returned from FilterBacalhauJobResultsReturned and is used to iterate over the raw logs and unpacked data for BacalhauJobResultsReturned events raised by the LilypadEvents contract.
type LilypadEventsBacalhauJobResultsReturnedIterator struct {
	Event *LilypadEventsBacalhauJobResultsReturned // Event containing the contract specifics and raw log

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
func (it *LilypadEventsBacalhauJobResultsReturnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadEventsBacalhauJobResultsReturned)
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
		it.Event = new(LilypadEventsBacalhauJobResultsReturned)
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
func (it *LilypadEventsBacalhauJobResultsReturnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadEventsBacalhauJobResultsReturnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadEventsBacalhauJobResultsReturned represents a BacalhauJobResultsReturned event raised by the LilypadEvents contract.
type LilypadEventsBacalhauJobResultsReturned struct {
	RequestorContract common.Address
	JobId             *big.Int
	JobName           string
	Results           string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBacalhauJobResultsReturned is a free log retrieval operation binding the contract event 0xbd272723382793172868359a6bd8728230f58381cdbc51499a89c6e63ba0a07f.
//
// Solidity: event BacalhauJobResultsReturned(address requestorContract, uint256 jobId, string jobName, string results)
func (_LilypadEvents *LilypadEventsFilterer) FilterBacalhauJobResultsReturned(opts *bind.FilterOpts) (*LilypadEventsBacalhauJobResultsReturnedIterator, error) {

	logs, sub, err := _LilypadEvents.contract.FilterLogs(opts, "BacalhauJobResultsReturned")
	if err != nil {
		return nil, err
	}
	return &LilypadEventsBacalhauJobResultsReturnedIterator{contract: _LilypadEvents.contract, event: "BacalhauJobResultsReturned", logs: logs, sub: sub}, nil
}

// WatchBacalhauJobResultsReturned is a free log subscription operation binding the contract event 0xbd272723382793172868359a6bd8728230f58381cdbc51499a89c6e63ba0a07f.
//
// Solidity: event BacalhauJobResultsReturned(address requestorContract, uint256 jobId, string jobName, string results)
func (_LilypadEvents *LilypadEventsFilterer) WatchBacalhauJobResultsReturned(opts *bind.WatchOpts, sink chan<- *LilypadEventsBacalhauJobResultsReturned) (event.Subscription, error) {

	logs, sub, err := _LilypadEvents.contract.WatchLogs(opts, "BacalhauJobResultsReturned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadEventsBacalhauJobResultsReturned)
				if err := _LilypadEvents.contract.UnpackLog(event, "BacalhauJobResultsReturned", log); err != nil {
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

// ParseBacalhauJobResultsReturned is a log parse operation binding the contract event 0xbd272723382793172868359a6bd8728230f58381cdbc51499a89c6e63ba0a07f.
//
// Solidity: event BacalhauJobResultsReturned(address requestorContract, uint256 jobId, string jobName, string results)
func (_LilypadEvents *LilypadEventsFilterer) ParseBacalhauJobResultsReturned(log types.Log) (*LilypadEventsBacalhauJobResultsReturned, error) {
	event := new(LilypadEventsBacalhauJobResultsReturned)
	if err := _LilypadEvents.contract.UnpackLog(event, "BacalhauJobResultsReturned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadEventsNewBacalhauJobSubmittedIterator is returned from FilterNewBacalhauJobSubmitted and is used to iterate over the raw logs and unpacked data for NewBacalhauJobSubmitted events raised by the LilypadEvents contract.
type LilypadEventsNewBacalhauJobSubmittedIterator struct {
	Event *LilypadEventsNewBacalhauJobSubmitted // Event containing the contract specifics and raw log

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
func (it *LilypadEventsNewBacalhauJobSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadEventsNewBacalhauJobSubmitted)
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
		it.Event = new(LilypadEventsNewBacalhauJobSubmitted)
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
func (it *LilypadEventsNewBacalhauJobSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadEventsNewBacalhauJobSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadEventsNewBacalhauJobSubmitted represents a NewBacalhauJobSubmitted event raised by the LilypadEvents contract.
type LilypadEventsNewBacalhauJobSubmitted struct {
	RequestorContract common.Address
	JobId             *big.Int
	JobName           string
	Params            string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewBacalhauJobSubmitted is a free log retrieval operation binding the contract event 0x46e1125364fc85bb83a2688b3a3a653a703b28f056096ae099d104e66ab6ab09.
//
// Solidity: event NewBacalhauJobSubmitted(address indexed requestorContract, uint256 jobId, string jobName, string params)
func (_LilypadEvents *LilypadEventsFilterer) FilterNewBacalhauJobSubmitted(opts *bind.FilterOpts, requestorContract []common.Address) (*LilypadEventsNewBacalhauJobSubmittedIterator, error) {

	var requestorContractRule []interface{}
	for _, requestorContractItem := range requestorContract {
		requestorContractRule = append(requestorContractRule, requestorContractItem)
	}

	logs, sub, err := _LilypadEvents.contract.FilterLogs(opts, "NewBacalhauJobSubmitted", requestorContractRule)
	if err != nil {
		return nil, err
	}
	return &LilypadEventsNewBacalhauJobSubmittedIterator{contract: _LilypadEvents.contract, event: "NewBacalhauJobSubmitted", logs: logs, sub: sub}, nil
}

// WatchNewBacalhauJobSubmitted is a free log subscription operation binding the contract event 0x46e1125364fc85bb83a2688b3a3a653a703b28f056096ae099d104e66ab6ab09.
//
// Solidity: event NewBacalhauJobSubmitted(address indexed requestorContract, uint256 jobId, string jobName, string params)
func (_LilypadEvents *LilypadEventsFilterer) WatchNewBacalhauJobSubmitted(opts *bind.WatchOpts, sink chan<- *LilypadEventsNewBacalhauJobSubmitted, requestorContract []common.Address) (event.Subscription, error) {

	var requestorContractRule []interface{}
	for _, requestorContractItem := range requestorContract {
		requestorContractRule = append(requestorContractRule, requestorContractItem)
	}

	logs, sub, err := _LilypadEvents.contract.WatchLogs(opts, "NewBacalhauJobSubmitted", requestorContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadEventsNewBacalhauJobSubmitted)
				if err := _LilypadEvents.contract.UnpackLog(event, "NewBacalhauJobSubmitted", log); err != nil {
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

// ParseNewBacalhauJobSubmitted is a log parse operation binding the contract event 0x46e1125364fc85bb83a2688b3a3a653a703b28f056096ae099d104e66ab6ab09.
//
// Solidity: event NewBacalhauJobSubmitted(address indexed requestorContract, uint256 jobId, string jobName, string params)
func (_LilypadEvents *LilypadEventsFilterer) ParseNewBacalhauJobSubmitted(log types.Log) (*LilypadEventsNewBacalhauJobSubmitted, error) {
	event := new(LilypadEventsNewBacalhauJobSubmitted)
	if err := _LilypadEvents.contract.UnpackLog(event, "NewBacalhauJobSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}


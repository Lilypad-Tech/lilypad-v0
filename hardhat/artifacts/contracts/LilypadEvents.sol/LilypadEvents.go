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
	Requestor common.Address
	Id        *big.Int
	Result    string
}

// LilypadEventsBacalhauJobCalled is an auto generated low-level Go binding around an user-defined struct.
type LilypadEventsBacalhauJobCalled struct {
	Requestor common.Address
	Id        *big.Int
	Spec      string
}

// LilypadEventsMetaData contains all meta data concerning the LilypadEvents contract.
var LilypadEventsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requestorContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"results\",\"type\":\"string\"}],\"name\":\"BacalhauJobResultsReturned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requestorContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"spec\",\"type\":\"string\"}],\"name\":\"NewBacalhauJobSubmitted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bacalhauJobCalledHistory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"spec\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bacalhauJobHistory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchAllCalledJobs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"spec\",\"type\":\"string\"}],\"internalType\":\"structLilypadEvents.BacalhauJobCalled[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchAllJobs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"internalType\":\"structLilypadEvents.BacalhauJob[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_requestor\",\"type\":\"address\"}],\"name\":\"fetchJobsByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"internalType\":\"structLilypadEvents.BacalhauJob[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_jobId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_ipfsResult\",\"type\":\"string\"}],\"name\":\"returnBacalhauResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_spec\",\"type\":\"string\"}],\"name\":\"runBacalhauJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506200005d6040518060400160405280602081526020017f4465706c6f79696e67204c696c797061644576656e747320636f6e74726163748152506200006360201b62000b0d1760201c565b620001ed565b62000103816040516024016200007a9190620001c9565b6040516020818303038152906040527f41304fac000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506200010660201b60201c565b50565b60008151905060006a636f6e736f6c652e6c6f679050602083016000808483855afa5050505050565b600081519050919050565b600082825260208201905092915050565b60005b838110156200016b5780820151818401526020810190506200014e565b60008484015250505050565b6000601f19601f8301169050919050565b600062000195826200012f565b620001a181856200013a565b9350620001b38185602086016200014b565b620001be8162000177565b840191505092915050565b60006020820190508181036000830152620001e5818462000188565b905092915050565b61161780620001fd6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063807d6dcf1161005b578063807d6dcf146100ec578063a03222461461011e578063ce4af1d71461013a578063d3b7667c146101585761007d565b8063399fad371461008257806343bd5191146100b257806346b3c952146100d0575b600080fd5b61009c60048036038101906100979190610c65565b61018a565b6040516100a99190610e5c565b60405180910390f35b6100ba61031a565b6040516100c79190610f90565b60405180910390f35b6100ea60048036038101906100e591906110e7565b61046b565b005b6101066004803603810190610101919061116f565b6105ac565b60405161011593929190611204565b60405180910390f35b61013860048036038101906101339190611242565b61068e565b005b6101426108da565b60405161014f9190610e5c565b60405180910390f35b610172600480360381019061016d919061116f565b610a2b565b60405161018193929190611204565b60405180910390f35b6060600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b8282101561030f57838290600052602060002090600302016040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820154815260200160028201805461027e906112e0565b80601f01602080910402602001604051908101604052809291908181526020018280546102aa906112e0565b80156102f75780601f106102cc576101008083540402835291602001916102f7565b820191906000526020600020905b8154815290600101906020018083116102da57829003601f168201915b505050505081525050815260200190600101906101eb565b505050509050919050565b60606002805480602002602001604051908101604052809291908181526020016000905b8282101561046257838290600052602060002090600302016040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820180546103d1906112e0565b80601f01602080910402602001604051908101604052809291908181526020018280546103fd906112e0565b801561044a5780601f1061041f5761010080835404028352916020019161044a565b820191906000526020600020905b81548152906001019060200180831161042d57829003601f168201915b5050505050815250508152602001906001019061033e565b50505050905090565b61047481610b0d565b60006104806000610ba6565b9050600060405180606001604052808573ffffffffffffffffffffffffffffffffffffffff168152602001838152602001848152509050600281908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201908161054991906114bd565b5050503373ffffffffffffffffffffffffffffffffffffffff167f824b4467ad6ba18ef48fda0d874eb770caa8f8979d165b1d2d6ab73b8fce5cb1838560405161059492919061158f565b60405180910390a26105a66000610bb4565b50505050565b600181815481106105bc57600080fd5b90600052602060002090600302016000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600101549080600201805461060b906112e0565b80601f0160208091040260200160405190810160405280929190818152602001828054610637906112e0565b80156106845780601f1061065957610100808354040283529160200191610684565b820191906000526020600020905b81548152906001019060200180831161066757829003601f168201915b5050505050905083565b600060405180606001604052808573ffffffffffffffffffffffffffffffffffffffff168152602001848152602001838152509050600181908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201908161075591906114bd565b505050600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201908161082791906114bd565b5050507f81fb9923a490992abba509e72b444bd3f44af7259412c153f25604a515f1856130848460405161085d93929190611204565b60405180910390a18373ffffffffffffffffffffffffffffffffffffffff166335c764e33085856040518463ffffffff1660e01b81526004016108a293929190611204565b600060405180830381600087803b1580156108bc57600080fd5b505af11580156108d0573d6000803e3d6000fd5b5050505050505050565b60606001805480602002602001604051908101604052809291908181526020016000905b82821015610a2257838290600052602060002090600302016040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201548152602001600282018054610991906112e0565b80601f01602080910402602001604051908101604052809291908181526020018280546109bd906112e0565b8015610a0a5780601f106109df57610100808354040283529160200191610a0a565b820191906000526020600020905b8154815290600101906020018083116109ed57829003601f168201915b505050505081525050815260200190600101906108fe565b50505050905090565b60028181548110610a3b57600080fd5b90600052602060002090600302016000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015490806002018054610a8a906112e0565b80601f0160208091040260200160405190810160405280929190818152602001828054610ab6906112e0565b8015610b035780601f10610ad857610100808354040283529160200191610b03565b820191906000526020600020905b815481529060010190602001808311610ae657829003601f168201915b5050505050905083565b610ba381604051602401610b2191906115bf565b6040516020818303038152906040527f41304fac000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610bca565b50565b600081600001549050919050565b6001816000016000828254019250508190555050565b60008151905060006a636f6e736f6c652e6c6f679050602083016000808483855afa5050505050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610c3282610c07565b9050919050565b610c4281610c27565b8114610c4d57600080fd5b50565b600081359050610c5f81610c39565b92915050565b600060208284031215610c7b57610c7a610bfd565b5b6000610c8984828501610c50565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610cc781610c27565b82525050565b6000819050919050565b610ce081610ccd565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610d20578082015181840152602081019050610d05565b60008484015250505050565b6000601f19601f8301169050919050565b6000610d4882610ce6565b610d528185610cf1565b9350610d62818560208601610d02565b610d6b81610d2c565b840191505092915050565b6000606083016000830151610d8e6000860182610cbe565b506020830151610da16020860182610cd7565b5060408301518482036040860152610db98282610d3d565b9150508091505092915050565b6000610dd28383610d76565b905092915050565b6000602082019050919050565b6000610df282610c92565b610dfc8185610c9d565b935083602082028501610e0e85610cae565b8060005b85811015610e4a5784840389528151610e2b8582610dc6565b9450610e3683610dda565b925060208a01995050600181019050610e12565b50829750879550505050505092915050565b60006020820190508181036000830152610e768184610de7565b905092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000606083016000830151610ec26000860182610cbe565b506020830151610ed56020860182610cd7565b5060408301518482036040860152610eed8282610d3d565b9150508091505092915050565b6000610f068383610eaa565b905092915050565b6000602082019050919050565b6000610f2682610e7e565b610f308185610e89565b935083602082028501610f4285610e9a565b8060005b85811015610f7e5784840389528151610f5f8582610efa565b9450610f6a83610f0e565b925060208a01995050600181019050610f46565b50829750879550505050505092915050565b60006020820190508181036000830152610faa8184610f1b565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610ff482610d2c565b810181811067ffffffffffffffff8211171561101357611012610fbc565b5b80604052505050565b6000611026610bf3565b90506110328282610feb565b919050565b600067ffffffffffffffff82111561105257611051610fbc565b5b61105b82610d2c565b9050602081019050919050565b82818337600083830152505050565b600061108a61108584611037565b61101c565b9050828152602081018484840111156110a6576110a5610fb7565b5b6110b1848285611068565b509392505050565b600082601f8301126110ce576110cd610fb2565b5b81356110de848260208601611077565b91505092915050565b600080604083850312156110fe576110fd610bfd565b5b600061110c85828601610c50565b925050602083013567ffffffffffffffff81111561112d5761112c610c02565b5b611139858286016110b9565b9150509250929050565b61114c81610ccd565b811461115757600080fd5b50565b60008135905061116981611143565b92915050565b60006020828403121561118557611184610bfd565b5b60006111938482850161115a565b91505092915050565b6111a581610c27565b82525050565b6111b481610ccd565b82525050565b600082825260208201905092915050565b60006111d682610ce6565b6111e081856111ba565b93506111f0818560208601610d02565b6111f981610d2c565b840191505092915050565b6000606082019050611219600083018661119c565b61122660208301856111ab565b818103604083015261123881846111cb565b9050949350505050565b60008060006060848603121561125b5761125a610bfd565b5b600061126986828701610c50565b935050602061127a8682870161115a565b925050604084013567ffffffffffffffff81111561129b5761129a610c02565b5b6112a7868287016110b9565b9150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806112f857607f821691505b60208210810361130b5761130a6112b1565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026113737fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611336565b61137d8683611336565b95508019841693508086168417925050509392505050565b6000819050919050565b60006113ba6113b56113b084610ccd565b611395565b610ccd565b9050919050565b6000819050919050565b6113d48361139f565b6113e86113e0826113c1565b848454611343565b825550505050565b600090565b6113fd6113f0565b6114088184846113cb565b505050565b5b8181101561142c576114216000826113f5565b60018101905061140e565b5050565b601f8211156114715761144281611311565b61144b84611326565b8101602085101561145a578190505b61146e61146685611326565b83018261140d565b50505b505050565b600082821c905092915050565b600061149460001984600802611476565b1980831691505092915050565b60006114ad8383611483565b9150826002028217905092915050565b6114c682610ce6565b67ffffffffffffffff8111156114df576114de610fbc565b5b6114e982546112e0565b6114f4828285611430565b600060209050601f8311600181146115275760008415611515578287015190505b61151f85826114a1565b865550611587565b601f19841661153586611311565b60005b8281101561155d57848901518255600182019150602085019450602081019050611538565b8683101561157a5784890151611576601f891682611483565b8355505b6001600288020188555050505b505050505050565b60006040820190506115a460008301856111ab565b81810360208301526115b681846111cb565b90509392505050565b600060208201905081810360008301526115d981846111cb565b90509291505056fea26469706673582212202b69c25d129d7019aa685ea95111bf4a9265c6e8ed8e33b824563103c81dcc5d64736f6c63430008110033",
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

// BacalhauJobCalledHistory is a free data retrieval call binding the contract method 0xd3b7667c.
//
// Solidity: function bacalhauJobCalledHistory(uint256 ) view returns(address requestor, uint256 id, string spec)
func (_LilypadEvents *LilypadEventsCaller) BacalhauJobCalledHistory(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Requestor common.Address
	Id        *big.Int
	Spec      string
}, error) {
	var out []interface{}
	err := _LilypadEvents.contract.Call(opts, &out, "bacalhauJobCalledHistory", arg0)

	outstruct := new(struct {
		Requestor common.Address
		Id        *big.Int
		Spec      string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Requestor = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Id = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Spec = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// BacalhauJobCalledHistory is a free data retrieval call binding the contract method 0xd3b7667c.
//
// Solidity: function bacalhauJobCalledHistory(uint256 ) view returns(address requestor, uint256 id, string spec)
func (_LilypadEvents *LilypadEventsSession) BacalhauJobCalledHistory(arg0 *big.Int) (struct {
	Requestor common.Address
	Id        *big.Int
	Spec      string
}, error) {
	return _LilypadEvents.Contract.BacalhauJobCalledHistory(&_LilypadEvents.CallOpts, arg0)
}

// BacalhauJobCalledHistory is a free data retrieval call binding the contract method 0xd3b7667c.
//
// Solidity: function bacalhauJobCalledHistory(uint256 ) view returns(address requestor, uint256 id, string spec)
func (_LilypadEvents *LilypadEventsCallerSession) BacalhauJobCalledHistory(arg0 *big.Int) (struct {
	Requestor common.Address
	Id        *big.Int
	Spec      string
}, error) {
	return _LilypadEvents.Contract.BacalhauJobCalledHistory(&_LilypadEvents.CallOpts, arg0)
}

// BacalhauJobHistory is a free data retrieval call binding the contract method 0x807d6dcf.
//
// Solidity: function bacalhauJobHistory(uint256 ) view returns(address requestor, uint256 id, string result)
func (_LilypadEvents *LilypadEventsCaller) BacalhauJobHistory(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Requestor common.Address
	Id        *big.Int
	Result    string
}, error) {
	var out []interface{}
	err := _LilypadEvents.contract.Call(opts, &out, "bacalhauJobHistory", arg0)

	outstruct := new(struct {
		Requestor common.Address
		Id        *big.Int
		Result    string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Requestor = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Id = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Result = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// BacalhauJobHistory is a free data retrieval call binding the contract method 0x807d6dcf.
//
// Solidity: function bacalhauJobHistory(uint256 ) view returns(address requestor, uint256 id, string result)
func (_LilypadEvents *LilypadEventsSession) BacalhauJobHistory(arg0 *big.Int) (struct {
	Requestor common.Address
	Id        *big.Int
	Result    string
}, error) {
	return _LilypadEvents.Contract.BacalhauJobHistory(&_LilypadEvents.CallOpts, arg0)
}

// BacalhauJobHistory is a free data retrieval call binding the contract method 0x807d6dcf.
//
// Solidity: function bacalhauJobHistory(uint256 ) view returns(address requestor, uint256 id, string result)
func (_LilypadEvents *LilypadEventsCallerSession) BacalhauJobHistory(arg0 *big.Int) (struct {
	Requestor common.Address
	Id        *big.Int
	Result    string
}, error) {
	return _LilypadEvents.Contract.BacalhauJobHistory(&_LilypadEvents.CallOpts, arg0)
}

// FetchAllCalledJobs is a free data retrieval call binding the contract method 0x43bd5191.
//
// Solidity: function fetchAllCalledJobs() view returns((address,uint256,string)[])
func (_LilypadEvents *LilypadEventsCaller) FetchAllCalledJobs(opts *bind.CallOpts) ([]LilypadEventsBacalhauJobCalled, error) {
	var out []interface{}
	err := _LilypadEvents.contract.Call(opts, &out, "fetchAllCalledJobs")

	if err != nil {
		return *new([]LilypadEventsBacalhauJobCalled), err
	}

	out0 := *abi.ConvertType(out[0], new([]LilypadEventsBacalhauJobCalled)).(*[]LilypadEventsBacalhauJobCalled)

	return out0, err

}

// FetchAllCalledJobs is a free data retrieval call binding the contract method 0x43bd5191.
//
// Solidity: function fetchAllCalledJobs() view returns((address,uint256,string)[])
func (_LilypadEvents *LilypadEventsSession) FetchAllCalledJobs() ([]LilypadEventsBacalhauJobCalled, error) {
	return _LilypadEvents.Contract.FetchAllCalledJobs(&_LilypadEvents.CallOpts)
}

// FetchAllCalledJobs is a free data retrieval call binding the contract method 0x43bd5191.
//
// Solidity: function fetchAllCalledJobs() view returns((address,uint256,string)[])
func (_LilypadEvents *LilypadEventsCallerSession) FetchAllCalledJobs() ([]LilypadEventsBacalhauJobCalled, error) {
	return _LilypadEvents.Contract.FetchAllCalledJobs(&_LilypadEvents.CallOpts)
}

// FetchAllJobs is a free data retrieval call binding the contract method 0xce4af1d7.
//
// Solidity: function fetchAllJobs() view returns((address,uint256,string)[])
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
// Solidity: function fetchAllJobs() view returns((address,uint256,string)[])
func (_LilypadEvents *LilypadEventsSession) FetchAllJobs() ([]LilypadEventsBacalhauJob, error) {
	return _LilypadEvents.Contract.FetchAllJobs(&_LilypadEvents.CallOpts)
}

// FetchAllJobs is a free data retrieval call binding the contract method 0xce4af1d7.
//
// Solidity: function fetchAllJobs() view returns((address,uint256,string)[])
func (_LilypadEvents *LilypadEventsCallerSession) FetchAllJobs() ([]LilypadEventsBacalhauJob, error) {
	return _LilypadEvents.Contract.FetchAllJobs(&_LilypadEvents.CallOpts)
}

// FetchJobsByAddress is a free data retrieval call binding the contract method 0x399fad37.
//
// Solidity: function fetchJobsByAddress(address _requestor) view returns((address,uint256,string)[])
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
// Solidity: function fetchJobsByAddress(address _requestor) view returns((address,uint256,string)[])
func (_LilypadEvents *LilypadEventsSession) FetchJobsByAddress(_requestor common.Address) ([]LilypadEventsBacalhauJob, error) {
	return _LilypadEvents.Contract.FetchJobsByAddress(&_LilypadEvents.CallOpts, _requestor)
}

// FetchJobsByAddress is a free data retrieval call binding the contract method 0x399fad37.
//
// Solidity: function fetchJobsByAddress(address _requestor) view returns((address,uint256,string)[])
func (_LilypadEvents *LilypadEventsCallerSession) FetchJobsByAddress(_requestor common.Address) ([]LilypadEventsBacalhauJob, error) {
	return _LilypadEvents.Contract.FetchJobsByAddress(&_LilypadEvents.CallOpts, _requestor)
}

// ReturnBacalhauResults is a paid mutator transaction binding the contract method 0xa0322246.
//
// Solidity: function returnBacalhauResults(address _to, uint256 _jobId, string _ipfsResult) returns()
func (_LilypadEvents *LilypadEventsTransactor) ReturnBacalhauResults(opts *bind.TransactOpts, _to common.Address, _jobId *big.Int, _ipfsResult string) (*types.Transaction, error) {
	return _LilypadEvents.contract.Transact(opts, "returnBacalhauResults", _to, _jobId, _ipfsResult)
}

// ReturnBacalhauResults is a paid mutator transaction binding the contract method 0xa0322246.
//
// Solidity: function returnBacalhauResults(address _to, uint256 _jobId, string _ipfsResult) returns()
func (_LilypadEvents *LilypadEventsSession) ReturnBacalhauResults(_to common.Address, _jobId *big.Int, _ipfsResult string) (*types.Transaction, error) {
	return _LilypadEvents.Contract.ReturnBacalhauResults(&_LilypadEvents.TransactOpts, _to, _jobId, _ipfsResult)
}

// ReturnBacalhauResults is a paid mutator transaction binding the contract method 0xa0322246.
//
// Solidity: function returnBacalhauResults(address _to, uint256 _jobId, string _ipfsResult) returns()
func (_LilypadEvents *LilypadEventsTransactorSession) ReturnBacalhauResults(_to common.Address, _jobId *big.Int, _ipfsResult string) (*types.Transaction, error) {
	return _LilypadEvents.Contract.ReturnBacalhauResults(&_LilypadEvents.TransactOpts, _to, _jobId, _ipfsResult)
}

// RunBacalhauJob is a paid mutator transaction binding the contract method 0x46b3c952.
//
// Solidity: function runBacalhauJob(address _from, string _spec) returns()
func (_LilypadEvents *LilypadEventsTransactor) RunBacalhauJob(opts *bind.TransactOpts, _from common.Address, _spec string) (*types.Transaction, error) {
	return _LilypadEvents.contract.Transact(opts, "runBacalhauJob", _from, _spec)
}

// RunBacalhauJob is a paid mutator transaction binding the contract method 0x46b3c952.
//
// Solidity: function runBacalhauJob(address _from, string _spec) returns()
func (_LilypadEvents *LilypadEventsSession) RunBacalhauJob(_from common.Address, _spec string) (*types.Transaction, error) {
	return _LilypadEvents.Contract.RunBacalhauJob(&_LilypadEvents.TransactOpts, _from, _spec)
}

// RunBacalhauJob is a paid mutator transaction binding the contract method 0x46b3c952.
//
// Solidity: function runBacalhauJob(address _from, string _spec) returns()
func (_LilypadEvents *LilypadEventsTransactorSession) RunBacalhauJob(_from common.Address, _spec string) (*types.Transaction, error) {
	return _LilypadEvents.Contract.RunBacalhauJob(&_LilypadEvents.TransactOpts, _from, _spec)
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
	Results           string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBacalhauJobResultsReturned is a free log retrieval operation binding the contract event 0x81fb9923a490992abba509e72b444bd3f44af7259412c153f25604a515f18561.
//
// Solidity: event BacalhauJobResultsReturned(address requestorContract, uint256 jobId, string results)
func (_LilypadEvents *LilypadEventsFilterer) FilterBacalhauJobResultsReturned(opts *bind.FilterOpts) (*LilypadEventsBacalhauJobResultsReturnedIterator, error) {

	logs, sub, err := _LilypadEvents.contract.FilterLogs(opts, "BacalhauJobResultsReturned")
	if err != nil {
		return nil, err
	}
	return &LilypadEventsBacalhauJobResultsReturnedIterator{contract: _LilypadEvents.contract, event: "BacalhauJobResultsReturned", logs: logs, sub: sub}, nil
}

// WatchBacalhauJobResultsReturned is a free log subscription operation binding the contract event 0x81fb9923a490992abba509e72b444bd3f44af7259412c153f25604a515f18561.
//
// Solidity: event BacalhauJobResultsReturned(address requestorContract, uint256 jobId, string results)
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

// ParseBacalhauJobResultsReturned is a log parse operation binding the contract event 0x81fb9923a490992abba509e72b444bd3f44af7259412c153f25604a515f18561.
//
// Solidity: event BacalhauJobResultsReturned(address requestorContract, uint256 jobId, string results)
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
	Id                *big.Int
	Spec              string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewBacalhauJobSubmitted is a free log retrieval operation binding the contract event 0x824b4467ad6ba18ef48fda0d874eb770caa8f8979d165b1d2d6ab73b8fce5cb1.
//
// Solidity: event NewBacalhauJobSubmitted(address indexed requestorContract, uint256 id, string spec)
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

// WatchNewBacalhauJobSubmitted is a free log subscription operation binding the contract event 0x824b4467ad6ba18ef48fda0d874eb770caa8f8979d165b1d2d6ab73b8fce5cb1.
//
// Solidity: event NewBacalhauJobSubmitted(address indexed requestorContract, uint256 id, string spec)
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

// ParseNewBacalhauJobSubmitted is a log parse operation binding the contract event 0x824b4467ad6ba18ef48fda0d874eb770caa8f8979d165b1d2d6ab73b8fce5cb1.
//
// Solidity: event NewBacalhauJobSubmitted(address indexed requestorContract, uint256 id, string spec)
func (_LilypadEvents *LilypadEventsFilterer) ParseNewBacalhauJobSubmitted(log types.Log) (*LilypadEventsNewBacalhauJobSubmitted, error) {
	event := new(LilypadEventsNewBacalhauJobSubmitted)
	if err := _LilypadEvents.contract.UnpackLog(event, "NewBacalhauJobSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}


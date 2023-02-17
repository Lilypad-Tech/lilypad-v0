// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StableDiffusionCaller

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

// StableDiffusionCallerLilypadJob is an auto generated low-level Go binding around an user-defined struct.
type StableDiffusionCallerLilypadJob struct {
	JobId      *big.Int
	IpfsResult string
}

// StableDiffusionCallerMetaData contains all meta data concerning the StableDiffusionCaller contract.
var StableDiffusionCallerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ipfsResult\",\"type\":\"string\"}],\"name\":\"LilypadResponse\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_prompt\",\"type\":\"string\"}],\"name\":\"StableDiffusion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractLilypadEvents\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchAllJobs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ipfsResult\",\"type\":\"string\"}],\"internalType\":\"structStableDiffusionCaller.LilypadJob[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lilypadJobs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ipfsResult\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_jobId\",\"type\":\"uint256\"},{\"internalType\":\"enumLilypadResultType\",\"name\":\"_resultType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_result\",\"type\":\"string\"}],\"name\":\"lilypadReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_eventsAddress\",\"type\":\"address\"}],\"name\":\"setLPEventsAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001553380380620015538339818101604052810190620000379190620001f7565b6200006660405180606001604052806022815260200162001531602291396200007e60201b620005661760201c565b62000077816200012160201b60201c565b50620002e7565b6200011e81604051602401620000959190620002c3565b6040516020818303038152906040527f41304fac000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506200016460201b60201c565b50565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008151905060006a636f6e736f6c652e6c6f679050602083016000808483855afa5050505050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620001bf8262000192565b9050919050565b620001d181620001b2565b8114620001dd57600080fd5b50565b600081519050620001f181620001c6565b92915050565b60006020828403121562000210576200020f6200018d565b5b60006200022084828501620001e0565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156200026557808201518184015260208101905062000248565b60008484015250505050565b6000601f19601f8301169050919050565b60006200028f8262000229565b6200029b818562000234565b9350620002ad81856020860162000245565b620002b88162000271565b840191505092915050565b60006020820190508181036000830152620002df818462000282565b905092915050565b61123a80620002f76000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630972c2fa14610067578063511deb4714610083578063c3357f4e1461009f578063ce4af1d7146100bb578063e78cea92146100d9578063f1cb28ad146100f7575b600080fd5b610081600480360381019061007c919061069a565b610128565b005b61009d60048036038101906100989190610868565b61016b565b005b6100b960048036038101906100b4919061094b565b61029b565b005b6100c361038b565b6040516100d09190610b25565b60405180910390f35b6100e1610486565b6040516100ee9190610ba6565b60405180910390f35b610111600480360381019061010c9190610bc1565b6104aa565b60405161011f929190610c47565b60405180910390f35b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16146101c357600080fd5b600060038111156101d7576101d6610c77565b5b8260038111156101ea576101e9610c77565b5b146101f457600080fd5b6000604051806040016040528085815260200183815250905060018190806001815401808255809150506001900390600052602060002090600202016000909190919091506000820151816000015560208201518160010190816102589190610ea8565b5050507f5658d669440c482d6eed1ca4a3e301ae71f7b56ae0d72c6743316907e62d190a848360405161028c929190610c47565b60405180910390a15050505050565b600060405180610100016040528060d181526020016110c660d1913983836040518060a00160405280606e8152602001611197606e91396040516020016102e59493929190610fdb565b604051602081830303815290604052905060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d992cac8308360006040518463ffffffff1660e01b815260040161035493929190611065565b600060405180830381600087803b15801561036e57600080fd5b505af1158015610382573d6000803e3d6000fd5b50505050505050565b60606001805480602002602001604051908101604052809291908181526020016000905b8282101561047d5783829060005260206000209060020201604051806040016040529081600082015481526020016001820180546103ec90610cd5565b80601f016020809104026020016040519081016040528092919081815260200182805461041890610cd5565b80156104655780601f1061043a57610100808354040283529160200191610465565b820191906000526020600020905b81548152906001019060200180831161044857829003601f168201915b505050505081525050815260200190600101906103af565b50505050905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600181815481106104ba57600080fd5b90600052602060002090600202016000915090508060000154908060010180546104e390610cd5565b80601f016020809104026020016040519081016040528092919081815260200182805461050f90610cd5565b801561055c5780601f106105315761010080835404028352916020019161055c565b820191906000526020600020905b81548152906001019060200180831161053f57829003601f168201915b5050505050905082565b6105fc8160405160240161057a91906110a3565b6040516020818303038152906040527f41304fac000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506105ff565b50565b60008151905060006a636f6e736f6c652e6c6f679050602083016000808483855afa5050505050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006106678261063c565b9050919050565b6106778161065c565b811461068257600080fd5b50565b6000813590506106948161066e565b92915050565b6000602082840312156106b0576106af610632565b5b60006106be84828501610685565b91505092915050565b6000819050919050565b6106da816106c7565b81146106e557600080fd5b50565b6000813590506106f7816106d1565b92915050565b6004811061070a57600080fd5b50565b60008135905061071c816106fd565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6107758261072c565b810181811067ffffffffffffffff821117156107945761079361073d565b5b80604052505050565b60006107a7610628565b90506107b3828261076c565b919050565b600067ffffffffffffffff8211156107d3576107d261073d565b5b6107dc8261072c565b9050602081019050919050565b82818337600083830152505050565b600061080b610806846107b8565b61079d565b90508281526020810184848401111561082757610826610727565b5b6108328482856107e9565b509392505050565b600082601f83011261084f5761084e610722565b5b813561085f8482602086016107f8565b91505092915050565b6000806000806080858703121561088257610881610632565b5b600061089087828801610685565b94505060206108a1878288016106e8565b93505060406108b28782880161070d565b925050606085013567ffffffffffffffff8111156108d3576108d2610637565b5b6108df8782880161083a565b91505092959194509250565b600080fd5b600080fd5b60008083601f84011261090b5761090a610722565b5b8235905067ffffffffffffffff811115610928576109276108eb565b5b602083019150836001820283011115610944576109436108f0565b5b9250929050565b6000806020838503121561096257610961610632565b5b600083013567ffffffffffffffff8111156109805761097f610637565b5b61098c858286016108f5565b92509250509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6109cd816106c7565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610a0d5780820151818401526020810190506109f2565b60008484015250505050565b6000610a24826109d3565b610a2e81856109de565b9350610a3e8185602086016109ef565b610a478161072c565b840191505092915050565b6000604083016000830151610a6a60008601826109c4565b5060208301518482036020860152610a828282610a19565b9150508091505092915050565b6000610a9b8383610a52565b905092915050565b6000602082019050919050565b6000610abb82610998565b610ac581856109a3565b935083602082028501610ad7856109b4565b8060005b85811015610b135784840389528151610af48582610a8f565b9450610aff83610aa3565b925060208a01995050600181019050610adb565b50829750879550505050505092915050565b60006020820190508181036000830152610b3f8184610ab0565b905092915050565b6000819050919050565b6000610b6c610b67610b628461063c565b610b47565b61063c565b9050919050565b6000610b7e82610b51565b9050919050565b6000610b9082610b73565b9050919050565b610ba081610b85565b82525050565b6000602082019050610bbb6000830184610b97565b92915050565b600060208284031215610bd757610bd6610632565b5b6000610be5848285016106e8565b91505092915050565b610bf7816106c7565b82525050565b600082825260208201905092915050565b6000610c19826109d3565b610c238185610bfd565b9350610c338185602086016109ef565b610c3c8161072c565b840191505092915050565b6000604082019050610c5c6000830185610bee565b8181036020830152610c6e8184610c0e565b90509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610ced57607f821691505b602082108103610d0057610cff610ca6565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610d687fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610d2b565b610d728683610d2b565b95508019841693508086168417925050509392505050565b6000610da5610da0610d9b846106c7565b610b47565b6106c7565b9050919050565b6000819050919050565b610dbf83610d8a565b610dd3610dcb82610dac565b848454610d38565b825550505050565b600090565b610de8610ddb565b610df3818484610db6565b505050565b5b81811015610e1757610e0c600082610de0565b600181019050610df9565b5050565b601f821115610e5c57610e2d81610d06565b610e3684610d1b565b81016020851015610e45578190505b610e59610e5185610d1b565b830182610df8565b50505b505050565b600082821c905092915050565b6000610e7f60001984600802610e61565b1980831691505092915050565b6000610e988383610e6e565b9150826002028217905092915050565b610eb1826109d3565b67ffffffffffffffff811115610eca57610ec961073d565b5b610ed48254610cd5565b610edf828285610e1b565b600060209050601f831160018114610f125760008415610f00578287015190505b610f0a8582610e8c565b865550610f72565b601f198416610f2086610d06565b60005b82811015610f4857848901518255600182019150602085019450602081019050610f23565b86831015610f655784890151610f61601f891682610e6e565b8355505b6001600288020188555050505b505050505050565b600081905092915050565b6000610f90826109d3565b610f9a8185610f7a565b9350610faa8185602086016109ef565b80840191505092915050565b6000610fc28385610f7a565b9350610fcf8385846107e9565b82840190509392505050565b6000610fe78287610f85565b9150610ff4828587610fb6565b91506110008284610f85565b915081905095945050505050565b6110178161065c565b82525050565b6004811061102e5761102d610c77565b5b50565b600081905061103f8261101d565b919050565b600061104f82611031565b9050919050565b61105f81611044565b82525050565b600060608201905061107a600083018661100e565b818103602083015261108c8185610c0e565b905061109b6040830184611056565b949350505050565b600060208201905081810360008301526110bd8184610c0e565b90509291505056fe7b22456e67696e65223a2022646f636b6572222c225665726966696572223a20226e6f6f70222c225075626c6973686572223a202265737475617279222c22446f636b6572223a207b22496d616765223a2022676863722e696f2f626163616c6861752d70726f6a6563742f6578616d706c65732f737461626c652d646966667573696f6e2d6770753a302e302e31222c22456e747279706f696e74223a205b22707974686f6e222c20226d61696e2e7079222c20222d2d6f222c20222e2f6f757470757473222c20222d2d70222c2022225d7d2c225265736f7572636573223a207b22475055223a202231227d2c224f757470757473223a205b7b224e616d65223a20226f757470757473222c202250617468223a20222f6f757470757473227d5d2c224465616c223a207b22436f6e63757272656e6379223a20317d7da2646970667358221220726cf0b91079ffe706139369cf849dc069f313dc9e64eec4ded9202ed078827964736f6c634300081100334465706c6f79696e6720537461626c65446966667573696f6e20636f6e7472616374",
}

// StableDiffusionCallerABI is the input ABI used to generate the binding from.
// Deprecated: Use StableDiffusionCallerMetaData.ABI instead.
var StableDiffusionCallerABI = StableDiffusionCallerMetaData.ABI

// StableDiffusionCallerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StableDiffusionCallerMetaData.Bin instead.
var StableDiffusionCallerBin = StableDiffusionCallerMetaData.Bin

// DeployStableDiffusionCaller deploys a new Ethereum contract, binding an instance of StableDiffusionCaller to it.
func DeployStableDiffusionCaller(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeContract common.Address) (common.Address, *types.Transaction, *StableDiffusionCaller, error) {
	parsed, err := StableDiffusionCallerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StableDiffusionCallerBin), backend, bridgeContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StableDiffusionCaller{StableDiffusionCallerCaller: StableDiffusionCallerCaller{contract: contract}, StableDiffusionCallerTransactor: StableDiffusionCallerTransactor{contract: contract}, StableDiffusionCallerFilterer: StableDiffusionCallerFilterer{contract: contract}}, nil
}

// StableDiffusionCaller is an auto generated Go binding around an Ethereum contract.
type StableDiffusionCaller struct {
	StableDiffusionCallerCaller     // Read-only binding to the contract
	StableDiffusionCallerTransactor // Write-only binding to the contract
	StableDiffusionCallerFilterer   // Log filterer for contract events
}

// StableDiffusionCallerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StableDiffusionCallerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StableDiffusionCallerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StableDiffusionCallerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StableDiffusionCallerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StableDiffusionCallerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StableDiffusionCallerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StableDiffusionCallerSession struct {
	Contract     *StableDiffusionCaller // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// StableDiffusionCallerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StableDiffusionCallerCallerSession struct {
	Contract *StableDiffusionCallerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// StableDiffusionCallerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StableDiffusionCallerTransactorSession struct {
	Contract     *StableDiffusionCallerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// StableDiffusionCallerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StableDiffusionCallerRaw struct {
	Contract *StableDiffusionCaller // Generic contract binding to access the raw methods on
}

// StableDiffusionCallerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StableDiffusionCallerCallerRaw struct {
	Contract *StableDiffusionCallerCaller // Generic read-only contract binding to access the raw methods on
}

// StableDiffusionCallerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StableDiffusionCallerTransactorRaw struct {
	Contract *StableDiffusionCallerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStableDiffusionCaller creates a new instance of StableDiffusionCaller, bound to a specific deployed contract.
func NewStableDiffusionCaller(address common.Address, backend bind.ContractBackend) (*StableDiffusionCaller, error) {
	contract, err := bindStableDiffusionCaller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StableDiffusionCaller{StableDiffusionCallerCaller: StableDiffusionCallerCaller{contract: contract}, StableDiffusionCallerTransactor: StableDiffusionCallerTransactor{contract: contract}, StableDiffusionCallerFilterer: StableDiffusionCallerFilterer{contract: contract}}, nil
}

// NewStableDiffusionCallerCaller creates a new read-only instance of StableDiffusionCaller, bound to a specific deployed contract.
func NewStableDiffusionCallerCaller(address common.Address, caller bind.ContractCaller) (*StableDiffusionCallerCaller, error) {
	contract, err := bindStableDiffusionCaller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StableDiffusionCallerCaller{contract: contract}, nil
}

// NewStableDiffusionCallerTransactor creates a new write-only instance of StableDiffusionCaller, bound to a specific deployed contract.
func NewStableDiffusionCallerTransactor(address common.Address, transactor bind.ContractTransactor) (*StableDiffusionCallerTransactor, error) {
	contract, err := bindStableDiffusionCaller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StableDiffusionCallerTransactor{contract: contract}, nil
}

// NewStableDiffusionCallerFilterer creates a new log filterer instance of StableDiffusionCaller, bound to a specific deployed contract.
func NewStableDiffusionCallerFilterer(address common.Address, filterer bind.ContractFilterer) (*StableDiffusionCallerFilterer, error) {
	contract, err := bindStableDiffusionCaller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StableDiffusionCallerFilterer{contract: contract}, nil
}

// bindStableDiffusionCaller binds a generic wrapper to an already deployed contract.
func bindStableDiffusionCaller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StableDiffusionCallerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StableDiffusionCaller *StableDiffusionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StableDiffusionCaller.Contract.StableDiffusionCallerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StableDiffusionCaller *StableDiffusionCallerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StableDiffusionCaller.Contract.StableDiffusionCallerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StableDiffusionCaller *StableDiffusionCallerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StableDiffusionCaller.Contract.StableDiffusionCallerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StableDiffusionCaller *StableDiffusionCallerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StableDiffusionCaller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StableDiffusionCaller *StableDiffusionCallerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StableDiffusionCaller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StableDiffusionCaller *StableDiffusionCallerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StableDiffusionCaller.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_StableDiffusionCaller *StableDiffusionCallerCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StableDiffusionCaller.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_StableDiffusionCaller *StableDiffusionCallerSession) Bridge() (common.Address, error) {
	return _StableDiffusionCaller.Contract.Bridge(&_StableDiffusionCaller.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_StableDiffusionCaller *StableDiffusionCallerCallerSession) Bridge() (common.Address, error) {
	return _StableDiffusionCaller.Contract.Bridge(&_StableDiffusionCaller.CallOpts)
}

// FetchAllJobs is a free data retrieval call binding the contract method 0xce4af1d7.
//
// Solidity: function fetchAllJobs() view returns((uint256,string)[])
func (_StableDiffusionCaller *StableDiffusionCallerCaller) FetchAllJobs(opts *bind.CallOpts) ([]StableDiffusionCallerLilypadJob, error) {
	var out []interface{}
	err := _StableDiffusionCaller.contract.Call(opts, &out, "fetchAllJobs")

	if err != nil {
		return *new([]StableDiffusionCallerLilypadJob), err
	}

	out0 := *abi.ConvertType(out[0], new([]StableDiffusionCallerLilypadJob)).(*[]StableDiffusionCallerLilypadJob)

	return out0, err

}

// FetchAllJobs is a free data retrieval call binding the contract method 0xce4af1d7.
//
// Solidity: function fetchAllJobs() view returns((uint256,string)[])
func (_StableDiffusionCaller *StableDiffusionCallerSession) FetchAllJobs() ([]StableDiffusionCallerLilypadJob, error) {
	return _StableDiffusionCaller.Contract.FetchAllJobs(&_StableDiffusionCaller.CallOpts)
}

// FetchAllJobs is a free data retrieval call binding the contract method 0xce4af1d7.
//
// Solidity: function fetchAllJobs() view returns((uint256,string)[])
func (_StableDiffusionCaller *StableDiffusionCallerCallerSession) FetchAllJobs() ([]StableDiffusionCallerLilypadJob, error) {
	return _StableDiffusionCaller.Contract.FetchAllJobs(&_StableDiffusionCaller.CallOpts)
}

// LilypadJobs is a free data retrieval call binding the contract method 0xf1cb28ad.
//
// Solidity: function lilypadJobs(uint256 ) view returns(uint256 jobId, string ipfsResult)
func (_StableDiffusionCaller *StableDiffusionCallerCaller) LilypadJobs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	JobId      *big.Int
	IpfsResult string
}, error) {
	var out []interface{}
	err := _StableDiffusionCaller.contract.Call(opts, &out, "lilypadJobs", arg0)

	outstruct := new(struct {
		JobId      *big.Int
		IpfsResult string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.JobId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IpfsResult = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// LilypadJobs is a free data retrieval call binding the contract method 0xf1cb28ad.
//
// Solidity: function lilypadJobs(uint256 ) view returns(uint256 jobId, string ipfsResult)
func (_StableDiffusionCaller *StableDiffusionCallerSession) LilypadJobs(arg0 *big.Int) (struct {
	JobId      *big.Int
	IpfsResult string
}, error) {
	return _StableDiffusionCaller.Contract.LilypadJobs(&_StableDiffusionCaller.CallOpts, arg0)
}

// LilypadJobs is a free data retrieval call binding the contract method 0xf1cb28ad.
//
// Solidity: function lilypadJobs(uint256 ) view returns(uint256 jobId, string ipfsResult)
func (_StableDiffusionCaller *StableDiffusionCallerCallerSession) LilypadJobs(arg0 *big.Int) (struct {
	JobId      *big.Int
	IpfsResult string
}, error) {
	return _StableDiffusionCaller.Contract.LilypadJobs(&_StableDiffusionCaller.CallOpts, arg0)
}

// StableDiffusion is a paid mutator transaction binding the contract method 0xc3357f4e.
//
// Solidity: function StableDiffusion(string _prompt) returns()
func (_StableDiffusionCaller *StableDiffusionCallerTransactor) StableDiffusion(opts *bind.TransactOpts, _prompt string) (*types.Transaction, error) {
	return _StableDiffusionCaller.contract.Transact(opts, "StableDiffusion", _prompt)
}

// StableDiffusion is a paid mutator transaction binding the contract method 0xc3357f4e.
//
// Solidity: function StableDiffusion(string _prompt) returns()
func (_StableDiffusionCaller *StableDiffusionCallerSession) StableDiffusion(_prompt string) (*types.Transaction, error) {
	return _StableDiffusionCaller.Contract.StableDiffusion(&_StableDiffusionCaller.TransactOpts, _prompt)
}

// StableDiffusion is a paid mutator transaction binding the contract method 0xc3357f4e.
//
// Solidity: function StableDiffusion(string _prompt) returns()
func (_StableDiffusionCaller *StableDiffusionCallerTransactorSession) StableDiffusion(_prompt string) (*types.Transaction, error) {
	return _StableDiffusionCaller.Contract.StableDiffusion(&_StableDiffusionCaller.TransactOpts, _prompt)
}

// LilypadReceiver is a paid mutator transaction binding the contract method 0x511deb47.
//
// Solidity: function lilypadReceiver(address _from, uint256 _jobId, uint8 _resultType, string _result) returns()
func (_StableDiffusionCaller *StableDiffusionCallerTransactor) LilypadReceiver(opts *bind.TransactOpts, _from common.Address, _jobId *big.Int, _resultType uint8, _result string) (*types.Transaction, error) {
	return _StableDiffusionCaller.contract.Transact(opts, "lilypadReceiver", _from, _jobId, _resultType, _result)
}

// LilypadReceiver is a paid mutator transaction binding the contract method 0x511deb47.
//
// Solidity: function lilypadReceiver(address _from, uint256 _jobId, uint8 _resultType, string _result) returns()
func (_StableDiffusionCaller *StableDiffusionCallerSession) LilypadReceiver(_from common.Address, _jobId *big.Int, _resultType uint8, _result string) (*types.Transaction, error) {
	return _StableDiffusionCaller.Contract.LilypadReceiver(&_StableDiffusionCaller.TransactOpts, _from, _jobId, _resultType, _result)
}

// LilypadReceiver is a paid mutator transaction binding the contract method 0x511deb47.
//
// Solidity: function lilypadReceiver(address _from, uint256 _jobId, uint8 _resultType, string _result) returns()
func (_StableDiffusionCaller *StableDiffusionCallerTransactorSession) LilypadReceiver(_from common.Address, _jobId *big.Int, _resultType uint8, _result string) (*types.Transaction, error) {
	return _StableDiffusionCaller.Contract.LilypadReceiver(&_StableDiffusionCaller.TransactOpts, _from, _jobId, _resultType, _result)
}

// SetLPEventsAddress is a paid mutator transaction binding the contract method 0x0972c2fa.
//
// Solidity: function setLPEventsAddress(address _eventsAddress) returns()
func (_StableDiffusionCaller *StableDiffusionCallerTransactor) SetLPEventsAddress(opts *bind.TransactOpts, _eventsAddress common.Address) (*types.Transaction, error) {
	return _StableDiffusionCaller.contract.Transact(opts, "setLPEventsAddress", _eventsAddress)
}

// SetLPEventsAddress is a paid mutator transaction binding the contract method 0x0972c2fa.
//
// Solidity: function setLPEventsAddress(address _eventsAddress) returns()
func (_StableDiffusionCaller *StableDiffusionCallerSession) SetLPEventsAddress(_eventsAddress common.Address) (*types.Transaction, error) {
	return _StableDiffusionCaller.Contract.SetLPEventsAddress(&_StableDiffusionCaller.TransactOpts, _eventsAddress)
}

// SetLPEventsAddress is a paid mutator transaction binding the contract method 0x0972c2fa.
//
// Solidity: function setLPEventsAddress(address _eventsAddress) returns()
func (_StableDiffusionCaller *StableDiffusionCallerTransactorSession) SetLPEventsAddress(_eventsAddress common.Address) (*types.Transaction, error) {
	return _StableDiffusionCaller.Contract.SetLPEventsAddress(&_StableDiffusionCaller.TransactOpts, _eventsAddress)
}

// StableDiffusionCallerLilypadResponseIterator is returned from FilterLilypadResponse and is used to iterate over the raw logs and unpacked data for LilypadResponse events raised by the StableDiffusionCaller contract.
type StableDiffusionCallerLilypadResponseIterator struct {
	Event *StableDiffusionCallerLilypadResponse // Event containing the contract specifics and raw log

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
func (it *StableDiffusionCallerLilypadResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableDiffusionCallerLilypadResponse)
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
		it.Event = new(StableDiffusionCallerLilypadResponse)
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
func (it *StableDiffusionCallerLilypadResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableDiffusionCallerLilypadResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableDiffusionCallerLilypadResponse represents a LilypadResponse event raised by the StableDiffusionCaller contract.
type StableDiffusionCallerLilypadResponse struct {
	JobId      *big.Int
	IpfsResult string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLilypadResponse is a free log retrieval operation binding the contract event 0x5658d669440c482d6eed1ca4a3e301ae71f7b56ae0d72c6743316907e62d190a.
//
// Solidity: event LilypadResponse(uint256 jobId, string ipfsResult)
func (_StableDiffusionCaller *StableDiffusionCallerFilterer) FilterLilypadResponse(opts *bind.FilterOpts) (*StableDiffusionCallerLilypadResponseIterator, error) {

	logs, sub, err := _StableDiffusionCaller.contract.FilterLogs(opts, "LilypadResponse")
	if err != nil {
		return nil, err
	}
	return &StableDiffusionCallerLilypadResponseIterator{contract: _StableDiffusionCaller.contract, event: "LilypadResponse", logs: logs, sub: sub}, nil
}

// WatchLilypadResponse is a free log subscription operation binding the contract event 0x5658d669440c482d6eed1ca4a3e301ae71f7b56ae0d72c6743316907e62d190a.
//
// Solidity: event LilypadResponse(uint256 jobId, string ipfsResult)
func (_StableDiffusionCaller *StableDiffusionCallerFilterer) WatchLilypadResponse(opts *bind.WatchOpts, sink chan<- *StableDiffusionCallerLilypadResponse) (event.Subscription, error) {

	logs, sub, err := _StableDiffusionCaller.contract.WatchLogs(opts, "LilypadResponse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableDiffusionCallerLilypadResponse)
				if err := _StableDiffusionCaller.contract.UnpackLog(event, "LilypadResponse", log); err != nil {
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

// ParseLilypadResponse is a log parse operation binding the contract event 0x5658d669440c482d6eed1ca4a3e301ae71f7b56ae0d72c6743316907e62d190a.
//
// Solidity: event LilypadResponse(uint256 jobId, string ipfsResult)
func (_StableDiffusionCaller *StableDiffusionCallerFilterer) ParseLilypadResponse(log types.Log) (*StableDiffusionCallerLilypadResponse, error) {
	event := new(StableDiffusionCallerLilypadResponse)
	if err := _StableDiffusionCaller.contract.UnpackLog(event, "LilypadResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}


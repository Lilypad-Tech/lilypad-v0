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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ipfsResult\",\"type\":\"string\"}],\"name\":\"LilypadResponse\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_prompt\",\"type\":\"string\"}],\"name\":\"StableDiffusion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractLilypadEvents\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchAllJobs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ipfsResult\",\"type\":\"string\"}],\"internalType\":\"structStableDiffusionCaller.LilypadJob[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lilypadJobs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ipfsResult\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_jobId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_ipfsResult\",\"type\":\"string\"}],\"name\":\"lilypadReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_eventsAddress\",\"type\":\"address\"}],\"name\":\"setLPEventsAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001460380380620014608339818101604052810190620000379190620001f7565b620000666040518060600160405280602281526020016200143e602291396200007e60201b620005311760201c565b62000077816200012160201b60201c565b50620002e7565b6200011e81604051602401620000959190620002c3565b6040516020818303038152906040527f41304fac000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506200016460201b60201c565b50565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008151905060006a636f6e736f6c652e6c6f679050602083016000808483855afa5050505050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620001bf8262000192565b9050919050565b620001d181620001b2565b8114620001dd57600080fd5b50565b600081519050620001f181620001c6565b92915050565b60006020828403121562000210576200020f6200018d565b5b60006200022084828501620001e0565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156200026557808201518184015260208101905062000248565b60008484015250505050565b6000601f19601f8301169050919050565b60006200028f8262000229565b6200029b818562000234565b9350620002ad81856020860162000245565b620002b88162000271565b840191505092915050565b60006020820190508181036000830152620002df818462000282565b905092915050565b61114780620002f76000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630972c2fa1461006757806335c764e314610083578063c3357f4e1461009f578063ce4af1d7146100bb578063e78cea92146100d9578063f1cb28ad146100f7575b600080fd5b610081600480360381019061007c9190610665565b610128565b005b61009d6004803603810190610098919061080e565b61016b565b005b6100b960048036038101906100b491906108dd565b610269565b005b6100c3610356565b6040516100d09190610ab7565b60405180910390f35b6100e1610451565b6040516100ee9190610b38565b60405180910390f35b610111600480360381019061010c9190610b53565b610475565b60405161011f929190610bd9565b60405180910390f35b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16146101c357600080fd5b6000604051806040016040528084815260200183815250905060018190806001815401808255809150506001900390600052602060002090600202016000909190919091506000820151816000015560208201518160010190816102279190610e0b565b5050507f5658d669440c482d6eed1ca4a3e301ae71f7b56ae0d72c6743316907e62d190a838360405161025b929190610bd9565b60405180910390a150505050565b600060405180610100016040528060d18152602001610fd360d1913983836040518060a00160405280606e81526020016110a4606e91396040516020016102b39493929190610f3e565b604051602081830303815290604052905060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166346b3c95230836040518363ffffffff1660e01b815260040161031f929190610f80565b600060405180830381600087803b15801561033957600080fd5b505af115801561034d573d6000803e3d6000fd5b50505050505050565b60606001805480602002602001604051908101604052809291908181526020016000905b828210156104485783829060005260206000209060020201604051806040016040529081600082015481526020016001820180546103b790610c38565b80601f01602080910402602001604051908101604052809291908181526020018280546103e390610c38565b80156104305780601f1061040557610100808354040283529160200191610430565b820191906000526020600020905b81548152906001019060200180831161041357829003601f168201915b5050505050815250508152602001906001019061037a565b50505050905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6001818154811061048557600080fd5b90600052602060002090600202016000915090508060000154908060010180546104ae90610c38565b80601f01602080910402602001604051908101604052809291908181526020018280546104da90610c38565b80156105275780601f106104fc57610100808354040283529160200191610527565b820191906000526020600020905b81548152906001019060200180831161050a57829003601f168201915b5050505050905082565b6105c7816040516024016105459190610fb0565b6040516020818303038152906040527f41304fac000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506105ca565b50565b60008151905060006a636f6e736f6c652e6c6f679050602083016000808483855afa5050505050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061063282610607565b9050919050565b61064281610627565b811461064d57600080fd5b50565b60008135905061065f81610639565b92915050565b60006020828403121561067b5761067a6105fd565b5b600061068984828501610650565b91505092915050565b6000819050919050565b6106a581610692565b81146106b057600080fd5b50565b6000813590506106c28161069c565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61071b826106d2565b810181811067ffffffffffffffff8211171561073a576107396106e3565b5b80604052505050565b600061074d6105f3565b90506107598282610712565b919050565b600067ffffffffffffffff821115610779576107786106e3565b5b610782826106d2565b9050602081019050919050565b82818337600083830152505050565b60006107b16107ac8461075e565b610743565b9050828152602081018484840111156107cd576107cc6106cd565b5b6107d884828561078f565b509392505050565b600082601f8301126107f5576107f46106c8565b5b813561080584826020860161079e565b91505092915050565b600080600060608486031215610827576108266105fd565b5b600061083586828701610650565b9350506020610846868287016106b3565b925050604084013567ffffffffffffffff81111561086757610866610602565b5b610873868287016107e0565b9150509250925092565b600080fd5b600080fd5b60008083601f84011261089d5761089c6106c8565b5b8235905067ffffffffffffffff8111156108ba576108b961087d565b5b6020830191508360018202830111156108d6576108d5610882565b5b9250929050565b600080602083850312156108f4576108f36105fd565b5b600083013567ffffffffffffffff81111561091257610911610602565b5b61091e85828601610887565b92509250509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61095f81610692565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561099f578082015181840152602081019050610984565b60008484015250505050565b60006109b682610965565b6109c08185610970565b93506109d0818560208601610981565b6109d9816106d2565b840191505092915050565b60006040830160008301516109fc6000860182610956565b5060208301518482036020860152610a1482826109ab565b9150508091505092915050565b6000610a2d83836109e4565b905092915050565b6000602082019050919050565b6000610a4d8261092a565b610a578185610935565b935083602082028501610a6985610946565b8060005b85811015610aa55784840389528151610a868582610a21565b9450610a9183610a35565b925060208a01995050600181019050610a6d565b50829750879550505050505092915050565b60006020820190508181036000830152610ad18184610a42565b905092915050565b6000819050919050565b6000610afe610af9610af484610607565b610ad9565b610607565b9050919050565b6000610b1082610ae3565b9050919050565b6000610b2282610b05565b9050919050565b610b3281610b17565b82525050565b6000602082019050610b4d6000830184610b29565b92915050565b600060208284031215610b6957610b686105fd565b5b6000610b77848285016106b3565b91505092915050565b610b8981610692565b82525050565b600082825260208201905092915050565b6000610bab82610965565b610bb58185610b8f565b9350610bc5818560208601610981565b610bce816106d2565b840191505092915050565b6000604082019050610bee6000830185610b80565b8181036020830152610c008184610ba0565b90509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610c5057607f821691505b602082108103610c6357610c62610c09565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610ccb7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610c8e565b610cd58683610c8e565b95508019841693508086168417925050509392505050565b6000610d08610d03610cfe84610692565b610ad9565b610692565b9050919050565b6000819050919050565b610d2283610ced565b610d36610d2e82610d0f565b848454610c9b565b825550505050565b600090565b610d4b610d3e565b610d56818484610d19565b505050565b5b81811015610d7a57610d6f600082610d43565b600181019050610d5c565b5050565b601f821115610dbf57610d9081610c69565b610d9984610c7e565b81016020851015610da8578190505b610dbc610db485610c7e565b830182610d5b565b50505b505050565b600082821c905092915050565b6000610de260001984600802610dc4565b1980831691505092915050565b6000610dfb8383610dd1565b9150826002028217905092915050565b610e1482610965565b67ffffffffffffffff811115610e2d57610e2c6106e3565b5b610e378254610c38565b610e42828285610d7e565b600060209050601f831160018114610e755760008415610e63578287015190505b610e6d8582610def565b865550610ed5565b601f198416610e8386610c69565b60005b82811015610eab57848901518255600182019150602085019450602081019050610e86565b86831015610ec85784890151610ec4601f891682610dd1565b8355505b6001600288020188555050505b505050505050565b600081905092915050565b6000610ef382610965565b610efd8185610edd565b9350610f0d818560208601610981565b80840191505092915050565b6000610f258385610edd565b9350610f3283858461078f565b82840190509392505050565b6000610f4a8287610ee8565b9150610f57828587610f19565b9150610f638284610ee8565b915081905095945050505050565b610f7a81610627565b82525050565b6000604082019050610f956000830185610f71565b8181036020830152610fa78184610ba0565b90509392505050565b60006020820190508181036000830152610fca8184610ba0565b90509291505056fe7b22456e67696e65223a2022646f636b6572222c225665726966696572223a20226e6f6f70222c225075626c6973686572223a202265737475617279222c22446f636b6572223a207b22496d616765223a2022676863722e696f2f626163616c6861752d70726f6a6563742f6578616d706c65732f737461626c652d646966667573696f6e2d6770753a302e302e31222c22456e747279706f696e74223a205b22707974686f6e222c20226d61696e2e7079222c20222d2d6f222c20222e2f6f757470757473222c20222d2d70222c2022225d7d2c225265736f7572636573223a207b22475055223a202231227d2c224f757470757473223a205b7b224e616d65223a20226f757470757473222c202250617468223a20222f6f757470757473227d5d2c224465616c223a207b22436f6e63757272656e6379223a20317d7da26469706673582212201b6d3087c79618349814465450640a56d1db47008ab5f3018c1b6e450c24994c64736f6c634300081100334465706c6f79696e6720537461626c65446966667573696f6e20636f6e7472616374",
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

// LilypadReceiver is a paid mutator transaction binding the contract method 0x35c764e3.
//
// Solidity: function lilypadReceiver(address _from, uint256 _jobId, string _ipfsResult) returns()
func (_StableDiffusionCaller *StableDiffusionCallerTransactor) LilypadReceiver(opts *bind.TransactOpts, _from common.Address, _jobId *big.Int, _ipfsResult string) (*types.Transaction, error) {
	return _StableDiffusionCaller.contract.Transact(opts, "lilypadReceiver", _from, _jobId, _ipfsResult)
}

// LilypadReceiver is a paid mutator transaction binding the contract method 0x35c764e3.
//
// Solidity: function lilypadReceiver(address _from, uint256 _jobId, string _ipfsResult) returns()
func (_StableDiffusionCaller *StableDiffusionCallerSession) LilypadReceiver(_from common.Address, _jobId *big.Int, _ipfsResult string) (*types.Transaction, error) {
	return _StableDiffusionCaller.Contract.LilypadReceiver(&_StableDiffusionCaller.TransactOpts, _from, _jobId, _ipfsResult)
}

// LilypadReceiver is a paid mutator transaction binding the contract method 0x35c764e3.
//
// Solidity: function lilypadReceiver(address _from, uint256 _jobId, string _ipfsResult) returns()
func (_StableDiffusionCaller *StableDiffusionCallerTransactorSession) LilypadReceiver(_from common.Address, _jobId *big.Int, _ipfsResult string) (*types.Transaction, error) {
	return _StableDiffusionCaller.Contract.LilypadReceiver(&_StableDiffusionCaller.TransactOpts, _from, _jobId, _ipfsResult)
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


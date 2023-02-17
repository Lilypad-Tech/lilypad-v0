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
	Id         *big.Int
	Result     string
	ResultType uint8
}

// LilypadEventsBacalhauJobCalled is an auto generated low-level Go binding around an user-defined struct.
type LilypadEventsBacalhauJobCalled struct {
	Requestor  common.Address
	Id         *big.Int
	Spec       string
	ResultType uint8
}

// LilypadEventsMetaData contains all meta data concerning the LilypadEvents contract.
var LilypadEventsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requestorContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"results\",\"type\":\"string\"}],\"name\":\"BacalhauJobResultsReturned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requestorContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"spec\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumLilypadResultType\",\"name\":\"resultType\",\"type\":\"uint8\"}],\"name\":\"NewBacalhauJobSubmitted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bacalhauJobCalledHistory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"spec\",\"type\":\"string\"},{\"internalType\":\"enumLilypadResultType\",\"name\":\"resultType\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bacalhauJobHistory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"},{\"internalType\":\"enumLilypadResultType\",\"name\":\"resultType\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchAllCalledJobs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"spec\",\"type\":\"string\"},{\"internalType\":\"enumLilypadResultType\",\"name\":\"resultType\",\"type\":\"uint8\"}],\"internalType\":\"structLilypadEvents.BacalhauJobCalled[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchAllJobs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"},{\"internalType\":\"enumLilypadResultType\",\"name\":\"resultType\",\"type\":\"uint8\"}],\"internalType\":\"structLilypadEvents.BacalhauJob[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_requestor\",\"type\":\"address\"}],\"name\":\"fetchJobsByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"},{\"internalType\":\"enumLilypadResultType\",\"name\":\"resultType\",\"type\":\"uint8\"}],\"internalType\":\"structLilypadEvents.BacalhauJob[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_jobId\",\"type\":\"uint256\"},{\"internalType\":\"enumLilypadResultType\",\"name\":\"_resultType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_result\",\"type\":\"string\"}],\"name\":\"returnBacalhauResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_spec\",\"type\":\"string\"},{\"internalType\":\"enumLilypadResultType\",\"name\":\"_resultType\",\"type\":\"uint8\"}],\"name\":\"runBacalhauJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506200005d6040518060400160405280602081526020017f4465706c6f79696e67204c696c797061644576656e747320636f6e74726163748152506200006360201b62000cac1760201c565b620001ed565b62000103816040516024016200007a9190620001c9565b6040516020818303038152906040527f41304fac000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506200010660201b60201c565b50565b60008151905060006a636f6e736f6c652e6c6f679050602083016000808483855afa5050505050565b600081519050919050565b600082825260208201905092915050565b60005b838110156200016b5780820151818401526020810190506200014e565b60008484015250505050565b6000601f19601f8301169050919050565b600062000195826200012f565b620001a181856200013a565b9350620001b38185602086016200014b565b620001be8162000177565b840191505092915050565b60006020820190508181036000830152620001e5818462000188565b905092915050565b61195480620001fd6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063ce4af1d71161005b578063ce4af1d714610103578063d3b7667c14610121578063d992cac814610154578063f280ccb8146101705761007d565b8063399fad371461008257806343bd5191146100b2578063807d6dcf146100d0575b600080fd5b61009c60048036038101906100979190610e04565b61018c565b6040516100a99190611085565b60405180910390f35b6100ba610357565b6040516100c791906111cc565b60405180910390f35b6100ea60048036038101906100e5919061121a565b6104e3565b6040516100fa94939291906112be565b60405180910390f35b61010b6105d8565b6040516101189190611085565b60405180910390f35b61013b6004803603810190610136919061121a565b610764565b60405161014b94939291906112be565b60405180910390f35b61016e60048036038101906101699190611464565b610859565b005b61018a600480360381019061018591906114d3565b6109e5565b005b6060600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b8282101561034c57838290600052602060002090600402016040518060800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820154815260200160028201805461028090611585565b80601f01602080910402602001604051908101604052809291908181526020018280546102ac90611585565b80156102f95780601f106102ce576101008083540402835291602001916102f9565b820191906000526020600020905b8154815290600101906020018083116102dc57829003601f168201915b505050505081526020016003820160009054906101000a900460ff16600381111561032757610326610f15565b5b600381111561033957610338610f15565b5b81525050815260200190600101906101ed565b505050509050919050565b60606002805480602002602001604051908101604052809291908181526020016000905b828210156104da57838290600052602060002090600402016040518060800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820154815260200160028201805461040e90611585565b80601f016020809104026020016040519081016040528092919081815260200182805461043a90611585565b80156104875780601f1061045c57610100808354040283529160200191610487565b820191906000526020600020905b81548152906001019060200180831161046a57829003601f168201915b505050505081526020016003820160009054906101000a900460ff1660038111156104b5576104b4610f15565b5b60038111156104c7576104c6610f15565b5b815250508152602001906001019061037b565b50505050905090565b600181815481106104f357600080fd5b90600052602060002090600402016000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600101549080600201805461054290611585565b80601f016020809104026020016040519081016040528092919081815260200182805461056e90611585565b80156105bb5780601f10610590576101008083540402835291602001916105bb565b820191906000526020600020905b81548152906001019060200180831161059e57829003601f168201915b5050505050908060030160009054906101000a900460ff16905084565b60606001805480602002602001604051908101604052809291908181526020016000905b8282101561075b57838290600052602060002090600402016040518060800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820154815260200160028201805461068f90611585565b80601f01602080910402602001604051908101604052809291908181526020018280546106bb90611585565b80156107085780601f106106dd57610100808354040283529160200191610708565b820191906000526020600020905b8154815290600101906020018083116106eb57829003601f168201915b505050505081526020016003820160009054906101000a900460ff16600381111561073657610735610f15565b5b600381111561074857610747610f15565b5b81525050815260200190600101906105fc565b50505050905090565b6002818154811061077457600080fd5b90600052602060002090600402016000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020180546107c390611585565b80601f01602080910402602001604051908101604052809291908181526020018280546107ef90611585565b801561083c5780601f106108115761010080835404028352916020019161083c565b820191906000526020600020905b81548152906001019060200180831161081f57829003601f168201915b5050505050908060030160009054906101000a900460ff16905084565b61086282610cac565b600061086e6000610d45565b9050600060405180608001604052808673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018581526020018460038111156108b8576108b7610f15565b5b8152509050600281908060018154018082558091505060019003906000526020600020906004020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201908161094f9190611762565b5060608201518160030160006101000a81548160ff0219169083600381111561097b5761097a610f15565b5b021790555050503373ffffffffffffffffffffffffffffffffffffffff167f223daa373e66780d5887e4cf0d785cd8847ea1cc5317faf2c069159746beeae78386866040516109cc93929190611834565b60405180910390a26109de6000610d53565b5050505050565b600060405180608001604052808673ffffffffffffffffffffffffffffffffffffffff168152602001858152602001838152602001846003811115610a2d57610a2c610f15565b5b8152509050600181908060018154018082558091505060019003906000526020600020906004020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101556040820151816002019081610ac49190611762565b5060608201518160030160006101000a81548160ff02191690836003811115610af057610aef610f15565b5b02179055505050600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081908060018154018082558091505060019003906000526020600020906004020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101556040820151816002019081610bc69190611762565b5060608201518160030160006101000a81548160ff02191690836003811115610bf257610bf1610f15565b5b021790555050507f81fb9923a490992abba509e72b444bd3f44af7259412c153f25604a515f18561308584604051610c2c93929190611872565b60405180910390a18473ffffffffffffffffffffffffffffffffffffffff1663511deb47308686866040518563ffffffff1660e01b8152600401610c7394939291906118b0565b600060405180830381600087803b158015610c8d57600080fd5b505af1158015610ca1573d6000803e3d6000fd5b505050505050505050565b610d4281604051602401610cc091906118fc565b6040516020818303038152906040527f41304fac000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610d69565b50565b600081600001549050919050565b6001816000016000828254019250508190555050565b60008151905060006a636f6e736f6c652e6c6f679050602083016000808483855afa5050505050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610dd182610da6565b9050919050565b610de181610dc6565b8114610dec57600080fd5b50565b600081359050610dfe81610dd8565b92915050565b600060208284031215610e1a57610e19610d9c565b5b6000610e2884828501610def565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610e6681610dc6565b82525050565b6000819050919050565b610e7f81610e6c565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610ebf578082015181840152602081019050610ea4565b60008484015250505050565b6000601f19601f8301169050919050565b6000610ee782610e85565b610ef18185610e90565b9350610f01818560208601610ea1565b610f0a81610ecb565b840191505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60048110610f5557610f54610f15565b5b50565b6000819050610f6682610f44565b919050565b6000610f7682610f58565b9050919050565b610f8681610f6b565b82525050565b6000608083016000830151610fa46000860182610e5d565b506020830151610fb76020860182610e76565b5060408301518482036040860152610fcf8282610edc565b9150506060830151610fe46060860182610f7d565b508091505092915050565b6000610ffb8383610f8c565b905092915050565b6000602082019050919050565b600061101b82610e31565b6110258185610e3c565b93508360208202850161103785610e4d565b8060005b8581101561107357848403895281516110548582610fef565b945061105f83611003565b925060208a0199505060018101905061103b565b50829750879550505050505092915050565b6000602082019050818103600083015261109f8184611010565b905092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60006080830160008301516110eb6000860182610e5d565b5060208301516110fe6020860182610e76565b50604083015184820360408601526111168282610edc565b915050606083015161112b6060860182610f7d565b508091505092915050565b600061114283836110d3565b905092915050565b6000602082019050919050565b6000611162826110a7565b61116c81856110b2565b93508360208202850161117e856110c3565b8060005b858110156111ba578484038952815161119b8582611136565b94506111a68361114a565b925060208a01995050600181019050611182565b50829750879550505050505092915050565b600060208201905081810360008301526111e68184611157565b905092915050565b6111f781610e6c565b811461120257600080fd5b50565b600081359050611214816111ee565b92915050565b6000602082840312156112305761122f610d9c565b5b600061123e84828501611205565b91505092915050565b61125081610dc6565b82525050565b61125f81610e6c565b82525050565b600082825260208201905092915050565b600061128182610e85565b61128b8185611265565b935061129b818560208601610ea1565b6112a481610ecb565b840191505092915050565b6112b881610f6b565b82525050565b60006080820190506112d36000830187611247565b6112e06020830186611256565b81810360408301526112f28185611276565b905061130160608301846112af565b95945050505050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61134c82610ecb565b810181811067ffffffffffffffff8211171561136b5761136a611314565b5b80604052505050565b600061137e610d92565b905061138a8282611343565b919050565b600067ffffffffffffffff8211156113aa576113a9611314565b5b6113b382610ecb565b9050602081019050919050565b82818337600083830152505050565b60006113e26113dd8461138f565b611374565b9050828152602081018484840111156113fe576113fd61130f565b5b6114098482856113c0565b509392505050565b600082601f8301126114265761142561130a565b5b81356114368482602086016113cf565b91505092915050565b6004811061144c57600080fd5b50565b60008135905061145e8161143f565b92915050565b60008060006060848603121561147d5761147c610d9c565b5b600061148b86828701610def565b935050602084013567ffffffffffffffff8111156114ac576114ab610da1565b5b6114b886828701611411565b92505060406114c98682870161144f565b9150509250925092565b600080600080608085870312156114ed576114ec610d9c565b5b60006114fb87828801610def565b945050602061150c87828801611205565b935050604061151d8782880161144f565b925050606085013567ffffffffffffffff81111561153e5761153d610da1565b5b61154a87828801611411565b91505092959194509250565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061159d57607f821691505b6020821081036115b0576115af611556565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026116187fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826115db565b61162286836115db565b95508019841693508086168417925050509392505050565b6000819050919050565b600061165f61165a61165584610e6c565b61163a565b610e6c565b9050919050565b6000819050919050565b61167983611644565b61168d61168582611666565b8484546115e8565b825550505050565b600090565b6116a2611695565b6116ad818484611670565b505050565b5b818110156116d1576116c660008261169a565b6001810190506116b3565b5050565b601f821115611716576116e7816115b6565b6116f0846115cb565b810160208510156116ff578190505b61171361170b856115cb565b8301826116b2565b50505b505050565b600082821c905092915050565b60006117396000198460080261171b565b1980831691505092915050565b60006117528383611728565b9150826002028217905092915050565b61176b82610e85565b67ffffffffffffffff81111561178457611783611314565b5b61178e8254611585565b6117998282856116d5565b600060209050601f8311600181146117cc57600084156117ba578287015190505b6117c48582611746565b86555061182c565b601f1984166117da866115b6565b60005b82811015611802578489015182556001820191506020850194506020810190506117dd565b8683101561181f578489015161181b601f891682611728565b8355505b6001600288020188555050505b505050505050565b60006060820190506118496000830186611256565b818103602083015261185b8185611276565b905061186a60408301846112af565b949350505050565b60006060820190506118876000830186611247565b6118946020830185611256565b81810360408301526118a68184611276565b9050949350505050565b60006080820190506118c56000830187611247565b6118d26020830186611256565b6118df60408301856112af565b81810360608301526118f18184611276565b905095945050505050565b600060208201905081810360008301526119168184611276565b90509291505056fea2646970667358221220f171c3acf66362f8639a2f57bb7c6c71b9164370ca4a85689f86e2f2d09de2ee64736f6c63430008110033",
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
// Solidity: function bacalhauJobCalledHistory(uint256 ) view returns(address requestor, uint256 id, string spec, uint8 resultType)
func (_LilypadEvents *LilypadEventsCaller) BacalhauJobCalledHistory(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Requestor  common.Address
	Id         *big.Int
	Spec       string
	ResultType uint8
}, error) {
	var out []interface{}
	err := _LilypadEvents.contract.Call(opts, &out, "bacalhauJobCalledHistory", arg0)

	outstruct := new(struct {
		Requestor  common.Address
		Id         *big.Int
		Spec       string
		ResultType uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Requestor = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Id = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Spec = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ResultType = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// BacalhauJobCalledHistory is a free data retrieval call binding the contract method 0xd3b7667c.
//
// Solidity: function bacalhauJobCalledHistory(uint256 ) view returns(address requestor, uint256 id, string spec, uint8 resultType)
func (_LilypadEvents *LilypadEventsSession) BacalhauJobCalledHistory(arg0 *big.Int) (struct {
	Requestor  common.Address
	Id         *big.Int
	Spec       string
	ResultType uint8
}, error) {
	return _LilypadEvents.Contract.BacalhauJobCalledHistory(&_LilypadEvents.CallOpts, arg0)
}

// BacalhauJobCalledHistory is a free data retrieval call binding the contract method 0xd3b7667c.
//
// Solidity: function bacalhauJobCalledHistory(uint256 ) view returns(address requestor, uint256 id, string spec, uint8 resultType)
func (_LilypadEvents *LilypadEventsCallerSession) BacalhauJobCalledHistory(arg0 *big.Int) (struct {
	Requestor  common.Address
	Id         *big.Int
	Spec       string
	ResultType uint8
}, error) {
	return _LilypadEvents.Contract.BacalhauJobCalledHistory(&_LilypadEvents.CallOpts, arg0)
}

// BacalhauJobHistory is a free data retrieval call binding the contract method 0x807d6dcf.
//
// Solidity: function bacalhauJobHistory(uint256 ) view returns(address requestor, uint256 id, string result, uint8 resultType)
func (_LilypadEvents *LilypadEventsCaller) BacalhauJobHistory(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Requestor  common.Address
	Id         *big.Int
	Result     string
	ResultType uint8
}, error) {
	var out []interface{}
	err := _LilypadEvents.contract.Call(opts, &out, "bacalhauJobHistory", arg0)

	outstruct := new(struct {
		Requestor  common.Address
		Id         *big.Int
		Result     string
		ResultType uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Requestor = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Id = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Result = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ResultType = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// BacalhauJobHistory is a free data retrieval call binding the contract method 0x807d6dcf.
//
// Solidity: function bacalhauJobHistory(uint256 ) view returns(address requestor, uint256 id, string result, uint8 resultType)
func (_LilypadEvents *LilypadEventsSession) BacalhauJobHistory(arg0 *big.Int) (struct {
	Requestor  common.Address
	Id         *big.Int
	Result     string
	ResultType uint8
}, error) {
	return _LilypadEvents.Contract.BacalhauJobHistory(&_LilypadEvents.CallOpts, arg0)
}

// BacalhauJobHistory is a free data retrieval call binding the contract method 0x807d6dcf.
//
// Solidity: function bacalhauJobHistory(uint256 ) view returns(address requestor, uint256 id, string result, uint8 resultType)
func (_LilypadEvents *LilypadEventsCallerSession) BacalhauJobHistory(arg0 *big.Int) (struct {
	Requestor  common.Address
	Id         *big.Int
	Result     string
	ResultType uint8
}, error) {
	return _LilypadEvents.Contract.BacalhauJobHistory(&_LilypadEvents.CallOpts, arg0)
}

// FetchAllCalledJobs is a free data retrieval call binding the contract method 0x43bd5191.
//
// Solidity: function fetchAllCalledJobs() view returns((address,uint256,string,uint8)[])
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
// Solidity: function fetchAllCalledJobs() view returns((address,uint256,string,uint8)[])
func (_LilypadEvents *LilypadEventsSession) FetchAllCalledJobs() ([]LilypadEventsBacalhauJobCalled, error) {
	return _LilypadEvents.Contract.FetchAllCalledJobs(&_LilypadEvents.CallOpts)
}

// FetchAllCalledJobs is a free data retrieval call binding the contract method 0x43bd5191.
//
// Solidity: function fetchAllCalledJobs() view returns((address,uint256,string,uint8)[])
func (_LilypadEvents *LilypadEventsCallerSession) FetchAllCalledJobs() ([]LilypadEventsBacalhauJobCalled, error) {
	return _LilypadEvents.Contract.FetchAllCalledJobs(&_LilypadEvents.CallOpts)
}

// FetchAllJobs is a free data retrieval call binding the contract method 0xce4af1d7.
//
// Solidity: function fetchAllJobs() view returns((address,uint256,string,uint8)[])
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
// Solidity: function fetchAllJobs() view returns((address,uint256,string,uint8)[])
func (_LilypadEvents *LilypadEventsSession) FetchAllJobs() ([]LilypadEventsBacalhauJob, error) {
	return _LilypadEvents.Contract.FetchAllJobs(&_LilypadEvents.CallOpts)
}

// FetchAllJobs is a free data retrieval call binding the contract method 0xce4af1d7.
//
// Solidity: function fetchAllJobs() view returns((address,uint256,string,uint8)[])
func (_LilypadEvents *LilypadEventsCallerSession) FetchAllJobs() ([]LilypadEventsBacalhauJob, error) {
	return _LilypadEvents.Contract.FetchAllJobs(&_LilypadEvents.CallOpts)
}

// FetchJobsByAddress is a free data retrieval call binding the contract method 0x399fad37.
//
// Solidity: function fetchJobsByAddress(address _requestor) view returns((address,uint256,string,uint8)[])
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
// Solidity: function fetchJobsByAddress(address _requestor) view returns((address,uint256,string,uint8)[])
func (_LilypadEvents *LilypadEventsSession) FetchJobsByAddress(_requestor common.Address) ([]LilypadEventsBacalhauJob, error) {
	return _LilypadEvents.Contract.FetchJobsByAddress(&_LilypadEvents.CallOpts, _requestor)
}

// FetchJobsByAddress is a free data retrieval call binding the contract method 0x399fad37.
//
// Solidity: function fetchJobsByAddress(address _requestor) view returns((address,uint256,string,uint8)[])
func (_LilypadEvents *LilypadEventsCallerSession) FetchJobsByAddress(_requestor common.Address) ([]LilypadEventsBacalhauJob, error) {
	return _LilypadEvents.Contract.FetchJobsByAddress(&_LilypadEvents.CallOpts, _requestor)
}

// ReturnBacalhauResults is a paid mutator transaction binding the contract method 0xf280ccb8.
//
// Solidity: function returnBacalhauResults(address _to, uint256 _jobId, uint8 _resultType, string _result) returns()
func (_LilypadEvents *LilypadEventsTransactor) ReturnBacalhauResults(opts *bind.TransactOpts, _to common.Address, _jobId *big.Int, _resultType uint8, _result string) (*types.Transaction, error) {
	return _LilypadEvents.contract.Transact(opts, "returnBacalhauResults", _to, _jobId, _resultType, _result)
}

// ReturnBacalhauResults is a paid mutator transaction binding the contract method 0xf280ccb8.
//
// Solidity: function returnBacalhauResults(address _to, uint256 _jobId, uint8 _resultType, string _result) returns()
func (_LilypadEvents *LilypadEventsSession) ReturnBacalhauResults(_to common.Address, _jobId *big.Int, _resultType uint8, _result string) (*types.Transaction, error) {
	return _LilypadEvents.Contract.ReturnBacalhauResults(&_LilypadEvents.TransactOpts, _to, _jobId, _resultType, _result)
}

// ReturnBacalhauResults is a paid mutator transaction binding the contract method 0xf280ccb8.
//
// Solidity: function returnBacalhauResults(address _to, uint256 _jobId, uint8 _resultType, string _result) returns()
func (_LilypadEvents *LilypadEventsTransactorSession) ReturnBacalhauResults(_to common.Address, _jobId *big.Int, _resultType uint8, _result string) (*types.Transaction, error) {
	return _LilypadEvents.Contract.ReturnBacalhauResults(&_LilypadEvents.TransactOpts, _to, _jobId, _resultType, _result)
}

// RunBacalhauJob is a paid mutator transaction binding the contract method 0xd992cac8.
//
// Solidity: function runBacalhauJob(address _from, string _spec, uint8 _resultType) returns()
func (_LilypadEvents *LilypadEventsTransactor) RunBacalhauJob(opts *bind.TransactOpts, _from common.Address, _spec string, _resultType uint8) (*types.Transaction, error) {
	return _LilypadEvents.contract.Transact(opts, "runBacalhauJob", _from, _spec, _resultType)
}

// RunBacalhauJob is a paid mutator transaction binding the contract method 0xd992cac8.
//
// Solidity: function runBacalhauJob(address _from, string _spec, uint8 _resultType) returns()
func (_LilypadEvents *LilypadEventsSession) RunBacalhauJob(_from common.Address, _spec string, _resultType uint8) (*types.Transaction, error) {
	return _LilypadEvents.Contract.RunBacalhauJob(&_LilypadEvents.TransactOpts, _from, _spec, _resultType)
}

// RunBacalhauJob is a paid mutator transaction binding the contract method 0xd992cac8.
//
// Solidity: function runBacalhauJob(address _from, string _spec, uint8 _resultType) returns()
func (_LilypadEvents *LilypadEventsTransactorSession) RunBacalhauJob(_from common.Address, _spec string, _resultType uint8) (*types.Transaction, error) {
	return _LilypadEvents.Contract.RunBacalhauJob(&_LilypadEvents.TransactOpts, _from, _spec, _resultType)
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
	ResultType        uint8
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewBacalhauJobSubmitted is a free log retrieval operation binding the contract event 0x223daa373e66780d5887e4cf0d785cd8847ea1cc5317faf2c069159746beeae7.
//
// Solidity: event NewBacalhauJobSubmitted(address indexed requestorContract, uint256 id, string spec, uint8 resultType)
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

// WatchNewBacalhauJobSubmitted is a free log subscription operation binding the contract event 0x223daa373e66780d5887e4cf0d785cd8847ea1cc5317faf2c069159746beeae7.
//
// Solidity: event NewBacalhauJobSubmitted(address indexed requestorContract, uint256 id, string spec, uint8 resultType)
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

// ParseNewBacalhauJobSubmitted is a log parse operation binding the contract event 0x223daa373e66780d5887e4cf0d785cd8847ea1cc5317faf2c069159746beeae7.
//
// Solidity: event NewBacalhauJobSubmitted(address indexed requestorContract, uint256 id, string spec, uint8 resultType)
func (_LilypadEvents *LilypadEventsFilterer) ParseNewBacalhauJobSubmitted(log types.Log) (*LilypadEventsNewBacalhauJobSubmitted, error) {
	event := new(LilypadEventsNewBacalhauJobSubmitted)
	if err := _LilypadEvents.contract.UnpackLog(event, "NewBacalhauJobSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}


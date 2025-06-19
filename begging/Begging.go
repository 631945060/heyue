// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BeggingContractV1_flatten_sol_BeggingContractV1

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
	_ = abi.ConvertType
)

// BeggingContractV1FlattenSolBeggingContractV1MetaData contains all meta data concerning the BeggingContractV1FlattenSolBeggingContractV1 contract.
var BeggingContractV1FlattenSolBeggingContractV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"donate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"donations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getChainlinkDataFeedLatestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"}],\"name\":\"getDonation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"priceFeeds\",\"outputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_priceFeedAddress\",\"type\":\"address\"}],\"name\":\"setPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b503360015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610c858061005c5f395ff3fe60806040526004361061007a575f3560e01c80639dcb511a1161004d5780639dcb511a14610122578063cc6cb19a1461015e578063e7078f921461019a578063ed88c68e146101d65761007a565b80633ccfd60b1461007e578063410a1d321461009457806376e11286146100d05780638da5cb5b146100f8575b5f5ffd5b348015610089575f5ffd5b506100926101e0565b005b34801561009f575f5ffd5b506100ba60048036038101906100b5919061071a565b61031d565b6040516100c7919061075d565b60405180910390f35b3480156100db575f5ffd5b506100f660048036038101906100f19190610776565b610362565b005b348015610103575f5ffd5b5061010c610470565b60405161011991906107c3565b60405180910390f35b34801561012d575f5ffd5b506101486004803603810190610143919061071a565b610495565b6040516101559190610837565b60405180910390f35b348015610169575f5ffd5b50610184600480360381019061017f919061071a565b6104c5565b604051610191919061075d565b60405180910390f35b3480156101a5575f5ffd5b506101c060048036038101906101bb919061071a565b6104d9565b6040516101cd9190610868565b60405180910390f35b6101de610626565b005b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461026f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026690610901565b60405180910390fd5b5f4790505f81116102b5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102ac90610969565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f19350505050158015610319573d5f5f3e3d5ffd5b5050565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e8906109d1565b60405180910390fd5b8060025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6002602052805f5260405f205f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f602052805f5260405f205f915090505481565b5f5f60025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036105a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161059f90610a39565b60405180910390fd5b5f8173ffffffffffffffffffffffffffffffffffffffff1663feaf968c6040518163ffffffff1660e01b815260040160a060405180830381865afa1580156105f2573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106169190610aea565b5050509150508092505050919050565b5f3411610668576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065f90610bd1565b60405180910390fd5b345f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546106b39190610c1c565b92505081905550565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6106e9826106c0565b9050919050565b6106f9816106df565b8114610703575f5ffd5b50565b5f81359050610714816106f0565b92915050565b5f6020828403121561072f5761072e6106bc565b5b5f61073c84828501610706565b91505092915050565b5f819050919050565b61075781610745565b82525050565b5f6020820190506107705f83018461074e565b92915050565b5f5f6040838503121561078c5761078b6106bc565b5b5f61079985828601610706565b92505060206107aa85828601610706565b9150509250929050565b6107bd816106df565b82525050565b5f6020820190506107d65f8301846107b4565b92915050565b5f819050919050565b5f6107ff6107fa6107f5846106c0565b6107dc565b6106c0565b9050919050565b5f610810826107e5565b9050919050565b5f61082182610806565b9050919050565b61083181610817565b82525050565b5f60208201905061084a5f830184610828565b92915050565b5f819050919050565b61086281610850565b82525050565b5f60208201905061087b5f830184610859565b92915050565b5f82825260208201905092915050565b7f4f6e6c792074686520636f6e7472616374206f776e65722063616e20776974685f8201527f647261772066756e64732e000000000000000000000000000000000000000000602082015250565b5f6108eb602b83610881565b91506108f682610891565b604082019050919050565b5f6020820190508181035f830152610918816108df565b9050919050565b7f4e6f2066756e647320746f2077697468647261772e00000000000000000000005f82015250565b5f610953601583610881565b915061095e8261091f565b602082019050919050565b5f6020820190508181035f83015261098081610947565b9050919050565b7f4f6e6c79206f776e65722063616e2073657420707269636520666565640000005f82015250565b5f6109bb601d83610881565b91506109c682610987565b602082019050919050565b5f6020820190508181035f8301526109e8816109af565b9050919050565b7f50726963652066656564206e6f742073657400000000000000000000000000005f82015250565b5f610a23601283610881565b9150610a2e826109ef565b602082019050919050565b5f6020820190508181035f830152610a5081610a17565b9050919050565b5f69ffffffffffffffffffff82169050919050565b610a7581610a57565b8114610a7f575f5ffd5b50565b5f81519050610a9081610a6c565b92915050565b610a9f81610850565b8114610aa9575f5ffd5b50565b5f81519050610aba81610a96565b92915050565b610ac981610745565b8114610ad3575f5ffd5b50565b5f81519050610ae481610ac0565b92915050565b5f5f5f5f5f60a08688031215610b0357610b026106bc565b5b5f610b1088828901610a82565b9550506020610b2188828901610aac565b9450506040610b3288828901610ad6565b9350506060610b4388828901610ad6565b9250506080610b5488828901610a82565b9150509295509295909350565b7f446f6e6174696f6e20616d6f756e74206d7573742062652067726561746572205f8201527f7468616e20302e00000000000000000000000000000000000000000000000000602082015250565b5f610bbb602783610881565b9150610bc682610b61565b604082019050919050565b5f6020820190508181035f830152610be881610baf565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610c2682610745565b9150610c3183610745565b9250828201905080821115610c4957610c48610bef565b5b9291505056fea264697066735822122087df52dcb3f7f9bc3e113fcc684ce1412bdc8d4660e15fab0acaf8d1f7f8689564736f6c634300081e0033",
}

// BeggingContractV1FlattenSolBeggingContractV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use BeggingContractV1FlattenSolBeggingContractV1MetaData.ABI instead.
var BeggingContractV1FlattenSolBeggingContractV1ABI = BeggingContractV1FlattenSolBeggingContractV1MetaData.ABI

// BeggingContractV1FlattenSolBeggingContractV1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BeggingContractV1FlattenSolBeggingContractV1MetaData.Bin instead.
var BeggingContractV1FlattenSolBeggingContractV1Bin = BeggingContractV1FlattenSolBeggingContractV1MetaData.Bin

// DeployBeggingContractV1FlattenSolBeggingContractV1 deploys a new Ethereum contract, binding an instance of BeggingContractV1FlattenSolBeggingContractV1 to it.
func DeployBeggingContractV1FlattenSolBeggingContractV1(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BeggingContractV1FlattenSolBeggingContractV1, error) {
	parsed, err := BeggingContractV1FlattenSolBeggingContractV1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BeggingContractV1FlattenSolBeggingContractV1Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BeggingContractV1FlattenSolBeggingContractV1{BeggingContractV1FlattenSolBeggingContractV1Caller: BeggingContractV1FlattenSolBeggingContractV1Caller{contract: contract}, BeggingContractV1FlattenSolBeggingContractV1Transactor: BeggingContractV1FlattenSolBeggingContractV1Transactor{contract: contract}, BeggingContractV1FlattenSolBeggingContractV1Filterer: BeggingContractV1FlattenSolBeggingContractV1Filterer{contract: contract}}, nil
}

// BeggingContractV1FlattenSolBeggingContractV1 is an auto generated Go binding around an Ethereum contract.
type BeggingContractV1FlattenSolBeggingContractV1 struct {
	BeggingContractV1FlattenSolBeggingContractV1Caller     // Read-only binding to the contract
	BeggingContractV1FlattenSolBeggingContractV1Transactor // Write-only binding to the contract
	BeggingContractV1FlattenSolBeggingContractV1Filterer   // Log filterer for contract events
}

// BeggingContractV1FlattenSolBeggingContractV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type BeggingContractV1FlattenSolBeggingContractV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeggingContractV1FlattenSolBeggingContractV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BeggingContractV1FlattenSolBeggingContractV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeggingContractV1FlattenSolBeggingContractV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BeggingContractV1FlattenSolBeggingContractV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeggingContractV1FlattenSolBeggingContractV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BeggingContractV1FlattenSolBeggingContractV1Session struct {
	Contract     *BeggingContractV1FlattenSolBeggingContractV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts                             // Transaction auth options to use throughout this session
}

// BeggingContractV1FlattenSolBeggingContractV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BeggingContractV1FlattenSolBeggingContractV1CallerSession struct {
	Contract *BeggingContractV1FlattenSolBeggingContractV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                       // Call options to use throughout this session
}

// BeggingContractV1FlattenSolBeggingContractV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BeggingContractV1FlattenSolBeggingContractV1TransactorSession struct {
	Contract     *BeggingContractV1FlattenSolBeggingContractV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                       // Transaction auth options to use throughout this session
}

// BeggingContractV1FlattenSolBeggingContractV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type BeggingContractV1FlattenSolBeggingContractV1Raw struct {
	Contract *BeggingContractV1FlattenSolBeggingContractV1 // Generic contract binding to access the raw methods on
}

// BeggingContractV1FlattenSolBeggingContractV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BeggingContractV1FlattenSolBeggingContractV1CallerRaw struct {
	Contract *BeggingContractV1FlattenSolBeggingContractV1Caller // Generic read-only contract binding to access the raw methods on
}

// BeggingContractV1FlattenSolBeggingContractV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BeggingContractV1FlattenSolBeggingContractV1TransactorRaw struct {
	Contract *BeggingContractV1FlattenSolBeggingContractV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBeggingContractV1FlattenSolBeggingContractV1 creates a new instance of BeggingContractV1FlattenSolBeggingContractV1, bound to a specific deployed contract.
func NewBeggingContractV1FlattenSolBeggingContractV1(address common.Address, backend bind.ContractBackend) (*BeggingContractV1FlattenSolBeggingContractV1, error) {
	contract, err := bindBeggingContractV1FlattenSolBeggingContractV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BeggingContractV1FlattenSolBeggingContractV1{BeggingContractV1FlattenSolBeggingContractV1Caller: BeggingContractV1FlattenSolBeggingContractV1Caller{contract: contract}, BeggingContractV1FlattenSolBeggingContractV1Transactor: BeggingContractV1FlattenSolBeggingContractV1Transactor{contract: contract}, BeggingContractV1FlattenSolBeggingContractV1Filterer: BeggingContractV1FlattenSolBeggingContractV1Filterer{contract: contract}}, nil
}

// NewBeggingContractV1FlattenSolBeggingContractV1Caller creates a new read-only instance of BeggingContractV1FlattenSolBeggingContractV1, bound to a specific deployed contract.
func NewBeggingContractV1FlattenSolBeggingContractV1Caller(address common.Address, caller bind.ContractCaller) (*BeggingContractV1FlattenSolBeggingContractV1Caller, error) {
	contract, err := bindBeggingContractV1FlattenSolBeggingContractV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BeggingContractV1FlattenSolBeggingContractV1Caller{contract: contract}, nil
}

// NewBeggingContractV1FlattenSolBeggingContractV1Transactor creates a new write-only instance of BeggingContractV1FlattenSolBeggingContractV1, bound to a specific deployed contract.
func NewBeggingContractV1FlattenSolBeggingContractV1Transactor(address common.Address, transactor bind.ContractTransactor) (*BeggingContractV1FlattenSolBeggingContractV1Transactor, error) {
	contract, err := bindBeggingContractV1FlattenSolBeggingContractV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BeggingContractV1FlattenSolBeggingContractV1Transactor{contract: contract}, nil
}

// NewBeggingContractV1FlattenSolBeggingContractV1Filterer creates a new log filterer instance of BeggingContractV1FlattenSolBeggingContractV1, bound to a specific deployed contract.
func NewBeggingContractV1FlattenSolBeggingContractV1Filterer(address common.Address, filterer bind.ContractFilterer) (*BeggingContractV1FlattenSolBeggingContractV1Filterer, error) {
	contract, err := bindBeggingContractV1FlattenSolBeggingContractV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BeggingContractV1FlattenSolBeggingContractV1Filterer{contract: contract}, nil
}

// bindBeggingContractV1FlattenSolBeggingContractV1 binds a generic wrapper to an already deployed contract.
func bindBeggingContractV1FlattenSolBeggingContractV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BeggingContractV1FlattenSolBeggingContractV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeggingContractV1FlattenSolBeggingContractV1.Contract.BeggingContractV1FlattenSolBeggingContractV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeggingContractV1FlattenSolBeggingContractV1.Contract.BeggingContractV1FlattenSolBeggingContractV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeggingContractV1FlattenSolBeggingContractV1.Contract.BeggingContractV1FlattenSolBeggingContractV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeggingContractV1FlattenSolBeggingContractV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeggingContractV1FlattenSolBeggingContractV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeggingContractV1FlattenSolBeggingContractV1.Contract.contract.Transact(opts, method, params...)
}

// Donations is a free data retrieval call binding the contract method 0xcc6cb19a.
//
// Solidity: function donations(address ) view returns(uint256)
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1Caller) Donations(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeggingContractV1FlattenSolBeggingContractV1.contract.Call(opts, &out, "donations", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Donations is a free data retrieval call binding the contract method 0xcc6cb19a.
//
// Solidity: function donations(address ) view returns(uint256)
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1Session) Donations(arg0 common.Address) (*big.Int, error) {
	return _BeggingContractV1FlattenSolBeggingContractV1.Contract.Donations(&_BeggingContractV1FlattenSolBeggingContractV1.CallOpts, arg0)
}

// Donations is a free data retrieval call binding the contract method 0xcc6cb19a.
//
// Solidity: function donations(address ) view returns(uint256)
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1CallerSession) Donations(arg0 common.Address) (*big.Int, error) {
	return _BeggingContractV1FlattenSolBeggingContractV1.Contract.Donations(&_BeggingContractV1FlattenSolBeggingContractV1.CallOpts, arg0)
}

// GetChainlinkDataFeedLatestAnswer is a free data retrieval call binding the contract method 0xe7078f92.
//
// Solidity: function getChainlinkDataFeedLatestAnswer(address tokenAddress) view returns(int256)
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1Caller) GetChainlinkDataFeedLatestAnswer(opts *bind.CallOpts, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeggingContractV1FlattenSolBeggingContractV1.contract.Call(opts, &out, "getChainlinkDataFeedLatestAnswer", tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainlinkDataFeedLatestAnswer is a free data retrieval call binding the contract method 0xe7078f92.
//
// Solidity: function getChainlinkDataFeedLatestAnswer(address tokenAddress) view returns(int256)
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1Session) GetChainlinkDataFeedLatestAnswer(tokenAddress common.Address) (*big.Int, error) {
	return _BeggingContractV1FlattenSolBeggingContractV1.Contract.GetChainlinkDataFeedLatestAnswer(&_BeggingContractV1FlattenSolBeggingContractV1.CallOpts, tokenAddress)
}

// GetChainlinkDataFeedLatestAnswer is a free data retrieval call binding the contract method 0xe7078f92.
//
// Solidity: function getChainlinkDataFeedLatestAnswer(address tokenAddress) view returns(int256)
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1CallerSession) GetChainlinkDataFeedLatestAnswer(tokenAddress common.Address) (*big.Int, error) {
	return _BeggingContractV1FlattenSolBeggingContractV1.Contract.GetChainlinkDataFeedLatestAnswer(&_BeggingContractV1FlattenSolBeggingContractV1.CallOpts, tokenAddress)
}

// GetDonation is a free data retrieval call binding the contract method 0x410a1d32.
//
// Solidity: function getDonation(address donor) view returns(uint256)
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1Caller) GetDonation(opts *bind.CallOpts, donor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeggingContractV1FlattenSolBeggingContractV1.contract.Call(opts, &out, "getDonation", donor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDonation is a free data retrieval call binding the contract method 0x410a1d32.
//
// Solidity: function getDonation(address donor) view returns(uint256)
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1Session) GetDonation(donor common.Address) (*big.Int, error) {
	return _BeggingContractV1FlattenSolBeggingContractV1.Contract.GetDonation(&_BeggingContractV1FlattenSolBeggingContractV1.CallOpts, donor)
}

// GetDonation is a free data retrieval call binding the contract method 0x410a1d32.
//
// Solidity: function getDonation(address donor) view returns(uint256)
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1CallerSession) GetDonation(donor common.Address) (*big.Int, error) {
	return _BeggingContractV1FlattenSolBeggingContractV1.Contract.GetDonation(&_BeggingContractV1FlattenSolBeggingContractV1.CallOpts, donor)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeggingContractV1FlattenSolBeggingContractV1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1Session) Owner() (common.Address, error) {
	return _BeggingContractV1FlattenSolBeggingContractV1.Contract.Owner(&_BeggingContractV1FlattenSolBeggingContractV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1CallerSession) Owner() (common.Address, error) {
	return _BeggingContractV1FlattenSolBeggingContractV1.Contract.Owner(&_BeggingContractV1FlattenSolBeggingContractV1.CallOpts)
}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address ) view returns(address)
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1Caller) PriceFeeds(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _BeggingContractV1FlattenSolBeggingContractV1.contract.Call(opts, &out, "priceFeeds", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address ) view returns(address)
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1Session) PriceFeeds(arg0 common.Address) (common.Address, error) {
	return _BeggingContractV1FlattenSolBeggingContractV1.Contract.PriceFeeds(&_BeggingContractV1FlattenSolBeggingContractV1.CallOpts, arg0)
}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address ) view returns(address)
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1CallerSession) PriceFeeds(arg0 common.Address) (common.Address, error) {
	return _BeggingContractV1FlattenSolBeggingContractV1.Contract.PriceFeeds(&_BeggingContractV1FlattenSolBeggingContractV1.CallOpts, arg0)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1Transactor) Donate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeggingContractV1FlattenSolBeggingContractV1.contract.Transact(opts, "donate")
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1Session) Donate() (*types.Transaction, error) {
	return _BeggingContractV1FlattenSolBeggingContractV1.Contract.Donate(&_BeggingContractV1FlattenSolBeggingContractV1.TransactOpts)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1TransactorSession) Donate() (*types.Transaction, error) {
	return _BeggingContractV1FlattenSolBeggingContractV1.Contract.Donate(&_BeggingContractV1FlattenSolBeggingContractV1.TransactOpts)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x76e11286.
//
// Solidity: function setPriceFeed(address tokenAddress, address _priceFeedAddress) returns()
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1Transactor) SetPriceFeed(opts *bind.TransactOpts, tokenAddress common.Address, _priceFeedAddress common.Address) (*types.Transaction, error) {
	return _BeggingContractV1FlattenSolBeggingContractV1.contract.Transact(opts, "setPriceFeed", tokenAddress, _priceFeedAddress)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x76e11286.
//
// Solidity: function setPriceFeed(address tokenAddress, address _priceFeedAddress) returns()
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1Session) SetPriceFeed(tokenAddress common.Address, _priceFeedAddress common.Address) (*types.Transaction, error) {
	return _BeggingContractV1FlattenSolBeggingContractV1.Contract.SetPriceFeed(&_BeggingContractV1FlattenSolBeggingContractV1.TransactOpts, tokenAddress, _priceFeedAddress)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x76e11286.
//
// Solidity: function setPriceFeed(address tokenAddress, address _priceFeedAddress) returns()
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1TransactorSession) SetPriceFeed(tokenAddress common.Address, _priceFeedAddress common.Address) (*types.Transaction, error) {
	return _BeggingContractV1FlattenSolBeggingContractV1.Contract.SetPriceFeed(&_BeggingContractV1FlattenSolBeggingContractV1.TransactOpts, tokenAddress, _priceFeedAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1Transactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeggingContractV1FlattenSolBeggingContractV1.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1Session) Withdraw() (*types.Transaction, error) {
	return _BeggingContractV1FlattenSolBeggingContractV1.Contract.Withdraw(&_BeggingContractV1FlattenSolBeggingContractV1.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1TransactorSession) Withdraw() (*types.Transaction, error) {
	return _BeggingContractV1FlattenSolBeggingContractV1.Contract.Withdraw(&_BeggingContractV1FlattenSolBeggingContractV1.TransactOpts)
}

// BeggingContractV1FlattenSolBeggingContractV1ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BeggingContractV1FlattenSolBeggingContractV1 contract.
type BeggingContractV1FlattenSolBeggingContractV1ApprovalIterator struct {
	Event *BeggingContractV1FlattenSolBeggingContractV1Approval // Event containing the contract specifics and raw log

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
func (it *BeggingContractV1FlattenSolBeggingContractV1ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeggingContractV1FlattenSolBeggingContractV1Approval)
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
		it.Event = new(BeggingContractV1FlattenSolBeggingContractV1Approval)
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
func (it *BeggingContractV1FlattenSolBeggingContractV1ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeggingContractV1FlattenSolBeggingContractV1ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeggingContractV1FlattenSolBeggingContractV1Approval represents a Approval event raised by the BeggingContractV1FlattenSolBeggingContractV1 contract.
type BeggingContractV1FlattenSolBeggingContractV1Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BeggingContractV1FlattenSolBeggingContractV1ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BeggingContractV1FlattenSolBeggingContractV1.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BeggingContractV1FlattenSolBeggingContractV1ApprovalIterator{contract: _BeggingContractV1FlattenSolBeggingContractV1.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BeggingContractV1FlattenSolBeggingContractV1Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BeggingContractV1FlattenSolBeggingContractV1.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeggingContractV1FlattenSolBeggingContractV1Approval)
				if err := _BeggingContractV1FlattenSolBeggingContractV1.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BeggingContractV1FlattenSolBeggingContractV1 *BeggingContractV1FlattenSolBeggingContractV1Filterer) ParseApproval(log types.Log) (*BeggingContractV1FlattenSolBeggingContractV1Approval, error) {
	event := new(BeggingContractV1FlattenSolBeggingContractV1Approval)
	if err := _BeggingContractV1FlattenSolBeggingContractV1.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package multicall

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

// MulticallMetaData contains all meta data concerning the Multicall contract.
var MulticallMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"depositedBalances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"userFirst\",\"type\":\"bool\"}],\"name\":\"depositedBalancesGeneric\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"depositedEtherGeneric\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"functionSignature\",\"type\":\"string\"}],\"name\":\"getFunctionSelector\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spenderContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"tokenAllowances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"allowances\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"tokenBalances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MulticallABI is the input ABI used to generate the binding from.
// Deprecated: Use MulticallMetaData.ABI instead.
var MulticallABI = MulticallMetaData.ABI

// Multicall is an auto generated Go binding around an Ethereum contract.
type Multicall struct {
	MulticallCaller     // Read-only binding to the contract
	MulticallTransactor // Write-only binding to the contract
	MulticallFilterer   // Log filterer for contract events
}

// MulticallCaller is an auto generated read-only Go binding around an Ethereum contract.
type MulticallCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MulticallTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MulticallFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MulticallSession struct {
	Contract     *Multicall        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MulticallCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MulticallCallerSession struct {
	Contract *MulticallCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MulticallTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MulticallTransactorSession struct {
	Contract     *MulticallTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MulticallRaw is an auto generated low-level Go binding around an Ethereum contract.
type MulticallRaw struct {
	Contract *Multicall // Generic contract binding to access the raw methods on
}

// MulticallCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MulticallCallerRaw struct {
	Contract *MulticallCaller // Generic read-only contract binding to access the raw methods on
}

// MulticallTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MulticallTransactorRaw struct {
	Contract *MulticallTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMulticall creates a new instance of Multicall, bound to a specific deployed contract.
func NewMulticall(address common.Address, backend bind.ContractBackend) (*Multicall, error) {
	contract, err := bindMulticall(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Multicall{MulticallCaller: MulticallCaller{contract: contract}, MulticallTransactor: MulticallTransactor{contract: contract}, MulticallFilterer: MulticallFilterer{contract: contract}}, nil
}

// NewMulticallCaller creates a new read-only instance of Multicall, bound to a specific deployed contract.
func NewMulticallCaller(address common.Address, caller bind.ContractCaller) (*MulticallCaller, error) {
	contract, err := bindMulticall(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallCaller{contract: contract}, nil
}

// NewMulticallTransactor creates a new write-only instance of Multicall, bound to a specific deployed contract.
func NewMulticallTransactor(address common.Address, transactor bind.ContractTransactor) (*MulticallTransactor, error) {
	contract, err := bindMulticall(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallTransactor{contract: contract}, nil
}

// NewMulticallFilterer creates a new log filterer instance of Multicall, bound to a specific deployed contract.
func NewMulticallFilterer(address common.Address, filterer bind.ContractFilterer) (*MulticallFilterer, error) {
	contract, err := bindMulticall(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MulticallFilterer{contract: contract}, nil
}

// bindMulticall binds a generic wrapper to an already deployed contract.
func bindMulticall(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MulticallABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multicall *MulticallRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multicall.Contract.MulticallCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multicall *MulticallRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multicall.Contract.MulticallTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multicall *MulticallRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multicall.Contract.MulticallTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multicall *MulticallCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multicall.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multicall *MulticallTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multicall.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multicall *MulticallTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multicall.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Multicall *MulticallCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Multicall.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Multicall *MulticallSession) Admin() (common.Address, error) {
	return _Multicall.Contract.Admin(&_Multicall.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Multicall *MulticallCallerSession) Admin() (common.Address, error) {
	return _Multicall.Contract.Admin(&_Multicall.CallOpts)
}

// DepositedBalances is a free data retrieval call binding the contract method 0x1621c4e5.
//
// Solidity: function depositedBalances(address exchange, address user, address[] tokens) view returns(uint256[] balances)
func (_Multicall *MulticallCaller) DepositedBalances(opts *bind.CallOpts, exchange common.Address, user common.Address, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Multicall.contract.Call(opts, &out, "depositedBalances", exchange, user, tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// DepositedBalances is a free data retrieval call binding the contract method 0x1621c4e5.
//
// Solidity: function depositedBalances(address exchange, address user, address[] tokens) view returns(uint256[] balances)
func (_Multicall *MulticallSession) DepositedBalances(exchange common.Address, user common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _Multicall.Contract.DepositedBalances(&_Multicall.CallOpts, exchange, user, tokens)
}

// DepositedBalances is a free data retrieval call binding the contract method 0x1621c4e5.
//
// Solidity: function depositedBalances(address exchange, address user, address[] tokens) view returns(uint256[] balances)
func (_Multicall *MulticallCallerSession) DepositedBalances(exchange common.Address, user common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _Multicall.Contract.DepositedBalances(&_Multicall.CallOpts, exchange, user, tokens)
}

// DepositedBalancesGeneric is a free data retrieval call binding the contract method 0x83dfb65a.
//
// Solidity: function depositedBalancesGeneric(address exchange, bytes4 selector, address user, address[] tokens, bool userFirst) view returns(uint256[] balances)
func (_Multicall *MulticallCaller) DepositedBalancesGeneric(opts *bind.CallOpts, exchange common.Address, selector [4]byte, user common.Address, tokens []common.Address, userFirst bool) ([]*big.Int, error) {
	var out []interface{}
	err := _Multicall.contract.Call(opts, &out, "depositedBalancesGeneric", exchange, selector, user, tokens, userFirst)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// DepositedBalancesGeneric is a free data retrieval call binding the contract method 0x83dfb65a.
//
// Solidity: function depositedBalancesGeneric(address exchange, bytes4 selector, address user, address[] tokens, bool userFirst) view returns(uint256[] balances)
func (_Multicall *MulticallSession) DepositedBalancesGeneric(exchange common.Address, selector [4]byte, user common.Address, tokens []common.Address, userFirst bool) ([]*big.Int, error) {
	return _Multicall.Contract.DepositedBalancesGeneric(&_Multicall.CallOpts, exchange, selector, user, tokens, userFirst)
}

// DepositedBalancesGeneric is a free data retrieval call binding the contract method 0x83dfb65a.
//
// Solidity: function depositedBalancesGeneric(address exchange, bytes4 selector, address user, address[] tokens, bool userFirst) view returns(uint256[] balances)
func (_Multicall *MulticallCallerSession) DepositedBalancesGeneric(exchange common.Address, selector [4]byte, user common.Address, tokens []common.Address, userFirst bool) ([]*big.Int, error) {
	return _Multicall.Contract.DepositedBalancesGeneric(&_Multicall.CallOpts, exchange, selector, user, tokens, userFirst)
}

// DepositedEtherGeneric is a free data retrieval call binding the contract method 0xfe8b2707.
//
// Solidity: function depositedEtherGeneric(address exchange, bytes4 selector, address user) view returns(uint256)
func (_Multicall *MulticallCaller) DepositedEtherGeneric(opts *bind.CallOpts, exchange common.Address, selector [4]byte, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Multicall.contract.Call(opts, &out, "depositedEtherGeneric", exchange, selector, user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositedEtherGeneric is a free data retrieval call binding the contract method 0xfe8b2707.
//
// Solidity: function depositedEtherGeneric(address exchange, bytes4 selector, address user) view returns(uint256)
func (_Multicall *MulticallSession) DepositedEtherGeneric(exchange common.Address, selector [4]byte, user common.Address) (*big.Int, error) {
	return _Multicall.Contract.DepositedEtherGeneric(&_Multicall.CallOpts, exchange, selector, user)
}

// DepositedEtherGeneric is a free data retrieval call binding the contract method 0xfe8b2707.
//
// Solidity: function depositedEtherGeneric(address exchange, bytes4 selector, address user) view returns(uint256)
func (_Multicall *MulticallCallerSession) DepositedEtherGeneric(exchange common.Address, selector [4]byte, user common.Address) (*big.Int, error) {
	return _Multicall.Contract.DepositedEtherGeneric(&_Multicall.CallOpts, exchange, selector, user)
}

// GetFunctionSelector is a free data retrieval call binding the contract method 0xde38c3d0.
//
// Solidity: function getFunctionSelector(string functionSignature) pure returns(bytes4)
func (_Multicall *MulticallCaller) GetFunctionSelector(opts *bind.CallOpts, functionSignature string) ([4]byte, error) {
	var out []interface{}
	err := _Multicall.contract.Call(opts, &out, "getFunctionSelector", functionSignature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// GetFunctionSelector is a free data retrieval call binding the contract method 0xde38c3d0.
//
// Solidity: function getFunctionSelector(string functionSignature) pure returns(bytes4)
func (_Multicall *MulticallSession) GetFunctionSelector(functionSignature string) ([4]byte, error) {
	return _Multicall.Contract.GetFunctionSelector(&_Multicall.CallOpts, functionSignature)
}

// GetFunctionSelector is a free data retrieval call binding the contract method 0xde38c3d0.
//
// Solidity: function getFunctionSelector(string functionSignature) pure returns(bytes4)
func (_Multicall *MulticallCallerSession) GetFunctionSelector(functionSignature string) ([4]byte, error) {
	return _Multicall.Contract.GetFunctionSelector(&_Multicall.CallOpts, functionSignature)
}

// TokenAllowances is a free data retrieval call binding the contract method 0xed583506.
//
// Solidity: function tokenAllowances(address spenderContract, address user, address[] tokens) view returns(uint256[] allowances)
func (_Multicall *MulticallCaller) TokenAllowances(opts *bind.CallOpts, spenderContract common.Address, user common.Address, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Multicall.contract.Call(opts, &out, "tokenAllowances", spenderContract, user, tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokenAllowances is a free data retrieval call binding the contract method 0xed583506.
//
// Solidity: function tokenAllowances(address spenderContract, address user, address[] tokens) view returns(uint256[] allowances)
func (_Multicall *MulticallSession) TokenAllowances(spenderContract common.Address, user common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _Multicall.Contract.TokenAllowances(&_Multicall.CallOpts, spenderContract, user, tokens)
}

// TokenAllowances is a free data retrieval call binding the contract method 0xed583506.
//
// Solidity: function tokenAllowances(address spenderContract, address user, address[] tokens) view returns(uint256[] allowances)
func (_Multicall *MulticallCallerSession) TokenAllowances(spenderContract common.Address, user common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _Multicall.Contract.TokenAllowances(&_Multicall.CallOpts, spenderContract, user, tokens)
}

// TokenBalances is a free data retrieval call binding the contract method 0x3ad206cc.
//
// Solidity: function tokenBalances(address user, address[] tokens) view returns(uint256[] balances)
func (_Multicall *MulticallCaller) TokenBalances(opts *bind.CallOpts, user common.Address, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Multicall.contract.Call(opts, &out, "tokenBalances", user, tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokenBalances is a free data retrieval call binding the contract method 0x3ad206cc.
//
// Solidity: function tokenBalances(address user, address[] tokens) view returns(uint256[] balances)
func (_Multicall *MulticallSession) TokenBalances(user common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _Multicall.Contract.TokenBalances(&_Multicall.CallOpts, user, tokens)
}

// TokenBalances is a free data retrieval call binding the contract method 0x3ad206cc.
//
// Solidity: function tokenBalances(address user, address[] tokens) view returns(uint256[] balances)
func (_Multicall *MulticallCallerSession) TokenBalances(user common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _Multicall.Contract.TokenBalances(&_Multicall.CallOpts, user, tokens)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Multicall *MulticallTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multicall.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Multicall *MulticallSession) Withdraw() (*types.Transaction, error) {
	return _Multicall.Contract.Withdraw(&_Multicall.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Multicall *MulticallTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Multicall.Contract.Withdraw(&_Multicall.TransactOpts)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(address token, uint256 amount) returns()
func (_Multicall *MulticallTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Multicall.contract.Transact(opts, "withdrawToken", token, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(address token, uint256 amount) returns()
func (_Multicall *MulticallSession) WithdrawToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Multicall.Contract.WithdrawToken(&_Multicall.TransactOpts, token, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(address token, uint256 amount) returns()
func (_Multicall *MulticallTransactorSession) WithdrawToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Multicall.Contract.WithdrawToken(&_Multicall.TransactOpts, token, amount)
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balances

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

// BalancesMetaData contains all meta data concerning the Balances contract.
var BalancesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"depositedBalances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"userFirst\",\"type\":\"bool\"}],\"name\":\"depositedBalancesGeneric\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"depositedEtherGeneric\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"functionSignature\",\"type\":\"string\"}],\"name\":\"getFunctionSelector\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spenderContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"tokenAllowances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"allowances\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"tokenBalances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BalancesABI is the input ABI used to generate the binding from.
// Deprecated: Use BalancesMetaData.ABI instead.
var BalancesABI = BalancesMetaData.ABI

// Balances is an auto generated Go binding around an Ethereum contract.
type Balances struct {
	BalancesCaller     // Read-only binding to the contract
	BalancesTransactor // Write-only binding to the contract
	BalancesFilterer   // Log filterer for contract events
}

// BalancesCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalancesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalancesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalancesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalancesSession struct {
	Contract     *Balances         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalancesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalancesCallerSession struct {
	Contract *BalancesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BalancesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalancesTransactorSession struct {
	Contract     *BalancesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BalancesRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalancesRaw struct {
	Contract *Balances // Generic contract binding to access the raw methods on
}

// BalancesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalancesCallerRaw struct {
	Contract *BalancesCaller // Generic read-only contract binding to access the raw methods on
}

// BalancesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalancesTransactorRaw struct {
	Contract *BalancesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalances creates a new instance of Balances, bound to a specific deployed contract.
func NewBalances(address common.Address, backend bind.ContractBackend) (*Balances, error) {
	contract, err := bindBalances(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Balances{BalancesCaller: BalancesCaller{contract: contract}, BalancesTransactor: BalancesTransactor{contract: contract}, BalancesFilterer: BalancesFilterer{contract: contract}}, nil
}

// NewBalancesCaller creates a new read-only instance of Balances, bound to a specific deployed contract.
func NewBalancesCaller(address common.Address, caller bind.ContractCaller) (*BalancesCaller, error) {
	contract, err := bindBalances(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalancesCaller{contract: contract}, nil
}

// NewBalancesTransactor creates a new write-only instance of Balances, bound to a specific deployed contract.
func NewBalancesTransactor(address common.Address, transactor bind.ContractTransactor) (*BalancesTransactor, error) {
	contract, err := bindBalances(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalancesTransactor{contract: contract}, nil
}

// NewBalancesFilterer creates a new log filterer instance of Balances, bound to a specific deployed contract.
func NewBalancesFilterer(address common.Address, filterer bind.ContractFilterer) (*BalancesFilterer, error) {
	contract, err := bindBalances(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalancesFilterer{contract: contract}, nil
}

// bindBalances binds a generic wrapper to an already deployed contract.
func bindBalances(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalancesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balances *BalancesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balances.Contract.BalancesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balances *BalancesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balances.Contract.BalancesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balances *BalancesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balances.Contract.BalancesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balances *BalancesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balances.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balances *BalancesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balances.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balances *BalancesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balances.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Balances *BalancesCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Balances.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Balances *BalancesSession) Admin() (common.Address, error) {
	return _Balances.Contract.Admin(&_Balances.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Balances *BalancesCallerSession) Admin() (common.Address, error) {
	return _Balances.Contract.Admin(&_Balances.CallOpts)
}

// DepositedBalances is a free data retrieval call binding the contract method 0x1621c4e5.
//
// Solidity: function depositedBalances(address exchange, address user, address[] tokens) view returns(uint256[] balances)
func (_Balances *BalancesCaller) DepositedBalances(opts *bind.CallOpts, exchange common.Address, user common.Address, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Balances.contract.Call(opts, &out, "depositedBalances", exchange, user, tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// DepositedBalances is a free data retrieval call binding the contract method 0x1621c4e5.
//
// Solidity: function depositedBalances(address exchange, address user, address[] tokens) view returns(uint256[] balances)
func (_Balances *BalancesSession) DepositedBalances(exchange common.Address, user common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _Balances.Contract.DepositedBalances(&_Balances.CallOpts, exchange, user, tokens)
}

// DepositedBalances is a free data retrieval call binding the contract method 0x1621c4e5.
//
// Solidity: function depositedBalances(address exchange, address user, address[] tokens) view returns(uint256[] balances)
func (_Balances *BalancesCallerSession) DepositedBalances(exchange common.Address, user common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _Balances.Contract.DepositedBalances(&_Balances.CallOpts, exchange, user, tokens)
}

// DepositedBalancesGeneric is a free data retrieval call binding the contract method 0x83dfb65a.
//
// Solidity: function depositedBalancesGeneric(address exchange, bytes4 selector, address user, address[] tokens, bool userFirst) view returns(uint256[] balances)
func (_Balances *BalancesCaller) DepositedBalancesGeneric(opts *bind.CallOpts, exchange common.Address, selector [4]byte, user common.Address, tokens []common.Address, userFirst bool) ([]*big.Int, error) {
	var out []interface{}
	err := _Balances.contract.Call(opts, &out, "depositedBalancesGeneric", exchange, selector, user, tokens, userFirst)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// DepositedBalancesGeneric is a free data retrieval call binding the contract method 0x83dfb65a.
//
// Solidity: function depositedBalancesGeneric(address exchange, bytes4 selector, address user, address[] tokens, bool userFirst) view returns(uint256[] balances)
func (_Balances *BalancesSession) DepositedBalancesGeneric(exchange common.Address, selector [4]byte, user common.Address, tokens []common.Address, userFirst bool) ([]*big.Int, error) {
	return _Balances.Contract.DepositedBalancesGeneric(&_Balances.CallOpts, exchange, selector, user, tokens, userFirst)
}

// DepositedBalancesGeneric is a free data retrieval call binding the contract method 0x83dfb65a.
//
// Solidity: function depositedBalancesGeneric(address exchange, bytes4 selector, address user, address[] tokens, bool userFirst) view returns(uint256[] balances)
func (_Balances *BalancesCallerSession) DepositedBalancesGeneric(exchange common.Address, selector [4]byte, user common.Address, tokens []common.Address, userFirst bool) ([]*big.Int, error) {
	return _Balances.Contract.DepositedBalancesGeneric(&_Balances.CallOpts, exchange, selector, user, tokens, userFirst)
}

// DepositedEtherGeneric is a free data retrieval call binding the contract method 0xfe8b2707.
//
// Solidity: function depositedEtherGeneric(address exchange, bytes4 selector, address user) view returns(uint256)
func (_Balances *BalancesCaller) DepositedEtherGeneric(opts *bind.CallOpts, exchange common.Address, selector [4]byte, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Balances.contract.Call(opts, &out, "depositedEtherGeneric", exchange, selector, user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositedEtherGeneric is a free data retrieval call binding the contract method 0xfe8b2707.
//
// Solidity: function depositedEtherGeneric(address exchange, bytes4 selector, address user) view returns(uint256)
func (_Balances *BalancesSession) DepositedEtherGeneric(exchange common.Address, selector [4]byte, user common.Address) (*big.Int, error) {
	return _Balances.Contract.DepositedEtherGeneric(&_Balances.CallOpts, exchange, selector, user)
}

// DepositedEtherGeneric is a free data retrieval call binding the contract method 0xfe8b2707.
//
// Solidity: function depositedEtherGeneric(address exchange, bytes4 selector, address user) view returns(uint256)
func (_Balances *BalancesCallerSession) DepositedEtherGeneric(exchange common.Address, selector [4]byte, user common.Address) (*big.Int, error) {
	return _Balances.Contract.DepositedEtherGeneric(&_Balances.CallOpts, exchange, selector, user)
}

// GetFunctionSelector is a free data retrieval call binding the contract method 0xde38c3d0.
//
// Solidity: function getFunctionSelector(string functionSignature) pure returns(bytes4)
func (_Balances *BalancesCaller) GetFunctionSelector(opts *bind.CallOpts, functionSignature string) ([4]byte, error) {
	var out []interface{}
	err := _Balances.contract.Call(opts, &out, "getFunctionSelector", functionSignature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// GetFunctionSelector is a free data retrieval call binding the contract method 0xde38c3d0.
//
// Solidity: function getFunctionSelector(string functionSignature) pure returns(bytes4)
func (_Balances *BalancesSession) GetFunctionSelector(functionSignature string) ([4]byte, error) {
	return _Balances.Contract.GetFunctionSelector(&_Balances.CallOpts, functionSignature)
}

// GetFunctionSelector is a free data retrieval call binding the contract method 0xde38c3d0.
//
// Solidity: function getFunctionSelector(string functionSignature) pure returns(bytes4)
func (_Balances *BalancesCallerSession) GetFunctionSelector(functionSignature string) ([4]byte, error) {
	return _Balances.Contract.GetFunctionSelector(&_Balances.CallOpts, functionSignature)
}

// TokenAllowances is a free data retrieval call binding the contract method 0xed583506.
//
// Solidity: function tokenAllowances(address spenderContract, address user, address[] tokens) view returns(uint256[] allowances)
func (_Balances *BalancesCaller) TokenAllowances(opts *bind.CallOpts, spenderContract common.Address, user common.Address, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Balances.contract.Call(opts, &out, "tokenAllowances", spenderContract, user, tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokenAllowances is a free data retrieval call binding the contract method 0xed583506.
//
// Solidity: function tokenAllowances(address spenderContract, address user, address[] tokens) view returns(uint256[] allowances)
func (_Balances *BalancesSession) TokenAllowances(spenderContract common.Address, user common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _Balances.Contract.TokenAllowances(&_Balances.CallOpts, spenderContract, user, tokens)
}

// TokenAllowances is a free data retrieval call binding the contract method 0xed583506.
//
// Solidity: function tokenAllowances(address spenderContract, address user, address[] tokens) view returns(uint256[] allowances)
func (_Balances *BalancesCallerSession) TokenAllowances(spenderContract common.Address, user common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _Balances.Contract.TokenAllowances(&_Balances.CallOpts, spenderContract, user, tokens)
}

// TokenBalances is a free data retrieval call binding the contract method 0x3ad206cc.
//
// Solidity: function tokenBalances(address user, address[] tokens) view returns(uint256[] balances)
func (_Balances *BalancesCaller) TokenBalances(opts *bind.CallOpts, user common.Address, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Balances.contract.Call(opts, &out, "tokenBalances", user, tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokenBalances is a free data retrieval call binding the contract method 0x3ad206cc.
//
// Solidity: function tokenBalances(address user, address[] tokens) view returns(uint256[] balances)
func (_Balances *BalancesSession) TokenBalances(user common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _Balances.Contract.TokenBalances(&_Balances.CallOpts, user, tokens)
}

// TokenBalances is a free data retrieval call binding the contract method 0x3ad206cc.
//
// Solidity: function tokenBalances(address user, address[] tokens) view returns(uint256[] balances)
func (_Balances *BalancesCallerSession) TokenBalances(user common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _Balances.Contract.TokenBalances(&_Balances.CallOpts, user, tokens)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Balances *BalancesTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balances.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Balances *BalancesSession) Withdraw() (*types.Transaction, error) {
	return _Balances.Contract.Withdraw(&_Balances.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Balances *BalancesTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Balances.Contract.Withdraw(&_Balances.TransactOpts)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(address token, uint256 amount) returns()
func (_Balances *BalancesTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Balances.contract.Transact(opts, "withdrawToken", token, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(address token, uint256 amount) returns()
func (_Balances *BalancesSession) WithdrawToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Balances.Contract.WithdrawToken(&_Balances.TransactOpts, token, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(address token, uint256 amount) returns()
func (_Balances *BalancesTransactorSession) WithdrawToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Balances.Contract.WithdrawToken(&_Balances.TransactOpts, token, amount)
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// DePocketBridgeMetaData contains all meta data concerning the DePocketBridge contract.
var DePocketBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"userAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenAddresses\",\"type\":\"address[]\"}],\"name\":\"allBalancesForManyAccounts\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"tokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokenAddresses\",\"type\":\"address[]\"}],\"name\":\"walletBalances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"01afd5f3": "allBalancesForManyAccounts(address[],address[])",
		"1049334f": "tokenBalance(address,address)",
		"77a7d968": "walletBalances(address,address[])",
	},
	Bin: "0x608060405234801561001057600080fd5b50610790806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806301afd5f3146100465780631049334f1461006f57806377a7d96814610090575b600080fd5b6100596100543660046105c2565b6100a3565b6040516100669190610682565b60405180910390f35b61008261007d36600461053c565b61025d565b604051908152602001610066565b61005961009e36600461056f565b61039f565b606060006100b185846106de565b67ffffffffffffffff8111156100c9576100c9610744565b6040519080825280602002602001820160405280156100f2578160200160208202803683370190505b50905060005b858110156102535760005b8481101561024057600086868381811061011f5761011f61072e565b9050602002016020810190610134919061051a565b6001600160a01b0316146101ca576101938888848181106101575761015761072e565b905060200201602081019061016c919061051a565b87878481811061017e5761017e61072e565b905060200201602081019061007d919061051a565b838261019f88866106de565b6101a991906106c6565b815181106101b9576101b961072e565b60200260200101818152505061022e565b8787838181106101dc576101dc61072e565b90506020020160208101906101f1919061051a565b6001600160a01b031631838261020788866106de565b61021191906106c6565b815181106102215761022161072e565b6020026020010181815250505b80610238816106fd565b915050610103565b508061024b816106fd565b9150506100f8565b5095945050505050565b6000813b1561039557604080516001600160a01b0385811660248084019190915283518084039091018152604490920183526020820180516001600160e01b03166370a0823160e01b17905291518492600092908416916102be9190610647565b600060405180830381855afa9150503d80600081146102f9576040519150601f19603f3d011682016040523d82523d6000602084013e6102fe565b606091505b50509050801561038a576040516370a0823160e01b81526001600160a01b0386811660048301528316906370a082319060240160206040518083038186803b15801561034957600080fd5b505afa15801561035d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610381919061062e565b92505050610399565b600092505050610399565b5060005b92915050565b6060816103ab57600080fd5b60008267ffffffffffffffff8111156103c6576103c6610744565b6040519080825280602002602001820160405280156103ef578160200160208202803683370190505b50905060005b838110156104a95760008585838181106104115761041161072e565b9050602002016020810190610426919061051a565b6001600160a01b03161461046d5761044a8686868481811061017e5761017e61072e565b82828151811061045c5761045c61072e565b602002602001018181525050610497565b856001600160a01b03163182828151811061048a5761048a61072e565b6020026020010181815250505b806104a1816106fd565b9150506103f5565b50949350505050565b80356001600160a01b03811681146104c957600080fd5b919050565b60008083601f8401126104e057600080fd5b50813567ffffffffffffffff8111156104f857600080fd5b6020830191508360208260051b850101111561051357600080fd5b9250929050565b60006020828403121561052c57600080fd5b610535826104b2565b9392505050565b6000806040838503121561054f57600080fd5b610558836104b2565b9150610566602084016104b2565b90509250929050565b60008060006040848603121561058457600080fd5b61058d846104b2565b9250602084013567ffffffffffffffff8111156105a957600080fd5b6105b5868287016104ce565b9497909650939450505050565b600080600080604085870312156105d857600080fd5b843567ffffffffffffffff808211156105f057600080fd5b6105fc888389016104ce565b9096509450602087013591508082111561061557600080fd5b50610622878288016104ce565b95989497509550505050565b60006020828403121561064057600080fd5b5051919050565b6000825160005b81811015610668576020818601810151858301520161064e565b81811115610677576000828501525b509190910192915050565b6020808252825182820181905260009190848201906040850190845b818110156106ba5783518352928401929184019160010161069e565b50909695505050505050565b600082198211156106d9576106d9610718565b500190565b60008160001904831182151516156106f8576106f8610718565b500290565b600060001982141561071157610711610718565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fdfea2646970667358221220ab5e2a69953f2ec98792c2f5c45128e573d0eba753fed1ee3fa422aafcad48e664736f6c63430008070033",
}

// DePocketBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use DePocketBridgeMetaData.ABI instead.
var DePocketBridgeABI = DePocketBridgeMetaData.ABI

// Deprecated: Use DePocketBridgeMetaData.Sigs instead.
// DePocketBridgeFuncSigs maps the 4-byte function signature to its string representation.
var DePocketBridgeFuncSigs = DePocketBridgeMetaData.Sigs

// DePocketBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DePocketBridgeMetaData.Bin instead.
var DePocketBridgeBin = DePocketBridgeMetaData.Bin

// DeployDePocketBridge deploys a new Ethereum contract, binding an instance of DePocketBridge to it.
func DeployDePocketBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DePocketBridge, error) {
	parsed, err := DePocketBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DePocketBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DePocketBridge{DePocketBridgeCaller: DePocketBridgeCaller{contract: contract}, DePocketBridgeTransactor: DePocketBridgeTransactor{contract: contract}, DePocketBridgeFilterer: DePocketBridgeFilterer{contract: contract}}, nil
}

// DePocketBridge is an auto generated Go binding around an Ethereum contract.
type DePocketBridge struct {
	DePocketBridgeCaller     // Read-only binding to the contract
	DePocketBridgeTransactor // Write-only binding to the contract
	DePocketBridgeFilterer   // Log filterer for contract events
}

// DePocketBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type DePocketBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DePocketBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DePocketBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DePocketBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DePocketBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DePocketBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DePocketBridgeSession struct {
	Contract     *DePocketBridge   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DePocketBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DePocketBridgeCallerSession struct {
	Contract *DePocketBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DePocketBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DePocketBridgeTransactorSession struct {
	Contract     *DePocketBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DePocketBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type DePocketBridgeRaw struct {
	Contract *DePocketBridge // Generic contract binding to access the raw methods on
}

// DePocketBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DePocketBridgeCallerRaw struct {
	Contract *DePocketBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// DePocketBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DePocketBridgeTransactorRaw struct {
	Contract *DePocketBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDePocketBridge creates a new instance of DePocketBridge, bound to a specific deployed contract.
func NewDePocketBridge(address common.Address, backend bind.ContractBackend) (*DePocketBridge, error) {
	contract, err := bindDePocketBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DePocketBridge{DePocketBridgeCaller: DePocketBridgeCaller{contract: contract}, DePocketBridgeTransactor: DePocketBridgeTransactor{contract: contract}, DePocketBridgeFilterer: DePocketBridgeFilterer{contract: contract}}, nil
}

// NewDePocketBridgeCaller creates a new read-only instance of DePocketBridge, bound to a specific deployed contract.
func NewDePocketBridgeCaller(address common.Address, caller bind.ContractCaller) (*DePocketBridgeCaller, error) {
	contract, err := bindDePocketBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DePocketBridgeCaller{contract: contract}, nil
}

// NewDePocketBridgeTransactor creates a new write-only instance of DePocketBridge, bound to a specific deployed contract.
func NewDePocketBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*DePocketBridgeTransactor, error) {
	contract, err := bindDePocketBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DePocketBridgeTransactor{contract: contract}, nil
}

// NewDePocketBridgeFilterer creates a new log filterer instance of DePocketBridge, bound to a specific deployed contract.
func NewDePocketBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*DePocketBridgeFilterer, error) {
	contract, err := bindDePocketBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DePocketBridgeFilterer{contract: contract}, nil
}

// bindDePocketBridge binds a generic wrapper to an already deployed contract.
func bindDePocketBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DePocketBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DePocketBridge *DePocketBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DePocketBridge.Contract.DePocketBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DePocketBridge *DePocketBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DePocketBridge.Contract.DePocketBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DePocketBridge *DePocketBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DePocketBridge.Contract.DePocketBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DePocketBridge *DePocketBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DePocketBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DePocketBridge *DePocketBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DePocketBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DePocketBridge *DePocketBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DePocketBridge.Contract.contract.Transact(opts, method, params...)
}

// AllBalancesForManyAccounts is a free data retrieval call binding the contract method 0x01afd5f3.
//
// Solidity: function allBalancesForManyAccounts(address[] userAddresses, address[] tokenAddresses) view returns(uint256[])
func (_DePocketBridge *DePocketBridgeCaller) AllBalancesForManyAccounts(opts *bind.CallOpts, userAddresses []common.Address, tokenAddresses []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _DePocketBridge.contract.Call(opts, &out, "allBalancesForManyAccounts", userAddresses, tokenAddresses)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllBalancesForManyAccounts is a free data retrieval call binding the contract method 0x01afd5f3.
//
// Solidity: function allBalancesForManyAccounts(address[] userAddresses, address[] tokenAddresses) view returns(uint256[])
func (_DePocketBridge *DePocketBridgeSession) AllBalancesForManyAccounts(userAddresses []common.Address, tokenAddresses []common.Address) ([]*big.Int, error) {
	return _DePocketBridge.Contract.AllBalancesForManyAccounts(&_DePocketBridge.CallOpts, userAddresses, tokenAddresses)
}

// AllBalancesForManyAccounts is a free data retrieval call binding the contract method 0x01afd5f3.
//
// Solidity: function allBalancesForManyAccounts(address[] userAddresses, address[] tokenAddresses) view returns(uint256[])
func (_DePocketBridge *DePocketBridgeCallerSession) AllBalancesForManyAccounts(userAddresses []common.Address, tokenAddresses []common.Address) ([]*big.Int, error) {
	return _DePocketBridge.Contract.AllBalancesForManyAccounts(&_DePocketBridge.CallOpts, userAddresses, tokenAddresses)
}

// TokenBalance is a free data retrieval call binding the contract method 0x1049334f.
//
// Solidity: function tokenBalance(address userAddress, address tokenAddress) view returns(uint256)
func (_DePocketBridge *DePocketBridgeCaller) TokenBalance(opts *bind.CallOpts, userAddress common.Address, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DePocketBridge.contract.Call(opts, &out, "tokenBalance", userAddress, tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenBalance is a free data retrieval call binding the contract method 0x1049334f.
//
// Solidity: function tokenBalance(address userAddress, address tokenAddress) view returns(uint256)
func (_DePocketBridge *DePocketBridgeSession) TokenBalance(userAddress common.Address, tokenAddress common.Address) (*big.Int, error) {
	return _DePocketBridge.Contract.TokenBalance(&_DePocketBridge.CallOpts, userAddress, tokenAddress)
}

// TokenBalance is a free data retrieval call binding the contract method 0x1049334f.
//
// Solidity: function tokenBalance(address userAddress, address tokenAddress) view returns(uint256)
func (_DePocketBridge *DePocketBridgeCallerSession) TokenBalance(userAddress common.Address, tokenAddress common.Address) (*big.Int, error) {
	return _DePocketBridge.Contract.TokenBalance(&_DePocketBridge.CallOpts, userAddress, tokenAddress)
}

// WalletBalances is a free data retrieval call binding the contract method 0x77a7d968.
//
// Solidity: function walletBalances(address userAddress, address[] tokenAddresses) view returns(uint256[])
func (_DePocketBridge *DePocketBridgeCaller) WalletBalances(opts *bind.CallOpts, userAddress common.Address, tokenAddresses []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _DePocketBridge.contract.Call(opts, &out, "walletBalances", userAddress, tokenAddresses)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// WalletBalances is a free data retrieval call binding the contract method 0x77a7d968.
//
// Solidity: function walletBalances(address userAddress, address[] tokenAddresses) view returns(uint256[])
func (_DePocketBridge *DePocketBridgeSession) WalletBalances(userAddress common.Address, tokenAddresses []common.Address) ([]*big.Int, error) {
	return _DePocketBridge.Contract.WalletBalances(&_DePocketBridge.CallOpts, userAddress, tokenAddresses)
}

// WalletBalances is a free data retrieval call binding the contract method 0x77a7d968.
//
// Solidity: function walletBalances(address userAddress, address[] tokenAddresses) view returns(uint256[])
func (_DePocketBridge *DePocketBridgeCallerSession) WalletBalances(userAddress common.Address, tokenAddresses []common.Address) ([]*big.Int, error) {
	return _DePocketBridge.Contract.WalletBalances(&_DePocketBridge.CallOpts, userAddress, tokenAddresses)
}

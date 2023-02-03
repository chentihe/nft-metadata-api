// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package clonex

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

// ClonexMetaData contains all meta data concerning the Clonex contract.
var ClonexMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fileId\",\"type\":\"string\"}],\"name\":\"CloneXRevealed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_tokenUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mintTransfer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"secureBaseUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"setMintvialAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"setRandomizerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ClonexABI is the input ABI used to generate the binding from.
// Deprecated: Use ClonexMetaData.ABI instead.
var ClonexABI = ClonexMetaData.ABI

// Clonex is an auto generated Go binding around an Ethereum contract.
type Clonex struct {
	ClonexCaller     // Read-only binding to the contract
	ClonexTransactor // Write-only binding to the contract
	ClonexFilterer   // Log filterer for contract events
}

// ClonexCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClonexCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClonexTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClonexTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClonexFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClonexFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClonexSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClonexSession struct {
	Contract     *Clonex           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClonexCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClonexCallerSession struct {
	Contract *ClonexCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ClonexTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClonexTransactorSession struct {
	Contract     *ClonexTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClonexRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClonexRaw struct {
	Contract *Clonex // Generic contract binding to access the raw methods on
}

// ClonexCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClonexCallerRaw struct {
	Contract *ClonexCaller // Generic read-only contract binding to access the raw methods on
}

// ClonexTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClonexTransactorRaw struct {
	Contract *ClonexTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClonex creates a new instance of Clonex, bound to a specific deployed contract.
func NewClonex(address common.Address, backend bind.ContractBackend) (*Clonex, error) {
	contract, err := bindClonex(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Clonex{ClonexCaller: ClonexCaller{contract: contract}, ClonexTransactor: ClonexTransactor{contract: contract}, ClonexFilterer: ClonexFilterer{contract: contract}}, nil
}

// NewClonexCaller creates a new read-only instance of Clonex, bound to a specific deployed contract.
func NewClonexCaller(address common.Address, caller bind.ContractCaller) (*ClonexCaller, error) {
	contract, err := bindClonex(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClonexCaller{contract: contract}, nil
}

// NewClonexTransactor creates a new write-only instance of Clonex, bound to a specific deployed contract.
func NewClonexTransactor(address common.Address, transactor bind.ContractTransactor) (*ClonexTransactor, error) {
	contract, err := bindClonex(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClonexTransactor{contract: contract}, nil
}

// NewClonexFilterer creates a new log filterer instance of Clonex, bound to a specific deployed contract.
func NewClonexFilterer(address common.Address, filterer bind.ContractFilterer) (*ClonexFilterer, error) {
	contract, err := bindClonex(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClonexFilterer{contract: contract}, nil
}

// bindClonex binds a generic wrapper to an already deployed contract.
func bindClonex(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClonexABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clonex *ClonexRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Clonex.Contract.ClonexCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clonex *ClonexRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clonex.Contract.ClonexTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clonex *ClonexRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clonex.Contract.ClonexTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clonex *ClonexCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Clonex.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clonex *ClonexTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clonex.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clonex *ClonexTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clonex.Contract.contract.Transact(opts, method, params...)
}

// TokenUri is a free data retrieval call binding the contract method 0xdf6552b9.
//
// Solidity: function _tokenUri() view returns(string)
func (_Clonex *ClonexCaller) TokenUri(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Clonex.contract.Call(opts, &out, "_tokenUri")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenUri is a free data retrieval call binding the contract method 0xdf6552b9.
//
// Solidity: function _tokenUri() view returns(string)
func (_Clonex *ClonexSession) TokenUri() (string, error) {
	return _Clonex.Contract.TokenUri(&_Clonex.CallOpts)
}

// TokenUri is a free data retrieval call binding the contract method 0xdf6552b9.
//
// Solidity: function _tokenUri() view returns(string)
func (_Clonex *ClonexCallerSession) TokenUri() (string, error) {
	return _Clonex.Contract.TokenUri(&_Clonex.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Clonex *ClonexCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Clonex.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Clonex *ClonexSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Clonex.Contract.BalanceOf(&_Clonex.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Clonex *ClonexCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Clonex.Contract.BalanceOf(&_Clonex.CallOpts, owner)
}

// ContractLocked is a free data retrieval call binding the contract method 0x324cb3cb.
//
// Solidity: function contractLocked() view returns(bool)
func (_Clonex *ClonexCaller) ContractLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Clonex.contract.Call(opts, &out, "contractLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractLocked is a free data retrieval call binding the contract method 0x324cb3cb.
//
// Solidity: function contractLocked() view returns(bool)
func (_Clonex *ClonexSession) ContractLocked() (bool, error) {
	return _Clonex.Contract.ContractLocked(&_Clonex.CallOpts)
}

// ContractLocked is a free data retrieval call binding the contract method 0x324cb3cb.
//
// Solidity: function contractLocked() view returns(bool)
func (_Clonex *ClonexCallerSession) ContractLocked() (bool, error) {
	return _Clonex.Contract.ContractLocked(&_Clonex.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Clonex *ClonexCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Clonex.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Clonex *ClonexSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Clonex.Contract.GetApproved(&_Clonex.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Clonex *ClonexCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Clonex.Contract.GetApproved(&_Clonex.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Clonex *ClonexCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Clonex.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Clonex *ClonexSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Clonex.Contract.IsApprovedForAll(&_Clonex.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Clonex *ClonexCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Clonex.Contract.IsApprovedForAll(&_Clonex.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Clonex *ClonexCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Clonex.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Clonex *ClonexSession) Name() (string, error) {
	return _Clonex.Contract.Name(&_Clonex.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Clonex *ClonexCallerSession) Name() (string, error) {
	return _Clonex.Contract.Name(&_Clonex.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Clonex *ClonexCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Clonex.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Clonex *ClonexSession) Owner() (common.Address, error) {
	return _Clonex.Contract.Owner(&_Clonex.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Clonex *ClonexCallerSession) Owner() (common.Address, error) {
	return _Clonex.Contract.Owner(&_Clonex.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Clonex *ClonexCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Clonex.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Clonex *ClonexSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Clonex.Contract.OwnerOf(&_Clonex.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Clonex *ClonexCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Clonex.Contract.OwnerOf(&_Clonex.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Clonex *ClonexCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Clonex.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Clonex *ClonexSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Clonex.Contract.SupportsInterface(&_Clonex.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Clonex *ClonexCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Clonex.Contract.SupportsInterface(&_Clonex.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Clonex *ClonexCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Clonex.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Clonex *ClonexSession) Symbol() (string, error) {
	return _Clonex.Contract.Symbol(&_Clonex.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Clonex *ClonexCallerSession) Symbol() (string, error) {
	return _Clonex.Contract.Symbol(&_Clonex.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Clonex *ClonexCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Clonex.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Clonex *ClonexSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Clonex.Contract.TokenByIndex(&_Clonex.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Clonex *ClonexCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Clonex.Contract.TokenByIndex(&_Clonex.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Clonex *ClonexCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Clonex.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Clonex *ClonexSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Clonex.Contract.TokenOfOwnerByIndex(&_Clonex.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Clonex *ClonexCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Clonex.Contract.TokenOfOwnerByIndex(&_Clonex.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Clonex *ClonexCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Clonex.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Clonex *ClonexSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Clonex.Contract.TokenURI(&_Clonex.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Clonex *ClonexCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Clonex.Contract.TokenURI(&_Clonex.CallOpts, tokenId)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[])
func (_Clonex *ClonexCaller) TokensOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Clonex.contract.Call(opts, &out, "tokensOfOwner", _owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[])
func (_Clonex *ClonexSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _Clonex.Contract.TokensOfOwner(&_Clonex.CallOpts, _owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[])
func (_Clonex *ClonexCallerSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _Clonex.Contract.TokensOfOwner(&_Clonex.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Clonex *ClonexCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clonex.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Clonex *ClonexSession) TotalSupply() (*big.Int, error) {
	return _Clonex.Contract.TotalSupply(&_Clonex.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Clonex *ClonexCallerSession) TotalSupply() (*big.Int, error) {
	return _Clonex.Contract.TotalSupply(&_Clonex.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Clonex *ClonexTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Clonex.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Clonex *ClonexSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Clonex.Contract.Approve(&_Clonex.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Clonex *ClonexTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Clonex.Contract.Approve(&_Clonex.TransactOpts, to, tokenId)
}

// LockContract is a paid mutator transaction binding the contract method 0x753868e3.
//
// Solidity: function lockContract() returns()
func (_Clonex *ClonexTransactor) LockContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clonex.contract.Transact(opts, "lockContract")
}

// LockContract is a paid mutator transaction binding the contract method 0x753868e3.
//
// Solidity: function lockContract() returns()
func (_Clonex *ClonexSession) LockContract() (*types.Transaction, error) {
	return _Clonex.Contract.LockContract(&_Clonex.TransactOpts)
}

// LockContract is a paid mutator transaction binding the contract method 0x753868e3.
//
// Solidity: function lockContract() returns()
func (_Clonex *ClonexTransactorSession) LockContract() (*types.Transaction, error) {
	return _Clonex.Contract.LockContract(&_Clonex.TransactOpts)
}

// MintTransfer is a paid mutator transaction binding the contract method 0x937f2608.
//
// Solidity: function mintTransfer(address to) returns(uint256)
func (_Clonex *ClonexTransactor) MintTransfer(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Clonex.contract.Transact(opts, "mintTransfer", to)
}

// MintTransfer is a paid mutator transaction binding the contract method 0x937f2608.
//
// Solidity: function mintTransfer(address to) returns(uint256)
func (_Clonex *ClonexSession) MintTransfer(to common.Address) (*types.Transaction, error) {
	return _Clonex.Contract.MintTransfer(&_Clonex.TransactOpts, to)
}

// MintTransfer is a paid mutator transaction binding the contract method 0x937f2608.
//
// Solidity: function mintTransfer(address to) returns(uint256)
func (_Clonex *ClonexTransactorSession) MintTransfer(to common.Address) (*types.Transaction, error) {
	return _Clonex.Contract.MintTransfer(&_Clonex.TransactOpts, to)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Clonex *ClonexTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clonex.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Clonex *ClonexSession) RenounceOwnership() (*types.Transaction, error) {
	return _Clonex.Contract.RenounceOwnership(&_Clonex.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Clonex *ClonexTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Clonex.Contract.RenounceOwnership(&_Clonex.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Clonex *ClonexTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Clonex.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Clonex *ClonexSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Clonex.Contract.SafeTransferFrom(&_Clonex.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Clonex *ClonexTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Clonex.Contract.SafeTransferFrom(&_Clonex.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Clonex *ClonexTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Clonex.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Clonex *ClonexSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Clonex.Contract.SafeTransferFrom0(&_Clonex.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Clonex *ClonexTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Clonex.Contract.SafeTransferFrom0(&_Clonex.TransactOpts, from, to, tokenId, _data)
}

// SecureBaseUri is a paid mutator transaction binding the contract method 0xc46b0d84.
//
// Solidity: function secureBaseUri(string newUri) returns()
func (_Clonex *ClonexTransactor) SecureBaseUri(opts *bind.TransactOpts, newUri string) (*types.Transaction, error) {
	return _Clonex.contract.Transact(opts, "secureBaseUri", newUri)
}

// SecureBaseUri is a paid mutator transaction binding the contract method 0xc46b0d84.
//
// Solidity: function secureBaseUri(string newUri) returns()
func (_Clonex *ClonexSession) SecureBaseUri(newUri string) (*types.Transaction, error) {
	return _Clonex.Contract.SecureBaseUri(&_Clonex.TransactOpts, newUri)
}

// SecureBaseUri is a paid mutator transaction binding the contract method 0xc46b0d84.
//
// Solidity: function secureBaseUri(string newUri) returns()
func (_Clonex *ClonexTransactorSession) SecureBaseUri(newUri string) (*types.Transaction, error) {
	return _Clonex.Contract.SecureBaseUri(&_Clonex.TransactOpts, newUri)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Clonex *ClonexTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Clonex.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Clonex *ClonexSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Clonex.Contract.SetApprovalForAll(&_Clonex.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Clonex *ClonexTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Clonex.Contract.SetApprovalForAll(&_Clonex.TransactOpts, operator, approved)
}

// SetMintvialAddress is a paid mutator transaction binding the contract method 0xad6c9962.
//
// Solidity: function setMintvialAddress(address newAddress) returns()
func (_Clonex *ClonexTransactor) SetMintvialAddress(opts *bind.TransactOpts, newAddress common.Address) (*types.Transaction, error) {
	return _Clonex.contract.Transact(opts, "setMintvialAddress", newAddress)
}

// SetMintvialAddress is a paid mutator transaction binding the contract method 0xad6c9962.
//
// Solidity: function setMintvialAddress(address newAddress) returns()
func (_Clonex *ClonexSession) SetMintvialAddress(newAddress common.Address) (*types.Transaction, error) {
	return _Clonex.Contract.SetMintvialAddress(&_Clonex.TransactOpts, newAddress)
}

// SetMintvialAddress is a paid mutator transaction binding the contract method 0xad6c9962.
//
// Solidity: function setMintvialAddress(address newAddress) returns()
func (_Clonex *ClonexTransactorSession) SetMintvialAddress(newAddress common.Address) (*types.Transaction, error) {
	return _Clonex.Contract.SetMintvialAddress(&_Clonex.TransactOpts, newAddress)
}

// SetRandomizerAddress is a paid mutator transaction binding the contract method 0xf1abd7de.
//
// Solidity: function setRandomizerAddress(address newAddress) returns()
func (_Clonex *ClonexTransactor) SetRandomizerAddress(opts *bind.TransactOpts, newAddress common.Address) (*types.Transaction, error) {
	return _Clonex.contract.Transact(opts, "setRandomizerAddress", newAddress)
}

// SetRandomizerAddress is a paid mutator transaction binding the contract method 0xf1abd7de.
//
// Solidity: function setRandomizerAddress(address newAddress) returns()
func (_Clonex *ClonexSession) SetRandomizerAddress(newAddress common.Address) (*types.Transaction, error) {
	return _Clonex.Contract.SetRandomizerAddress(&_Clonex.TransactOpts, newAddress)
}

// SetRandomizerAddress is a paid mutator transaction binding the contract method 0xf1abd7de.
//
// Solidity: function setRandomizerAddress(address newAddress) returns()
func (_Clonex *ClonexTransactorSession) SetRandomizerAddress(newAddress common.Address) (*types.Transaction, error) {
	return _Clonex.Contract.SetRandomizerAddress(&_Clonex.TransactOpts, newAddress)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Clonex *ClonexTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Clonex.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Clonex *ClonexSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Clonex.Contract.TransferFrom(&_Clonex.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Clonex *ClonexTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Clonex.Contract.TransferFrom(&_Clonex.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Clonex *ClonexTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Clonex.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Clonex *ClonexSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Clonex.Contract.TransferOwnership(&_Clonex.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Clonex *ClonexTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Clonex.Contract.TransferOwnership(&_Clonex.TransactOpts, newOwner)
}

// ClonexApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Clonex contract.
type ClonexApprovalIterator struct {
	Event *ClonexApproval // Event containing the contract specifics and raw log

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
func (it *ClonexApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClonexApproval)
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
		it.Event = new(ClonexApproval)
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
func (it *ClonexApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClonexApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClonexApproval represents a Approval event raised by the Clonex contract.
type ClonexApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Clonex *ClonexFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ClonexApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Clonex.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ClonexApprovalIterator{contract: _Clonex.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Clonex *ClonexFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ClonexApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Clonex.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClonexApproval)
				if err := _Clonex.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Clonex *ClonexFilterer) ParseApproval(log types.Log) (*ClonexApproval, error) {
	event := new(ClonexApproval)
	if err := _Clonex.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClonexApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Clonex contract.
type ClonexApprovalForAllIterator struct {
	Event *ClonexApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ClonexApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClonexApprovalForAll)
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
		it.Event = new(ClonexApprovalForAll)
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
func (it *ClonexApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClonexApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClonexApprovalForAll represents a ApprovalForAll event raised by the Clonex contract.
type ClonexApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Clonex *ClonexFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ClonexApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Clonex.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ClonexApprovalForAllIterator{contract: _Clonex.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Clonex *ClonexFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ClonexApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Clonex.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClonexApprovalForAll)
				if err := _Clonex.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Clonex *ClonexFilterer) ParseApprovalForAll(log types.Log) (*ClonexApprovalForAll, error) {
	event := new(ClonexApprovalForAll)
	if err := _Clonex.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClonexCloneXRevealedIterator is returned from FilterCloneXRevealed and is used to iterate over the raw logs and unpacked data for CloneXRevealed events raised by the Clonex contract.
type ClonexCloneXRevealedIterator struct {
	Event *ClonexCloneXRevealed // Event containing the contract specifics and raw log

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
func (it *ClonexCloneXRevealedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClonexCloneXRevealed)
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
		it.Event = new(ClonexCloneXRevealed)
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
func (it *ClonexCloneXRevealedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClonexCloneXRevealedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClonexCloneXRevealed represents a CloneXRevealed event raised by the Clonex contract.
type ClonexCloneXRevealed struct {
	TokenId *big.Int
	FileId  string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCloneXRevealed is a free log retrieval operation binding the contract event 0x2ce54db2c0bc64dc675f5fae90636ec2f0c88dbd8e7c6a19c9caca9193741b15.
//
// Solidity: event CloneXRevealed(uint256 tokenId, string fileId)
func (_Clonex *ClonexFilterer) FilterCloneXRevealed(opts *bind.FilterOpts) (*ClonexCloneXRevealedIterator, error) {

	logs, sub, err := _Clonex.contract.FilterLogs(opts, "CloneXRevealed")
	if err != nil {
		return nil, err
	}
	return &ClonexCloneXRevealedIterator{contract: _Clonex.contract, event: "CloneXRevealed", logs: logs, sub: sub}, nil
}

// WatchCloneXRevealed is a free log subscription operation binding the contract event 0x2ce54db2c0bc64dc675f5fae90636ec2f0c88dbd8e7c6a19c9caca9193741b15.
//
// Solidity: event CloneXRevealed(uint256 tokenId, string fileId)
func (_Clonex *ClonexFilterer) WatchCloneXRevealed(opts *bind.WatchOpts, sink chan<- *ClonexCloneXRevealed) (event.Subscription, error) {

	logs, sub, err := _Clonex.contract.WatchLogs(opts, "CloneXRevealed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClonexCloneXRevealed)
				if err := _Clonex.contract.UnpackLog(event, "CloneXRevealed", log); err != nil {
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

// ParseCloneXRevealed is a log parse operation binding the contract event 0x2ce54db2c0bc64dc675f5fae90636ec2f0c88dbd8e7c6a19c9caca9193741b15.
//
// Solidity: event CloneXRevealed(uint256 tokenId, string fileId)
func (_Clonex *ClonexFilterer) ParseCloneXRevealed(log types.Log) (*ClonexCloneXRevealed, error) {
	event := new(ClonexCloneXRevealed)
	if err := _Clonex.contract.UnpackLog(event, "CloneXRevealed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClonexOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Clonex contract.
type ClonexOwnershipTransferredIterator struct {
	Event *ClonexOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ClonexOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClonexOwnershipTransferred)
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
		it.Event = new(ClonexOwnershipTransferred)
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
func (it *ClonexOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClonexOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClonexOwnershipTransferred represents a OwnershipTransferred event raised by the Clonex contract.
type ClonexOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Clonex *ClonexFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ClonexOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Clonex.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ClonexOwnershipTransferredIterator{contract: _Clonex.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Clonex *ClonexFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ClonexOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Clonex.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClonexOwnershipTransferred)
				if err := _Clonex.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Clonex *ClonexFilterer) ParseOwnershipTransferred(log types.Log) (*ClonexOwnershipTransferred, error) {
	event := new(ClonexOwnershipTransferred)
	if err := _Clonex.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClonexTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Clonex contract.
type ClonexTransferIterator struct {
	Event *ClonexTransfer // Event containing the contract specifics and raw log

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
func (it *ClonexTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClonexTransfer)
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
		it.Event = new(ClonexTransfer)
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
func (it *ClonexTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClonexTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClonexTransfer represents a Transfer event raised by the Clonex contract.
type ClonexTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Clonex *ClonexFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ClonexTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Clonex.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ClonexTransferIterator{contract: _Clonex.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Clonex *ClonexFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ClonexTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Clonex.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClonexTransfer)
				if err := _Clonex.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Clonex *ClonexFilterer) ParseTransfer(log types.Log) (*ClonexTransfer, error) {
	event := new(ClonexTransfer)
	if err := _Clonex.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

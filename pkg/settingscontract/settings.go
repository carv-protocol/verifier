// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package settingscontract

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

// ISettingsSettingParams is an auto generated low-level Go binding around an user-defined struct.
type ISettingsSettingParams struct {
	MaxVrfActiveNodes               *big.Int
	NodeMinOnlineDuration           *big.Int
	NodeVerifyDuration              *big.Int
	NodeSlashReward                 *big.Int
	MinTeeStakeAmount               *big.Int
	TeeSlashAmount                  *big.Int
	TeeUnstakeDuration              *big.Int
	MinCommissionRateModifyInterval *big.Int
	NodeMaxMissVerifyCount          uint64
	MaxNodeWeights                  uint16
}

// SettingscontractMetaData contains all meta data concerning the Settingscontract contract.
var SettingscontractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxVrfActiveNodes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeMinOnlineDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeVerifyDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeSlashReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTeeStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"teeSlashAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"teeUnstakeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minCommissionRateModifyInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nodeMaxMissVerifyCount\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"maxNodeWeights\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structISettings.SettingParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"UpdateSettings\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"maxNodeWeights\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxVrfActiveNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minCommissionRateModifyInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minTeeStakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeMaxMissVerifyCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeMinOnlineDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeSlashReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeVerifyDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeSlashAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeUnstakeDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxVrfActiveNodes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeMinOnlineDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeVerifyDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeSlashReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTeeStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"teeSlashAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"teeUnstakeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minCommissionRateModifyInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nodeMaxMissVerifyCount\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"maxNodeWeights\",\"type\":\"uint16\"}],\"internalType\":\"structISettings.SettingParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"updateSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50338061003757604051631e4fbdf760e01b81526000600482015260240160405180910390fd5b61004081610046565b50610096565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b61052a806100a56000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c806373cf675e1161008c578063b416dd8511610066578063b416dd85146101b7578063caa9eeba146101c0578063ee6b2820146101c9578063f2fde38b146101dc57600080fd5b806373cf675e1461018a57806388cef9c2146101935780638da5cb5b1461019c57600080fd5b8063512ab483116100c8578063512ab48314610165578063600d09491461016e578063715018a614610177578063729fe45f1461018157600080fd5b80630b097ac3146100ef57806315fbbc9d146101215780633c6174bc1461014e575b600080fd5b6009546101039067ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020015b60405180910390f35b60095461013b9068010000000000000000900461ffff1681565b60405161ffff9091168152602001610118565b61015760055481565b604051908152602001610118565b61015760085481565b61015760035481565b61017f6101ef565b005b61015760025481565b61015760065481565b61015760045481565b6000546040516001600160a01b039091168152602001610118565b61015760015481565b61015760075481565b61017f6101d73660046103ba565b610203565b61017f6101ea3660046103d3565b610217565b6101f7610257565b6102016000610284565b565b61020b610257565b610214816102e1565b50565b61021f610257565b6001600160a01b03811661024e57604051631e4fbdf760e01b8152600060048201526024015b60405180910390fd5b61021481610284565b6000546001600160a01b031633146102015760405163118cdaa760e01b8152336004820152602401610245565b600080546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b8035600155602081013560025560408101356003556060810135600455608081013560055560a081013560065560c081013560075560e081013560085561033061012082016101008301610420565b6009805467ffffffffffffffff191667ffffffffffffffff929092169190911790556103646101408201610120830161044d565b600960086101000a81548161ffff021916908361ffff1602179055507fce7570d0b935bf7c8cc4eb76a6e3405085353109cf57c8f63d46ae3939281249816040516103af9190610468565b60405180910390a150565b600061014082840312156103cd57600080fd5b50919050565b6000602082840312156103e557600080fd5b81356001600160a01b03811681146103fc57600080fd5b9392505050565b803567ffffffffffffffff8116811461041b57600080fd5b919050565b60006020828403121561043257600080fd5b6103fc82610403565b803561ffff8116811461041b57600080fd5b60006020828403121561045f57600080fd5b6103fc8261043b565b600061014082019050823582526020830135602083015260408301356040830152606083013560608301526080830135608083015260a083013560a083015260c083013560c083015260e083013560e08301526101006104c9818501610403565b67ffffffffffffffff16908301526101206104e584820161043b565b61ffff1692019190915291905056fea26469706673582212207e4a571d114c3698360bade80103c3e939b20b24eeac38a2596208226b8263ab64736f6c63430008140033",
}

// SettingscontractABI is the input ABI used to generate the binding from.
// Deprecated: Use SettingscontractMetaData.ABI instead.
var SettingscontractABI = SettingscontractMetaData.ABI

// SettingscontractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SettingscontractMetaData.Bin instead.
var SettingscontractBin = SettingscontractMetaData.Bin

// DeploySettingscontract deploys a new Ethereum contract, binding an instance of Settingscontract to it.
func DeploySettingscontract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Settingscontract, error) {
	parsed, err := SettingscontractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SettingscontractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Settingscontract{SettingscontractCaller: SettingscontractCaller{contract: contract}, SettingscontractTransactor: SettingscontractTransactor{contract: contract}, SettingscontractFilterer: SettingscontractFilterer{contract: contract}}, nil
}

// Settingscontract is an auto generated Go binding around an Ethereum contract.
type Settingscontract struct {
	SettingscontractCaller     // Read-only binding to the contract
	SettingscontractTransactor // Write-only binding to the contract
	SettingscontractFilterer   // Log filterer for contract events
}

// SettingscontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SettingscontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SettingscontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SettingscontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SettingscontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SettingscontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SettingscontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SettingscontractSession struct {
	Contract     *Settingscontract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SettingscontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SettingscontractCallerSession struct {
	Contract *SettingscontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// SettingscontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SettingscontractTransactorSession struct {
	Contract     *SettingscontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SettingscontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SettingscontractRaw struct {
	Contract *Settingscontract // Generic contract binding to access the raw methods on
}

// SettingscontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SettingscontractCallerRaw struct {
	Contract *SettingscontractCaller // Generic read-only contract binding to access the raw methods on
}

// SettingscontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SettingscontractTransactorRaw struct {
	Contract *SettingscontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSettingscontract creates a new instance of Settingscontract, bound to a specific deployed contract.
func NewSettingscontract(address common.Address, backend bind.ContractBackend) (*Settingscontract, error) {
	contract, err := bindSettingscontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Settingscontract{SettingscontractCaller: SettingscontractCaller{contract: contract}, SettingscontractTransactor: SettingscontractTransactor{contract: contract}, SettingscontractFilterer: SettingscontractFilterer{contract: contract}}, nil
}

// NewSettingscontractCaller creates a new read-only instance of Settingscontract, bound to a specific deployed contract.
func NewSettingscontractCaller(address common.Address, caller bind.ContractCaller) (*SettingscontractCaller, error) {
	contract, err := bindSettingscontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SettingscontractCaller{contract: contract}, nil
}

// NewSettingscontractTransactor creates a new write-only instance of Settingscontract, bound to a specific deployed contract.
func NewSettingscontractTransactor(address common.Address, transactor bind.ContractTransactor) (*SettingscontractTransactor, error) {
	contract, err := bindSettingscontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SettingscontractTransactor{contract: contract}, nil
}

// NewSettingscontractFilterer creates a new log filterer instance of Settingscontract, bound to a specific deployed contract.
func NewSettingscontractFilterer(address common.Address, filterer bind.ContractFilterer) (*SettingscontractFilterer, error) {
	contract, err := bindSettingscontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SettingscontractFilterer{contract: contract}, nil
}

// bindSettingscontract binds a generic wrapper to an already deployed contract.
func bindSettingscontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SettingscontractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Settingscontract *SettingscontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Settingscontract.Contract.SettingscontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Settingscontract *SettingscontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Settingscontract.Contract.SettingscontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Settingscontract *SettingscontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Settingscontract.Contract.SettingscontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Settingscontract *SettingscontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Settingscontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Settingscontract *SettingscontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Settingscontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Settingscontract *SettingscontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Settingscontract.Contract.contract.Transact(opts, method, params...)
}

// MaxNodeWeights is a free data retrieval call binding the contract method 0x15fbbc9d.
//
// Solidity: function maxNodeWeights() view returns(uint16)
func (_Settingscontract *SettingscontractCaller) MaxNodeWeights(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Settingscontract.contract.Call(opts, &out, "maxNodeWeights")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MaxNodeWeights is a free data retrieval call binding the contract method 0x15fbbc9d.
//
// Solidity: function maxNodeWeights() view returns(uint16)
func (_Settingscontract *SettingscontractSession) MaxNodeWeights() (uint16, error) {
	return _Settingscontract.Contract.MaxNodeWeights(&_Settingscontract.CallOpts)
}

// MaxNodeWeights is a free data retrieval call binding the contract method 0x15fbbc9d.
//
// Solidity: function maxNodeWeights() view returns(uint16)
func (_Settingscontract *SettingscontractCallerSession) MaxNodeWeights() (uint16, error) {
	return _Settingscontract.Contract.MaxNodeWeights(&_Settingscontract.CallOpts)
}

// MaxVrfActiveNodes is a free data retrieval call binding the contract method 0xb416dd85.
//
// Solidity: function maxVrfActiveNodes() view returns(uint256)
func (_Settingscontract *SettingscontractCaller) MaxVrfActiveNodes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Settingscontract.contract.Call(opts, &out, "maxVrfActiveNodes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxVrfActiveNodes is a free data retrieval call binding the contract method 0xb416dd85.
//
// Solidity: function maxVrfActiveNodes() view returns(uint256)
func (_Settingscontract *SettingscontractSession) MaxVrfActiveNodes() (*big.Int, error) {
	return _Settingscontract.Contract.MaxVrfActiveNodes(&_Settingscontract.CallOpts)
}

// MaxVrfActiveNodes is a free data retrieval call binding the contract method 0xb416dd85.
//
// Solidity: function maxVrfActiveNodes() view returns(uint256)
func (_Settingscontract *SettingscontractCallerSession) MaxVrfActiveNodes() (*big.Int, error) {
	return _Settingscontract.Contract.MaxVrfActiveNodes(&_Settingscontract.CallOpts)
}

// MinCommissionRateModifyInterval is a free data retrieval call binding the contract method 0x512ab483.
//
// Solidity: function minCommissionRateModifyInterval() view returns(uint256)
func (_Settingscontract *SettingscontractCaller) MinCommissionRateModifyInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Settingscontract.contract.Call(opts, &out, "minCommissionRateModifyInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinCommissionRateModifyInterval is a free data retrieval call binding the contract method 0x512ab483.
//
// Solidity: function minCommissionRateModifyInterval() view returns(uint256)
func (_Settingscontract *SettingscontractSession) MinCommissionRateModifyInterval() (*big.Int, error) {
	return _Settingscontract.Contract.MinCommissionRateModifyInterval(&_Settingscontract.CallOpts)
}

// MinCommissionRateModifyInterval is a free data retrieval call binding the contract method 0x512ab483.
//
// Solidity: function minCommissionRateModifyInterval() view returns(uint256)
func (_Settingscontract *SettingscontractCallerSession) MinCommissionRateModifyInterval() (*big.Int, error) {
	return _Settingscontract.Contract.MinCommissionRateModifyInterval(&_Settingscontract.CallOpts)
}

// MinTeeStakeAmount is a free data retrieval call binding the contract method 0x3c6174bc.
//
// Solidity: function minTeeStakeAmount() view returns(uint256)
func (_Settingscontract *SettingscontractCaller) MinTeeStakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Settingscontract.contract.Call(opts, &out, "minTeeStakeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinTeeStakeAmount is a free data retrieval call binding the contract method 0x3c6174bc.
//
// Solidity: function minTeeStakeAmount() view returns(uint256)
func (_Settingscontract *SettingscontractSession) MinTeeStakeAmount() (*big.Int, error) {
	return _Settingscontract.Contract.MinTeeStakeAmount(&_Settingscontract.CallOpts)
}

// MinTeeStakeAmount is a free data retrieval call binding the contract method 0x3c6174bc.
//
// Solidity: function minTeeStakeAmount() view returns(uint256)
func (_Settingscontract *SettingscontractCallerSession) MinTeeStakeAmount() (*big.Int, error) {
	return _Settingscontract.Contract.MinTeeStakeAmount(&_Settingscontract.CallOpts)
}

// NodeMaxMissVerifyCount is a free data retrieval call binding the contract method 0x0b097ac3.
//
// Solidity: function nodeMaxMissVerifyCount() view returns(uint64)
func (_Settingscontract *SettingscontractCaller) NodeMaxMissVerifyCount(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Settingscontract.contract.Call(opts, &out, "nodeMaxMissVerifyCount")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NodeMaxMissVerifyCount is a free data retrieval call binding the contract method 0x0b097ac3.
//
// Solidity: function nodeMaxMissVerifyCount() view returns(uint64)
func (_Settingscontract *SettingscontractSession) NodeMaxMissVerifyCount() (uint64, error) {
	return _Settingscontract.Contract.NodeMaxMissVerifyCount(&_Settingscontract.CallOpts)
}

// NodeMaxMissVerifyCount is a free data retrieval call binding the contract method 0x0b097ac3.
//
// Solidity: function nodeMaxMissVerifyCount() view returns(uint64)
func (_Settingscontract *SettingscontractCallerSession) NodeMaxMissVerifyCount() (uint64, error) {
	return _Settingscontract.Contract.NodeMaxMissVerifyCount(&_Settingscontract.CallOpts)
}

// NodeMinOnlineDuration is a free data retrieval call binding the contract method 0x729fe45f.
//
// Solidity: function nodeMinOnlineDuration() view returns(uint256)
func (_Settingscontract *SettingscontractCaller) NodeMinOnlineDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Settingscontract.contract.Call(opts, &out, "nodeMinOnlineDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodeMinOnlineDuration is a free data retrieval call binding the contract method 0x729fe45f.
//
// Solidity: function nodeMinOnlineDuration() view returns(uint256)
func (_Settingscontract *SettingscontractSession) NodeMinOnlineDuration() (*big.Int, error) {
	return _Settingscontract.Contract.NodeMinOnlineDuration(&_Settingscontract.CallOpts)
}

// NodeMinOnlineDuration is a free data retrieval call binding the contract method 0x729fe45f.
//
// Solidity: function nodeMinOnlineDuration() view returns(uint256)
func (_Settingscontract *SettingscontractCallerSession) NodeMinOnlineDuration() (*big.Int, error) {
	return _Settingscontract.Contract.NodeMinOnlineDuration(&_Settingscontract.CallOpts)
}

// NodeSlashReward is a free data retrieval call binding the contract method 0x88cef9c2.
//
// Solidity: function nodeSlashReward() view returns(uint256)
func (_Settingscontract *SettingscontractCaller) NodeSlashReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Settingscontract.contract.Call(opts, &out, "nodeSlashReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodeSlashReward is a free data retrieval call binding the contract method 0x88cef9c2.
//
// Solidity: function nodeSlashReward() view returns(uint256)
func (_Settingscontract *SettingscontractSession) NodeSlashReward() (*big.Int, error) {
	return _Settingscontract.Contract.NodeSlashReward(&_Settingscontract.CallOpts)
}

// NodeSlashReward is a free data retrieval call binding the contract method 0x88cef9c2.
//
// Solidity: function nodeSlashReward() view returns(uint256)
func (_Settingscontract *SettingscontractCallerSession) NodeSlashReward() (*big.Int, error) {
	return _Settingscontract.Contract.NodeSlashReward(&_Settingscontract.CallOpts)
}

// NodeVerifyDuration is a free data retrieval call binding the contract method 0x600d0949.
//
// Solidity: function nodeVerifyDuration() view returns(uint256)
func (_Settingscontract *SettingscontractCaller) NodeVerifyDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Settingscontract.contract.Call(opts, &out, "nodeVerifyDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodeVerifyDuration is a free data retrieval call binding the contract method 0x600d0949.
//
// Solidity: function nodeVerifyDuration() view returns(uint256)
func (_Settingscontract *SettingscontractSession) NodeVerifyDuration() (*big.Int, error) {
	return _Settingscontract.Contract.NodeVerifyDuration(&_Settingscontract.CallOpts)
}

// NodeVerifyDuration is a free data retrieval call binding the contract method 0x600d0949.
//
// Solidity: function nodeVerifyDuration() view returns(uint256)
func (_Settingscontract *SettingscontractCallerSession) NodeVerifyDuration() (*big.Int, error) {
	return _Settingscontract.Contract.NodeVerifyDuration(&_Settingscontract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Settingscontract *SettingscontractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Settingscontract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Settingscontract *SettingscontractSession) Owner() (common.Address, error) {
	return _Settingscontract.Contract.Owner(&_Settingscontract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Settingscontract *SettingscontractCallerSession) Owner() (common.Address, error) {
	return _Settingscontract.Contract.Owner(&_Settingscontract.CallOpts)
}

// TeeSlashAmount is a free data retrieval call binding the contract method 0x73cf675e.
//
// Solidity: function teeSlashAmount() view returns(uint256)
func (_Settingscontract *SettingscontractCaller) TeeSlashAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Settingscontract.contract.Call(opts, &out, "teeSlashAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TeeSlashAmount is a free data retrieval call binding the contract method 0x73cf675e.
//
// Solidity: function teeSlashAmount() view returns(uint256)
func (_Settingscontract *SettingscontractSession) TeeSlashAmount() (*big.Int, error) {
	return _Settingscontract.Contract.TeeSlashAmount(&_Settingscontract.CallOpts)
}

// TeeSlashAmount is a free data retrieval call binding the contract method 0x73cf675e.
//
// Solidity: function teeSlashAmount() view returns(uint256)
func (_Settingscontract *SettingscontractCallerSession) TeeSlashAmount() (*big.Int, error) {
	return _Settingscontract.Contract.TeeSlashAmount(&_Settingscontract.CallOpts)
}

// TeeUnstakeDuration is a free data retrieval call binding the contract method 0xcaa9eeba.
//
// Solidity: function teeUnstakeDuration() view returns(uint256)
func (_Settingscontract *SettingscontractCaller) TeeUnstakeDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Settingscontract.contract.Call(opts, &out, "teeUnstakeDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TeeUnstakeDuration is a free data retrieval call binding the contract method 0xcaa9eeba.
//
// Solidity: function teeUnstakeDuration() view returns(uint256)
func (_Settingscontract *SettingscontractSession) TeeUnstakeDuration() (*big.Int, error) {
	return _Settingscontract.Contract.TeeUnstakeDuration(&_Settingscontract.CallOpts)
}

// TeeUnstakeDuration is a free data retrieval call binding the contract method 0xcaa9eeba.
//
// Solidity: function teeUnstakeDuration() view returns(uint256)
func (_Settingscontract *SettingscontractCallerSession) TeeUnstakeDuration() (*big.Int, error) {
	return _Settingscontract.Contract.TeeUnstakeDuration(&_Settingscontract.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Settingscontract *SettingscontractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Settingscontract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Settingscontract *SettingscontractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Settingscontract.Contract.RenounceOwnership(&_Settingscontract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Settingscontract *SettingscontractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Settingscontract.Contract.RenounceOwnership(&_Settingscontract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Settingscontract *SettingscontractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Settingscontract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Settingscontract *SettingscontractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Settingscontract.Contract.TransferOwnership(&_Settingscontract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Settingscontract *SettingscontractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Settingscontract.Contract.TransferOwnership(&_Settingscontract.TransactOpts, newOwner)
}

// UpdateSettings is a paid mutator transaction binding the contract method 0xee6b2820.
//
// Solidity: function updateSettings((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint16) params) returns()
func (_Settingscontract *SettingscontractTransactor) UpdateSettings(opts *bind.TransactOpts, params ISettingsSettingParams) (*types.Transaction, error) {
	return _Settingscontract.contract.Transact(opts, "updateSettings", params)
}

// UpdateSettings is a paid mutator transaction binding the contract method 0xee6b2820.
//
// Solidity: function updateSettings((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint16) params) returns()
func (_Settingscontract *SettingscontractSession) UpdateSettings(params ISettingsSettingParams) (*types.Transaction, error) {
	return _Settingscontract.Contract.UpdateSettings(&_Settingscontract.TransactOpts, params)
}

// UpdateSettings is a paid mutator transaction binding the contract method 0xee6b2820.
//
// Solidity: function updateSettings((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint16) params) returns()
func (_Settingscontract *SettingscontractTransactorSession) UpdateSettings(params ISettingsSettingParams) (*types.Transaction, error) {
	return _Settingscontract.Contract.UpdateSettings(&_Settingscontract.TransactOpts, params)
}

// SettingscontractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Settingscontract contract.
type SettingscontractOwnershipTransferredIterator struct {
	Event *SettingscontractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SettingscontractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SettingscontractOwnershipTransferred)
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
		it.Event = new(SettingscontractOwnershipTransferred)
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
func (it *SettingscontractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SettingscontractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SettingscontractOwnershipTransferred represents a OwnershipTransferred event raised by the Settingscontract contract.
type SettingscontractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Settingscontract *SettingscontractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SettingscontractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Settingscontract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SettingscontractOwnershipTransferredIterator{contract: _Settingscontract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Settingscontract *SettingscontractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SettingscontractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Settingscontract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SettingscontractOwnershipTransferred)
				if err := _Settingscontract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Settingscontract *SettingscontractFilterer) ParseOwnershipTransferred(log types.Log) (*SettingscontractOwnershipTransferred, error) {
	event := new(SettingscontractOwnershipTransferred)
	if err := _Settingscontract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SettingscontractUpdateSettingsIterator is returned from FilterUpdateSettings and is used to iterate over the raw logs and unpacked data for UpdateSettings events raised by the Settingscontract contract.
type SettingscontractUpdateSettingsIterator struct {
	Event *SettingscontractUpdateSettings // Event containing the contract specifics and raw log

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
func (it *SettingscontractUpdateSettingsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SettingscontractUpdateSettings)
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
		it.Event = new(SettingscontractUpdateSettings)
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
func (it *SettingscontractUpdateSettingsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SettingscontractUpdateSettingsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SettingscontractUpdateSettings represents a UpdateSettings event raised by the Settingscontract contract.
type SettingscontractUpdateSettings struct {
	Params ISettingsSettingParams
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdateSettings is a free log retrieval operation binding the contract event 0xce7570d0b935bf7c8cc4eb76a6e3405085353109cf57c8f63d46ae3939281249.
//
// Solidity: event UpdateSettings((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint16) params)
func (_Settingscontract *SettingscontractFilterer) FilterUpdateSettings(opts *bind.FilterOpts) (*SettingscontractUpdateSettingsIterator, error) {

	logs, sub, err := _Settingscontract.contract.FilterLogs(opts, "UpdateSettings")
	if err != nil {
		return nil, err
	}
	return &SettingscontractUpdateSettingsIterator{contract: _Settingscontract.contract, event: "UpdateSettings", logs: logs, sub: sub}, nil
}

// WatchUpdateSettings is a free log subscription operation binding the contract event 0xce7570d0b935bf7c8cc4eb76a6e3405085353109cf57c8f63d46ae3939281249.
//
// Solidity: event UpdateSettings((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint16) params)
func (_Settingscontract *SettingscontractFilterer) WatchUpdateSettings(opts *bind.WatchOpts, sink chan<- *SettingscontractUpdateSettings) (event.Subscription, error) {

	logs, sub, err := _Settingscontract.contract.WatchLogs(opts, "UpdateSettings")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SettingscontractUpdateSettings)
				if err := _Settingscontract.contract.UnpackLog(event, "UpdateSettings", log); err != nil {
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

// ParseUpdateSettings is a log parse operation binding the contract event 0xce7570d0b935bf7c8cc4eb76a6e3405085353109cf57c8f63d46ae3939281249.
//
// Solidity: event UpdateSettings((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint16) params)
func (_Settingscontract *SettingscontractFilterer) ParseUpdateSettings(log types.Log) (*SettingscontractUpdateSettings, error) {
	event := new(SettingscontractUpdateSettings)
	if err := _Settingscontract.contract.UnpackLog(event, "UpdateSettings", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

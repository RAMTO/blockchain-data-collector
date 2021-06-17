// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lmc

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// LmcABI is the input ABI used to generate the binding from.
const LmcABI = "[{\"inputs\":[{\"internalType\":\"contractIERC20Detailed\",\"name\":\"_stakingToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endBlock\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_rewardsTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_rewardPerBlock\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_albtAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stakeLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_contractStakeLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bonusAmount\",\"type\":\"uint256\"}],\"name\":\"BonusTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Exited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"ExitedAndUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newEndBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"newRewardsPerBlock\",\"type\":\"uint256[]\"}],\"name\":\"Extended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_lockScheme\",\"type\":\"address\"}],\"name\":\"StakedAndLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rewardsAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"WithdrawLPRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accumulatedRewardMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractStakeLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_rewardsPerBlock\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_currentRemainingRewards\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_newRemainingRewards\",\"type\":\"uint256[]\"}],\"name\":\"extend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardTokensCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenIndex\",\"type\":\"uint256\"}],\"name\":\"getUserAccumulatedReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getUserOwedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getUserRewardDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"getUserRewardDebtLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"getUserTokensOwedLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasStakingStarted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRewardBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lockSchemes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockSchemesExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"receiversWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsPoolFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardsTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"contractIERC20Detailed\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateRewardMultipliers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userAccruedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"firstStakedBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lpTokenContract\",\"type\":\"address\"}],\"name\":\"withdrawLPRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_lockScheme\",\"type\":\"address\"}],\"name\":\"stakeAndLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exitAndUnlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"}],\"name\":\"setReceiverWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakePool\",\"type\":\"address\"}],\"name\":\"exitAndStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transferTo\",\"type\":\"address\"}],\"name\":\"exitAndTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_lockSchemes\",\"type\":\"address[]\"}],\"name\":\"setLockSchemes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Lmc is an auto generated Go binding around an Ethereum contract.
type Lmc struct {
	LmcCaller     // Read-only binding to the contract
	LmcTransactor // Write-only binding to the contract
	LmcFilterer   // Log filterer for contract events
}

// LmcCaller is an auto generated read-only Go binding around an Ethereum contract.
type LmcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LmcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LmcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LmcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LmcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LmcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LmcSession struct {
	Contract     *Lmc              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LmcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LmcCallerSession struct {
	Contract *LmcCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LmcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LmcTransactorSession struct {
	Contract     *LmcTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LmcRaw is an auto generated low-level Go binding around an Ethereum contract.
type LmcRaw struct {
	Contract *Lmc // Generic contract binding to access the raw methods on
}

// LmcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LmcCallerRaw struct {
	Contract *LmcCaller // Generic read-only contract binding to access the raw methods on
}

// LmcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LmcTransactorRaw struct {
	Contract *LmcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLmc creates a new instance of Lmc, bound to a specific deployed contract.
func NewLmc(address common.Address, backend bind.ContractBackend) (*Lmc, error) {
	contract, err := bindLmc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lmc{LmcCaller: LmcCaller{contract: contract}, LmcTransactor: LmcTransactor{contract: contract}, LmcFilterer: LmcFilterer{contract: contract}}, nil
}

// NewLmcCaller creates a new read-only instance of Lmc, bound to a specific deployed contract.
func NewLmcCaller(address common.Address, caller bind.ContractCaller) (*LmcCaller, error) {
	contract, err := bindLmc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LmcCaller{contract: contract}, nil
}

// NewLmcTransactor creates a new write-only instance of Lmc, bound to a specific deployed contract.
func NewLmcTransactor(address common.Address, transactor bind.ContractTransactor) (*LmcTransactor, error) {
	contract, err := bindLmc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LmcTransactor{contract: contract}, nil
}

// NewLmcFilterer creates a new log filterer instance of Lmc, bound to a specific deployed contract.
func NewLmcFilterer(address common.Address, filterer bind.ContractFilterer) (*LmcFilterer, error) {
	contract, err := bindLmc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LmcFilterer{contract: contract}, nil
}

// bindLmc binds a generic wrapper to an already deployed contract.
func bindLmc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LmcABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lmc *LmcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lmc.Contract.LmcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lmc *LmcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lmc.Contract.LmcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lmc *LmcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lmc.Contract.LmcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lmc *LmcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lmc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lmc *LmcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lmc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lmc *LmcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lmc.Contract.contract.Transact(opts, method, params...)
}

// AccumulatedRewardMultiplier is a free data retrieval call binding the contract method 0xfb58cad1.
//
// Solidity: function accumulatedRewardMultiplier(uint256 ) view returns(uint256)
func (_Lmc *LmcCaller) AccumulatedRewardMultiplier(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lmc.contract.Call(opts, &out, "accumulatedRewardMultiplier", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedRewardMultiplier is a free data retrieval call binding the contract method 0xfb58cad1.
//
// Solidity: function accumulatedRewardMultiplier(uint256 ) view returns(uint256)
func (_Lmc *LmcSession) AccumulatedRewardMultiplier(arg0 *big.Int) (*big.Int, error) {
	return _Lmc.Contract.AccumulatedRewardMultiplier(&_Lmc.CallOpts, arg0)
}

// AccumulatedRewardMultiplier is a free data retrieval call binding the contract method 0xfb58cad1.
//
// Solidity: function accumulatedRewardMultiplier(uint256 ) view returns(uint256)
func (_Lmc *LmcCallerSession) AccumulatedRewardMultiplier(arg0 *big.Int) (*big.Int, error) {
	return _Lmc.Contract.AccumulatedRewardMultiplier(&_Lmc.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _userAddress) view returns(uint256)
func (_Lmc *LmcCaller) BalanceOf(opts *bind.CallOpts, _userAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lmc.contract.Call(opts, &out, "balanceOf", _userAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _userAddress) view returns(uint256)
func (_Lmc *LmcSession) BalanceOf(_userAddress common.Address) (*big.Int, error) {
	return _Lmc.Contract.BalanceOf(&_Lmc.CallOpts, _userAddress)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _userAddress) view returns(uint256)
func (_Lmc *LmcCallerSession) BalanceOf(_userAddress common.Address) (*big.Int, error) {
	return _Lmc.Contract.BalanceOf(&_Lmc.CallOpts, _userAddress)
}

// ContractStakeLimit is a free data retrieval call binding the contract method 0x03d1dae0.
//
// Solidity: function contractStakeLimit() view returns(uint256)
func (_Lmc *LmcCaller) ContractStakeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lmc.contract.Call(opts, &out, "contractStakeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractStakeLimit is a free data retrieval call binding the contract method 0x03d1dae0.
//
// Solidity: function contractStakeLimit() view returns(uint256)
func (_Lmc *LmcSession) ContractStakeLimit() (*big.Int, error) {
	return _Lmc.Contract.ContractStakeLimit(&_Lmc.CallOpts)
}

// ContractStakeLimit is a free data retrieval call binding the contract method 0x03d1dae0.
//
// Solidity: function contractStakeLimit() view returns(uint256)
func (_Lmc *LmcCallerSession) ContractStakeLimit() (*big.Int, error) {
	return _Lmc.Contract.ContractStakeLimit(&_Lmc.CallOpts)
}

// EndBlock is a free data retrieval call binding the contract method 0x083c6323.
//
// Solidity: function endBlock() view returns(uint256)
func (_Lmc *LmcCaller) EndBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lmc.contract.Call(opts, &out, "endBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndBlock is a free data retrieval call binding the contract method 0x083c6323.
//
// Solidity: function endBlock() view returns(uint256)
func (_Lmc *LmcSession) EndBlock() (*big.Int, error) {
	return _Lmc.Contract.EndBlock(&_Lmc.CallOpts)
}

// EndBlock is a free data retrieval call binding the contract method 0x083c6323.
//
// Solidity: function endBlock() view returns(uint256)
func (_Lmc *LmcCallerSession) EndBlock() (*big.Int, error) {
	return _Lmc.Contract.EndBlock(&_Lmc.CallOpts)
}

// GetRewardTokensCount is a free data retrieval call binding the contract method 0x2d9e88e1.
//
// Solidity: function getRewardTokensCount() view returns(uint256)
func (_Lmc *LmcCaller) GetRewardTokensCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lmc.contract.Call(opts, &out, "getRewardTokensCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardTokensCount is a free data retrieval call binding the contract method 0x2d9e88e1.
//
// Solidity: function getRewardTokensCount() view returns(uint256)
func (_Lmc *LmcSession) GetRewardTokensCount() (*big.Int, error) {
	return _Lmc.Contract.GetRewardTokensCount(&_Lmc.CallOpts)
}

// GetRewardTokensCount is a free data retrieval call binding the contract method 0x2d9e88e1.
//
// Solidity: function getRewardTokensCount() view returns(uint256)
func (_Lmc *LmcCallerSession) GetRewardTokensCount() (*big.Int, error) {
	return _Lmc.Contract.GetRewardTokensCount(&_Lmc.CallOpts)
}

// GetUserAccumulatedReward is a free data retrieval call binding the contract method 0xdf9d777f.
//
// Solidity: function getUserAccumulatedReward(address _userAddress, uint256 tokenIndex) view returns(uint256)
func (_Lmc *LmcCaller) GetUserAccumulatedReward(opts *bind.CallOpts, _userAddress common.Address, tokenIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lmc.contract.Call(opts, &out, "getUserAccumulatedReward", _userAddress, tokenIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserAccumulatedReward is a free data retrieval call binding the contract method 0xdf9d777f.
//
// Solidity: function getUserAccumulatedReward(address _userAddress, uint256 tokenIndex) view returns(uint256)
func (_Lmc *LmcSession) GetUserAccumulatedReward(_userAddress common.Address, tokenIndex *big.Int) (*big.Int, error) {
	return _Lmc.Contract.GetUserAccumulatedReward(&_Lmc.CallOpts, _userAddress, tokenIndex)
}

// GetUserAccumulatedReward is a free data retrieval call binding the contract method 0xdf9d777f.
//
// Solidity: function getUserAccumulatedReward(address _userAddress, uint256 tokenIndex) view returns(uint256)
func (_Lmc *LmcCallerSession) GetUserAccumulatedReward(_userAddress common.Address, tokenIndex *big.Int) (*big.Int, error) {
	return _Lmc.Contract.GetUserAccumulatedReward(&_Lmc.CallOpts, _userAddress, tokenIndex)
}

// GetUserOwedTokens is a free data retrieval call binding the contract method 0xce415302.
//
// Solidity: function getUserOwedTokens(address _userAddress, uint256 _index) view returns(uint256)
func (_Lmc *LmcCaller) GetUserOwedTokens(opts *bind.CallOpts, _userAddress common.Address, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lmc.contract.Call(opts, &out, "getUserOwedTokens", _userAddress, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserOwedTokens is a free data retrieval call binding the contract method 0xce415302.
//
// Solidity: function getUserOwedTokens(address _userAddress, uint256 _index) view returns(uint256)
func (_Lmc *LmcSession) GetUserOwedTokens(_userAddress common.Address, _index *big.Int) (*big.Int, error) {
	return _Lmc.Contract.GetUserOwedTokens(&_Lmc.CallOpts, _userAddress, _index)
}

// GetUserOwedTokens is a free data retrieval call binding the contract method 0xce415302.
//
// Solidity: function getUserOwedTokens(address _userAddress, uint256 _index) view returns(uint256)
func (_Lmc *LmcCallerSession) GetUserOwedTokens(_userAddress common.Address, _index *big.Int) (*big.Int, error) {
	return _Lmc.Contract.GetUserOwedTokens(&_Lmc.CallOpts, _userAddress, _index)
}

// GetUserRewardDebt is a free data retrieval call binding the contract method 0xf27d0264.
//
// Solidity: function getUserRewardDebt(address _userAddress, uint256 _index) view returns(uint256)
func (_Lmc *LmcCaller) GetUserRewardDebt(opts *bind.CallOpts, _userAddress common.Address, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lmc.contract.Call(opts, &out, "getUserRewardDebt", _userAddress, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserRewardDebt is a free data retrieval call binding the contract method 0xf27d0264.
//
// Solidity: function getUserRewardDebt(address _userAddress, uint256 _index) view returns(uint256)
func (_Lmc *LmcSession) GetUserRewardDebt(_userAddress common.Address, _index *big.Int) (*big.Int, error) {
	return _Lmc.Contract.GetUserRewardDebt(&_Lmc.CallOpts, _userAddress, _index)
}

// GetUserRewardDebt is a free data retrieval call binding the contract method 0xf27d0264.
//
// Solidity: function getUserRewardDebt(address _userAddress, uint256 _index) view returns(uint256)
func (_Lmc *LmcCallerSession) GetUserRewardDebt(_userAddress common.Address, _index *big.Int) (*big.Int, error) {
	return _Lmc.Contract.GetUserRewardDebt(&_Lmc.CallOpts, _userAddress, _index)
}

// GetUserRewardDebtLength is a free data retrieval call binding the contract method 0x0084c927.
//
// Solidity: function getUserRewardDebtLength(address _userAddress) view returns(uint256)
func (_Lmc *LmcCaller) GetUserRewardDebtLength(opts *bind.CallOpts, _userAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lmc.contract.Call(opts, &out, "getUserRewardDebtLength", _userAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserRewardDebtLength is a free data retrieval call binding the contract method 0x0084c927.
//
// Solidity: function getUserRewardDebtLength(address _userAddress) view returns(uint256)
func (_Lmc *LmcSession) GetUserRewardDebtLength(_userAddress common.Address) (*big.Int, error) {
	return _Lmc.Contract.GetUserRewardDebtLength(&_Lmc.CallOpts, _userAddress)
}

// GetUserRewardDebtLength is a free data retrieval call binding the contract method 0x0084c927.
//
// Solidity: function getUserRewardDebtLength(address _userAddress) view returns(uint256)
func (_Lmc *LmcCallerSession) GetUserRewardDebtLength(_userAddress common.Address) (*big.Int, error) {
	return _Lmc.Contract.GetUserRewardDebtLength(&_Lmc.CallOpts, _userAddress)
}

// GetUserTokensOwedLength is a free data retrieval call binding the contract method 0xa1292aea.
//
// Solidity: function getUserTokensOwedLength(address _userAddress) view returns(uint256)
func (_Lmc *LmcCaller) GetUserTokensOwedLength(opts *bind.CallOpts, _userAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lmc.contract.Call(opts, &out, "getUserTokensOwedLength", _userAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserTokensOwedLength is a free data retrieval call binding the contract method 0xa1292aea.
//
// Solidity: function getUserTokensOwedLength(address _userAddress) view returns(uint256)
func (_Lmc *LmcSession) GetUserTokensOwedLength(_userAddress common.Address) (*big.Int, error) {
	return _Lmc.Contract.GetUserTokensOwedLength(&_Lmc.CallOpts, _userAddress)
}

// GetUserTokensOwedLength is a free data retrieval call binding the contract method 0xa1292aea.
//
// Solidity: function getUserTokensOwedLength(address _userAddress) view returns(uint256)
func (_Lmc *LmcCallerSession) GetUserTokensOwedLength(_userAddress common.Address) (*big.Int, error) {
	return _Lmc.Contract.GetUserTokensOwedLength(&_Lmc.CallOpts, _userAddress)
}

// HasStakingStarted is a free data retrieval call binding the contract method 0x57b4f01f.
//
// Solidity: function hasStakingStarted() view returns(bool)
func (_Lmc *LmcCaller) HasStakingStarted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lmc.contract.Call(opts, &out, "hasStakingStarted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasStakingStarted is a free data retrieval call binding the contract method 0x57b4f01f.
//
// Solidity: function hasStakingStarted() view returns(bool)
func (_Lmc *LmcSession) HasStakingStarted() (bool, error) {
	return _Lmc.Contract.HasStakingStarted(&_Lmc.CallOpts)
}

// HasStakingStarted is a free data retrieval call binding the contract method 0x57b4f01f.
//
// Solidity: function hasStakingStarted() view returns(bool)
func (_Lmc *LmcCallerSession) HasStakingStarted() (bool, error) {
	return _Lmc.Contract.HasStakingStarted(&_Lmc.CallOpts)
}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_Lmc *LmcCaller) LastRewardBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lmc.contract.Call(opts, &out, "lastRewardBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_Lmc *LmcSession) LastRewardBlock() (*big.Int, error) {
	return _Lmc.Contract.LastRewardBlock(&_Lmc.CallOpts)
}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_Lmc *LmcCallerSession) LastRewardBlock() (*big.Int, error) {
	return _Lmc.Contract.LastRewardBlock(&_Lmc.CallOpts)
}

// LockSchemes is a free data retrieval call binding the contract method 0xbb9c9eb4.
//
// Solidity: function lockSchemes(uint256 ) view returns(address)
func (_Lmc *LmcCaller) LockSchemes(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Lmc.contract.Call(opts, &out, "lockSchemes", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LockSchemes is a free data retrieval call binding the contract method 0xbb9c9eb4.
//
// Solidity: function lockSchemes(uint256 ) view returns(address)
func (_Lmc *LmcSession) LockSchemes(arg0 *big.Int) (common.Address, error) {
	return _Lmc.Contract.LockSchemes(&_Lmc.CallOpts, arg0)
}

// LockSchemes is a free data retrieval call binding the contract method 0xbb9c9eb4.
//
// Solidity: function lockSchemes(uint256 ) view returns(address)
func (_Lmc *LmcCallerSession) LockSchemes(arg0 *big.Int) (common.Address, error) {
	return _Lmc.Contract.LockSchemes(&_Lmc.CallOpts, arg0)
}

// LockSchemesExist is a free data retrieval call binding the contract method 0x205acb31.
//
// Solidity: function lockSchemesExist(address ) view returns(bool)
func (_Lmc *LmcCaller) LockSchemesExist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Lmc.contract.Call(opts, &out, "lockSchemesExist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LockSchemesExist is a free data retrieval call binding the contract method 0x205acb31.
//
// Solidity: function lockSchemesExist(address ) view returns(bool)
func (_Lmc *LmcSession) LockSchemesExist(arg0 common.Address) (bool, error) {
	return _Lmc.Contract.LockSchemesExist(&_Lmc.CallOpts, arg0)
}

// LockSchemesExist is a free data retrieval call binding the contract method 0x205acb31.
//
// Solidity: function lockSchemesExist(address ) view returns(bool)
func (_Lmc *LmcCallerSession) LockSchemesExist(arg0 common.Address) (bool, error) {
	return _Lmc.Contract.LockSchemesExist(&_Lmc.CallOpts, arg0)
}

// ReceiversWhitelist is a free data retrieval call binding the contract method 0x363291dc.
//
// Solidity: function receiversWhitelist(address ) view returns(bool)
func (_Lmc *LmcCaller) ReceiversWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Lmc.contract.Call(opts, &out, "receiversWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ReceiversWhitelist is a free data retrieval call binding the contract method 0x363291dc.
//
// Solidity: function receiversWhitelist(address ) view returns(bool)
func (_Lmc *LmcSession) ReceiversWhitelist(arg0 common.Address) (bool, error) {
	return _Lmc.Contract.ReceiversWhitelist(&_Lmc.CallOpts, arg0)
}

// ReceiversWhitelist is a free data retrieval call binding the contract method 0x363291dc.
//
// Solidity: function receiversWhitelist(address ) view returns(bool)
func (_Lmc *LmcCallerSession) ReceiversWhitelist(arg0 common.Address) (bool, error) {
	return _Lmc.Contract.ReceiversWhitelist(&_Lmc.CallOpts, arg0)
}

// RewardPerBlock is a free data retrieval call binding the contract method 0x791f39cd.
//
// Solidity: function rewardPerBlock(uint256 ) view returns(uint256)
func (_Lmc *LmcCaller) RewardPerBlock(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lmc.contract.Call(opts, &out, "rewardPerBlock", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerBlock is a free data retrieval call binding the contract method 0x791f39cd.
//
// Solidity: function rewardPerBlock(uint256 ) view returns(uint256)
func (_Lmc *LmcSession) RewardPerBlock(arg0 *big.Int) (*big.Int, error) {
	return _Lmc.Contract.RewardPerBlock(&_Lmc.CallOpts, arg0)
}

// RewardPerBlock is a free data retrieval call binding the contract method 0x791f39cd.
//
// Solidity: function rewardPerBlock(uint256 ) view returns(uint256)
func (_Lmc *LmcCallerSession) RewardPerBlock(arg0 *big.Int) (*big.Int, error) {
	return _Lmc.Contract.RewardPerBlock(&_Lmc.CallOpts, arg0)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Lmc *LmcCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lmc.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Lmc *LmcSession) RewardToken() (common.Address, error) {
	return _Lmc.Contract.RewardToken(&_Lmc.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Lmc *LmcCallerSession) RewardToken() (common.Address, error) {
	return _Lmc.Contract.RewardToken(&_Lmc.CallOpts)
}

// RewardsPoolFactory is a free data retrieval call binding the contract method 0x56409b81.
//
// Solidity: function rewardsPoolFactory() view returns(address)
func (_Lmc *LmcCaller) RewardsPoolFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lmc.contract.Call(opts, &out, "rewardsPoolFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsPoolFactory is a free data retrieval call binding the contract method 0x56409b81.
//
// Solidity: function rewardsPoolFactory() view returns(address)
func (_Lmc *LmcSession) RewardsPoolFactory() (common.Address, error) {
	return _Lmc.Contract.RewardsPoolFactory(&_Lmc.CallOpts)
}

// RewardsPoolFactory is a free data retrieval call binding the contract method 0x56409b81.
//
// Solidity: function rewardsPoolFactory() view returns(address)
func (_Lmc *LmcCallerSession) RewardsPoolFactory() (common.Address, error) {
	return _Lmc.Contract.RewardsPoolFactory(&_Lmc.CallOpts)
}

// RewardsTokens is a free data retrieval call binding the contract method 0xb6d0dcd8.
//
// Solidity: function rewardsTokens(uint256 ) view returns(address)
func (_Lmc *LmcCaller) RewardsTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Lmc.contract.Call(opts, &out, "rewardsTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsTokens is a free data retrieval call binding the contract method 0xb6d0dcd8.
//
// Solidity: function rewardsTokens(uint256 ) view returns(address)
func (_Lmc *LmcSession) RewardsTokens(arg0 *big.Int) (common.Address, error) {
	return _Lmc.Contract.RewardsTokens(&_Lmc.CallOpts, arg0)
}

// RewardsTokens is a free data retrieval call binding the contract method 0xb6d0dcd8.
//
// Solidity: function rewardsTokens(uint256 ) view returns(address)
func (_Lmc *LmcCallerSession) RewardsTokens(arg0 *big.Int) (common.Address, error) {
	return _Lmc.Contract.RewardsTokens(&_Lmc.CallOpts, arg0)
}

// StakeLimit is a free data retrieval call binding the contract method 0x45ef79af.
//
// Solidity: function stakeLimit() view returns(uint256)
func (_Lmc *LmcCaller) StakeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lmc.contract.Call(opts, &out, "stakeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeLimit is a free data retrieval call binding the contract method 0x45ef79af.
//
// Solidity: function stakeLimit() view returns(uint256)
func (_Lmc *LmcSession) StakeLimit() (*big.Int, error) {
	return _Lmc.Contract.StakeLimit(&_Lmc.CallOpts)
}

// StakeLimit is a free data retrieval call binding the contract method 0x45ef79af.
//
// Solidity: function stakeLimit() view returns(uint256)
func (_Lmc *LmcCallerSession) StakeLimit() (*big.Int, error) {
	return _Lmc.Contract.StakeLimit(&_Lmc.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_Lmc *LmcCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lmc.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_Lmc *LmcSession) StakingToken() (common.Address, error) {
	return _Lmc.Contract.StakingToken(&_Lmc.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_Lmc *LmcCallerSession) StakingToken() (common.Address, error) {
	return _Lmc.Contract.StakingToken(&_Lmc.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Lmc *LmcCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lmc.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Lmc *LmcSession) StartBlock() (*big.Int, error) {
	return _Lmc.Contract.StartBlock(&_Lmc.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Lmc *LmcCallerSession) StartBlock() (*big.Int, error) {
	return _Lmc.Contract.StartBlock(&_Lmc.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_Lmc *LmcCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lmc.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_Lmc *LmcSession) TotalStaked() (*big.Int, error) {
	return _Lmc.Contract.TotalStaked(&_Lmc.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_Lmc *LmcCallerSession) TotalStaked() (*big.Int, error) {
	return _Lmc.Contract.TotalStaked(&_Lmc.CallOpts)
}

// UserAccruedRewards is a free data retrieval call binding the contract method 0x408651be.
//
// Solidity: function userAccruedRewards(address ) view returns(uint256)
func (_Lmc *LmcCaller) UserAccruedRewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lmc.contract.Call(opts, &out, "userAccruedRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserAccruedRewards is a free data retrieval call binding the contract method 0x408651be.
//
// Solidity: function userAccruedRewards(address ) view returns(uint256)
func (_Lmc *LmcSession) UserAccruedRewards(arg0 common.Address) (*big.Int, error) {
	return _Lmc.Contract.UserAccruedRewards(&_Lmc.CallOpts, arg0)
}

// UserAccruedRewards is a free data retrieval call binding the contract method 0x408651be.
//
// Solidity: function userAccruedRewards(address ) view returns(uint256)
func (_Lmc *LmcCallerSession) UserAccruedRewards(arg0 common.Address) (*big.Int, error) {
	return _Lmc.Contract.UserAccruedRewards(&_Lmc.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 firstStakedBlockNumber, uint256 amountStaked)
func (_Lmc *LmcCaller) UserInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	FirstStakedBlockNumber *big.Int
	AmountStaked           *big.Int
}, error) {
	var out []interface{}
	err := _Lmc.contract.Call(opts, &out, "userInfo", arg0)

	outstruct := new(struct {
		FirstStakedBlockNumber *big.Int
		AmountStaked           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FirstStakedBlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountStaked = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 firstStakedBlockNumber, uint256 amountStaked)
func (_Lmc *LmcSession) UserInfo(arg0 common.Address) (struct {
	FirstStakedBlockNumber *big.Int
	AmountStaked           *big.Int
}, error) {
	return _Lmc.Contract.UserInfo(&_Lmc.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 firstStakedBlockNumber, uint256 amountStaked)
func (_Lmc *LmcCallerSession) UserInfo(arg0 common.Address) (struct {
	FirstStakedBlockNumber *big.Int
	AmountStaked           *big.Int
}, error) {
	return _Lmc.Contract.UserInfo(&_Lmc.CallOpts, arg0)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Lmc *LmcTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lmc.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Lmc *LmcSession) Claim() (*types.Transaction, error) {
	return _Lmc.Contract.Claim(&_Lmc.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Lmc *LmcTransactorSession) Claim() (*types.Transaction, error) {
	return _Lmc.Contract.Claim(&_Lmc.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Lmc *LmcTransactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lmc.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Lmc *LmcSession) Exit() (*types.Transaction, error) {
	return _Lmc.Contract.Exit(&_Lmc.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Lmc *LmcTransactorSession) Exit() (*types.Transaction, error) {
	return _Lmc.Contract.Exit(&_Lmc.TransactOpts)
}

// ExitAndStake is a paid mutator transaction binding the contract method 0x5999e473.
//
// Solidity: function exitAndStake(address _stakePool) returns()
func (_Lmc *LmcTransactor) ExitAndStake(opts *bind.TransactOpts, _stakePool common.Address) (*types.Transaction, error) {
	return _Lmc.contract.Transact(opts, "exitAndStake", _stakePool)
}

// ExitAndStake is a paid mutator transaction binding the contract method 0x5999e473.
//
// Solidity: function exitAndStake(address _stakePool) returns()
func (_Lmc *LmcSession) ExitAndStake(_stakePool common.Address) (*types.Transaction, error) {
	return _Lmc.Contract.ExitAndStake(&_Lmc.TransactOpts, _stakePool)
}

// ExitAndStake is a paid mutator transaction binding the contract method 0x5999e473.
//
// Solidity: function exitAndStake(address _stakePool) returns()
func (_Lmc *LmcTransactorSession) ExitAndStake(_stakePool common.Address) (*types.Transaction, error) {
	return _Lmc.Contract.ExitAndStake(&_Lmc.TransactOpts, _stakePool)
}

// ExitAndTransfer is a paid mutator transaction binding the contract method 0x2240e63c.
//
// Solidity: function exitAndTransfer(address transferTo) returns()
func (_Lmc *LmcTransactor) ExitAndTransfer(opts *bind.TransactOpts, transferTo common.Address) (*types.Transaction, error) {
	return _Lmc.contract.Transact(opts, "exitAndTransfer", transferTo)
}

// ExitAndTransfer is a paid mutator transaction binding the contract method 0x2240e63c.
//
// Solidity: function exitAndTransfer(address transferTo) returns()
func (_Lmc *LmcSession) ExitAndTransfer(transferTo common.Address) (*types.Transaction, error) {
	return _Lmc.Contract.ExitAndTransfer(&_Lmc.TransactOpts, transferTo)
}

// ExitAndTransfer is a paid mutator transaction binding the contract method 0x2240e63c.
//
// Solidity: function exitAndTransfer(address transferTo) returns()
func (_Lmc *LmcTransactorSession) ExitAndTransfer(transferTo common.Address) (*types.Transaction, error) {
	return _Lmc.Contract.ExitAndTransfer(&_Lmc.TransactOpts, transferTo)
}

// ExitAndUnlock is a paid mutator transaction binding the contract method 0x3411ef51.
//
// Solidity: function exitAndUnlock() returns()
func (_Lmc *LmcTransactor) ExitAndUnlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lmc.contract.Transact(opts, "exitAndUnlock")
}

// ExitAndUnlock is a paid mutator transaction binding the contract method 0x3411ef51.
//
// Solidity: function exitAndUnlock() returns()
func (_Lmc *LmcSession) ExitAndUnlock() (*types.Transaction, error) {
	return _Lmc.Contract.ExitAndUnlock(&_Lmc.TransactOpts)
}

// ExitAndUnlock is a paid mutator transaction binding the contract method 0x3411ef51.
//
// Solidity: function exitAndUnlock() returns()
func (_Lmc *LmcTransactorSession) ExitAndUnlock() (*types.Transaction, error) {
	return _Lmc.Contract.ExitAndUnlock(&_Lmc.TransactOpts)
}

// Extend is a paid mutator transaction binding the contract method 0x20e67c76.
//
// Solidity: function extend(uint256 _endBlock, uint256[] _rewardsPerBlock, uint256[] _currentRemainingRewards, uint256[] _newRemainingRewards) returns()
func (_Lmc *LmcTransactor) Extend(opts *bind.TransactOpts, _endBlock *big.Int, _rewardsPerBlock []*big.Int, _currentRemainingRewards []*big.Int, _newRemainingRewards []*big.Int) (*types.Transaction, error) {
	return _Lmc.contract.Transact(opts, "extend", _endBlock, _rewardsPerBlock, _currentRemainingRewards, _newRemainingRewards)
}

// Extend is a paid mutator transaction binding the contract method 0x20e67c76.
//
// Solidity: function extend(uint256 _endBlock, uint256[] _rewardsPerBlock, uint256[] _currentRemainingRewards, uint256[] _newRemainingRewards) returns()
func (_Lmc *LmcSession) Extend(_endBlock *big.Int, _rewardsPerBlock []*big.Int, _currentRemainingRewards []*big.Int, _newRemainingRewards []*big.Int) (*types.Transaction, error) {
	return _Lmc.Contract.Extend(&_Lmc.TransactOpts, _endBlock, _rewardsPerBlock, _currentRemainingRewards, _newRemainingRewards)
}

// Extend is a paid mutator transaction binding the contract method 0x20e67c76.
//
// Solidity: function extend(uint256 _endBlock, uint256[] _rewardsPerBlock, uint256[] _currentRemainingRewards, uint256[] _newRemainingRewards) returns()
func (_Lmc *LmcTransactorSession) Extend(_endBlock *big.Int, _rewardsPerBlock []*big.Int, _currentRemainingRewards []*big.Int, _newRemainingRewards []*big.Int) (*types.Transaction, error) {
	return _Lmc.Contract.Extend(&_Lmc.TransactOpts, _endBlock, _rewardsPerBlock, _currentRemainingRewards, _newRemainingRewards)
}

// SetLockSchemes is a paid mutator transaction binding the contract method 0xc8523c1d.
//
// Solidity: function setLockSchemes(address[] _lockSchemes) returns()
func (_Lmc *LmcTransactor) SetLockSchemes(opts *bind.TransactOpts, _lockSchemes []common.Address) (*types.Transaction, error) {
	return _Lmc.contract.Transact(opts, "setLockSchemes", _lockSchemes)
}

// SetLockSchemes is a paid mutator transaction binding the contract method 0xc8523c1d.
//
// Solidity: function setLockSchemes(address[] _lockSchemes) returns()
func (_Lmc *LmcSession) SetLockSchemes(_lockSchemes []common.Address) (*types.Transaction, error) {
	return _Lmc.Contract.SetLockSchemes(&_Lmc.TransactOpts, _lockSchemes)
}

// SetLockSchemes is a paid mutator transaction binding the contract method 0xc8523c1d.
//
// Solidity: function setLockSchemes(address[] _lockSchemes) returns()
func (_Lmc *LmcTransactorSession) SetLockSchemes(_lockSchemes []common.Address) (*types.Transaction, error) {
	return _Lmc.Contract.SetLockSchemes(&_Lmc.TransactOpts, _lockSchemes)
}

// SetReceiverWhitelisted is a paid mutator transaction binding the contract method 0xa861a7a3.
//
// Solidity: function setReceiverWhitelisted(address receiver, bool whitelisted) returns()
func (_Lmc *LmcTransactor) SetReceiverWhitelisted(opts *bind.TransactOpts, receiver common.Address, whitelisted bool) (*types.Transaction, error) {
	return _Lmc.contract.Transact(opts, "setReceiverWhitelisted", receiver, whitelisted)
}

// SetReceiverWhitelisted is a paid mutator transaction binding the contract method 0xa861a7a3.
//
// Solidity: function setReceiverWhitelisted(address receiver, bool whitelisted) returns()
func (_Lmc *LmcSession) SetReceiverWhitelisted(receiver common.Address, whitelisted bool) (*types.Transaction, error) {
	return _Lmc.Contract.SetReceiverWhitelisted(&_Lmc.TransactOpts, receiver, whitelisted)
}

// SetReceiverWhitelisted is a paid mutator transaction binding the contract method 0xa861a7a3.
//
// Solidity: function setReceiverWhitelisted(address receiver, bool whitelisted) returns()
func (_Lmc *LmcTransactorSession) SetReceiverWhitelisted(receiver common.Address, whitelisted bool) (*types.Transaction, error) {
	return _Lmc.Contract.SetReceiverWhitelisted(&_Lmc.TransactOpts, receiver, whitelisted)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _tokenAmount) returns()
func (_Lmc *LmcTransactor) Stake(opts *bind.TransactOpts, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _Lmc.contract.Transact(opts, "stake", _tokenAmount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _tokenAmount) returns()
func (_Lmc *LmcSession) Stake(_tokenAmount *big.Int) (*types.Transaction, error) {
	return _Lmc.Contract.Stake(&_Lmc.TransactOpts, _tokenAmount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _tokenAmount) returns()
func (_Lmc *LmcTransactorSession) Stake(_tokenAmount *big.Int) (*types.Transaction, error) {
	return _Lmc.Contract.Stake(&_Lmc.TransactOpts, _tokenAmount)
}

// StakeAndLock is a paid mutator transaction binding the contract method 0x4715a949.
//
// Solidity: function stakeAndLock(uint256 _tokenAmount, address _lockScheme) returns()
func (_Lmc *LmcTransactor) StakeAndLock(opts *bind.TransactOpts, _tokenAmount *big.Int, _lockScheme common.Address) (*types.Transaction, error) {
	return _Lmc.contract.Transact(opts, "stakeAndLock", _tokenAmount, _lockScheme)
}

// StakeAndLock is a paid mutator transaction binding the contract method 0x4715a949.
//
// Solidity: function stakeAndLock(uint256 _tokenAmount, address _lockScheme) returns()
func (_Lmc *LmcSession) StakeAndLock(_tokenAmount *big.Int, _lockScheme common.Address) (*types.Transaction, error) {
	return _Lmc.Contract.StakeAndLock(&_Lmc.TransactOpts, _tokenAmount, _lockScheme)
}

// StakeAndLock is a paid mutator transaction binding the contract method 0x4715a949.
//
// Solidity: function stakeAndLock(uint256 _tokenAmount, address _lockScheme) returns()
func (_Lmc *LmcTransactorSession) StakeAndLock(_tokenAmount *big.Int, _lockScheme common.Address) (*types.Transaction, error) {
	return _Lmc.Contract.StakeAndLock(&_Lmc.TransactOpts, _tokenAmount, _lockScheme)
}

// UpdateRewardMultipliers is a paid mutator transaction binding the contract method 0xdd2da220.
//
// Solidity: function updateRewardMultipliers() returns()
func (_Lmc *LmcTransactor) UpdateRewardMultipliers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lmc.contract.Transact(opts, "updateRewardMultipliers")
}

// UpdateRewardMultipliers is a paid mutator transaction binding the contract method 0xdd2da220.
//
// Solidity: function updateRewardMultipliers() returns()
func (_Lmc *LmcSession) UpdateRewardMultipliers() (*types.Transaction, error) {
	return _Lmc.Contract.UpdateRewardMultipliers(&_Lmc.TransactOpts)
}

// UpdateRewardMultipliers is a paid mutator transaction binding the contract method 0xdd2da220.
//
// Solidity: function updateRewardMultipliers() returns()
func (_Lmc *LmcTransactorSession) UpdateRewardMultipliers() (*types.Transaction, error) {
	return _Lmc.Contract.UpdateRewardMultipliers(&_Lmc.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _tokenAmount) returns()
func (_Lmc *LmcTransactor) Withdraw(opts *bind.TransactOpts, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _Lmc.contract.Transact(opts, "withdraw", _tokenAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _tokenAmount) returns()
func (_Lmc *LmcSession) Withdraw(_tokenAmount *big.Int) (*types.Transaction, error) {
	return _Lmc.Contract.Withdraw(&_Lmc.TransactOpts, _tokenAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _tokenAmount) returns()
func (_Lmc *LmcTransactorSession) Withdraw(_tokenAmount *big.Int) (*types.Transaction, error) {
	return _Lmc.Contract.Withdraw(&_Lmc.TransactOpts, _tokenAmount)
}

// WithdrawLPRewards is a paid mutator transaction binding the contract method 0xa1002a0f.
//
// Solidity: function withdrawLPRewards(address recipient, address lpTokenContract) returns()
func (_Lmc *LmcTransactor) WithdrawLPRewards(opts *bind.TransactOpts, recipient common.Address, lpTokenContract common.Address) (*types.Transaction, error) {
	return _Lmc.contract.Transact(opts, "withdrawLPRewards", recipient, lpTokenContract)
}

// WithdrawLPRewards is a paid mutator transaction binding the contract method 0xa1002a0f.
//
// Solidity: function withdrawLPRewards(address recipient, address lpTokenContract) returns()
func (_Lmc *LmcSession) WithdrawLPRewards(recipient common.Address, lpTokenContract common.Address) (*types.Transaction, error) {
	return _Lmc.Contract.WithdrawLPRewards(&_Lmc.TransactOpts, recipient, lpTokenContract)
}

// WithdrawLPRewards is a paid mutator transaction binding the contract method 0xa1002a0f.
//
// Solidity: function withdrawLPRewards(address recipient, address lpTokenContract) returns()
func (_Lmc *LmcTransactorSession) WithdrawLPRewards(recipient common.Address, lpTokenContract common.Address) (*types.Transaction, error) {
	return _Lmc.Contract.WithdrawLPRewards(&_Lmc.TransactOpts, recipient, lpTokenContract)
}

// LmcBonusTransferredIterator is returned from FilterBonusTransferred and is used to iterate over the raw logs and unpacked data for BonusTransferred events raised by the Lmc contract.
type LmcBonusTransferredIterator struct {
	Event *LmcBonusTransferred // Event containing the contract specifics and raw log

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
func (it *LmcBonusTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LmcBonusTransferred)
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
		it.Event = new(LmcBonusTransferred)
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
func (it *LmcBonusTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LmcBonusTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LmcBonusTransferred represents a BonusTransferred event raised by the Lmc contract.
type LmcBonusTransferred struct {
	UserAddress common.Address
	BonusAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBonusTransferred is a free log retrieval operation binding the contract event 0x9e40dbc4af813371d457949ccb8f04e60076420f1659f45d6bc93588c0dd82df.
//
// Solidity: event BonusTransferred(address indexed _userAddress, uint256 _bonusAmount)
func (_Lmc *LmcFilterer) FilterBonusTransferred(opts *bind.FilterOpts, _userAddress []common.Address) (*LmcBonusTransferredIterator, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}

	logs, sub, err := _Lmc.contract.FilterLogs(opts, "BonusTransferred", _userAddressRule)
	if err != nil {
		return nil, err
	}
	return &LmcBonusTransferredIterator{contract: _Lmc.contract, event: "BonusTransferred", logs: logs, sub: sub}, nil
}

// WatchBonusTransferred is a free log subscription operation binding the contract event 0x9e40dbc4af813371d457949ccb8f04e60076420f1659f45d6bc93588c0dd82df.
//
// Solidity: event BonusTransferred(address indexed _userAddress, uint256 _bonusAmount)
func (_Lmc *LmcFilterer) WatchBonusTransferred(opts *bind.WatchOpts, sink chan<- *LmcBonusTransferred, _userAddress []common.Address) (event.Subscription, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}

	logs, sub, err := _Lmc.contract.WatchLogs(opts, "BonusTransferred", _userAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LmcBonusTransferred)
				if err := _Lmc.contract.UnpackLog(event, "BonusTransferred", log); err != nil {
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

// ParseBonusTransferred is a log parse operation binding the contract event 0x9e40dbc4af813371d457949ccb8f04e60076420f1659f45d6bc93588c0dd82df.
//
// Solidity: event BonusTransferred(address indexed _userAddress, uint256 _bonusAmount)
func (_Lmc *LmcFilterer) ParseBonusTransferred(log types.Log) (*LmcBonusTransferred, error) {
	event := new(LmcBonusTransferred)
	if err := _Lmc.contract.UnpackLog(event, "BonusTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LmcClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the Lmc contract.
type LmcClaimedIterator struct {
	Event *LmcClaimed // Event containing the contract specifics and raw log

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
func (it *LmcClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LmcClaimed)
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
		it.Event = new(LmcClaimed)
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
func (it *LmcClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LmcClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LmcClaimed represents a Claimed event raised by the Lmc contract.
type LmcClaimed struct {
	User   common.Address
	Amount *big.Int
	Token  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x7e6632ca16a0ac6cf28448500b1a17d96c8b8163ad4c4a9b44ef5386cc02779e.
//
// Solidity: event Claimed(address indexed user, uint256 amount, address token)
func (_Lmc *LmcFilterer) FilterClaimed(opts *bind.FilterOpts, user []common.Address) (*LmcClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lmc.contract.FilterLogs(opts, "Claimed", userRule)
	if err != nil {
		return nil, err
	}
	return &LmcClaimedIterator{contract: _Lmc.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x7e6632ca16a0ac6cf28448500b1a17d96c8b8163ad4c4a9b44ef5386cc02779e.
//
// Solidity: event Claimed(address indexed user, uint256 amount, address token)
func (_Lmc *LmcFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *LmcClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lmc.contract.WatchLogs(opts, "Claimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LmcClaimed)
				if err := _Lmc.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x7e6632ca16a0ac6cf28448500b1a17d96c8b8163ad4c4a9b44ef5386cc02779e.
//
// Solidity: event Claimed(address indexed user, uint256 amount, address token)
func (_Lmc *LmcFilterer) ParseClaimed(log types.Log) (*LmcClaimed, error) {
	event := new(LmcClaimed)
	if err := _Lmc.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LmcExitedIterator is returned from FilterExited and is used to iterate over the raw logs and unpacked data for Exited events raised by the Lmc contract.
type LmcExitedIterator struct {
	Event *LmcExited // Event containing the contract specifics and raw log

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
func (it *LmcExitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LmcExited)
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
		it.Event = new(LmcExited)
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
func (it *LmcExitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LmcExitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LmcExited represents a Exited event raised by the Lmc contract.
type LmcExited struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExited is a free log retrieval operation binding the contract event 0x920bb94eb3842a728db98228c375ff6b00c5bc5a54fac6736155517a0a20a61a.
//
// Solidity: event Exited(address indexed user, uint256 amount)
func (_Lmc *LmcFilterer) FilterExited(opts *bind.FilterOpts, user []common.Address) (*LmcExitedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lmc.contract.FilterLogs(opts, "Exited", userRule)
	if err != nil {
		return nil, err
	}
	return &LmcExitedIterator{contract: _Lmc.contract, event: "Exited", logs: logs, sub: sub}, nil
}

// WatchExited is a free log subscription operation binding the contract event 0x920bb94eb3842a728db98228c375ff6b00c5bc5a54fac6736155517a0a20a61a.
//
// Solidity: event Exited(address indexed user, uint256 amount)
func (_Lmc *LmcFilterer) WatchExited(opts *bind.WatchOpts, sink chan<- *LmcExited, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lmc.contract.WatchLogs(opts, "Exited", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LmcExited)
				if err := _Lmc.contract.UnpackLog(event, "Exited", log); err != nil {
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

// ParseExited is a log parse operation binding the contract event 0x920bb94eb3842a728db98228c375ff6b00c5bc5a54fac6736155517a0a20a61a.
//
// Solidity: event Exited(address indexed user, uint256 amount)
func (_Lmc *LmcFilterer) ParseExited(log types.Log) (*LmcExited, error) {
	event := new(LmcExited)
	if err := _Lmc.contract.UnpackLog(event, "Exited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LmcExitedAndUnlockedIterator is returned from FilterExitedAndUnlocked and is used to iterate over the raw logs and unpacked data for ExitedAndUnlocked events raised by the Lmc contract.
type LmcExitedAndUnlockedIterator struct {
	Event *LmcExitedAndUnlocked // Event containing the contract specifics and raw log

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
func (it *LmcExitedAndUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LmcExitedAndUnlocked)
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
		it.Event = new(LmcExitedAndUnlocked)
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
func (it *LmcExitedAndUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LmcExitedAndUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LmcExitedAndUnlocked represents a ExitedAndUnlocked event raised by the Lmc contract.
type LmcExitedAndUnlocked struct {
	UserAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExitedAndUnlocked is a free log retrieval operation binding the contract event 0x73d36a08b171571d289bbb4c400b9c74176d5be660fae67a23ed20ba1120d583.
//
// Solidity: event ExitedAndUnlocked(address indexed _userAddress)
func (_Lmc *LmcFilterer) FilterExitedAndUnlocked(opts *bind.FilterOpts, _userAddress []common.Address) (*LmcExitedAndUnlockedIterator, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}

	logs, sub, err := _Lmc.contract.FilterLogs(opts, "ExitedAndUnlocked", _userAddressRule)
	if err != nil {
		return nil, err
	}
	return &LmcExitedAndUnlockedIterator{contract: _Lmc.contract, event: "ExitedAndUnlocked", logs: logs, sub: sub}, nil
}

// WatchExitedAndUnlocked is a free log subscription operation binding the contract event 0x73d36a08b171571d289bbb4c400b9c74176d5be660fae67a23ed20ba1120d583.
//
// Solidity: event ExitedAndUnlocked(address indexed _userAddress)
func (_Lmc *LmcFilterer) WatchExitedAndUnlocked(opts *bind.WatchOpts, sink chan<- *LmcExitedAndUnlocked, _userAddress []common.Address) (event.Subscription, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}

	logs, sub, err := _Lmc.contract.WatchLogs(opts, "ExitedAndUnlocked", _userAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LmcExitedAndUnlocked)
				if err := _Lmc.contract.UnpackLog(event, "ExitedAndUnlocked", log); err != nil {
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

// ParseExitedAndUnlocked is a log parse operation binding the contract event 0x73d36a08b171571d289bbb4c400b9c74176d5be660fae67a23ed20ba1120d583.
//
// Solidity: event ExitedAndUnlocked(address indexed _userAddress)
func (_Lmc *LmcFilterer) ParseExitedAndUnlocked(log types.Log) (*LmcExitedAndUnlocked, error) {
	event := new(LmcExitedAndUnlocked)
	if err := _Lmc.contract.UnpackLog(event, "ExitedAndUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LmcExtendedIterator is returned from FilterExtended and is used to iterate over the raw logs and unpacked data for Extended events raised by the Lmc contract.
type LmcExtendedIterator struct {
	Event *LmcExtended // Event containing the contract specifics and raw log

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
func (it *LmcExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LmcExtended)
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
		it.Event = new(LmcExtended)
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
func (it *LmcExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LmcExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LmcExtended represents a Extended event raised by the Lmc contract.
type LmcExtended struct {
	NewEndBlock        *big.Int
	NewRewardsPerBlock []*big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterExtended is a free log retrieval operation binding the contract event 0x137c92cc7579cc4d6a2b109467cd475c205d1c136363ca854cc46d72f840d5de.
//
// Solidity: event Extended(uint256 newEndBlock, uint256[] newRewardsPerBlock)
func (_Lmc *LmcFilterer) FilterExtended(opts *bind.FilterOpts) (*LmcExtendedIterator, error) {

	logs, sub, err := _Lmc.contract.FilterLogs(opts, "Extended")
	if err != nil {
		return nil, err
	}
	return &LmcExtendedIterator{contract: _Lmc.contract, event: "Extended", logs: logs, sub: sub}, nil
}

// WatchExtended is a free log subscription operation binding the contract event 0x137c92cc7579cc4d6a2b109467cd475c205d1c136363ca854cc46d72f840d5de.
//
// Solidity: event Extended(uint256 newEndBlock, uint256[] newRewardsPerBlock)
func (_Lmc *LmcFilterer) WatchExtended(opts *bind.WatchOpts, sink chan<- *LmcExtended) (event.Subscription, error) {

	logs, sub, err := _Lmc.contract.WatchLogs(opts, "Extended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LmcExtended)
				if err := _Lmc.contract.UnpackLog(event, "Extended", log); err != nil {
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

// ParseExtended is a log parse operation binding the contract event 0x137c92cc7579cc4d6a2b109467cd475c205d1c136363ca854cc46d72f840d5de.
//
// Solidity: event Extended(uint256 newEndBlock, uint256[] newRewardsPerBlock)
func (_Lmc *LmcFilterer) ParseExtended(log types.Log) (*LmcExtended, error) {
	event := new(LmcExtended)
	if err := _Lmc.contract.UnpackLog(event, "Extended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LmcStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Lmc contract.
type LmcStakedIterator struct {
	Event *LmcStaked // Event containing the contract specifics and raw log

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
func (it *LmcStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LmcStaked)
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
		it.Event = new(LmcStaked)
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
func (it *LmcStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LmcStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LmcStaked represents a Staked event raised by the Lmc contract.
type LmcStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Lmc *LmcFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*LmcStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lmc.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &LmcStakedIterator{contract: _Lmc.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Lmc *LmcFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *LmcStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lmc.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LmcStaked)
				if err := _Lmc.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Lmc *LmcFilterer) ParseStaked(log types.Log) (*LmcStaked, error) {
	event := new(LmcStaked)
	if err := _Lmc.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LmcStakedAndLockedIterator is returned from FilterStakedAndLocked and is used to iterate over the raw logs and unpacked data for StakedAndLocked events raised by the Lmc contract.
type LmcStakedAndLockedIterator struct {
	Event *LmcStakedAndLocked // Event containing the contract specifics and raw log

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
func (it *LmcStakedAndLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LmcStakedAndLocked)
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
		it.Event = new(LmcStakedAndLocked)
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
func (it *LmcStakedAndLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LmcStakedAndLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LmcStakedAndLocked represents a StakedAndLocked event raised by the Lmc contract.
type LmcStakedAndLocked struct {
	UserAddress common.Address
	TokenAmount *big.Int
	LockScheme  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStakedAndLocked is a free log retrieval operation binding the contract event 0x5ae9eac9eae389addf073e8e829ae488ab31d186b6234575446951ba6e53390e.
//
// Solidity: event StakedAndLocked(address indexed _userAddress, uint256 _tokenAmount, address _lockScheme)
func (_Lmc *LmcFilterer) FilterStakedAndLocked(opts *bind.FilterOpts, _userAddress []common.Address) (*LmcStakedAndLockedIterator, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}

	logs, sub, err := _Lmc.contract.FilterLogs(opts, "StakedAndLocked", _userAddressRule)
	if err != nil {
		return nil, err
	}
	return &LmcStakedAndLockedIterator{contract: _Lmc.contract, event: "StakedAndLocked", logs: logs, sub: sub}, nil
}

// WatchStakedAndLocked is a free log subscription operation binding the contract event 0x5ae9eac9eae389addf073e8e829ae488ab31d186b6234575446951ba6e53390e.
//
// Solidity: event StakedAndLocked(address indexed _userAddress, uint256 _tokenAmount, address _lockScheme)
func (_Lmc *LmcFilterer) WatchStakedAndLocked(opts *bind.WatchOpts, sink chan<- *LmcStakedAndLocked, _userAddress []common.Address) (event.Subscription, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}

	logs, sub, err := _Lmc.contract.WatchLogs(opts, "StakedAndLocked", _userAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LmcStakedAndLocked)
				if err := _Lmc.contract.UnpackLog(event, "StakedAndLocked", log); err != nil {
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

// ParseStakedAndLocked is a log parse operation binding the contract event 0x5ae9eac9eae389addf073e8e829ae488ab31d186b6234575446951ba6e53390e.
//
// Solidity: event StakedAndLocked(address indexed _userAddress, uint256 _tokenAmount, address _lockScheme)
func (_Lmc *LmcFilterer) ParseStakedAndLocked(log types.Log) (*LmcStakedAndLocked, error) {
	event := new(LmcStakedAndLocked)
	if err := _Lmc.contract.UnpackLog(event, "StakedAndLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LmcWithdrawLPRewardsIterator is returned from FilterWithdrawLPRewards and is used to iterate over the raw logs and unpacked data for WithdrawLPRewards events raised by the Lmc contract.
type LmcWithdrawLPRewardsIterator struct {
	Event *LmcWithdrawLPRewards // Event containing the contract specifics and raw log

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
func (it *LmcWithdrawLPRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LmcWithdrawLPRewards)
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
		it.Event = new(LmcWithdrawLPRewards)
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
func (it *LmcWithdrawLPRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LmcWithdrawLPRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LmcWithdrawLPRewards represents a WithdrawLPRewards event raised by the Lmc contract.
type LmcWithdrawLPRewards struct {
	RewardsAmount *big.Int
	Recipient     common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawLPRewards is a free log retrieval operation binding the contract event 0xdf1b625465761dd1a4330a9b3a21d4943183f301036528a3544f8d891fdce2c2.
//
// Solidity: event WithdrawLPRewards(uint256 indexed rewardsAmount, address indexed recipient)
func (_Lmc *LmcFilterer) FilterWithdrawLPRewards(opts *bind.FilterOpts, rewardsAmount []*big.Int, recipient []common.Address) (*LmcWithdrawLPRewardsIterator, error) {

	var rewardsAmountRule []interface{}
	for _, rewardsAmountItem := range rewardsAmount {
		rewardsAmountRule = append(rewardsAmountRule, rewardsAmountItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Lmc.contract.FilterLogs(opts, "WithdrawLPRewards", rewardsAmountRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &LmcWithdrawLPRewardsIterator{contract: _Lmc.contract, event: "WithdrawLPRewards", logs: logs, sub: sub}, nil
}

// WatchWithdrawLPRewards is a free log subscription operation binding the contract event 0xdf1b625465761dd1a4330a9b3a21d4943183f301036528a3544f8d891fdce2c2.
//
// Solidity: event WithdrawLPRewards(uint256 indexed rewardsAmount, address indexed recipient)
func (_Lmc *LmcFilterer) WatchWithdrawLPRewards(opts *bind.WatchOpts, sink chan<- *LmcWithdrawLPRewards, rewardsAmount []*big.Int, recipient []common.Address) (event.Subscription, error) {

	var rewardsAmountRule []interface{}
	for _, rewardsAmountItem := range rewardsAmount {
		rewardsAmountRule = append(rewardsAmountRule, rewardsAmountItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Lmc.contract.WatchLogs(opts, "WithdrawLPRewards", rewardsAmountRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LmcWithdrawLPRewards)
				if err := _Lmc.contract.UnpackLog(event, "WithdrawLPRewards", log); err != nil {
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

// ParseWithdrawLPRewards is a log parse operation binding the contract event 0xdf1b625465761dd1a4330a9b3a21d4943183f301036528a3544f8d891fdce2c2.
//
// Solidity: event WithdrawLPRewards(uint256 indexed rewardsAmount, address indexed recipient)
func (_Lmc *LmcFilterer) ParseWithdrawLPRewards(log types.Log) (*LmcWithdrawLPRewards, error) {
	event := new(LmcWithdrawLPRewards)
	if err := _Lmc.contract.UnpackLog(event, "WithdrawLPRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LmcWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Lmc contract.
type LmcWithdrawnIterator struct {
	Event *LmcWithdrawn // Event containing the contract specifics and raw log

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
func (it *LmcWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LmcWithdrawn)
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
		it.Event = new(LmcWithdrawn)
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
func (it *LmcWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LmcWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LmcWithdrawn represents a Withdrawn event raised by the Lmc contract.
type LmcWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Lmc *LmcFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*LmcWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lmc.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &LmcWithdrawnIterator{contract: _Lmc.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Lmc *LmcFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *LmcWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lmc.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LmcWithdrawn)
				if err := _Lmc.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Lmc *LmcFilterer) ParseWithdrawn(log types.Log) (*LmcWithdrawn, error) {
	event := new(LmcWithdrawn)
	if err := _Lmc.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

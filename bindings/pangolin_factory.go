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

// IPangolinFactoryMetaData contains all meta data concerning the IPangolinFactory contract.
var IPangolinFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IPangolinFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IPangolinFactoryMetaData.ABI instead.
var IPangolinFactoryABI = IPangolinFactoryMetaData.ABI

// IPangolinFactory is an auto generated Go binding around an Ethereum contract.
type IPangolinFactory struct {
	IPangolinFactoryCaller     // Read-only binding to the contract
	IPangolinFactoryTransactor // Write-only binding to the contract
	IPangolinFactoryFilterer   // Log filterer for contract events
}

// IPangolinFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPangolinFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPangolinFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPangolinFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPangolinFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPangolinFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPangolinFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPangolinFactorySession struct {
	Contract     *IPangolinFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPangolinFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPangolinFactoryCallerSession struct {
	Contract *IPangolinFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IPangolinFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPangolinFactoryTransactorSession struct {
	Contract     *IPangolinFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IPangolinFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPangolinFactoryRaw struct {
	Contract *IPangolinFactory // Generic contract binding to access the raw methods on
}

// IPangolinFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPangolinFactoryCallerRaw struct {
	Contract *IPangolinFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IPangolinFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPangolinFactoryTransactorRaw struct {
	Contract *IPangolinFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPangolinFactory creates a new instance of IPangolinFactory, bound to a specific deployed contract.
func NewIPangolinFactory(address common.Address, backend bind.ContractBackend) (*IPangolinFactory, error) {
	contract, err := bindIPangolinFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPangolinFactory{IPangolinFactoryCaller: IPangolinFactoryCaller{contract: contract}, IPangolinFactoryTransactor: IPangolinFactoryTransactor{contract: contract}, IPangolinFactoryFilterer: IPangolinFactoryFilterer{contract: contract}}, nil
}

// NewIPangolinFactoryCaller creates a new read-only instance of IPangolinFactory, bound to a specific deployed contract.
func NewIPangolinFactoryCaller(address common.Address, caller bind.ContractCaller) (*IPangolinFactoryCaller, error) {
	contract, err := bindIPangolinFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPangolinFactoryCaller{contract: contract}, nil
}

// NewIPangolinFactoryTransactor creates a new write-only instance of IPangolinFactory, bound to a specific deployed contract.
func NewIPangolinFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IPangolinFactoryTransactor, error) {
	contract, err := bindIPangolinFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPangolinFactoryTransactor{contract: contract}, nil
}

// NewIPangolinFactoryFilterer creates a new log filterer instance of IPangolinFactory, bound to a specific deployed contract.
func NewIPangolinFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IPangolinFactoryFilterer, error) {
	contract, err := bindIPangolinFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPangolinFactoryFilterer{contract: contract}, nil
}

// bindIPangolinFactory binds a generic wrapper to an already deployed contract.
func bindIPangolinFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPangolinFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPangolinFactory *IPangolinFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPangolinFactory.Contract.IPangolinFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPangolinFactory *IPangolinFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPangolinFactory.Contract.IPangolinFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPangolinFactory *IPangolinFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPangolinFactory.Contract.IPangolinFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPangolinFactory *IPangolinFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPangolinFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPangolinFactory *IPangolinFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPangolinFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPangolinFactory *IPangolinFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPangolinFactory.Contract.contract.Transact(opts, method, params...)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_IPangolinFactory *IPangolinFactoryCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IPangolinFactory.contract.Call(opts, &out, "allPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_IPangolinFactory *IPangolinFactorySession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _IPangolinFactory.Contract.AllPairs(&_IPangolinFactory.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_IPangolinFactory *IPangolinFactoryCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _IPangolinFactory.Contract.AllPairs(&_IPangolinFactory.CallOpts, arg0)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_IPangolinFactory *IPangolinFactoryCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPangolinFactory.contract.Call(opts, &out, "allPairsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_IPangolinFactory *IPangolinFactorySession) AllPairsLength() (*big.Int, error) {
	return _IPangolinFactory.Contract.AllPairsLength(&_IPangolinFactory.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_IPangolinFactory *IPangolinFactoryCallerSession) AllPairsLength() (*big.Int, error) {
	return _IPangolinFactory.Contract.AllPairsLength(&_IPangolinFactory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_IPangolinFactory *IPangolinFactoryCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPangolinFactory.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_IPangolinFactory *IPangolinFactorySession) FeeTo() (common.Address, error) {
	return _IPangolinFactory.Contract.FeeTo(&_IPangolinFactory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_IPangolinFactory *IPangolinFactoryCallerSession) FeeTo() (common.Address, error) {
	return _IPangolinFactory.Contract.FeeTo(&_IPangolinFactory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_IPangolinFactory *IPangolinFactoryCaller) FeeToSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPangolinFactory.contract.Call(opts, &out, "feeToSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_IPangolinFactory *IPangolinFactorySession) FeeToSetter() (common.Address, error) {
	return _IPangolinFactory.Contract.FeeToSetter(&_IPangolinFactory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_IPangolinFactory *IPangolinFactoryCallerSession) FeeToSetter() (common.Address, error) {
	return _IPangolinFactory.Contract.FeeToSetter(&_IPangolinFactory.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_IPangolinFactory *IPangolinFactoryCaller) GetPair(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	var out []interface{}
	err := _IPangolinFactory.contract.Call(opts, &out, "getPair", tokenA, tokenB)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_IPangolinFactory *IPangolinFactorySession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _IPangolinFactory.Contract.GetPair(&_IPangolinFactory.CallOpts, tokenA, tokenB)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_IPangolinFactory *IPangolinFactoryCallerSession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _IPangolinFactory.Contract.GetPair(&_IPangolinFactory.CallOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_IPangolinFactory *IPangolinFactoryTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _IPangolinFactory.contract.Transact(opts, "createPair", tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_IPangolinFactory *IPangolinFactorySession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _IPangolinFactory.Contract.CreatePair(&_IPangolinFactory.TransactOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_IPangolinFactory *IPangolinFactoryTransactorSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _IPangolinFactory.Contract.CreatePair(&_IPangolinFactory.TransactOpts, tokenA, tokenB)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_IPangolinFactory *IPangolinFactoryTransactor) SetFeeTo(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IPangolinFactory.contract.Transact(opts, "setFeeTo", arg0)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_IPangolinFactory *IPangolinFactorySession) SetFeeTo(arg0 common.Address) (*types.Transaction, error) {
	return _IPangolinFactory.Contract.SetFeeTo(&_IPangolinFactory.TransactOpts, arg0)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_IPangolinFactory *IPangolinFactoryTransactorSession) SetFeeTo(arg0 common.Address) (*types.Transaction, error) {
	return _IPangolinFactory.Contract.SetFeeTo(&_IPangolinFactory.TransactOpts, arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_IPangolinFactory *IPangolinFactoryTransactor) SetFeeToSetter(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IPangolinFactory.contract.Transact(opts, "setFeeToSetter", arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_IPangolinFactory *IPangolinFactorySession) SetFeeToSetter(arg0 common.Address) (*types.Transaction, error) {
	return _IPangolinFactory.Contract.SetFeeToSetter(&_IPangolinFactory.TransactOpts, arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_IPangolinFactory *IPangolinFactoryTransactorSession) SetFeeToSetter(arg0 common.Address) (*types.Transaction, error) {
	return _IPangolinFactory.Contract.SetFeeToSetter(&_IPangolinFactory.TransactOpts, arg0)
}

// IPangolinFactoryPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the IPangolinFactory contract.
type IPangolinFactoryPairCreatedIterator struct {
	Event *IPangolinFactoryPairCreated // Event containing the contract specifics and raw log

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
func (it *IPangolinFactoryPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPangolinFactoryPairCreated)
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
		it.Event = new(IPangolinFactoryPairCreated)
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
func (it *IPangolinFactoryPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPangolinFactoryPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPangolinFactoryPairCreated represents a PairCreated event raised by the IPangolinFactory contract.
type IPangolinFactoryPairCreated struct {
	Token0 common.Address
	Token1 common.Address
	Pair   common.Address
	Arg3   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_IPangolinFactory *IPangolinFactoryFilterer) FilterPairCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*IPangolinFactoryPairCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _IPangolinFactory.contract.FilterLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &IPangolinFactoryPairCreatedIterator{contract: _IPangolinFactory.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_IPangolinFactory *IPangolinFactoryFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *IPangolinFactoryPairCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _IPangolinFactory.contract.WatchLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPangolinFactoryPairCreated)
				if err := _IPangolinFactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_IPangolinFactory *IPangolinFactoryFilterer) ParsePairCreated(log types.Log) (*IPangolinFactoryPairCreated, error) {
	event := new(IPangolinFactoryPairCreated)
	if err := _IPangolinFactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"PriceSetter__InsufficientPriceDifference\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceSetter__SymbolCantBeEmpty\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PriceChange\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MIN_DIF\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"priceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610392806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80638a42ebe914610046578063ab9a78df1461005b578063e7926e7b14610098575b600080fd5b6100596100543660046101cd565b6100a0565b005b61008661006936600461025b565b805160208183018101805160008252928201919093012091525481565b60405190815260200160405180910390f35b610086600281565b828260008190036100c4576040516303c8340b60e41b815260040160405180910390fd5b60008086866040516100d792919061030c565b9081526040519081900360200190205490506100f4600282610332565b6100fe82866101a1565b610109906064610332565b1161012757604051631fecbef960e01b815260040160405180910390fd5b836000878760405161013a92919061030c565b9081526040519081900360200181209190915561015a908790879061030c565b60408051918290038220868352426020840152917feff87fbb70cdf1275d2da00eaaa9ea0fdb7cbdd63cd5660e9e01ffc0a0243ec9910160405180910390a2505050505050565b6000818310156101ba576101b58383610349565b6101c4565b6101c48284610349565b90505b92915050565b6000806000604084860312156101e257600080fd5b833567ffffffffffffffff808211156101fa57600080fd5b818601915086601f83011261020e57600080fd5b81358181111561021d57600080fd5b87602082850101111561022f57600080fd5b6020928301989097509590910135949350505050565b634e487b7160e01b600052604160045260246000fd5b60006020828403121561026d57600080fd5b813567ffffffffffffffff8082111561028557600080fd5b818401915084601f83011261029957600080fd5b8135818111156102ab576102ab610245565b604051601f8201601f19908116603f011681019083821181831017156102d3576102d3610245565b816040528281528760208487010111156102ec57600080fd5b826020860160208301376000928101602001929092525095945050505050565b8183823760009101908152919050565b634e487b7160e01b600052601160045260246000fd5b80820281158282048414176101c7576101c761031c565b818103818111156101c7576101c761031c56fea2646970667358221220a0e4ba4694ea868058a7554ea4520698d38d772b82f6329571bca4e78ac50d3a64736f6c63430008110033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// MINDIF is a free data retrieval call binding the contract method 0xe7926e7b.
//
// Solidity: function MIN_DIF() view returns(uint256)
func (_Api *ApiCaller) MINDIF(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "MIN_DIF")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINDIF is a free data retrieval call binding the contract method 0xe7926e7b.
//
// Solidity: function MIN_DIF() view returns(uint256)
func (_Api *ApiSession) MINDIF() (*big.Int, error) {
	return _Api.Contract.MINDIF(&_Api.CallOpts)
}

// MINDIF is a free data retrieval call binding the contract method 0xe7926e7b.
//
// Solidity: function MIN_DIF() view returns(uint256)
func (_Api *ApiCallerSession) MINDIF() (*big.Int, error) {
	return _Api.Contract.MINDIF(&_Api.CallOpts)
}

// PriceOf is a free data retrieval call binding the contract method 0xab9a78df.
//
// Solidity: function priceOf(string ) view returns(uint256)
func (_Api *ApiCaller) PriceOf(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "priceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceOf is a free data retrieval call binding the contract method 0xab9a78df.
//
// Solidity: function priceOf(string ) view returns(uint256)
func (_Api *ApiSession) PriceOf(arg0 string) (*big.Int, error) {
	return _Api.Contract.PriceOf(&_Api.CallOpts, arg0)
}

// PriceOf is a free data retrieval call binding the contract method 0xab9a78df.
//
// Solidity: function priceOf(string ) view returns(uint256)
func (_Api *ApiCallerSession) PriceOf(arg0 string) (*big.Int, error) {
	return _Api.Contract.PriceOf(&_Api.CallOpts, arg0)
}

// Set is a paid mutator transaction binding the contract method 0x8a42ebe9.
//
// Solidity: function set(string _symbol, uint256 _price) returns()
func (_Api *ApiTransactor) Set(opts *bind.TransactOpts, _symbol string, _price *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "set", _symbol, _price)
}

// Set is a paid mutator transaction binding the contract method 0x8a42ebe9.
//
// Solidity: function set(string _symbol, uint256 _price) returns()
func (_Api *ApiSession) Set(_symbol string, _price *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Set(&_Api.TransactOpts, _symbol, _price)
}

// Set is a paid mutator transaction binding the contract method 0x8a42ebe9.
//
// Solidity: function set(string _symbol, uint256 _price) returns()
func (_Api *ApiTransactorSession) Set(_symbol string, _price *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Set(&_Api.TransactOpts, _symbol, _price)
}

// ApiPriceChangeIterator is returned from FilterPriceChange and is used to iterate over the raw logs and unpacked data for PriceChange events raised by the Api contract.
type ApiPriceChangeIterator struct {
	Event *ApiPriceChange // Event containing the contract specifics and raw log

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
func (it *ApiPriceChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiPriceChange)
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
		it.Event = new(ApiPriceChange)
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
func (it *ApiPriceChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiPriceChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiPriceChange represents a PriceChange event raised by the Api contract.
type ApiPriceChange struct {
	Symbol    common.Hash
	Price     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPriceChange is a free log retrieval operation binding the contract event 0xeff87fbb70cdf1275d2da00eaaa9ea0fdb7cbdd63cd5660e9e01ffc0a0243ec9.
//
// Solidity: event PriceChange(string indexed symbol, uint256 price, uint256 timestamp)
func (_Api *ApiFilterer) FilterPriceChange(opts *bind.FilterOpts, symbol []string) (*ApiPriceChangeIterator, error) {

	var symbolRule []interface{}
	for _, symbolItem := range symbol {
		symbolRule = append(symbolRule, symbolItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "PriceChange", symbolRule)
	if err != nil {
		return nil, err
	}
	return &ApiPriceChangeIterator{contract: _Api.contract, event: "PriceChange", logs: logs, sub: sub}, nil
}

// WatchPriceChange is a free log subscription operation binding the contract event 0xeff87fbb70cdf1275d2da00eaaa9ea0fdb7cbdd63cd5660e9e01ffc0a0243ec9.
//
// Solidity: event PriceChange(string indexed symbol, uint256 price, uint256 timestamp)
func (_Api *ApiFilterer) WatchPriceChange(opts *bind.WatchOpts, sink chan<- *ApiPriceChange, symbol []string) (event.Subscription, error) {

	var symbolRule []interface{}
	for _, symbolItem := range symbol {
		symbolRule = append(symbolRule, symbolItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "PriceChange", symbolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiPriceChange)
				if err := _Api.contract.UnpackLog(event, "PriceChange", log); err != nil {
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

// ParsePriceChange is a log parse operation binding the contract event 0xeff87fbb70cdf1275d2da00eaaa9ea0fdb7cbdd63cd5660e9e01ffc0a0243ec9.
//
// Solidity: event PriceChange(string indexed symbol, uint256 price, uint256 timestamp)
func (_Api *ApiFilterer) ParsePriceChange(log types.Log) (*ApiPriceChange, error) {
	event := new(ApiPriceChange)
	if err := _Api.contract.UnpackLog(event, "PriceChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

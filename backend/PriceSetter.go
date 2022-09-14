// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pricesetter

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

// PricesetterMetaData contains all meta data concerning the Pricesetter contract.
var PricesetterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"PriceSetter__InsufficientPriceDifference\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceSetter__SymbolCantBeEmpty\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"PriceChange\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MIN_DIF\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"priceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610633806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80638a42ebe914610046578063ab9a78df14610062578063e7926e7b14610092575b600080fd5b610060600480360381019061005b91906102fb565b6100b0565b005b61007c6004803603810190610077919061049c565b6101e9565b60405161008991906104f4565b60405180910390f35b61009a610217565b6040516100a791906104f4565b60405180910390f35b8282600082829050036100ef576040517f3c8340b000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080868660405161010292919061053f565b90815260200160405180910390205490506002816101209190610587565b606461012c838761021c565b6101369190610587565b1161016d576040517f1fecbef900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b836000878760405161018092919061053f565b90815260200160405180910390208190555085856040516101a292919061053f565b60405180910390207faf0b3c558c126d46d4e7d77be350be2aa80a5fcaa4513f906b2878d5058be6f3856040516101d991906104f4565b60405180910390a2505050505050565b6000818051602081018201805184825260208301602085012081835280955050505050506000915090505481565b600281565b60008183101561023757828261023291906105c9565b610244565b818361024391906105c9565b5b905092915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f84011261028557610284610260565b5b8235905067ffffffffffffffff8111156102a2576102a1610265565b5b6020830191508360018202830111156102be576102bd61026a565b5b9250929050565b6000819050919050565b6102d8816102c5565b81146102e357600080fd5b50565b6000813590506102f5816102cf565b92915050565b60008060006040848603121561031457610313610256565b5b600084013567ffffffffffffffff8111156103325761033161025b565b5b61033e8682870161026f565b93509350506020610351868287016102e6565b9150509250925092565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6103a982610360565b810181811067ffffffffffffffff821117156103c8576103c7610371565b5b80604052505050565b60006103db61024c565b90506103e782826103a0565b919050565b600067ffffffffffffffff82111561040757610406610371565b5b61041082610360565b9050602081019050919050565b82818337600083830152505050565b600061043f61043a846103ec565b6103d1565b90508281526020810184848401111561045b5761045a61035b565b5b61046684828561041d565b509392505050565b600082601f83011261048357610482610260565b5b813561049384826020860161042c565b91505092915050565b6000602082840312156104b2576104b1610256565b5b600082013567ffffffffffffffff8111156104d0576104cf61025b565b5b6104dc8482850161046e565b91505092915050565b6104ee816102c5565b82525050565b600060208201905061050960008301846104e5565b92915050565b600081905092915050565b6000610526838561050f565b935061053383858461041d565b82840190509392505050565b600061054c82848661051a565b91508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610592826102c5565b915061059d836102c5565b92508282026105ab816102c5565b915082820484148315176105c2576105c1610558565b5b5092915050565b60006105d4826102c5565b91506105df836102c5565b92508282039050818111156105f7576105f6610558565b5b9291505056fea26469706673582212201027e5c9051ca211dd6b685ad066303af83b103e2657f135854a94ba2f06eba064736f6c63430008110033",
}

// PricesetterABI is the input ABI used to generate the binding from.
// Deprecated: Use PricesetterMetaData.ABI instead.
var PricesetterABI = PricesetterMetaData.ABI

// PricesetterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PricesetterMetaData.Bin instead.
var PricesetterBin = PricesetterMetaData.Bin

// DeployPricesetter deploys a new Ethereum contract, binding an instance of Pricesetter to it.
func DeployPricesetter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pricesetter, error) {
	parsed, err := PricesetterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PricesetterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pricesetter{PricesetterCaller: PricesetterCaller{contract: contract}, PricesetterTransactor: PricesetterTransactor{contract: contract}, PricesetterFilterer: PricesetterFilterer{contract: contract}}, nil
}

// Pricesetter is an auto generated Go binding around an Ethereum contract.
type Pricesetter struct {
	PricesetterCaller     // Read-only binding to the contract
	PricesetterTransactor // Write-only binding to the contract
	PricesetterFilterer   // Log filterer for contract events
}

// PricesetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type PricesetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PricesetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PricesetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PricesetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PricesetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PricesetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PricesetterSession struct {
	Contract     *Pricesetter      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PricesetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PricesetterCallerSession struct {
	Contract *PricesetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PricesetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PricesetterTransactorSession struct {
	Contract     *PricesetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PricesetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type PricesetterRaw struct {
	Contract *Pricesetter // Generic contract binding to access the raw methods on
}

// PricesetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PricesetterCallerRaw struct {
	Contract *PricesetterCaller // Generic read-only contract binding to access the raw methods on
}

// PricesetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PricesetterTransactorRaw struct {
	Contract *PricesetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPricesetter creates a new instance of Pricesetter, bound to a specific deployed contract.
func NewPricesetter(address common.Address, backend bind.ContractBackend) (*Pricesetter, error) {
	contract, err := bindPricesetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pricesetter{PricesetterCaller: PricesetterCaller{contract: contract}, PricesetterTransactor: PricesetterTransactor{contract: contract}, PricesetterFilterer: PricesetterFilterer{contract: contract}}, nil
}

// NewPricesetterCaller creates a new read-only instance of Pricesetter, bound to a specific deployed contract.
func NewPricesetterCaller(address common.Address, caller bind.ContractCaller) (*PricesetterCaller, error) {
	contract, err := bindPricesetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PricesetterCaller{contract: contract}, nil
}

// NewPricesetterTransactor creates a new write-only instance of Pricesetter, bound to a specific deployed contract.
func NewPricesetterTransactor(address common.Address, transactor bind.ContractTransactor) (*PricesetterTransactor, error) {
	contract, err := bindPricesetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PricesetterTransactor{contract: contract}, nil
}

// NewPricesetterFilterer creates a new log filterer instance of Pricesetter, bound to a specific deployed contract.
func NewPricesetterFilterer(address common.Address, filterer bind.ContractFilterer) (*PricesetterFilterer, error) {
	contract, err := bindPricesetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PricesetterFilterer{contract: contract}, nil
}

// bindPricesetter binds a generic wrapper to an already deployed contract.
func bindPricesetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PricesetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pricesetter *PricesetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pricesetter.Contract.PricesetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pricesetter *PricesetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pricesetter.Contract.PricesetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pricesetter *PricesetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pricesetter.Contract.PricesetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pricesetter *PricesetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pricesetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pricesetter *PricesetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pricesetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pricesetter *PricesetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pricesetter.Contract.contract.Transact(opts, method, params...)
}

// MINDIF is a free data retrieval call binding the contract method 0xe7926e7b.
//
// Solidity: function MIN_DIF() view returns(uint256)
func (_Pricesetter *PricesetterCaller) MINDIF(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pricesetter.contract.Call(opts, &out, "MIN_DIF")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINDIF is a free data retrieval call binding the contract method 0xe7926e7b.
//
// Solidity: function MIN_DIF() view returns(uint256)
func (_Pricesetter *PricesetterSession) MINDIF() (*big.Int, error) {
	return _Pricesetter.Contract.MINDIF(&_Pricesetter.CallOpts)
}

// MINDIF is a free data retrieval call binding the contract method 0xe7926e7b.
//
// Solidity: function MIN_DIF() view returns(uint256)
func (_Pricesetter *PricesetterCallerSession) MINDIF() (*big.Int, error) {
	return _Pricesetter.Contract.MINDIF(&_Pricesetter.CallOpts)
}

// PriceOf is a free data retrieval call binding the contract method 0xab9a78df.
//
// Solidity: function priceOf(string ) view returns(uint256)
func (_Pricesetter *PricesetterCaller) PriceOf(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Pricesetter.contract.Call(opts, &out, "priceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceOf is a free data retrieval call binding the contract method 0xab9a78df.
//
// Solidity: function priceOf(string ) view returns(uint256)
func (_Pricesetter *PricesetterSession) PriceOf(arg0 string) (*big.Int, error) {
	return _Pricesetter.Contract.PriceOf(&_Pricesetter.CallOpts, arg0)
}

// PriceOf is a free data retrieval call binding the contract method 0xab9a78df.
//
// Solidity: function priceOf(string ) view returns(uint256)
func (_Pricesetter *PricesetterCallerSession) PriceOf(arg0 string) (*big.Int, error) {
	return _Pricesetter.Contract.PriceOf(&_Pricesetter.CallOpts, arg0)
}

// Set is a paid mutator transaction binding the contract method 0x8a42ebe9.
//
// Solidity: function set(string _symbol, uint256 _price) returns()
func (_Pricesetter *PricesetterTransactor) Set(opts *bind.TransactOpts, _symbol string, _price *big.Int) (*types.Transaction, error) {
	return _Pricesetter.contract.Transact(opts, "set", _symbol, _price)
}

// Set is a paid mutator transaction binding the contract method 0x8a42ebe9.
//
// Solidity: function set(string _symbol, uint256 _price) returns()
func (_Pricesetter *PricesetterSession) Set(_symbol string, _price *big.Int) (*types.Transaction, error) {
	return _Pricesetter.Contract.Set(&_Pricesetter.TransactOpts, _symbol, _price)
}

// Set is a paid mutator transaction binding the contract method 0x8a42ebe9.
//
// Solidity: function set(string _symbol, uint256 _price) returns()
func (_Pricesetter *PricesetterTransactorSession) Set(_symbol string, _price *big.Int) (*types.Transaction, error) {
	return _Pricesetter.Contract.Set(&_Pricesetter.TransactOpts, _symbol, _price)
}

// PricesetterPriceChangeIterator is returned from FilterPriceChange and is used to iterate over the raw logs and unpacked data for PriceChange events raised by the Pricesetter contract.
type PricesetterPriceChangeIterator struct {
	Event *PricesetterPriceChange // Event containing the contract specifics and raw log

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
func (it *PricesetterPriceChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PricesetterPriceChange)
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
		it.Event = new(PricesetterPriceChange)
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
func (it *PricesetterPriceChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PricesetterPriceChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PricesetterPriceChange represents a PriceChange event raised by the Pricesetter contract.
type PricesetterPriceChange struct {
	Symbol common.Hash
	Price  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPriceChange is a free log retrieval operation binding the contract event 0xaf0b3c558c126d46d4e7d77be350be2aa80a5fcaa4513f906b2878d5058be6f3.
//
// Solidity: event PriceChange(string indexed symbol, uint256 price)
func (_Pricesetter *PricesetterFilterer) FilterPriceChange(opts *bind.FilterOpts, symbol []string) (*PricesetterPriceChangeIterator, error) {

	var symbolRule []interface{}
	for _, symbolItem := range symbol {
		symbolRule = append(symbolRule, symbolItem)
	}

	logs, sub, err := _Pricesetter.contract.FilterLogs(opts, "PriceChange", symbolRule)
	if err != nil {
		return nil, err
	}
	return &PricesetterPriceChangeIterator{contract: _Pricesetter.contract, event: "PriceChange", logs: logs, sub: sub}, nil
}

// WatchPriceChange is a free log subscription operation binding the contract event 0xaf0b3c558c126d46d4e7d77be350be2aa80a5fcaa4513f906b2878d5058be6f3.
//
// Solidity: event PriceChange(string indexed symbol, uint256 price)
func (_Pricesetter *PricesetterFilterer) WatchPriceChange(opts *bind.WatchOpts, sink chan<- *PricesetterPriceChange, symbol []string) (event.Subscription, error) {

	var symbolRule []interface{}
	for _, symbolItem := range symbol {
		symbolRule = append(symbolRule, symbolItem)
	}

	logs, sub, err := _Pricesetter.contract.WatchLogs(opts, "PriceChange", symbolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PricesetterPriceChange)
				if err := _Pricesetter.contract.UnpackLog(event, "PriceChange", log); err != nil {
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

// ParsePriceChange is a log parse operation binding the contract event 0xaf0b3c558c126d46d4e7d77be350be2aa80a5fcaa4513f906b2878d5058be6f3.
//
// Solidity: event PriceChange(string indexed symbol, uint256 price)
func (_Pricesetter *PricesetterFilterer) ParsePriceChange(log types.Log) (*PricesetterPriceChange, error) {
	event := new(PricesetterPriceChange)
	if err := _Pricesetter.contract.UnpackLog(event, "PriceChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

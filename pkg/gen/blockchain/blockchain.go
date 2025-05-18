// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockchain

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

// CompetitionRegistryMetaData contains all meta data concerning the CompetitionRegistry contract.
var CompetitionRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"CompetitionRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"recordCompetition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recorded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"validateCompetition\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061058c8061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c8063501895ae1461004e57806360b2267b14610077578063a098e6ae146100a9578063a91c2213146100bc575b5f5ffd5b61006161005c366004610293565b6100d1565b60405161006e91906102aa565b60405180910390f35b610099610085366004610293565b5f6020819052908152604090205460ff1681565b604051901515815260200161006e565b6100996100b73660046102df565b610168565b6100cf6100ca3660046102df565b6101c5565b005b60016020525f9081526040902080546100e990610356565b80601f016020809104026020016040519081016040528092919081815260200182805461011590610356565b80156101605780601f1061013757610100808354040283529160200191610160565b820191905f5260205f20905b81548152906001019060200180831161014357829003601f168201915b505050505081565b5f8381526020819052604081205460ff1680156101bd5750828260405161019092919061038e565b604080519182900382205f87815260016020529190912090916101b3919061039d565b6040518091039020145b949350505050565b5f8381526020819052604090205460ff16156102275760405162461bcd60e51b815260206004820152601c60248201527f436f6d7065746974696f6e20616c7265616479207265636f7264656400000000604482015260640160405180910390fd5b5f83815260208181526040808320805460ff19166001908117909155909152902061025382848361046e565b50827f1c05c0829b9e8a2ee03538ba5c7cc384bde635f39b98fb6e99b182ff124a44728383604051610286929190610528565b60405180910390a2505050565b5f602082840312156102a3575f5ffd5b5035919050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f5f5f604084860312156102f1575f5ffd5b83359250602084013567ffffffffffffffff81111561030e575f5ffd5b8401601f8101861361031e575f5ffd5b803567ffffffffffffffff811115610334575f5ffd5b866020828401011115610345575f5ffd5b939660209190910195509293505050565b600181811c9082168061036a57607f821691505b60208210810361038857634e487b7160e01b5f52602260045260245ffd5b50919050565b818382375f9101908152919050565b5f5f83546103aa81610356565b6001821680156103c157600181146103d657610403565b60ff1983168652811515820286019350610403565b865f5260205f205f5b838110156103fb578154888201526001909101906020016103df565b505081860193505b509195945050505050565b634e487b7160e01b5f52604160045260245ffd5b601f82111561046957805f5260205f20601f840160051c810160208510156104475750805b601f840160051c820191505b81811015610466575f8155600101610453565b50505b505050565b67ffffffffffffffff8311156104865761048661040e565b61049a836104948354610356565b83610422565b5f601f8411600181146104cb575f85156104b45750838201355b5f19600387901b1c1916600186901b178355610466565b5f83815260208120601f198716915b828110156104fa57868501358255602094850194600190920191016104da565b5086821015610516575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b60208152816020820152818360408301375f818301604090810191909152601f909201601f1916010191905056fea264697066735822122083d3e1d40a2300ee9daf08ec1c80e1ecafb514040f353a905235532bea66866764736f6c634300081e0033",
}

// CompetitionRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use CompetitionRegistryMetaData.ABI instead.
var CompetitionRegistryABI = CompetitionRegistryMetaData.ABI

// CompetitionRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CompetitionRegistryMetaData.Bin instead.
var CompetitionRegistryBin = CompetitionRegistryMetaData.Bin

// DeployCompetitionRegistry deploys a new Ethereum contract, binding an instance of CompetitionRegistry to it.
func DeployCompetitionRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CompetitionRegistry, error) {
	parsed, err := CompetitionRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CompetitionRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CompetitionRegistry{CompetitionRegistryCaller: CompetitionRegistryCaller{contract: contract}, CompetitionRegistryTransactor: CompetitionRegistryTransactor{contract: contract}, CompetitionRegistryFilterer: CompetitionRegistryFilterer{contract: contract}}, nil
}

// CompetitionRegistry is an auto generated Go binding around an Ethereum contract.
type CompetitionRegistry struct {
	CompetitionRegistryCaller     // Read-only binding to the contract
	CompetitionRegistryTransactor // Write-only binding to the contract
	CompetitionRegistryFilterer   // Log filterer for contract events
}

// CompetitionRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompetitionRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompetitionRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompetitionRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompetitionRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompetitionRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompetitionRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompetitionRegistrySession struct {
	Contract     *CompetitionRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CompetitionRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompetitionRegistryCallerSession struct {
	Contract *CompetitionRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// CompetitionRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompetitionRegistryTransactorSession struct {
	Contract     *CompetitionRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// CompetitionRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompetitionRegistryRaw struct {
	Contract *CompetitionRegistry // Generic contract binding to access the raw methods on
}

// CompetitionRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompetitionRegistryCallerRaw struct {
	Contract *CompetitionRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// CompetitionRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompetitionRegistryTransactorRaw struct {
	Contract *CompetitionRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompetitionRegistry creates a new instance of CompetitionRegistry, bound to a specific deployed contract.
func NewCompetitionRegistry(address common.Address, backend bind.ContractBackend) (*CompetitionRegistry, error) {
	contract, err := bindCompetitionRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CompetitionRegistry{CompetitionRegistryCaller: CompetitionRegistryCaller{contract: contract}, CompetitionRegistryTransactor: CompetitionRegistryTransactor{contract: contract}, CompetitionRegistryFilterer: CompetitionRegistryFilterer{contract: contract}}, nil
}

// NewCompetitionRegistryCaller creates a new read-only instance of CompetitionRegistry, bound to a specific deployed contract.
func NewCompetitionRegistryCaller(address common.Address, caller bind.ContractCaller) (*CompetitionRegistryCaller, error) {
	contract, err := bindCompetitionRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompetitionRegistryCaller{contract: contract}, nil
}

// NewCompetitionRegistryTransactor creates a new write-only instance of CompetitionRegistry, bound to a specific deployed contract.
func NewCompetitionRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*CompetitionRegistryTransactor, error) {
	contract, err := bindCompetitionRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompetitionRegistryTransactor{contract: contract}, nil
}

// NewCompetitionRegistryFilterer creates a new log filterer instance of CompetitionRegistry, bound to a specific deployed contract.
func NewCompetitionRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*CompetitionRegistryFilterer, error) {
	contract, err := bindCompetitionRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompetitionRegistryFilterer{contract: contract}, nil
}

// bindCompetitionRegistry binds a generic wrapper to an already deployed contract.
func bindCompetitionRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CompetitionRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompetitionRegistry *CompetitionRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompetitionRegistry.Contract.CompetitionRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompetitionRegistry *CompetitionRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompetitionRegistry.Contract.CompetitionRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompetitionRegistry *CompetitionRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompetitionRegistry.Contract.CompetitionRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompetitionRegistry *CompetitionRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompetitionRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompetitionRegistry *CompetitionRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompetitionRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompetitionRegistry *CompetitionRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompetitionRegistry.Contract.contract.Transact(opts, method, params...)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_CompetitionRegistry *CompetitionRegistryCaller) Hashes(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _CompetitionRegistry.contract.Call(opts, &out, "hashes", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_CompetitionRegistry *CompetitionRegistrySession) Hashes(arg0 *big.Int) (string, error) {
	return _CompetitionRegistry.Contract.Hashes(&_CompetitionRegistry.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_CompetitionRegistry *CompetitionRegistryCallerSession) Hashes(arg0 *big.Int) (string, error) {
	return _CompetitionRegistry.Contract.Hashes(&_CompetitionRegistry.CallOpts, arg0)
}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_CompetitionRegistry *CompetitionRegistryCaller) Recorded(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _CompetitionRegistry.contract.Call(opts, &out, "recorded", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_CompetitionRegistry *CompetitionRegistrySession) Recorded(arg0 *big.Int) (bool, error) {
	return _CompetitionRegistry.Contract.Recorded(&_CompetitionRegistry.CallOpts, arg0)
}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_CompetitionRegistry *CompetitionRegistryCallerSession) Recorded(arg0 *big.Int) (bool, error) {
	return _CompetitionRegistry.Contract.Recorded(&_CompetitionRegistry.CallOpts, arg0)
}

// ValidateCompetition is a free data retrieval call binding the contract method 0xa098e6ae.
//
// Solidity: function validateCompetition(uint256 id, string recordHash) view returns(bool)
func (_CompetitionRegistry *CompetitionRegistryCaller) ValidateCompetition(opts *bind.CallOpts, id *big.Int, recordHash string) (bool, error) {
	var out []interface{}
	err := _CompetitionRegistry.contract.Call(opts, &out, "validateCompetition", id, recordHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateCompetition is a free data retrieval call binding the contract method 0xa098e6ae.
//
// Solidity: function validateCompetition(uint256 id, string recordHash) view returns(bool)
func (_CompetitionRegistry *CompetitionRegistrySession) ValidateCompetition(id *big.Int, recordHash string) (bool, error) {
	return _CompetitionRegistry.Contract.ValidateCompetition(&_CompetitionRegistry.CallOpts, id, recordHash)
}

// ValidateCompetition is a free data retrieval call binding the contract method 0xa098e6ae.
//
// Solidity: function validateCompetition(uint256 id, string recordHash) view returns(bool)
func (_CompetitionRegistry *CompetitionRegistryCallerSession) ValidateCompetition(id *big.Int, recordHash string) (bool, error) {
	return _CompetitionRegistry.Contract.ValidateCompetition(&_CompetitionRegistry.CallOpts, id, recordHash)
}

// RecordCompetition is a paid mutator transaction binding the contract method 0xa91c2213.
//
// Solidity: function recordCompetition(uint256 id, string recordHash) returns()
func (_CompetitionRegistry *CompetitionRegistryTransactor) RecordCompetition(opts *bind.TransactOpts, id *big.Int, recordHash string) (*types.Transaction, error) {
	return _CompetitionRegistry.contract.Transact(opts, "recordCompetition", id, recordHash)
}

// RecordCompetition is a paid mutator transaction binding the contract method 0xa91c2213.
//
// Solidity: function recordCompetition(uint256 id, string recordHash) returns()
func (_CompetitionRegistry *CompetitionRegistrySession) RecordCompetition(id *big.Int, recordHash string) (*types.Transaction, error) {
	return _CompetitionRegistry.Contract.RecordCompetition(&_CompetitionRegistry.TransactOpts, id, recordHash)
}

// RecordCompetition is a paid mutator transaction binding the contract method 0xa91c2213.
//
// Solidity: function recordCompetition(uint256 id, string recordHash) returns()
func (_CompetitionRegistry *CompetitionRegistryTransactorSession) RecordCompetition(id *big.Int, recordHash string) (*types.Transaction, error) {
	return _CompetitionRegistry.Contract.RecordCompetition(&_CompetitionRegistry.TransactOpts, id, recordHash)
}

// CompetitionRegistryCompetitionRecordedIterator is returned from FilterCompetitionRecorded and is used to iterate over the raw logs and unpacked data for CompetitionRecorded events raised by the CompetitionRegistry contract.
type CompetitionRegistryCompetitionRecordedIterator struct {
	Event *CompetitionRegistryCompetitionRecorded // Event containing the contract specifics and raw log

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
func (it *CompetitionRegistryCompetitionRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompetitionRegistryCompetitionRecorded)
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
		it.Event = new(CompetitionRegistryCompetitionRecorded)
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
func (it *CompetitionRegistryCompetitionRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompetitionRegistryCompetitionRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompetitionRegistryCompetitionRecorded represents a CompetitionRecorded event raised by the CompetitionRegistry contract.
type CompetitionRegistryCompetitionRecorded struct {
	Id         *big.Int
	RecordHash string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCompetitionRecorded is a free log retrieval operation binding the contract event 0x1c05c0829b9e8a2ee03538ba5c7cc384bde635f39b98fb6e99b182ff124a4472.
//
// Solidity: event CompetitionRecorded(uint256 indexed id, string recordHash)
func (_CompetitionRegistry *CompetitionRegistryFilterer) FilterCompetitionRecorded(opts *bind.FilterOpts, id []*big.Int) (*CompetitionRegistryCompetitionRecordedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _CompetitionRegistry.contract.FilterLogs(opts, "CompetitionRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return &CompetitionRegistryCompetitionRecordedIterator{contract: _CompetitionRegistry.contract, event: "CompetitionRecorded", logs: logs, sub: sub}, nil
}

// WatchCompetitionRecorded is a free log subscription operation binding the contract event 0x1c05c0829b9e8a2ee03538ba5c7cc384bde635f39b98fb6e99b182ff124a4472.
//
// Solidity: event CompetitionRecorded(uint256 indexed id, string recordHash)
func (_CompetitionRegistry *CompetitionRegistryFilterer) WatchCompetitionRecorded(opts *bind.WatchOpts, sink chan<- *CompetitionRegistryCompetitionRecorded, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _CompetitionRegistry.contract.WatchLogs(opts, "CompetitionRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompetitionRegistryCompetitionRecorded)
				if err := _CompetitionRegistry.contract.UnpackLog(event, "CompetitionRecorded", log); err != nil {
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

// ParseCompetitionRecorded is a log parse operation binding the contract event 0x1c05c0829b9e8a2ee03538ba5c7cc384bde635f39b98fb6e99b182ff124a4472.
//
// Solidity: event CompetitionRecorded(uint256 indexed id, string recordHash)
func (_CompetitionRegistry *CompetitionRegistryFilterer) ParseCompetitionRecorded(log types.Log) (*CompetitionRegistryCompetitionRecorded, error) {
	event := new(CompetitionRegistryCompetitionRecorded)
	if err := _CompetitionRegistry.contract.UnpackLog(event, "CompetitionRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompetitionTeamsRegistryMetaData contains all meta data concerning the CompetitionTeamsRegistry contract.
var CompetitionTeamsRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"CompetitionTeamsRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"recordCompetitionTeams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recorded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"validateCompetitionTeams\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506105968061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c8063501895ae1461004e57806360b2267b14610077578063b7a549e0146100a9578063e5fe541e146100be575b5f5ffd5b61006161005c36600461029d565b6100d1565b60405161006e91906102b4565b60405180910390f35b61009961008536600461029d565b5f6020819052908152604090205460ff1681565b604051901515815260200161006e565b6100bc6100b73660046102e9565b610168565b005b6100996100cc3660046102e9565b610240565b60016020525f9081526040902080546100e990610360565b80601f016020809104026020016040519081016040528092919081815260200182805461011590610360565b80156101605780601f1061013757610100808354040283529160200191610160565b820191905f5260205f20905b81548152906001019060200180831161014357829003601f168201915b505050505081565b5f8381526020819052604090205460ff16156101d45760405162461bcd60e51b815260206004820152602160248201527f436f6d7065746974696f6e5465616d7320616c7265616479207265636f7264656044820152601960fa1b606482015260840160405180910390fd5b5f83815260208181526040808320805460ff1916600190811790915590915290206102008284836103f8565b50827fb5afd97b4f740b109d4828fcda8584f7e87aebb685df30a0aae5f00b7659114883836040516102339291906104b2565b60405180910390a2505050565b5f8381526020819052604081205460ff168015610295575082826040516102689291906104e0565b604080519182900382205f878152600160205291909120909161028b91906104ef565b6040518091039020145b949350505050565b5f602082840312156102ad575f5ffd5b5035919050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f5f5f604084860312156102fb575f5ffd5b83359250602084013567ffffffffffffffff811115610318575f5ffd5b8401601f81018613610328575f5ffd5b803567ffffffffffffffff81111561033e575f5ffd5b86602082840101111561034f575f5ffd5b939660209190910195509293505050565b600181811c9082168061037457607f821691505b60208210810361039257634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52604160045260245ffd5b601f8211156103f357805f5260205f20601f840160051c810160208510156103d15750805b601f840160051c820191505b818110156103f0575f81556001016103dd565b50505b505050565b67ffffffffffffffff83111561041057610410610398565b6104248361041e8354610360565b836103ac565b5f601f841160018114610455575f851561043e5750838201355b5f19600387901b1c1916600186901b1783556103f0565b5f83815260208120601f198716915b828110156104845786850135825560209485019460019092019101610464565b50868210156104a0575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b60208152816020820152818360408301375f818301604090810191909152601f909201601f19160101919050565b818382375f9101908152919050565b5f5f83546104fc81610360565b600182168015610513576001811461052857610555565b60ff1983168652811515820286019350610555565b865f5260205f205f5b8381101561054d57815488820152600190910190602001610531565b505081860193505b50919594505050505056fea2646970667358221220c54993009f71e17dc48bc9431909c0de87f9a0b18448e272a2d93018316730e064736f6c634300081e0033",
}

// CompetitionTeamsRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use CompetitionTeamsRegistryMetaData.ABI instead.
var CompetitionTeamsRegistryABI = CompetitionTeamsRegistryMetaData.ABI

// CompetitionTeamsRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CompetitionTeamsRegistryMetaData.Bin instead.
var CompetitionTeamsRegistryBin = CompetitionTeamsRegistryMetaData.Bin

// DeployCompetitionTeamsRegistry deploys a new Ethereum contract, binding an instance of CompetitionTeamsRegistry to it.
func DeployCompetitionTeamsRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CompetitionTeamsRegistry, error) {
	parsed, err := CompetitionTeamsRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CompetitionTeamsRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CompetitionTeamsRegistry{CompetitionTeamsRegistryCaller: CompetitionTeamsRegistryCaller{contract: contract}, CompetitionTeamsRegistryTransactor: CompetitionTeamsRegistryTransactor{contract: contract}, CompetitionTeamsRegistryFilterer: CompetitionTeamsRegistryFilterer{contract: contract}}, nil
}

// CompetitionTeamsRegistry is an auto generated Go binding around an Ethereum contract.
type CompetitionTeamsRegistry struct {
	CompetitionTeamsRegistryCaller     // Read-only binding to the contract
	CompetitionTeamsRegistryTransactor // Write-only binding to the contract
	CompetitionTeamsRegistryFilterer   // Log filterer for contract events
}

// CompetitionTeamsRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompetitionTeamsRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompetitionTeamsRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompetitionTeamsRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompetitionTeamsRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompetitionTeamsRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompetitionTeamsRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompetitionTeamsRegistrySession struct {
	Contract     *CompetitionTeamsRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CompetitionTeamsRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompetitionTeamsRegistryCallerSession struct {
	Contract *CompetitionTeamsRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// CompetitionTeamsRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompetitionTeamsRegistryTransactorSession struct {
	Contract     *CompetitionTeamsRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// CompetitionTeamsRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompetitionTeamsRegistryRaw struct {
	Contract *CompetitionTeamsRegistry // Generic contract binding to access the raw methods on
}

// CompetitionTeamsRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompetitionTeamsRegistryCallerRaw struct {
	Contract *CompetitionTeamsRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// CompetitionTeamsRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompetitionTeamsRegistryTransactorRaw struct {
	Contract *CompetitionTeamsRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompetitionTeamsRegistry creates a new instance of CompetitionTeamsRegistry, bound to a specific deployed contract.
func NewCompetitionTeamsRegistry(address common.Address, backend bind.ContractBackend) (*CompetitionTeamsRegistry, error) {
	contract, err := bindCompetitionTeamsRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CompetitionTeamsRegistry{CompetitionTeamsRegistryCaller: CompetitionTeamsRegistryCaller{contract: contract}, CompetitionTeamsRegistryTransactor: CompetitionTeamsRegistryTransactor{contract: contract}, CompetitionTeamsRegistryFilterer: CompetitionTeamsRegistryFilterer{contract: contract}}, nil
}

// NewCompetitionTeamsRegistryCaller creates a new read-only instance of CompetitionTeamsRegistry, bound to a specific deployed contract.
func NewCompetitionTeamsRegistryCaller(address common.Address, caller bind.ContractCaller) (*CompetitionTeamsRegistryCaller, error) {
	contract, err := bindCompetitionTeamsRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompetitionTeamsRegistryCaller{contract: contract}, nil
}

// NewCompetitionTeamsRegistryTransactor creates a new write-only instance of CompetitionTeamsRegistry, bound to a specific deployed contract.
func NewCompetitionTeamsRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*CompetitionTeamsRegistryTransactor, error) {
	contract, err := bindCompetitionTeamsRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompetitionTeamsRegistryTransactor{contract: contract}, nil
}

// NewCompetitionTeamsRegistryFilterer creates a new log filterer instance of CompetitionTeamsRegistry, bound to a specific deployed contract.
func NewCompetitionTeamsRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*CompetitionTeamsRegistryFilterer, error) {
	contract, err := bindCompetitionTeamsRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompetitionTeamsRegistryFilterer{contract: contract}, nil
}

// bindCompetitionTeamsRegistry binds a generic wrapper to an already deployed contract.
func bindCompetitionTeamsRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CompetitionTeamsRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompetitionTeamsRegistry.Contract.CompetitionTeamsRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompetitionTeamsRegistry.Contract.CompetitionTeamsRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompetitionTeamsRegistry.Contract.CompetitionTeamsRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompetitionTeamsRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompetitionTeamsRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompetitionTeamsRegistry.Contract.contract.Transact(opts, method, params...)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistryCaller) Hashes(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _CompetitionTeamsRegistry.contract.Call(opts, &out, "hashes", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistrySession) Hashes(arg0 *big.Int) (string, error) {
	return _CompetitionTeamsRegistry.Contract.Hashes(&_CompetitionTeamsRegistry.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistryCallerSession) Hashes(arg0 *big.Int) (string, error) {
	return _CompetitionTeamsRegistry.Contract.Hashes(&_CompetitionTeamsRegistry.CallOpts, arg0)
}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistryCaller) Recorded(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _CompetitionTeamsRegistry.contract.Call(opts, &out, "recorded", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistrySession) Recorded(arg0 *big.Int) (bool, error) {
	return _CompetitionTeamsRegistry.Contract.Recorded(&_CompetitionTeamsRegistry.CallOpts, arg0)
}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistryCallerSession) Recorded(arg0 *big.Int) (bool, error) {
	return _CompetitionTeamsRegistry.Contract.Recorded(&_CompetitionTeamsRegistry.CallOpts, arg0)
}

// ValidateCompetitionTeams is a free data retrieval call binding the contract method 0xe5fe541e.
//
// Solidity: function validateCompetitionTeams(uint256 id, string recordHash) view returns(bool)
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistryCaller) ValidateCompetitionTeams(opts *bind.CallOpts, id *big.Int, recordHash string) (bool, error) {
	var out []interface{}
	err := _CompetitionTeamsRegistry.contract.Call(opts, &out, "validateCompetitionTeams", id, recordHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateCompetitionTeams is a free data retrieval call binding the contract method 0xe5fe541e.
//
// Solidity: function validateCompetitionTeams(uint256 id, string recordHash) view returns(bool)
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistrySession) ValidateCompetitionTeams(id *big.Int, recordHash string) (bool, error) {
	return _CompetitionTeamsRegistry.Contract.ValidateCompetitionTeams(&_CompetitionTeamsRegistry.CallOpts, id, recordHash)
}

// ValidateCompetitionTeams is a free data retrieval call binding the contract method 0xe5fe541e.
//
// Solidity: function validateCompetitionTeams(uint256 id, string recordHash) view returns(bool)
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistryCallerSession) ValidateCompetitionTeams(id *big.Int, recordHash string) (bool, error) {
	return _CompetitionTeamsRegistry.Contract.ValidateCompetitionTeams(&_CompetitionTeamsRegistry.CallOpts, id, recordHash)
}

// RecordCompetitionTeams is a paid mutator transaction binding the contract method 0xb7a549e0.
//
// Solidity: function recordCompetitionTeams(uint256 id, string recordHash) returns()
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistryTransactor) RecordCompetitionTeams(opts *bind.TransactOpts, id *big.Int, recordHash string) (*types.Transaction, error) {
	return _CompetitionTeamsRegistry.contract.Transact(opts, "recordCompetitionTeams", id, recordHash)
}

// RecordCompetitionTeams is a paid mutator transaction binding the contract method 0xb7a549e0.
//
// Solidity: function recordCompetitionTeams(uint256 id, string recordHash) returns()
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistrySession) RecordCompetitionTeams(id *big.Int, recordHash string) (*types.Transaction, error) {
	return _CompetitionTeamsRegistry.Contract.RecordCompetitionTeams(&_CompetitionTeamsRegistry.TransactOpts, id, recordHash)
}

// RecordCompetitionTeams is a paid mutator transaction binding the contract method 0xb7a549e0.
//
// Solidity: function recordCompetitionTeams(uint256 id, string recordHash) returns()
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistryTransactorSession) RecordCompetitionTeams(id *big.Int, recordHash string) (*types.Transaction, error) {
	return _CompetitionTeamsRegistry.Contract.RecordCompetitionTeams(&_CompetitionTeamsRegistry.TransactOpts, id, recordHash)
}

// CompetitionTeamsRegistryCompetitionTeamsRecordedIterator is returned from FilterCompetitionTeamsRecorded and is used to iterate over the raw logs and unpacked data for CompetitionTeamsRecorded events raised by the CompetitionTeamsRegistry contract.
type CompetitionTeamsRegistryCompetitionTeamsRecordedIterator struct {
	Event *CompetitionTeamsRegistryCompetitionTeamsRecorded // Event containing the contract specifics and raw log

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
func (it *CompetitionTeamsRegistryCompetitionTeamsRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompetitionTeamsRegistryCompetitionTeamsRecorded)
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
		it.Event = new(CompetitionTeamsRegistryCompetitionTeamsRecorded)
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
func (it *CompetitionTeamsRegistryCompetitionTeamsRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompetitionTeamsRegistryCompetitionTeamsRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompetitionTeamsRegistryCompetitionTeamsRecorded represents a CompetitionTeamsRecorded event raised by the CompetitionTeamsRegistry contract.
type CompetitionTeamsRegistryCompetitionTeamsRecorded struct {
	Id         *big.Int
	RecordHash string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCompetitionTeamsRecorded is a free log retrieval operation binding the contract event 0xb5afd97b4f740b109d4828fcda8584f7e87aebb685df30a0aae5f00b76591148.
//
// Solidity: event CompetitionTeamsRecorded(uint256 indexed id, string recordHash)
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistryFilterer) FilterCompetitionTeamsRecorded(opts *bind.FilterOpts, id []*big.Int) (*CompetitionTeamsRegistryCompetitionTeamsRecordedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _CompetitionTeamsRegistry.contract.FilterLogs(opts, "CompetitionTeamsRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return &CompetitionTeamsRegistryCompetitionTeamsRecordedIterator{contract: _CompetitionTeamsRegistry.contract, event: "CompetitionTeamsRecorded", logs: logs, sub: sub}, nil
}

// WatchCompetitionTeamsRecorded is a free log subscription operation binding the contract event 0xb5afd97b4f740b109d4828fcda8584f7e87aebb685df30a0aae5f00b76591148.
//
// Solidity: event CompetitionTeamsRecorded(uint256 indexed id, string recordHash)
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistryFilterer) WatchCompetitionTeamsRecorded(opts *bind.WatchOpts, sink chan<- *CompetitionTeamsRegistryCompetitionTeamsRecorded, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _CompetitionTeamsRegistry.contract.WatchLogs(opts, "CompetitionTeamsRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompetitionTeamsRegistryCompetitionTeamsRecorded)
				if err := _CompetitionTeamsRegistry.contract.UnpackLog(event, "CompetitionTeamsRecorded", log); err != nil {
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

// ParseCompetitionTeamsRecorded is a log parse operation binding the contract event 0xb5afd97b4f740b109d4828fcda8584f7e87aebb685df30a0aae5f00b76591148.
//
// Solidity: event CompetitionTeamsRecorded(uint256 indexed id, string recordHash)
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistryFilterer) ParseCompetitionTeamsRecorded(log types.Log) (*CompetitionTeamsRegistryCompetitionTeamsRecorded, error) {
	event := new(CompetitionTeamsRegistryCompetitionTeamsRecorded)
	if err := _CompetitionTeamsRegistry.contract.UnpackLog(event, "CompetitionTeamsRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MatchParticipantRegistryMetaData contains all meta data concerning the MatchParticipantRegistry contract.
var MatchParticipantRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"MatchParticipantRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"recordMatchParticipant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recorded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"validateMatchParticipant\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506105968061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c8063501895ae1461004e57806360b2267b1461007757806363c3542d146100a9578063ac2a85a6146100bc575b5f5ffd5b61006161005c36600461029d565b6100d1565b60405161006e91906102b4565b60405180910390f35b61009961008536600461029d565b5f6020819052908152604090205460ff1681565b604051901515815260200161006e565b6100996100b73660046102e9565b610168565b6100cf6100ca3660046102e9565b6101c5565b005b60016020525f9081526040902080546100e990610360565b80601f016020809104026020016040519081016040528092919081815260200182805461011590610360565b80156101605780601f1061013757610100808354040283529160200191610160565b820191905f5260205f20905b81548152906001019060200180831161014357829003601f168201915b505050505081565b5f8381526020819052604081205460ff1680156101bd57508282604051610190929190610398565b604080519182900382205f87815260016020529190912090916101b391906103a7565b6040518091039020145b949350505050565b5f8381526020819052604090205460ff16156102315760405162461bcd60e51b815260206004820152602160248201527f4d617463685061727469636970616e7420616c7265616479207265636f7264656044820152601960fa1b606482015260840160405180910390fd5b5f83815260208181526040808320805460ff19166001908117909155909152902061025d828483610478565b50827f4a4c497b8c1e81bc7bb92e9151f380955f34fa46ab887d42cdd5100df6d6e82f8383604051610290929190610532565b60405180910390a2505050565b5f602082840312156102ad575f5ffd5b5035919050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f5f5f604084860312156102fb575f5ffd5b83359250602084013567ffffffffffffffff811115610318575f5ffd5b8401601f81018613610328575f5ffd5b803567ffffffffffffffff81111561033e575f5ffd5b86602082840101111561034f575f5ffd5b939660209190910195509293505050565b600181811c9082168061037457607f821691505b60208210810361039257634e487b7160e01b5f52602260045260245ffd5b50919050565b818382375f9101908152919050565b5f5f83546103b481610360565b6001821680156103cb57600181146103e05761040d565b60ff198316865281151582028601935061040d565b865f5260205f205f5b83811015610405578154888201526001909101906020016103e9565b505081860193505b509195945050505050565b634e487b7160e01b5f52604160045260245ffd5b601f82111561047357805f5260205f20601f840160051c810160208510156104515750805b601f840160051c820191505b81811015610470575f815560010161045d565b50505b505050565b67ffffffffffffffff83111561049057610490610418565b6104a48361049e8354610360565b8361042c565b5f601f8411600181146104d5575f85156104be5750838201355b5f19600387901b1c1916600186901b178355610470565b5f83815260208120601f198716915b8281101561050457868501358255602094850194600190920191016104e4565b5086821015610520575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b60208152816020820152818360408301375f818301604090810191909152601f909201601f1916010191905056fea264697066735822122073b39842d18ddfa896f05dbb916b99cd2d9ec4030a3a3ff559fc985e3353ca1d64736f6c634300081e0033",
}

// MatchParticipantRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use MatchParticipantRegistryMetaData.ABI instead.
var MatchParticipantRegistryABI = MatchParticipantRegistryMetaData.ABI

// MatchParticipantRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MatchParticipantRegistryMetaData.Bin instead.
var MatchParticipantRegistryBin = MatchParticipantRegistryMetaData.Bin

// DeployMatchParticipantRegistry deploys a new Ethereum contract, binding an instance of MatchParticipantRegistry to it.
func DeployMatchParticipantRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MatchParticipantRegistry, error) {
	parsed, err := MatchParticipantRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MatchParticipantRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MatchParticipantRegistry{MatchParticipantRegistryCaller: MatchParticipantRegistryCaller{contract: contract}, MatchParticipantRegistryTransactor: MatchParticipantRegistryTransactor{contract: contract}, MatchParticipantRegistryFilterer: MatchParticipantRegistryFilterer{contract: contract}}, nil
}

// MatchParticipantRegistry is an auto generated Go binding around an Ethereum contract.
type MatchParticipantRegistry struct {
	MatchParticipantRegistryCaller     // Read-only binding to the contract
	MatchParticipantRegistryTransactor // Write-only binding to the contract
	MatchParticipantRegistryFilterer   // Log filterer for contract events
}

// MatchParticipantRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MatchParticipantRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MatchParticipantRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MatchParticipantRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MatchParticipantRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MatchParticipantRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MatchParticipantRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MatchParticipantRegistrySession struct {
	Contract     *MatchParticipantRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MatchParticipantRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MatchParticipantRegistryCallerSession struct {
	Contract *MatchParticipantRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// MatchParticipantRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MatchParticipantRegistryTransactorSession struct {
	Contract     *MatchParticipantRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// MatchParticipantRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MatchParticipantRegistryRaw struct {
	Contract *MatchParticipantRegistry // Generic contract binding to access the raw methods on
}

// MatchParticipantRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MatchParticipantRegistryCallerRaw struct {
	Contract *MatchParticipantRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// MatchParticipantRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MatchParticipantRegistryTransactorRaw struct {
	Contract *MatchParticipantRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMatchParticipantRegistry creates a new instance of MatchParticipantRegistry, bound to a specific deployed contract.
func NewMatchParticipantRegistry(address common.Address, backend bind.ContractBackend) (*MatchParticipantRegistry, error) {
	contract, err := bindMatchParticipantRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MatchParticipantRegistry{MatchParticipantRegistryCaller: MatchParticipantRegistryCaller{contract: contract}, MatchParticipantRegistryTransactor: MatchParticipantRegistryTransactor{contract: contract}, MatchParticipantRegistryFilterer: MatchParticipantRegistryFilterer{contract: contract}}, nil
}

// NewMatchParticipantRegistryCaller creates a new read-only instance of MatchParticipantRegistry, bound to a specific deployed contract.
func NewMatchParticipantRegistryCaller(address common.Address, caller bind.ContractCaller) (*MatchParticipantRegistryCaller, error) {
	contract, err := bindMatchParticipantRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MatchParticipantRegistryCaller{contract: contract}, nil
}

// NewMatchParticipantRegistryTransactor creates a new write-only instance of MatchParticipantRegistry, bound to a specific deployed contract.
func NewMatchParticipantRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*MatchParticipantRegistryTransactor, error) {
	contract, err := bindMatchParticipantRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MatchParticipantRegistryTransactor{contract: contract}, nil
}

// NewMatchParticipantRegistryFilterer creates a new log filterer instance of MatchParticipantRegistry, bound to a specific deployed contract.
func NewMatchParticipantRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*MatchParticipantRegistryFilterer, error) {
	contract, err := bindMatchParticipantRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MatchParticipantRegistryFilterer{contract: contract}, nil
}

// bindMatchParticipantRegistry binds a generic wrapper to an already deployed contract.
func bindMatchParticipantRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MatchParticipantRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MatchParticipantRegistry *MatchParticipantRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MatchParticipantRegistry.Contract.MatchParticipantRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MatchParticipantRegistry *MatchParticipantRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MatchParticipantRegistry.Contract.MatchParticipantRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MatchParticipantRegistry *MatchParticipantRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MatchParticipantRegistry.Contract.MatchParticipantRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MatchParticipantRegistry *MatchParticipantRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MatchParticipantRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MatchParticipantRegistry *MatchParticipantRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MatchParticipantRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MatchParticipantRegistry *MatchParticipantRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MatchParticipantRegistry.Contract.contract.Transact(opts, method, params...)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_MatchParticipantRegistry *MatchParticipantRegistryCaller) Hashes(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _MatchParticipantRegistry.contract.Call(opts, &out, "hashes", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_MatchParticipantRegistry *MatchParticipantRegistrySession) Hashes(arg0 *big.Int) (string, error) {
	return _MatchParticipantRegistry.Contract.Hashes(&_MatchParticipantRegistry.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_MatchParticipantRegistry *MatchParticipantRegistryCallerSession) Hashes(arg0 *big.Int) (string, error) {
	return _MatchParticipantRegistry.Contract.Hashes(&_MatchParticipantRegistry.CallOpts, arg0)
}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_MatchParticipantRegistry *MatchParticipantRegistryCaller) Recorded(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _MatchParticipantRegistry.contract.Call(opts, &out, "recorded", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_MatchParticipantRegistry *MatchParticipantRegistrySession) Recorded(arg0 *big.Int) (bool, error) {
	return _MatchParticipantRegistry.Contract.Recorded(&_MatchParticipantRegistry.CallOpts, arg0)
}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_MatchParticipantRegistry *MatchParticipantRegistryCallerSession) Recorded(arg0 *big.Int) (bool, error) {
	return _MatchParticipantRegistry.Contract.Recorded(&_MatchParticipantRegistry.CallOpts, arg0)
}

// ValidateMatchParticipant is a free data retrieval call binding the contract method 0x63c3542d.
//
// Solidity: function validateMatchParticipant(uint256 id, string recordHash) view returns(bool)
func (_MatchParticipantRegistry *MatchParticipantRegistryCaller) ValidateMatchParticipant(opts *bind.CallOpts, id *big.Int, recordHash string) (bool, error) {
	var out []interface{}
	err := _MatchParticipantRegistry.contract.Call(opts, &out, "validateMatchParticipant", id, recordHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateMatchParticipant is a free data retrieval call binding the contract method 0x63c3542d.
//
// Solidity: function validateMatchParticipant(uint256 id, string recordHash) view returns(bool)
func (_MatchParticipantRegistry *MatchParticipantRegistrySession) ValidateMatchParticipant(id *big.Int, recordHash string) (bool, error) {
	return _MatchParticipantRegistry.Contract.ValidateMatchParticipant(&_MatchParticipantRegistry.CallOpts, id, recordHash)
}

// ValidateMatchParticipant is a free data retrieval call binding the contract method 0x63c3542d.
//
// Solidity: function validateMatchParticipant(uint256 id, string recordHash) view returns(bool)
func (_MatchParticipantRegistry *MatchParticipantRegistryCallerSession) ValidateMatchParticipant(id *big.Int, recordHash string) (bool, error) {
	return _MatchParticipantRegistry.Contract.ValidateMatchParticipant(&_MatchParticipantRegistry.CallOpts, id, recordHash)
}

// RecordMatchParticipant is a paid mutator transaction binding the contract method 0xac2a85a6.
//
// Solidity: function recordMatchParticipant(uint256 id, string recordHash) returns()
func (_MatchParticipantRegistry *MatchParticipantRegistryTransactor) RecordMatchParticipant(opts *bind.TransactOpts, id *big.Int, recordHash string) (*types.Transaction, error) {
	return _MatchParticipantRegistry.contract.Transact(opts, "recordMatchParticipant", id, recordHash)
}

// RecordMatchParticipant is a paid mutator transaction binding the contract method 0xac2a85a6.
//
// Solidity: function recordMatchParticipant(uint256 id, string recordHash) returns()
func (_MatchParticipantRegistry *MatchParticipantRegistrySession) RecordMatchParticipant(id *big.Int, recordHash string) (*types.Transaction, error) {
	return _MatchParticipantRegistry.Contract.RecordMatchParticipant(&_MatchParticipantRegistry.TransactOpts, id, recordHash)
}

// RecordMatchParticipant is a paid mutator transaction binding the contract method 0xac2a85a6.
//
// Solidity: function recordMatchParticipant(uint256 id, string recordHash) returns()
func (_MatchParticipantRegistry *MatchParticipantRegistryTransactorSession) RecordMatchParticipant(id *big.Int, recordHash string) (*types.Transaction, error) {
	return _MatchParticipantRegistry.Contract.RecordMatchParticipant(&_MatchParticipantRegistry.TransactOpts, id, recordHash)
}

// MatchParticipantRegistryMatchParticipantRecordedIterator is returned from FilterMatchParticipantRecorded and is used to iterate over the raw logs and unpacked data for MatchParticipantRecorded events raised by the MatchParticipantRegistry contract.
type MatchParticipantRegistryMatchParticipantRecordedIterator struct {
	Event *MatchParticipantRegistryMatchParticipantRecorded // Event containing the contract specifics and raw log

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
func (it *MatchParticipantRegistryMatchParticipantRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MatchParticipantRegistryMatchParticipantRecorded)
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
		it.Event = new(MatchParticipantRegistryMatchParticipantRecorded)
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
func (it *MatchParticipantRegistryMatchParticipantRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MatchParticipantRegistryMatchParticipantRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MatchParticipantRegistryMatchParticipantRecorded represents a MatchParticipantRecorded event raised by the MatchParticipantRegistry contract.
type MatchParticipantRegistryMatchParticipantRecorded struct {
	Id         *big.Int
	RecordHash string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMatchParticipantRecorded is a free log retrieval operation binding the contract event 0x4a4c497b8c1e81bc7bb92e9151f380955f34fa46ab887d42cdd5100df6d6e82f.
//
// Solidity: event MatchParticipantRecorded(uint256 indexed id, string recordHash)
func (_MatchParticipantRegistry *MatchParticipantRegistryFilterer) FilterMatchParticipantRecorded(opts *bind.FilterOpts, id []*big.Int) (*MatchParticipantRegistryMatchParticipantRecordedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MatchParticipantRegistry.contract.FilterLogs(opts, "MatchParticipantRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return &MatchParticipantRegistryMatchParticipantRecordedIterator{contract: _MatchParticipantRegistry.contract, event: "MatchParticipantRecorded", logs: logs, sub: sub}, nil
}

// WatchMatchParticipantRecorded is a free log subscription operation binding the contract event 0x4a4c497b8c1e81bc7bb92e9151f380955f34fa46ab887d42cdd5100df6d6e82f.
//
// Solidity: event MatchParticipantRecorded(uint256 indexed id, string recordHash)
func (_MatchParticipantRegistry *MatchParticipantRegistryFilterer) WatchMatchParticipantRecorded(opts *bind.WatchOpts, sink chan<- *MatchParticipantRegistryMatchParticipantRecorded, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MatchParticipantRegistry.contract.WatchLogs(opts, "MatchParticipantRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MatchParticipantRegistryMatchParticipantRecorded)
				if err := _MatchParticipantRegistry.contract.UnpackLog(event, "MatchParticipantRecorded", log); err != nil {
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

// ParseMatchParticipantRecorded is a log parse operation binding the contract event 0x4a4c497b8c1e81bc7bb92e9151f380955f34fa46ab887d42cdd5100df6d6e82f.
//
// Solidity: event MatchParticipantRecorded(uint256 indexed id, string recordHash)
func (_MatchParticipantRegistry *MatchParticipantRegistryFilterer) ParseMatchParticipantRecorded(log types.Log) (*MatchParticipantRegistryMatchParticipantRecorded, error) {
	event := new(MatchParticipantRegistryMatchParticipantRecorded)
	if err := _MatchParticipantRegistry.contract.UnpackLog(event, "MatchParticipantRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MatchRegistryMetaData contains all meta data concerning the MatchRegistry contract.
var MatchRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"MatchRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"recordMatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recorded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"validateMatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506105858061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c8063501895ae1461004e57806360b2267b14610077578063daec364f146100a9578063fb77d0d8146100be575b5f5ffd5b61006161005c36600461028c565b6100d1565b60405161006e91906102a3565b60405180910390f35b61009961008536600461028c565b5f6020819052908152604090205460ff1681565b604051901515815260200161006e565b6100bc6100b73660046102d8565b610168565b005b6100996100cc3660046102d8565b61022f565b60016020525f9081526040902080546100e99061034f565b80601f01602080910402602001604051908101604052809291908181526020018280546101159061034f565b80156101605780601f1061013757610100808354040283529160200191610160565b820191905f5260205f20905b81548152906001019060200180831161014357829003601f168201915b505050505081565b5f8381526020819052604090205460ff16156101c35760405162461bcd60e51b815260206004820152601660248201527513585d18da08185b1c9958591e481c9958dbdc99195960521b604482015260640160405180910390fd5b5f83815260208181526040808320805460ff1916600190811790915590915290206101ef8284836103e7565b50827f480d6ba8c2fc3a872f1ea9cb5995e6c5e2eddf416be39c71e88a62a33bd8135a83836040516102229291906104a1565b60405180910390a2505050565b5f8381526020819052604081205460ff168015610284575082826040516102579291906104cf565b604080519182900382205f878152600160205291909120909161027a91906104de565b6040518091039020145b949350505050565b5f6020828403121561029c575f5ffd5b5035919050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f5f5f604084860312156102ea575f5ffd5b83359250602084013567ffffffffffffffff811115610307575f5ffd5b8401601f81018613610317575f5ffd5b803567ffffffffffffffff81111561032d575f5ffd5b86602082840101111561033e575f5ffd5b939660209190910195509293505050565b600181811c9082168061036357607f821691505b60208210810361038157634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52604160045260245ffd5b601f8211156103e257805f5260205f20601f840160051c810160208510156103c05750805b601f840160051c820191505b818110156103df575f81556001016103cc565b50505b505050565b67ffffffffffffffff8311156103ff576103ff610387565b6104138361040d835461034f565b8361039b565b5f601f841160018114610444575f851561042d5750838201355b5f19600387901b1c1916600186901b1783556103df565b5f83815260208120601f198716915b828110156104735786850135825560209485019460019092019101610453565b508682101561048f575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b60208152816020820152818360408301375f818301604090810191909152601f909201601f19160101919050565b818382375f9101908152919050565b5f5f83546104eb8161034f565b600182168015610502576001811461051757610544565b60ff1983168652811515820286019350610544565b865f5260205f205f5b8381101561053c57815488820152600190910190602001610520565b505081860193505b50919594505050505056fea2646970667358221220098bf52451f8fa89fb5860229a5cc43ec578e63512b3b4df339451e5686ba37364736f6c634300081e0033",
}

// MatchRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use MatchRegistryMetaData.ABI instead.
var MatchRegistryABI = MatchRegistryMetaData.ABI

// MatchRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MatchRegistryMetaData.Bin instead.
var MatchRegistryBin = MatchRegistryMetaData.Bin

// DeployMatchRegistry deploys a new Ethereum contract, binding an instance of MatchRegistry to it.
func DeployMatchRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MatchRegistry, error) {
	parsed, err := MatchRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MatchRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MatchRegistry{MatchRegistryCaller: MatchRegistryCaller{contract: contract}, MatchRegistryTransactor: MatchRegistryTransactor{contract: contract}, MatchRegistryFilterer: MatchRegistryFilterer{contract: contract}}, nil
}

// MatchRegistry is an auto generated Go binding around an Ethereum contract.
type MatchRegistry struct {
	MatchRegistryCaller     // Read-only binding to the contract
	MatchRegistryTransactor // Write-only binding to the contract
	MatchRegistryFilterer   // Log filterer for contract events
}

// MatchRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MatchRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MatchRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MatchRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MatchRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MatchRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MatchRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MatchRegistrySession struct {
	Contract     *MatchRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MatchRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MatchRegistryCallerSession struct {
	Contract *MatchRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MatchRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MatchRegistryTransactorSession struct {
	Contract     *MatchRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MatchRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MatchRegistryRaw struct {
	Contract *MatchRegistry // Generic contract binding to access the raw methods on
}

// MatchRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MatchRegistryCallerRaw struct {
	Contract *MatchRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// MatchRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MatchRegistryTransactorRaw struct {
	Contract *MatchRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMatchRegistry creates a new instance of MatchRegistry, bound to a specific deployed contract.
func NewMatchRegistry(address common.Address, backend bind.ContractBackend) (*MatchRegistry, error) {
	contract, err := bindMatchRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MatchRegistry{MatchRegistryCaller: MatchRegistryCaller{contract: contract}, MatchRegistryTransactor: MatchRegistryTransactor{contract: contract}, MatchRegistryFilterer: MatchRegistryFilterer{contract: contract}}, nil
}

// NewMatchRegistryCaller creates a new read-only instance of MatchRegistry, bound to a specific deployed contract.
func NewMatchRegistryCaller(address common.Address, caller bind.ContractCaller) (*MatchRegistryCaller, error) {
	contract, err := bindMatchRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MatchRegistryCaller{contract: contract}, nil
}

// NewMatchRegistryTransactor creates a new write-only instance of MatchRegistry, bound to a specific deployed contract.
func NewMatchRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*MatchRegistryTransactor, error) {
	contract, err := bindMatchRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MatchRegistryTransactor{contract: contract}, nil
}

// NewMatchRegistryFilterer creates a new log filterer instance of MatchRegistry, bound to a specific deployed contract.
func NewMatchRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*MatchRegistryFilterer, error) {
	contract, err := bindMatchRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MatchRegistryFilterer{contract: contract}, nil
}

// bindMatchRegistry binds a generic wrapper to an already deployed contract.
func bindMatchRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MatchRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MatchRegistry *MatchRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MatchRegistry.Contract.MatchRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MatchRegistry *MatchRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MatchRegistry.Contract.MatchRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MatchRegistry *MatchRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MatchRegistry.Contract.MatchRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MatchRegistry *MatchRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MatchRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MatchRegistry *MatchRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MatchRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MatchRegistry *MatchRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MatchRegistry.Contract.contract.Transact(opts, method, params...)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_MatchRegistry *MatchRegistryCaller) Hashes(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _MatchRegistry.contract.Call(opts, &out, "hashes", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_MatchRegistry *MatchRegistrySession) Hashes(arg0 *big.Int) (string, error) {
	return _MatchRegistry.Contract.Hashes(&_MatchRegistry.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_MatchRegistry *MatchRegistryCallerSession) Hashes(arg0 *big.Int) (string, error) {
	return _MatchRegistry.Contract.Hashes(&_MatchRegistry.CallOpts, arg0)
}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_MatchRegistry *MatchRegistryCaller) Recorded(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _MatchRegistry.contract.Call(opts, &out, "recorded", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_MatchRegistry *MatchRegistrySession) Recorded(arg0 *big.Int) (bool, error) {
	return _MatchRegistry.Contract.Recorded(&_MatchRegistry.CallOpts, arg0)
}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_MatchRegistry *MatchRegistryCallerSession) Recorded(arg0 *big.Int) (bool, error) {
	return _MatchRegistry.Contract.Recorded(&_MatchRegistry.CallOpts, arg0)
}

// ValidateMatch is a free data retrieval call binding the contract method 0xfb77d0d8.
//
// Solidity: function validateMatch(uint256 id, string recordHash) view returns(bool)
func (_MatchRegistry *MatchRegistryCaller) ValidateMatch(opts *bind.CallOpts, id *big.Int, recordHash string) (bool, error) {
	var out []interface{}
	err := _MatchRegistry.contract.Call(opts, &out, "validateMatch", id, recordHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateMatch is a free data retrieval call binding the contract method 0xfb77d0d8.
//
// Solidity: function validateMatch(uint256 id, string recordHash) view returns(bool)
func (_MatchRegistry *MatchRegistrySession) ValidateMatch(id *big.Int, recordHash string) (bool, error) {
	return _MatchRegistry.Contract.ValidateMatch(&_MatchRegistry.CallOpts, id, recordHash)
}

// ValidateMatch is a free data retrieval call binding the contract method 0xfb77d0d8.
//
// Solidity: function validateMatch(uint256 id, string recordHash) view returns(bool)
func (_MatchRegistry *MatchRegistryCallerSession) ValidateMatch(id *big.Int, recordHash string) (bool, error) {
	return _MatchRegistry.Contract.ValidateMatch(&_MatchRegistry.CallOpts, id, recordHash)
}

// RecordMatch is a paid mutator transaction binding the contract method 0xdaec364f.
//
// Solidity: function recordMatch(uint256 id, string recordHash) returns()
func (_MatchRegistry *MatchRegistryTransactor) RecordMatch(opts *bind.TransactOpts, id *big.Int, recordHash string) (*types.Transaction, error) {
	return _MatchRegistry.contract.Transact(opts, "recordMatch", id, recordHash)
}

// RecordMatch is a paid mutator transaction binding the contract method 0xdaec364f.
//
// Solidity: function recordMatch(uint256 id, string recordHash) returns()
func (_MatchRegistry *MatchRegistrySession) RecordMatch(id *big.Int, recordHash string) (*types.Transaction, error) {
	return _MatchRegistry.Contract.RecordMatch(&_MatchRegistry.TransactOpts, id, recordHash)
}

// RecordMatch is a paid mutator transaction binding the contract method 0xdaec364f.
//
// Solidity: function recordMatch(uint256 id, string recordHash) returns()
func (_MatchRegistry *MatchRegistryTransactorSession) RecordMatch(id *big.Int, recordHash string) (*types.Transaction, error) {
	return _MatchRegistry.Contract.RecordMatch(&_MatchRegistry.TransactOpts, id, recordHash)
}

// MatchRegistryMatchRecordedIterator is returned from FilterMatchRecorded and is used to iterate over the raw logs and unpacked data for MatchRecorded events raised by the MatchRegistry contract.
type MatchRegistryMatchRecordedIterator struct {
	Event *MatchRegistryMatchRecorded // Event containing the contract specifics and raw log

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
func (it *MatchRegistryMatchRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MatchRegistryMatchRecorded)
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
		it.Event = new(MatchRegistryMatchRecorded)
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
func (it *MatchRegistryMatchRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MatchRegistryMatchRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MatchRegistryMatchRecorded represents a MatchRecorded event raised by the MatchRegistry contract.
type MatchRegistryMatchRecorded struct {
	Id         *big.Int
	RecordHash string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMatchRecorded is a free log retrieval operation binding the contract event 0x480d6ba8c2fc3a872f1ea9cb5995e6c5e2eddf416be39c71e88a62a33bd8135a.
//
// Solidity: event MatchRecorded(uint256 indexed id, string recordHash)
func (_MatchRegistry *MatchRegistryFilterer) FilterMatchRecorded(opts *bind.FilterOpts, id []*big.Int) (*MatchRegistryMatchRecordedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MatchRegistry.contract.FilterLogs(opts, "MatchRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return &MatchRegistryMatchRecordedIterator{contract: _MatchRegistry.contract, event: "MatchRecorded", logs: logs, sub: sub}, nil
}

// WatchMatchRecorded is a free log subscription operation binding the contract event 0x480d6ba8c2fc3a872f1ea9cb5995e6c5e2eddf416be39c71e88a62a33bd8135a.
//
// Solidity: event MatchRecorded(uint256 indexed id, string recordHash)
func (_MatchRegistry *MatchRegistryFilterer) WatchMatchRecorded(opts *bind.WatchOpts, sink chan<- *MatchRegistryMatchRecorded, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MatchRegistry.contract.WatchLogs(opts, "MatchRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MatchRegistryMatchRecorded)
				if err := _MatchRegistry.contract.UnpackLog(event, "MatchRecorded", log); err != nil {
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

// ParseMatchRecorded is a log parse operation binding the contract event 0x480d6ba8c2fc3a872f1ea9cb5995e6c5e2eddf416be39c71e88a62a33bd8135a.
//
// Solidity: event MatchRecorded(uint256 indexed id, string recordHash)
func (_MatchRegistry *MatchRegistryFilterer) ParseMatchRecorded(log types.Log) (*MatchRegistryMatchRecorded, error) {
	event := new(MatchRegistryMatchRecorded)
	if err := _MatchRegistry.contract.UnpackLog(event, "MatchRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PersonRegistryMetaData contains all meta data concerning the PersonRegistry contract.
var PersonRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"PersonRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"recordPerson\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recorded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"validatePerson\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506105888061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c80632b796d5e1461004e578063501895ae146100765780635a204a5a1461009657806360b2267b146100ab575b5f5ffd5b61006161005c36600461028f565b6100cd565b60405190151581526020015b60405180910390f35b610089610084366004610306565b61012a565b60405161006d919061031d565b6100a96100a436600461028f565b6101c1565b005b6100616100b9366004610306565b5f6020819052908152604090205460ff1681565b5f8381526020819052604081205460ff168015610122575082826040516100f5929190610352565b604080519182900382205f87815260016020529190912090916101189190610399565b6040518091039020145b949350505050565b60016020525f90815260409020805461014290610361565b80601f016020809104026020016040519081016040528092919081815260200182805461016e90610361565b80156101b95780601f10610190576101008083540402835291602001916101b9565b820191905f5260205f20905b81548152906001019060200180831161019c57829003601f168201915b505050505081565b5f8381526020819052604090205460ff16156102235760405162461bcd60e51b815260206004820152601760248201527f506572736f6e20616c7265616479207265636f72646564000000000000000000604482015260640160405180910390fd5b5f83815260208181526040808320805460ff19166001908117909155909152902061024f82848361046a565b50827fd908cd34fe6486ae21421c8f652420da6374812e726c445a45c3ea56cdc1ad488383604051610282929190610524565b60405180910390a2505050565b5f5f5f604084860312156102a1575f5ffd5b83359250602084013567ffffffffffffffff8111156102be575f5ffd5b8401601f810186136102ce575f5ffd5b803567ffffffffffffffff8111156102e4575f5ffd5b8660208284010111156102f5575f5ffd5b939660209190910195509293505050565b5f60208284031215610316575f5ffd5b5035919050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b818382375f9101908152919050565b600181811c9082168061037557607f821691505b60208210810361039357634e487b7160e01b5f52602260045260245ffd5b50919050565b5f5f83546103a681610361565b6001821680156103bd57600181146103d2576103ff565b60ff19831686528115158202860193506103ff565b865f5260205f205f5b838110156103f7578154888201526001909101906020016103db565b505081860193505b509195945050505050565b634e487b7160e01b5f52604160045260245ffd5b601f82111561046557805f5260205f20601f840160051c810160208510156104435750805b601f840160051c820191505b81811015610462575f815560010161044f565b50505b505050565b67ffffffffffffffff8311156104825761048261040a565b610496836104908354610361565b8361041e565b5f601f8411600181146104c7575f85156104b05750838201355b5f19600387901b1c1916600186901b178355610462565b5f83815260208120601f198716915b828110156104f657868501358255602094850194600190920191016104d6565b5086821015610512575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b60208152816020820152818360408301375f818301604090810191909152601f909201601f1916010191905056fea2646970667358221220d7d42889adddd3fef725ef7089ea36b9c666fa69d4877a07943933800a49d4ee64736f6c634300081e0033",
}

// PersonRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use PersonRegistryMetaData.ABI instead.
var PersonRegistryABI = PersonRegistryMetaData.ABI

// PersonRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PersonRegistryMetaData.Bin instead.
var PersonRegistryBin = PersonRegistryMetaData.Bin

// DeployPersonRegistry deploys a new Ethereum contract, binding an instance of PersonRegistry to it.
func DeployPersonRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PersonRegistry, error) {
	parsed, err := PersonRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PersonRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PersonRegistry{PersonRegistryCaller: PersonRegistryCaller{contract: contract}, PersonRegistryTransactor: PersonRegistryTransactor{contract: contract}, PersonRegistryFilterer: PersonRegistryFilterer{contract: contract}}, nil
}

// PersonRegistry is an auto generated Go binding around an Ethereum contract.
type PersonRegistry struct {
	PersonRegistryCaller     // Read-only binding to the contract
	PersonRegistryTransactor // Write-only binding to the contract
	PersonRegistryFilterer   // Log filterer for contract events
}

// PersonRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PersonRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PersonRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PersonRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PersonRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PersonRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PersonRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PersonRegistrySession struct {
	Contract     *PersonRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PersonRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PersonRegistryCallerSession struct {
	Contract *PersonRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PersonRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PersonRegistryTransactorSession struct {
	Contract     *PersonRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PersonRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PersonRegistryRaw struct {
	Contract *PersonRegistry // Generic contract binding to access the raw methods on
}

// PersonRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PersonRegistryCallerRaw struct {
	Contract *PersonRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// PersonRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PersonRegistryTransactorRaw struct {
	Contract *PersonRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPersonRegistry creates a new instance of PersonRegistry, bound to a specific deployed contract.
func NewPersonRegistry(address common.Address, backend bind.ContractBackend) (*PersonRegistry, error) {
	contract, err := bindPersonRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PersonRegistry{PersonRegistryCaller: PersonRegistryCaller{contract: contract}, PersonRegistryTransactor: PersonRegistryTransactor{contract: contract}, PersonRegistryFilterer: PersonRegistryFilterer{contract: contract}}, nil
}

// NewPersonRegistryCaller creates a new read-only instance of PersonRegistry, bound to a specific deployed contract.
func NewPersonRegistryCaller(address common.Address, caller bind.ContractCaller) (*PersonRegistryCaller, error) {
	contract, err := bindPersonRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PersonRegistryCaller{contract: contract}, nil
}

// NewPersonRegistryTransactor creates a new write-only instance of PersonRegistry, bound to a specific deployed contract.
func NewPersonRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*PersonRegistryTransactor, error) {
	contract, err := bindPersonRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PersonRegistryTransactor{contract: contract}, nil
}

// NewPersonRegistryFilterer creates a new log filterer instance of PersonRegistry, bound to a specific deployed contract.
func NewPersonRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*PersonRegistryFilterer, error) {
	contract, err := bindPersonRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PersonRegistryFilterer{contract: contract}, nil
}

// bindPersonRegistry binds a generic wrapper to an already deployed contract.
func bindPersonRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PersonRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PersonRegistry *PersonRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PersonRegistry.Contract.PersonRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PersonRegistry *PersonRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PersonRegistry.Contract.PersonRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PersonRegistry *PersonRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PersonRegistry.Contract.PersonRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PersonRegistry *PersonRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PersonRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PersonRegistry *PersonRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PersonRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PersonRegistry *PersonRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PersonRegistry.Contract.contract.Transact(opts, method, params...)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_PersonRegistry *PersonRegistryCaller) Hashes(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _PersonRegistry.contract.Call(opts, &out, "hashes", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_PersonRegistry *PersonRegistrySession) Hashes(arg0 *big.Int) (string, error) {
	return _PersonRegistry.Contract.Hashes(&_PersonRegistry.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_PersonRegistry *PersonRegistryCallerSession) Hashes(arg0 *big.Int) (string, error) {
	return _PersonRegistry.Contract.Hashes(&_PersonRegistry.CallOpts, arg0)
}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_PersonRegistry *PersonRegistryCaller) Recorded(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _PersonRegistry.contract.Call(opts, &out, "recorded", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_PersonRegistry *PersonRegistrySession) Recorded(arg0 *big.Int) (bool, error) {
	return _PersonRegistry.Contract.Recorded(&_PersonRegistry.CallOpts, arg0)
}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_PersonRegistry *PersonRegistryCallerSession) Recorded(arg0 *big.Int) (bool, error) {
	return _PersonRegistry.Contract.Recorded(&_PersonRegistry.CallOpts, arg0)
}

// ValidatePerson is a free data retrieval call binding the contract method 0x2b796d5e.
//
// Solidity: function validatePerson(uint256 id, string recordHash) view returns(bool)
func (_PersonRegistry *PersonRegistryCaller) ValidatePerson(opts *bind.CallOpts, id *big.Int, recordHash string) (bool, error) {
	var out []interface{}
	err := _PersonRegistry.contract.Call(opts, &out, "validatePerson", id, recordHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidatePerson is a free data retrieval call binding the contract method 0x2b796d5e.
//
// Solidity: function validatePerson(uint256 id, string recordHash) view returns(bool)
func (_PersonRegistry *PersonRegistrySession) ValidatePerson(id *big.Int, recordHash string) (bool, error) {
	return _PersonRegistry.Contract.ValidatePerson(&_PersonRegistry.CallOpts, id, recordHash)
}

// ValidatePerson is a free data retrieval call binding the contract method 0x2b796d5e.
//
// Solidity: function validatePerson(uint256 id, string recordHash) view returns(bool)
func (_PersonRegistry *PersonRegistryCallerSession) ValidatePerson(id *big.Int, recordHash string) (bool, error) {
	return _PersonRegistry.Contract.ValidatePerson(&_PersonRegistry.CallOpts, id, recordHash)
}

// RecordPerson is a paid mutator transaction binding the contract method 0x5a204a5a.
//
// Solidity: function recordPerson(uint256 id, string recordHash) returns()
func (_PersonRegistry *PersonRegistryTransactor) RecordPerson(opts *bind.TransactOpts, id *big.Int, recordHash string) (*types.Transaction, error) {
	return _PersonRegistry.contract.Transact(opts, "recordPerson", id, recordHash)
}

// RecordPerson is a paid mutator transaction binding the contract method 0x5a204a5a.
//
// Solidity: function recordPerson(uint256 id, string recordHash) returns()
func (_PersonRegistry *PersonRegistrySession) RecordPerson(id *big.Int, recordHash string) (*types.Transaction, error) {
	return _PersonRegistry.Contract.RecordPerson(&_PersonRegistry.TransactOpts, id, recordHash)
}

// RecordPerson is a paid mutator transaction binding the contract method 0x5a204a5a.
//
// Solidity: function recordPerson(uint256 id, string recordHash) returns()
func (_PersonRegistry *PersonRegistryTransactorSession) RecordPerson(id *big.Int, recordHash string) (*types.Transaction, error) {
	return _PersonRegistry.Contract.RecordPerson(&_PersonRegistry.TransactOpts, id, recordHash)
}

// PersonRegistryPersonRecordedIterator is returned from FilterPersonRecorded and is used to iterate over the raw logs and unpacked data for PersonRecorded events raised by the PersonRegistry contract.
type PersonRegistryPersonRecordedIterator struct {
	Event *PersonRegistryPersonRecorded // Event containing the contract specifics and raw log

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
func (it *PersonRegistryPersonRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PersonRegistryPersonRecorded)
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
		it.Event = new(PersonRegistryPersonRecorded)
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
func (it *PersonRegistryPersonRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PersonRegistryPersonRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PersonRegistryPersonRecorded represents a PersonRecorded event raised by the PersonRegistry contract.
type PersonRegistryPersonRecorded struct {
	Id         *big.Int
	RecordHash string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPersonRecorded is a free log retrieval operation binding the contract event 0xd908cd34fe6486ae21421c8f652420da6374812e726c445a45c3ea56cdc1ad48.
//
// Solidity: event PersonRecorded(uint256 indexed id, string recordHash)
func (_PersonRegistry *PersonRegistryFilterer) FilterPersonRecorded(opts *bind.FilterOpts, id []*big.Int) (*PersonRegistryPersonRecordedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _PersonRegistry.contract.FilterLogs(opts, "PersonRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return &PersonRegistryPersonRecordedIterator{contract: _PersonRegistry.contract, event: "PersonRecorded", logs: logs, sub: sub}, nil
}

// WatchPersonRecorded is a free log subscription operation binding the contract event 0xd908cd34fe6486ae21421c8f652420da6374812e726c445a45c3ea56cdc1ad48.
//
// Solidity: event PersonRecorded(uint256 indexed id, string recordHash)
func (_PersonRegistry *PersonRegistryFilterer) WatchPersonRecorded(opts *bind.WatchOpts, sink chan<- *PersonRegistryPersonRecorded, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _PersonRegistry.contract.WatchLogs(opts, "PersonRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PersonRegistryPersonRecorded)
				if err := _PersonRegistry.contract.UnpackLog(event, "PersonRecorded", log); err != nil {
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

// ParsePersonRecorded is a log parse operation binding the contract event 0xd908cd34fe6486ae21421c8f652420da6374812e726c445a45c3ea56cdc1ad48.
//
// Solidity: event PersonRecorded(uint256 indexed id, string recordHash)
func (_PersonRegistry *PersonRegistryFilterer) ParsePersonRecorded(log types.Log) (*PersonRegistryPersonRecorded, error) {
	event := new(PersonRegistryPersonRecorded)
	if err := _PersonRegistry.contract.UnpackLog(event, "PersonRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizeRegistryMetaData contains all meta data concerning the PrizeRegistry contract.
var PrizeRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"PrizeRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"recordPrize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recorded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"validatePrize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506105858061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c8063501895ae1461004e57806360b2267b1461007757806363a66be2146100a9578063ab84ec9f146100bc575b5f5ffd5b61006161005c36600461028c565b6100d1565b60405161006e91906102a3565b60405180910390f35b61009961008536600461028c565b5f6020819052908152604090205460ff1681565b604051901515815260200161006e565b6100996100b73660046102d8565b610168565b6100cf6100ca3660046102d8565b6101c5565b005b60016020525f9081526040902080546100e99061034f565b80601f01602080910402602001604051908101604052809291908181526020018280546101159061034f565b80156101605780601f1061013757610100808354040283529160200191610160565b820191905f5260205f20905b81548152906001019060200180831161014357829003601f168201915b505050505081565b5f8381526020819052604081205460ff1680156101bd57508282604051610190929190610387565b604080519182900382205f87815260016020529190912090916101b39190610396565b6040518091039020145b949350505050565b5f8381526020819052604090205460ff16156102205760405162461bcd60e51b8152602060048201526016602482015275141c9a5e9948185b1c9958591e481c9958dbdc99195960521b604482015260640160405180910390fd5b5f83815260208181526040808320805460ff19166001908117909155909152902061024c828483610467565b50827f2f8720dba077ce08b29ef00ca011bd0420304cd53e1795486dbf734022d4c659838360405161027f929190610521565b60405180910390a2505050565b5f6020828403121561029c575f5ffd5b5035919050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f5f5f604084860312156102ea575f5ffd5b83359250602084013567ffffffffffffffff811115610307575f5ffd5b8401601f81018613610317575f5ffd5b803567ffffffffffffffff81111561032d575f5ffd5b86602082840101111561033e575f5ffd5b939660209190910195509293505050565b600181811c9082168061036357607f821691505b60208210810361038157634e487b7160e01b5f52602260045260245ffd5b50919050565b818382375f9101908152919050565b5f5f83546103a38161034f565b6001821680156103ba57600181146103cf576103fc565b60ff19831686528115158202860193506103fc565b865f5260205f205f5b838110156103f4578154888201526001909101906020016103d8565b505081860193505b509195945050505050565b634e487b7160e01b5f52604160045260245ffd5b601f82111561046257805f5260205f20601f840160051c810160208510156104405750805b601f840160051c820191505b8181101561045f575f815560010161044c565b50505b505050565b67ffffffffffffffff83111561047f5761047f610407565b6104938361048d835461034f565b8361041b565b5f601f8411600181146104c4575f85156104ad5750838201355b5f19600387901b1c1916600186901b17835561045f565b5f83815260208120601f198716915b828110156104f357868501358255602094850194600190920191016104d3565b508682101561050f575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b60208152816020820152818360408301375f818301604090810191909152601f909201601f1916010191905056fea2646970667358221220da41a58e7f8f5c9fe308e7dfc3fee7d5f7dd761cb9de0de04f5470427e4a55c464736f6c634300081e0033",
}

// PrizeRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use PrizeRegistryMetaData.ABI instead.
var PrizeRegistryABI = PrizeRegistryMetaData.ABI

// PrizeRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PrizeRegistryMetaData.Bin instead.
var PrizeRegistryBin = PrizeRegistryMetaData.Bin

// DeployPrizeRegistry deploys a new Ethereum contract, binding an instance of PrizeRegistry to it.
func DeployPrizeRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PrizeRegistry, error) {
	parsed, err := PrizeRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PrizeRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PrizeRegistry{PrizeRegistryCaller: PrizeRegistryCaller{contract: contract}, PrizeRegistryTransactor: PrizeRegistryTransactor{contract: contract}, PrizeRegistryFilterer: PrizeRegistryFilterer{contract: contract}}, nil
}

// PrizeRegistry is an auto generated Go binding around an Ethereum contract.
type PrizeRegistry struct {
	PrizeRegistryCaller     // Read-only binding to the contract
	PrizeRegistryTransactor // Write-only binding to the contract
	PrizeRegistryFilterer   // Log filterer for contract events
}

// PrizeRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrizeRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrizeRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrizeRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrizeRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrizeRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrizeRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrizeRegistrySession struct {
	Contract     *PrizeRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PrizeRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrizeRegistryCallerSession struct {
	Contract *PrizeRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PrizeRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrizeRegistryTransactorSession struct {
	Contract     *PrizeRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PrizeRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrizeRegistryRaw struct {
	Contract *PrizeRegistry // Generic contract binding to access the raw methods on
}

// PrizeRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrizeRegistryCallerRaw struct {
	Contract *PrizeRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// PrizeRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrizeRegistryTransactorRaw struct {
	Contract *PrizeRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrizeRegistry creates a new instance of PrizeRegistry, bound to a specific deployed contract.
func NewPrizeRegistry(address common.Address, backend bind.ContractBackend) (*PrizeRegistry, error) {
	contract, err := bindPrizeRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PrizeRegistry{PrizeRegistryCaller: PrizeRegistryCaller{contract: contract}, PrizeRegistryTransactor: PrizeRegistryTransactor{contract: contract}, PrizeRegistryFilterer: PrizeRegistryFilterer{contract: contract}}, nil
}

// NewPrizeRegistryCaller creates a new read-only instance of PrizeRegistry, bound to a specific deployed contract.
func NewPrizeRegistryCaller(address common.Address, caller bind.ContractCaller) (*PrizeRegistryCaller, error) {
	contract, err := bindPrizeRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrizeRegistryCaller{contract: contract}, nil
}

// NewPrizeRegistryTransactor creates a new write-only instance of PrizeRegistry, bound to a specific deployed contract.
func NewPrizeRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*PrizeRegistryTransactor, error) {
	contract, err := bindPrizeRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrizeRegistryTransactor{contract: contract}, nil
}

// NewPrizeRegistryFilterer creates a new log filterer instance of PrizeRegistry, bound to a specific deployed contract.
func NewPrizeRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*PrizeRegistryFilterer, error) {
	contract, err := bindPrizeRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrizeRegistryFilterer{contract: contract}, nil
}

// bindPrizeRegistry binds a generic wrapper to an already deployed contract.
func bindPrizeRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PrizeRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrizeRegistry *PrizeRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrizeRegistry.Contract.PrizeRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrizeRegistry *PrizeRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrizeRegistry.Contract.PrizeRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrizeRegistry *PrizeRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrizeRegistry.Contract.PrizeRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrizeRegistry *PrizeRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrizeRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrizeRegistry *PrizeRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrizeRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrizeRegistry *PrizeRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrizeRegistry.Contract.contract.Transact(opts, method, params...)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_PrizeRegistry *PrizeRegistryCaller) Hashes(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _PrizeRegistry.contract.Call(opts, &out, "hashes", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_PrizeRegistry *PrizeRegistrySession) Hashes(arg0 *big.Int) (string, error) {
	return _PrizeRegistry.Contract.Hashes(&_PrizeRegistry.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_PrizeRegistry *PrizeRegistryCallerSession) Hashes(arg0 *big.Int) (string, error) {
	return _PrizeRegistry.Contract.Hashes(&_PrizeRegistry.CallOpts, arg0)
}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_PrizeRegistry *PrizeRegistryCaller) Recorded(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _PrizeRegistry.contract.Call(opts, &out, "recorded", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_PrizeRegistry *PrizeRegistrySession) Recorded(arg0 *big.Int) (bool, error) {
	return _PrizeRegistry.Contract.Recorded(&_PrizeRegistry.CallOpts, arg0)
}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_PrizeRegistry *PrizeRegistryCallerSession) Recorded(arg0 *big.Int) (bool, error) {
	return _PrizeRegistry.Contract.Recorded(&_PrizeRegistry.CallOpts, arg0)
}

// ValidatePrize is a free data retrieval call binding the contract method 0x63a66be2.
//
// Solidity: function validatePrize(uint256 id, string recordHash) view returns(bool)
func (_PrizeRegistry *PrizeRegistryCaller) ValidatePrize(opts *bind.CallOpts, id *big.Int, recordHash string) (bool, error) {
	var out []interface{}
	err := _PrizeRegistry.contract.Call(opts, &out, "validatePrize", id, recordHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidatePrize is a free data retrieval call binding the contract method 0x63a66be2.
//
// Solidity: function validatePrize(uint256 id, string recordHash) view returns(bool)
func (_PrizeRegistry *PrizeRegistrySession) ValidatePrize(id *big.Int, recordHash string) (bool, error) {
	return _PrizeRegistry.Contract.ValidatePrize(&_PrizeRegistry.CallOpts, id, recordHash)
}

// ValidatePrize is a free data retrieval call binding the contract method 0x63a66be2.
//
// Solidity: function validatePrize(uint256 id, string recordHash) view returns(bool)
func (_PrizeRegistry *PrizeRegistryCallerSession) ValidatePrize(id *big.Int, recordHash string) (bool, error) {
	return _PrizeRegistry.Contract.ValidatePrize(&_PrizeRegistry.CallOpts, id, recordHash)
}

// RecordPrize is a paid mutator transaction binding the contract method 0xab84ec9f.
//
// Solidity: function recordPrize(uint256 id, string recordHash) returns()
func (_PrizeRegistry *PrizeRegistryTransactor) RecordPrize(opts *bind.TransactOpts, id *big.Int, recordHash string) (*types.Transaction, error) {
	return _PrizeRegistry.contract.Transact(opts, "recordPrize", id, recordHash)
}

// RecordPrize is a paid mutator transaction binding the contract method 0xab84ec9f.
//
// Solidity: function recordPrize(uint256 id, string recordHash) returns()
func (_PrizeRegistry *PrizeRegistrySession) RecordPrize(id *big.Int, recordHash string) (*types.Transaction, error) {
	return _PrizeRegistry.Contract.RecordPrize(&_PrizeRegistry.TransactOpts, id, recordHash)
}

// RecordPrize is a paid mutator transaction binding the contract method 0xab84ec9f.
//
// Solidity: function recordPrize(uint256 id, string recordHash) returns()
func (_PrizeRegistry *PrizeRegistryTransactorSession) RecordPrize(id *big.Int, recordHash string) (*types.Transaction, error) {
	return _PrizeRegistry.Contract.RecordPrize(&_PrizeRegistry.TransactOpts, id, recordHash)
}

// PrizeRegistryPrizeRecordedIterator is returned from FilterPrizeRecorded and is used to iterate over the raw logs and unpacked data for PrizeRecorded events raised by the PrizeRegistry contract.
type PrizeRegistryPrizeRecordedIterator struct {
	Event *PrizeRegistryPrizeRecorded // Event containing the contract specifics and raw log

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
func (it *PrizeRegistryPrizeRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizeRegistryPrizeRecorded)
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
		it.Event = new(PrizeRegistryPrizeRecorded)
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
func (it *PrizeRegistryPrizeRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizeRegistryPrizeRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizeRegistryPrizeRecorded represents a PrizeRecorded event raised by the PrizeRegistry contract.
type PrizeRegistryPrizeRecorded struct {
	Id         *big.Int
	RecordHash string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPrizeRecorded is a free log retrieval operation binding the contract event 0x2f8720dba077ce08b29ef00ca011bd0420304cd53e1795486dbf734022d4c659.
//
// Solidity: event PrizeRecorded(uint256 indexed id, string recordHash)
func (_PrizeRegistry *PrizeRegistryFilterer) FilterPrizeRecorded(opts *bind.FilterOpts, id []*big.Int) (*PrizeRegistryPrizeRecordedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _PrizeRegistry.contract.FilterLogs(opts, "PrizeRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return &PrizeRegistryPrizeRecordedIterator{contract: _PrizeRegistry.contract, event: "PrizeRecorded", logs: logs, sub: sub}, nil
}

// WatchPrizeRecorded is a free log subscription operation binding the contract event 0x2f8720dba077ce08b29ef00ca011bd0420304cd53e1795486dbf734022d4c659.
//
// Solidity: event PrizeRecorded(uint256 indexed id, string recordHash)
func (_PrizeRegistry *PrizeRegistryFilterer) WatchPrizeRecorded(opts *bind.WatchOpts, sink chan<- *PrizeRegistryPrizeRecorded, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _PrizeRegistry.contract.WatchLogs(opts, "PrizeRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizeRegistryPrizeRecorded)
				if err := _PrizeRegistry.contract.UnpackLog(event, "PrizeRecorded", log); err != nil {
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

// ParsePrizeRecorded is a log parse operation binding the contract event 0x2f8720dba077ce08b29ef00ca011bd0420304cd53e1795486dbf734022d4c659.
//
// Solidity: event PrizeRecorded(uint256 indexed id, string recordHash)
func (_PrizeRegistry *PrizeRegistryFilterer) ParsePrizeRecorded(log types.Log) (*PrizeRegistryPrizeRecorded, error) {
	event := new(PrizeRegistryPrizeRecorded)
	if err := _PrizeRegistry.contract.UnpackLog(event, "PrizeRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SportRegistryMetaData contains all meta data concerning the SportRegistry contract.
var SportRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"SportRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"recordSport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recorded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"validateSport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506105818061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c80632806ebd41461004e578063501895ae1461007657806360b2267b146100965780639f3c80b8146100b8575b5f5ffd5b61006161005c366004610288565b6100cd565b60405190151581526020015b60405180910390f35b6100896100843660046102ff565b61012a565b60405161006d9190610316565b6100616100a43660046102ff565b5f6020819052908152604090205460ff1681565b6100cb6100c6366004610288565b6101c1565b005b5f8381526020819052604081205460ff168015610122575082826040516100f592919061034b565b604080519182900382205f87815260016020529190912090916101189190610392565b6040518091039020145b949350505050565b60016020525f9081526040902080546101429061035a565b80601f016020809104026020016040519081016040528092919081815260200182805461016e9061035a565b80156101b95780601f10610190576101008083540402835291602001916101b9565b820191905f5260205f20905b81548152906001019060200180831161019c57829003601f168201915b505050505081565b5f8381526020819052604090205460ff161561021c5760405162461bcd60e51b815260206004820152601660248201527514dc1bdc9d08185b1c9958591e481c9958dbdc99195960521b604482015260640160405180910390fd5b5f83815260208181526040808320805460ff191660019081179091559091529020610248828483610463565b50827f85f8c8eab2cb1e62ec63641891be02349d978d3bfda05de067d6cfd4eb5b7aaf838360405161027b92919061051d565b60405180910390a2505050565b5f5f5f6040848603121561029a575f5ffd5b83359250602084013567ffffffffffffffff8111156102b7575f5ffd5b8401601f810186136102c7575f5ffd5b803567ffffffffffffffff8111156102dd575f5ffd5b8660208284010111156102ee575f5ffd5b939660209190910195509293505050565b5f6020828403121561030f575f5ffd5b5035919050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b818382375f9101908152919050565b600181811c9082168061036e57607f821691505b60208210810361038c57634e487b7160e01b5f52602260045260245ffd5b50919050565b5f5f835461039f8161035a565b6001821680156103b657600181146103cb576103f8565b60ff19831686528115158202860193506103f8565b865f5260205f205f5b838110156103f0578154888201526001909101906020016103d4565b505081860193505b509195945050505050565b634e487b7160e01b5f52604160045260245ffd5b601f82111561045e57805f5260205f20601f840160051c8101602085101561043c5750805b601f840160051c820191505b8181101561045b575f8155600101610448565b50505b505050565b67ffffffffffffffff83111561047b5761047b610403565b61048f83610489835461035a565b83610417565b5f601f8411600181146104c0575f85156104a95750838201355b5f19600387901b1c1916600186901b17835561045b565b5f83815260208120601f198716915b828110156104ef57868501358255602094850194600190920191016104cf565b508682101561050b575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b60208152816020820152818360408301375f818301604090810191909152601f909201601f1916010191905056fea2646970667358221220af045c8195ef3f50a55601e3522c587c974b0cd7810840b7d2bd00875bd21f1264736f6c634300081e0033",
}

// SportRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use SportRegistryMetaData.ABI instead.
var SportRegistryABI = SportRegistryMetaData.ABI

// SportRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SportRegistryMetaData.Bin instead.
var SportRegistryBin = SportRegistryMetaData.Bin

// DeploySportRegistry deploys a new Ethereum contract, binding an instance of SportRegistry to it.
func DeploySportRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SportRegistry, error) {
	parsed, err := SportRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SportRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SportRegistry{SportRegistryCaller: SportRegistryCaller{contract: contract}, SportRegistryTransactor: SportRegistryTransactor{contract: contract}, SportRegistryFilterer: SportRegistryFilterer{contract: contract}}, nil
}

// SportRegistry is an auto generated Go binding around an Ethereum contract.
type SportRegistry struct {
	SportRegistryCaller     // Read-only binding to the contract
	SportRegistryTransactor // Write-only binding to the contract
	SportRegistryFilterer   // Log filterer for contract events
}

// SportRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SportRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SportRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SportRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SportRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SportRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SportRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SportRegistrySession struct {
	Contract     *SportRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SportRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SportRegistryCallerSession struct {
	Contract *SportRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SportRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SportRegistryTransactorSession struct {
	Contract     *SportRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SportRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SportRegistryRaw struct {
	Contract *SportRegistry // Generic contract binding to access the raw methods on
}

// SportRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SportRegistryCallerRaw struct {
	Contract *SportRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// SportRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SportRegistryTransactorRaw struct {
	Contract *SportRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSportRegistry creates a new instance of SportRegistry, bound to a specific deployed contract.
func NewSportRegistry(address common.Address, backend bind.ContractBackend) (*SportRegistry, error) {
	contract, err := bindSportRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SportRegistry{SportRegistryCaller: SportRegistryCaller{contract: contract}, SportRegistryTransactor: SportRegistryTransactor{contract: contract}, SportRegistryFilterer: SportRegistryFilterer{contract: contract}}, nil
}

// NewSportRegistryCaller creates a new read-only instance of SportRegistry, bound to a specific deployed contract.
func NewSportRegistryCaller(address common.Address, caller bind.ContractCaller) (*SportRegistryCaller, error) {
	contract, err := bindSportRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SportRegistryCaller{contract: contract}, nil
}

// NewSportRegistryTransactor creates a new write-only instance of SportRegistry, bound to a specific deployed contract.
func NewSportRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*SportRegistryTransactor, error) {
	contract, err := bindSportRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SportRegistryTransactor{contract: contract}, nil
}

// NewSportRegistryFilterer creates a new log filterer instance of SportRegistry, bound to a specific deployed contract.
func NewSportRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*SportRegistryFilterer, error) {
	contract, err := bindSportRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SportRegistryFilterer{contract: contract}, nil
}

// bindSportRegistry binds a generic wrapper to an already deployed contract.
func bindSportRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SportRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SportRegistry *SportRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SportRegistry.Contract.SportRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SportRegistry *SportRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SportRegistry.Contract.SportRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SportRegistry *SportRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SportRegistry.Contract.SportRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SportRegistry *SportRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SportRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SportRegistry *SportRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SportRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SportRegistry *SportRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SportRegistry.Contract.contract.Transact(opts, method, params...)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_SportRegistry *SportRegistryCaller) Hashes(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _SportRegistry.contract.Call(opts, &out, "hashes", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_SportRegistry *SportRegistrySession) Hashes(arg0 *big.Int) (string, error) {
	return _SportRegistry.Contract.Hashes(&_SportRegistry.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_SportRegistry *SportRegistryCallerSession) Hashes(arg0 *big.Int) (string, error) {
	return _SportRegistry.Contract.Hashes(&_SportRegistry.CallOpts, arg0)
}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_SportRegistry *SportRegistryCaller) Recorded(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _SportRegistry.contract.Call(opts, &out, "recorded", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_SportRegistry *SportRegistrySession) Recorded(arg0 *big.Int) (bool, error) {
	return _SportRegistry.Contract.Recorded(&_SportRegistry.CallOpts, arg0)
}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_SportRegistry *SportRegistryCallerSession) Recorded(arg0 *big.Int) (bool, error) {
	return _SportRegistry.Contract.Recorded(&_SportRegistry.CallOpts, arg0)
}

// ValidateSport is a free data retrieval call binding the contract method 0x2806ebd4.
//
// Solidity: function validateSport(uint256 id, string recordHash) view returns(bool)
func (_SportRegistry *SportRegistryCaller) ValidateSport(opts *bind.CallOpts, id *big.Int, recordHash string) (bool, error) {
	var out []interface{}
	err := _SportRegistry.contract.Call(opts, &out, "validateSport", id, recordHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateSport is a free data retrieval call binding the contract method 0x2806ebd4.
//
// Solidity: function validateSport(uint256 id, string recordHash) view returns(bool)
func (_SportRegistry *SportRegistrySession) ValidateSport(id *big.Int, recordHash string) (bool, error) {
	return _SportRegistry.Contract.ValidateSport(&_SportRegistry.CallOpts, id, recordHash)
}

// ValidateSport is a free data retrieval call binding the contract method 0x2806ebd4.
//
// Solidity: function validateSport(uint256 id, string recordHash) view returns(bool)
func (_SportRegistry *SportRegistryCallerSession) ValidateSport(id *big.Int, recordHash string) (bool, error) {
	return _SportRegistry.Contract.ValidateSport(&_SportRegistry.CallOpts, id, recordHash)
}

// RecordSport is a paid mutator transaction binding the contract method 0x9f3c80b8.
//
// Solidity: function recordSport(uint256 id, string recordHash) returns()
func (_SportRegistry *SportRegistryTransactor) RecordSport(opts *bind.TransactOpts, id *big.Int, recordHash string) (*types.Transaction, error) {
	return _SportRegistry.contract.Transact(opts, "recordSport", id, recordHash)
}

// RecordSport is a paid mutator transaction binding the contract method 0x9f3c80b8.
//
// Solidity: function recordSport(uint256 id, string recordHash) returns()
func (_SportRegistry *SportRegistrySession) RecordSport(id *big.Int, recordHash string) (*types.Transaction, error) {
	return _SportRegistry.Contract.RecordSport(&_SportRegistry.TransactOpts, id, recordHash)
}

// RecordSport is a paid mutator transaction binding the contract method 0x9f3c80b8.
//
// Solidity: function recordSport(uint256 id, string recordHash) returns()
func (_SportRegistry *SportRegistryTransactorSession) RecordSport(id *big.Int, recordHash string) (*types.Transaction, error) {
	return _SportRegistry.Contract.RecordSport(&_SportRegistry.TransactOpts, id, recordHash)
}

// SportRegistrySportRecordedIterator is returned from FilterSportRecorded and is used to iterate over the raw logs and unpacked data for SportRecorded events raised by the SportRegistry contract.
type SportRegistrySportRecordedIterator struct {
	Event *SportRegistrySportRecorded // Event containing the contract specifics and raw log

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
func (it *SportRegistrySportRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SportRegistrySportRecorded)
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
		it.Event = new(SportRegistrySportRecorded)
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
func (it *SportRegistrySportRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SportRegistrySportRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SportRegistrySportRecorded represents a SportRecorded event raised by the SportRegistry contract.
type SportRegistrySportRecorded struct {
	Id         *big.Int
	RecordHash string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSportRecorded is a free log retrieval operation binding the contract event 0x85f8c8eab2cb1e62ec63641891be02349d978d3bfda05de067d6cfd4eb5b7aaf.
//
// Solidity: event SportRecorded(uint256 indexed id, string recordHash)
func (_SportRegistry *SportRegistryFilterer) FilterSportRecorded(opts *bind.FilterOpts, id []*big.Int) (*SportRegistrySportRecordedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SportRegistry.contract.FilterLogs(opts, "SportRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return &SportRegistrySportRecordedIterator{contract: _SportRegistry.contract, event: "SportRecorded", logs: logs, sub: sub}, nil
}

// WatchSportRecorded is a free log subscription operation binding the contract event 0x85f8c8eab2cb1e62ec63641891be02349d978d3bfda05de067d6cfd4eb5b7aaf.
//
// Solidity: event SportRecorded(uint256 indexed id, string recordHash)
func (_SportRegistry *SportRegistryFilterer) WatchSportRecorded(opts *bind.WatchOpts, sink chan<- *SportRegistrySportRecorded, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SportRegistry.contract.WatchLogs(opts, "SportRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SportRegistrySportRecorded)
				if err := _SportRegistry.contract.UnpackLog(event, "SportRecorded", log); err != nil {
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

// ParseSportRecorded is a log parse operation binding the contract event 0x85f8c8eab2cb1e62ec63641891be02349d978d3bfda05de067d6cfd4eb5b7aaf.
//
// Solidity: event SportRecorded(uint256 indexed id, string recordHash)
func (_SportRegistry *SportRegistryFilterer) ParseSportRecorded(log types.Log) (*SportRegistrySportRecorded, error) {
	event := new(SportRegistrySportRecorded)
	if err := _SportRegistry.contract.UnpackLog(event, "SportRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeamAchievementsRegistryMetaData contains all meta data concerning the TeamAchievementsRegistry contract.
var TeamAchievementsRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"TeamAchievementsRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"recordTeamAchievements\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recorded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"validateTeamAchievements\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506105968061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c8063501895ae1461004e57806360b2267b146100775780636d4f4153146100a95780637659e820146100be575b5f5ffd5b61006161005c36600461029d565b6100d1565b60405161006e91906102b4565b60405180910390f35b61009961008536600461029d565b5f6020819052908152604090205460ff1681565b604051901515815260200161006e565b6100bc6100b73660046102e9565b610168565b005b6100996100cc3660046102e9565b610240565b60016020525f9081526040902080546100e990610360565b80601f016020809104026020016040519081016040528092919081815260200182805461011590610360565b80156101605780601f1061013757610100808354040283529160200191610160565b820191905f5260205f20905b81548152906001019060200180831161014357829003601f168201915b505050505081565b5f8381526020819052604090205460ff16156101d45760405162461bcd60e51b815260206004820152602160248201527f5465616d416368696576656d656e747320616c7265616479207265636f7264656044820152601960fa1b606482015260840160405180910390fd5b5f83815260208181526040808320805460ff1916600190811790915590915290206102008284836103f8565b50827fb0c1034ffefd8772b51b2995a0cf2233ab4a1a7135761b86fbee097ceaac232883836040516102339291906104b2565b60405180910390a2505050565b5f8381526020819052604081205460ff168015610295575082826040516102689291906104e0565b604080519182900382205f878152600160205291909120909161028b91906104ef565b6040518091039020145b949350505050565b5f602082840312156102ad575f5ffd5b5035919050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f5f5f604084860312156102fb575f5ffd5b83359250602084013567ffffffffffffffff811115610318575f5ffd5b8401601f81018613610328575f5ffd5b803567ffffffffffffffff81111561033e575f5ffd5b86602082840101111561034f575f5ffd5b939660209190910195509293505050565b600181811c9082168061037457607f821691505b60208210810361039257634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52604160045260245ffd5b601f8211156103f357805f5260205f20601f840160051c810160208510156103d15750805b601f840160051c820191505b818110156103f0575f81556001016103dd565b50505b505050565b67ffffffffffffffff83111561041057610410610398565b6104248361041e8354610360565b836103ac565b5f601f841160018114610455575f851561043e5750838201355b5f19600387901b1c1916600186901b1783556103f0565b5f83815260208120601f198716915b828110156104845786850135825560209485019460019092019101610464565b50868210156104a0575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b60208152816020820152818360408301375f818301604090810191909152601f909201601f19160101919050565b818382375f9101908152919050565b5f5f83546104fc81610360565b600182168015610513576001811461052857610555565b60ff1983168652811515820286019350610555565b865f5260205f205f5b8381101561054d57815488820152600190910190602001610531565b505081860193505b50919594505050505056fea2646970667358221220dc135980ab05ef1e1b363e67a939aa1b207318179255be5a16a7761e1a14503064736f6c634300081e0033",
}

// TeamAchievementsRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use TeamAchievementsRegistryMetaData.ABI instead.
var TeamAchievementsRegistryABI = TeamAchievementsRegistryMetaData.ABI

// TeamAchievementsRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TeamAchievementsRegistryMetaData.Bin instead.
var TeamAchievementsRegistryBin = TeamAchievementsRegistryMetaData.Bin

// DeployTeamAchievementsRegistry deploys a new Ethereum contract, binding an instance of TeamAchievementsRegistry to it.
func DeployTeamAchievementsRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TeamAchievementsRegistry, error) {
	parsed, err := TeamAchievementsRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TeamAchievementsRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TeamAchievementsRegistry{TeamAchievementsRegistryCaller: TeamAchievementsRegistryCaller{contract: contract}, TeamAchievementsRegistryTransactor: TeamAchievementsRegistryTransactor{contract: contract}, TeamAchievementsRegistryFilterer: TeamAchievementsRegistryFilterer{contract: contract}}, nil
}

// TeamAchievementsRegistry is an auto generated Go binding around an Ethereum contract.
type TeamAchievementsRegistry struct {
	TeamAchievementsRegistryCaller     // Read-only binding to the contract
	TeamAchievementsRegistryTransactor // Write-only binding to the contract
	TeamAchievementsRegistryFilterer   // Log filterer for contract events
}

// TeamAchievementsRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeamAchievementsRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeamAchievementsRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeamAchievementsRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeamAchievementsRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeamAchievementsRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeamAchievementsRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeamAchievementsRegistrySession struct {
	Contract     *TeamAchievementsRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TeamAchievementsRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeamAchievementsRegistryCallerSession struct {
	Contract *TeamAchievementsRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// TeamAchievementsRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeamAchievementsRegistryTransactorSession struct {
	Contract     *TeamAchievementsRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// TeamAchievementsRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeamAchievementsRegistryRaw struct {
	Contract *TeamAchievementsRegistry // Generic contract binding to access the raw methods on
}

// TeamAchievementsRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeamAchievementsRegistryCallerRaw struct {
	Contract *TeamAchievementsRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// TeamAchievementsRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeamAchievementsRegistryTransactorRaw struct {
	Contract *TeamAchievementsRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeamAchievementsRegistry creates a new instance of TeamAchievementsRegistry, bound to a specific deployed contract.
func NewTeamAchievementsRegistry(address common.Address, backend bind.ContractBackend) (*TeamAchievementsRegistry, error) {
	contract, err := bindTeamAchievementsRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeamAchievementsRegistry{TeamAchievementsRegistryCaller: TeamAchievementsRegistryCaller{contract: contract}, TeamAchievementsRegistryTransactor: TeamAchievementsRegistryTransactor{contract: contract}, TeamAchievementsRegistryFilterer: TeamAchievementsRegistryFilterer{contract: contract}}, nil
}

// NewTeamAchievementsRegistryCaller creates a new read-only instance of TeamAchievementsRegistry, bound to a specific deployed contract.
func NewTeamAchievementsRegistryCaller(address common.Address, caller bind.ContractCaller) (*TeamAchievementsRegistryCaller, error) {
	contract, err := bindTeamAchievementsRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeamAchievementsRegistryCaller{contract: contract}, nil
}

// NewTeamAchievementsRegistryTransactor creates a new write-only instance of TeamAchievementsRegistry, bound to a specific deployed contract.
func NewTeamAchievementsRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*TeamAchievementsRegistryTransactor, error) {
	contract, err := bindTeamAchievementsRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeamAchievementsRegistryTransactor{contract: contract}, nil
}

// NewTeamAchievementsRegistryFilterer creates a new log filterer instance of TeamAchievementsRegistry, bound to a specific deployed contract.
func NewTeamAchievementsRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*TeamAchievementsRegistryFilterer, error) {
	contract, err := bindTeamAchievementsRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeamAchievementsRegistryFilterer{contract: contract}, nil
}

// bindTeamAchievementsRegistry binds a generic wrapper to an already deployed contract.
func bindTeamAchievementsRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeamAchievementsRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeamAchievementsRegistry *TeamAchievementsRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeamAchievementsRegistry.Contract.TeamAchievementsRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeamAchievementsRegistry *TeamAchievementsRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeamAchievementsRegistry.Contract.TeamAchievementsRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeamAchievementsRegistry *TeamAchievementsRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeamAchievementsRegistry.Contract.TeamAchievementsRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeamAchievementsRegistry *TeamAchievementsRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeamAchievementsRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeamAchievementsRegistry *TeamAchievementsRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeamAchievementsRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeamAchievementsRegistry *TeamAchievementsRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeamAchievementsRegistry.Contract.contract.Transact(opts, method, params...)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_TeamAchievementsRegistry *TeamAchievementsRegistryCaller) Hashes(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _TeamAchievementsRegistry.contract.Call(opts, &out, "hashes", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_TeamAchievementsRegistry *TeamAchievementsRegistrySession) Hashes(arg0 *big.Int) (string, error) {
	return _TeamAchievementsRegistry.Contract.Hashes(&_TeamAchievementsRegistry.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_TeamAchievementsRegistry *TeamAchievementsRegistryCallerSession) Hashes(arg0 *big.Int) (string, error) {
	return _TeamAchievementsRegistry.Contract.Hashes(&_TeamAchievementsRegistry.CallOpts, arg0)
}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_TeamAchievementsRegistry *TeamAchievementsRegistryCaller) Recorded(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _TeamAchievementsRegistry.contract.Call(opts, &out, "recorded", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_TeamAchievementsRegistry *TeamAchievementsRegistrySession) Recorded(arg0 *big.Int) (bool, error) {
	return _TeamAchievementsRegistry.Contract.Recorded(&_TeamAchievementsRegistry.CallOpts, arg0)
}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_TeamAchievementsRegistry *TeamAchievementsRegistryCallerSession) Recorded(arg0 *big.Int) (bool, error) {
	return _TeamAchievementsRegistry.Contract.Recorded(&_TeamAchievementsRegistry.CallOpts, arg0)
}

// ValidateTeamAchievements is a free data retrieval call binding the contract method 0x7659e820.
//
// Solidity: function validateTeamAchievements(uint256 id, string recordHash) view returns(bool)
func (_TeamAchievementsRegistry *TeamAchievementsRegistryCaller) ValidateTeamAchievements(opts *bind.CallOpts, id *big.Int, recordHash string) (bool, error) {
	var out []interface{}
	err := _TeamAchievementsRegistry.contract.Call(opts, &out, "validateTeamAchievements", id, recordHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateTeamAchievements is a free data retrieval call binding the contract method 0x7659e820.
//
// Solidity: function validateTeamAchievements(uint256 id, string recordHash) view returns(bool)
func (_TeamAchievementsRegistry *TeamAchievementsRegistrySession) ValidateTeamAchievements(id *big.Int, recordHash string) (bool, error) {
	return _TeamAchievementsRegistry.Contract.ValidateTeamAchievements(&_TeamAchievementsRegistry.CallOpts, id, recordHash)
}

// ValidateTeamAchievements is a free data retrieval call binding the contract method 0x7659e820.
//
// Solidity: function validateTeamAchievements(uint256 id, string recordHash) view returns(bool)
func (_TeamAchievementsRegistry *TeamAchievementsRegistryCallerSession) ValidateTeamAchievements(id *big.Int, recordHash string) (bool, error) {
	return _TeamAchievementsRegistry.Contract.ValidateTeamAchievements(&_TeamAchievementsRegistry.CallOpts, id, recordHash)
}

// RecordTeamAchievements is a paid mutator transaction binding the contract method 0x6d4f4153.
//
// Solidity: function recordTeamAchievements(uint256 id, string recordHash) returns()
func (_TeamAchievementsRegistry *TeamAchievementsRegistryTransactor) RecordTeamAchievements(opts *bind.TransactOpts, id *big.Int, recordHash string) (*types.Transaction, error) {
	return _TeamAchievementsRegistry.contract.Transact(opts, "recordTeamAchievements", id, recordHash)
}

// RecordTeamAchievements is a paid mutator transaction binding the contract method 0x6d4f4153.
//
// Solidity: function recordTeamAchievements(uint256 id, string recordHash) returns()
func (_TeamAchievementsRegistry *TeamAchievementsRegistrySession) RecordTeamAchievements(id *big.Int, recordHash string) (*types.Transaction, error) {
	return _TeamAchievementsRegistry.Contract.RecordTeamAchievements(&_TeamAchievementsRegistry.TransactOpts, id, recordHash)
}

// RecordTeamAchievements is a paid mutator transaction binding the contract method 0x6d4f4153.
//
// Solidity: function recordTeamAchievements(uint256 id, string recordHash) returns()
func (_TeamAchievementsRegistry *TeamAchievementsRegistryTransactorSession) RecordTeamAchievements(id *big.Int, recordHash string) (*types.Transaction, error) {
	return _TeamAchievementsRegistry.Contract.RecordTeamAchievements(&_TeamAchievementsRegistry.TransactOpts, id, recordHash)
}

// TeamAchievementsRegistryTeamAchievementsRecordedIterator is returned from FilterTeamAchievementsRecorded and is used to iterate over the raw logs and unpacked data for TeamAchievementsRecorded events raised by the TeamAchievementsRegistry contract.
type TeamAchievementsRegistryTeamAchievementsRecordedIterator struct {
	Event *TeamAchievementsRegistryTeamAchievementsRecorded // Event containing the contract specifics and raw log

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
func (it *TeamAchievementsRegistryTeamAchievementsRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeamAchievementsRegistryTeamAchievementsRecorded)
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
		it.Event = new(TeamAchievementsRegistryTeamAchievementsRecorded)
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
func (it *TeamAchievementsRegistryTeamAchievementsRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeamAchievementsRegistryTeamAchievementsRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeamAchievementsRegistryTeamAchievementsRecorded represents a TeamAchievementsRecorded event raised by the TeamAchievementsRegistry contract.
type TeamAchievementsRegistryTeamAchievementsRecorded struct {
	Id         *big.Int
	RecordHash string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTeamAchievementsRecorded is a free log retrieval operation binding the contract event 0xb0c1034ffefd8772b51b2995a0cf2233ab4a1a7135761b86fbee097ceaac2328.
//
// Solidity: event TeamAchievementsRecorded(uint256 indexed id, string recordHash)
func (_TeamAchievementsRegistry *TeamAchievementsRegistryFilterer) FilterTeamAchievementsRecorded(opts *bind.FilterOpts, id []*big.Int) (*TeamAchievementsRegistryTeamAchievementsRecordedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TeamAchievementsRegistry.contract.FilterLogs(opts, "TeamAchievementsRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return &TeamAchievementsRegistryTeamAchievementsRecordedIterator{contract: _TeamAchievementsRegistry.contract, event: "TeamAchievementsRecorded", logs: logs, sub: sub}, nil
}

// WatchTeamAchievementsRecorded is a free log subscription operation binding the contract event 0xb0c1034ffefd8772b51b2995a0cf2233ab4a1a7135761b86fbee097ceaac2328.
//
// Solidity: event TeamAchievementsRecorded(uint256 indexed id, string recordHash)
func (_TeamAchievementsRegistry *TeamAchievementsRegistryFilterer) WatchTeamAchievementsRecorded(opts *bind.WatchOpts, sink chan<- *TeamAchievementsRegistryTeamAchievementsRecorded, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TeamAchievementsRegistry.contract.WatchLogs(opts, "TeamAchievementsRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeamAchievementsRegistryTeamAchievementsRecorded)
				if err := _TeamAchievementsRegistry.contract.UnpackLog(event, "TeamAchievementsRecorded", log); err != nil {
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

// ParseTeamAchievementsRecorded is a log parse operation binding the contract event 0xb0c1034ffefd8772b51b2995a0cf2233ab4a1a7135761b86fbee097ceaac2328.
//
// Solidity: event TeamAchievementsRecorded(uint256 indexed id, string recordHash)
func (_TeamAchievementsRegistry *TeamAchievementsRegistryFilterer) ParseTeamAchievementsRecorded(log types.Log) (*TeamAchievementsRegistryTeamAchievementsRecorded, error) {
	event := new(TeamAchievementsRegistryTeamAchievementsRecorded)
	if err := _TeamAchievementsRegistry.contract.UnpackLog(event, "TeamAchievementsRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeamPersonRegistryMetaData contains all meta data concerning the TeamPersonRegistry contract.
var TeamPersonRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"TeamPersonRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"recordTeamPerson\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recorded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"validateTeamPerson\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506105888061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c80631101f6751461004e5780634bea644614610076578063501895ae1461008b57806360b2267b146100ab575b5f5ffd5b61006161005c36600461028f565b6100cd565b60405190151581526020015b60405180910390f35b61008961008436600461028f565b61012a565b005b61009e610099366004610306565b6101f8565b60405161006d919061031d565b6100616100b9366004610306565b5f6020819052908152604090205460ff1681565b5f8381526020819052604081205460ff168015610122575082826040516100f5929190610352565b604080519182900382205f87815260016020529190912090916101189190610399565b6040518091039020145b949350505050565b5f8381526020819052604090205460ff161561018c5760405162461bcd60e51b815260206004820152601b60248201527f5465616d506572736f6e20616c7265616479207265636f726465640000000000604482015260640160405180910390fd5b5f83815260208181526040808320805460ff1916600190811790915590915290206101b882848361046a565b50827fa949f116d73f2c1a318d60c6fecf77c0f66e5394ae0345f20f9c5429a1924fc983836040516101eb929190610524565b60405180910390a2505050565b60016020525f90815260409020805461021090610361565b80601f016020809104026020016040519081016040528092919081815260200182805461023c90610361565b80156102875780601f1061025e57610100808354040283529160200191610287565b820191905f5260205f20905b81548152906001019060200180831161026a57829003601f168201915b505050505081565b5f5f5f604084860312156102a1575f5ffd5b83359250602084013567ffffffffffffffff8111156102be575f5ffd5b8401601f810186136102ce575f5ffd5b803567ffffffffffffffff8111156102e4575f5ffd5b8660208284010111156102f5575f5ffd5b939660209190910195509293505050565b5f60208284031215610316575f5ffd5b5035919050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b818382375f9101908152919050565b600181811c9082168061037557607f821691505b60208210810361039357634e487b7160e01b5f52602260045260245ffd5b50919050565b5f5f83546103a681610361565b6001821680156103bd57600181146103d2576103ff565b60ff19831686528115158202860193506103ff565b865f5260205f205f5b838110156103f7578154888201526001909101906020016103db565b505081860193505b509195945050505050565b634e487b7160e01b5f52604160045260245ffd5b601f82111561046557805f5260205f20601f840160051c810160208510156104435750805b601f840160051c820191505b81811015610462575f815560010161044f565b50505b505050565b67ffffffffffffffff8311156104825761048261040a565b610496836104908354610361565b8361041e565b5f601f8411600181146104c7575f85156104b05750838201355b5f19600387901b1c1916600186901b178355610462565b5f83815260208120601f198716915b828110156104f657868501358255602094850194600190920191016104d6565b5086821015610512575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b60208152816020820152818360408301375f818301604090810191909152601f909201601f1916010191905056fea26469706673582212200f071b05c35bcbc2a045940e675d14cc76c1dd22762819220d50d954825b03ca64736f6c634300081e0033",
}

// TeamPersonRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use TeamPersonRegistryMetaData.ABI instead.
var TeamPersonRegistryABI = TeamPersonRegistryMetaData.ABI

// TeamPersonRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TeamPersonRegistryMetaData.Bin instead.
var TeamPersonRegistryBin = TeamPersonRegistryMetaData.Bin

// DeployTeamPersonRegistry deploys a new Ethereum contract, binding an instance of TeamPersonRegistry to it.
func DeployTeamPersonRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TeamPersonRegistry, error) {
	parsed, err := TeamPersonRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TeamPersonRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TeamPersonRegistry{TeamPersonRegistryCaller: TeamPersonRegistryCaller{contract: contract}, TeamPersonRegistryTransactor: TeamPersonRegistryTransactor{contract: contract}, TeamPersonRegistryFilterer: TeamPersonRegistryFilterer{contract: contract}}, nil
}

// TeamPersonRegistry is an auto generated Go binding around an Ethereum contract.
type TeamPersonRegistry struct {
	TeamPersonRegistryCaller     // Read-only binding to the contract
	TeamPersonRegistryTransactor // Write-only binding to the contract
	TeamPersonRegistryFilterer   // Log filterer for contract events
}

// TeamPersonRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeamPersonRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeamPersonRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeamPersonRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeamPersonRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeamPersonRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeamPersonRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeamPersonRegistrySession struct {
	Contract     *TeamPersonRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TeamPersonRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeamPersonRegistryCallerSession struct {
	Contract *TeamPersonRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TeamPersonRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeamPersonRegistryTransactorSession struct {
	Contract     *TeamPersonRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TeamPersonRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeamPersonRegistryRaw struct {
	Contract *TeamPersonRegistry // Generic contract binding to access the raw methods on
}

// TeamPersonRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeamPersonRegistryCallerRaw struct {
	Contract *TeamPersonRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// TeamPersonRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeamPersonRegistryTransactorRaw struct {
	Contract *TeamPersonRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeamPersonRegistry creates a new instance of TeamPersonRegistry, bound to a specific deployed contract.
func NewTeamPersonRegistry(address common.Address, backend bind.ContractBackend) (*TeamPersonRegistry, error) {
	contract, err := bindTeamPersonRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeamPersonRegistry{TeamPersonRegistryCaller: TeamPersonRegistryCaller{contract: contract}, TeamPersonRegistryTransactor: TeamPersonRegistryTransactor{contract: contract}, TeamPersonRegistryFilterer: TeamPersonRegistryFilterer{contract: contract}}, nil
}

// NewTeamPersonRegistryCaller creates a new read-only instance of TeamPersonRegistry, bound to a specific deployed contract.
func NewTeamPersonRegistryCaller(address common.Address, caller bind.ContractCaller) (*TeamPersonRegistryCaller, error) {
	contract, err := bindTeamPersonRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeamPersonRegistryCaller{contract: contract}, nil
}

// NewTeamPersonRegistryTransactor creates a new write-only instance of TeamPersonRegistry, bound to a specific deployed contract.
func NewTeamPersonRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*TeamPersonRegistryTransactor, error) {
	contract, err := bindTeamPersonRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeamPersonRegistryTransactor{contract: contract}, nil
}

// NewTeamPersonRegistryFilterer creates a new log filterer instance of TeamPersonRegistry, bound to a specific deployed contract.
func NewTeamPersonRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*TeamPersonRegistryFilterer, error) {
	contract, err := bindTeamPersonRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeamPersonRegistryFilterer{contract: contract}, nil
}

// bindTeamPersonRegistry binds a generic wrapper to an already deployed contract.
func bindTeamPersonRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeamPersonRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeamPersonRegistry *TeamPersonRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeamPersonRegistry.Contract.TeamPersonRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeamPersonRegistry *TeamPersonRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeamPersonRegistry.Contract.TeamPersonRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeamPersonRegistry *TeamPersonRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeamPersonRegistry.Contract.TeamPersonRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeamPersonRegistry *TeamPersonRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeamPersonRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeamPersonRegistry *TeamPersonRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeamPersonRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeamPersonRegistry *TeamPersonRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeamPersonRegistry.Contract.contract.Transact(opts, method, params...)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_TeamPersonRegistry *TeamPersonRegistryCaller) Hashes(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _TeamPersonRegistry.contract.Call(opts, &out, "hashes", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_TeamPersonRegistry *TeamPersonRegistrySession) Hashes(arg0 *big.Int) (string, error) {
	return _TeamPersonRegistry.Contract.Hashes(&_TeamPersonRegistry.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_TeamPersonRegistry *TeamPersonRegistryCallerSession) Hashes(arg0 *big.Int) (string, error) {
	return _TeamPersonRegistry.Contract.Hashes(&_TeamPersonRegistry.CallOpts, arg0)
}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_TeamPersonRegistry *TeamPersonRegistryCaller) Recorded(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _TeamPersonRegistry.contract.Call(opts, &out, "recorded", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_TeamPersonRegistry *TeamPersonRegistrySession) Recorded(arg0 *big.Int) (bool, error) {
	return _TeamPersonRegistry.Contract.Recorded(&_TeamPersonRegistry.CallOpts, arg0)
}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_TeamPersonRegistry *TeamPersonRegistryCallerSession) Recorded(arg0 *big.Int) (bool, error) {
	return _TeamPersonRegistry.Contract.Recorded(&_TeamPersonRegistry.CallOpts, arg0)
}

// ValidateTeamPerson is a free data retrieval call binding the contract method 0x1101f675.
//
// Solidity: function validateTeamPerson(uint256 id, string recordHash) view returns(bool)
func (_TeamPersonRegistry *TeamPersonRegistryCaller) ValidateTeamPerson(opts *bind.CallOpts, id *big.Int, recordHash string) (bool, error) {
	var out []interface{}
	err := _TeamPersonRegistry.contract.Call(opts, &out, "validateTeamPerson", id, recordHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateTeamPerson is a free data retrieval call binding the contract method 0x1101f675.
//
// Solidity: function validateTeamPerson(uint256 id, string recordHash) view returns(bool)
func (_TeamPersonRegistry *TeamPersonRegistrySession) ValidateTeamPerson(id *big.Int, recordHash string) (bool, error) {
	return _TeamPersonRegistry.Contract.ValidateTeamPerson(&_TeamPersonRegistry.CallOpts, id, recordHash)
}

// ValidateTeamPerson is a free data retrieval call binding the contract method 0x1101f675.
//
// Solidity: function validateTeamPerson(uint256 id, string recordHash) view returns(bool)
func (_TeamPersonRegistry *TeamPersonRegistryCallerSession) ValidateTeamPerson(id *big.Int, recordHash string) (bool, error) {
	return _TeamPersonRegistry.Contract.ValidateTeamPerson(&_TeamPersonRegistry.CallOpts, id, recordHash)
}

// RecordTeamPerson is a paid mutator transaction binding the contract method 0x4bea6446.
//
// Solidity: function recordTeamPerson(uint256 id, string recordHash) returns()
func (_TeamPersonRegistry *TeamPersonRegistryTransactor) RecordTeamPerson(opts *bind.TransactOpts, id *big.Int, recordHash string) (*types.Transaction, error) {
	return _TeamPersonRegistry.contract.Transact(opts, "recordTeamPerson", id, recordHash)
}

// RecordTeamPerson is a paid mutator transaction binding the contract method 0x4bea6446.
//
// Solidity: function recordTeamPerson(uint256 id, string recordHash) returns()
func (_TeamPersonRegistry *TeamPersonRegistrySession) RecordTeamPerson(id *big.Int, recordHash string) (*types.Transaction, error) {
	return _TeamPersonRegistry.Contract.RecordTeamPerson(&_TeamPersonRegistry.TransactOpts, id, recordHash)
}

// RecordTeamPerson is a paid mutator transaction binding the contract method 0x4bea6446.
//
// Solidity: function recordTeamPerson(uint256 id, string recordHash) returns()
func (_TeamPersonRegistry *TeamPersonRegistryTransactorSession) RecordTeamPerson(id *big.Int, recordHash string) (*types.Transaction, error) {
	return _TeamPersonRegistry.Contract.RecordTeamPerson(&_TeamPersonRegistry.TransactOpts, id, recordHash)
}

// TeamPersonRegistryTeamPersonRecordedIterator is returned from FilterTeamPersonRecorded and is used to iterate over the raw logs and unpacked data for TeamPersonRecorded events raised by the TeamPersonRegistry contract.
type TeamPersonRegistryTeamPersonRecordedIterator struct {
	Event *TeamPersonRegistryTeamPersonRecorded // Event containing the contract specifics and raw log

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
func (it *TeamPersonRegistryTeamPersonRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeamPersonRegistryTeamPersonRecorded)
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
		it.Event = new(TeamPersonRegistryTeamPersonRecorded)
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
func (it *TeamPersonRegistryTeamPersonRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeamPersonRegistryTeamPersonRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeamPersonRegistryTeamPersonRecorded represents a TeamPersonRecorded event raised by the TeamPersonRegistry contract.
type TeamPersonRegistryTeamPersonRecorded struct {
	Id         *big.Int
	RecordHash string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTeamPersonRecorded is a free log retrieval operation binding the contract event 0xa949f116d73f2c1a318d60c6fecf77c0f66e5394ae0345f20f9c5429a1924fc9.
//
// Solidity: event TeamPersonRecorded(uint256 indexed id, string recordHash)
func (_TeamPersonRegistry *TeamPersonRegistryFilterer) FilterTeamPersonRecorded(opts *bind.FilterOpts, id []*big.Int) (*TeamPersonRegistryTeamPersonRecordedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TeamPersonRegistry.contract.FilterLogs(opts, "TeamPersonRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return &TeamPersonRegistryTeamPersonRecordedIterator{contract: _TeamPersonRegistry.contract, event: "TeamPersonRecorded", logs: logs, sub: sub}, nil
}

// WatchTeamPersonRecorded is a free log subscription operation binding the contract event 0xa949f116d73f2c1a318d60c6fecf77c0f66e5394ae0345f20f9c5429a1924fc9.
//
// Solidity: event TeamPersonRecorded(uint256 indexed id, string recordHash)
func (_TeamPersonRegistry *TeamPersonRegistryFilterer) WatchTeamPersonRecorded(opts *bind.WatchOpts, sink chan<- *TeamPersonRegistryTeamPersonRecorded, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TeamPersonRegistry.contract.WatchLogs(opts, "TeamPersonRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeamPersonRegistryTeamPersonRecorded)
				if err := _TeamPersonRegistry.contract.UnpackLog(event, "TeamPersonRecorded", log); err != nil {
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

// ParseTeamPersonRecorded is a log parse operation binding the contract event 0xa949f116d73f2c1a318d60c6fecf77c0f66e5394ae0345f20f9c5429a1924fc9.
//
// Solidity: event TeamPersonRecorded(uint256 indexed id, string recordHash)
func (_TeamPersonRegistry *TeamPersonRegistryFilterer) ParseTeamPersonRecorded(log types.Log) (*TeamPersonRegistryTeamPersonRecorded, error) {
	event := new(TeamPersonRegistryTeamPersonRecorded)
	if err := _TeamPersonRegistry.contract.UnpackLog(event, "TeamPersonRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeamRegistryMetaData contains all meta data concerning the TeamRegistry contract.
var TeamRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"TeamRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"recordTeam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recorded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recordHash\",\"type\":\"string\"}],\"name\":\"validateTeam\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506105848061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c8063501895ae1461004e57806360b2267b1461007757806378768229146100a957806395fe1e7e146100bc575b5f5ffd5b61006161005c36600461028b565b6100d1565b60405161006e91906102a2565b60405180910390f35b61009961008536600461028b565b5f6020819052908152604090205460ff1681565b604051901515815260200161006e565b6100996100b73660046102d7565b610168565b6100cf6100ca3660046102d7565b6101c5565b005b60016020525f9081526040902080546100e99061034e565b80601f01602080910402602001604051908101604052809291908181526020018280546101159061034e565b80156101605780601f1061013757610100808354040283529160200191610160565b820191905f5260205f20905b81548152906001019060200180831161014357829003601f168201915b505050505081565b5f8381526020819052604081205460ff1680156101bd57508282604051610190929190610386565b604080519182900382205f87815260016020529190912090916101b39190610395565b6040518091039020145b949350505050565b5f8381526020819052604090205460ff161561021f5760405162461bcd60e51b81526020600482015260156024820152741519585b48185b1c9958591e481c9958dbdc991959605a1b604482015260640160405180910390fd5b5f83815260208181526040808320805460ff19166001908117909155909152902061024b828483610466565b50827f4c10eb58093e6afcff1f799a75ce04b158a68b5b1b9fe4d67ae71d86556e51c3838360405161027e929190610520565b60405180910390a2505050565b5f6020828403121561029b575f5ffd5b5035919050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f5f5f604084860312156102e9575f5ffd5b83359250602084013567ffffffffffffffff811115610306575f5ffd5b8401601f81018613610316575f5ffd5b803567ffffffffffffffff81111561032c575f5ffd5b86602082840101111561033d575f5ffd5b939660209190910195509293505050565b600181811c9082168061036257607f821691505b60208210810361038057634e487b7160e01b5f52602260045260245ffd5b50919050565b818382375f9101908152919050565b5f5f83546103a28161034e565b6001821680156103b957600181146103ce576103fb565b60ff19831686528115158202860193506103fb565b865f5260205f205f5b838110156103f3578154888201526001909101906020016103d7565b505081860193505b509195945050505050565b634e487b7160e01b5f52604160045260245ffd5b601f82111561046157805f5260205f20601f840160051c8101602085101561043f5750805b601f840160051c820191505b8181101561045e575f815560010161044b565b50505b505050565b67ffffffffffffffff83111561047e5761047e610406565b6104928361048c835461034e565b8361041a565b5f601f8411600181146104c3575f85156104ac5750838201355b5f19600387901b1c1916600186901b17835561045e565b5f83815260208120601f198716915b828110156104f257868501358255602094850194600190920191016104d2565b508682101561050e575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b60208152816020820152818360408301375f818301604090810191909152601f909201601f1916010191905056fea2646970667358221220f5371004ca7f0cebbf43f1c7274a5706b286ac83f4e64385a6d0543afa18d24d64736f6c634300081e0033",
}

// TeamRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use TeamRegistryMetaData.ABI instead.
var TeamRegistryABI = TeamRegistryMetaData.ABI

// TeamRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TeamRegistryMetaData.Bin instead.
var TeamRegistryBin = TeamRegistryMetaData.Bin

// DeployTeamRegistry deploys a new Ethereum contract, binding an instance of TeamRegistry to it.
func DeployTeamRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TeamRegistry, error) {
	parsed, err := TeamRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TeamRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TeamRegistry{TeamRegistryCaller: TeamRegistryCaller{contract: contract}, TeamRegistryTransactor: TeamRegistryTransactor{contract: contract}, TeamRegistryFilterer: TeamRegistryFilterer{contract: contract}}, nil
}

// TeamRegistry is an auto generated Go binding around an Ethereum contract.
type TeamRegistry struct {
	TeamRegistryCaller     // Read-only binding to the contract
	TeamRegistryTransactor // Write-only binding to the contract
	TeamRegistryFilterer   // Log filterer for contract events
}

// TeamRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeamRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeamRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeamRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeamRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeamRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeamRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeamRegistrySession struct {
	Contract     *TeamRegistry     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeamRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeamRegistryCallerSession struct {
	Contract *TeamRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TeamRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeamRegistryTransactorSession struct {
	Contract     *TeamRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TeamRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeamRegistryRaw struct {
	Contract *TeamRegistry // Generic contract binding to access the raw methods on
}

// TeamRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeamRegistryCallerRaw struct {
	Contract *TeamRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// TeamRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeamRegistryTransactorRaw struct {
	Contract *TeamRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeamRegistry creates a new instance of TeamRegistry, bound to a specific deployed contract.
func NewTeamRegistry(address common.Address, backend bind.ContractBackend) (*TeamRegistry, error) {
	contract, err := bindTeamRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeamRegistry{TeamRegistryCaller: TeamRegistryCaller{contract: contract}, TeamRegistryTransactor: TeamRegistryTransactor{contract: contract}, TeamRegistryFilterer: TeamRegistryFilterer{contract: contract}}, nil
}

// NewTeamRegistryCaller creates a new read-only instance of TeamRegistry, bound to a specific deployed contract.
func NewTeamRegistryCaller(address common.Address, caller bind.ContractCaller) (*TeamRegistryCaller, error) {
	contract, err := bindTeamRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeamRegistryCaller{contract: contract}, nil
}

// NewTeamRegistryTransactor creates a new write-only instance of TeamRegistry, bound to a specific deployed contract.
func NewTeamRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*TeamRegistryTransactor, error) {
	contract, err := bindTeamRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeamRegistryTransactor{contract: contract}, nil
}

// NewTeamRegistryFilterer creates a new log filterer instance of TeamRegistry, bound to a specific deployed contract.
func NewTeamRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*TeamRegistryFilterer, error) {
	contract, err := bindTeamRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeamRegistryFilterer{contract: contract}, nil
}

// bindTeamRegistry binds a generic wrapper to an already deployed contract.
func bindTeamRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeamRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeamRegistry *TeamRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeamRegistry.Contract.TeamRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeamRegistry *TeamRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeamRegistry.Contract.TeamRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeamRegistry *TeamRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeamRegistry.Contract.TeamRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeamRegistry *TeamRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeamRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeamRegistry *TeamRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeamRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeamRegistry *TeamRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeamRegistry.Contract.contract.Transact(opts, method, params...)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_TeamRegistry *TeamRegistryCaller) Hashes(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _TeamRegistry.contract.Call(opts, &out, "hashes", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_TeamRegistry *TeamRegistrySession) Hashes(arg0 *big.Int) (string, error) {
	return _TeamRegistry.Contract.Hashes(&_TeamRegistry.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(string)
func (_TeamRegistry *TeamRegistryCallerSession) Hashes(arg0 *big.Int) (string, error) {
	return _TeamRegistry.Contract.Hashes(&_TeamRegistry.CallOpts, arg0)
}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_TeamRegistry *TeamRegistryCaller) Recorded(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _TeamRegistry.contract.Call(opts, &out, "recorded", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_TeamRegistry *TeamRegistrySession) Recorded(arg0 *big.Int) (bool, error) {
	return _TeamRegistry.Contract.Recorded(&_TeamRegistry.CallOpts, arg0)
}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_TeamRegistry *TeamRegistryCallerSession) Recorded(arg0 *big.Int) (bool, error) {
	return _TeamRegistry.Contract.Recorded(&_TeamRegistry.CallOpts, arg0)
}

// ValidateTeam is a free data retrieval call binding the contract method 0x78768229.
//
// Solidity: function validateTeam(uint256 id, string recordHash) view returns(bool)
func (_TeamRegistry *TeamRegistryCaller) ValidateTeam(opts *bind.CallOpts, id *big.Int, recordHash string) (bool, error) {
	var out []interface{}
	err := _TeamRegistry.contract.Call(opts, &out, "validateTeam", id, recordHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateTeam is a free data retrieval call binding the contract method 0x78768229.
//
// Solidity: function validateTeam(uint256 id, string recordHash) view returns(bool)
func (_TeamRegistry *TeamRegistrySession) ValidateTeam(id *big.Int, recordHash string) (bool, error) {
	return _TeamRegistry.Contract.ValidateTeam(&_TeamRegistry.CallOpts, id, recordHash)
}

// ValidateTeam is a free data retrieval call binding the contract method 0x78768229.
//
// Solidity: function validateTeam(uint256 id, string recordHash) view returns(bool)
func (_TeamRegistry *TeamRegistryCallerSession) ValidateTeam(id *big.Int, recordHash string) (bool, error) {
	return _TeamRegistry.Contract.ValidateTeam(&_TeamRegistry.CallOpts, id, recordHash)
}

// RecordTeam is a paid mutator transaction binding the contract method 0x95fe1e7e.
//
// Solidity: function recordTeam(uint256 id, string recordHash) returns()
func (_TeamRegistry *TeamRegistryTransactor) RecordTeam(opts *bind.TransactOpts, id *big.Int, recordHash string) (*types.Transaction, error) {
	return _TeamRegistry.contract.Transact(opts, "recordTeam", id, recordHash)
}

// RecordTeam is a paid mutator transaction binding the contract method 0x95fe1e7e.
//
// Solidity: function recordTeam(uint256 id, string recordHash) returns()
func (_TeamRegistry *TeamRegistrySession) RecordTeam(id *big.Int, recordHash string) (*types.Transaction, error) {
	return _TeamRegistry.Contract.RecordTeam(&_TeamRegistry.TransactOpts, id, recordHash)
}

// RecordTeam is a paid mutator transaction binding the contract method 0x95fe1e7e.
//
// Solidity: function recordTeam(uint256 id, string recordHash) returns()
func (_TeamRegistry *TeamRegistryTransactorSession) RecordTeam(id *big.Int, recordHash string) (*types.Transaction, error) {
	return _TeamRegistry.Contract.RecordTeam(&_TeamRegistry.TransactOpts, id, recordHash)
}

// TeamRegistryTeamRecordedIterator is returned from FilterTeamRecorded and is used to iterate over the raw logs and unpacked data for TeamRecorded events raised by the TeamRegistry contract.
type TeamRegistryTeamRecordedIterator struct {
	Event *TeamRegistryTeamRecorded // Event containing the contract specifics and raw log

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
func (it *TeamRegistryTeamRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeamRegistryTeamRecorded)
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
		it.Event = new(TeamRegistryTeamRecorded)
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
func (it *TeamRegistryTeamRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeamRegistryTeamRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeamRegistryTeamRecorded represents a TeamRecorded event raised by the TeamRegistry contract.
type TeamRegistryTeamRecorded struct {
	Id         *big.Int
	RecordHash string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTeamRecorded is a free log retrieval operation binding the contract event 0x4c10eb58093e6afcff1f799a75ce04b158a68b5b1b9fe4d67ae71d86556e51c3.
//
// Solidity: event TeamRecorded(uint256 indexed id, string recordHash)
func (_TeamRegistry *TeamRegistryFilterer) FilterTeamRecorded(opts *bind.FilterOpts, id []*big.Int) (*TeamRegistryTeamRecordedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TeamRegistry.contract.FilterLogs(opts, "TeamRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return &TeamRegistryTeamRecordedIterator{contract: _TeamRegistry.contract, event: "TeamRecorded", logs: logs, sub: sub}, nil
}

// WatchTeamRecorded is a free log subscription operation binding the contract event 0x4c10eb58093e6afcff1f799a75ce04b158a68b5b1b9fe4d67ae71d86556e51c3.
//
// Solidity: event TeamRecorded(uint256 indexed id, string recordHash)
func (_TeamRegistry *TeamRegistryFilterer) WatchTeamRecorded(opts *bind.WatchOpts, sink chan<- *TeamRegistryTeamRecorded, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TeamRegistry.contract.WatchLogs(opts, "TeamRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeamRegistryTeamRecorded)
				if err := _TeamRegistry.contract.UnpackLog(event, "TeamRecorded", log); err != nil {
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

// ParseTeamRecorded is a log parse operation binding the contract event 0x4c10eb58093e6afcff1f799a75ce04b158a68b5b1b9fe4d67ae71d86556e51c3.
//
// Solidity: event TeamRecorded(uint256 indexed id, string recordHash)
func (_TeamRegistry *TeamRegistryFilterer) ParseTeamRecorded(log types.Log) (*TeamRegistryTeamRecorded, error) {
	event := new(TeamRegistryTeamRecorded)
	if err := _TeamRegistry.contract.UnpackLog(event, "TeamRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

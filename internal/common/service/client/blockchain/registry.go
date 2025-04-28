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

// CompetitionTeamsRegistryMetaData contains all meta data concerning the CompetitionTeamsRegistry contract.
var CompetitionTeamsRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"CompetitionTeamsRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"recordCompetitionTeams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recorded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"validateCompetitionTeams\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061023e8061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c8063334e90af1461004e5780634516146e14610063578063501895ae1461008b57806360b2267b146100b8575b5f5ffd5b61006161005c3660046101d1565b6100da565b005b6100766100713660046101d1565b6101a1565b60405190151581526020015b60405180910390f35b6100aa6100993660046101f1565b60016020525f908152604090205481565b604051908152602001610082565b6100766100c63660046101f1565b5f6020819052908152604090205460ff1681565b5f8281526020819052604090205460ff16156101465760405162461bcd60e51b815260206004820152602160248201527f436f6d7065746974696f6e5465616d7320616c7265616479207265636f7264656044820152601960fa1b606482015260840160405180910390fd5b5f82815260208181526040808320805460ff19166001908117909155825291829020839055905182815283917fcca48f312d5dfa82c50dfc7934b24c6f1b3bdc730987f6c36180693a526c9d78910160405180910390a25050565b5f8281526020819052604081205460ff1680156101ca57505f8381526001602052604090205482145b9392505050565b5f5f604083850312156101e2575f5ffd5b50508035926020909101359150565b5f60208284031215610201575f5ffd5b503591905056fea26469706673582212200f7485969a8dc4b42a4c48e6be5ab4105ce23a0d819f894673851cfc8cf1269964736f6c634300081d0033",
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
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistryCaller) Hashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _CompetitionTeamsRegistry.contract.Call(opts, &out, "hashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistrySession) Hashes(arg0 *big.Int) ([32]byte, error) {
	return _CompetitionTeamsRegistry.Contract.Hashes(&_CompetitionTeamsRegistry.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistryCallerSession) Hashes(arg0 *big.Int) ([32]byte, error) {
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

// ValidateCompetitionTeams is a free data retrieval call binding the contract method 0x4516146e.
//
// Solidity: function validateCompetitionTeams(uint256 id, bytes32 recordHash) view returns(bool)
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistryCaller) ValidateCompetitionTeams(opts *bind.CallOpts, id *big.Int, recordHash [32]byte) (bool, error) {
	var out []interface{}
	err := _CompetitionTeamsRegistry.contract.Call(opts, &out, "validateCompetitionTeams", id, recordHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateCompetitionTeams is a free data retrieval call binding the contract method 0x4516146e.
//
// Solidity: function validateCompetitionTeams(uint256 id, bytes32 recordHash) view returns(bool)
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistrySession) ValidateCompetitionTeams(id *big.Int, recordHash [32]byte) (bool, error) {
	return _CompetitionTeamsRegistry.Contract.ValidateCompetitionTeams(&_CompetitionTeamsRegistry.CallOpts, id, recordHash)
}

// ValidateCompetitionTeams is a free data retrieval call binding the contract method 0x4516146e.
//
// Solidity: function validateCompetitionTeams(uint256 id, bytes32 recordHash) view returns(bool)
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistryCallerSession) ValidateCompetitionTeams(id *big.Int, recordHash [32]byte) (bool, error) {
	return _CompetitionTeamsRegistry.Contract.ValidateCompetitionTeams(&_CompetitionTeamsRegistry.CallOpts, id, recordHash)
}

// RecordCompetitionTeams is a paid mutator transaction binding the contract method 0x334e90af.
//
// Solidity: function recordCompetitionTeams(uint256 id, bytes32 recordHash) returns()
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistryTransactor) RecordCompetitionTeams(opts *bind.TransactOpts, id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
	return _CompetitionTeamsRegistry.contract.Transact(opts, "recordCompetitionTeams", id, recordHash)
}

// RecordCompetitionTeams is a paid mutator transaction binding the contract method 0x334e90af.
//
// Solidity: function recordCompetitionTeams(uint256 id, bytes32 recordHash) returns()
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistrySession) RecordCompetitionTeams(id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
	return _CompetitionTeamsRegistry.Contract.RecordCompetitionTeams(&_CompetitionTeamsRegistry.TransactOpts, id, recordHash)
}

// RecordCompetitionTeams is a paid mutator transaction binding the contract method 0x334e90af.
//
// Solidity: function recordCompetitionTeams(uint256 id, bytes32 recordHash) returns()
func (_CompetitionTeamsRegistry *CompetitionTeamsRegistryTransactorSession) RecordCompetitionTeams(id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
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
	RecordHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCompetitionTeamsRecorded is a free log retrieval operation binding the contract event 0xcca48f312d5dfa82c50dfc7934b24c6f1b3bdc730987f6c36180693a526c9d78.
//
// Solidity: event CompetitionTeamsRecorded(uint256 indexed id, bytes32 recordHash)
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

// WatchCompetitionTeamsRecorded is a free log subscription operation binding the contract event 0xcca48f312d5dfa82c50dfc7934b24c6f1b3bdc730987f6c36180693a526c9d78.
//
// Solidity: event CompetitionTeamsRecorded(uint256 indexed id, bytes32 recordHash)
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

// ParseCompetitionTeamsRecorded is a log parse operation binding the contract event 0xcca48f312d5dfa82c50dfc7934b24c6f1b3bdc730987f6c36180693a526c9d78.
//
// Solidity: event CompetitionTeamsRecorded(uint256 indexed id, bytes32 recordHash)
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"MatchParticipantRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"recordMatchParticipant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recorded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"validateMatchParticipant\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061023e8061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c8063501895ae1461004e57806360b2267b1461008057806370a9bd7c146100b2578063a33b5273146100c5575b5f5ffd5b61006d61005c3660046101d1565b60016020525f908152604090205481565b6040519081526020015b60405180910390f35b6100a261008e3660046101d1565b5f6020819052908152604090205460ff1681565b6040519015158152602001610077565b6100a26100c03660046101e8565b6100da565b6100d86100d33660046101e8565b61010a565b005b5f8281526020819052604081205460ff16801561010357505f8381526001602052604090205482145b9392505050565b5f8281526020819052604090205460ff16156101765760405162461bcd60e51b815260206004820152602160248201527f4d617463685061727469636970616e7420616c7265616479207265636f7264656044820152601960fa1b606482015260840160405180910390fd5b5f82815260208181526040808320805460ff19166001908117909155825291829020839055905182815283917f8150f40e33aac4e36316694127820faf2edce46c48e34b3ea472d2027b1e97b2910160405180910390a25050565b5f602082840312156101e1575f5ffd5b5035919050565b5f5f604083850312156101f9575f5ffd5b5050803592602090910135915056fea2646970667358221220f2de66d6d9fef9d6e2564cc48984d4cb0f3eda92af0fae8e72bd297117cc816864736f6c634300081d0033",
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
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_MatchParticipantRegistry *MatchParticipantRegistryCaller) Hashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MatchParticipantRegistry.contract.Call(opts, &out, "hashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_MatchParticipantRegistry *MatchParticipantRegistrySession) Hashes(arg0 *big.Int) ([32]byte, error) {
	return _MatchParticipantRegistry.Contract.Hashes(&_MatchParticipantRegistry.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_MatchParticipantRegistry *MatchParticipantRegistryCallerSession) Hashes(arg0 *big.Int) ([32]byte, error) {
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

// ValidateMatchParticipant is a free data retrieval call binding the contract method 0x70a9bd7c.
//
// Solidity: function validateMatchParticipant(uint256 id, bytes32 recordHash) view returns(bool)
func (_MatchParticipantRegistry *MatchParticipantRegistryCaller) ValidateMatchParticipant(opts *bind.CallOpts, id *big.Int, recordHash [32]byte) (bool, error) {
	var out []interface{}
	err := _MatchParticipantRegistry.contract.Call(opts, &out, "validateMatchParticipant", id, recordHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateMatchParticipant is a free data retrieval call binding the contract method 0x70a9bd7c.
//
// Solidity: function validateMatchParticipant(uint256 id, bytes32 recordHash) view returns(bool)
func (_MatchParticipantRegistry *MatchParticipantRegistrySession) ValidateMatchParticipant(id *big.Int, recordHash [32]byte) (bool, error) {
	return _MatchParticipantRegistry.Contract.ValidateMatchParticipant(&_MatchParticipantRegistry.CallOpts, id, recordHash)
}

// ValidateMatchParticipant is a free data retrieval call binding the contract method 0x70a9bd7c.
//
// Solidity: function validateMatchParticipant(uint256 id, bytes32 recordHash) view returns(bool)
func (_MatchParticipantRegistry *MatchParticipantRegistryCallerSession) ValidateMatchParticipant(id *big.Int, recordHash [32]byte) (bool, error) {
	return _MatchParticipantRegistry.Contract.ValidateMatchParticipant(&_MatchParticipantRegistry.CallOpts, id, recordHash)
}

// RecordMatchParticipant is a paid mutator transaction binding the contract method 0xa33b5273.
//
// Solidity: function recordMatchParticipant(uint256 id, bytes32 recordHash) returns()
func (_MatchParticipantRegistry *MatchParticipantRegistryTransactor) RecordMatchParticipant(opts *bind.TransactOpts, id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
	return _MatchParticipantRegistry.contract.Transact(opts, "recordMatchParticipant", id, recordHash)
}

// RecordMatchParticipant is a paid mutator transaction binding the contract method 0xa33b5273.
//
// Solidity: function recordMatchParticipant(uint256 id, bytes32 recordHash) returns()
func (_MatchParticipantRegistry *MatchParticipantRegistrySession) RecordMatchParticipant(id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
	return _MatchParticipantRegistry.Contract.RecordMatchParticipant(&_MatchParticipantRegistry.TransactOpts, id, recordHash)
}

// RecordMatchParticipant is a paid mutator transaction binding the contract method 0xa33b5273.
//
// Solidity: function recordMatchParticipant(uint256 id, bytes32 recordHash) returns()
func (_MatchParticipantRegistry *MatchParticipantRegistryTransactorSession) RecordMatchParticipant(id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
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
	RecordHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMatchParticipantRecorded is a free log retrieval operation binding the contract event 0x8150f40e33aac4e36316694127820faf2edce46c48e34b3ea472d2027b1e97b2.
//
// Solidity: event MatchParticipantRecorded(uint256 indexed id, bytes32 recordHash)
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

// WatchMatchParticipantRecorded is a free log subscription operation binding the contract event 0x8150f40e33aac4e36316694127820faf2edce46c48e34b3ea472d2027b1e97b2.
//
// Solidity: event MatchParticipantRecorded(uint256 indexed id, bytes32 recordHash)
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

// ParseMatchParticipantRecorded is a log parse operation binding the contract event 0x8150f40e33aac4e36316694127820faf2edce46c48e34b3ea472d2027b1e97b2.
//
// Solidity: event MatchParticipantRecorded(uint256 indexed id, bytes32 recordHash)
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"MatchRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"recordMatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recorded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"validateMatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061022d8061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c8063501895ae1461004e57806360b2267b1461008057806394e40e9a146100b2578063dc01da5b146100c5575b5f5ffd5b61006d61005c3660046101c0565b60016020525f908152604090205481565b6040519081526020015b60405180910390f35b6100a261008e3660046101c0565b5f6020819052908152604090205460ff1681565b6040519015158152602001610077565b6100a26100c03660046101d7565b6100da565b6100d86100d33660046101d7565b61010a565b005b5f8281526020819052604081205460ff16801561010357505f8381526001602052604090205482145b9392505050565b5f8281526020819052604090205460ff16156101655760405162461bcd60e51b815260206004820152601660248201527513585d18da08185b1c9958591e481c9958dbdc99195960521b604482015260640160405180910390fd5b5f82815260208181526040808320805460ff19166001908117909155825291829020839055905182815283917fc8000df4b446153e43b5c2ff7a39fdcb7f2b5bec162aa92b4e17b740d604c171910160405180910390a25050565b5f602082840312156101d0575f5ffd5b5035919050565b5f5f604083850312156101e8575f5ffd5b5050803592602090910135915056fea2646970667358221220f354bc9d8fa86198e0def1eb842a62f724a8e105439e482ba4265e0f969173ee64736f6c634300081d0033",
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
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_MatchRegistry *MatchRegistryCaller) Hashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MatchRegistry.contract.Call(opts, &out, "hashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_MatchRegistry *MatchRegistrySession) Hashes(arg0 *big.Int) ([32]byte, error) {
	return _MatchRegistry.Contract.Hashes(&_MatchRegistry.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_MatchRegistry *MatchRegistryCallerSession) Hashes(arg0 *big.Int) ([32]byte, error) {
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

// ValidateMatch is a free data retrieval call binding the contract method 0x94e40e9a.
//
// Solidity: function validateMatch(uint256 id, bytes32 recordHash) view returns(bool)
func (_MatchRegistry *MatchRegistryCaller) ValidateMatch(opts *bind.CallOpts, id *big.Int, recordHash [32]byte) (bool, error) {
	var out []interface{}
	err := _MatchRegistry.contract.Call(opts, &out, "validateMatch", id, recordHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateMatch is a free data retrieval call binding the contract method 0x94e40e9a.
//
// Solidity: function validateMatch(uint256 id, bytes32 recordHash) view returns(bool)
func (_MatchRegistry *MatchRegistrySession) ValidateMatch(id *big.Int, recordHash [32]byte) (bool, error) {
	return _MatchRegistry.Contract.ValidateMatch(&_MatchRegistry.CallOpts, id, recordHash)
}

// ValidateMatch is a free data retrieval call binding the contract method 0x94e40e9a.
//
// Solidity: function validateMatch(uint256 id, bytes32 recordHash) view returns(bool)
func (_MatchRegistry *MatchRegistryCallerSession) ValidateMatch(id *big.Int, recordHash [32]byte) (bool, error) {
	return _MatchRegistry.Contract.ValidateMatch(&_MatchRegistry.CallOpts, id, recordHash)
}

// RecordMatch is a paid mutator transaction binding the contract method 0xdc01da5b.
//
// Solidity: function recordMatch(uint256 id, bytes32 recordHash) returns()
func (_MatchRegistry *MatchRegistryTransactor) RecordMatch(opts *bind.TransactOpts, id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
	return _MatchRegistry.contract.Transact(opts, "recordMatch", id, recordHash)
}

// RecordMatch is a paid mutator transaction binding the contract method 0xdc01da5b.
//
// Solidity: function recordMatch(uint256 id, bytes32 recordHash) returns()
func (_MatchRegistry *MatchRegistrySession) RecordMatch(id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
	return _MatchRegistry.Contract.RecordMatch(&_MatchRegistry.TransactOpts, id, recordHash)
}

// RecordMatch is a paid mutator transaction binding the contract method 0xdc01da5b.
//
// Solidity: function recordMatch(uint256 id, bytes32 recordHash) returns()
func (_MatchRegistry *MatchRegistryTransactorSession) RecordMatch(id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
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
	RecordHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMatchRecorded is a free log retrieval operation binding the contract event 0xc8000df4b446153e43b5c2ff7a39fdcb7f2b5bec162aa92b4e17b740d604c171.
//
// Solidity: event MatchRecorded(uint256 indexed id, bytes32 recordHash)
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

// WatchMatchRecorded is a free log subscription operation binding the contract event 0xc8000df4b446153e43b5c2ff7a39fdcb7f2b5bec162aa92b4e17b740d604c171.
//
// Solidity: event MatchRecorded(uint256 indexed id, bytes32 recordHash)
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

// ParseMatchRecorded is a log parse operation binding the contract event 0xc8000df4b446153e43b5c2ff7a39fdcb7f2b5bec162aa92b4e17b740d604c171.
//
// Solidity: event MatchRecorded(uint256 indexed id, bytes32 recordHash)
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"PersonRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"recordPerson\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recorded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"validatePerson\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506102348061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c8063501895ae1461004e57806360b2267b1461008057806372013409146100b2578063e600835d146100c7575b5f5ffd5b61006d61005c3660046101c7565b60016020525f908152604090205481565b6040519081526020015b60405180910390f35b6100a261008e3660046101c7565b5f6020819052908152604090205460ff1681565b6040519015158152602001610077565b6100c56100c03660046101de565b6100da565b005b6100a26100d53660046101de565b610197565b5f8281526020819052604090205460ff161561013c5760405162461bcd60e51b815260206004820152601760248201527f506572736f6e20616c7265616479207265636f72646564000000000000000000604482015260640160405180910390fd5b5f82815260208181526040808320805460ff19166001908117909155825291829020839055905182815283917fccc35dc586d377da368a5e0ae6527ddbd36d7520272350d11ddde3995c8e2ff9910160405180910390a25050565b5f8281526020819052604081205460ff1680156101c057505f8381526001602052604090205482145b9392505050565b5f602082840312156101d7575f5ffd5b5035919050565b5f5f604083850312156101ef575f5ffd5b5050803592602090910135915056fea2646970667358221220e2e7380765fd8ac737ea29d8474d9be4a116c3bf9f950a3965663dad2e39c1f164736f6c634300081d0033",
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
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_PersonRegistry *PersonRegistryCaller) Hashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _PersonRegistry.contract.Call(opts, &out, "hashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_PersonRegistry *PersonRegistrySession) Hashes(arg0 *big.Int) ([32]byte, error) {
	return _PersonRegistry.Contract.Hashes(&_PersonRegistry.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_PersonRegistry *PersonRegistryCallerSession) Hashes(arg0 *big.Int) ([32]byte, error) {
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

// ValidatePerson is a free data retrieval call binding the contract method 0xe600835d.
//
// Solidity: function validatePerson(uint256 id, bytes32 recordHash) view returns(bool)
func (_PersonRegistry *PersonRegistryCaller) ValidatePerson(opts *bind.CallOpts, id *big.Int, recordHash [32]byte) (bool, error) {
	var out []interface{}
	err := _PersonRegistry.contract.Call(opts, &out, "validatePerson", id, recordHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidatePerson is a free data retrieval call binding the contract method 0xe600835d.
//
// Solidity: function validatePerson(uint256 id, bytes32 recordHash) view returns(bool)
func (_PersonRegistry *PersonRegistrySession) ValidatePerson(id *big.Int, recordHash [32]byte) (bool, error) {
	return _PersonRegistry.Contract.ValidatePerson(&_PersonRegistry.CallOpts, id, recordHash)
}

// ValidatePerson is a free data retrieval call binding the contract method 0xe600835d.
//
// Solidity: function validatePerson(uint256 id, bytes32 recordHash) view returns(bool)
func (_PersonRegistry *PersonRegistryCallerSession) ValidatePerson(id *big.Int, recordHash [32]byte) (bool, error) {
	return _PersonRegistry.Contract.ValidatePerson(&_PersonRegistry.CallOpts, id, recordHash)
}

// RecordPerson is a paid mutator transaction binding the contract method 0x72013409.
//
// Solidity: function recordPerson(uint256 id, bytes32 recordHash) returns()
func (_PersonRegistry *PersonRegistryTransactor) RecordPerson(opts *bind.TransactOpts, id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
	return _PersonRegistry.contract.Transact(opts, "recordPerson", id, recordHash)
}

// RecordPerson is a paid mutator transaction binding the contract method 0x72013409.
//
// Solidity: function recordPerson(uint256 id, bytes32 recordHash) returns()
func (_PersonRegistry *PersonRegistrySession) RecordPerson(id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
	return _PersonRegistry.Contract.RecordPerson(&_PersonRegistry.TransactOpts, id, recordHash)
}

// RecordPerson is a paid mutator transaction binding the contract method 0x72013409.
//
// Solidity: function recordPerson(uint256 id, bytes32 recordHash) returns()
func (_PersonRegistry *PersonRegistryTransactorSession) RecordPerson(id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
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
	RecordHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPersonRecorded is a free log retrieval operation binding the contract event 0xccc35dc586d377da368a5e0ae6527ddbd36d7520272350d11ddde3995c8e2ff9.
//
// Solidity: event PersonRecorded(uint256 indexed id, bytes32 recordHash)
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

// WatchPersonRecorded is a free log subscription operation binding the contract event 0xccc35dc586d377da368a5e0ae6527ddbd36d7520272350d11ddde3995c8e2ff9.
//
// Solidity: event PersonRecorded(uint256 indexed id, bytes32 recordHash)
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

// ParsePersonRecorded is a log parse operation binding the contract event 0xccc35dc586d377da368a5e0ae6527ddbd36d7520272350d11ddde3995c8e2ff9.
//
// Solidity: event PersonRecorded(uint256 indexed id, bytes32 recordHash)
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"PrizeRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"recordPrize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recorded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"validatePrize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061022d8061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c8063501895ae1461004e57806360b2267b146100805780636e7a2303146100b2578063d6b2e79c146100c5575b5f5ffd5b61006d61005c3660046101c0565b60016020525f908152604090205481565b6040519081526020015b60405180910390f35b6100a261008e3660046101c0565b5f6020819052908152604090205460ff1681565b6040519015158152602001610077565b6100a26100c03660046101d7565b6100da565b6100d86100d33660046101d7565b61010a565b005b5f8281526020819052604081205460ff16801561010357505f8381526001602052604090205482145b9392505050565b5f8281526020819052604090205460ff16156101655760405162461bcd60e51b8152602060048201526016602482015275141c9a5e9948185b1c9958591e481c9958dbdc99195960521b604482015260640160405180910390fd5b5f82815260208181526040808320805460ff19166001908117909155825291829020839055905182815283917fe10db0f42bd2a99cbdf969539fe5a713c22dd4f295a12c3eff519e78e1909d36910160405180910390a25050565b5f602082840312156101d0575f5ffd5b5035919050565b5f5f604083850312156101e8575f5ffd5b5050803592602090910135915056fea2646970667358221220ea48d0de0a117bbb53e24a54d53d39dd897296a3f0a05494230c77e92d74bbc764736f6c634300081d0033",
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
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_PrizeRegistry *PrizeRegistryCaller) Hashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _PrizeRegistry.contract.Call(opts, &out, "hashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_PrizeRegistry *PrizeRegistrySession) Hashes(arg0 *big.Int) ([32]byte, error) {
	return _PrizeRegistry.Contract.Hashes(&_PrizeRegistry.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_PrizeRegistry *PrizeRegistryCallerSession) Hashes(arg0 *big.Int) ([32]byte, error) {
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

// ValidatePrize is a free data retrieval call binding the contract method 0x6e7a2303.
//
// Solidity: function validatePrize(uint256 id, bytes32 recordHash) view returns(bool)
func (_PrizeRegistry *PrizeRegistryCaller) ValidatePrize(opts *bind.CallOpts, id *big.Int, recordHash [32]byte) (bool, error) {
	var out []interface{}
	err := _PrizeRegistry.contract.Call(opts, &out, "validatePrize", id, recordHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidatePrize is a free data retrieval call binding the contract method 0x6e7a2303.
//
// Solidity: function validatePrize(uint256 id, bytes32 recordHash) view returns(bool)
func (_PrizeRegistry *PrizeRegistrySession) ValidatePrize(id *big.Int, recordHash [32]byte) (bool, error) {
	return _PrizeRegistry.Contract.ValidatePrize(&_PrizeRegistry.CallOpts, id, recordHash)
}

// ValidatePrize is a free data retrieval call binding the contract method 0x6e7a2303.
//
// Solidity: function validatePrize(uint256 id, bytes32 recordHash) view returns(bool)
func (_PrizeRegistry *PrizeRegistryCallerSession) ValidatePrize(id *big.Int, recordHash [32]byte) (bool, error) {
	return _PrizeRegistry.Contract.ValidatePrize(&_PrizeRegistry.CallOpts, id, recordHash)
}

// RecordPrize is a paid mutator transaction binding the contract method 0xd6b2e79c.
//
// Solidity: function recordPrize(uint256 id, bytes32 recordHash) returns()
func (_PrizeRegistry *PrizeRegistryTransactor) RecordPrize(opts *bind.TransactOpts, id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
	return _PrizeRegistry.contract.Transact(opts, "recordPrize", id, recordHash)
}

// RecordPrize is a paid mutator transaction binding the contract method 0xd6b2e79c.
//
// Solidity: function recordPrize(uint256 id, bytes32 recordHash) returns()
func (_PrizeRegistry *PrizeRegistrySession) RecordPrize(id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
	return _PrizeRegistry.Contract.RecordPrize(&_PrizeRegistry.TransactOpts, id, recordHash)
}

// RecordPrize is a paid mutator transaction binding the contract method 0xd6b2e79c.
//
// Solidity: function recordPrize(uint256 id, bytes32 recordHash) returns()
func (_PrizeRegistry *PrizeRegistryTransactorSession) RecordPrize(id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
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
	RecordHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPrizeRecorded is a free log retrieval operation binding the contract event 0xe10db0f42bd2a99cbdf969539fe5a713c22dd4f295a12c3eff519e78e1909d36.
//
// Solidity: event PrizeRecorded(uint256 indexed id, bytes32 recordHash)
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

// WatchPrizeRecorded is a free log subscription operation binding the contract event 0xe10db0f42bd2a99cbdf969539fe5a713c22dd4f295a12c3eff519e78e1909d36.
//
// Solidity: event PrizeRecorded(uint256 indexed id, bytes32 recordHash)
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

// ParsePrizeRecorded is a log parse operation binding the contract event 0xe10db0f42bd2a99cbdf969539fe5a713c22dd4f295a12c3eff519e78e1909d36.
//
// Solidity: event PrizeRecorded(uint256 indexed id, bytes32 recordHash)
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"SportRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"recordSport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recorded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"validateSport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061022d8061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c8063501895ae1461004e57806360b2267b14610080578063bec12103146100b2578063d76c9386146100c5575b5f5ffd5b61006d61005c3660046101c0565b60016020525f908152604090205481565b6040519081526020015b60405180910390f35b6100a261008e3660046101c0565b5f6020819052908152604090205460ff1681565b6040519015158152602001610077565b6100a26100c03660046101d7565b6100da565b6100d86100d33660046101d7565b61010a565b005b5f8281526020819052604081205460ff16801561010357505f8381526001602052604090205482145b9392505050565b5f8281526020819052604090205460ff16156101655760405162461bcd60e51b815260206004820152601660248201527514dc1bdc9d08185b1c9958591e481c9958dbdc99195960521b604482015260640160405180910390fd5b5f82815260208181526040808320805460ff19166001908117909155825291829020839055905182815283917f3a90657270ecf8ae1de2f87142a28a59f2ff9d42307e84a7e388ccab1413610b910160405180910390a25050565b5f602082840312156101d0575f5ffd5b5035919050565b5f5f604083850312156101e8575f5ffd5b5050803592602090910135915056fea2646970667358221220766dd5526676e7de22096c934e5ad557bff4573f8e5e7194b31a4becc07137db64736f6c634300081d0033",
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
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_SportRegistry *SportRegistryCaller) Hashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SportRegistry.contract.Call(opts, &out, "hashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_SportRegistry *SportRegistrySession) Hashes(arg0 *big.Int) ([32]byte, error) {
	return _SportRegistry.Contract.Hashes(&_SportRegistry.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_SportRegistry *SportRegistryCallerSession) Hashes(arg0 *big.Int) ([32]byte, error) {
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

// ValidateSport is a free data retrieval call binding the contract method 0xbec12103.
//
// Solidity: function validateSport(uint256 id, bytes32 recordHash) view returns(bool)
func (_SportRegistry *SportRegistryCaller) ValidateSport(opts *bind.CallOpts, id *big.Int, recordHash [32]byte) (bool, error) {
	var out []interface{}
	err := _SportRegistry.contract.Call(opts, &out, "validateSport", id, recordHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateSport is a free data retrieval call binding the contract method 0xbec12103.
//
// Solidity: function validateSport(uint256 id, bytes32 recordHash) view returns(bool)
func (_SportRegistry *SportRegistrySession) ValidateSport(id *big.Int, recordHash [32]byte) (bool, error) {
	return _SportRegistry.Contract.ValidateSport(&_SportRegistry.CallOpts, id, recordHash)
}

// ValidateSport is a free data retrieval call binding the contract method 0xbec12103.
//
// Solidity: function validateSport(uint256 id, bytes32 recordHash) view returns(bool)
func (_SportRegistry *SportRegistryCallerSession) ValidateSport(id *big.Int, recordHash [32]byte) (bool, error) {
	return _SportRegistry.Contract.ValidateSport(&_SportRegistry.CallOpts, id, recordHash)
}

// RecordSport is a paid mutator transaction binding the contract method 0xd76c9386.
//
// Solidity: function recordSport(uint256 id, bytes32 recordHash) returns()
func (_SportRegistry *SportRegistryTransactor) RecordSport(opts *bind.TransactOpts, id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
	return _SportRegistry.contract.Transact(opts, "recordSport", id, recordHash)
}

// RecordSport is a paid mutator transaction binding the contract method 0xd76c9386.
//
// Solidity: function recordSport(uint256 id, bytes32 recordHash) returns()
func (_SportRegistry *SportRegistrySession) RecordSport(id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
	return _SportRegistry.Contract.RecordSport(&_SportRegistry.TransactOpts, id, recordHash)
}

// RecordSport is a paid mutator transaction binding the contract method 0xd76c9386.
//
// Solidity: function recordSport(uint256 id, bytes32 recordHash) returns()
func (_SportRegistry *SportRegistryTransactorSession) RecordSport(id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
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
	RecordHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSportRecorded is a free log retrieval operation binding the contract event 0x3a90657270ecf8ae1de2f87142a28a59f2ff9d42307e84a7e388ccab1413610b.
//
// Solidity: event SportRecorded(uint256 indexed id, bytes32 recordHash)
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

// WatchSportRecorded is a free log subscription operation binding the contract event 0x3a90657270ecf8ae1de2f87142a28a59f2ff9d42307e84a7e388ccab1413610b.
//
// Solidity: event SportRecorded(uint256 indexed id, bytes32 recordHash)
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

// ParseSportRecorded is a log parse operation binding the contract event 0x3a90657270ecf8ae1de2f87142a28a59f2ff9d42307e84a7e388ccab1413610b.
//
// Solidity: event SportRecorded(uint256 indexed id, bytes32 recordHash)
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"TeamAchievementsRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"recordTeamAchievements\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recorded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"validateTeamAchievements\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061023e8061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c8063501895ae1461004e57806360b2267b1461008057806367f7f8e9146100b257806390601114146100c5575b5f5ffd5b61006d61005c3660046101d1565b60016020525f908152604090205481565b6040519081526020015b60405180910390f35b6100a261008e3660046101d1565b5f6020819052908152604090205460ff1681565b6040519015158152602001610077565b6100a26100c03660046101e8565b6100da565b6100d86100d33660046101e8565b61010a565b005b5f8281526020819052604081205460ff16801561010357505f8381526001602052604090205482145b9392505050565b5f8281526020819052604090205460ff16156101765760405162461bcd60e51b815260206004820152602160248201527f5465616d416368696576656d656e747320616c7265616479207265636f7264656044820152601960fa1b606482015260840160405180910390fd5b5f82815260208181526040808320805460ff19166001908117909155825291829020839055905182815283917f4ccb76ae22e046343f35e06efd22087741eef767301af42bd138cae899df4899910160405180910390a25050565b5f602082840312156101e1575f5ffd5b5035919050565b5f5f604083850312156101f9575f5ffd5b5050803592602090910135915056fea264697066735822122001f2730e409ea4cd63cd9c1a6b0c4f7998f59010eef90d5aa3a926cc4c53451b64736f6c634300081d0033",
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
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_TeamAchievementsRegistry *TeamAchievementsRegistryCaller) Hashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TeamAchievementsRegistry.contract.Call(opts, &out, "hashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_TeamAchievementsRegistry *TeamAchievementsRegistrySession) Hashes(arg0 *big.Int) ([32]byte, error) {
	return _TeamAchievementsRegistry.Contract.Hashes(&_TeamAchievementsRegistry.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_TeamAchievementsRegistry *TeamAchievementsRegistryCallerSession) Hashes(arg0 *big.Int) ([32]byte, error) {
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

// ValidateTeamAchievements is a free data retrieval call binding the contract method 0x67f7f8e9.
//
// Solidity: function validateTeamAchievements(uint256 id, bytes32 recordHash) view returns(bool)
func (_TeamAchievementsRegistry *TeamAchievementsRegistryCaller) ValidateTeamAchievements(opts *bind.CallOpts, id *big.Int, recordHash [32]byte) (bool, error) {
	var out []interface{}
	err := _TeamAchievementsRegistry.contract.Call(opts, &out, "validateTeamAchievements", id, recordHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateTeamAchievements is a free data retrieval call binding the contract method 0x67f7f8e9.
//
// Solidity: function validateTeamAchievements(uint256 id, bytes32 recordHash) view returns(bool)
func (_TeamAchievementsRegistry *TeamAchievementsRegistrySession) ValidateTeamAchievements(id *big.Int, recordHash [32]byte) (bool, error) {
	return _TeamAchievementsRegistry.Contract.ValidateTeamAchievements(&_TeamAchievementsRegistry.CallOpts, id, recordHash)
}

// ValidateTeamAchievements is a free data retrieval call binding the contract method 0x67f7f8e9.
//
// Solidity: function validateTeamAchievements(uint256 id, bytes32 recordHash) view returns(bool)
func (_TeamAchievementsRegistry *TeamAchievementsRegistryCallerSession) ValidateTeamAchievements(id *big.Int, recordHash [32]byte) (bool, error) {
	return _TeamAchievementsRegistry.Contract.ValidateTeamAchievements(&_TeamAchievementsRegistry.CallOpts, id, recordHash)
}

// RecordTeamAchievements is a paid mutator transaction binding the contract method 0x90601114.
//
// Solidity: function recordTeamAchievements(uint256 id, bytes32 recordHash) returns()
func (_TeamAchievementsRegistry *TeamAchievementsRegistryTransactor) RecordTeamAchievements(opts *bind.TransactOpts, id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
	return _TeamAchievementsRegistry.contract.Transact(opts, "recordTeamAchievements", id, recordHash)
}

// RecordTeamAchievements is a paid mutator transaction binding the contract method 0x90601114.
//
// Solidity: function recordTeamAchievements(uint256 id, bytes32 recordHash) returns()
func (_TeamAchievementsRegistry *TeamAchievementsRegistrySession) RecordTeamAchievements(id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
	return _TeamAchievementsRegistry.Contract.RecordTeamAchievements(&_TeamAchievementsRegistry.TransactOpts, id, recordHash)
}

// RecordTeamAchievements is a paid mutator transaction binding the contract method 0x90601114.
//
// Solidity: function recordTeamAchievements(uint256 id, bytes32 recordHash) returns()
func (_TeamAchievementsRegistry *TeamAchievementsRegistryTransactorSession) RecordTeamAchievements(id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
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
	RecordHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTeamAchievementsRecorded is a free log retrieval operation binding the contract event 0x4ccb76ae22e046343f35e06efd22087741eef767301af42bd138cae899df4899.
//
// Solidity: event TeamAchievementsRecorded(uint256 indexed id, bytes32 recordHash)
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

// WatchTeamAchievementsRecorded is a free log subscription operation binding the contract event 0x4ccb76ae22e046343f35e06efd22087741eef767301af42bd138cae899df4899.
//
// Solidity: event TeamAchievementsRecorded(uint256 indexed id, bytes32 recordHash)
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

// ParseTeamAchievementsRecorded is a log parse operation binding the contract event 0x4ccb76ae22e046343f35e06efd22087741eef767301af42bd138cae899df4899.
//
// Solidity: event TeamAchievementsRecorded(uint256 indexed id, bytes32 recordHash)
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"TeamPersonRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"recordTeamPerson\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recorded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"validateTeamPerson\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506102348061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c8063501895ae1461004e57806360b2267b14610080578063c8461d7d146100b2578063ce44aa83146100c5575b5f5ffd5b61006d61005c3660046101c7565b60016020525f908152604090205481565b6040519081526020015b60405180910390f35b6100a261008e3660046101c7565b5f6020819052908152604090205460ff1681565b6040519015158152602001610077565b6100a26100c03660046101de565b6100da565b6100d86100d33660046101de565b61010a565b005b5f8281526020819052604081205460ff16801561010357505f8381526001602052604090205482145b9392505050565b5f8281526020819052604090205460ff161561016c5760405162461bcd60e51b815260206004820152601b60248201527f5465616d506572736f6e20616c7265616479207265636f726465640000000000604482015260640160405180910390fd5b5f82815260208181526040808320805460ff19166001908117909155825291829020839055905182815283917f19b871a1e9d52c1eeecd8e47c380c8aeabac9753b1e4dbaefab31819d47e13d8910160405180910390a25050565b5f602082840312156101d7575f5ffd5b5035919050565b5f5f604083850312156101ef575f5ffd5b5050803592602090910135915056fea2646970667358221220ddbccef7e90bfb87145cac3096979cc9894a1cb47e908fa630ab1ecaa2af7f3564736f6c634300081d0033",
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
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_TeamPersonRegistry *TeamPersonRegistryCaller) Hashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TeamPersonRegistry.contract.Call(opts, &out, "hashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_TeamPersonRegistry *TeamPersonRegistrySession) Hashes(arg0 *big.Int) ([32]byte, error) {
	return _TeamPersonRegistry.Contract.Hashes(&_TeamPersonRegistry.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_TeamPersonRegistry *TeamPersonRegistryCallerSession) Hashes(arg0 *big.Int) ([32]byte, error) {
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

// ValidateTeamPerson is a free data retrieval call binding the contract method 0xc8461d7d.
//
// Solidity: function validateTeamPerson(uint256 id, bytes32 recordHash) view returns(bool)
func (_TeamPersonRegistry *TeamPersonRegistryCaller) ValidateTeamPerson(opts *bind.CallOpts, id *big.Int, recordHash [32]byte) (bool, error) {
	var out []interface{}
	err := _TeamPersonRegistry.contract.Call(opts, &out, "validateTeamPerson", id, recordHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateTeamPerson is a free data retrieval call binding the contract method 0xc8461d7d.
//
// Solidity: function validateTeamPerson(uint256 id, bytes32 recordHash) view returns(bool)
func (_TeamPersonRegistry *TeamPersonRegistrySession) ValidateTeamPerson(id *big.Int, recordHash [32]byte) (bool, error) {
	return _TeamPersonRegistry.Contract.ValidateTeamPerson(&_TeamPersonRegistry.CallOpts, id, recordHash)
}

// ValidateTeamPerson is a free data retrieval call binding the contract method 0xc8461d7d.
//
// Solidity: function validateTeamPerson(uint256 id, bytes32 recordHash) view returns(bool)
func (_TeamPersonRegistry *TeamPersonRegistryCallerSession) ValidateTeamPerson(id *big.Int, recordHash [32]byte) (bool, error) {
	return _TeamPersonRegistry.Contract.ValidateTeamPerson(&_TeamPersonRegistry.CallOpts, id, recordHash)
}

// RecordTeamPerson is a paid mutator transaction binding the contract method 0xce44aa83.
//
// Solidity: function recordTeamPerson(uint256 id, bytes32 recordHash) returns()
func (_TeamPersonRegistry *TeamPersonRegistryTransactor) RecordTeamPerson(opts *bind.TransactOpts, id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
	return _TeamPersonRegistry.contract.Transact(opts, "recordTeamPerson", id, recordHash)
}

// RecordTeamPerson is a paid mutator transaction binding the contract method 0xce44aa83.
//
// Solidity: function recordTeamPerson(uint256 id, bytes32 recordHash) returns()
func (_TeamPersonRegistry *TeamPersonRegistrySession) RecordTeamPerson(id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
	return _TeamPersonRegistry.Contract.RecordTeamPerson(&_TeamPersonRegistry.TransactOpts, id, recordHash)
}

// RecordTeamPerson is a paid mutator transaction binding the contract method 0xce44aa83.
//
// Solidity: function recordTeamPerson(uint256 id, bytes32 recordHash) returns()
func (_TeamPersonRegistry *TeamPersonRegistryTransactorSession) RecordTeamPerson(id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
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
	RecordHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTeamPersonRecorded is a free log retrieval operation binding the contract event 0x19b871a1e9d52c1eeecd8e47c380c8aeabac9753b1e4dbaefab31819d47e13d8.
//
// Solidity: event TeamPersonRecorded(uint256 indexed id, bytes32 recordHash)
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

// WatchTeamPersonRecorded is a free log subscription operation binding the contract event 0x19b871a1e9d52c1eeecd8e47c380c8aeabac9753b1e4dbaefab31819d47e13d8.
//
// Solidity: event TeamPersonRecorded(uint256 indexed id, bytes32 recordHash)
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

// ParseTeamPersonRecorded is a log parse operation binding the contract event 0x19b871a1e9d52c1eeecd8e47c380c8aeabac9753b1e4dbaefab31819d47e13d8.
//
// Solidity: event TeamPersonRecorded(uint256 indexed id, bytes32 recordHash)
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"TeamRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"recordTeam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recorded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"validateTeam\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061022c8061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c806317c22dcd1461004e578063501895ae1461007657806360b2267b146100a35780637e55300b146100c5575b5f5ffd5b61006161005c3660046101bf565b6100da565b60405190151581526020015b60405180910390f35b6100956100843660046101df565b60016020525f908152604090205481565b60405190815260200161006d565b6100616100b13660046101df565b5f6020819052908152604090205460ff1681565b6100d86100d33660046101bf565b61010a565b005b5f8281526020819052604081205460ff16801561010357505f8381526001602052604090205482145b9392505050565b5f8281526020819052604090205460ff16156101645760405162461bcd60e51b81526020600482015260156024820152741519585b48185b1c9958591e481c9958dbdc991959605a1b604482015260640160405180910390fd5b5f82815260208181526040808320805460ff19166001908117909155825291829020839055905182815283917fc7c2936543b717eb7c516edee4f9e4f4055071ab079302b387ee8c28e512c0ae910160405180910390a25050565b5f5f604083850312156101d0575f5ffd5b50508035926020909101359150565b5f602082840312156101ef575f5ffd5b503591905056fea26469706673582212203a750cb892017c889818950bb824efd7023d69e1e40ca571f3d3966160f14adc64736f6c634300081d0033",
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
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_TeamRegistry *TeamRegistryCaller) Hashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TeamRegistry.contract.Call(opts, &out, "hashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_TeamRegistry *TeamRegistrySession) Hashes(arg0 *big.Int) ([32]byte, error) {
	return _TeamRegistry.Contract.Hashes(&_TeamRegistry.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_TeamRegistry *TeamRegistryCallerSession) Hashes(arg0 *big.Int) ([32]byte, error) {
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

// ValidateTeam is a free data retrieval call binding the contract method 0x17c22dcd.
//
// Solidity: function validateTeam(uint256 id, bytes32 recordHash) view returns(bool)
func (_TeamRegistry *TeamRegistryCaller) ValidateTeam(opts *bind.CallOpts, id *big.Int, recordHash [32]byte) (bool, error) {
	var out []interface{}
	err := _TeamRegistry.contract.Call(opts, &out, "validateTeam", id, recordHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateTeam is a free data retrieval call binding the contract method 0x17c22dcd.
//
// Solidity: function validateTeam(uint256 id, bytes32 recordHash) view returns(bool)
func (_TeamRegistry *TeamRegistrySession) ValidateTeam(id *big.Int, recordHash [32]byte) (bool, error) {
	return _TeamRegistry.Contract.ValidateTeam(&_TeamRegistry.CallOpts, id, recordHash)
}

// ValidateTeam is a free data retrieval call binding the contract method 0x17c22dcd.
//
// Solidity: function validateTeam(uint256 id, bytes32 recordHash) view returns(bool)
func (_TeamRegistry *TeamRegistryCallerSession) ValidateTeam(id *big.Int, recordHash [32]byte) (bool, error) {
	return _TeamRegistry.Contract.ValidateTeam(&_TeamRegistry.CallOpts, id, recordHash)
}

// RecordTeam is a paid mutator transaction binding the contract method 0x7e55300b.
//
// Solidity: function recordTeam(uint256 id, bytes32 recordHash) returns()
func (_TeamRegistry *TeamRegistryTransactor) RecordTeam(opts *bind.TransactOpts, id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
	return _TeamRegistry.contract.Transact(opts, "recordTeam", id, recordHash)
}

// RecordTeam is a paid mutator transaction binding the contract method 0x7e55300b.
//
// Solidity: function recordTeam(uint256 id, bytes32 recordHash) returns()
func (_TeamRegistry *TeamRegistrySession) RecordTeam(id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
	return _TeamRegistry.Contract.RecordTeam(&_TeamRegistry.TransactOpts, id, recordHash)
}

// RecordTeam is a paid mutator transaction binding the contract method 0x7e55300b.
//
// Solidity: function recordTeam(uint256 id, bytes32 recordHash) returns()
func (_TeamRegistry *TeamRegistryTransactorSession) RecordTeam(id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
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
	RecordHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTeamRecorded is a free log retrieval operation binding the contract event 0xc7c2936543b717eb7c516edee4f9e4f4055071ab079302b387ee8c28e512c0ae.
//
// Solidity: event TeamRecorded(uint256 indexed id, bytes32 recordHash)
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

// WatchTeamRecorded is a free log subscription operation binding the contract event 0xc7c2936543b717eb7c516edee4f9e4f4055071ab079302b387ee8c28e512c0ae.
//
// Solidity: event TeamRecorded(uint256 indexed id, bytes32 recordHash)
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

// ParseTeamRecorded is a log parse operation binding the contract event 0xc7c2936543b717eb7c516edee4f9e4f4055071ab079302b387ee8c28e512c0ae.
//
// Solidity: event TeamRecorded(uint256 indexed id, bytes32 recordHash)
func (_TeamRegistry *TeamRegistryFilterer) ParseTeamRecorded(log types.Log) (*TeamRegistryTeamRecorded, error) {
	event := new(TeamRegistryTeamRecorded)
	if err := _TeamRegistry.contract.UnpackLog(event, "TeamRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

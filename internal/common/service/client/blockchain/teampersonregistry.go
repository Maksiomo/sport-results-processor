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

// BlockchainMetaData contains all meta data concerning the Blockchain contract.
var BlockchainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"TeamPersonRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"recordTeamPerson\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recorded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"validateTeamPerson\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506102348061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c8063501895ae1461004e57806360b2267b14610080578063c8461d7d146100b2578063ce44aa83146100c5575b5f5ffd5b61006d61005c3660046101c7565b60016020525f908152604090205481565b6040519081526020015b60405180910390f35b6100a261008e3660046101c7565b5f6020819052908152604090205460ff1681565b6040519015158152602001610077565b6100a26100c03660046101de565b6100da565b6100d86100d33660046101de565b61010a565b005b5f8281526020819052604081205460ff16801561010357505f8381526001602052604090205482145b9392505050565b5f8281526020819052604090205460ff161561016c5760405162461bcd60e51b815260206004820152601b60248201527f5465616d506572736f6e20616c7265616479207265636f726465640000000000604482015260640160405180910390fd5b5f82815260208181526040808320805460ff19166001908117909155825291829020839055905182815283917f19b871a1e9d52c1eeecd8e47c380c8aeabac9753b1e4dbaefab31819d47e13d8910160405180910390a25050565b5f602082840312156101d7575f5ffd5b5035919050565b5f5f604083850312156101ef575f5ffd5b5050803592602090910135915056fea2646970667358221220955e8412fd72e872fb7809d81ca0bce0bdac6bb7a0d8c26aac78e3ae640a61de64736f6c634300081d0033",
}

// BlockchainABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockchainMetaData.ABI instead.
var BlockchainABI = BlockchainMetaData.ABI

// BlockchainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BlockchainMetaData.Bin instead.
var BlockchainBin = BlockchainMetaData.Bin

// DeployBlockchain deploys a new Ethereum contract, binding an instance of Blockchain to it.
func DeployBlockchain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Blockchain, error) {
	parsed, err := BlockchainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BlockchainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Blockchain{BlockchainCaller: BlockchainCaller{contract: contract}, BlockchainTransactor: BlockchainTransactor{contract: contract}, BlockchainFilterer: BlockchainFilterer{contract: contract}}, nil
}

// Blockchain is an auto generated Go binding around an Ethereum contract.
type Blockchain struct {
	BlockchainCaller     // Read-only binding to the contract
	BlockchainTransactor // Write-only binding to the contract
	BlockchainFilterer   // Log filterer for contract events
}

// BlockchainCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockchainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockchainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockchainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockchainSession struct {
	Contract     *Blockchain       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlockchainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockchainCallerSession struct {
	Contract *BlockchainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BlockchainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockchainTransactorSession struct {
	Contract     *BlockchainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BlockchainRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockchainRaw struct {
	Contract *Blockchain // Generic contract binding to access the raw methods on
}

// BlockchainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockchainCallerRaw struct {
	Contract *BlockchainCaller // Generic read-only contract binding to access the raw methods on
}

// BlockchainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockchainTransactorRaw struct {
	Contract *BlockchainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockchain creates a new instance of Blockchain, bound to a specific deployed contract.
func NewBlockchain(address common.Address, backend bind.ContractBackend) (*Blockchain, error) {
	contract, err := bindBlockchain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Blockchain{BlockchainCaller: BlockchainCaller{contract: contract}, BlockchainTransactor: BlockchainTransactor{contract: contract}, BlockchainFilterer: BlockchainFilterer{contract: contract}}, nil
}

// NewBlockchainCaller creates a new read-only instance of Blockchain, bound to a specific deployed contract.
func NewBlockchainCaller(address common.Address, caller bind.ContractCaller) (*BlockchainCaller, error) {
	contract, err := bindBlockchain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockchainCaller{contract: contract}, nil
}

// NewBlockchainTransactor creates a new write-only instance of Blockchain, bound to a specific deployed contract.
func NewBlockchainTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockchainTransactor, error) {
	contract, err := bindBlockchain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockchainTransactor{contract: contract}, nil
}

// NewBlockchainFilterer creates a new log filterer instance of Blockchain, bound to a specific deployed contract.
func NewBlockchainFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockchainFilterer, error) {
	contract, err := bindBlockchain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockchainFilterer{contract: contract}, nil
}

// bindBlockchain binds a generic wrapper to an already deployed contract.
func bindBlockchain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlockchainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockchain *BlockchainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blockchain.Contract.BlockchainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockchain *BlockchainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockchain.Contract.BlockchainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockchain *BlockchainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blockchain.Contract.BlockchainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockchain *BlockchainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blockchain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockchain *BlockchainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockchain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockchain *BlockchainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blockchain.Contract.contract.Transact(opts, method, params...)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_Blockchain *BlockchainCaller) Hashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "hashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_Blockchain *BlockchainSession) Hashes(arg0 *big.Int) ([32]byte, error) {
	return _Blockchain.Contract.Hashes(&_Blockchain.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0x501895ae.
//
// Solidity: function hashes(uint256 ) view returns(bytes32)
func (_Blockchain *BlockchainCallerSession) Hashes(arg0 *big.Int) ([32]byte, error) {
	return _Blockchain.Contract.Hashes(&_Blockchain.CallOpts, arg0)
}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_Blockchain *BlockchainCaller) Recorded(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "recorded", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_Blockchain *BlockchainSession) Recorded(arg0 *big.Int) (bool, error) {
	return _Blockchain.Contract.Recorded(&_Blockchain.CallOpts, arg0)
}

// Recorded is a free data retrieval call binding the contract method 0x60b2267b.
//
// Solidity: function recorded(uint256 ) view returns(bool)
func (_Blockchain *BlockchainCallerSession) Recorded(arg0 *big.Int) (bool, error) {
	return _Blockchain.Contract.Recorded(&_Blockchain.CallOpts, arg0)
}

// ValidateTeamPerson is a free data retrieval call binding the contract method 0xc8461d7d.
//
// Solidity: function validateTeamPerson(uint256 id, bytes32 recordHash) view returns(bool)
func (_Blockchain *BlockchainCaller) ValidateTeamPerson(opts *bind.CallOpts, id *big.Int, recordHash [32]byte) (bool, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "validateTeamPerson", id, recordHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateTeamPerson is a free data retrieval call binding the contract method 0xc8461d7d.
//
// Solidity: function validateTeamPerson(uint256 id, bytes32 recordHash) view returns(bool)
func (_Blockchain *BlockchainSession) ValidateTeamPerson(id *big.Int, recordHash [32]byte) (bool, error) {
	return _Blockchain.Contract.ValidateTeamPerson(&_Blockchain.CallOpts, id, recordHash)
}

// ValidateTeamPerson is a free data retrieval call binding the contract method 0xc8461d7d.
//
// Solidity: function validateTeamPerson(uint256 id, bytes32 recordHash) view returns(bool)
func (_Blockchain *BlockchainCallerSession) ValidateTeamPerson(id *big.Int, recordHash [32]byte) (bool, error) {
	return _Blockchain.Contract.ValidateTeamPerson(&_Blockchain.CallOpts, id, recordHash)
}

// RecordTeamPerson is a paid mutator transaction binding the contract method 0xce44aa83.
//
// Solidity: function recordTeamPerson(uint256 id, bytes32 recordHash) returns()
func (_Blockchain *BlockchainTransactor) RecordTeamPerson(opts *bind.TransactOpts, id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "recordTeamPerson", id, recordHash)
}

// RecordTeamPerson is a paid mutator transaction binding the contract method 0xce44aa83.
//
// Solidity: function recordTeamPerson(uint256 id, bytes32 recordHash) returns()
func (_Blockchain *BlockchainSession) RecordTeamPerson(id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
	return _Blockchain.Contract.RecordTeamPerson(&_Blockchain.TransactOpts, id, recordHash)
}

// RecordTeamPerson is a paid mutator transaction binding the contract method 0xce44aa83.
//
// Solidity: function recordTeamPerson(uint256 id, bytes32 recordHash) returns()
func (_Blockchain *BlockchainTransactorSession) RecordTeamPerson(id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
	return _Blockchain.Contract.RecordTeamPerson(&_Blockchain.TransactOpts, id, recordHash)
}

// BlockchainTeamPersonRecordedIterator is returned from FilterTeamPersonRecorded and is used to iterate over the raw logs and unpacked data for TeamPersonRecorded events raised by the Blockchain contract.
type BlockchainTeamPersonRecordedIterator struct {
	Event *BlockchainTeamPersonRecorded // Event containing the contract specifics and raw log

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
func (it *BlockchainTeamPersonRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainTeamPersonRecorded)
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
		it.Event = new(BlockchainTeamPersonRecorded)
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
func (it *BlockchainTeamPersonRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainTeamPersonRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainTeamPersonRecorded represents a TeamPersonRecorded event raised by the Blockchain contract.
type BlockchainTeamPersonRecorded struct {
	Id         *big.Int
	RecordHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTeamPersonRecorded is a free log retrieval operation binding the contract event 0x19b871a1e9d52c1eeecd8e47c380c8aeabac9753b1e4dbaefab31819d47e13d8.
//
// Solidity: event TeamPersonRecorded(uint256 indexed id, bytes32 recordHash)
func (_Blockchain *BlockchainFilterer) FilterTeamPersonRecorded(opts *bind.FilterOpts, id []*big.Int) (*BlockchainTeamPersonRecordedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "TeamPersonRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainTeamPersonRecordedIterator{contract: _Blockchain.contract, event: "TeamPersonRecorded", logs: logs, sub: sub}, nil
}

// WatchTeamPersonRecorded is a free log subscription operation binding the contract event 0x19b871a1e9d52c1eeecd8e47c380c8aeabac9753b1e4dbaefab31819d47e13d8.
//
// Solidity: event TeamPersonRecorded(uint256 indexed id, bytes32 recordHash)
func (_Blockchain *BlockchainFilterer) WatchTeamPersonRecorded(opts *bind.WatchOpts, sink chan<- *BlockchainTeamPersonRecorded, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "TeamPersonRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainTeamPersonRecorded)
				if err := _Blockchain.contract.UnpackLog(event, "TeamPersonRecorded", log); err != nil {
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
func (_Blockchain *BlockchainFilterer) ParseTeamPersonRecorded(log types.Log) (*BlockchainTeamPersonRecorded, error) {
	event := new(BlockchainTeamPersonRecorded)
	if err := _Blockchain.contract.UnpackLog(event, "TeamPersonRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

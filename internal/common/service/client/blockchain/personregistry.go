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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"PersonRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"recordPerson\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recorded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recordHash\",\"type\":\"bytes32\"}],\"name\":\"validatePerson\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506102348061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c8063501895ae1461004e57806360b2267b1461008057806372013409146100b2578063e600835d146100c7575b5f5ffd5b61006d61005c3660046101c7565b60016020525f908152604090205481565b6040519081526020015b60405180910390f35b6100a261008e3660046101c7565b5f6020819052908152604090205460ff1681565b6040519015158152602001610077565b6100c56100c03660046101de565b6100da565b005b6100a26100d53660046101de565b610197565b5f8281526020819052604090205460ff161561013c5760405162461bcd60e51b815260206004820152601760248201527f506572736f6e20616c7265616479207265636f72646564000000000000000000604482015260640160405180910390fd5b5f82815260208181526040808320805460ff19166001908117909155825291829020839055905182815283917fccc35dc586d377da368a5e0ae6527ddbd36d7520272350d11ddde3995c8e2ff9910160405180910390a25050565b5f8281526020819052604081205460ff1680156101c057505f8381526001602052604090205482145b9392505050565b5f602082840312156101d7575f5ffd5b5035919050565b5f5f604083850312156101ef575f5ffd5b5050803592602090910135915056fea2646970667358221220d4c1d990030454a8da8da2072bf5f0e967691a0a03460538c83e35dec8b4b15e64736f6c634300081d0033",
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

// ValidatePerson is a free data retrieval call binding the contract method 0xe600835d.
//
// Solidity: function validatePerson(uint256 id, bytes32 recordHash) view returns(bool)
func (_Blockchain *BlockchainCaller) ValidatePerson(opts *bind.CallOpts, id *big.Int, recordHash [32]byte) (bool, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "validatePerson", id, recordHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidatePerson is a free data retrieval call binding the contract method 0xe600835d.
//
// Solidity: function validatePerson(uint256 id, bytes32 recordHash) view returns(bool)
func (_Blockchain *BlockchainSession) ValidatePerson(id *big.Int, recordHash [32]byte) (bool, error) {
	return _Blockchain.Contract.ValidatePerson(&_Blockchain.CallOpts, id, recordHash)
}

// ValidatePerson is a free data retrieval call binding the contract method 0xe600835d.
//
// Solidity: function validatePerson(uint256 id, bytes32 recordHash) view returns(bool)
func (_Blockchain *BlockchainCallerSession) ValidatePerson(id *big.Int, recordHash [32]byte) (bool, error) {
	return _Blockchain.Contract.ValidatePerson(&_Blockchain.CallOpts, id, recordHash)
}

// RecordPerson is a paid mutator transaction binding the contract method 0x72013409.
//
// Solidity: function recordPerson(uint256 id, bytes32 recordHash) returns()
func (_Blockchain *BlockchainTransactor) RecordPerson(opts *bind.TransactOpts, id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "recordPerson", id, recordHash)
}

// RecordPerson is a paid mutator transaction binding the contract method 0x72013409.
//
// Solidity: function recordPerson(uint256 id, bytes32 recordHash) returns()
func (_Blockchain *BlockchainSession) RecordPerson(id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
	return _Blockchain.Contract.RecordPerson(&_Blockchain.TransactOpts, id, recordHash)
}

// RecordPerson is a paid mutator transaction binding the contract method 0x72013409.
//
// Solidity: function recordPerson(uint256 id, bytes32 recordHash) returns()
func (_Blockchain *BlockchainTransactorSession) RecordPerson(id *big.Int, recordHash [32]byte) (*types.Transaction, error) {
	return _Blockchain.Contract.RecordPerson(&_Blockchain.TransactOpts, id, recordHash)
}

// BlockchainPersonRecordedIterator is returned from FilterPersonRecorded and is used to iterate over the raw logs and unpacked data for PersonRecorded events raised by the Blockchain contract.
type BlockchainPersonRecordedIterator struct {
	Event *BlockchainPersonRecorded // Event containing the contract specifics and raw log

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
func (it *BlockchainPersonRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainPersonRecorded)
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
		it.Event = new(BlockchainPersonRecorded)
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
func (it *BlockchainPersonRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainPersonRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainPersonRecorded represents a PersonRecorded event raised by the Blockchain contract.
type BlockchainPersonRecorded struct {
	Id         *big.Int
	RecordHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPersonRecorded is a free log retrieval operation binding the contract event 0xccc35dc586d377da368a5e0ae6527ddbd36d7520272350d11ddde3995c8e2ff9.
//
// Solidity: event PersonRecorded(uint256 indexed id, bytes32 recordHash)
func (_Blockchain *BlockchainFilterer) FilterPersonRecorded(opts *bind.FilterOpts, id []*big.Int) (*BlockchainPersonRecordedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "PersonRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainPersonRecordedIterator{contract: _Blockchain.contract, event: "PersonRecorded", logs: logs, sub: sub}, nil
}

// WatchPersonRecorded is a free log subscription operation binding the contract event 0xccc35dc586d377da368a5e0ae6527ddbd36d7520272350d11ddde3995c8e2ff9.
//
// Solidity: event PersonRecorded(uint256 indexed id, bytes32 recordHash)
func (_Blockchain *BlockchainFilterer) WatchPersonRecorded(opts *bind.WatchOpts, sink chan<- *BlockchainPersonRecorded, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "PersonRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainPersonRecorded)
				if err := _Blockchain.contract.UnpackLog(event, "PersonRecorded", log); err != nil {
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
func (_Blockchain *BlockchainFilterer) ParsePersonRecorded(log types.Log) (*BlockchainPersonRecorded, error) {
	event := new(BlockchainPersonRecorded)
	if err := _Blockchain.contract.UnpackLog(event, "PersonRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

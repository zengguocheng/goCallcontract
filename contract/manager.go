// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package manager

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// ManagerABI is the input ABI used to generate the binding from.
const ManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"submitAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"contractUser\",\"outputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"}],\"name\":\"SubmitContract\",\"type\":\"event\"}]"

// ManagerBin is the compiled bytecode used for deploying new contracts.
const ManagerBin = `608060405234801561001057600080fd5b50610326806100206000396000f30060806040526004361061004c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063ef97ba3914610051578063f121f45514610094575b600080fd5b34801561005d57600080fd5b50610092600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061011e565b005b3480156100a057600080fd5b506100d5600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506102b6565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b6000808273ffffffffffffffffffffffffffffffffffffffff161415151561014557600080fd5b813b905060008163ffffffff1611151561015e57600080fd5b60408051908101604052803373ffffffffffffffffffffffffffffffffffffffff168152602001428152506000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101559050507f54755321060261a6fba43c679c62e5d67bd9cadc6df62c627b6d0364340497c18233604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a15050565b60006020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600101549050825600a165627a7a72305820a765dd3538fd0247af6c01204a3ca997cdc01675a9d0a85ca0f7b4fa6963703a0029`

// DeployManager deploys a new Ethereum contract, binding an instance of Manager to it.
func DeployManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Manager, error) {
	parsed, err := abi.JSON(strings.NewReader(ManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Manager{ManagerCaller: ManagerCaller{contract: contract}, ManagerTransactor: ManagerTransactor{contract: contract}, ManagerFilterer: ManagerFilterer{contract: contract}}, nil
}

// Manager is an auto generated Go binding around an Ethereum contract.
type Manager struct {
	ManagerCaller     // Read-only binding to the contract
	ManagerTransactor // Write-only binding to the contract
	ManagerFilterer   // Log filterer for contract events
}

// ManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ManagerSession struct {
	Contract     *Manager          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ManagerCallerSession struct {
	Contract *ManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ManagerTransactorSession struct {
	Contract     *ManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ManagerRaw struct {
	Contract *Manager // Generic contract binding to access the raw methods on
}

// ManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ManagerCallerRaw struct {
	Contract *ManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ManagerTransactorRaw struct {
	Contract *ManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewManager creates a new instance of Manager, bound to a specific deployed contract.
func NewManager(address common.Address, backend bind.ContractBackend) (*Manager, error) {
	contract, err := bindManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Manager{ManagerCaller: ManagerCaller{contract: contract}, ManagerTransactor: ManagerTransactor{contract: contract}, ManagerFilterer: ManagerFilterer{contract: contract}}, nil
}

// NewManagerCaller creates a new read-only instance of Manager, bound to a specific deployed contract.
func NewManagerCaller(address common.Address, caller bind.ContractCaller) (*ManagerCaller, error) {
	contract, err := bindManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ManagerCaller{contract: contract}, nil
}

// NewManagerTransactor creates a new write-only instance of Manager, bound to a specific deployed contract.
func NewManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ManagerTransactor, error) {
	contract, err := bindManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ManagerTransactor{contract: contract}, nil
}

// NewManagerFilterer creates a new log filterer instance of Manager, bound to a specific deployed contract.
func NewManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ManagerFilterer, error) {
	contract, err := bindManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ManagerFilterer{contract: contract}, nil
}

// bindManager binds a generic wrapper to an already deployed contract.
func bindManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Manager *ManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Manager.Contract.ManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Manager *ManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manager.Contract.ManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Manager *ManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Manager.Contract.ManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Manager *ManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Manager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Manager *ManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Manager *ManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Manager.Contract.contract.Transact(opts, method, params...)
}

// ContractUser is a free data retrieval call binding the contract method 0xf121f455.
//
// Solidity: function contractUser( address) constant returns(from address, timestamp uint256)
func (_Manager *ManagerCaller) ContractUser(opts *bind.CallOpts, arg0 common.Address) (struct {
	From      common.Address
	Timestamp *big.Int
}, error) {
	ret := new(struct {
		From      common.Address
		Timestamp *big.Int
	})
	out := ret
	err := _Manager.contract.Call(opts, out, "contractUser", arg0)
	return *ret, err
}

// ContractUser is a free data retrieval call binding the contract method 0xf121f455.
//
// Solidity: function contractUser( address) constant returns(from address, timestamp uint256)
func (_Manager *ManagerSession) ContractUser(arg0 common.Address) (struct {
	From      common.Address
	Timestamp *big.Int
}, error) {
	return _Manager.Contract.ContractUser(&_Manager.CallOpts, arg0)
}

// ContractUser is a free data retrieval call binding the contract method 0xf121f455.
//
// Solidity: function contractUser( address) constant returns(from address, timestamp uint256)
func (_Manager *ManagerCallerSession) ContractUser(arg0 common.Address) (struct {
	From      common.Address
	Timestamp *big.Int
}, error) {
	return _Manager.Contract.ContractUser(&_Manager.CallOpts, arg0)
}

// SubmitAddress is a paid mutator transaction binding the contract method 0xef97ba39.
//
// Solidity: function submitAddress(_addr address) returns()
func (_Manager *ManagerTransactor) SubmitAddress(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "submitAddress", _addr)
}

// SubmitAddress is a paid mutator transaction binding the contract method 0xef97ba39.
//
// Solidity: function submitAddress(_addr address) returns()
func (_Manager *ManagerSession) SubmitAddress(_addr common.Address) (*types.Transaction, error) {
	return _Manager.Contract.SubmitAddress(&_Manager.TransactOpts, _addr)
}

// SubmitAddress is a paid mutator transaction binding the contract method 0xef97ba39.
//
// Solidity: function submitAddress(_addr address) returns()
func (_Manager *ManagerTransactorSession) SubmitAddress(_addr common.Address) (*types.Transaction, error) {
	return _Manager.Contract.SubmitAddress(&_Manager.TransactOpts, _addr)
}

// ManagerSubmitContractIterator is returned from FilterSubmitContract and is used to iterate over the raw logs and unpacked data for SubmitContract events raised by the Manager contract.
type ManagerSubmitContractIterator struct {
	Event *ManagerSubmitContract // Event containing the contract specifics and raw log

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
func (it *ManagerSubmitContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerSubmitContract)
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
		it.Event = new(ManagerSubmitContract)
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
func (it *ManagerSubmitContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerSubmitContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerSubmitContract represents a SubmitContract event raised by the Manager contract.
type ManagerSubmitContract struct {
	Addr common.Address
	From common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSubmitContract is a free log retrieval operation binding the contract event 0x54755321060261a6fba43c679c62e5d67bd9cadc6df62c627b6d0364340497c1.
//
// Solidity: e SubmitContract(addr address, from address)
func (_Manager *ManagerFilterer) FilterSubmitContract(opts *bind.FilterOpts) (*ManagerSubmitContractIterator, error) {

	logs, sub, err := _Manager.contract.FilterLogs(opts, "SubmitContract")
	if err != nil {
		return nil, err
	}
	return &ManagerSubmitContractIterator{contract: _Manager.contract, event: "SubmitContract", logs: logs, sub: sub}, nil
}

// WatchSubmitContract is a free log subscription operation binding the contract event 0x54755321060261a6fba43c679c62e5d67bd9cadc6df62c627b6d0364340497c1.
//
// Solidity: e SubmitContract(addr address, from address)
func (_Manager *ManagerFilterer) WatchSubmitContract(opts *bind.WatchOpts, sink chan<- *ManagerSubmitContract) (event.Subscription, error) {

	logs, sub, err := _Manager.contract.WatchLogs(opts, "SubmitContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerSubmitContract)
				if err := _Manager.contract.UnpackLog(event, "SubmitContract", log); err != nil {
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

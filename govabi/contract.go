// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package govabi

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

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"ProposalCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"ProposalExecutionExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"ProposalRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"ProposalResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"TasksErased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startIdx\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endIdx\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"handled\",\"type\":\"uint256\"}],\"name\":\"TasksHandled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegatedTo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"VoteCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"diff\",\"type\":\"uint256\"}],\"name\":\"VoteWeightOverridden\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"diff\",\"type\":\"uint256\"}],\"name\":\"VoteWeightUnOverridden\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegatedTo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"choices\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_bytes32\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastProposalID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxExecutionPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxOptions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minVotesRatio\",\"type\":\"uint256\"}],\"name\":\"minVotesAbsolute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"overriddenWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proposalBurntFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proposalFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"taskErasingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"taskHandlingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governableContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proposalVerifier\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"proposalParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pType\",\"type\":\"uint256\"},{\"internalType\":\"enumProposal.ExecType\",\"name\":\"executable\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"minVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAgreement\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"opinionScales\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"options\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"proposalContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"votingStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingMinEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingMaxEndTime\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionID\",\"type\":\"uint256\"}],\"name\":\"proposalOptionState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"agreementRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"agreement\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"proposalState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"winnerOptionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatedTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"getVote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"choices\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tasksCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getTask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"assignment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatedTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"choices\",\"type\":\"uint256[]\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposalContract\",\"type\":\"address\"}],\"name\":\"createProposal\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"cancelProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIdx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"handleTasks\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"tasksCleanup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"calculateVotingTally\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"proposalResolved\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"winnerID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatedTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"cancelVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"voterAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatedTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"name\":\"recountVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_Contract *ContractCaller) Bytes32ToString(opts *bind.CallOpts, _bytes32 [32]byte) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "bytes32ToString", _bytes32)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_Contract *ContractSession) Bytes32ToString(_bytes32 [32]byte) (string, error) {
	return _Contract.Contract.Bytes32ToString(&_Contract.CallOpts, _bytes32)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_Contract *ContractCallerSession) Bytes32ToString(_bytes32 [32]byte) (string, error) {
	return _Contract.Contract.Bytes32ToString(&_Contract.CallOpts, _bytes32)
}

// CalculateVotingTally is a free data retrieval call binding the contract method 0x14262d79.
//
// Solidity: function calculateVotingTally(uint256 proposalID) view returns(bool proposalResolved, uint256 winnerID, uint256 votes)
func (_Contract *ContractCaller) CalculateVotingTally(opts *bind.CallOpts, proposalID *big.Int) (struct {
	ProposalResolved bool
	WinnerID         *big.Int
	Votes            *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "calculateVotingTally", proposalID)

	outstruct := new(struct {
		ProposalResolved bool
		WinnerID         *big.Int
		Votes            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProposalResolved = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.WinnerID = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Votes = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalculateVotingTally is a free data retrieval call binding the contract method 0x14262d79.
//
// Solidity: function calculateVotingTally(uint256 proposalID) view returns(bool proposalResolved, uint256 winnerID, uint256 votes)
func (_Contract *ContractSession) CalculateVotingTally(proposalID *big.Int) (struct {
	ProposalResolved bool
	WinnerID         *big.Int
	Votes            *big.Int
}, error) {
	return _Contract.Contract.CalculateVotingTally(&_Contract.CallOpts, proposalID)
}

// CalculateVotingTally is a free data retrieval call binding the contract method 0x14262d79.
//
// Solidity: function calculateVotingTally(uint256 proposalID) view returns(bool proposalResolved, uint256 winnerID, uint256 votes)
func (_Contract *ContractCallerSession) CalculateVotingTally(proposalID *big.Int) (struct {
	ProposalResolved bool
	WinnerID         *big.Int
	Votes            *big.Int
}, error) {
	return _Contract.Contract.CalculateVotingTally(&_Contract.CallOpts, proposalID)
}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 i) view returns(bool active, uint256 assignment, uint256 proposalID)
func (_Contract *ContractCaller) GetTask(opts *bind.CallOpts, i *big.Int) (struct {
	Active     bool
	Assignment *big.Int
	ProposalID *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getTask", i)

	outstruct := new(struct {
		Active     bool
		Assignment *big.Int
		ProposalID *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Active = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Assignment = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ProposalID = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 i) view returns(bool active, uint256 assignment, uint256 proposalID)
func (_Contract *ContractSession) GetTask(i *big.Int) (struct {
	Active     bool
	Assignment *big.Int
	ProposalID *big.Int
}, error) {
	return _Contract.Contract.GetTask(&_Contract.CallOpts, i)
}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 i) view returns(bool active, uint256 assignment, uint256 proposalID)
func (_Contract *ContractCallerSession) GetTask(i *big.Int) (struct {
	Active     bool
	Assignment *big.Int
	ProposalID *big.Int
}, error) {
	return _Contract.Contract.GetTask(&_Contract.CallOpts, i)
}

// GetVote is a free data retrieval call binding the contract method 0xb9e6842b.
//
// Solidity: function getVote(address from, address delegatedTo, uint256 proposalID) view returns(uint256 weight, uint256[] choices)
func (_Contract *ContractCaller) GetVote(opts *bind.CallOpts, from common.Address, delegatedTo common.Address, proposalID *big.Int) (struct {
	Weight  *big.Int
	Choices []*big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getVote", from, delegatedTo, proposalID)

	outstruct := new(struct {
		Weight  *big.Int
		Choices []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Weight = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Choices = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetVote is a free data retrieval call binding the contract method 0xb9e6842b.
//
// Solidity: function getVote(address from, address delegatedTo, uint256 proposalID) view returns(uint256 weight, uint256[] choices)
func (_Contract *ContractSession) GetVote(from common.Address, delegatedTo common.Address, proposalID *big.Int) (struct {
	Weight  *big.Int
	Choices []*big.Int
}, error) {
	return _Contract.Contract.GetVote(&_Contract.CallOpts, from, delegatedTo, proposalID)
}

// GetVote is a free data retrieval call binding the contract method 0xb9e6842b.
//
// Solidity: function getVote(address from, address delegatedTo, uint256 proposalID) view returns(uint256 weight, uint256[] choices)
func (_Contract *ContractCallerSession) GetVote(from common.Address, delegatedTo common.Address, proposalID *big.Int) (struct {
	Weight  *big.Int
	Choices []*big.Int
}, error) {
	return _Contract.Contract.GetVote(&_Contract.CallOpts, from, delegatedTo, proposalID)
}

// LastProposalID is a free data retrieval call binding the contract method 0xa1d373d7.
//
// Solidity: function lastProposalID() view returns(uint256)
func (_Contract *ContractCaller) LastProposalID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "lastProposalID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastProposalID is a free data retrieval call binding the contract method 0xa1d373d7.
//
// Solidity: function lastProposalID() view returns(uint256)
func (_Contract *ContractSession) LastProposalID() (*big.Int, error) {
	return _Contract.Contract.LastProposalID(&_Contract.CallOpts)
}

// LastProposalID is a free data retrieval call binding the contract method 0xa1d373d7.
//
// Solidity: function lastProposalID() view returns(uint256)
func (_Contract *ContractCallerSession) LastProposalID() (*big.Int, error) {
	return _Contract.Contract.LastProposalID(&_Contract.CallOpts)
}

// MaxExecutionPeriod is a free data retrieval call binding the contract method 0xc25c2f26.
//
// Solidity: function maxExecutionPeriod() pure returns(uint256)
func (_Contract *ContractCaller) MaxExecutionPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxExecutionPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxExecutionPeriod is a free data retrieval call binding the contract method 0xc25c2f26.
//
// Solidity: function maxExecutionPeriod() pure returns(uint256)
func (_Contract *ContractSession) MaxExecutionPeriod() (*big.Int, error) {
	return _Contract.Contract.MaxExecutionPeriod(&_Contract.CallOpts)
}

// MaxExecutionPeriod is a free data retrieval call binding the contract method 0xc25c2f26.
//
// Solidity: function maxExecutionPeriod() pure returns(uint256)
func (_Contract *ContractCallerSession) MaxExecutionPeriod() (*big.Int, error) {
	return _Contract.Contract.MaxExecutionPeriod(&_Contract.CallOpts)
}

// MaxOptions is a free data retrieval call binding the contract method 0x5df17077.
//
// Solidity: function maxOptions() pure returns(uint256)
func (_Contract *ContractCaller) MaxOptions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxOptions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxOptions is a free data retrieval call binding the contract method 0x5df17077.
//
// Solidity: function maxOptions() pure returns(uint256)
func (_Contract *ContractSession) MaxOptions() (*big.Int, error) {
	return _Contract.Contract.MaxOptions(&_Contract.CallOpts)
}

// MaxOptions is a free data retrieval call binding the contract method 0x5df17077.
//
// Solidity: function maxOptions() pure returns(uint256)
func (_Contract *ContractCallerSession) MaxOptions() (*big.Int, error) {
	return _Contract.Contract.MaxOptions(&_Contract.CallOpts)
}

// MinVotesAbsolute is a free data retrieval call binding the contract method 0x1191e270.
//
// Solidity: function minVotesAbsolute(uint256 totalWeight, uint256 minVotesRatio) pure returns(uint256)
func (_Contract *ContractCaller) MinVotesAbsolute(opts *bind.CallOpts, totalWeight *big.Int, minVotesRatio *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minVotesAbsolute", totalWeight, minVotesRatio)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinVotesAbsolute is a free data retrieval call binding the contract method 0x1191e270.
//
// Solidity: function minVotesAbsolute(uint256 totalWeight, uint256 minVotesRatio) pure returns(uint256)
func (_Contract *ContractSession) MinVotesAbsolute(totalWeight *big.Int, minVotesRatio *big.Int) (*big.Int, error) {
	return _Contract.Contract.MinVotesAbsolute(&_Contract.CallOpts, totalWeight, minVotesRatio)
}

// MinVotesAbsolute is a free data retrieval call binding the contract method 0x1191e270.
//
// Solidity: function minVotesAbsolute(uint256 totalWeight, uint256 minVotesRatio) pure returns(uint256)
func (_Contract *ContractCallerSession) MinVotesAbsolute(totalWeight *big.Int, minVotesRatio *big.Int) (*big.Int, error) {
	return _Contract.Contract.MinVotesAbsolute(&_Contract.CallOpts, totalWeight, minVotesRatio)
}

// OverriddenWeight is a free data retrieval call binding the contract method 0x2177d6fc.
//
// Solidity: function overriddenWeight(address , uint256 ) view returns(uint256)
func (_Contract *ContractCaller) OverriddenWeight(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "overriddenWeight", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OverriddenWeight is a free data retrieval call binding the contract method 0x2177d6fc.
//
// Solidity: function overriddenWeight(address , uint256 ) view returns(uint256)
func (_Contract *ContractSession) OverriddenWeight(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.OverriddenWeight(&_Contract.CallOpts, arg0, arg1)
}

// OverriddenWeight is a free data retrieval call binding the contract method 0x2177d6fc.
//
// Solidity: function overriddenWeight(address , uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) OverriddenWeight(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.OverriddenWeight(&_Contract.CallOpts, arg0, arg1)
}

// ProposalBurntFee is a free data retrieval call binding the contract method 0x8a444bd7.
//
// Solidity: function proposalBurntFee() pure returns(uint256)
func (_Contract *ContractCaller) ProposalBurntFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "proposalBurntFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalBurntFee is a free data retrieval call binding the contract method 0x8a444bd7.
//
// Solidity: function proposalBurntFee() pure returns(uint256)
func (_Contract *ContractSession) ProposalBurntFee() (*big.Int, error) {
	return _Contract.Contract.ProposalBurntFee(&_Contract.CallOpts)
}

// ProposalBurntFee is a free data retrieval call binding the contract method 0x8a444bd7.
//
// Solidity: function proposalBurntFee() pure returns(uint256)
func (_Contract *ContractCallerSession) ProposalBurntFee() (*big.Int, error) {
	return _Contract.Contract.ProposalBurntFee(&_Contract.CallOpts)
}

// ProposalFee is a free data retrieval call binding the contract method 0xc27cabb5.
//
// Solidity: function proposalFee() pure returns(uint256)
func (_Contract *ContractCaller) ProposalFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "proposalFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalFee is a free data retrieval call binding the contract method 0xc27cabb5.
//
// Solidity: function proposalFee() pure returns(uint256)
func (_Contract *ContractSession) ProposalFee() (*big.Int, error) {
	return _Contract.Contract.ProposalFee(&_Contract.CallOpts)
}

// ProposalFee is a free data retrieval call binding the contract method 0xc27cabb5.
//
// Solidity: function proposalFee() pure returns(uint256)
func (_Contract *ContractCallerSession) ProposalFee() (*big.Int, error) {
	return _Contract.Contract.ProposalFee(&_Contract.CallOpts)
}

// ProposalOptionState is a free data retrieval call binding the contract method 0x5f89801e.
//
// Solidity: function proposalOptionState(uint256 proposalID, uint256 optionID) view returns(uint256 votes, uint256 agreementRatio, uint256 agreement)
func (_Contract *ContractCaller) ProposalOptionState(opts *bind.CallOpts, proposalID *big.Int, optionID *big.Int) (struct {
	Votes          *big.Int
	AgreementRatio *big.Int
	Agreement      *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "proposalOptionState", proposalID, optionID)

	outstruct := new(struct {
		Votes          *big.Int
		AgreementRatio *big.Int
		Agreement      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Votes = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AgreementRatio = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Agreement = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProposalOptionState is a free data retrieval call binding the contract method 0x5f89801e.
//
// Solidity: function proposalOptionState(uint256 proposalID, uint256 optionID) view returns(uint256 votes, uint256 agreementRatio, uint256 agreement)
func (_Contract *ContractSession) ProposalOptionState(proposalID *big.Int, optionID *big.Int) (struct {
	Votes          *big.Int
	AgreementRatio *big.Int
	Agreement      *big.Int
}, error) {
	return _Contract.Contract.ProposalOptionState(&_Contract.CallOpts, proposalID, optionID)
}

// ProposalOptionState is a free data retrieval call binding the contract method 0x5f89801e.
//
// Solidity: function proposalOptionState(uint256 proposalID, uint256 optionID) view returns(uint256 votes, uint256 agreementRatio, uint256 agreement)
func (_Contract *ContractCallerSession) ProposalOptionState(proposalID *big.Int, optionID *big.Int) (struct {
	Votes          *big.Int
	AgreementRatio *big.Int
	Agreement      *big.Int
}, error) {
	return _Contract.Contract.ProposalOptionState(&_Contract.CallOpts, proposalID, optionID)
}

// ProposalParams is a free data retrieval call binding the contract method 0xcfa1afa3.
//
// Solidity: function proposalParams(uint256 proposalID) view returns(uint256 pType, uint8 executable, uint256 minVotes, uint256 minAgreement, uint256[] opinionScales, bytes32[] options, address proposalContract, uint256 votingStartTime, uint256 votingMinEndTime, uint256 votingMaxEndTime)
func (_Contract *ContractCaller) ProposalParams(opts *bind.CallOpts, proposalID *big.Int) (struct {
	PType            *big.Int
	Executable       uint8
	MinVotes         *big.Int
	MinAgreement     *big.Int
	OpinionScales    []*big.Int
	Options          [][32]byte
	ProposalContract common.Address
	VotingStartTime  *big.Int
	VotingMinEndTime *big.Int
	VotingMaxEndTime *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "proposalParams", proposalID)

	outstruct := new(struct {
		PType            *big.Int
		Executable       uint8
		MinVotes         *big.Int
		MinAgreement     *big.Int
		OpinionScales    []*big.Int
		Options          [][32]byte
		ProposalContract common.Address
		VotingStartTime  *big.Int
		VotingMinEndTime *big.Int
		VotingMaxEndTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PType = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Executable = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.MinVotes = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MinAgreement = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.OpinionScales = *abi.ConvertType(out[4], new([]*big.Int)).(*[]*big.Int)
	outstruct.Options = *abi.ConvertType(out[5], new([][32]byte)).(*[][32]byte)
	outstruct.ProposalContract = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.VotingStartTime = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.VotingMinEndTime = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.VotingMaxEndTime = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProposalParams is a free data retrieval call binding the contract method 0xcfa1afa3.
//
// Solidity: function proposalParams(uint256 proposalID) view returns(uint256 pType, uint8 executable, uint256 minVotes, uint256 minAgreement, uint256[] opinionScales, bytes32[] options, address proposalContract, uint256 votingStartTime, uint256 votingMinEndTime, uint256 votingMaxEndTime)
func (_Contract *ContractSession) ProposalParams(proposalID *big.Int) (struct {
	PType            *big.Int
	Executable       uint8
	MinVotes         *big.Int
	MinAgreement     *big.Int
	OpinionScales    []*big.Int
	Options          [][32]byte
	ProposalContract common.Address
	VotingStartTime  *big.Int
	VotingMinEndTime *big.Int
	VotingMaxEndTime *big.Int
}, error) {
	return _Contract.Contract.ProposalParams(&_Contract.CallOpts, proposalID)
}

// ProposalParams is a free data retrieval call binding the contract method 0xcfa1afa3.
//
// Solidity: function proposalParams(uint256 proposalID) view returns(uint256 pType, uint8 executable, uint256 minVotes, uint256 minAgreement, uint256[] opinionScales, bytes32[] options, address proposalContract, uint256 votingStartTime, uint256 votingMinEndTime, uint256 votingMaxEndTime)
func (_Contract *ContractCallerSession) ProposalParams(proposalID *big.Int) (struct {
	PType            *big.Int
	Executable       uint8
	MinVotes         *big.Int
	MinAgreement     *big.Int
	OpinionScales    []*big.Int
	Options          [][32]byte
	ProposalContract common.Address
	VotingStartTime  *big.Int
	VotingMinEndTime *big.Int
	VotingMaxEndTime *big.Int
}, error) {
	return _Contract.Contract.ProposalParams(&_Contract.CallOpts, proposalID)
}

// ProposalState is a free data retrieval call binding the contract method 0xd26331d4.
//
// Solidity: function proposalState(uint256 proposalID) view returns(uint256 winnerOptionID, uint256 votes, uint256 status)
func (_Contract *ContractCaller) ProposalState(opts *bind.CallOpts, proposalID *big.Int) (struct {
	WinnerOptionID *big.Int
	Votes          *big.Int
	Status         *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "proposalState", proposalID)

	outstruct := new(struct {
		WinnerOptionID *big.Int
		Votes          *big.Int
		Status         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WinnerOptionID = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Votes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProposalState is a free data retrieval call binding the contract method 0xd26331d4.
//
// Solidity: function proposalState(uint256 proposalID) view returns(uint256 winnerOptionID, uint256 votes, uint256 status)
func (_Contract *ContractSession) ProposalState(proposalID *big.Int) (struct {
	WinnerOptionID *big.Int
	Votes          *big.Int
	Status         *big.Int
}, error) {
	return _Contract.Contract.ProposalState(&_Contract.CallOpts, proposalID)
}

// ProposalState is a free data retrieval call binding the contract method 0xd26331d4.
//
// Solidity: function proposalState(uint256 proposalID) view returns(uint256 winnerOptionID, uint256 votes, uint256 status)
func (_Contract *ContractCallerSession) ProposalState(proposalID *big.Int) (struct {
	WinnerOptionID *big.Int
	Votes          *big.Int
	Status         *big.Int
}, error) {
	return _Contract.Contract.ProposalState(&_Contract.CallOpts, proposalID)
}

// TaskErasingReward is a free data retrieval call binding the contract method 0xaea2ddbd.
//
// Solidity: function taskErasingReward() pure returns(uint256)
func (_Contract *ContractCaller) TaskErasingReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "taskErasingReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaskErasingReward is a free data retrieval call binding the contract method 0xaea2ddbd.
//
// Solidity: function taskErasingReward() pure returns(uint256)
func (_Contract *ContractSession) TaskErasingReward() (*big.Int, error) {
	return _Contract.Contract.TaskErasingReward(&_Contract.CallOpts)
}

// TaskErasingReward is a free data retrieval call binding the contract method 0xaea2ddbd.
//
// Solidity: function taskErasingReward() pure returns(uint256)
func (_Contract *ContractCallerSession) TaskErasingReward() (*big.Int, error) {
	return _Contract.Contract.TaskErasingReward(&_Contract.CallOpts)
}

// TaskHandlingReward is a free data retrieval call binding the contract method 0xe3a4d709.
//
// Solidity: function taskHandlingReward() pure returns(uint256)
func (_Contract *ContractCaller) TaskHandlingReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "taskHandlingReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaskHandlingReward is a free data retrieval call binding the contract method 0xe3a4d709.
//
// Solidity: function taskHandlingReward() pure returns(uint256)
func (_Contract *ContractSession) TaskHandlingReward() (*big.Int, error) {
	return _Contract.Contract.TaskHandlingReward(&_Contract.CallOpts)
}

// TaskHandlingReward is a free data retrieval call binding the contract method 0xe3a4d709.
//
// Solidity: function taskHandlingReward() pure returns(uint256)
func (_Contract *ContractCallerSession) TaskHandlingReward() (*big.Int, error) {
	return _Contract.Contract.TaskHandlingReward(&_Contract.CallOpts)
}

// TasksCount is a free data retrieval call binding the contract method 0xbb6a0f07.
//
// Solidity: function tasksCount() view returns(uint256)
func (_Contract *ContractCaller) TasksCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "tasksCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TasksCount is a free data retrieval call binding the contract method 0xbb6a0f07.
//
// Solidity: function tasksCount() view returns(uint256)
func (_Contract *ContractSession) TasksCount() (*big.Int, error) {
	return _Contract.Contract.TasksCount(&_Contract.CallOpts)
}

// TasksCount is a free data retrieval call binding the contract method 0xbb6a0f07.
//
// Solidity: function tasksCount() view returns(uint256)
func (_Contract *ContractCallerSession) TasksCount() (*big.Int, error) {
	return _Contract.Contract.TasksCount(&_Contract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(bytes4)
func (_Contract *ContractCaller) Version(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "version")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(bytes4)
func (_Contract *ContractSession) Version() ([4]byte, error) {
	return _Contract.Contract.Version(&_Contract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(bytes4)
func (_Contract *ContractCallerSession) Version() ([4]byte, error) {
	return _Contract.Contract.Version(&_Contract.CallOpts)
}

// CancelProposal is a paid mutator transaction binding the contract method 0xe0a8f6f5.
//
// Solidity: function cancelProposal(uint256 proposalID) returns()
func (_Contract *ContractTransactor) CancelProposal(opts *bind.TransactOpts, proposalID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "cancelProposal", proposalID)
}

// CancelProposal is a paid mutator transaction binding the contract method 0xe0a8f6f5.
//
// Solidity: function cancelProposal(uint256 proposalID) returns()
func (_Contract *ContractSession) CancelProposal(proposalID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CancelProposal(&_Contract.TransactOpts, proposalID)
}

// CancelProposal is a paid mutator transaction binding the contract method 0xe0a8f6f5.
//
// Solidity: function cancelProposal(uint256 proposalID) returns()
func (_Contract *ContractTransactorSession) CancelProposal(proposalID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CancelProposal(&_Contract.TransactOpts, proposalID)
}

// CancelVote is a paid mutator transaction binding the contract method 0x21edf2eb.
//
// Solidity: function cancelVote(address delegatedTo, uint256 proposalID) returns()
func (_Contract *ContractTransactor) CancelVote(opts *bind.TransactOpts, delegatedTo common.Address, proposalID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "cancelVote", delegatedTo, proposalID)
}

// CancelVote is a paid mutator transaction binding the contract method 0x21edf2eb.
//
// Solidity: function cancelVote(address delegatedTo, uint256 proposalID) returns()
func (_Contract *ContractSession) CancelVote(delegatedTo common.Address, proposalID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CancelVote(&_Contract.TransactOpts, delegatedTo, proposalID)
}

// CancelVote is a paid mutator transaction binding the contract method 0x21edf2eb.
//
// Solidity: function cancelVote(address delegatedTo, uint256 proposalID) returns()
func (_Contract *ContractTransactorSession) CancelVote(delegatedTo common.Address, proposalID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CancelVote(&_Contract.TransactOpts, delegatedTo, proposalID)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x07eecaf5.
//
// Solidity: function createProposal(address proposalContract) payable returns()
func (_Contract *ContractTransactor) CreateProposal(opts *bind.TransactOpts, proposalContract common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createProposal", proposalContract)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x07eecaf5.
//
// Solidity: function createProposal(address proposalContract) payable returns()
func (_Contract *ContractSession) CreateProposal(proposalContract common.Address) (*types.Transaction, error) {
	return _Contract.Contract.CreateProposal(&_Contract.TransactOpts, proposalContract)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x07eecaf5.
//
// Solidity: function createProposal(address proposalContract) payable returns()
func (_Contract *ContractTransactorSession) CreateProposal(proposalContract common.Address) (*types.Transaction, error) {
	return _Contract.Contract.CreateProposal(&_Contract.TransactOpts, proposalContract)
}

// HandleTasks is a paid mutator transaction binding the contract method 0x6b2ad7d8.
//
// Solidity: function handleTasks(uint256 startIdx, uint256 quantity) returns()
func (_Contract *ContractTransactor) HandleTasks(opts *bind.TransactOpts, startIdx *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "handleTasks", startIdx, quantity)
}

// HandleTasks is a paid mutator transaction binding the contract method 0x6b2ad7d8.
//
// Solidity: function handleTasks(uint256 startIdx, uint256 quantity) returns()
func (_Contract *ContractSession) HandleTasks(startIdx *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.HandleTasks(&_Contract.TransactOpts, startIdx, quantity)
}

// HandleTasks is a paid mutator transaction binding the contract method 0x6b2ad7d8.
//
// Solidity: function handleTasks(uint256 startIdx, uint256 quantity) returns()
func (_Contract *ContractTransactorSession) HandleTasks(startIdx *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.HandleTasks(&_Contract.TransactOpts, startIdx, quantity)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _governableContract, address _proposalVerifier) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, _governableContract common.Address, _proposalVerifier common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", _governableContract, _proposalVerifier)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _governableContract, address _proposalVerifier) returns()
func (_Contract *ContractSession) Initialize(_governableContract common.Address, _proposalVerifier common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _governableContract, _proposalVerifier)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _governableContract, address _proposalVerifier) returns()
func (_Contract *ContractTransactorSession) Initialize(_governableContract common.Address, _proposalVerifier common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _governableContract, _proposalVerifier)
}

// RecountVote is a paid mutator transaction binding the contract method 0x69921506.
//
// Solidity: function recountVote(address voterAddr, address delegatedTo, uint256 proposalID) returns()
func (_Contract *ContractTransactor) RecountVote(opts *bind.TransactOpts, voterAddr common.Address, delegatedTo common.Address, proposalID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "recountVote", voterAddr, delegatedTo, proposalID)
}

// RecountVote is a paid mutator transaction binding the contract method 0x69921506.
//
// Solidity: function recountVote(address voterAddr, address delegatedTo, uint256 proposalID) returns()
func (_Contract *ContractSession) RecountVote(voterAddr common.Address, delegatedTo common.Address, proposalID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RecountVote(&_Contract.TransactOpts, voterAddr, delegatedTo, proposalID)
}

// RecountVote is a paid mutator transaction binding the contract method 0x69921506.
//
// Solidity: function recountVote(address voterAddr, address delegatedTo, uint256 proposalID) returns()
func (_Contract *ContractTransactorSession) RecountVote(voterAddr common.Address, delegatedTo common.Address, proposalID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RecountVote(&_Contract.TransactOpts, voterAddr, delegatedTo, proposalID)
}

// TasksCleanup is a paid mutator transaction binding the contract method 0x491adeee.
//
// Solidity: function tasksCleanup(uint256 quantity) returns()
func (_Contract *ContractTransactor) TasksCleanup(opts *bind.TransactOpts, quantity *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "tasksCleanup", quantity)
}

// TasksCleanup is a paid mutator transaction binding the contract method 0x491adeee.
//
// Solidity: function tasksCleanup(uint256 quantity) returns()
func (_Contract *ContractSession) TasksCleanup(quantity *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TasksCleanup(&_Contract.TransactOpts, quantity)
}

// TasksCleanup is a paid mutator transaction binding the contract method 0x491adeee.
//
// Solidity: function tasksCleanup(uint256 quantity) returns()
func (_Contract *ContractTransactorSession) TasksCleanup(quantity *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TasksCleanup(&_Contract.TransactOpts, quantity)
}

// Vote is a paid mutator transaction binding the contract method 0x172c18b1.
//
// Solidity: function vote(address delegatedTo, uint256 proposalID, uint256[] choices) returns()
func (_Contract *ContractTransactor) Vote(opts *bind.TransactOpts, delegatedTo common.Address, proposalID *big.Int, choices []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "vote", delegatedTo, proposalID, choices)
}

// Vote is a paid mutator transaction binding the contract method 0x172c18b1.
//
// Solidity: function vote(address delegatedTo, uint256 proposalID, uint256[] choices) returns()
func (_Contract *ContractSession) Vote(delegatedTo common.Address, proposalID *big.Int, choices []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Vote(&_Contract.TransactOpts, delegatedTo, proposalID, choices)
}

// Vote is a paid mutator transaction binding the contract method 0x172c18b1.
//
// Solidity: function vote(address delegatedTo, uint256 proposalID, uint256[] choices) returns()
func (_Contract *ContractTransactorSession) Vote(delegatedTo common.Address, proposalID *big.Int, choices []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Vote(&_Contract.TransactOpts, delegatedTo, proposalID, choices)
}

// ContractProposalCanceledIterator is returned from FilterProposalCanceled and is used to iterate over the raw logs and unpacked data for ProposalCanceled events raised by the Contract contract.
type ContractProposalCanceledIterator struct {
	Event *ContractProposalCanceled // Event containing the contract specifics and raw log

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
func (it *ContractProposalCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractProposalCanceled)
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
		it.Event = new(ContractProposalCanceled)
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
func (it *ContractProposalCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractProposalCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractProposalCanceled represents a ProposalCanceled event raised by the Contract contract.
type ContractProposalCanceled struct {
	ProposalID *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalCanceled is a free log retrieval operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalID)
func (_Contract *ContractFilterer) FilterProposalCanceled(opts *bind.FilterOpts) (*ContractProposalCanceledIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return &ContractProposalCanceledIterator{contract: _Contract.contract, event: "ProposalCanceled", logs: logs, sub: sub}, nil
}

// WatchProposalCanceled is a free log subscription operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalID)
func (_Contract *ContractFilterer) WatchProposalCanceled(opts *bind.WatchOpts, sink chan<- *ContractProposalCanceled) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractProposalCanceled)
				if err := _Contract.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
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

// ParseProposalCanceled is a log parse operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalID)
func (_Contract *ContractFilterer) ParseProposalCanceled(log types.Log) (*ContractProposalCanceled, error) {
	event := new(ContractProposalCanceled)
	if err := _Contract.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the Contract contract.
type ContractProposalCreatedIterator struct {
	Event *ContractProposalCreated // Event containing the contract specifics and raw log

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
func (it *ContractProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractProposalCreated)
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
		it.Event = new(ContractProposalCreated)
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
func (it *ContractProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractProposalCreated represents a ProposalCreated event raised by the Contract contract.
type ContractProposalCreated struct {
	ProposalID *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0xc2c021f5d73c63c481d336fbbafec58f694fc45095f00b02d2deb8cca59afe07.
//
// Solidity: event ProposalCreated(uint256 proposalID)
func (_Contract *ContractFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*ContractProposalCreatedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &ContractProposalCreatedIterator{contract: _Contract.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0xc2c021f5d73c63c481d336fbbafec58f694fc45095f00b02d2deb8cca59afe07.
//
// Solidity: event ProposalCreated(uint256 proposalID)
func (_Contract *ContractFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *ContractProposalCreated) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractProposalCreated)
				if err := _Contract.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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

// ParseProposalCreated is a log parse operation binding the contract event 0xc2c021f5d73c63c481d336fbbafec58f694fc45095f00b02d2deb8cca59afe07.
//
// Solidity: event ProposalCreated(uint256 proposalID)
func (_Contract *ContractFilterer) ParseProposalCreated(log types.Log) (*ContractProposalCreated, error) {
	event := new(ContractProposalCreated)
	if err := _Contract.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractProposalExecutionExpiredIterator is returned from FilterProposalExecutionExpired and is used to iterate over the raw logs and unpacked data for ProposalExecutionExpired events raised by the Contract contract.
type ContractProposalExecutionExpiredIterator struct {
	Event *ContractProposalExecutionExpired // Event containing the contract specifics and raw log

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
func (it *ContractProposalExecutionExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractProposalExecutionExpired)
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
		it.Event = new(ContractProposalExecutionExpired)
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
func (it *ContractProposalExecutionExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractProposalExecutionExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractProposalExecutionExpired represents a ProposalExecutionExpired event raised by the Contract contract.
type ContractProposalExecutionExpired struct {
	ProposalID *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecutionExpired is a free log retrieval operation binding the contract event 0xe8365dd25802fb5113a4ebd6fbe5fee885b5ea470b6b1467f3f4df69e490ed87.
//
// Solidity: event ProposalExecutionExpired(uint256 proposalID)
func (_Contract *ContractFilterer) FilterProposalExecutionExpired(opts *bind.FilterOpts) (*ContractProposalExecutionExpiredIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ProposalExecutionExpired")
	if err != nil {
		return nil, err
	}
	return &ContractProposalExecutionExpiredIterator{contract: _Contract.contract, event: "ProposalExecutionExpired", logs: logs, sub: sub}, nil
}

// WatchProposalExecutionExpired is a free log subscription operation binding the contract event 0xe8365dd25802fb5113a4ebd6fbe5fee885b5ea470b6b1467f3f4df69e490ed87.
//
// Solidity: event ProposalExecutionExpired(uint256 proposalID)
func (_Contract *ContractFilterer) WatchProposalExecutionExpired(opts *bind.WatchOpts, sink chan<- *ContractProposalExecutionExpired) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ProposalExecutionExpired")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractProposalExecutionExpired)
				if err := _Contract.contract.UnpackLog(event, "ProposalExecutionExpired", log); err != nil {
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

// ParseProposalExecutionExpired is a log parse operation binding the contract event 0xe8365dd25802fb5113a4ebd6fbe5fee885b5ea470b6b1467f3f4df69e490ed87.
//
// Solidity: event ProposalExecutionExpired(uint256 proposalID)
func (_Contract *ContractFilterer) ParseProposalExecutionExpired(log types.Log) (*ContractProposalExecutionExpired, error) {
	event := new(ContractProposalExecutionExpired)
	if err := _Contract.contract.UnpackLog(event, "ProposalExecutionExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractProposalRejectedIterator is returned from FilterProposalRejected and is used to iterate over the raw logs and unpacked data for ProposalRejected events raised by the Contract contract.
type ContractProposalRejectedIterator struct {
	Event *ContractProposalRejected // Event containing the contract specifics and raw log

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
func (it *ContractProposalRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractProposalRejected)
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
		it.Event = new(ContractProposalRejected)
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
func (it *ContractProposalRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractProposalRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractProposalRejected represents a ProposalRejected event raised by the Contract contract.
type ContractProposalRejected struct {
	ProposalID *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalRejected is a free log retrieval operation binding the contract event 0xd92fba445edb3153b571e6df782d7a66fd0ce668519273670820ee3a86da0ef4.
//
// Solidity: event ProposalRejected(uint256 proposalID)
func (_Contract *ContractFilterer) FilterProposalRejected(opts *bind.FilterOpts) (*ContractProposalRejectedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ProposalRejected")
	if err != nil {
		return nil, err
	}
	return &ContractProposalRejectedIterator{contract: _Contract.contract, event: "ProposalRejected", logs: logs, sub: sub}, nil
}

// WatchProposalRejected is a free log subscription operation binding the contract event 0xd92fba445edb3153b571e6df782d7a66fd0ce668519273670820ee3a86da0ef4.
//
// Solidity: event ProposalRejected(uint256 proposalID)
func (_Contract *ContractFilterer) WatchProposalRejected(opts *bind.WatchOpts, sink chan<- *ContractProposalRejected) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ProposalRejected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractProposalRejected)
				if err := _Contract.contract.UnpackLog(event, "ProposalRejected", log); err != nil {
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

// ParseProposalRejected is a log parse operation binding the contract event 0xd92fba445edb3153b571e6df782d7a66fd0ce668519273670820ee3a86da0ef4.
//
// Solidity: event ProposalRejected(uint256 proposalID)
func (_Contract *ContractFilterer) ParseProposalRejected(log types.Log) (*ContractProposalRejected, error) {
	event := new(ContractProposalRejected)
	if err := _Contract.contract.UnpackLog(event, "ProposalRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractProposalResolvedIterator is returned from FilterProposalResolved and is used to iterate over the raw logs and unpacked data for ProposalResolved events raised by the Contract contract.
type ContractProposalResolvedIterator struct {
	Event *ContractProposalResolved // Event containing the contract specifics and raw log

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
func (it *ContractProposalResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractProposalResolved)
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
		it.Event = new(ContractProposalResolved)
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
func (it *ContractProposalResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractProposalResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractProposalResolved represents a ProposalResolved event raised by the Contract contract.
type ContractProposalResolved struct {
	ProposalID *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalResolved is a free log retrieval operation binding the contract event 0x663674d96fd5c2a954bf75ad2e6795f9c9701eb687a7a8f3297c7a299467c941.
//
// Solidity: event ProposalResolved(uint256 proposalID)
func (_Contract *ContractFilterer) FilterProposalResolved(opts *bind.FilterOpts) (*ContractProposalResolvedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ProposalResolved")
	if err != nil {
		return nil, err
	}
	return &ContractProposalResolvedIterator{contract: _Contract.contract, event: "ProposalResolved", logs: logs, sub: sub}, nil
}

// WatchProposalResolved is a free log subscription operation binding the contract event 0x663674d96fd5c2a954bf75ad2e6795f9c9701eb687a7a8f3297c7a299467c941.
//
// Solidity: event ProposalResolved(uint256 proposalID)
func (_Contract *ContractFilterer) WatchProposalResolved(opts *bind.WatchOpts, sink chan<- *ContractProposalResolved) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ProposalResolved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractProposalResolved)
				if err := _Contract.contract.UnpackLog(event, "ProposalResolved", log); err != nil {
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

// ParseProposalResolved is a log parse operation binding the contract event 0x663674d96fd5c2a954bf75ad2e6795f9c9701eb687a7a8f3297c7a299467c941.
//
// Solidity: event ProposalResolved(uint256 proposalID)
func (_Contract *ContractFilterer) ParseProposalResolved(log types.Log) (*ContractProposalResolved, error) {
	event := new(ContractProposalResolved)
	if err := _Contract.contract.UnpackLog(event, "ProposalResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTasksErasedIterator is returned from FilterTasksErased and is used to iterate over the raw logs and unpacked data for TasksErased events raised by the Contract contract.
type ContractTasksErasedIterator struct {
	Event *ContractTasksErased // Event containing the contract specifics and raw log

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
func (it *ContractTasksErasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTasksErased)
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
		it.Event = new(ContractTasksErased)
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
func (it *ContractTasksErasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTasksErasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTasksErased represents a TasksErased event raised by the Contract contract.
type ContractTasksErased struct {
	Quantity *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTasksErased is a free log retrieval operation binding the contract event 0xdd64b087b4726a6d5006bfae34869d1044e415cd6fc53054b69619ee337b730d.
//
// Solidity: event TasksErased(uint256 quantity)
func (_Contract *ContractFilterer) FilterTasksErased(opts *bind.FilterOpts) (*ContractTasksErasedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TasksErased")
	if err != nil {
		return nil, err
	}
	return &ContractTasksErasedIterator{contract: _Contract.contract, event: "TasksErased", logs: logs, sub: sub}, nil
}

// WatchTasksErased is a free log subscription operation binding the contract event 0xdd64b087b4726a6d5006bfae34869d1044e415cd6fc53054b69619ee337b730d.
//
// Solidity: event TasksErased(uint256 quantity)
func (_Contract *ContractFilterer) WatchTasksErased(opts *bind.WatchOpts, sink chan<- *ContractTasksErased) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TasksErased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTasksErased)
				if err := _Contract.contract.UnpackLog(event, "TasksErased", log); err != nil {
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

// ParseTasksErased is a log parse operation binding the contract event 0xdd64b087b4726a6d5006bfae34869d1044e415cd6fc53054b69619ee337b730d.
//
// Solidity: event TasksErased(uint256 quantity)
func (_Contract *ContractFilterer) ParseTasksErased(log types.Log) (*ContractTasksErased, error) {
	event := new(ContractTasksErased)
	if err := _Contract.contract.UnpackLog(event, "TasksErased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTasksHandledIterator is returned from FilterTasksHandled and is used to iterate over the raw logs and unpacked data for TasksHandled events raised by the Contract contract.
type ContractTasksHandledIterator struct {
	Event *ContractTasksHandled // Event containing the contract specifics and raw log

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
func (it *ContractTasksHandledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTasksHandled)
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
		it.Event = new(ContractTasksHandled)
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
func (it *ContractTasksHandledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTasksHandledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTasksHandled represents a TasksHandled event raised by the Contract contract.
type ContractTasksHandled struct {
	StartIdx *big.Int
	EndIdx   *big.Int
	Handled  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTasksHandled is a free log retrieval operation binding the contract event 0x010b51ac92c2d2e8fe9c84de31065ca878ff8274a1e5e6d31799520b7d16bb58.
//
// Solidity: event TasksHandled(uint256 startIdx, uint256 endIdx, uint256 handled)
func (_Contract *ContractFilterer) FilterTasksHandled(opts *bind.FilterOpts) (*ContractTasksHandledIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TasksHandled")
	if err != nil {
		return nil, err
	}
	return &ContractTasksHandledIterator{contract: _Contract.contract, event: "TasksHandled", logs: logs, sub: sub}, nil
}

// WatchTasksHandled is a free log subscription operation binding the contract event 0x010b51ac92c2d2e8fe9c84de31065ca878ff8274a1e5e6d31799520b7d16bb58.
//
// Solidity: event TasksHandled(uint256 startIdx, uint256 endIdx, uint256 handled)
func (_Contract *ContractFilterer) WatchTasksHandled(opts *bind.WatchOpts, sink chan<- *ContractTasksHandled) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TasksHandled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTasksHandled)
				if err := _Contract.contract.UnpackLog(event, "TasksHandled", log); err != nil {
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

// ParseTasksHandled is a log parse operation binding the contract event 0x010b51ac92c2d2e8fe9c84de31065ca878ff8274a1e5e6d31799520b7d16bb58.
//
// Solidity: event TasksHandled(uint256 startIdx, uint256 endIdx, uint256 handled)
func (_Contract *ContractFilterer) ParseTasksHandled(log types.Log) (*ContractTasksHandled, error) {
	event := new(ContractTasksHandled)
	if err := _Contract.contract.UnpackLog(event, "TasksHandled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractVoteCanceledIterator is returned from FilterVoteCanceled and is used to iterate over the raw logs and unpacked data for VoteCanceled events raised by the Contract contract.
type ContractVoteCanceledIterator struct {
	Event *ContractVoteCanceled // Event containing the contract specifics and raw log

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
func (it *ContractVoteCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractVoteCanceled)
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
		it.Event = new(ContractVoteCanceled)
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
func (it *ContractVoteCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractVoteCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractVoteCanceled represents a VoteCanceled event raised by the Contract contract.
type ContractVoteCanceled struct {
	Voter       common.Address
	DelegatedTo common.Address
	ProposalID  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVoteCanceled is a free log retrieval operation binding the contract event 0x666685d133047310e2a2e8c4f6794b6dccb4e9ad9c6903ac753fb10d8918b649.
//
// Solidity: event VoteCanceled(address voter, address delegatedTo, uint256 proposalID)
func (_Contract *ContractFilterer) FilterVoteCanceled(opts *bind.FilterOpts) (*ContractVoteCanceledIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "VoteCanceled")
	if err != nil {
		return nil, err
	}
	return &ContractVoteCanceledIterator{contract: _Contract.contract, event: "VoteCanceled", logs: logs, sub: sub}, nil
}

// WatchVoteCanceled is a free log subscription operation binding the contract event 0x666685d133047310e2a2e8c4f6794b6dccb4e9ad9c6903ac753fb10d8918b649.
//
// Solidity: event VoteCanceled(address voter, address delegatedTo, uint256 proposalID)
func (_Contract *ContractFilterer) WatchVoteCanceled(opts *bind.WatchOpts, sink chan<- *ContractVoteCanceled) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "VoteCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractVoteCanceled)
				if err := _Contract.contract.UnpackLog(event, "VoteCanceled", log); err != nil {
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

// ParseVoteCanceled is a log parse operation binding the contract event 0x666685d133047310e2a2e8c4f6794b6dccb4e9ad9c6903ac753fb10d8918b649.
//
// Solidity: event VoteCanceled(address voter, address delegatedTo, uint256 proposalID)
func (_Contract *ContractFilterer) ParseVoteCanceled(log types.Log) (*ContractVoteCanceled, error) {
	event := new(ContractVoteCanceled)
	if err := _Contract.contract.UnpackLog(event, "VoteCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractVoteWeightOverriddenIterator is returned from FilterVoteWeightOverridden and is used to iterate over the raw logs and unpacked data for VoteWeightOverridden events raised by the Contract contract.
type ContractVoteWeightOverriddenIterator struct {
	Event *ContractVoteWeightOverridden // Event containing the contract specifics and raw log

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
func (it *ContractVoteWeightOverriddenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractVoteWeightOverridden)
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
		it.Event = new(ContractVoteWeightOverridden)
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
func (it *ContractVoteWeightOverriddenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractVoteWeightOverriddenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractVoteWeightOverridden represents a VoteWeightOverridden event raised by the Contract contract.
type ContractVoteWeightOverridden struct {
	Voter common.Address
	Diff  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVoteWeightOverridden is a free log retrieval operation binding the contract event 0x68fe85c5f71a2900fddf574935a27d4f1cb28af34d4fa2742b202684b45d3d14.
//
// Solidity: event VoteWeightOverridden(address voter, uint256 diff)
func (_Contract *ContractFilterer) FilterVoteWeightOverridden(opts *bind.FilterOpts) (*ContractVoteWeightOverriddenIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "VoteWeightOverridden")
	if err != nil {
		return nil, err
	}
	return &ContractVoteWeightOverriddenIterator{contract: _Contract.contract, event: "VoteWeightOverridden", logs: logs, sub: sub}, nil
}

// WatchVoteWeightOverridden is a free log subscription operation binding the contract event 0x68fe85c5f71a2900fddf574935a27d4f1cb28af34d4fa2742b202684b45d3d14.
//
// Solidity: event VoteWeightOverridden(address voter, uint256 diff)
func (_Contract *ContractFilterer) WatchVoteWeightOverridden(opts *bind.WatchOpts, sink chan<- *ContractVoteWeightOverridden) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "VoteWeightOverridden")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractVoteWeightOverridden)
				if err := _Contract.contract.UnpackLog(event, "VoteWeightOverridden", log); err != nil {
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

// ParseVoteWeightOverridden is a log parse operation binding the contract event 0x68fe85c5f71a2900fddf574935a27d4f1cb28af34d4fa2742b202684b45d3d14.
//
// Solidity: event VoteWeightOverridden(address voter, uint256 diff)
func (_Contract *ContractFilterer) ParseVoteWeightOverridden(log types.Log) (*ContractVoteWeightOverridden, error) {
	event := new(ContractVoteWeightOverridden)
	if err := _Contract.contract.UnpackLog(event, "VoteWeightOverridden", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractVoteWeightUnOverriddenIterator is returned from FilterVoteWeightUnOverridden and is used to iterate over the raw logs and unpacked data for VoteWeightUnOverridden events raised by the Contract contract.
type ContractVoteWeightUnOverriddenIterator struct {
	Event *ContractVoteWeightUnOverridden // Event containing the contract specifics and raw log

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
func (it *ContractVoteWeightUnOverriddenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractVoteWeightUnOverridden)
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
		it.Event = new(ContractVoteWeightUnOverridden)
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
func (it *ContractVoteWeightUnOverriddenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractVoteWeightUnOverriddenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractVoteWeightUnOverridden represents a VoteWeightUnOverridden event raised by the Contract contract.
type ContractVoteWeightUnOverridden struct {
	Voter common.Address
	Diff  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVoteWeightUnOverridden is a free log retrieval operation binding the contract event 0x11d7313426c62d856bd4ea20cdef8b93af4b40d2ea5b8f6f962fc705dbdcdbef.
//
// Solidity: event VoteWeightUnOverridden(address voter, uint256 diff)
func (_Contract *ContractFilterer) FilterVoteWeightUnOverridden(opts *bind.FilterOpts) (*ContractVoteWeightUnOverriddenIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "VoteWeightUnOverridden")
	if err != nil {
		return nil, err
	}
	return &ContractVoteWeightUnOverriddenIterator{contract: _Contract.contract, event: "VoteWeightUnOverridden", logs: logs, sub: sub}, nil
}

// WatchVoteWeightUnOverridden is a free log subscription operation binding the contract event 0x11d7313426c62d856bd4ea20cdef8b93af4b40d2ea5b8f6f962fc705dbdcdbef.
//
// Solidity: event VoteWeightUnOverridden(address voter, uint256 diff)
func (_Contract *ContractFilterer) WatchVoteWeightUnOverridden(opts *bind.WatchOpts, sink chan<- *ContractVoteWeightUnOverridden) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "VoteWeightUnOverridden")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractVoteWeightUnOverridden)
				if err := _Contract.contract.UnpackLog(event, "VoteWeightUnOverridden", log); err != nil {
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

// ParseVoteWeightUnOverridden is a log parse operation binding the contract event 0x11d7313426c62d856bd4ea20cdef8b93af4b40d2ea5b8f6f962fc705dbdcdbef.
//
// Solidity: event VoteWeightUnOverridden(address voter, uint256 diff)
func (_Contract *ContractFilterer) ParseVoteWeightUnOverridden(log types.Log) (*ContractVoteWeightUnOverridden, error) {
	event := new(ContractVoteWeightUnOverridden)
	if err := _Contract.contract.UnpackLog(event, "VoteWeightUnOverridden", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the Contract contract.
type ContractVotedIterator struct {
	Event *ContractVoted // Event containing the contract specifics and raw log

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
func (it *ContractVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractVoted)
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
		it.Event = new(ContractVoted)
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
func (it *ContractVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractVoted represents a Voted event raised by the Contract contract.
type ContractVoted struct {
	Voter       common.Address
	DelegatedTo common.Address
	ProposalID  *big.Int
	Choices     []*big.Int
	Weight      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x6e5f0f6e0ce2bdcdb0a82952fc6eb90c4c22f0b6228e4619b5dc2118e1166a12.
//
// Solidity: event Voted(address voter, address delegatedTo, uint256 proposalID, uint256[] choices, uint256 weight)
func (_Contract *ContractFilterer) FilterVoted(opts *bind.FilterOpts) (*ContractVotedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Voted")
	if err != nil {
		return nil, err
	}
	return &ContractVotedIterator{contract: _Contract.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x6e5f0f6e0ce2bdcdb0a82952fc6eb90c4c22f0b6228e4619b5dc2118e1166a12.
//
// Solidity: event Voted(address voter, address delegatedTo, uint256 proposalID, uint256[] choices, uint256 weight)
func (_Contract *ContractFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *ContractVoted) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Voted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractVoted)
				if err := _Contract.contract.UnpackLog(event, "Voted", log); err != nil {
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

// ParseVoted is a log parse operation binding the contract event 0x6e5f0f6e0ce2bdcdb0a82952fc6eb90c4c22f0b6228e4619b5dc2118e1166a12.
//
// Solidity: event Voted(address voter, address delegatedTo, uint256 proposalID, uint256[] choices, uint256 weight)
func (_Contract *ContractFilterer) ParseVoted(log types.Log) (*ContractVoted, error) {
	event := new(ContractVoted)
	if err := _Contract.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

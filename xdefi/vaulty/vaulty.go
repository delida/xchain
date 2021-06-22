// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vaulty

import (
	"math/big"
	"strings"

	"github.com/MOACChain/MoacLib/common"
	"github.com/MOACChain/MoacLib/types"
	moaccore "github.com/MOACChain/xchain"
	"github.com/MOACChain/xchain/accounts/abi"
	"github.com/MOACChain/xchain/accounts/abi/bind"
	"github.com/MOACChain/xchain/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = moaccore.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RoleAccessRole is an auto generated low-level Go binding around an user-defined struct.
type RoleAccessRole struct {
	Role     [32]byte
	Describe string
}

// VaultYtokenPairInfo is an auto generated low-level Go binding around an user-defined struct.
type VaultYtokenPairInfo struct {
	SourceToken        common.Address
	SourceTokenChainid *big.Int
	SourceTokenSymbol  string
	MappedToken        common.Address
	MappedTokenChainid *big.Int
	MappedTokenSymbol  string
	Paused             bool
}

// VaultYABI is the input ABI used to generate the binding from.
const VaultYABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newFeeAccount\",\"type\":\"address\"}],\"name\":\"FeeAccountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newFiatCurrency\",\"type\":\"string\"}],\"name\":\"FiatCurrencyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFiatFeeAmount\",\"type\":\"uint256\"}],\"name\":\"FiatFeeAmountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"nonces\",\"type\":\"uint256[]\"}],\"name\":\"IgnoreNonces\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPriceOracle\",\"type\":\"address\"}],\"name\":\"PriceOracleChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Refund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"step\",\"type\":\"uint256\"}],\"name\":\"SkipNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTipAccount\",\"type\":\"address\"}],\"name\":\"TipAccountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTipRate\",\"type\":\"uint256\"}],\"name\":\"TipRateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mappedChainid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tip\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"TokenBurn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mappedChainid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tip\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"TokenDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"CreatedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"name\":\"addRoleMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoles\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"describe\",\"type\":\"string\"}],\"internalType\":\"structRoleAccess.Role[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getTip\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"omitNonces\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"}],\"name\":\"pauseTokenMapping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"name\":\"removeRoleMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeAccount\",\"type\":\"address\"}],\"name\":\"setFeeAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newFiatCurrency\",\"type\":\"string\"}],\"name\":\"setFiatCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setFiatFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPriceoracle\",\"type\":\"address\"}],\"name\":\"setPriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTipAccount\",\"type\":\"address\"}],\"name\":\"setTipAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newTipRate\",\"type\":\"uint256\"}],\"name\":\"setTipRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenMappingWatermark\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpauseAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"}],\"name\":\"unpauseTokenMapping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sourceTokenSymbol_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"mappedTokenSymbol_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tipRate\",\"type\":\"uint256\"}],\"name\":\"setupTokenMapping\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sourceTokenChainid\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourceTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mappedTokenChainid\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"mappedTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"internalType\":\"structVaultY.tokenPairInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"batchMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"grantMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tipY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tipX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"tipCashout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"tipBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nonces\",\"type\":\"uint256[]\"}],\"name\":\"addNoncesToOmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nonces\",\"type\":\"uint256[]\"}],\"name\":\"removeNoncesToOmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"skip\",\"type\":\"uint256\"}],\"name\":\"skipMintWatermark\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// VaultY is an auto generated Go binding around an Ethereum contract.
type VaultY struct {
	VaultYCaller     // Read-only binding to the contract
	VaultYTransactor // Write-only binding to the contract
	VaultYFilterer   // Log filterer for contract events
}

// VaultYCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultYCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultYTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultYTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultYFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultYFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultYSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultYSession struct {
	Contract     *VaultY           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultYCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultYCallerSession struct {
	Contract *VaultYCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VaultYTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultYTransactorSession struct {
	Contract     *VaultYTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultYRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultYRaw struct {
	Contract *VaultY // Generic contract binding to access the raw methods on
}

// VaultYCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultYCallerRaw struct {
	Contract *VaultYCaller // Generic read-only contract binding to access the raw methods on
}

// VaultYTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultYTransactorRaw struct {
	Contract *VaultYTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultY creates a new instance of VaultY, bound to a specific deployed contract.
func NewVaultY(address common.Address, backend bind.ContractBackend) (*VaultY, error) {
	contract, err := bindVaultY(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultY{VaultYCaller: VaultYCaller{contract: contract}, VaultYTransactor: VaultYTransactor{contract: contract}, VaultYFilterer: VaultYFilterer{contract: contract}}, nil
}

// NewVaultYCaller creates a new read-only instance of VaultY, bound to a specific deployed contract.
func NewVaultYCaller(address common.Address, caller bind.ContractCaller) (*VaultYCaller, error) {
	contract, err := bindVaultY(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultYCaller{contract: contract}, nil
}

// NewVaultYTransactor creates a new write-only instance of VaultY, bound to a specific deployed contract.
func NewVaultYTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultYTransactor, error) {
	contract, err := bindVaultY(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultYTransactor{contract: contract}, nil
}

// NewVaultYFilterer creates a new log filterer instance of VaultY, bound to a specific deployed contract.
func NewVaultYFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultYFilterer, error) {
	contract, err := bindVaultY(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultYFilterer{contract: contract}, nil
}

// bindVaultY binds a generic wrapper to an already deployed contract.
func bindVaultY(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultYABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultY *VaultYRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultY.Contract.VaultYCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultY *VaultYRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultY.Contract.VaultYTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultY *VaultYRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultY.Contract.VaultYTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultY *VaultYCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultY.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultY *VaultYTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultY.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultY *VaultYTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultY.Contract.contract.Transact(opts, method, params...)
}

// CreatedAt is a free data retrieval call binding the contract method 0x748140a3.
//
// Solidity: function CreatedAt() view returns(uint256)
func (_VaultY *VaultYCaller) CreatedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "CreatedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreatedAt is a free data retrieval call binding the contract method 0x748140a3.
//
// Solidity: function CreatedAt() view returns(uint256)
func (_VaultY *VaultYSession) CreatedAt() (*big.Int, error) {
	return _VaultY.Contract.CreatedAt(&_VaultY.CallOpts)
}

// CreatedAt is a free data retrieval call binding the contract method 0x748140a3.
//
// Solidity: function CreatedAt() view returns(uint256)
func (_VaultY *VaultYCallerSession) CreatedAt() (*big.Int, error) {
	return _VaultY.Contract.CreatedAt(&_VaultY.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VaultY *VaultYCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VaultY *VaultYSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _VaultY.Contract.DEFAULTADMINROLE(&_VaultY.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VaultY *VaultYCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _VaultY.Contract.DEFAULTADMINROLE(&_VaultY.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VaultY *VaultYCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VaultY *VaultYSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _VaultY.Contract.GetRoleAdmin(&_VaultY.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VaultY *VaultYCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _VaultY.Contract.GetRoleAdmin(&_VaultY.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_VaultY *VaultYCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_VaultY *VaultYSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _VaultY.Contract.GetRoleMember(&_VaultY.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_VaultY *VaultYCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _VaultY.Contract.GetRoleMember(&_VaultY.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_VaultY *VaultYCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_VaultY *VaultYSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _VaultY.Contract.GetRoleMemberCount(&_VaultY.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_VaultY *VaultYCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _VaultY.Contract.GetRoleMemberCount(&_VaultY.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_VaultY *VaultYCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_VaultY *VaultYSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _VaultY.Contract.GetRoleMembers(&_VaultY.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_VaultY *VaultYCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _VaultY.Contract.GetRoleMembers(&_VaultY.CallOpts, role)
}

// GetRoles is a free data retrieval call binding the contract method 0x71061398.
//
// Solidity: function getRoles() pure returns((bytes32,string)[])
func (_VaultY *VaultYCaller) GetRoles(opts *bind.CallOpts) ([]RoleAccessRole, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "getRoles")

	if err != nil {
		return *new([]RoleAccessRole), err
	}

	out0 := *abi.ConvertType(out[0], new([]RoleAccessRole)).(*[]RoleAccessRole)

	return out0, err

}

// GetRoles is a free data retrieval call binding the contract method 0x71061398.
//
// Solidity: function getRoles() pure returns((bytes32,string)[])
func (_VaultY *VaultYSession) GetRoles() ([]RoleAccessRole, error) {
	return _VaultY.Contract.GetRoles(&_VaultY.CallOpts)
}

// GetRoles is a free data retrieval call binding the contract method 0x71061398.
//
// Solidity: function getRoles() pure returns((bytes32,string)[])
func (_VaultY *VaultYCallerSession) GetRoles() ([]RoleAccessRole, error) {
	return _VaultY.Contract.GetRoles(&_VaultY.CallOpts)
}

// GetTip is a free data retrieval call binding the contract method 0xb58f745b.
//
// Solidity: function getTip(address sourceToken, address mappedToken, uint256 amount) view returns(uint256)
func (_VaultY *VaultYCaller) GetTip(opts *bind.CallOpts, sourceToken common.Address, mappedToken common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "getTip", sourceToken, mappedToken, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTip is a free data retrieval call binding the contract method 0xb58f745b.
//
// Solidity: function getTip(address sourceToken, address mappedToken, uint256 amount) view returns(uint256)
func (_VaultY *VaultYSession) GetTip(sourceToken common.Address, mappedToken common.Address, amount *big.Int) (*big.Int, error) {
	return _VaultY.Contract.GetTip(&_VaultY.CallOpts, sourceToken, mappedToken, amount)
}

// GetTip is a free data retrieval call binding the contract method 0xb58f745b.
//
// Solidity: function getTip(address sourceToken, address mappedToken, uint256 amount) view returns(uint256)
func (_VaultY *VaultYCallerSession) GetTip(sourceToken common.Address, mappedToken common.Address, amount *big.Int) (*big.Int, error) {
	return _VaultY.Contract.GetTip(&_VaultY.CallOpts, sourceToken, mappedToken, amount)
}

// GetTokenPairs is a free data retrieval call binding the contract method 0xe24e4fdb.
//
// Solidity: function getTokenPairs() view returns((address,uint256,string,address,uint256,string,bool)[])
func (_VaultY *VaultYCaller) GetTokenPairs(opts *bind.CallOpts) ([]VaultYtokenPairInfo, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "getTokenPairs")

	if err != nil {
		return *new([]VaultYtokenPairInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]VaultYtokenPairInfo)).(*[]VaultYtokenPairInfo)

	return out0, err

}

// GetTokenPairs is a free data retrieval call binding the contract method 0xe24e4fdb.
//
// Solidity: function getTokenPairs() view returns((address,uint256,string,address,uint256,string,bool)[])
func (_VaultY *VaultYSession) GetTokenPairs() ([]VaultYtokenPairInfo, error) {
	return _VaultY.Contract.GetTokenPairs(&_VaultY.CallOpts)
}

// GetTokenPairs is a free data retrieval call binding the contract method 0xe24e4fdb.
//
// Solidity: function getTokenPairs() view returns((address,uint256,string,address,uint256,string,bool)[])
func (_VaultY *VaultYCallerSession) GetTokenPairs() ([]VaultYtokenPairInfo, error) {
	return _VaultY.Contract.GetTokenPairs(&_VaultY.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VaultY *VaultYCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VaultY *VaultYSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _VaultY.Contract.HasRole(&_VaultY.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VaultY *VaultYCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _VaultY.Contract.HasRole(&_VaultY.CallOpts, role, account)
}

// OmitNonces is a free data retrieval call binding the contract method 0x1bcd5a4e.
//
// Solidity: function omitNonces(uint256 ) view returns(bool)
func (_VaultY *VaultYCaller) OmitNonces(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "omitNonces", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OmitNonces is a free data retrieval call binding the contract method 0x1bcd5a4e.
//
// Solidity: function omitNonces(uint256 ) view returns(bool)
func (_VaultY *VaultYSession) OmitNonces(arg0 *big.Int) (bool, error) {
	return _VaultY.Contract.OmitNonces(&_VaultY.CallOpts, arg0)
}

// OmitNonces is a free data retrieval call binding the contract method 0x1bcd5a4e.
//
// Solidity: function omitNonces(uint256 ) view returns(bool)
func (_VaultY *VaultYCallerSession) OmitNonces(arg0 *big.Int) (bool, error) {
	return _VaultY.Contract.OmitNonces(&_VaultY.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VaultY *VaultYCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VaultY *VaultYSession) Paused() (bool, error) {
	return _VaultY.Contract.Paused(&_VaultY.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VaultY *VaultYCallerSession) Paused() (bool, error) {
	return _VaultY.Contract.Paused(&_VaultY.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VaultY *VaultYCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VaultY *VaultYSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VaultY.Contract.SupportsInterface(&_VaultY.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VaultY *VaultYCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VaultY.Contract.SupportsInterface(&_VaultY.CallOpts, interfaceId)
}

// TipBalance is a free data retrieval call binding the contract method 0xefc03425.
//
// Solidity: function tipBalance(address token) view returns(uint256)
func (_VaultY *VaultYCaller) TipBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "tipBalance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TipBalance is a free data retrieval call binding the contract method 0xefc03425.
//
// Solidity: function tipBalance(address token) view returns(uint256)
func (_VaultY *VaultYSession) TipBalance(token common.Address) (*big.Int, error) {
	return _VaultY.Contract.TipBalance(&_VaultY.CallOpts, token)
}

// TipBalance is a free data retrieval call binding the contract method 0xefc03425.
//
// Solidity: function tipBalance(address token) view returns(uint256)
func (_VaultY *VaultYCallerSession) TipBalance(token common.Address) (*big.Int, error) {
	return _VaultY.Contract.TipBalance(&_VaultY.CallOpts, token)
}

// TokenMappingWatermark is a free data retrieval call binding the contract method 0x55275c6f.
//
// Solidity: function tokenMappingWatermark(address , address ) view returns(uint256)
func (_VaultY *VaultYCaller) TokenMappingWatermark(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "tokenMappingWatermark", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMappingWatermark is a free data retrieval call binding the contract method 0x55275c6f.
//
// Solidity: function tokenMappingWatermark(address , address ) view returns(uint256)
func (_VaultY *VaultYSession) TokenMappingWatermark(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VaultY.Contract.TokenMappingWatermark(&_VaultY.CallOpts, arg0, arg1)
}

// TokenMappingWatermark is a free data retrieval call binding the contract method 0x55275c6f.
//
// Solidity: function tokenMappingWatermark(address , address ) view returns(uint256)
func (_VaultY *VaultYCallerSession) TokenMappingWatermark(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VaultY.Contract.TokenMappingWatermark(&_VaultY.CallOpts, arg0, arg1)
}

// TokenPairs is a free data retrieval call binding the contract method 0xa15369d1.
//
// Solidity: function tokenPairs(uint256 ) view returns(address sourceToken, address mappedToken)
func (_VaultY *VaultYCaller) TokenPairs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SourceToken common.Address
	MappedToken common.Address
}, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "tokenPairs", arg0)

	outstruct := new(struct {
		SourceToken common.Address
		MappedToken common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SourceToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.MappedToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// TokenPairs is a free data retrieval call binding the contract method 0xa15369d1.
//
// Solidity: function tokenPairs(uint256 ) view returns(address sourceToken, address mappedToken)
func (_VaultY *VaultYSession) TokenPairs(arg0 *big.Int) (struct {
	SourceToken common.Address
	MappedToken common.Address
}, error) {
	return _VaultY.Contract.TokenPairs(&_VaultY.CallOpts, arg0)
}

// TokenPairs is a free data retrieval call binding the contract method 0xa15369d1.
//
// Solidity: function tokenPairs(uint256 ) view returns(address sourceToken, address mappedToken)
func (_VaultY *VaultYCallerSession) TokenPairs(arg0 *big.Int) (struct {
	SourceToken common.Address
	MappedToken common.Address
}, error) {
	return _VaultY.Contract.TokenPairs(&_VaultY.CallOpts, arg0)
}

// AddNoncesToOmit is a paid mutator transaction binding the contract method 0x379e271e.
//
// Solidity: function addNoncesToOmit(uint256[] nonces) returns()
func (_VaultY *VaultYTransactor) AddNoncesToOmit(opts *bind.TransactOpts, nonces []*big.Int) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "addNoncesToOmit", nonces)
}

// AddNoncesToOmit is a paid mutator transaction binding the contract method 0x379e271e.
//
// Solidity: function addNoncesToOmit(uint256[] nonces) returns()
func (_VaultY *VaultYSession) AddNoncesToOmit(nonces []*big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.AddNoncesToOmit(&_VaultY.TransactOpts, nonces)
}

// AddNoncesToOmit is a paid mutator transaction binding the contract method 0x379e271e.
//
// Solidity: function addNoncesToOmit(uint256[] nonces) returns()
func (_VaultY *VaultYTransactorSession) AddNoncesToOmit(nonces []*big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.AddNoncesToOmit(&_VaultY.TransactOpts, nonces)
}

// AddRoleMember is a paid mutator transaction binding the contract method 0x1b65471f.
//
// Solidity: function addRoleMember(bytes32 role, address member) returns(bool)
func (_VaultY *VaultYTransactor) AddRoleMember(opts *bind.TransactOpts, role [32]byte, member common.Address) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "addRoleMember", role, member)
}

// AddRoleMember is a paid mutator transaction binding the contract method 0x1b65471f.
//
// Solidity: function addRoleMember(bytes32 role, address member) returns(bool)
func (_VaultY *VaultYSession) AddRoleMember(role [32]byte, member common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.AddRoleMember(&_VaultY.TransactOpts, role, member)
}

// AddRoleMember is a paid mutator transaction binding the contract method 0x1b65471f.
//
// Solidity: function addRoleMember(bytes32 role, address member) returns(bool)
func (_VaultY *VaultYTransactorSession) AddRoleMember(role [32]byte, member common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.AddRoleMember(&_VaultY.TransactOpts, role, member)
}

// BatchMint is a paid mutator transaction binding the contract method 0x21b97da9.
//
// Solidity: function batchMint(bytes signature, bytes input) returns()
func (_VaultY *VaultYTransactor) BatchMint(opts *bind.TransactOpts, signature []byte, input []byte) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "batchMint", signature, input)
}

// BatchMint is a paid mutator transaction binding the contract method 0x21b97da9.
//
// Solidity: function batchMint(bytes signature, bytes input) returns()
func (_VaultY *VaultYSession) BatchMint(signature []byte, input []byte) (*types.Transaction, error) {
	return _VaultY.Contract.BatchMint(&_VaultY.TransactOpts, signature, input)
}

// BatchMint is a paid mutator transaction binding the contract method 0x21b97da9.
//
// Solidity: function batchMint(bytes signature, bytes input) returns()
func (_VaultY *VaultYTransactorSession) BatchMint(signature []byte, input []byte) (*types.Transaction, error) {
	return _VaultY.Contract.BatchMint(&_VaultY.TransactOpts, signature, input)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address mappedToken, uint256 amount) returns()
func (_VaultY *VaultYTransactor) Burn(opts *bind.TransactOpts, mappedToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "burn", mappedToken, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address mappedToken, uint256 amount) returns()
func (_VaultY *VaultYSession) Burn(mappedToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.Burn(&_VaultY.TransactOpts, mappedToken, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address mappedToken, uint256 amount) returns()
func (_VaultY *VaultYTransactorSession) Burn(mappedToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.Burn(&_VaultY.TransactOpts, mappedToken, amount)
}

// GrantMinter is a paid mutator transaction binding the contract method 0x261707fa.
//
// Solidity: function grantMinter(address minter) returns(bool)
func (_VaultY *VaultYTransactor) GrantMinter(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "grantMinter", minter)
}

// GrantMinter is a paid mutator transaction binding the contract method 0x261707fa.
//
// Solidity: function grantMinter(address minter) returns(bool)
func (_VaultY *VaultYSession) GrantMinter(minter common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.GrantMinter(&_VaultY.TransactOpts, minter)
}

// GrantMinter is a paid mutator transaction binding the contract method 0x261707fa.
//
// Solidity: function grantMinter(address minter) returns(bool)
func (_VaultY *VaultYTransactorSession) GrantMinter(minter common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.GrantMinter(&_VaultY.TransactOpts, minter)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VaultY *VaultYTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VaultY *VaultYSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.GrantRole(&_VaultY.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VaultY *VaultYTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.GrantRole(&_VaultY.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0x3e401bf8.
//
// Solidity: function mint(address sourceToken, address mappedToken, address to, uint256 amount, uint256 tipX, uint256 nonce) returns()
func (_VaultY *VaultYTransactor) Mint(opts *bind.TransactOpts, sourceToken common.Address, mappedToken common.Address, to common.Address, amount *big.Int, tipX *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "mint", sourceToken, mappedToken, to, amount, tipX, nonce)
}

// Mint is a paid mutator transaction binding the contract method 0x3e401bf8.
//
// Solidity: function mint(address sourceToken, address mappedToken, address to, uint256 amount, uint256 tipX, uint256 nonce) returns()
func (_VaultY *VaultYSession) Mint(sourceToken common.Address, mappedToken common.Address, to common.Address, amount *big.Int, tipX *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.Mint(&_VaultY.TransactOpts, sourceToken, mappedToken, to, amount, tipX, nonce)
}

// Mint is a paid mutator transaction binding the contract method 0x3e401bf8.
//
// Solidity: function mint(address sourceToken, address mappedToken, address to, uint256 amount, uint256 tipX, uint256 nonce) returns()
func (_VaultY *VaultYTransactorSession) Mint(sourceToken common.Address, mappedToken common.Address, to common.Address, amount *big.Int, tipX *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.Mint(&_VaultY.TransactOpts, sourceToken, mappedToken, to, amount, tipX, nonce)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_VaultY *VaultYTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_VaultY *VaultYSession) PauseAll() (*types.Transaction, error) {
	return _VaultY.Contract.PauseAll(&_VaultY.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_VaultY *VaultYTransactorSession) PauseAll() (*types.Transaction, error) {
	return _VaultY.Contract.PauseAll(&_VaultY.TransactOpts)
}

// PauseTokenMapping is a paid mutator transaction binding the contract method 0x0020d792.
//
// Solidity: function pauseTokenMapping(address sourceToken, address mappedToken) returns()
func (_VaultY *VaultYTransactor) PauseTokenMapping(opts *bind.TransactOpts, sourceToken common.Address, mappedToken common.Address) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "pauseTokenMapping", sourceToken, mappedToken)
}

// PauseTokenMapping is a paid mutator transaction binding the contract method 0x0020d792.
//
// Solidity: function pauseTokenMapping(address sourceToken, address mappedToken) returns()
func (_VaultY *VaultYSession) PauseTokenMapping(sourceToken common.Address, mappedToken common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.PauseTokenMapping(&_VaultY.TransactOpts, sourceToken, mappedToken)
}

// PauseTokenMapping is a paid mutator transaction binding the contract method 0x0020d792.
//
// Solidity: function pauseTokenMapping(address sourceToken, address mappedToken) returns()
func (_VaultY *VaultYTransactorSession) PauseTokenMapping(sourceToken common.Address, mappedToken common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.PauseTokenMapping(&_VaultY.TransactOpts, sourceToken, mappedToken)
}

// Refund is a paid mutator transaction binding the contract method 0x82ad6f35.
//
// Solidity: function refund(address token, address to, uint256 amount) returns()
func (_VaultY *VaultYTransactor) Refund(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "refund", token, to, amount)
}

// Refund is a paid mutator transaction binding the contract method 0x82ad6f35.
//
// Solidity: function refund(address token, address to, uint256 amount) returns()
func (_VaultY *VaultYSession) Refund(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.Refund(&_VaultY.TransactOpts, token, to, amount)
}

// Refund is a paid mutator transaction binding the contract method 0x82ad6f35.
//
// Solidity: function refund(address token, address to, uint256 amount) returns()
func (_VaultY *VaultYTransactorSession) Refund(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.Refund(&_VaultY.TransactOpts, token, to, amount)
}

// RemoveNoncesToOmit is a paid mutator transaction binding the contract method 0xcb79b053.
//
// Solidity: function removeNoncesToOmit(uint256[] nonces) returns()
func (_VaultY *VaultYTransactor) RemoveNoncesToOmit(opts *bind.TransactOpts, nonces []*big.Int) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "removeNoncesToOmit", nonces)
}

// RemoveNoncesToOmit is a paid mutator transaction binding the contract method 0xcb79b053.
//
// Solidity: function removeNoncesToOmit(uint256[] nonces) returns()
func (_VaultY *VaultYSession) RemoveNoncesToOmit(nonces []*big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.RemoveNoncesToOmit(&_VaultY.TransactOpts, nonces)
}

// RemoveNoncesToOmit is a paid mutator transaction binding the contract method 0xcb79b053.
//
// Solidity: function removeNoncesToOmit(uint256[] nonces) returns()
func (_VaultY *VaultYTransactorSession) RemoveNoncesToOmit(nonces []*big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.RemoveNoncesToOmit(&_VaultY.TransactOpts, nonces)
}

// RemoveRoleMember is a paid mutator transaction binding the contract method 0x4dd8fac8.
//
// Solidity: function removeRoleMember(bytes32 role, address member) returns(bool)
func (_VaultY *VaultYTransactor) RemoveRoleMember(opts *bind.TransactOpts, role [32]byte, member common.Address) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "removeRoleMember", role, member)
}

// RemoveRoleMember is a paid mutator transaction binding the contract method 0x4dd8fac8.
//
// Solidity: function removeRoleMember(bytes32 role, address member) returns(bool)
func (_VaultY *VaultYSession) RemoveRoleMember(role [32]byte, member common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.RemoveRoleMember(&_VaultY.TransactOpts, role, member)
}

// RemoveRoleMember is a paid mutator transaction binding the contract method 0x4dd8fac8.
//
// Solidity: function removeRoleMember(bytes32 role, address member) returns(bool)
func (_VaultY *VaultYTransactorSession) RemoveRoleMember(role [32]byte, member common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.RemoveRoleMember(&_VaultY.TransactOpts, role, member)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_VaultY *VaultYTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_VaultY *VaultYSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.RenounceRole(&_VaultY.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_VaultY *VaultYTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.RenounceRole(&_VaultY.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VaultY *VaultYTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VaultY *VaultYSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.RevokeRole(&_VaultY.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VaultY *VaultYTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.RevokeRole(&_VaultY.TransactOpts, role, account)
}

// SetFeeAccount is a paid mutator transaction binding the contract method 0x4b023cf8.
//
// Solidity: function setFeeAccount(address newFeeAccount) returns()
func (_VaultY *VaultYTransactor) SetFeeAccount(opts *bind.TransactOpts, newFeeAccount common.Address) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "setFeeAccount", newFeeAccount)
}

// SetFeeAccount is a paid mutator transaction binding the contract method 0x4b023cf8.
//
// Solidity: function setFeeAccount(address newFeeAccount) returns()
func (_VaultY *VaultYSession) SetFeeAccount(newFeeAccount common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.SetFeeAccount(&_VaultY.TransactOpts, newFeeAccount)
}

// SetFeeAccount is a paid mutator transaction binding the contract method 0x4b023cf8.
//
// Solidity: function setFeeAccount(address newFeeAccount) returns()
func (_VaultY *VaultYTransactorSession) SetFeeAccount(newFeeAccount common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.SetFeeAccount(&_VaultY.TransactOpts, newFeeAccount)
}

// SetFiatCurrency is a paid mutator transaction binding the contract method 0x10804064.
//
// Solidity: function setFiatCurrency(string newFiatCurrency) returns()
func (_VaultY *VaultYTransactor) SetFiatCurrency(opts *bind.TransactOpts, newFiatCurrency string) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "setFiatCurrency", newFiatCurrency)
}

// SetFiatCurrency is a paid mutator transaction binding the contract method 0x10804064.
//
// Solidity: function setFiatCurrency(string newFiatCurrency) returns()
func (_VaultY *VaultYSession) SetFiatCurrency(newFiatCurrency string) (*types.Transaction, error) {
	return _VaultY.Contract.SetFiatCurrency(&_VaultY.TransactOpts, newFiatCurrency)
}

// SetFiatCurrency is a paid mutator transaction binding the contract method 0x10804064.
//
// Solidity: function setFiatCurrency(string newFiatCurrency) returns()
func (_VaultY *VaultYTransactorSession) SetFiatCurrency(newFiatCurrency string) (*types.Transaction, error) {
	return _VaultY.Contract.SetFiatCurrency(&_VaultY.TransactOpts, newFiatCurrency)
}

// SetFiatFeeAmount is a paid mutator transaction binding the contract method 0xb7fac573.
//
// Solidity: function setFiatFeeAmount(uint256 amount) returns()
func (_VaultY *VaultYTransactor) SetFiatFeeAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "setFiatFeeAmount", amount)
}

// SetFiatFeeAmount is a paid mutator transaction binding the contract method 0xb7fac573.
//
// Solidity: function setFiatFeeAmount(uint256 amount) returns()
func (_VaultY *VaultYSession) SetFiatFeeAmount(amount *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.SetFiatFeeAmount(&_VaultY.TransactOpts, amount)
}

// SetFiatFeeAmount is a paid mutator transaction binding the contract method 0xb7fac573.
//
// Solidity: function setFiatFeeAmount(uint256 amount) returns()
func (_VaultY *VaultYTransactorSession) SetFiatFeeAmount(amount *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.SetFiatFeeAmount(&_VaultY.TransactOpts, amount)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address newPriceoracle) returns()
func (_VaultY *VaultYTransactor) SetPriceOracle(opts *bind.TransactOpts, newPriceoracle common.Address) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "setPriceOracle", newPriceoracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address newPriceoracle) returns()
func (_VaultY *VaultYSession) SetPriceOracle(newPriceoracle common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.SetPriceOracle(&_VaultY.TransactOpts, newPriceoracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address newPriceoracle) returns()
func (_VaultY *VaultYTransactorSession) SetPriceOracle(newPriceoracle common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.SetPriceOracle(&_VaultY.TransactOpts, newPriceoracle)
}

// SetTipAccount is a paid mutator transaction binding the contract method 0x0cddaa22.
//
// Solidity: function setTipAccount(address newTipAccount) returns()
func (_VaultY *VaultYTransactor) SetTipAccount(opts *bind.TransactOpts, newTipAccount common.Address) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "setTipAccount", newTipAccount)
}

// SetTipAccount is a paid mutator transaction binding the contract method 0x0cddaa22.
//
// Solidity: function setTipAccount(address newTipAccount) returns()
func (_VaultY *VaultYSession) SetTipAccount(newTipAccount common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.SetTipAccount(&_VaultY.TransactOpts, newTipAccount)
}

// SetTipAccount is a paid mutator transaction binding the contract method 0x0cddaa22.
//
// Solidity: function setTipAccount(address newTipAccount) returns()
func (_VaultY *VaultYTransactorSession) SetTipAccount(newTipAccount common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.SetTipAccount(&_VaultY.TransactOpts, newTipAccount)
}

// SetTipRate is a paid mutator transaction binding the contract method 0xda78defd.
//
// Solidity: function setTipRate(address sourceToken, address mappedToken, uint256 newTipRate) returns()
func (_VaultY *VaultYTransactor) SetTipRate(opts *bind.TransactOpts, sourceToken common.Address, mappedToken common.Address, newTipRate *big.Int) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "setTipRate", sourceToken, mappedToken, newTipRate)
}

// SetTipRate is a paid mutator transaction binding the contract method 0xda78defd.
//
// Solidity: function setTipRate(address sourceToken, address mappedToken, uint256 newTipRate) returns()
func (_VaultY *VaultYSession) SetTipRate(sourceToken common.Address, mappedToken common.Address, newTipRate *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.SetTipRate(&_VaultY.TransactOpts, sourceToken, mappedToken, newTipRate)
}

// SetTipRate is a paid mutator transaction binding the contract method 0xda78defd.
//
// Solidity: function setTipRate(address sourceToken, address mappedToken, uint256 newTipRate) returns()
func (_VaultY *VaultYTransactorSession) SetTipRate(sourceToken common.Address, mappedToken common.Address, newTipRate *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.SetTipRate(&_VaultY.TransactOpts, sourceToken, mappedToken, newTipRate)
}

// SetupTokenMapping is a paid mutator transaction binding the contract method 0x9c0d9230.
//
// Solidity: function setupTokenMapping(uint256 sourceChainid, address sourceToken, address mappedToken, string sourceTokenSymbol_, string mappedTokenSymbol_, uint256 tipRate) returns(bool)
func (_VaultY *VaultYTransactor) SetupTokenMapping(opts *bind.TransactOpts, sourceChainid *big.Int, sourceToken common.Address, mappedToken common.Address, sourceTokenSymbol_ string, mappedTokenSymbol_ string, tipRate *big.Int) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "setupTokenMapping", sourceChainid, sourceToken, mappedToken, sourceTokenSymbol_, mappedTokenSymbol_, tipRate)
}

// SetupTokenMapping is a paid mutator transaction binding the contract method 0x9c0d9230.
//
// Solidity: function setupTokenMapping(uint256 sourceChainid, address sourceToken, address mappedToken, string sourceTokenSymbol_, string mappedTokenSymbol_, uint256 tipRate) returns(bool)
func (_VaultY *VaultYSession) SetupTokenMapping(sourceChainid *big.Int, sourceToken common.Address, mappedToken common.Address, sourceTokenSymbol_ string, mappedTokenSymbol_ string, tipRate *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.SetupTokenMapping(&_VaultY.TransactOpts, sourceChainid, sourceToken, mappedToken, sourceTokenSymbol_, mappedTokenSymbol_, tipRate)
}

// SetupTokenMapping is a paid mutator transaction binding the contract method 0x9c0d9230.
//
// Solidity: function setupTokenMapping(uint256 sourceChainid, address sourceToken, address mappedToken, string sourceTokenSymbol_, string mappedTokenSymbol_, uint256 tipRate) returns(bool)
func (_VaultY *VaultYTransactorSession) SetupTokenMapping(sourceChainid *big.Int, sourceToken common.Address, mappedToken common.Address, sourceTokenSymbol_ string, mappedTokenSymbol_ string, tipRate *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.SetupTokenMapping(&_VaultY.TransactOpts, sourceChainid, sourceToken, mappedToken, sourceTokenSymbol_, mappedTokenSymbol_, tipRate)
}

// SkipMintWatermark is a paid mutator transaction binding the contract method 0xec232e86.
//
// Solidity: function skipMintWatermark(address sourceToken, address mappedToken, uint256 skip) returns()
func (_VaultY *VaultYTransactor) SkipMintWatermark(opts *bind.TransactOpts, sourceToken common.Address, mappedToken common.Address, skip *big.Int) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "skipMintWatermark", sourceToken, mappedToken, skip)
}

// SkipMintWatermark is a paid mutator transaction binding the contract method 0xec232e86.
//
// Solidity: function skipMintWatermark(address sourceToken, address mappedToken, uint256 skip) returns()
func (_VaultY *VaultYSession) SkipMintWatermark(sourceToken common.Address, mappedToken common.Address, skip *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.SkipMintWatermark(&_VaultY.TransactOpts, sourceToken, mappedToken, skip)
}

// SkipMintWatermark is a paid mutator transaction binding the contract method 0xec232e86.
//
// Solidity: function skipMintWatermark(address sourceToken, address mappedToken, uint256 skip) returns()
func (_VaultY *VaultYTransactorSession) SkipMintWatermark(sourceToken common.Address, mappedToken common.Address, skip *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.SkipMintWatermark(&_VaultY.TransactOpts, sourceToken, mappedToken, skip)
}

// TipCashout is a paid mutator transaction binding the contract method 0x1a46b010.
//
// Solidity: function tipCashout(address token, address to, uint256 amount) returns()
func (_VaultY *VaultYTransactor) TipCashout(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "tipCashout", token, to, amount)
}

// TipCashout is a paid mutator transaction binding the contract method 0x1a46b010.
//
// Solidity: function tipCashout(address token, address to, uint256 amount) returns()
func (_VaultY *VaultYSession) TipCashout(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.TipCashout(&_VaultY.TransactOpts, token, to, amount)
}

// TipCashout is a paid mutator transaction binding the contract method 0x1a46b010.
//
// Solidity: function tipCashout(address token, address to, uint256 amount) returns()
func (_VaultY *VaultYTransactorSession) TipCashout(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.TipCashout(&_VaultY.TransactOpts, token, to, amount)
}

// UnpauseAll is a paid mutator transaction binding the contract method 0x8a2ddd03.
//
// Solidity: function unpauseAll() returns()
func (_VaultY *VaultYTransactor) UnpauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "unpauseAll")
}

// UnpauseAll is a paid mutator transaction binding the contract method 0x8a2ddd03.
//
// Solidity: function unpauseAll() returns()
func (_VaultY *VaultYSession) UnpauseAll() (*types.Transaction, error) {
	return _VaultY.Contract.UnpauseAll(&_VaultY.TransactOpts)
}

// UnpauseAll is a paid mutator transaction binding the contract method 0x8a2ddd03.
//
// Solidity: function unpauseAll() returns()
func (_VaultY *VaultYTransactorSession) UnpauseAll() (*types.Transaction, error) {
	return _VaultY.Contract.UnpauseAll(&_VaultY.TransactOpts)
}

// UnpauseTokenMapping is a paid mutator transaction binding the contract method 0x7b32a34d.
//
// Solidity: function unpauseTokenMapping(address sourceToken, address mappedToken) returns()
func (_VaultY *VaultYTransactor) UnpauseTokenMapping(opts *bind.TransactOpts, sourceToken common.Address, mappedToken common.Address) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "unpauseTokenMapping", sourceToken, mappedToken)
}

// UnpauseTokenMapping is a paid mutator transaction binding the contract method 0x7b32a34d.
//
// Solidity: function unpauseTokenMapping(address sourceToken, address mappedToken) returns()
func (_VaultY *VaultYSession) UnpauseTokenMapping(sourceToken common.Address, mappedToken common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.UnpauseTokenMapping(&_VaultY.TransactOpts, sourceToken, mappedToken)
}

// UnpauseTokenMapping is a paid mutator transaction binding the contract method 0x7b32a34d.
//
// Solidity: function unpauseTokenMapping(address sourceToken, address mappedToken) returns()
func (_VaultY *VaultYTransactorSession) UnpauseTokenMapping(sourceToken common.Address, mappedToken common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.UnpauseTokenMapping(&_VaultY.TransactOpts, sourceToken, mappedToken)
}

// Withdraw is a paid mutator transaction binding the contract method 0xac6f7e3b.
//
// Solidity: function withdraw(address sourceToken, address mappedToken, address to, uint256 amount, uint256 tipY, uint256 nonce) returns()
func (_VaultY *VaultYTransactor) Withdraw(opts *bind.TransactOpts, sourceToken common.Address, mappedToken common.Address, to common.Address, amount *big.Int, tipY *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "withdraw", sourceToken, mappedToken, to, amount, tipY, nonce)
}

// Withdraw is a paid mutator transaction binding the contract method 0xac6f7e3b.
//
// Solidity: function withdraw(address sourceToken, address mappedToken, address to, uint256 amount, uint256 tipY, uint256 nonce) returns()
func (_VaultY *VaultYSession) Withdraw(sourceToken common.Address, mappedToken common.Address, to common.Address, amount *big.Int, tipY *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.Withdraw(&_VaultY.TransactOpts, sourceToken, mappedToken, to, amount, tipY, nonce)
}

// Withdraw is a paid mutator transaction binding the contract method 0xac6f7e3b.
//
// Solidity: function withdraw(address sourceToken, address mappedToken, address to, uint256 amount, uint256 tipY, uint256 nonce) returns()
func (_VaultY *VaultYTransactorSession) Withdraw(sourceToken common.Address, mappedToken common.Address, to common.Address, amount *big.Int, tipY *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.Withdraw(&_VaultY.TransactOpts, sourceToken, mappedToken, to, amount, tipY, nonce)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_VaultY *VaultYTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _VaultY.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_VaultY *VaultYSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _VaultY.Contract.Fallback(&_VaultY.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_VaultY *VaultYTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _VaultY.Contract.Fallback(&_VaultY.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VaultY *VaultYTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultY.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VaultY *VaultYSession) Receive() (*types.Transaction, error) {
	return _VaultY.Contract.Receive(&_VaultY.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VaultY *VaultYTransactorSession) Receive() (*types.Transaction, error) {
	return _VaultY.Contract.Receive(&_VaultY.TransactOpts)
}

// VaultYFeeAccountChangedIterator is returned from FilterFeeAccountChanged and is used to iterate over the raw logs and unpacked data for FeeAccountChanged events raised by the VaultY contract.
type VaultYFeeAccountChangedIterator struct {
	Event *VaultYFeeAccountChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultYFeeAccountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultYFeeAccountChanged)
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
		it.Event = new(VaultYFeeAccountChanged)
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
func (it *VaultYFeeAccountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultYFeeAccountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultYFeeAccountChanged represents a FeeAccountChanged event raised by the VaultY contract.
type VaultYFeeAccountChanged struct {
	Sender        common.Address
	NewFeeAccount common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeeAccountChanged is a free log retrieval operation binding the contract event 0x508dcd584678b460e90a73b67f0c71c3c9eb2506cc1100a01ec04f6a19fac648.
//
// Solidity: event FeeAccountChanged(address indexed sender, address newFeeAccount)
func (_VaultY *VaultYFilterer) FilterFeeAccountChanged(opts *bind.FilterOpts, sender []common.Address) (*VaultYFeeAccountChangedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultY.contract.FilterLogs(opts, "FeeAccountChanged", senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultYFeeAccountChangedIterator{contract: _VaultY.contract, event: "FeeAccountChanged", logs: logs, sub: sub}, nil
}

// WatchFeeAccountChanged is a free log subscription operation binding the contract event 0x508dcd584678b460e90a73b67f0c71c3c9eb2506cc1100a01ec04f6a19fac648.
//
// Solidity: event FeeAccountChanged(address indexed sender, address newFeeAccount)
func (_VaultY *VaultYFilterer) WatchFeeAccountChanged(opts *bind.WatchOpts, sink chan<- *VaultYFeeAccountChanged, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultY.contract.WatchLogs(opts, "FeeAccountChanged", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultYFeeAccountChanged)
				if err := _VaultY.contract.UnpackLog(event, "FeeAccountChanged", log); err != nil {
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

// ParseFeeAccountChanged is a log parse operation binding the contract event 0x508dcd584678b460e90a73b67f0c71c3c9eb2506cc1100a01ec04f6a19fac648.
//
// Solidity: event FeeAccountChanged(address indexed sender, address newFeeAccount)
func (_VaultY *VaultYFilterer) ParseFeeAccountChanged(log types.Log) (*VaultYFeeAccountChanged, error) {
	event := new(VaultYFeeAccountChanged)
	if err := _VaultY.contract.UnpackLog(event, "FeeAccountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultYFiatCurrencyChangedIterator is returned from FilterFiatCurrencyChanged and is used to iterate over the raw logs and unpacked data for FiatCurrencyChanged events raised by the VaultY contract.
type VaultYFiatCurrencyChangedIterator struct {
	Event *VaultYFiatCurrencyChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultYFiatCurrencyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultYFiatCurrencyChanged)
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
		it.Event = new(VaultYFiatCurrencyChanged)
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
func (it *VaultYFiatCurrencyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultYFiatCurrencyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultYFiatCurrencyChanged represents a FiatCurrencyChanged event raised by the VaultY contract.
type VaultYFiatCurrencyChanged struct {
	Sender          common.Address
	NewFiatCurrency string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFiatCurrencyChanged is a free log retrieval operation binding the contract event 0x7ed7441c7c9c83f1fd4152cec0a268d64e2c8784a2a05edf810befb09eb2a328.
//
// Solidity: event FiatCurrencyChanged(address indexed sender, string newFiatCurrency)
func (_VaultY *VaultYFilterer) FilterFiatCurrencyChanged(opts *bind.FilterOpts, sender []common.Address) (*VaultYFiatCurrencyChangedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultY.contract.FilterLogs(opts, "FiatCurrencyChanged", senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultYFiatCurrencyChangedIterator{contract: _VaultY.contract, event: "FiatCurrencyChanged", logs: logs, sub: sub}, nil
}

// WatchFiatCurrencyChanged is a free log subscription operation binding the contract event 0x7ed7441c7c9c83f1fd4152cec0a268d64e2c8784a2a05edf810befb09eb2a328.
//
// Solidity: event FiatCurrencyChanged(address indexed sender, string newFiatCurrency)
func (_VaultY *VaultYFilterer) WatchFiatCurrencyChanged(opts *bind.WatchOpts, sink chan<- *VaultYFiatCurrencyChanged, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultY.contract.WatchLogs(opts, "FiatCurrencyChanged", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultYFiatCurrencyChanged)
				if err := _VaultY.contract.UnpackLog(event, "FiatCurrencyChanged", log); err != nil {
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

// ParseFiatCurrencyChanged is a log parse operation binding the contract event 0x7ed7441c7c9c83f1fd4152cec0a268d64e2c8784a2a05edf810befb09eb2a328.
//
// Solidity: event FiatCurrencyChanged(address indexed sender, string newFiatCurrency)
func (_VaultY *VaultYFilterer) ParseFiatCurrencyChanged(log types.Log) (*VaultYFiatCurrencyChanged, error) {
	event := new(VaultYFiatCurrencyChanged)
	if err := _VaultY.contract.UnpackLog(event, "FiatCurrencyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultYFiatFeeAmountChangedIterator is returned from FilterFiatFeeAmountChanged and is used to iterate over the raw logs and unpacked data for FiatFeeAmountChanged events raised by the VaultY contract.
type VaultYFiatFeeAmountChangedIterator struct {
	Event *VaultYFiatFeeAmountChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultYFiatFeeAmountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultYFiatFeeAmountChanged)
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
		it.Event = new(VaultYFiatFeeAmountChanged)
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
func (it *VaultYFiatFeeAmountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultYFiatFeeAmountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultYFiatFeeAmountChanged represents a FiatFeeAmountChanged event raised by the VaultY contract.
type VaultYFiatFeeAmountChanged struct {
	Sender           common.Address
	NewFiatFeeAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFiatFeeAmountChanged is a free log retrieval operation binding the contract event 0x2460d22a2a151edef78f3e2e7bf595d43fabf127aea7b47450f81fafef9d515b.
//
// Solidity: event FiatFeeAmountChanged(address indexed sender, uint256 newFiatFeeAmount)
func (_VaultY *VaultYFilterer) FilterFiatFeeAmountChanged(opts *bind.FilterOpts, sender []common.Address) (*VaultYFiatFeeAmountChangedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultY.contract.FilterLogs(opts, "FiatFeeAmountChanged", senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultYFiatFeeAmountChangedIterator{contract: _VaultY.contract, event: "FiatFeeAmountChanged", logs: logs, sub: sub}, nil
}

// WatchFiatFeeAmountChanged is a free log subscription operation binding the contract event 0x2460d22a2a151edef78f3e2e7bf595d43fabf127aea7b47450f81fafef9d515b.
//
// Solidity: event FiatFeeAmountChanged(address indexed sender, uint256 newFiatFeeAmount)
func (_VaultY *VaultYFilterer) WatchFiatFeeAmountChanged(opts *bind.WatchOpts, sink chan<- *VaultYFiatFeeAmountChanged, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultY.contract.WatchLogs(opts, "FiatFeeAmountChanged", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultYFiatFeeAmountChanged)
				if err := _VaultY.contract.UnpackLog(event, "FiatFeeAmountChanged", log); err != nil {
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

// ParseFiatFeeAmountChanged is a log parse operation binding the contract event 0x2460d22a2a151edef78f3e2e7bf595d43fabf127aea7b47450f81fafef9d515b.
//
// Solidity: event FiatFeeAmountChanged(address indexed sender, uint256 newFiatFeeAmount)
func (_VaultY *VaultYFilterer) ParseFiatFeeAmountChanged(log types.Log) (*VaultYFiatFeeAmountChanged, error) {
	event := new(VaultYFiatFeeAmountChanged)
	if err := _VaultY.contract.UnpackLog(event, "FiatFeeAmountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultYIgnoreNoncesIterator is returned from FilterIgnoreNonces and is used to iterate over the raw logs and unpacked data for IgnoreNonces events raised by the VaultY contract.
type VaultYIgnoreNoncesIterator struct {
	Event *VaultYIgnoreNonces // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultYIgnoreNoncesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultYIgnoreNonces)
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
		it.Event = new(VaultYIgnoreNonces)
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
func (it *VaultYIgnoreNoncesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultYIgnoreNoncesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultYIgnoreNonces represents a IgnoreNonces event raised by the VaultY contract.
type VaultYIgnoreNonces struct {
	Sender common.Address
	Nonces []*big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIgnoreNonces is a free log retrieval operation binding the contract event 0x33fe1cf6879c44bf1b70173d9ba045fa144fdcaafc5135a64285cd6e8b18989c.
//
// Solidity: event IgnoreNonces(address sender, uint256[] nonces)
func (_VaultY *VaultYFilterer) FilterIgnoreNonces(opts *bind.FilterOpts) (*VaultYIgnoreNoncesIterator, error) {

	logs, sub, err := _VaultY.contract.FilterLogs(opts, "IgnoreNonces")
	if err != nil {
		return nil, err
	}
	return &VaultYIgnoreNoncesIterator{contract: _VaultY.contract, event: "IgnoreNonces", logs: logs, sub: sub}, nil
}

// WatchIgnoreNonces is a free log subscription operation binding the contract event 0x33fe1cf6879c44bf1b70173d9ba045fa144fdcaafc5135a64285cd6e8b18989c.
//
// Solidity: event IgnoreNonces(address sender, uint256[] nonces)
func (_VaultY *VaultYFilterer) WatchIgnoreNonces(opts *bind.WatchOpts, sink chan<- *VaultYIgnoreNonces) (event.Subscription, error) {

	logs, sub, err := _VaultY.contract.WatchLogs(opts, "IgnoreNonces")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultYIgnoreNonces)
				if err := _VaultY.contract.UnpackLog(event, "IgnoreNonces", log); err != nil {
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

// ParseIgnoreNonces is a log parse operation binding the contract event 0x33fe1cf6879c44bf1b70173d9ba045fa144fdcaafc5135a64285cd6e8b18989c.
//
// Solidity: event IgnoreNonces(address sender, uint256[] nonces)
func (_VaultY *VaultYFilterer) ParseIgnoreNonces(log types.Log) (*VaultYIgnoreNonces, error) {
	event := new(VaultYIgnoreNonces)
	if err := _VaultY.contract.UnpackLog(event, "IgnoreNonces", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultYPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the VaultY contract.
type VaultYPausedIterator struct {
	Event *VaultYPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultYPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultYPaused)
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
		it.Event = new(VaultYPaused)
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
func (it *VaultYPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultYPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultYPaused represents a Paused event raised by the VaultY contract.
type VaultYPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VaultY *VaultYFilterer) FilterPaused(opts *bind.FilterOpts) (*VaultYPausedIterator, error) {

	logs, sub, err := _VaultY.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &VaultYPausedIterator{contract: _VaultY.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VaultY *VaultYFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *VaultYPaused) (event.Subscription, error) {

	logs, sub, err := _VaultY.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultYPaused)
				if err := _VaultY.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VaultY *VaultYFilterer) ParsePaused(log types.Log) (*VaultYPaused, error) {
	event := new(VaultYPaused)
	if err := _VaultY.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultYPriceOracleChangedIterator is returned from FilterPriceOracleChanged and is used to iterate over the raw logs and unpacked data for PriceOracleChanged events raised by the VaultY contract.
type VaultYPriceOracleChangedIterator struct {
	Event *VaultYPriceOracleChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultYPriceOracleChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultYPriceOracleChanged)
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
		it.Event = new(VaultYPriceOracleChanged)
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
func (it *VaultYPriceOracleChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultYPriceOracleChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultYPriceOracleChanged represents a PriceOracleChanged event raised by the VaultY contract.
type VaultYPriceOracleChanged struct {
	Sender         common.Address
	NewPriceOracle common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPriceOracleChanged is a free log retrieval operation binding the contract event 0x40bddd72ea96b80dae14e3d13e8ce2c4ecd2500d88c6c0004d24a00deab28f9c.
//
// Solidity: event PriceOracleChanged(address indexed sender, address newPriceOracle)
func (_VaultY *VaultYFilterer) FilterPriceOracleChanged(opts *bind.FilterOpts, sender []common.Address) (*VaultYPriceOracleChangedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultY.contract.FilterLogs(opts, "PriceOracleChanged", senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultYPriceOracleChangedIterator{contract: _VaultY.contract, event: "PriceOracleChanged", logs: logs, sub: sub}, nil
}

// WatchPriceOracleChanged is a free log subscription operation binding the contract event 0x40bddd72ea96b80dae14e3d13e8ce2c4ecd2500d88c6c0004d24a00deab28f9c.
//
// Solidity: event PriceOracleChanged(address indexed sender, address newPriceOracle)
func (_VaultY *VaultYFilterer) WatchPriceOracleChanged(opts *bind.WatchOpts, sink chan<- *VaultYPriceOracleChanged, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultY.contract.WatchLogs(opts, "PriceOracleChanged", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultYPriceOracleChanged)
				if err := _VaultY.contract.UnpackLog(event, "PriceOracleChanged", log); err != nil {
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

// ParsePriceOracleChanged is a log parse operation binding the contract event 0x40bddd72ea96b80dae14e3d13e8ce2c4ecd2500d88c6c0004d24a00deab28f9c.
//
// Solidity: event PriceOracleChanged(address indexed sender, address newPriceOracle)
func (_VaultY *VaultYFilterer) ParsePriceOracleChanged(log types.Log) (*VaultYPriceOracleChanged, error) {
	event := new(VaultYPriceOracleChanged)
	if err := _VaultY.contract.UnpackLog(event, "PriceOracleChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultYRefundIterator is returned from FilterRefund and is used to iterate over the raw logs and unpacked data for Refund events raised by the VaultY contract.
type VaultYRefundIterator struct {
	Event *VaultYRefund // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultYRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultYRefund)
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
		it.Event = new(VaultYRefund)
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
func (it *VaultYRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultYRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultYRefund represents a Refund event raised by the VaultY contract.
type VaultYRefund struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefund is a free log retrieval operation binding the contract event 0xf40cc8c1a1d17359049ba500cfc894596a692cffc9d03943cd92ec2e159cf6ae.
//
// Solidity: event Refund(address token, address to, uint256 amount)
func (_VaultY *VaultYFilterer) FilterRefund(opts *bind.FilterOpts) (*VaultYRefundIterator, error) {

	logs, sub, err := _VaultY.contract.FilterLogs(opts, "Refund")
	if err != nil {
		return nil, err
	}
	return &VaultYRefundIterator{contract: _VaultY.contract, event: "Refund", logs: logs, sub: sub}, nil
}

// WatchRefund is a free log subscription operation binding the contract event 0xf40cc8c1a1d17359049ba500cfc894596a692cffc9d03943cd92ec2e159cf6ae.
//
// Solidity: event Refund(address token, address to, uint256 amount)
func (_VaultY *VaultYFilterer) WatchRefund(opts *bind.WatchOpts, sink chan<- *VaultYRefund) (event.Subscription, error) {

	logs, sub, err := _VaultY.contract.WatchLogs(opts, "Refund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultYRefund)
				if err := _VaultY.contract.UnpackLog(event, "Refund", log); err != nil {
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

// ParseRefund is a log parse operation binding the contract event 0xf40cc8c1a1d17359049ba500cfc894596a692cffc9d03943cd92ec2e159cf6ae.
//
// Solidity: event Refund(address token, address to, uint256 amount)
func (_VaultY *VaultYFilterer) ParseRefund(log types.Log) (*VaultYRefund, error) {
	event := new(VaultYRefund)
	if err := _VaultY.contract.UnpackLog(event, "Refund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultYRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the VaultY contract.
type VaultYRoleAdminChangedIterator struct {
	Event *VaultYRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultYRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultYRoleAdminChanged)
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
		it.Event = new(VaultYRoleAdminChanged)
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
func (it *VaultYRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultYRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultYRoleAdminChanged represents a RoleAdminChanged event raised by the VaultY contract.
type VaultYRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VaultY *VaultYFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*VaultYRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _VaultY.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &VaultYRoleAdminChangedIterator{contract: _VaultY.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VaultY *VaultYFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *VaultYRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _VaultY.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultYRoleAdminChanged)
				if err := _VaultY.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VaultY *VaultYFilterer) ParseRoleAdminChanged(log types.Log) (*VaultYRoleAdminChanged, error) {
	event := new(VaultYRoleAdminChanged)
	if err := _VaultY.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultYRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the VaultY contract.
type VaultYRoleGrantedIterator struct {
	Event *VaultYRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultYRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultYRoleGranted)
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
		it.Event = new(VaultYRoleGranted)
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
func (it *VaultYRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultYRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultYRoleGranted represents a RoleGranted event raised by the VaultY contract.
type VaultYRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultY *VaultYFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VaultYRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultY.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultYRoleGrantedIterator{contract: _VaultY.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultY *VaultYFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *VaultYRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultY.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultYRoleGranted)
				if err := _VaultY.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultY *VaultYFilterer) ParseRoleGranted(log types.Log) (*VaultYRoleGranted, error) {
	event := new(VaultYRoleGranted)
	if err := _VaultY.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultYRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the VaultY contract.
type VaultYRoleRevokedIterator struct {
	Event *VaultYRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultYRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultYRoleRevoked)
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
		it.Event = new(VaultYRoleRevoked)
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
func (it *VaultYRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultYRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultYRoleRevoked represents a RoleRevoked event raised by the VaultY contract.
type VaultYRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultY *VaultYFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VaultYRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultY.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultYRoleRevokedIterator{contract: _VaultY.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultY *VaultYFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *VaultYRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultY.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultYRoleRevoked)
				if err := _VaultY.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultY *VaultYFilterer) ParseRoleRevoked(log types.Log) (*VaultYRoleRevoked, error) {
	event := new(VaultYRoleRevoked)
	if err := _VaultY.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultYSkipNonceIterator is returned from FilterSkipNonce and is used to iterate over the raw logs and unpacked data for SkipNonce events raised by the VaultY contract.
type VaultYSkipNonceIterator struct {
	Event *VaultYSkipNonce // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultYSkipNonceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultYSkipNonce)
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
		it.Event = new(VaultYSkipNonce)
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
func (it *VaultYSkipNonceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultYSkipNonceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultYSkipNonce represents a SkipNonce event raised by the VaultY contract.
type VaultYSkipNonce struct {
	Start *big.Int
	Step  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSkipNonce is a free log retrieval operation binding the contract event 0x503d8507b9fedcafb1c6a87c33912dcfc181e9f19b17345e8b53c26b3e9fc29e.
//
// Solidity: event SkipNonce(uint256 start, uint256 step)
func (_VaultY *VaultYFilterer) FilterSkipNonce(opts *bind.FilterOpts) (*VaultYSkipNonceIterator, error) {

	logs, sub, err := _VaultY.contract.FilterLogs(opts, "SkipNonce")
	if err != nil {
		return nil, err
	}
	return &VaultYSkipNonceIterator{contract: _VaultY.contract, event: "SkipNonce", logs: logs, sub: sub}, nil
}

// WatchSkipNonce is a free log subscription operation binding the contract event 0x503d8507b9fedcafb1c6a87c33912dcfc181e9f19b17345e8b53c26b3e9fc29e.
//
// Solidity: event SkipNonce(uint256 start, uint256 step)
func (_VaultY *VaultYFilterer) WatchSkipNonce(opts *bind.WatchOpts, sink chan<- *VaultYSkipNonce) (event.Subscription, error) {

	logs, sub, err := _VaultY.contract.WatchLogs(opts, "SkipNonce")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultYSkipNonce)
				if err := _VaultY.contract.UnpackLog(event, "SkipNonce", log); err != nil {
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

// ParseSkipNonce is a log parse operation binding the contract event 0x503d8507b9fedcafb1c6a87c33912dcfc181e9f19b17345e8b53c26b3e9fc29e.
//
// Solidity: event SkipNonce(uint256 start, uint256 step)
func (_VaultY *VaultYFilterer) ParseSkipNonce(log types.Log) (*VaultYSkipNonce, error) {
	event := new(VaultYSkipNonce)
	if err := _VaultY.contract.UnpackLog(event, "SkipNonce", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultYStakeAddedIterator is returned from FilterStakeAdded and is used to iterate over the raw logs and unpacked data for StakeAdded events raised by the VaultY contract.
type VaultYStakeAddedIterator struct {
	Event *VaultYStakeAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultYStakeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultYStakeAdded)
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
		it.Event = new(VaultYStakeAdded)
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
func (it *VaultYStakeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultYStakeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultYStakeAdded represents a StakeAdded event raised by the VaultY contract.
type VaultYStakeAdded struct {
	From   common.Address
	Epoch  *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeAdded is a free log retrieval operation binding the contract event 0x270d6dd254edd1d985c81cf7861b8f28fb06b6d719df04d90464034d43412440.
//
// Solidity: event StakeAdded(address indexed from, uint256 indexed epoch, uint256 amount)
func (_VaultY *VaultYFilterer) FilterStakeAdded(opts *bind.FilterOpts, from []common.Address, epoch []*big.Int) (*VaultYStakeAddedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _VaultY.contract.FilterLogs(opts, "StakeAdded", fromRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &VaultYStakeAddedIterator{contract: _VaultY.contract, event: "StakeAdded", logs: logs, sub: sub}, nil
}

// WatchStakeAdded is a free log subscription operation binding the contract event 0x270d6dd254edd1d985c81cf7861b8f28fb06b6d719df04d90464034d43412440.
//
// Solidity: event StakeAdded(address indexed from, uint256 indexed epoch, uint256 amount)
func (_VaultY *VaultYFilterer) WatchStakeAdded(opts *bind.WatchOpts, sink chan<- *VaultYStakeAdded, from []common.Address, epoch []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _VaultY.contract.WatchLogs(opts, "StakeAdded", fromRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultYStakeAdded)
				if err := _VaultY.contract.UnpackLog(event, "StakeAdded", log); err != nil {
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

// ParseStakeAdded is a log parse operation binding the contract event 0x270d6dd254edd1d985c81cf7861b8f28fb06b6d719df04d90464034d43412440.
//
// Solidity: event StakeAdded(address indexed from, uint256 indexed epoch, uint256 amount)
func (_VaultY *VaultYFilterer) ParseStakeAdded(log types.Log) (*VaultYStakeAdded, error) {
	event := new(VaultYStakeAdded)
	if err := _VaultY.contract.UnpackLog(event, "StakeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultYStakeRemovedIterator is returned from FilterStakeRemoved and is used to iterate over the raw logs and unpacked data for StakeRemoved events raised by the VaultY contract.
type VaultYStakeRemovedIterator struct {
	Event *VaultYStakeRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultYStakeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultYStakeRemoved)
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
		it.Event = new(VaultYStakeRemoved)
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
func (it *VaultYStakeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultYStakeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultYStakeRemoved represents a StakeRemoved event raised by the VaultY contract.
type VaultYStakeRemoved struct {
	From   common.Address
	Epoch  *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeRemoved is a free log retrieval operation binding the contract event 0xdcca95406ac9554449be02d88dcdf9c877f96e4c02bdad4bd5cadefc98a20e3d.
//
// Solidity: event StakeRemoved(address indexed from, uint256 indexed epoch, uint256 amount)
func (_VaultY *VaultYFilterer) FilterStakeRemoved(opts *bind.FilterOpts, from []common.Address, epoch []*big.Int) (*VaultYStakeRemovedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _VaultY.contract.FilterLogs(opts, "StakeRemoved", fromRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &VaultYStakeRemovedIterator{contract: _VaultY.contract, event: "StakeRemoved", logs: logs, sub: sub}, nil
}

// WatchStakeRemoved is a free log subscription operation binding the contract event 0xdcca95406ac9554449be02d88dcdf9c877f96e4c02bdad4bd5cadefc98a20e3d.
//
// Solidity: event StakeRemoved(address indexed from, uint256 indexed epoch, uint256 amount)
func (_VaultY *VaultYFilterer) WatchStakeRemoved(opts *bind.WatchOpts, sink chan<- *VaultYStakeRemoved, from []common.Address, epoch []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _VaultY.contract.WatchLogs(opts, "StakeRemoved", fromRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultYStakeRemoved)
				if err := _VaultY.contract.UnpackLog(event, "StakeRemoved", log); err != nil {
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

// ParseStakeRemoved is a log parse operation binding the contract event 0xdcca95406ac9554449be02d88dcdf9c877f96e4c02bdad4bd5cadefc98a20e3d.
//
// Solidity: event StakeRemoved(address indexed from, uint256 indexed epoch, uint256 amount)
func (_VaultY *VaultYFilterer) ParseStakeRemoved(log types.Log) (*VaultYStakeRemoved, error) {
	event := new(VaultYStakeRemoved)
	if err := _VaultY.contract.UnpackLog(event, "StakeRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultYTipAccountChangedIterator is returned from FilterTipAccountChanged and is used to iterate over the raw logs and unpacked data for TipAccountChanged events raised by the VaultY contract.
type VaultYTipAccountChangedIterator struct {
	Event *VaultYTipAccountChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultYTipAccountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultYTipAccountChanged)
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
		it.Event = new(VaultYTipAccountChanged)
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
func (it *VaultYTipAccountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultYTipAccountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultYTipAccountChanged represents a TipAccountChanged event raised by the VaultY contract.
type VaultYTipAccountChanged struct {
	Sender        common.Address
	NewTipAccount common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTipAccountChanged is a free log retrieval operation binding the contract event 0x028dd631374b48803cfa88ca1e6693001c280e4ffa9e573bf1002723cb03ae15.
//
// Solidity: event TipAccountChanged(address indexed sender, address newTipAccount)
func (_VaultY *VaultYFilterer) FilterTipAccountChanged(opts *bind.FilterOpts, sender []common.Address) (*VaultYTipAccountChangedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultY.contract.FilterLogs(opts, "TipAccountChanged", senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultYTipAccountChangedIterator{contract: _VaultY.contract, event: "TipAccountChanged", logs: logs, sub: sub}, nil
}

// WatchTipAccountChanged is a free log subscription operation binding the contract event 0x028dd631374b48803cfa88ca1e6693001c280e4ffa9e573bf1002723cb03ae15.
//
// Solidity: event TipAccountChanged(address indexed sender, address newTipAccount)
func (_VaultY *VaultYFilterer) WatchTipAccountChanged(opts *bind.WatchOpts, sink chan<- *VaultYTipAccountChanged, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultY.contract.WatchLogs(opts, "TipAccountChanged", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultYTipAccountChanged)
				if err := _VaultY.contract.UnpackLog(event, "TipAccountChanged", log); err != nil {
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

// ParseTipAccountChanged is a log parse operation binding the contract event 0x028dd631374b48803cfa88ca1e6693001c280e4ffa9e573bf1002723cb03ae15.
//
// Solidity: event TipAccountChanged(address indexed sender, address newTipAccount)
func (_VaultY *VaultYFilterer) ParseTipAccountChanged(log types.Log) (*VaultYTipAccountChanged, error) {
	event := new(VaultYTipAccountChanged)
	if err := _VaultY.contract.UnpackLog(event, "TipAccountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultYTipRateChangedIterator is returned from FilterTipRateChanged and is used to iterate over the raw logs and unpacked data for TipRateChanged events raised by the VaultY contract.
type VaultYTipRateChangedIterator struct {
	Event *VaultYTipRateChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultYTipRateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultYTipRateChanged)
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
		it.Event = new(VaultYTipRateChanged)
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
func (it *VaultYTipRateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultYTipRateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultYTipRateChanged represents a TipRateChanged event raised by the VaultY contract.
type VaultYTipRateChanged struct {
	Sender      common.Address
	SourceToken common.Address
	MappedToken common.Address
	NewTipRate  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTipRateChanged is a free log retrieval operation binding the contract event 0x7551456b2f0d91a3f37c1ac5ab2c19b7ed477502a8236b6e280061e05bd144a8.
//
// Solidity: event TipRateChanged(address indexed sender, address sourceToken, address mappedToken, uint256 newTipRate)
func (_VaultY *VaultYFilterer) FilterTipRateChanged(opts *bind.FilterOpts, sender []common.Address) (*VaultYTipRateChangedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultY.contract.FilterLogs(opts, "TipRateChanged", senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultYTipRateChangedIterator{contract: _VaultY.contract, event: "TipRateChanged", logs: logs, sub: sub}, nil
}

// WatchTipRateChanged is a free log subscription operation binding the contract event 0x7551456b2f0d91a3f37c1ac5ab2c19b7ed477502a8236b6e280061e05bd144a8.
//
// Solidity: event TipRateChanged(address indexed sender, address sourceToken, address mappedToken, uint256 newTipRate)
func (_VaultY *VaultYFilterer) WatchTipRateChanged(opts *bind.WatchOpts, sink chan<- *VaultYTipRateChanged, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultY.contract.WatchLogs(opts, "TipRateChanged", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultYTipRateChanged)
				if err := _VaultY.contract.UnpackLog(event, "TipRateChanged", log); err != nil {
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

// ParseTipRateChanged is a log parse operation binding the contract event 0x7551456b2f0d91a3f37c1ac5ab2c19b7ed477502a8236b6e280061e05bd144a8.
//
// Solidity: event TipRateChanged(address indexed sender, address sourceToken, address mappedToken, uint256 newTipRate)
func (_VaultY *VaultYFilterer) ParseTipRateChanged(log types.Log) (*VaultYTipRateChanged, error) {
	event := new(VaultYTipRateChanged)
	if err := _VaultY.contract.UnpackLog(event, "TipRateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultYTokenBurnIterator is returned from FilterTokenBurn and is used to iterate over the raw logs and unpacked data for TokenBurn events raised by the VaultY contract.
type VaultYTokenBurnIterator struct {
	Event *VaultYTokenBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultYTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultYTokenBurn)
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
		it.Event = new(VaultYTokenBurn)
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
func (it *VaultYTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultYTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultYTokenBurn represents a TokenBurn event raised by the VaultY contract.
type VaultYTokenBurn struct {
	Vault         common.Address
	SourceChainid *big.Int
	MappedChainid *big.Int
	SourceToken   common.Address
	MappedToken   common.Address
	From          common.Address
	Amount        *big.Int
	Tip           *big.Int
	Nonce         *big.Int
	BlockNumber   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenBurn is a free log retrieval operation binding the contract event 0xe6064164b9ee4a11aab4f75dcadd3198373528b269d446c73ca48565eaf1c28f.
//
// Solidity: event TokenBurn(address vault, uint256 sourceChainid, uint256 mappedChainid, address indexed sourceToken, address indexed mappedToken, address from, uint256 amount, uint256 tip, uint256 indexed nonce, uint256 blockNumber)
func (_VaultY *VaultYFilterer) FilterTokenBurn(opts *bind.FilterOpts, sourceToken []common.Address, mappedToken []common.Address, nonce []*big.Int) (*VaultYTokenBurnIterator, error) {

	var sourceTokenRule []interface{}
	for _, sourceTokenItem := range sourceToken {
		sourceTokenRule = append(sourceTokenRule, sourceTokenItem)
	}
	var mappedTokenRule []interface{}
	for _, mappedTokenItem := range mappedToken {
		mappedTokenRule = append(mappedTokenRule, mappedTokenItem)
	}

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _VaultY.contract.FilterLogs(opts, "TokenBurn", sourceTokenRule, mappedTokenRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &VaultYTokenBurnIterator{contract: _VaultY.contract, event: "TokenBurn", logs: logs, sub: sub}, nil
}

// WatchTokenBurn is a free log subscription operation binding the contract event 0xe6064164b9ee4a11aab4f75dcadd3198373528b269d446c73ca48565eaf1c28f.
//
// Solidity: event TokenBurn(address vault, uint256 sourceChainid, uint256 mappedChainid, address indexed sourceToken, address indexed mappedToken, address from, uint256 amount, uint256 tip, uint256 indexed nonce, uint256 blockNumber)
func (_VaultY *VaultYFilterer) WatchTokenBurn(opts *bind.WatchOpts, sink chan<- *VaultYTokenBurn, sourceToken []common.Address, mappedToken []common.Address, nonce []*big.Int) (event.Subscription, error) {

	var sourceTokenRule []interface{}
	for _, sourceTokenItem := range sourceToken {
		sourceTokenRule = append(sourceTokenRule, sourceTokenItem)
	}
	var mappedTokenRule []interface{}
	for _, mappedTokenItem := range mappedToken {
		mappedTokenRule = append(mappedTokenRule, mappedTokenItem)
	}

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _VaultY.contract.WatchLogs(opts, "TokenBurn", sourceTokenRule, mappedTokenRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultYTokenBurn)
				if err := _VaultY.contract.UnpackLog(event, "TokenBurn", log); err != nil {
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

// ParseTokenBurn is a log parse operation binding the contract event 0xe6064164b9ee4a11aab4f75dcadd3198373528b269d446c73ca48565eaf1c28f.
//
// Solidity: event TokenBurn(address vault, uint256 sourceChainid, uint256 mappedChainid, address indexed sourceToken, address indexed mappedToken, address from, uint256 amount, uint256 tip, uint256 indexed nonce, uint256 blockNumber)
func (_VaultY *VaultYFilterer) ParseTokenBurn(log types.Log) (*VaultYTokenBurn, error) {
	event := new(VaultYTokenBurn)
	if err := _VaultY.contract.UnpackLog(event, "TokenBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultYTokenDepositIterator is returned from FilterTokenDeposit and is used to iterate over the raw logs and unpacked data for TokenDeposit events raised by the VaultY contract.
type VaultYTokenDepositIterator struct {
	Event *VaultYTokenDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultYTokenDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultYTokenDeposit)
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
		it.Event = new(VaultYTokenDeposit)
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
func (it *VaultYTokenDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultYTokenDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultYTokenDeposit represents a TokenDeposit event raised by the VaultY contract.
type VaultYTokenDeposit struct {
	Vault         common.Address
	SourceChainid *big.Int
	MappedChainid *big.Int
	SourceToken   common.Address
	MappedToken   common.Address
	From          common.Address
	Amount        *big.Int
	Tip           *big.Int
	Nonce         *big.Int
	BlockNumber   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenDeposit is a free log retrieval operation binding the contract event 0x41f138048b690e2a81e12b36d21c06d77c6a6ee2a5b4ed203690f54a4d5ca4bf.
//
// Solidity: event TokenDeposit(address vault, uint256 sourceChainid, uint256 mappedChainid, address indexed sourceToken, address indexed mappedToken, address from, uint256 amount, uint256 tip, uint256 indexed nonce, uint256 blockNumber)
func (_VaultY *VaultYFilterer) FilterTokenDeposit(opts *bind.FilterOpts, sourceToken []common.Address, mappedToken []common.Address, nonce []*big.Int) (*VaultYTokenDepositIterator, error) {

	var sourceTokenRule []interface{}
	for _, sourceTokenItem := range sourceToken {
		sourceTokenRule = append(sourceTokenRule, sourceTokenItem)
	}
	var mappedTokenRule []interface{}
	for _, mappedTokenItem := range mappedToken {
		mappedTokenRule = append(mappedTokenRule, mappedTokenItem)
	}

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _VaultY.contract.FilterLogs(opts, "TokenDeposit", sourceTokenRule, mappedTokenRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &VaultYTokenDepositIterator{contract: _VaultY.contract, event: "TokenDeposit", logs: logs, sub: sub}, nil
}

// WatchTokenDeposit is a free log subscription operation binding the contract event 0x41f138048b690e2a81e12b36d21c06d77c6a6ee2a5b4ed203690f54a4d5ca4bf.
//
// Solidity: event TokenDeposit(address vault, uint256 sourceChainid, uint256 mappedChainid, address indexed sourceToken, address indexed mappedToken, address from, uint256 amount, uint256 tip, uint256 indexed nonce, uint256 blockNumber)
func (_VaultY *VaultYFilterer) WatchTokenDeposit(opts *bind.WatchOpts, sink chan<- *VaultYTokenDeposit, sourceToken []common.Address, mappedToken []common.Address, nonce []*big.Int) (event.Subscription, error) {

	var sourceTokenRule []interface{}
	for _, sourceTokenItem := range sourceToken {
		sourceTokenRule = append(sourceTokenRule, sourceTokenItem)
	}
	var mappedTokenRule []interface{}
	for _, mappedTokenItem := range mappedToken {
		mappedTokenRule = append(mappedTokenRule, mappedTokenItem)
	}

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _VaultY.contract.WatchLogs(opts, "TokenDeposit", sourceTokenRule, mappedTokenRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultYTokenDeposit)
				if err := _VaultY.contract.UnpackLog(event, "TokenDeposit", log); err != nil {
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

// ParseTokenDeposit is a log parse operation binding the contract event 0x41f138048b690e2a81e12b36d21c06d77c6a6ee2a5b4ed203690f54a4d5ca4bf.
//
// Solidity: event TokenDeposit(address vault, uint256 sourceChainid, uint256 mappedChainid, address indexed sourceToken, address indexed mappedToken, address from, uint256 amount, uint256 tip, uint256 indexed nonce, uint256 blockNumber)
func (_VaultY *VaultYFilterer) ParseTokenDeposit(log types.Log) (*VaultYTokenDeposit, error) {
	event := new(VaultYTokenDeposit)
	if err := _VaultY.contract.UnpackLog(event, "TokenDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultYUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the VaultY contract.
type VaultYUnpausedIterator struct {
	Event *VaultYUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultYUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultYUnpaused)
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
		it.Event = new(VaultYUnpaused)
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
func (it *VaultYUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultYUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultYUnpaused represents a Unpaused event raised by the VaultY contract.
type VaultYUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VaultY *VaultYFilterer) FilterUnpaused(opts *bind.FilterOpts) (*VaultYUnpausedIterator, error) {

	logs, sub, err := _VaultY.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &VaultYUnpausedIterator{contract: _VaultY.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VaultY *VaultYFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *VaultYUnpaused) (event.Subscription, error) {

	logs, sub, err := _VaultY.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultYUnpaused)
				if err := _VaultY.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VaultY *VaultYFilterer) ParseUnpaused(log types.Log) (*VaultYUnpaused, error) {
	event := new(VaultYUnpaused)
	if err := _VaultY.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

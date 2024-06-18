// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// IProtocolServiceVerificationInfo is an auto generated low-level Go binding around an user-defined struct.
type IProtocolServiceVerificationInfo struct {
	Result uint8
	Index  uint32
	Signer common.Address
	V      uint8
	R      [32]byte
	S      [32]byte
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimMaliciousTeeRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"ClaimRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32[]\",\"name\":\"vrfChosen\",\"type\":\"uint32[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ConfirmVrfNodes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Delegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"NodeActivate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"NodeClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"NodeClear\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"commissionRate\",\"type\":\"uint32\"}],\"name\":\"NodeModifyCommissionRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"name\":\"NodeRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"attestationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumIProtocolService.AttestationResult\",\"name\":\"result\",\"type\":\"uint8\"}],\"name\":\"NodeReportVerification\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"attestationID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumIProtocolService.AttestationResult\",\"name\":\"result\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structIProtocolService.VerificationInfo[]\",\"name\":\"infos\",\"type\":\"tuple[]\"}],\"name\":\"NodeReportVerificationBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"NodeSetClaimer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"attestationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"NodeSlash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Redelegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"attestationIDs\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"attestationInfos\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"}],\"name\":\"TeeReportAttestations\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"attestationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TeeSlash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TeeStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TeeUnstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Undelegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"settings\",\"type\":\"address\"}],\"name\":\"UpdateSettingsAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vrf\",\"type\":\"address\"}],\"name\":\"UpdateVrfAddress\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_UINT32\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activeVrfNodeList\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"attestationVerifiedNode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"attestations\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"valid\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"invalid\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"malicious\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"vrfChosenID\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"slashed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"carvNft\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"carvToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"carvVrf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"checkClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"attestationID\",\"type\":\"bytes32\"}],\"name\":\"claimMaliciousTeeRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"delegation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegationWeights\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712DomainHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"}],\"name\":\"fulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"globalDailyActiveNodes\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"carvToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"carvNft_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"modifyAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tee\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"grant\",\"type\":\"bool\"}],\"name\":\"modifyTeeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"nodeAddrByID\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"nodeClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nodeClaimedTeeRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"nodeDailyActive\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"replacedNode\",\"type\":\"address\"}],\"name\":\"nodeEnter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"replacedNode\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"nodeEnterWithSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"nodeExitWithSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeIndex\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodeInfos\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"listIndex\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"lastConfirmDate\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"missedVerifyCount\",\"type\":\"uint64\"},{\"internalType\":\"int256\",\"name\":\"selfTotalRewards\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"selfClaimedRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegationRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastEnterTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"commissionRate\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"commissionRateLastModifyAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"commissionRate\",\"type\":\"uint32\"}],\"name\":\"nodeModifyCommissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"commissionRate\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"nodeModifyCommissionRateWithSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"nodeReportDailyActive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"attestationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"enumIProtocolService.AttestationResult\",\"name\":\"result\",\"type\":\"uint8\"}],\"name\":\"nodeReportVerification\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"attestationID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumIProtocolService.AttestationResult\",\"name\":\"result\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIProtocolService.VerificationInfo[]\",\"name\":\"infos\",\"type\":\"tuple[]\"}],\"name\":\"nodeReportVerificationBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"nodeSetRewardClaimerWithSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"attestationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"nodeSlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nodeSlashed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"redelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"request2AttestationIDs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settings\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"attestationInfos\",\"type\":\"string[]\"}],\"name\":\"teeReportAttestations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"teeRoles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"attestationID\",\"type\":\"bytes32\"}],\"name\":\"teeSlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"teeStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"teeStakeInfos\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"staked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastReportAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"todayIndex\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"todayOffset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenRewardInfos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"initialRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimedRewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"settings_\",\"type\":\"address\"}],\"name\":\"updateSettingsAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"carvVrf_\",\"type\":\"address\"}],\"name\":\"updateVrfAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vrfChosenIndex\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vrfChosenMap\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50615e9580620000216000396000f3fe608060405234801561001057600080fd5b50600436106103835760003560e01c80637f32e167116101de578063cc4ed7c51161010f578063e17ec726116100ad578063f39a19bf1161007c578063f39a19bf14610abb578063f851a44014610ace578063f94dc4bc14610ae7578063fbfa77cf14610af057600080fd5b8063e17ec72614610a5c578063e39a0beb14610a6f578063e93ee21214610a82578063ea71562714610a9557600080fd5b8063d46849db116100e9578063d46849db14610a26578063dc57a64714610a2e578063df5e1fdd14610a36578063e06174e414610a4957600080fd5b8063cc4ed7c5146109d5578063ce3f3d9514610a03578063cf756fdf14610a1357600080fd5b8063a91e54791161017c578063b50bd0d411610156578063b50bd0d414610965578063c018a80614610978578063c4a16c52146109a1578063c71d5b2f146109b457600080fd5b8063a91e5479146108ee578063ac9650d81461091c578063b1a61dda1461093c57600080fd5b80639ad4f752116101b85780639ad4f752146108845780639c509861146108975780639d9b4270146108aa578063a3b4c737146108bd57600080fd5b80637f32e167146107a05780638a4fd0c8146107b3578063940992a3146107c657600080fd5b806337c68423116102b857806368f418101161025657806373df96ed1161023057806373df96ed1461064b578063796ba4d71461065e578063797e18e8146106755780637a83fe2a1461078d57600080fd5b806368f41810146105b15780636c68c0e1146106015780637334c9ba1461061457600080fd5b806347b8706c1161029257806347b8706c146105655780634bbc4c97146105785780634c9a41f91461058b5780635890b9581461059e57600080fd5b806337c684231461050057806338ba4614146105085780633b85a3461461051b57600080fd5b8063195f3a9c116103255780632d1af52e116102ff5780632d1af52e1461049d57806334bee6ee146104a557806334d37fa0146104c857806335a55caa146104db57600080fd5b8063195f3a9c1461044c5780631cde81d51461045f57806323031abe1461048a57600080fd5b8063055b1edd11610361578063055b1edd146103fe57806308bbb824146104135780630962ef791461042657806309d65fd81461043957600080fd5b80630225de921461038857806303c13728146103b557806304901b93146103f3575b600080fd5b61039b61039636600461519c565b610b03565b60405163ffffffff90911681526020015b60405180910390f35b6103e36103c33660046151ca565b601460209081526000928352604080842090915290825290205460ff1681565b60405190151581526020016103ac565b61039b63ffffffff81565b61041161040c3660046151f6565b610b3d565b005b610411610421366004615213565b610bc5565b61041161043436600461519c565b610e89565b61039b61044736600461525c565b6110e4565b61041161045a366004615278565b61112d565b600554610472906001600160a01b031681565b6040516001600160a01b0390911681526020016103ac565b6104116104983660046151f6565b61113a565b6104116111db565b6103e36104b33660046151f6565b60016020526000908152604090205460ff1681565b6104116104d63660046151f6565b61139c565b6103e36104e936600461519c565b6000908152600f6020526040902060010154151590565b61039b61140d565b6104116105163660046152fe565b6114a3565b61054a61052936600461519c565b600f6020526000908152604090208054600182015460029092015490919083565b604080519384526020840192909252908201526060016103ac565b6104116105733660046153b1565b611753565b61041161058636600461519c565b6117d4565b6104116105993660046151f6565b611a16565b6104116105ac366004615422565b611ab0565b6105e46105bf3660046151f6565b60106020526000908152604090208054600182015460029092015460ff909116919083565b6040805193151584526020840192909252908201526060016103ac565b61041161060f36600461519c565b611b27565b6106386106223660046151f6565b600d6020526000908152604090205461ffff1681565b60405161ffff90911681526020016103ac565b6104116106593660046151f6565b611d75565b60075461039b90600160a01b900463ffffffff1681565b61070e6106833660046151f6565b600e60205260009081526040902080546001820154600283015460038401546004850154600586015460069096015463ffffffff808716976401000000008804821697600160401b810460ff1697600160481b8204841697600160681b90920467ffffffffffffffff169690959194909391926001600160a01b03811692600160a01b90910416908c565b6040805163ffffffff9d8e1681529b8d1660208d0152991515998b0199909952968a1660608a015267ffffffffffffffff909516608089015260a088019390935260c087019190915260e08601526101008501526001600160a01b031661012084015292909216610140820152610160810191909152610180016103ac565b600754610472906001600160a01b031681565b6104116107ae366004615450565b611d7f565b6104116107c13660046154a0565b611dec565b6108356107d436600461519c565b601160205260009081526040902080546001909101546001600160a01b03821691600160a01b810461ffff90811692600160b01b8304821692600160c01b810490921691600160d01b810463ffffffff1691600160f01b90910460ff169087565b604080516001600160a01b03909816885261ffff96871660208901529486169487019490945293909116606085015263ffffffff16608084015290151560a083015260c082015260e0016103ac565b610411610892366004615213565b611e61565b6104116108a536600461519c565b61222f565b6104116108b83660046154c2565b612485565b61039b6108cb3660046155c9565b601760209081526000928352604080842090915290825290205463ffffffff1681565b6103e36108fc366004615213565b601360209081526000928352604080842090915290825290205460ff1681565b61092f61092a3660046155fe565b61271d565b6040516103ac91906156c3565b61047261094a36600461519c565b600c602052600090815260409020546001600160a01b031681565b610411610973366004615725565b612811565b610472610986366004615278565b600b602052600090815260409020546001600160a01b031681565b6104116109af3660046153b1565b612ced565b6109c76109c23660046157a4565b612dc2565b6040519081526020016103ac565b6103e36109e33660046151ca565b601560209081526000928352604080842090915290825290205460ff1681565b60095461039b9063ffffffff1681565b610411610a213660046157c6565b612df3565b610411613031565b6109c761303c565b610411610a44366004615817565b6130cd565b600654610472906001600160a01b031681565b600454610472906001600160a01b031681565b610411610a7d36600461519c565b6134d3565b610411610a90366004615864565b6137d7565b61039b610aa3366004615278565b60166020526000908152604090205463ffffffff1681565b610411610ac93660046151f6565b613995565b600054610472906201000090046001600160a01b031681565b6109c760025481565b600354610472906001600160a01b031681565b60088181548110610b1357600080fd5b9060005260206000209060089182820401919006600402915054906101000a900463ffffffff1681565b6000546201000090046001600160a01b03163314610b8e5760405162461bcd60e51b81526020600482015260096024820152682737ba1030b236b4b760b91b60448201526064015b60405180910390fd5b600080546001600160a01b03909216620100000275ffffffffffffffffffffffffffffffffffffffff000019909216919091179055565b6005546040516331a9108f60e11b815260048101849052839133916001600160a01b0390911690636352211e90602401602060405180830381865afa158015610c12573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c369190615897565b6001600160a01b031614610c785760405162461bcd60e51b81526020600482015260096024820152682737ba1037bbb732b960b91b6044820152606401610b85565b6000838152600c60205260409020546001600160a01b031615610cdd5760405162461bcd60e51b815260206004820152601060248201527f416c72656164792064656c6567617465000000000000000000000000000000006044820152606401610b85565b600660009054906101000a90046001600160a01b03166001600160a01b03166315fbbc9d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d30573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d5491906158b4565b6001600160a01b0383166000908152600d602052604090205461ffff918216911610610db55760405162461bcd60e51b815260206004820152601060248201526f4d6178206e6f6465207765696768747360801b6044820152606401610b85565b6000838152600c6020908152604080832080546001600160a01b0319166001600160a01b0387169081179091558352600d9091528120805460019290610e0090849061ffff166158ee565b825461ffff9182166101009390930a9283029190920219909116179055506001600160a01b0382166000818152600e6020908152604080832060030154878452600f835292819020929092558151868152908101929092527fb024739a6a34bded9b50497210d3eb50ccd1b03ddf8ab3738496175272c5bb6591015b60405180910390a1505050565b6005546040516331a9108f60e11b815260048101839052829133916001600160a01b0390911690636352211e90602401602060405180830381865afa158015610ed6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610efa9190615897565b6001600160a01b031614610f3c5760405162461bcd60e51b81526020600482015260096024820152682737ba1037bbb732b960b91b6044820152606401610b85565b6000828152600f60209081526040808320600c909252909120546001600160a01b031615610fdb5780546000848152600c60209081526040808320546001600160a01b03168352600e909152902060030154610f989190615910565b816001016000828254610fab9190615923565b90915550506000838152600c60209081526040808320546001600160a01b03168352600e90915290206003015481555b806002015481600101541161101e5760405162461bcd60e51b8152602060048201526009602482015268139bc81c995dd85c9960ba1b6044820152606401610b85565b6000816002015482600101546110349190615910565b6001830154600284015560035460405163c917e63560e01b8152336004820152602481018390529192506001600160a01b03169063c917e63590604401600060405180830381600087803b15801561108b57600080fd5b505af115801561109f573d6000803e3d6000fd5b505060408051878152602081018590527f8b0944629ca33aba8dd5f33f7f8220efe77a2d5548a1651362c856c5ea586a6593500190505b60405180910390a150505050565b600a602052816000526040600020818154811061110057600080fd5b9060005260206000209060089182820401919006600402915091509054906101000a900463ffffffff1681565b6111373382613b66565b50565b6000546201000090046001600160a01b031633146111865760405162461bcd60e51b81526020600482015260096024820152682737ba1030b236b4b760b91b6044820152606401610b85565b600680546001600160a01b0319166001600160a01b0383169081179091556040519081527f6bdb6edfe23f0a97be9dda487ad7ee5fc6590e2058caf57a9abd481269d4a2e3906020015b60405180910390a150565b3360009081526001602052604090205460ff166112245760405162461bcd60e51b81526020600482015260076024820152664e6f742074656560c81b6044820152606401610b85565b336000908152601060209081526040918290206006548351636554f75d60e11b8152935191936001600160a01b039091169263caa9eeba926004808401938290030181865afa15801561127b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061129f9190615936565b60028201546112ae9042615910565b10156112e65760405162461bcd60e51b81526020600482015260076024820152664c6f636b696e6760c81b6044820152606401610b85565b600181018054825460ff191683556000909155600354604051633c1e15c160e01b8152336004820152602481018390526001600160a01b0390911690633c1e15c190604401600060405180830381600087803b15801561134557600080fd5b505af1158015611359573d6000803e3d6000fd5b505060408051338152602081018590527f8348f568c1ea68de8db3504a38def7ab3a5c7d12d8f913fff88c0bf7e10794c993500190505b60405180910390a15050565b6001600160a01b0381166000908152600e6020526040902054600160401b900460ff166113fb5760405162461bcd60e51b815260206004820152600d60248201526c496e616374697665206e6f646560981b6044820152606401610b85565b61140481613d0b565b61113781613f44565b600062015180600360009054906101000a90046001600160a01b03166001600160a01b031663e6fd48bc6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611466573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061148a9190615936565b6114949042615910565b61149e9190615965565b905090565b6007546001600160a01b031633146114ec5760405162461bcd60e51b815260206004820152600c60248201526b2737ba1031b0b93b103b393360a11b6044820152606401610b85565b600081511161153d5760405162461bcd60e51b815260206004820152601160248201527f57726f6e672072616e646f6d576f7264730000000000000000000000000000006044820152606401610b85565b6006546040805163600d094960e01b815290516000926001600160a01b03169163600d09499160048083019260209291908290030181865afa158015611587573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115ab9190615936565b6115b59042615923565b905060006115dc836000815181106115cf576115cf615979565b6020026020010151614165565b6009805491925063ffffffff9091169060006115f78361598f565b82546101009290920a63ffffffff818102199093169183160217909155600954166000908152600a60209081526040909120835161163a9350909184019061509d565b5060008481526012602090815260408083208054825181850281018501909352808352919290919083018282801561169157602002820191906000526020600020905b81548152602001906001019080831161167d575b5050505050905060005b8151811015611710576000601160008484815181106116bc576116bc615979565b602090810291909101810151825281019190915260400160002060018101869055600954815463ffffffff60d01b191663ffffffff909116600160d01b021790555080611708816159b2565b91505061169b565b507f455929120054502ca2ea8194b26e7bb3acb631d30177f6881ffa70581abd4a13858385604051611744939291906159cb565b60405180910390a15050505050565b604080517f05f6695b92c38b9ece7f9daacd48393686400addab760204389a9a8a3169913660208201526001600160a01b03881691810191909152606081018690526000906080016040516020818303038152906040528051906020012090506117c186828787878761432e565b6117cb8588614460565b50505050505050565b600081815260116020526040902054600160f01b900460ff166118275760405162461bcd60e51b815260206004820152600b60248201526a139bdd081cdb185cda195960aa1b6044820152606401610b85565b600081815260136020908152604080832033845290915290205460ff1661187f5760405162461bcd60e51b815260206004820152600c60248201526b139bdd081d995c9a599a595960a21b6044820152606401610b85565b33600090815260156020908152604080832084845290915290205460ff16156118ea5760405162461bcd60e51b815260206004820152600f60248201527f416c726561647920636c61696d656400000000000000000000000000000000006044820152606401610b85565b600654604080516339e7b3af60e11b815290516000926001600160a01b0316916373cf675e9160048083019260209291908290030181865afa158015611934573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119589190615936565b600354604051633c1e15c160e01b8152336004820152602481018390529192506001600160a01b031690633c1e15c190604401600060405180830381600087803b1580156119a557600080fd5b505af11580156119b9573d6000803e3d6000fd5b5050336000818152601560209081526040808320888452825291829020805460ff19166001179055815192835282018590527fc1c7a9f718ce1c74f4b880bf361273760fb29077721becb28e8260c44cd6a9ea9350019050611390565b6000546201000090046001600160a01b03163314611a625760405162461bcd60e51b81526020600482015260096024820152682737ba1030b236b4b760b91b6044820152606401610b85565b600780546001600160a01b0319166001600160a01b0383169081179091556040519081527f891cc243fd5e5ff1f52df8dae4e723c3a5ccb1ac26346ca41c55b82f7f7162eb906020016111d0565b6000546201000090046001600160a01b03163314611afc5760405162461bcd60e51b81526020600482015260096024820152682737ba1030b236b4b760b91b6044820152606401610b85565b6001600160a01b03919091166000908152600160205260409020805460ff1916911515919091179055565b6005546040516331a9108f60e11b815260048101839052829133916001600160a01b0390911690636352211e90602401602060405180830381865afa158015611b74573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b989190615897565b6001600160a01b031614611bda5760405162461bcd60e51b81526020600482015260096024820152682737ba1037bbb732b960b91b6044820152606401610b85565b6000828152600c60205260409020546001600160a01b031680611c2e5760405162461bcd60e51b815260206004820152600c60248201526b4e6f742064656c656761746560a01b6044820152606401610b85565b6000838152600c6020908152604080832080546001600160a01b03191690556001600160a01b0384168352600d9091528120805460019290611c7590849061ffff16615a28565b825461ffff9182166101009390930a9283029190920219909116179055506001600160a01b0381166000908152600e602052604090205460ff600160401b909104168015611cdd57506001600160a01b0381166000908152600d602052604090205461ffff16155b15611ceb57611ceb81614772565b6000838152600f6020908152604080832080546001600160a01b0386168552600e909352922060030154611d1f9190615910565b816001016000828254611d329190615923565b9091555050604080518581526001600160a01b03841660208201527fc29e8cd64b240d860ccd82cacb9453b926a03ba48a57753a663a6facde47105d91016110d6565b6111373382614460565b604080517f9ee3a52f2e148cbdce38a22a9a0aef5c9564b2c41cf59aa7d874db978602862d6020820152908101869052600090606001604051602081830303815290604052805190602001209050611ddb86828787878761432e565b611de485614772565b505050505050565b604080517f0515535cbd55f5cae822b3790792a15a624eab9dbd2eb518ade8888ae01dcaf8602082015263ffffffff88169181019190915260608101869052600090608001604051602081830303815290604052805190602001209050611e5786828787878761432e565b6117cb8588613b66565b6005546040516331a9108f60e11b815260048101849052839133916001600160a01b0390911690636352211e90602401602060405180830381865afa158015611eae573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ed29190615897565b6001600160a01b031614611f145760405162461bcd60e51b81526020600482015260096024820152682737ba1037bbb732b960b91b6044820152606401610b85565b6000838152600c60205260409020546001600160a01b03168015801590611f4d5750826001600160a01b0316816001600160a01b031614155b611fad5760405162461bcd60e51b815260206004820152602b60248201527f43616e6e6f7420726564656c6567617465206f7220726564656c65676174652060448201526a746f2073616d65206f6e6560a81b6064820152608401610b85565b600660009054906101000a90046001600160a01b03166001600160a01b03166315fbbc9d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612000573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061202491906158b4565b6001600160a01b0384166000908152600d602052604090205461ffff9182169116106120855760405162461bcd60e51b815260206004820152601060248201526f4d6178206e6f6465207765696768747360801b6044820152606401610b85565b6000848152600c6020908152604080832080546001600160a01b0319166001600160a01b0388169081179091558352600d90915281208054600192906120d090849061ffff166158ee565b82546101009290920a61ffff8181021990931691831602179091556001600160a01b0383166000908152600d60205260408120805460019450909261211791859116615a28565b825461ffff9182166101009390930a9283029190920219909116179055506001600160a01b0381166000908152600e602052604090205460ff600160401b90910416801561217f57506001600160a01b0381166000908152600d602052604090205461ffff16155b1561218d5761218d81614772565b6000848152600f6020908152604080832080546001600160a01b0386168552600e9093529220600301546121c19190615910565b8160010160008282546121d49190615923565b90915550506001600160a01b0384166000818152600e60209081526040918290206003015484558151888152908101929092527fa6494190270aff4668179671a838f4079a7b83070d0dc4aafc1252956fa93d139101611744565b3360009081526001602052604090205460ff166122785760405162461bcd60e51b81526020600482015260076024820152664e6f742074656560c81b6044820152606401610b85565b600480546003546040516323b872dd60e01b815233938101939093526001600160a01b0390811660248401526044830184905216906323b872dd906064016020604051808303816000875af11580156122d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122f99190615a43565b50600354604051630cd34c3760e21b8152600481018390526001600160a01b039091169063334d30dc90602401600060405180830381600087803b15801561234057600080fd5b505af1158015612354573d6000803e3d6000fd5b5050336000908152601060205260408120600181018054919450859350919061237e908490615923565b909155505060065460408051630f185d2f60e21b815290516001600160a01b0390921691633c6174bc916004808201926020929091908290030181865afa1580156123cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123f19190615936565b816001015410156124445760405162461bcd60e51b815260206004820152601a60248201527f4e6f74206d656574206d696e5465655374616b65416d6f756e740000000000006044820152606401610b85565b805460ff1916600117815560408051338152602081018490527f1b457cf8cf55b14c9e9f6ca27cea4ca8eb08d8d1e032db6923681cf45093bc659101611390565b3360009081526001602052604090205460ff166124ce5760405162461bcd60e51b81526020600482015260076024820152664e6f742074656560c81b6044820152606401610b85565b336000908152601060205260409020805460ff166125185760405162461bcd60e51b8152602060048201526007602482015266125b9d985b1a5960ca1b6044820152606401610b85565b4260028201556007546040805163e0c8628960e01b815290516000926001600160a01b03169163e0c86289916004808301926020929190829003018187875af1158015612569573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061258d9190615936565b90506000835167ffffffffffffffff8111156125ab576125ab615293565b6040519080825280602002602001820160405280156125d4578160200160208202803683370190505b50905060005b84518110156126c75760008582815181106125f7576125f7615979565b602090810291909101810151805190820120600081815260119092526040909120549091506001600160a01b0316156126725760405162461bcd60e51b815260206004820152601060248201527f416c7265616479207265706f72746564000000000000000000000000000000006044820152606401610b85565b600081815260116020526040902080546001600160a01b03191633178155835182908590859081106126a6576126a6615979565b602002602001018181525050505080806126bf906159b2565b9150506125da565b50600082815260126020908152604090912082516126e79284019061514c565b507f89a3b784b99180438f3b2027aa89e97c3c3ed66e8dc78a555d7013b39caf1a89338286856040516110d69493929190615a60565b6040805160008152602081019091526060908267ffffffffffffffff81111561274857612748615293565b60405190808252806020026020018201604052801561277b57816020015b60608152602001906001900390816127665790505b50915060005b83811015612808576127d83086868481811061279f5761279f615979565b90506020028101906127b19190615b14565b856040516020016127c493929190615b62565b60405160208183030381529060405261495b565b8382815181106127ea576127ea615979565b60200260200101819052508080612800906159b2565b915050612781565b50505b92915050565b600083815260116020908152604091829020825160e08101845281546001600160a01b0381168252600160a01b810461ffff90811694830194909452600160b01b8104841694820194909452600160c01b84049092166060830152600160d01b830463ffffffff166080830152600160f01b90920460ff16151560a0820152600182015460c08201526128a3906149d1565b60008080805b85811015612bf95760007fd781bff21a117fe0ccb0248062cc8bd0fb2afd7d5f49c20f0bb26cf724e0ef23898989858181106128e7576128e7615979565b6128fd92602060c0909202019081019150615b89565b8a8a8681811061290f5761290f615979565b905060c0020160200160208101906129279190615278565b60405160200161293a9493929190615bdc565b6040516020818303038152906040528051906020012090506129e842828a8a8681811061296957612969615979565b905060c00201604001602081019061298191906151f6565b8b8b8781811061299357612993615979565b905060c0020160600160208101906129ab9190615c0c565b8c8c888181106129bd576129bd615979565b905060c00201608001358d8d898181106129d9576129d9615979565b905060c0020160a0013561432e565b612a45898989858181106129fe576129fe615979565b905060c002016040016020810190612a1691906151f6565b8a8a86818110612a2857612a28615979565b905060c002016020016020810190612a409190615278565b614ac3565b6000888884818110612a5957612a59615979565b612a6f92602060c0909202019081019150615b89565b6002811115612a8057612a80615ba4565b03612a975784612a8f81615c27565b955050612b83565b6001888884818110612aab57612aab615979565b612ac192602060c0909202019081019150615b89565b6002811115612ad257612ad2615ba4565b03612ae95783612ae181615c27565b945050612b83565b6002888884818110612afd57612afd615979565b612b1392602060c0909202019081019150615b89565b6002811115612b2457612b24615ba4565b03612b3b5782612b3381615c27565b935050612b83565b60405162461bcd60e51b815260206004820152601960248201527f556e6b6e6f776e204174746573746174696f6e526573756c74000000000000006044820152606401610b85565b60008981526013602052604081206001918a8a86818110612ba657612ba6615979565b905060c002016040016020810190612bbe91906151f6565b6001600160a01b031681526020810191909152604001600020805460ff19169115159190911790555080612bf1816159b2565b9150506128a9565b50835483908590601490612c19908490600160a01b900461ffff166158ee565b92506101000a81548161ffff021916908361ffff160217905550818460000160168282829054906101000a900461ffff16612c5491906158ee565b92506101000a81548161ffff021916908361ffff160217905550808460000160188282829054906101000a900461ffff16612c8f91906158ee565b92506101000a81548161ffff021916908361ffff1602179055507f16093eb01e7213cc4445469a1c8ded426fe3c5f18a72b6c4151b07dd4916cdf5878787604051612cdc93929190615c3e565b60405180910390a150505050505050565b604080517fcec305acdefa978a0c15fc72c2f70dfec101c80bbebe137b563678828e5f7d7e60208201526001600160a01b0388169181019190915260608101869052600090608001604051602081830303815290604052805190602001209050612d5b86828787878761432e565b6001600160a01b038581166000818152600e602090815260409182902060050180546001600160a01b031916948c16948517905581519283528201929092527f1454d6e227f2d31063b1c094100ff944573c2c21ddb0972ffc83313638583f8f9101612cdc565b60126020528160005260406000208181548110612dde57600080fd5b90600052602060002001600091509150505481565b600054610100900460ff1615808015612e135750600054600160ff909116105b80612e2d5750303b158015612e2d575060005460ff166001145b612e9f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610b85565b6000805460ff191660011790558015612ec2576000805461ff0019166101001790555b600480546001600160a01b038088166001600160a01b031992831617909255600580548784169083161790556003805492861692909116919091179055612f0833614d0e565b604080518082018252600f81527f50726f746f636f6c5365727669636500000000000000000000000000000000006020918201528151808301835260058152640312e302e360dc1b9082015281517fc2f8787176b8ac6bf7215b4adcc1e069bf4ab82d9ab1df05a57a91d425935b6e818301527fee7ec1d45dbf7d40ce3d51a97619883c7aab09ac9d5aa566f33615283a410c19818401527f06c015bd22b4c69690933c1058878ebdfef31f9aaae40bbe86d8a09fe1b2972c606082015260808082018690528351808303909101815260a09091019092528151910120600255801561302a576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602001611744565b5050505050565b61303a33614772565b565b600062015180600360009054906101000a90046001600160a01b03166001600160a01b031663e6fd48bc6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613095573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130b99190615936565b6130c39042615910565b61149e9190615cee565b6000828152601160205260409020600181015442106131195760405162461bcd60e51b8152602060048201526008602482015267446561646c696e6560c01b6044820152606401610b85565b6001600160a01b0384166000908152600e60209081526040808320601483528184208785529092529091205460ff16156131875760405162461bcd60e51b815260206004820152600f60248201526e105b1c9958591e481cdb185cda1959608a1b6044820152606401610b85565b80548254600160d01b900463ffffffff9081166000908152600a60205260409020805492821692909186169081106131c1576131c1615979565b6000918252602090912060088204015460079091166004026101000a900463ffffffff161461321f5760405162461bcd60e51b815260206004820152600a6024820152692737ba1031b437b9b2b760b11b6044820152606401610b85565b60008481526013602090815260408083206001600160a01b038916845290915290205460ff16156132825760405162461bcd60e51b815260206004820152600d60248201526c139bd919481d995c9a599a5959609a1b6044820152606401610b85565b600654604080516344677ce160e11b815290516000926001600160a01b0316916388cef9c29160048083019260209291908290030181865afa1580156132cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132f09190615936565b82549091506001908390600d90613319908490600160681b900467ffffffffffffffff16615d02565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550808260010160008282546133539190615d23565b90915550506001600160a01b03868116600090815260146020908152604080832089845290915290819020805460ff19166001179055600354905163c917e63560e01b81523360048201526024810184905291169063c917e63590604401600060405180830381600087803b1580156133cb57600080fd5b505af11580156133df573d6000803e3d6000fd5b50505050600660009054906101000a90046001600160a01b03166001600160a01b0316630b097ac36040518163ffffffff1660e01b8152600401602060405180830381865afa158015613436573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061345a9190615d43565b825467ffffffffffffffff918216600160681b909104909116106134815761348186614772565b604080516001600160a01b0388168152602081018790529081018290527f2ec1d25826f3279dd30b2af581bf1fb9fdca750ac9df9cc193d194704ea0d76d9060600160405180910390a1505050505050565b60008181526011602052604090208054600160f01b900460ff161561352c5760405162461bcd60e51b815260206004820152600f60248201526e105b1c9958591e481cdb185cda1959608a1b6044820152606401610b85565b4281600101541061356a5760405162461bcd60e51b8152602060048201526008602482015267446561646c696e6560c01b6044820152606401610b85565b805461ffff600160c01b82048116600160a01b90920416106135ce5760405162461bcd60e51b815260206004820152601160248201527f56616c6964206174746573746174696f6e0000000000000000000000000000006044820152606401610b85565b80546001600160a01b03166000818152601060205260409020805460ff166136265760405162461bcd60e51b815260206004820152600b60248201526a496e76616c69642074656560a81b6044820152606401610b85565b825460009061ffff600160c01b820481169161365391600160b01b8204811691600160a01b9004166158ee565b61365d91906158ee565b61ffff16600660009054906101000a90046001600160a01b03166001600160a01b03166373cf675e6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156136b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136d89190615936565b6136e29190615d6d565b9050808260010160008282546136f89190615910565b909155505060065460408051630f185d2f60e21b815290516001600160a01b0390921691633c6174bc916004808201926020929091908290030181865afa158015613747573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061376b9190615936565b8260010154101561377f57815460ff191682555b835460ff60f01b1916600160f01b178455604080516001600160a01b0385168152602081018790529081018290527fcb2bfd0c073ab6c31355ed348da3a1b661a3a578be5a450eaa8d420a9323efaf90606001611744565b600083815260116020908152604091829020825160e08101845281546001600160a01b0381168252600160a01b810461ffff90811694830194909452600160b01b8104841694820194909452600160c01b84049092166060830152600160d01b830463ffffffff166080830152600160f01b90920460ff16151560a0820152600182015460c0820152613869906149d1565b613874843385614ac3565b600082600281111561388857613888615ba4565b036138cc57805460019082906014906138ad908490600160a01b900461ffff166158ee565b92506101000a81548161ffff021916908361ffff16021790555061393e565b60018260028111156138e0576138e0615ba4565b0361390557805460019082906016906138ad908490600160b01b900461ffff166158ee565b600282600281111561391957613919615ba4565b03612b3b57805460019082906018906138ad908490600160c01b900461ffff166158ee565b600084815260136020908152604080832033808552925291829020805460ff1916600117905590517f52b7fb186080c2874dd6f580c1c5c01ddd95c9c2f414fd8c318a09ead3dc5536916110d69187908690615d84565b6001600160a01b0381166000908152600e60205260409020805463ffffffff166139f05760405162461bcd60e51b815260206004820152600c60248201526b2737ba103932b3b4b9ba32b960a11b6044820152606401610b85565b336001600160a01b0383161480613a13575060058101546001600160a01b031633145b613a4e5760405162461bcd60e51b815260206004820152600c60248201526b43616e6e6f7420636c61696d60a01b6044820152606401610b85565b613a5782613f44565b8060020154816001015413613a9a5760405162461bcd60e51b8152602060048201526009602482015268139bc81c995dd85c9960ba1b6044820152606401610b85565b600081600201548260010154613ab09190615910565b6001830154600284015560035460405163c917e63560e01b8152336004820152602481018390529192506001600160a01b03169063c917e63590604401600060405180830381600087803b158015613b0757600080fd5b505af1158015613b1b573d6000803e3d6000fd5b5050604080516001600160a01b03871681523360208201529081018490527fea0babbef5f952932a55a6932b70dee7422d83300a5a59fae102bcf1e113632492506060019050610e7c565b6001600160a01b0382166000908152600e60205260409020805463ffffffff16613bc15760405162461bcd60e51b815260206004820152600c60248201526b2737ba103932b3b4b9ba32b960a11b6044820152606401610b85565b6006546040805163512ab48360e01b8152905142926001600160a01b03169163512ab4839160048083019260209291908290030181865afa158015613c0a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613c2e9190615936565b8260060154613c3d9190615923565b10613c9f5760405162461bcd60e51b815260206004820152602c60248201527f4e6f74206d656574206d696e20636f6d6d697373696f6e2072617465206d6f6460448201526b1a599e481a5b9d195c9d985b60a21b6064820152608401610b85565b60058101805463ffffffff60a01b1916600160a01b63ffffffff851690810291909117909155426006830155604080516001600160a01b038616815260208101929092527fb259504a8d6c681b25c45cd2a3825d7b986058852f5d01180dec6f9574f781f99101610e7c565b6000613d1561140d565b6001600160a01b038316600090815260176020908152604080832063ffffffff80861685529252909120549192501615613d4d575050565b600660009054906101000a90046001600160a01b03166001600160a01b031663729fe45f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613da0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613dc49190615936565b6001600160a01b0383166000908152600e6020526040902060040154613dea9042615910565b10158015613e755750600660009054906101000a90046001600160a01b03166001600160a01b031663729fe45f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613e46573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613e6a9190615936565b613e7261303c565b10155b15613f40576001600160a01b0382166000908152600d60209081526040808320546017835281842063ffffffff80871686529352908320805461ffff9092169390929091613ec591859116615db0565b82546101009290920a63ffffffff8181021990931691831602179091556001600160a01b0384166000908152600d6020908152604080832054868516845260169092528220805461ffff909216945092613f2191859116615db0565b92506101000a81548163ffffffff021916908363ffffffff1602179055505b5050565b6000613f4e61140d565b6001600160a01b0383166000908152600e60205260409020909150613f74600183615dcd565b815463ffffffff918216600160481b90910490911603613f9357505050565b8054600090613fb090600160481b900463ffffffff166001615db0565b90505b8263ffffffff168163ffffffff16101561412f5763ffffffff808216600090815260166020526040902054161561411d5763ffffffff818116600081815260166020526040808220546003549151631e88e34160e11b81526004810194909452919391909116916001600160a01b0390911690633d11c68290602401602060405180830381865afa15801561404c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906140709190615936565b61407a9190615965565b6001600160a01b038616600090815260176020908152604080832063ffffffff808816855292528220546005870154939450919291811691612710916140c891600160a01b90041685615d6d565b6140d29190615965565b6140dc9190615d6d565b9050808460010160008282546140f29190615dea565b9091555061410290508183615910565b8460030160008282546141159190615923565b909155505050505b806141278161598f565b915050613fb3565b5061413b600183615dcd565b815463ffffffff91909116600160481b026cffffffff000000000000000000199091161790555050565b600854606090600a106141f65760088054806020026020016040519081016040528092919081815260200182805480156141ea57602002820191906000526020600020906000905b82829054906101000a900463ffffffff1663ffffffff16815260200190600401906020826003010492830192600103820291508084116141ad5790505b50505050509050919050565b6008546000906142069084615cee565b600854909150600a90606410156142295760085461422690600a90615965565b90505b60008163ffffffff1667ffffffffffffffff81111561424a5761424a615293565b604051908082528060200260200182016040528015614273578160200160208202803683370190505b50905060005b8263ffffffff168163ffffffff1610156143255760088463ffffffff16815481106142a6576142a6615979565b90600052602060002090600891828204019190066004029054906101000a900463ffffffff16828263ffffffff16815181106142e4576142e4615979565b63ffffffff90921660209283029190910190910152600854614307856001615db0565b6143119190615e0a565b93508061431d8161598f565b915050614279565b50949350505050565b428610156143685760405162461bcd60e51b8152602060048201526007602482015266115e1c1a5c995960ca1b6044820152606401610b85565b60025460405161190160f01b602082015260228101919091526042810186905260009060620160408051601f1981840301815282825280516020918201206000845290830180835281905260ff8716918301919091526060820185905260808201849052915060019060a0016020604051602081039080840390855afa1580156143f6573d6000803e3d6000fd5b505050602060405103516001600160a01b0316856001600160a01b0316146117cb5760405162461bcd60e51b815260206004820152601060248201527f7369676e6572206e6f74206d61746368000000000000000000000000000000006044820152606401610b85565b6001600160a01b0382166000908152600d602052604090205461ffff166144c95760405162461bcd60e51b815260206004820152601460248201527f4e6f2064656c65676174696f6e576569676874730000000000000000000000006044820152606401610b85565b6001600160a01b0382166000908152600e602052604081205463ffffffff1690036144f7576144f782614d79565b6001600160a01b0382166000908152600e6020526040902054600160401b900460ff16156145575760405162461bcd60e51b815260206004820152600d60248201526c20b63932b0b23c9032b73a32b960991b6044820152606401610b85565b600660009054906101000a90046001600160a01b03166001600160a01b031663b416dd856040518163ffffffff1660e01b8152600401602060405180830381865afa1580156145aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906145ce9190615936565b600854101561465f576001600160a01b0382166000908152600e60205260408120546008805460018082018355938290527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee382820401805463ffffffff60079093166004026101000a83810219909116929094169093021790915554613f4091849161465a9190615dcd565b614f08565b6001600160a01b038082166000908152600d602052604080822054928516825290205461ffff9182169116116146c65760405162461bcd60e51b815260206004820152600c60248201526b4c657373207765696768747360a01b6044820152606401610b85565b6001600160a01b038082166000908152600e6020526040808220928516825290205481546008805463ffffffff938416939192640100000000900490911690811061471357614713615979565b90600052602060002090600891828204019190066004026101000a81548163ffffffff021916908363ffffffff160217905550614764838260000160049054906101000a900463ffffffff16614f08565b61476d82614f8c565b505050565b6001600160a01b0381166000908152600e6020526040902054600160401b900460ff166147d05760405162461bcd60e51b815260206004820152600c60248201526b105b1c9958591e48195e1a5d60a21b6044820152606401610b85565b6008546147df90600190615910565b6001600160a01b0382166000908152600e6020526040902054640100000000900463ffffffff161461491357600880546000919061481f90600190615910565b8154811061482f5761482f615979565b6000918252602080832060088084049091015460079093166004026101000a90920463ffffffff908116808552600b83526040808620546001600160a01b039081168752600e909452808620805494891687529520548454919650928216939264010000000090049091169081106148a9576148a9615979565b600091825260208083206008830401805460079093166004026101000a63ffffffff8181021990941695841602949094179093556001600160a01b0386168252600e9092526040902054825467ffffffff0000000019166401000000009182900490921602179055505b600880548061492457614924615e2d565b600082815260209020600860001990920191820401805463ffffffff600460078516026101000a0219169055905561113781614f8c565b6060600080846001600160a01b0316846040516149789190615e43565b600060405180830381855af49150503d80600081146149b3576040519150601f19603f3d011682016040523d82523d6000602084013e6149b8565b606091505b50915091506149c8858383615015565b95945050505050565b80516001600160a01b0316614a285760405162461bcd60e51b815260206004820152601560248201527f4174746573746174696f6e206e6f7420657869737400000000000000000000006044820152606401610b85565b428160c0015111614a7b5760405162461bcd60e51b815260206004820152600f60248201527f446561646c696e652070617373656400000000000000000000000000000000006044820152606401610b85565b6000816080015163ffffffff16116111375760405162461bcd60e51b815260206004820152600b60248201526a2bb0b4ba34b733903b393360a91b6044820152606401610b85565b6000838152601160209081526040808320815160e08101835281546001600160a01b03808216835261ffff600160a01b8304811684880152600160b01b8304811684870152600160c01b830416606084015263ffffffff600160d01b830416608084015260ff600160f01b9092048216151560a084015260019093015460c08301529187168552600e9093529220549091600160401b90910416614b965760405162461bcd60e51b815260206004820152600a6024820152694e6f742061637469766560b01b6044820152606401610b85565b6001600160a01b0383166000908152600d602052604090205461ffff16614bec5760405162461bcd60e51b815260206004820152600a6024820152694e6f207765696768747360b01b6044820152606401610b85565b60008481526013602090815260408083206001600160a01b038716845290915290205460ff1615614c5f5760405162461bcd60e51b815260206004820152600e60248201527f416c7265616479207665726966790000000000000000000000000000000000006044820152606401610b85565b6001600160a01b0383166000908152600e6020908152604080832054608085015163ffffffff9081168552600a909352922080549282169290918516908110614caa57614caa615979565b6000918252602090912060088204015460079091166004026101000a900463ffffffff1614614d085760405162461bcd60e51b815260206004820152600a6024820152692737ba1031b437b9b2b760b11b6044820152606401610b85565b50505050565b600054610100900460ff16610b8e5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608401610b85565b60075463ffffffff600160a01b909104811610614dd85760405162461bcd60e51b815260206004820152600e60248201527f496e646578206f766572666c6f770000000000000000000000000000000000006044820152606401610b85565b60078054600160a01b900463ffffffff16906014614df58361598f565b825463ffffffff9182166101009390930a9283029282021916919091179091556001600160a01b0383166000908152600e60205260409020600754815467ffffffffffffffff1916600160a01b9091049092169190911767ffffffff0000000017815590506001614e6461140d565b614e6e9190615dcd565b81546cffffffff0000000000000000001916600160481b63ffffffff9283160217825560078054600160a01b9081900483166000908152600b602090815260409182902080546001600160a01b0319166001600160a01b0389169081179091559354825194855292909204909316908201527fdcda586e1f19498517f947f53b7416526efbc24328864eb9071a9315ff8a86a99101611390565b6001600160a01b0382166000818152600e60209081526040918290208054600160401b68ffffffffff000000001990911664010000000063ffffffff88160268ff0000000000000000191617178155426004820155915192835290917f1b3699bd6aa5abdc4e768415f5023e2c35280c273edb4f358bad0c05f7b2701d9101610e7c565b6001600160a01b0381166000908152600e60205260409020805474ffffffffffffffff00000000ffffffffff00000000191667ffffffff00000000178155614fd382613d0b565b614fdc82613f44565b6040516001600160a01b03831681527f8a0859fa4a2e331800d512db6925d210facda82733207cb9fe49e7da954fc4aa90602001611390565b60608261502a5761502582615074565b61506d565b815115801561504157506001600160a01b0384163b155b1561506a57604051639996b31560e01b81526001600160a01b0385166004820152602401610b85565b50805b9392505050565b8051156150845780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b8280548282559060005260206000209060070160089004810192821561513c5791602002820160005b8382111561510a57835183826101000a81548163ffffffff021916908363ffffffff16021790555092602001926004016020816003010492830192600103026150c6565b801561513a5782816101000a81549063ffffffff021916905560040160208160030104928301926001030261510a565b505b50615148929150615187565b5090565b82805482825590600052602060002090810192821561513c579160200282015b8281111561513c57825182559160200191906001019061516c565b5b808211156151485760008155600101615188565b6000602082840312156151ae57600080fd5b5035919050565b6001600160a01b038116811461113757600080fd5b600080604083850312156151dd57600080fd5b82356151e8816151b5565b946020939093013593505050565b60006020828403121561520857600080fd5b813561506d816151b5565b6000806040838503121561522657600080fd5b823591506020830135615238816151b5565b809150509250929050565b803563ffffffff8116811461525757600080fd5b919050565b6000806040838503121561526f57600080fd5b6151e883615243565b60006020828403121561528a57600080fd5b61506d82615243565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff811182821017156152d2576152d2615293565b604052919050565b600067ffffffffffffffff8211156152f4576152f4615293565b5060051b60200190565b6000806040838503121561531157600080fd5b8235915060208084013567ffffffffffffffff81111561533057600080fd5b8401601f8101861361534157600080fd5b803561535461534f826152da565b6152a9565b81815260059190911b8201830190838101908883111561537357600080fd5b928401925b8284101561539157833582529284019290840190615378565b80955050505050509250929050565b803560ff8116811461525757600080fd5b60008060008060008060c087890312156153ca57600080fd5b86356153d5816151b5565b95506020870135945060408701356153ec816151b5565b93506153fa606088016153a0565b92506080870135915060a087013590509295509295509295565b801515811461113757600080fd5b6000806040838503121561543557600080fd5b8235615440816151b5565b9150602083013561523881615414565b600080600080600060a0868803121561546857600080fd5b85359450602086013561547a816151b5565b9350615488604087016153a0565b94979396509394606081013594506080013592915050565b60008060008060008060c087890312156154b957600080fd5b6153d587615243565b600060208083850312156154d557600080fd5b823567ffffffffffffffff808211156154ed57600080fd5b8185019150601f868184011261550257600080fd5b823561551061534f826152da565b81815260059190911b8401850190858101908983111561552f57600080fd5b8686015b838110156155bb5780358681111561554b5760008081fd5b8701603f81018c1361555d5760008081fd5b8881013560408882111561557357615573615293565b615584828901601f19168c016152a9565b8281528e828486010111156155995760008081fd5b828285018d83013760009281018c019290925250845250918701918701615533565b509998505050505050505050565b600080604083850312156155dc57600080fd5b82356155e7816151b5565b91506155f560208401615243565b90509250929050565b6000806020838503121561561157600080fd5b823567ffffffffffffffff8082111561562957600080fd5b818501915085601f83011261563d57600080fd5b81358181111561564c57600080fd5b8660208260051b850101111561566157600080fd5b60209290920196919550909350505050565b60005b8381101561568e578181015183820152602001615676565b50506000910152565b600081518084526156af816020860160208601615673565b601f01601f19169290920160200192915050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b8281101561571857603f19888603018452615706858351615697565b945092850192908501906001016156ea565b5092979650505050505050565b60008060006040848603121561573a57600080fd5b83359250602084013567ffffffffffffffff8082111561575957600080fd5b818601915086601f83011261576d57600080fd5b81358181111561577c57600080fd5b87602060c08302850101111561579157600080fd5b6020830194508093505050509250925092565b600080604083850312156157b757600080fd5b50508035926020909101359150565b600080600080608085870312156157dc57600080fd5b84356157e7816151b5565b935060208501356157f7816151b5565b92506040850135615807816151b5565b9396929550929360600135925050565b60008060006060848603121561582c57600080fd5b8335615837816151b5565b92506020840135915061584c60408501615243565b90509250925092565b80356003811061525757600080fd5b60008060006060848603121561587957600080fd5b8335925061588960208501615243565b915061584c60408501615855565b6000602082840312156158a957600080fd5b815161506d816151b5565b6000602082840312156158c657600080fd5b815161ffff8116811461506d57600080fd5b634e487b7160e01b600052601160045260246000fd5b61ffff818116838216019080821115615909576159096158d8565b5092915050565b8181038181111561280b5761280b6158d8565b8082018082111561280b5761280b6158d8565b60006020828403121561594857600080fd5b5051919050565b634e487b7160e01b600052601260045260246000fd5b6000826159745761597461594f565b500490565b634e487b7160e01b600052603260045260246000fd5b600063ffffffff8083168181036159a8576159a86158d8565b6001019392505050565b6000600182016159c4576159c46158d8565b5060010190565b6000606082018583526020606081850152818651808452608086019150828801935060005b81811015615a1257845163ffffffff16835293830193918301916001016159f0565b5050809350505050826040830152949350505050565b61ffff828116828216039080821115615909576159096158d8565b600060208284031215615a5557600080fd5b815161506d81615414565b6000608082016001600160a01b0387168352602060808185015281875180845260a086019150828901935060005b81811015615aaa57845183529383019391830191600101615a8e565b5050848103604086015286518082528282019350600581901b8201830183890160005b83811015615afb57601f19858403018752615ae9838351615697565b96860196925090850190600101615acd565b5050809550505050505082606083015295945050505050565b6000808335601e19843603018112615b2b57600080fd5b83018035915067ffffffffffffffff821115615b4657600080fd5b602001915036819003821315615b5b57600080fd5b9250929050565b828482376000838201600081528351615b7f818360208801615673565b0195945050505050565b600060208284031215615b9b57600080fd5b61506d82615855565b634e487b7160e01b600052602160045260246000fd5b60038110615bd857634e487b7160e01b600052602160045260246000fd5b9052565b8481526020810184905260808101615bf76040830185615bba565b63ffffffff8316606083015295945050505050565b600060208284031215615c1e57600080fd5b61506d826153a0565b600061ffff8083168181036159a8576159a86158d8565b838152604060208083018290528282018490526000919060609081850187855b88811015615cdf57615c7883615c7384615855565b615bba565b63ffffffff615c88858401615243565b168484015285820135615c9a816151b5565b6001600160a01b03168387015260ff615cb48387016153a0565b16838601526080828101359084015260a0808301359084015260c09283019290910190600101615c5e565b50909998505050505050505050565b600082615cfd57615cfd61594f565b500690565b67ffffffffffffffff818116838216019080821115615909576159096158d8565b8181036000831280158383131683831282161715615909576159096158d8565b600060208284031215615d5557600080fd5b815167ffffffffffffffff8116811461506d57600080fd5b808202811582820484141761280b5761280b6158d8565b6001600160a01b03841681526020810183905260608101615da86040830184615bba565b949350505050565b63ffffffff818116838216019080821115615909576159096158d8565b63ffffffff828116828216039080821115615909576159096158d8565b8082018281126000831280158216821582161715612808576128086158d8565b600063ffffffff80841680615e2157615e2161594f565b92169190910692915050565b634e487b7160e01b600052603160045260246000fd5b60008251615e55818460208701615673565b919091019291505056fea2646970667358221220c5017fba4c8d20b5b03b6f4baa650c143c02a0d43f097a6a367b97d9d4a420fe64736f6c63430008140033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

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
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// MAXUINT32 is a free data retrieval call binding the contract method 0x04901b93.
//
// Solidity: function MAX_UINT32() view returns(uint32)
func (_Contract *ContractCaller) MAXUINT32(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "MAX_UINT32")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXUINT32 is a free data retrieval call binding the contract method 0x04901b93.
//
// Solidity: function MAX_UINT32() view returns(uint32)
func (_Contract *ContractSession) MAXUINT32() (uint32, error) {
	return _Contract.Contract.MAXUINT32(&_Contract.CallOpts)
}

// MAXUINT32 is a free data retrieval call binding the contract method 0x04901b93.
//
// Solidity: function MAX_UINT32() view returns(uint32)
func (_Contract *ContractCallerSession) MAXUINT32() (uint32, error) {
	return _Contract.Contract.MAXUINT32(&_Contract.CallOpts)
}

// ActiveVrfNodeList is a free data retrieval call binding the contract method 0x0225de92.
//
// Solidity: function activeVrfNodeList(uint256 ) view returns(uint32)
func (_Contract *ContractCaller) ActiveVrfNodeList(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "activeVrfNodeList", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ActiveVrfNodeList is a free data retrieval call binding the contract method 0x0225de92.
//
// Solidity: function activeVrfNodeList(uint256 ) view returns(uint32)
func (_Contract *ContractSession) ActiveVrfNodeList(arg0 *big.Int) (uint32, error) {
	return _Contract.Contract.ActiveVrfNodeList(&_Contract.CallOpts, arg0)
}

// ActiveVrfNodeList is a free data retrieval call binding the contract method 0x0225de92.
//
// Solidity: function activeVrfNodeList(uint256 ) view returns(uint32)
func (_Contract *ContractCallerSession) ActiveVrfNodeList(arg0 *big.Int) (uint32, error) {
	return _Contract.Contract.ActiveVrfNodeList(&_Contract.CallOpts, arg0)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Contract *ContractCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Contract *ContractSession) Admin() (common.Address, error) {
	return _Contract.Contract.Admin(&_Contract.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Contract *ContractCallerSession) Admin() (common.Address, error) {
	return _Contract.Contract.Admin(&_Contract.CallOpts)
}

// AttestationVerifiedNode is a free data retrieval call binding the contract method 0xa91e5479.
//
// Solidity: function attestationVerifiedNode(bytes32 , address ) view returns(bool)
func (_Contract *ContractCaller) AttestationVerifiedNode(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "attestationVerifiedNode", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AttestationVerifiedNode is a free data retrieval call binding the contract method 0xa91e5479.
//
// Solidity: function attestationVerifiedNode(bytes32 , address ) view returns(bool)
func (_Contract *ContractSession) AttestationVerifiedNode(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Contract.Contract.AttestationVerifiedNode(&_Contract.CallOpts, arg0, arg1)
}

// AttestationVerifiedNode is a free data retrieval call binding the contract method 0xa91e5479.
//
// Solidity: function attestationVerifiedNode(bytes32 , address ) view returns(bool)
func (_Contract *ContractCallerSession) AttestationVerifiedNode(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Contract.Contract.AttestationVerifiedNode(&_Contract.CallOpts, arg0, arg1)
}

// Attestations is a free data retrieval call binding the contract method 0x940992a3.
//
// Solidity: function attestations(bytes32 ) view returns(address reporter, uint16 valid, uint16 invalid, uint16 malicious, uint32 vrfChosenID, bool slashed, uint256 deadline)
func (_Contract *ContractCaller) Attestations(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Reporter    common.Address
	Valid       uint16
	Invalid     uint16
	Malicious   uint16
	VrfChosenID uint32
	Slashed     bool
	Deadline    *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "attestations", arg0)

	outstruct := new(struct {
		Reporter    common.Address
		Valid       uint16
		Invalid     uint16
		Malicious   uint16
		VrfChosenID uint32
		Slashed     bool
		Deadline    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reporter = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Valid = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.Invalid = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.Malicious = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.VrfChosenID = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.Slashed = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.Deadline = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Attestations is a free data retrieval call binding the contract method 0x940992a3.
//
// Solidity: function attestations(bytes32 ) view returns(address reporter, uint16 valid, uint16 invalid, uint16 malicious, uint32 vrfChosenID, bool slashed, uint256 deadline)
func (_Contract *ContractSession) Attestations(arg0 [32]byte) (struct {
	Reporter    common.Address
	Valid       uint16
	Invalid     uint16
	Malicious   uint16
	VrfChosenID uint32
	Slashed     bool
	Deadline    *big.Int
}, error) {
	return _Contract.Contract.Attestations(&_Contract.CallOpts, arg0)
}

// Attestations is a free data retrieval call binding the contract method 0x940992a3.
//
// Solidity: function attestations(bytes32 ) view returns(address reporter, uint16 valid, uint16 invalid, uint16 malicious, uint32 vrfChosenID, bool slashed, uint256 deadline)
func (_Contract *ContractCallerSession) Attestations(arg0 [32]byte) (struct {
	Reporter    common.Address
	Valid       uint16
	Invalid     uint16
	Malicious   uint16
	VrfChosenID uint32
	Slashed     bool
	Deadline    *big.Int
}, error) {
	return _Contract.Contract.Attestations(&_Contract.CallOpts, arg0)
}

// CarvNft is a free data retrieval call binding the contract method 0x1cde81d5.
//
// Solidity: function carvNft() view returns(address)
func (_Contract *ContractCaller) CarvNft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "carvNft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CarvNft is a free data retrieval call binding the contract method 0x1cde81d5.
//
// Solidity: function carvNft() view returns(address)
func (_Contract *ContractSession) CarvNft() (common.Address, error) {
	return _Contract.Contract.CarvNft(&_Contract.CallOpts)
}

// CarvNft is a free data retrieval call binding the contract method 0x1cde81d5.
//
// Solidity: function carvNft() view returns(address)
func (_Contract *ContractCallerSession) CarvNft() (common.Address, error) {
	return _Contract.Contract.CarvNft(&_Contract.CallOpts)
}

// CarvToken is a free data retrieval call binding the contract method 0xe17ec726.
//
// Solidity: function carvToken() view returns(address)
func (_Contract *ContractCaller) CarvToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "carvToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CarvToken is a free data retrieval call binding the contract method 0xe17ec726.
//
// Solidity: function carvToken() view returns(address)
func (_Contract *ContractSession) CarvToken() (common.Address, error) {
	return _Contract.Contract.CarvToken(&_Contract.CallOpts)
}

// CarvToken is a free data retrieval call binding the contract method 0xe17ec726.
//
// Solidity: function carvToken() view returns(address)
func (_Contract *ContractCallerSession) CarvToken() (common.Address, error) {
	return _Contract.Contract.CarvToken(&_Contract.CallOpts)
}

// CarvVrf is a free data retrieval call binding the contract method 0x7a83fe2a.
//
// Solidity: function carvVrf() view returns(address)
func (_Contract *ContractCaller) CarvVrf(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "carvVrf")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CarvVrf is a free data retrieval call binding the contract method 0x7a83fe2a.
//
// Solidity: function carvVrf() view returns(address)
func (_Contract *ContractSession) CarvVrf() (common.Address, error) {
	return _Contract.Contract.CarvVrf(&_Contract.CallOpts)
}

// CarvVrf is a free data retrieval call binding the contract method 0x7a83fe2a.
//
// Solidity: function carvVrf() view returns(address)
func (_Contract *ContractCallerSession) CarvVrf() (common.Address, error) {
	return _Contract.Contract.CarvVrf(&_Contract.CallOpts)
}

// CheckClaimed is a free data retrieval call binding the contract method 0x35a55caa.
//
// Solidity: function checkClaimed(uint256 tokenID) view returns(bool)
func (_Contract *ContractCaller) CheckClaimed(opts *bind.CallOpts, tokenID *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "checkClaimed", tokenID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckClaimed is a free data retrieval call binding the contract method 0x35a55caa.
//
// Solidity: function checkClaimed(uint256 tokenID) view returns(bool)
func (_Contract *ContractSession) CheckClaimed(tokenID *big.Int) (bool, error) {
	return _Contract.Contract.CheckClaimed(&_Contract.CallOpts, tokenID)
}

// CheckClaimed is a free data retrieval call binding the contract method 0x35a55caa.
//
// Solidity: function checkClaimed(uint256 tokenID) view returns(bool)
func (_Contract *ContractCallerSession) CheckClaimed(tokenID *big.Int) (bool, error) {
	return _Contract.Contract.CheckClaimed(&_Contract.CallOpts, tokenID)
}

// Delegation is a free data retrieval call binding the contract method 0xb1a61dda.
//
// Solidity: function delegation(uint256 ) view returns(address)
func (_Contract *ContractCaller) Delegation(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "delegation", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xb1a61dda.
//
// Solidity: function delegation(uint256 ) view returns(address)
func (_Contract *ContractSession) Delegation(arg0 *big.Int) (common.Address, error) {
	return _Contract.Contract.Delegation(&_Contract.CallOpts, arg0)
}

// Delegation is a free data retrieval call binding the contract method 0xb1a61dda.
//
// Solidity: function delegation(uint256 ) view returns(address)
func (_Contract *ContractCallerSession) Delegation(arg0 *big.Int) (common.Address, error) {
	return _Contract.Contract.Delegation(&_Contract.CallOpts, arg0)
}

// DelegationWeights is a free data retrieval call binding the contract method 0x7334c9ba.
//
// Solidity: function delegationWeights(address ) view returns(uint16)
func (_Contract *ContractCaller) DelegationWeights(opts *bind.CallOpts, arg0 common.Address) (uint16, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "delegationWeights", arg0)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// DelegationWeights is a free data retrieval call binding the contract method 0x7334c9ba.
//
// Solidity: function delegationWeights(address ) view returns(uint16)
func (_Contract *ContractSession) DelegationWeights(arg0 common.Address) (uint16, error) {
	return _Contract.Contract.DelegationWeights(&_Contract.CallOpts, arg0)
}

// DelegationWeights is a free data retrieval call binding the contract method 0x7334c9ba.
//
// Solidity: function delegationWeights(address ) view returns(uint16)
func (_Contract *ContractCallerSession) DelegationWeights(arg0 common.Address) (uint16, error) {
	return _Contract.Contract.DelegationWeights(&_Contract.CallOpts, arg0)
}

// Eip712DomainHash is a free data retrieval call binding the contract method 0xf94dc4bc.
//
// Solidity: function eip712DomainHash() view returns(bytes32)
func (_Contract *ContractCaller) Eip712DomainHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "eip712DomainHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Eip712DomainHash is a free data retrieval call binding the contract method 0xf94dc4bc.
//
// Solidity: function eip712DomainHash() view returns(bytes32)
func (_Contract *ContractSession) Eip712DomainHash() ([32]byte, error) {
	return _Contract.Contract.Eip712DomainHash(&_Contract.CallOpts)
}

// Eip712DomainHash is a free data retrieval call binding the contract method 0xf94dc4bc.
//
// Solidity: function eip712DomainHash() view returns(bytes32)
func (_Contract *ContractCallerSession) Eip712DomainHash() ([32]byte, error) {
	return _Contract.Contract.Eip712DomainHash(&_Contract.CallOpts)
}

// GlobalDailyActiveNodes is a free data retrieval call binding the contract method 0xea715627.
//
// Solidity: function globalDailyActiveNodes(uint32 ) view returns(uint32)
func (_Contract *ContractCaller) GlobalDailyActiveNodes(opts *bind.CallOpts, arg0 uint32) (uint32, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "globalDailyActiveNodes", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GlobalDailyActiveNodes is a free data retrieval call binding the contract method 0xea715627.
//
// Solidity: function globalDailyActiveNodes(uint32 ) view returns(uint32)
func (_Contract *ContractSession) GlobalDailyActiveNodes(arg0 uint32) (uint32, error) {
	return _Contract.Contract.GlobalDailyActiveNodes(&_Contract.CallOpts, arg0)
}

// GlobalDailyActiveNodes is a free data retrieval call binding the contract method 0xea715627.
//
// Solidity: function globalDailyActiveNodes(uint32 ) view returns(uint32)
func (_Contract *ContractCallerSession) GlobalDailyActiveNodes(arg0 uint32) (uint32, error) {
	return _Contract.Contract.GlobalDailyActiveNodes(&_Contract.CallOpts, arg0)
}

// NodeAddrByID is a free data retrieval call binding the contract method 0xc018a806.
//
// Solidity: function nodeAddrByID(uint32 ) view returns(address)
func (_Contract *ContractCaller) NodeAddrByID(opts *bind.CallOpts, arg0 uint32) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "nodeAddrByID", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeAddrByID is a free data retrieval call binding the contract method 0xc018a806.
//
// Solidity: function nodeAddrByID(uint32 ) view returns(address)
func (_Contract *ContractSession) NodeAddrByID(arg0 uint32) (common.Address, error) {
	return _Contract.Contract.NodeAddrByID(&_Contract.CallOpts, arg0)
}

// NodeAddrByID is a free data retrieval call binding the contract method 0xc018a806.
//
// Solidity: function nodeAddrByID(uint32 ) view returns(address)
func (_Contract *ContractCallerSession) NodeAddrByID(arg0 uint32) (common.Address, error) {
	return _Contract.Contract.NodeAddrByID(&_Contract.CallOpts, arg0)
}

// NodeClaimedTeeRewards is a free data retrieval call binding the contract method 0xcc4ed7c5.
//
// Solidity: function nodeClaimedTeeRewards(address , bytes32 ) view returns(bool)
func (_Contract *ContractCaller) NodeClaimedTeeRewards(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "nodeClaimedTeeRewards", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NodeClaimedTeeRewards is a free data retrieval call binding the contract method 0xcc4ed7c5.
//
// Solidity: function nodeClaimedTeeRewards(address , bytes32 ) view returns(bool)
func (_Contract *ContractSession) NodeClaimedTeeRewards(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _Contract.Contract.NodeClaimedTeeRewards(&_Contract.CallOpts, arg0, arg1)
}

// NodeClaimedTeeRewards is a free data retrieval call binding the contract method 0xcc4ed7c5.
//
// Solidity: function nodeClaimedTeeRewards(address , bytes32 ) view returns(bool)
func (_Contract *ContractCallerSession) NodeClaimedTeeRewards(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _Contract.Contract.NodeClaimedTeeRewards(&_Contract.CallOpts, arg0, arg1)
}

// NodeDailyActive is a free data retrieval call binding the contract method 0xa3b4c737.
//
// Solidity: function nodeDailyActive(address , uint32 ) view returns(uint32)
func (_Contract *ContractCaller) NodeDailyActive(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (uint32, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "nodeDailyActive", arg0, arg1)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NodeDailyActive is a free data retrieval call binding the contract method 0xa3b4c737.
//
// Solidity: function nodeDailyActive(address , uint32 ) view returns(uint32)
func (_Contract *ContractSession) NodeDailyActive(arg0 common.Address, arg1 uint32) (uint32, error) {
	return _Contract.Contract.NodeDailyActive(&_Contract.CallOpts, arg0, arg1)
}

// NodeDailyActive is a free data retrieval call binding the contract method 0xa3b4c737.
//
// Solidity: function nodeDailyActive(address , uint32 ) view returns(uint32)
func (_Contract *ContractCallerSession) NodeDailyActive(arg0 common.Address, arg1 uint32) (uint32, error) {
	return _Contract.Contract.NodeDailyActive(&_Contract.CallOpts, arg0, arg1)
}

// NodeIndex is a free data retrieval call binding the contract method 0x796ba4d7.
//
// Solidity: function nodeIndex() view returns(uint32)
func (_Contract *ContractCaller) NodeIndex(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "nodeIndex")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NodeIndex is a free data retrieval call binding the contract method 0x796ba4d7.
//
// Solidity: function nodeIndex() view returns(uint32)
func (_Contract *ContractSession) NodeIndex() (uint32, error) {
	return _Contract.Contract.NodeIndex(&_Contract.CallOpts)
}

// NodeIndex is a free data retrieval call binding the contract method 0x796ba4d7.
//
// Solidity: function nodeIndex() view returns(uint32)
func (_Contract *ContractCallerSession) NodeIndex() (uint32, error) {
	return _Contract.Contract.NodeIndex(&_Contract.CallOpts)
}

// NodeInfos is a free data retrieval call binding the contract method 0x797e18e8.
//
// Solidity: function nodeInfos(address ) view returns(uint32 id, uint32 listIndex, bool active, uint32 lastConfirmDate, uint64 missedVerifyCount, int256 selfTotalRewards, uint256 selfClaimedRewards, uint256 delegationRewards, uint256 lastEnterTime, address claimer, uint32 commissionRate, uint256 commissionRateLastModifyAt)
func (_Contract *ContractCaller) NodeInfos(opts *bind.CallOpts, arg0 common.Address) (struct {
	Id                         uint32
	ListIndex                  uint32
	Active                     bool
	LastConfirmDate            uint32
	MissedVerifyCount          uint64
	SelfTotalRewards           *big.Int
	SelfClaimedRewards         *big.Int
	DelegationRewards          *big.Int
	LastEnterTime              *big.Int
	Claimer                    common.Address
	CommissionRate             uint32
	CommissionRateLastModifyAt *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "nodeInfos", arg0)

	outstruct := new(struct {
		Id                         uint32
		ListIndex                  uint32
		Active                     bool
		LastConfirmDate            uint32
		MissedVerifyCount          uint64
		SelfTotalRewards           *big.Int
		SelfClaimedRewards         *big.Int
		DelegationRewards          *big.Int
		LastEnterTime              *big.Int
		Claimer                    common.Address
		CommissionRate             uint32
		CommissionRateLastModifyAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.ListIndex = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Active = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.LastConfirmDate = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.MissedVerifyCount = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.SelfTotalRewards = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.SelfClaimedRewards = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.DelegationRewards = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.LastEnterTime = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.Claimer = *abi.ConvertType(out[9], new(common.Address)).(*common.Address)
	outstruct.CommissionRate = *abi.ConvertType(out[10], new(uint32)).(*uint32)
	outstruct.CommissionRateLastModifyAt = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// NodeInfos is a free data retrieval call binding the contract method 0x797e18e8.
//
// Solidity: function nodeInfos(address ) view returns(uint32 id, uint32 listIndex, bool active, uint32 lastConfirmDate, uint64 missedVerifyCount, int256 selfTotalRewards, uint256 selfClaimedRewards, uint256 delegationRewards, uint256 lastEnterTime, address claimer, uint32 commissionRate, uint256 commissionRateLastModifyAt)
func (_Contract *ContractSession) NodeInfos(arg0 common.Address) (struct {
	Id                         uint32
	ListIndex                  uint32
	Active                     bool
	LastConfirmDate            uint32
	MissedVerifyCount          uint64
	SelfTotalRewards           *big.Int
	SelfClaimedRewards         *big.Int
	DelegationRewards          *big.Int
	LastEnterTime              *big.Int
	Claimer                    common.Address
	CommissionRate             uint32
	CommissionRateLastModifyAt *big.Int
}, error) {
	return _Contract.Contract.NodeInfos(&_Contract.CallOpts, arg0)
}

// NodeInfos is a free data retrieval call binding the contract method 0x797e18e8.
//
// Solidity: function nodeInfos(address ) view returns(uint32 id, uint32 listIndex, bool active, uint32 lastConfirmDate, uint64 missedVerifyCount, int256 selfTotalRewards, uint256 selfClaimedRewards, uint256 delegationRewards, uint256 lastEnterTime, address claimer, uint32 commissionRate, uint256 commissionRateLastModifyAt)
func (_Contract *ContractCallerSession) NodeInfos(arg0 common.Address) (struct {
	Id                         uint32
	ListIndex                  uint32
	Active                     bool
	LastConfirmDate            uint32
	MissedVerifyCount          uint64
	SelfTotalRewards           *big.Int
	SelfClaimedRewards         *big.Int
	DelegationRewards          *big.Int
	LastEnterTime              *big.Int
	Claimer                    common.Address
	CommissionRate             uint32
	CommissionRateLastModifyAt *big.Int
}, error) {
	return _Contract.Contract.NodeInfos(&_Contract.CallOpts, arg0)
}

// NodeSlashed is a free data retrieval call binding the contract method 0x03c13728.
//
// Solidity: function nodeSlashed(address , bytes32 ) view returns(bool)
func (_Contract *ContractCaller) NodeSlashed(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "nodeSlashed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NodeSlashed is a free data retrieval call binding the contract method 0x03c13728.
//
// Solidity: function nodeSlashed(address , bytes32 ) view returns(bool)
func (_Contract *ContractSession) NodeSlashed(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _Contract.Contract.NodeSlashed(&_Contract.CallOpts, arg0, arg1)
}

// NodeSlashed is a free data retrieval call binding the contract method 0x03c13728.
//
// Solidity: function nodeSlashed(address , bytes32 ) view returns(bool)
func (_Contract *ContractCallerSession) NodeSlashed(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _Contract.Contract.NodeSlashed(&_Contract.CallOpts, arg0, arg1)
}

// Request2AttestationIDs is a free data retrieval call binding the contract method 0xc71d5b2f.
//
// Solidity: function request2AttestationIDs(uint256 , uint256 ) view returns(bytes32)
func (_Contract *ContractCaller) Request2AttestationIDs(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "request2AttestationIDs", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Request2AttestationIDs is a free data retrieval call binding the contract method 0xc71d5b2f.
//
// Solidity: function request2AttestationIDs(uint256 , uint256 ) view returns(bytes32)
func (_Contract *ContractSession) Request2AttestationIDs(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _Contract.Contract.Request2AttestationIDs(&_Contract.CallOpts, arg0, arg1)
}

// Request2AttestationIDs is a free data retrieval call binding the contract method 0xc71d5b2f.
//
// Solidity: function request2AttestationIDs(uint256 , uint256 ) view returns(bytes32)
func (_Contract *ContractCallerSession) Request2AttestationIDs(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _Contract.Contract.Request2AttestationIDs(&_Contract.CallOpts, arg0, arg1)
}

// Settings is a free data retrieval call binding the contract method 0xe06174e4.
//
// Solidity: function settings() view returns(address)
func (_Contract *ContractCaller) Settings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "settings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Settings is a free data retrieval call binding the contract method 0xe06174e4.
//
// Solidity: function settings() view returns(address)
func (_Contract *ContractSession) Settings() (common.Address, error) {
	return _Contract.Contract.Settings(&_Contract.CallOpts)
}

// Settings is a free data retrieval call binding the contract method 0xe06174e4.
//
// Solidity: function settings() view returns(address)
func (_Contract *ContractCallerSession) Settings() (common.Address, error) {
	return _Contract.Contract.Settings(&_Contract.CallOpts)
}

// TeeRoles is a free data retrieval call binding the contract method 0x34bee6ee.
//
// Solidity: function teeRoles(address ) view returns(bool)
func (_Contract *ContractCaller) TeeRoles(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "teeRoles", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TeeRoles is a free data retrieval call binding the contract method 0x34bee6ee.
//
// Solidity: function teeRoles(address ) view returns(bool)
func (_Contract *ContractSession) TeeRoles(arg0 common.Address) (bool, error) {
	return _Contract.Contract.TeeRoles(&_Contract.CallOpts, arg0)
}

// TeeRoles is a free data retrieval call binding the contract method 0x34bee6ee.
//
// Solidity: function teeRoles(address ) view returns(bool)
func (_Contract *ContractCallerSession) TeeRoles(arg0 common.Address) (bool, error) {
	return _Contract.Contract.TeeRoles(&_Contract.CallOpts, arg0)
}

// TeeStakeInfos is a free data retrieval call binding the contract method 0x68f41810.
//
// Solidity: function teeStakeInfos(address ) view returns(bool valid, uint256 staked, uint256 lastReportAt)
func (_Contract *ContractCaller) TeeStakeInfos(opts *bind.CallOpts, arg0 common.Address) (struct {
	Valid        bool
	Staked       *big.Int
	LastReportAt *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "teeStakeInfos", arg0)

	outstruct := new(struct {
		Valid        bool
		Staked       *big.Int
		LastReportAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Valid = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Staked = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastReportAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TeeStakeInfos is a free data retrieval call binding the contract method 0x68f41810.
//
// Solidity: function teeStakeInfos(address ) view returns(bool valid, uint256 staked, uint256 lastReportAt)
func (_Contract *ContractSession) TeeStakeInfos(arg0 common.Address) (struct {
	Valid        bool
	Staked       *big.Int
	LastReportAt *big.Int
}, error) {
	return _Contract.Contract.TeeStakeInfos(&_Contract.CallOpts, arg0)
}

// TeeStakeInfos is a free data retrieval call binding the contract method 0x68f41810.
//
// Solidity: function teeStakeInfos(address ) view returns(bool valid, uint256 staked, uint256 lastReportAt)
func (_Contract *ContractCallerSession) TeeStakeInfos(arg0 common.Address) (struct {
	Valid        bool
	Staked       *big.Int
	LastReportAt *big.Int
}, error) {
	return _Contract.Contract.TeeStakeInfos(&_Contract.CallOpts, arg0)
}

// TodayIndex is a free data retrieval call binding the contract method 0x37c68423.
//
// Solidity: function todayIndex() view returns(uint32)
func (_Contract *ContractCaller) TodayIndex(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "todayIndex")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TodayIndex is a free data retrieval call binding the contract method 0x37c68423.
//
// Solidity: function todayIndex() view returns(uint32)
func (_Contract *ContractSession) TodayIndex() (uint32, error) {
	return _Contract.Contract.TodayIndex(&_Contract.CallOpts)
}

// TodayIndex is a free data retrieval call binding the contract method 0x37c68423.
//
// Solidity: function todayIndex() view returns(uint32)
func (_Contract *ContractCallerSession) TodayIndex() (uint32, error) {
	return _Contract.Contract.TodayIndex(&_Contract.CallOpts)
}

// TodayOffset is a free data retrieval call binding the contract method 0xdc57a647.
//
// Solidity: function todayOffset() view returns(uint256)
func (_Contract *ContractCaller) TodayOffset(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "todayOffset")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TodayOffset is a free data retrieval call binding the contract method 0xdc57a647.
//
// Solidity: function todayOffset() view returns(uint256)
func (_Contract *ContractSession) TodayOffset() (*big.Int, error) {
	return _Contract.Contract.TodayOffset(&_Contract.CallOpts)
}

// TodayOffset is a free data retrieval call binding the contract method 0xdc57a647.
//
// Solidity: function todayOffset() view returns(uint256)
func (_Contract *ContractCallerSession) TodayOffset() (*big.Int, error) {
	return _Contract.Contract.TodayOffset(&_Contract.CallOpts)
}

// TokenRewardInfos is a free data retrieval call binding the contract method 0x3b85a346.
//
// Solidity: function tokenRewardInfos(uint256 ) view returns(uint256 initialRewards, uint256 totalRewards, uint256 claimedRewards)
func (_Contract *ContractCaller) TokenRewardInfos(opts *bind.CallOpts, arg0 *big.Int) (struct {
	InitialRewards *big.Int
	TotalRewards   *big.Int
	ClaimedRewards *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "tokenRewardInfos", arg0)

	outstruct := new(struct {
		InitialRewards *big.Int
		TotalRewards   *big.Int
		ClaimedRewards *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InitialRewards = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalRewards = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ClaimedRewards = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TokenRewardInfos is a free data retrieval call binding the contract method 0x3b85a346.
//
// Solidity: function tokenRewardInfos(uint256 ) view returns(uint256 initialRewards, uint256 totalRewards, uint256 claimedRewards)
func (_Contract *ContractSession) TokenRewardInfos(arg0 *big.Int) (struct {
	InitialRewards *big.Int
	TotalRewards   *big.Int
	ClaimedRewards *big.Int
}, error) {
	return _Contract.Contract.TokenRewardInfos(&_Contract.CallOpts, arg0)
}

// TokenRewardInfos is a free data retrieval call binding the contract method 0x3b85a346.
//
// Solidity: function tokenRewardInfos(uint256 ) view returns(uint256 initialRewards, uint256 totalRewards, uint256 claimedRewards)
func (_Contract *ContractCallerSession) TokenRewardInfos(arg0 *big.Int) (struct {
	InitialRewards *big.Int
	TotalRewards   *big.Int
	ClaimedRewards *big.Int
}, error) {
	return _Contract.Contract.TokenRewardInfos(&_Contract.CallOpts, arg0)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Contract *ContractCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Contract *ContractSession) Vault() (common.Address, error) {
	return _Contract.Contract.Vault(&_Contract.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Contract *ContractCallerSession) Vault() (common.Address, error) {
	return _Contract.Contract.Vault(&_Contract.CallOpts)
}

// VrfChosenIndex is a free data retrieval call binding the contract method 0xce3f3d95.
//
// Solidity: function vrfChosenIndex() view returns(uint32)
func (_Contract *ContractCaller) VrfChosenIndex(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "vrfChosenIndex")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// VrfChosenIndex is a free data retrieval call binding the contract method 0xce3f3d95.
//
// Solidity: function vrfChosenIndex() view returns(uint32)
func (_Contract *ContractSession) VrfChosenIndex() (uint32, error) {
	return _Contract.Contract.VrfChosenIndex(&_Contract.CallOpts)
}

// VrfChosenIndex is a free data retrieval call binding the contract method 0xce3f3d95.
//
// Solidity: function vrfChosenIndex() view returns(uint32)
func (_Contract *ContractCallerSession) VrfChosenIndex() (uint32, error) {
	return _Contract.Contract.VrfChosenIndex(&_Contract.CallOpts)
}

// VrfChosenMap is a free data retrieval call binding the contract method 0x09d65fd8.
//
// Solidity: function vrfChosenMap(uint32 , uint256 ) view returns(uint32)
func (_Contract *ContractCaller) VrfChosenMap(opts *bind.CallOpts, arg0 uint32, arg1 *big.Int) (uint32, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "vrfChosenMap", arg0, arg1)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// VrfChosenMap is a free data retrieval call binding the contract method 0x09d65fd8.
//
// Solidity: function vrfChosenMap(uint32 , uint256 ) view returns(uint32)
func (_Contract *ContractSession) VrfChosenMap(arg0 uint32, arg1 *big.Int) (uint32, error) {
	return _Contract.Contract.VrfChosenMap(&_Contract.CallOpts, arg0, arg1)
}

// VrfChosenMap is a free data retrieval call binding the contract method 0x09d65fd8.
//
// Solidity: function vrfChosenMap(uint32 , uint256 ) view returns(uint32)
func (_Contract *ContractCallerSession) VrfChosenMap(arg0 uint32, arg1 *big.Int) (uint32, error) {
	return _Contract.Contract.VrfChosenMap(&_Contract.CallOpts, arg0, arg1)
}

// ClaimMaliciousTeeRewards is a paid mutator transaction binding the contract method 0x4bbc4c97.
//
// Solidity: function claimMaliciousTeeRewards(bytes32 attestationID) returns()
func (_Contract *ContractTransactor) ClaimMaliciousTeeRewards(opts *bind.TransactOpts, attestationID [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimMaliciousTeeRewards", attestationID)
}

// ClaimMaliciousTeeRewards is a paid mutator transaction binding the contract method 0x4bbc4c97.
//
// Solidity: function claimMaliciousTeeRewards(bytes32 attestationID) returns()
func (_Contract *ContractSession) ClaimMaliciousTeeRewards(attestationID [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.ClaimMaliciousTeeRewards(&_Contract.TransactOpts, attestationID)
}

// ClaimMaliciousTeeRewards is a paid mutator transaction binding the contract method 0x4bbc4c97.
//
// Solidity: function claimMaliciousTeeRewards(bytes32 attestationID) returns()
func (_Contract *ContractTransactorSession) ClaimMaliciousTeeRewards(attestationID [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.ClaimMaliciousTeeRewards(&_Contract.TransactOpts, attestationID)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x0962ef79.
//
// Solidity: function claimRewards(uint256 tokenID) returns()
func (_Contract *ContractTransactor) ClaimRewards(opts *bind.TransactOpts, tokenID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimRewards", tokenID)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x0962ef79.
//
// Solidity: function claimRewards(uint256 tokenID) returns()
func (_Contract *ContractSession) ClaimRewards(tokenID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimRewards(&_Contract.TransactOpts, tokenID)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x0962ef79.
//
// Solidity: function claimRewards(uint256 tokenID) returns()
func (_Contract *ContractTransactorSession) ClaimRewards(tokenID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimRewards(&_Contract.TransactOpts, tokenID)
}

// Delegate is a paid mutator transaction binding the contract method 0x08bbb824.
//
// Solidity: function delegate(uint256 tokenID, address to) returns()
func (_Contract *ContractTransactor) Delegate(opts *bind.TransactOpts, tokenID *big.Int, to common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "delegate", tokenID, to)
}

// Delegate is a paid mutator transaction binding the contract method 0x08bbb824.
//
// Solidity: function delegate(uint256 tokenID, address to) returns()
func (_Contract *ContractSession) Delegate(tokenID *big.Int, to common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Delegate(&_Contract.TransactOpts, tokenID, to)
}

// Delegate is a paid mutator transaction binding the contract method 0x08bbb824.
//
// Solidity: function delegate(uint256 tokenID, address to) returns()
func (_Contract *ContractTransactorSession) Delegate(tokenID *big.Int, to common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Delegate(&_Contract.TransactOpts, tokenID, to)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x38ba4614.
//
// Solidity: function fulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_Contract *ContractTransactor) FulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "fulfillRandomWords", requestId, randomWords)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x38ba4614.
//
// Solidity: function fulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_Contract *ContractSession) FulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.FulfillRandomWords(&_Contract.TransactOpts, requestId, randomWords)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x38ba4614.
//
// Solidity: function fulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_Contract *ContractTransactorSession) FulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.FulfillRandomWords(&_Contract.TransactOpts, requestId, randomWords)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address carvToken_, address carvNft_, address vault_, uint256 chainID) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, carvToken_ common.Address, carvNft_ common.Address, vault_ common.Address, chainID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", carvToken_, carvNft_, vault_, chainID)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address carvToken_, address carvNft_, address vault_, uint256 chainID) returns()
func (_Contract *ContractSession) Initialize(carvToken_ common.Address, carvNft_ common.Address, vault_ common.Address, chainID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, carvToken_, carvNft_, vault_, chainID)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address carvToken_, address carvNft_, address vault_, uint256 chainID) returns()
func (_Contract *ContractTransactorSession) Initialize(carvToken_ common.Address, carvNft_ common.Address, vault_ common.Address, chainID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, carvToken_, carvNft_, vault_, chainID)
}

// ModifyAdmin is a paid mutator transaction binding the contract method 0x055b1edd.
//
// Solidity: function modifyAdmin(address newAdmin) returns()
func (_Contract *ContractTransactor) ModifyAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "modifyAdmin", newAdmin)
}

// ModifyAdmin is a paid mutator transaction binding the contract method 0x055b1edd.
//
// Solidity: function modifyAdmin(address newAdmin) returns()
func (_Contract *ContractSession) ModifyAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ModifyAdmin(&_Contract.TransactOpts, newAdmin)
}

// ModifyAdmin is a paid mutator transaction binding the contract method 0x055b1edd.
//
// Solidity: function modifyAdmin(address newAdmin) returns()
func (_Contract *ContractTransactorSession) ModifyAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ModifyAdmin(&_Contract.TransactOpts, newAdmin)
}

// ModifyTeeRole is a paid mutator transaction binding the contract method 0x5890b958.
//
// Solidity: function modifyTeeRole(address tee, bool grant) returns()
func (_Contract *ContractTransactor) ModifyTeeRole(opts *bind.TransactOpts, tee common.Address, grant bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "modifyTeeRole", tee, grant)
}

// ModifyTeeRole is a paid mutator transaction binding the contract method 0x5890b958.
//
// Solidity: function modifyTeeRole(address tee, bool grant) returns()
func (_Contract *ContractSession) ModifyTeeRole(tee common.Address, grant bool) (*types.Transaction, error) {
	return _Contract.Contract.ModifyTeeRole(&_Contract.TransactOpts, tee, grant)
}

// ModifyTeeRole is a paid mutator transaction binding the contract method 0x5890b958.
//
// Solidity: function modifyTeeRole(address tee, bool grant) returns()
func (_Contract *ContractTransactorSession) ModifyTeeRole(tee common.Address, grant bool) (*types.Transaction, error) {
	return _Contract.Contract.ModifyTeeRole(&_Contract.TransactOpts, tee, grant)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Contract *ContractTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Contract *ContractSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.Multicall(&_Contract.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Contract *ContractTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.Multicall(&_Contract.TransactOpts, data)
}

// NodeClaim is a paid mutator transaction binding the contract method 0xf39a19bf.
//
// Solidity: function nodeClaim(address node) returns()
func (_Contract *ContractTransactor) NodeClaim(opts *bind.TransactOpts, node common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "nodeClaim", node)
}

// NodeClaim is a paid mutator transaction binding the contract method 0xf39a19bf.
//
// Solidity: function nodeClaim(address node) returns()
func (_Contract *ContractSession) NodeClaim(node common.Address) (*types.Transaction, error) {
	return _Contract.Contract.NodeClaim(&_Contract.TransactOpts, node)
}

// NodeClaim is a paid mutator transaction binding the contract method 0xf39a19bf.
//
// Solidity: function nodeClaim(address node) returns()
func (_Contract *ContractTransactorSession) NodeClaim(node common.Address) (*types.Transaction, error) {
	return _Contract.Contract.NodeClaim(&_Contract.TransactOpts, node)
}

// NodeEnter is a paid mutator transaction binding the contract method 0x73df96ed.
//
// Solidity: function nodeEnter(address replacedNode) returns()
func (_Contract *ContractTransactor) NodeEnter(opts *bind.TransactOpts, replacedNode common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "nodeEnter", replacedNode)
}

// NodeEnter is a paid mutator transaction binding the contract method 0x73df96ed.
//
// Solidity: function nodeEnter(address replacedNode) returns()
func (_Contract *ContractSession) NodeEnter(replacedNode common.Address) (*types.Transaction, error) {
	return _Contract.Contract.NodeEnter(&_Contract.TransactOpts, replacedNode)
}

// NodeEnter is a paid mutator transaction binding the contract method 0x73df96ed.
//
// Solidity: function nodeEnter(address replacedNode) returns()
func (_Contract *ContractTransactorSession) NodeEnter(replacedNode common.Address) (*types.Transaction, error) {
	return _Contract.Contract.NodeEnter(&_Contract.TransactOpts, replacedNode)
}

// NodeEnterWithSignature is a paid mutator transaction binding the contract method 0x47b8706c.
//
// Solidity: function nodeEnterWithSignature(address replacedNode, uint256 expiredAt, address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contract *ContractTransactor) NodeEnterWithSignature(opts *bind.TransactOpts, replacedNode common.Address, expiredAt *big.Int, signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "nodeEnterWithSignature", replacedNode, expiredAt, signer, v, r, s)
}

// NodeEnterWithSignature is a paid mutator transaction binding the contract method 0x47b8706c.
//
// Solidity: function nodeEnterWithSignature(address replacedNode, uint256 expiredAt, address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contract *ContractSession) NodeEnterWithSignature(replacedNode common.Address, expiredAt *big.Int, signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.NodeEnterWithSignature(&_Contract.TransactOpts, replacedNode, expiredAt, signer, v, r, s)
}

// NodeEnterWithSignature is a paid mutator transaction binding the contract method 0x47b8706c.
//
// Solidity: function nodeEnterWithSignature(address replacedNode, uint256 expiredAt, address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contract *ContractTransactorSession) NodeEnterWithSignature(replacedNode common.Address, expiredAt *big.Int, signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.NodeEnterWithSignature(&_Contract.TransactOpts, replacedNode, expiredAt, signer, v, r, s)
}

// NodeExit is a paid mutator transaction binding the contract method 0xd46849db.
//
// Solidity: function nodeExit() returns()
func (_Contract *ContractTransactor) NodeExit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "nodeExit")
}

// NodeExit is a paid mutator transaction binding the contract method 0xd46849db.
//
// Solidity: function nodeExit() returns()
func (_Contract *ContractSession) NodeExit() (*types.Transaction, error) {
	return _Contract.Contract.NodeExit(&_Contract.TransactOpts)
}

// NodeExit is a paid mutator transaction binding the contract method 0xd46849db.
//
// Solidity: function nodeExit() returns()
func (_Contract *ContractTransactorSession) NodeExit() (*types.Transaction, error) {
	return _Contract.Contract.NodeExit(&_Contract.TransactOpts)
}

// NodeExitWithSignature is a paid mutator transaction binding the contract method 0x7f32e167.
//
// Solidity: function nodeExitWithSignature(uint256 expiredAt, address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contract *ContractTransactor) NodeExitWithSignature(opts *bind.TransactOpts, expiredAt *big.Int, signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "nodeExitWithSignature", expiredAt, signer, v, r, s)
}

// NodeExitWithSignature is a paid mutator transaction binding the contract method 0x7f32e167.
//
// Solidity: function nodeExitWithSignature(uint256 expiredAt, address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contract *ContractSession) NodeExitWithSignature(expiredAt *big.Int, signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.NodeExitWithSignature(&_Contract.TransactOpts, expiredAt, signer, v, r, s)
}

// NodeExitWithSignature is a paid mutator transaction binding the contract method 0x7f32e167.
//
// Solidity: function nodeExitWithSignature(uint256 expiredAt, address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contract *ContractTransactorSession) NodeExitWithSignature(expiredAt *big.Int, signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.NodeExitWithSignature(&_Contract.TransactOpts, expiredAt, signer, v, r, s)
}

// NodeModifyCommissionRate is a paid mutator transaction binding the contract method 0x195f3a9c.
//
// Solidity: function nodeModifyCommissionRate(uint32 commissionRate) returns()
func (_Contract *ContractTransactor) NodeModifyCommissionRate(opts *bind.TransactOpts, commissionRate uint32) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "nodeModifyCommissionRate", commissionRate)
}

// NodeModifyCommissionRate is a paid mutator transaction binding the contract method 0x195f3a9c.
//
// Solidity: function nodeModifyCommissionRate(uint32 commissionRate) returns()
func (_Contract *ContractSession) NodeModifyCommissionRate(commissionRate uint32) (*types.Transaction, error) {
	return _Contract.Contract.NodeModifyCommissionRate(&_Contract.TransactOpts, commissionRate)
}

// NodeModifyCommissionRate is a paid mutator transaction binding the contract method 0x195f3a9c.
//
// Solidity: function nodeModifyCommissionRate(uint32 commissionRate) returns()
func (_Contract *ContractTransactorSession) NodeModifyCommissionRate(commissionRate uint32) (*types.Transaction, error) {
	return _Contract.Contract.NodeModifyCommissionRate(&_Contract.TransactOpts, commissionRate)
}

// NodeModifyCommissionRateWithSignature is a paid mutator transaction binding the contract method 0x8a4fd0c8.
//
// Solidity: function nodeModifyCommissionRateWithSignature(uint32 commissionRate, uint256 expiredAt, address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contract *ContractTransactor) NodeModifyCommissionRateWithSignature(opts *bind.TransactOpts, commissionRate uint32, expiredAt *big.Int, signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "nodeModifyCommissionRateWithSignature", commissionRate, expiredAt, signer, v, r, s)
}

// NodeModifyCommissionRateWithSignature is a paid mutator transaction binding the contract method 0x8a4fd0c8.
//
// Solidity: function nodeModifyCommissionRateWithSignature(uint32 commissionRate, uint256 expiredAt, address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contract *ContractSession) NodeModifyCommissionRateWithSignature(commissionRate uint32, expiredAt *big.Int, signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.NodeModifyCommissionRateWithSignature(&_Contract.TransactOpts, commissionRate, expiredAt, signer, v, r, s)
}

// NodeModifyCommissionRateWithSignature is a paid mutator transaction binding the contract method 0x8a4fd0c8.
//
// Solidity: function nodeModifyCommissionRateWithSignature(uint32 commissionRate, uint256 expiredAt, address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contract *ContractTransactorSession) NodeModifyCommissionRateWithSignature(commissionRate uint32, expiredAt *big.Int, signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.NodeModifyCommissionRateWithSignature(&_Contract.TransactOpts, commissionRate, expiredAt, signer, v, r, s)
}

// NodeReportDailyActive is a paid mutator transaction binding the contract method 0x34d37fa0.
//
// Solidity: function nodeReportDailyActive(address node) returns()
func (_Contract *ContractTransactor) NodeReportDailyActive(opts *bind.TransactOpts, node common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "nodeReportDailyActive", node)
}

// NodeReportDailyActive is a paid mutator transaction binding the contract method 0x34d37fa0.
//
// Solidity: function nodeReportDailyActive(address node) returns()
func (_Contract *ContractSession) NodeReportDailyActive(node common.Address) (*types.Transaction, error) {
	return _Contract.Contract.NodeReportDailyActive(&_Contract.TransactOpts, node)
}

// NodeReportDailyActive is a paid mutator transaction binding the contract method 0x34d37fa0.
//
// Solidity: function nodeReportDailyActive(address node) returns()
func (_Contract *ContractTransactorSession) NodeReportDailyActive(node common.Address) (*types.Transaction, error) {
	return _Contract.Contract.NodeReportDailyActive(&_Contract.TransactOpts, node)
}

// NodeReportVerification is a paid mutator transaction binding the contract method 0xe93ee212.
//
// Solidity: function nodeReportVerification(bytes32 attestationID, uint32 index, uint8 result) returns()
func (_Contract *ContractTransactor) NodeReportVerification(opts *bind.TransactOpts, attestationID [32]byte, index uint32, result uint8) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "nodeReportVerification", attestationID, index, result)
}

// NodeReportVerification is a paid mutator transaction binding the contract method 0xe93ee212.
//
// Solidity: function nodeReportVerification(bytes32 attestationID, uint32 index, uint8 result) returns()
func (_Contract *ContractSession) NodeReportVerification(attestationID [32]byte, index uint32, result uint8) (*types.Transaction, error) {
	return _Contract.Contract.NodeReportVerification(&_Contract.TransactOpts, attestationID, index, result)
}

// NodeReportVerification is a paid mutator transaction binding the contract method 0xe93ee212.
//
// Solidity: function nodeReportVerification(bytes32 attestationID, uint32 index, uint8 result) returns()
func (_Contract *ContractTransactorSession) NodeReportVerification(attestationID [32]byte, index uint32, result uint8) (*types.Transaction, error) {
	return _Contract.Contract.NodeReportVerification(&_Contract.TransactOpts, attestationID, index, result)
}

// NodeReportVerificationBatch is a paid mutator transaction binding the contract method 0xb50bd0d4.
//
// Solidity: function nodeReportVerificationBatch(bytes32 attestationID, (uint8,uint32,address,uint8,bytes32,bytes32)[] infos) returns()
func (_Contract *ContractTransactor) NodeReportVerificationBatch(opts *bind.TransactOpts, attestationID [32]byte, infos []IProtocolServiceVerificationInfo) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "nodeReportVerificationBatch", attestationID, infos)
}

// NodeReportVerificationBatch is a paid mutator transaction binding the contract method 0xb50bd0d4.
//
// Solidity: function nodeReportVerificationBatch(bytes32 attestationID, (uint8,uint32,address,uint8,bytes32,bytes32)[] infos) returns()
func (_Contract *ContractSession) NodeReportVerificationBatch(attestationID [32]byte, infos []IProtocolServiceVerificationInfo) (*types.Transaction, error) {
	return _Contract.Contract.NodeReportVerificationBatch(&_Contract.TransactOpts, attestationID, infos)
}

// NodeReportVerificationBatch is a paid mutator transaction binding the contract method 0xb50bd0d4.
//
// Solidity: function nodeReportVerificationBatch(bytes32 attestationID, (uint8,uint32,address,uint8,bytes32,bytes32)[] infos) returns()
func (_Contract *ContractTransactorSession) NodeReportVerificationBatch(attestationID [32]byte, infos []IProtocolServiceVerificationInfo) (*types.Transaction, error) {
	return _Contract.Contract.NodeReportVerificationBatch(&_Contract.TransactOpts, attestationID, infos)
}

// NodeSetRewardClaimerWithSignature is a paid mutator transaction binding the contract method 0xc4a16c52.
//
// Solidity: function nodeSetRewardClaimerWithSignature(address claimer, uint256 expiredAt, address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contract *ContractTransactor) NodeSetRewardClaimerWithSignature(opts *bind.TransactOpts, claimer common.Address, expiredAt *big.Int, signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "nodeSetRewardClaimerWithSignature", claimer, expiredAt, signer, v, r, s)
}

// NodeSetRewardClaimerWithSignature is a paid mutator transaction binding the contract method 0xc4a16c52.
//
// Solidity: function nodeSetRewardClaimerWithSignature(address claimer, uint256 expiredAt, address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contract *ContractSession) NodeSetRewardClaimerWithSignature(claimer common.Address, expiredAt *big.Int, signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.NodeSetRewardClaimerWithSignature(&_Contract.TransactOpts, claimer, expiredAt, signer, v, r, s)
}

// NodeSetRewardClaimerWithSignature is a paid mutator transaction binding the contract method 0xc4a16c52.
//
// Solidity: function nodeSetRewardClaimerWithSignature(address claimer, uint256 expiredAt, address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contract *ContractTransactorSession) NodeSetRewardClaimerWithSignature(claimer common.Address, expiredAt *big.Int, signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.NodeSetRewardClaimerWithSignature(&_Contract.TransactOpts, claimer, expiredAt, signer, v, r, s)
}

// NodeSlash is a paid mutator transaction binding the contract method 0xdf5e1fdd.
//
// Solidity: function nodeSlash(address node, bytes32 attestationID, uint32 index) returns()
func (_Contract *ContractTransactor) NodeSlash(opts *bind.TransactOpts, node common.Address, attestationID [32]byte, index uint32) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "nodeSlash", node, attestationID, index)
}

// NodeSlash is a paid mutator transaction binding the contract method 0xdf5e1fdd.
//
// Solidity: function nodeSlash(address node, bytes32 attestationID, uint32 index) returns()
func (_Contract *ContractSession) NodeSlash(node common.Address, attestationID [32]byte, index uint32) (*types.Transaction, error) {
	return _Contract.Contract.NodeSlash(&_Contract.TransactOpts, node, attestationID, index)
}

// NodeSlash is a paid mutator transaction binding the contract method 0xdf5e1fdd.
//
// Solidity: function nodeSlash(address node, bytes32 attestationID, uint32 index) returns()
func (_Contract *ContractTransactorSession) NodeSlash(node common.Address, attestationID [32]byte, index uint32) (*types.Transaction, error) {
	return _Contract.Contract.NodeSlash(&_Contract.TransactOpts, node, attestationID, index)
}

// Redelegate is a paid mutator transaction binding the contract method 0x9ad4f752.
//
// Solidity: function redelegate(uint256 tokenID, address to) returns()
func (_Contract *ContractTransactor) Redelegate(opts *bind.TransactOpts, tokenID *big.Int, to common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "redelegate", tokenID, to)
}

// Redelegate is a paid mutator transaction binding the contract method 0x9ad4f752.
//
// Solidity: function redelegate(uint256 tokenID, address to) returns()
func (_Contract *ContractSession) Redelegate(tokenID *big.Int, to common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Redelegate(&_Contract.TransactOpts, tokenID, to)
}

// Redelegate is a paid mutator transaction binding the contract method 0x9ad4f752.
//
// Solidity: function redelegate(uint256 tokenID, address to) returns()
func (_Contract *ContractTransactorSession) Redelegate(tokenID *big.Int, to common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Redelegate(&_Contract.TransactOpts, tokenID, to)
}

// TeeReportAttestations is a paid mutator transaction binding the contract method 0x9d9b4270.
//
// Solidity: function teeReportAttestations(string[] attestationInfos) returns()
func (_Contract *ContractTransactor) TeeReportAttestations(opts *bind.TransactOpts, attestationInfos []string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "teeReportAttestations", attestationInfos)
}

// TeeReportAttestations is a paid mutator transaction binding the contract method 0x9d9b4270.
//
// Solidity: function teeReportAttestations(string[] attestationInfos) returns()
func (_Contract *ContractSession) TeeReportAttestations(attestationInfos []string) (*types.Transaction, error) {
	return _Contract.Contract.TeeReportAttestations(&_Contract.TransactOpts, attestationInfos)
}

// TeeReportAttestations is a paid mutator transaction binding the contract method 0x9d9b4270.
//
// Solidity: function teeReportAttestations(string[] attestationInfos) returns()
func (_Contract *ContractTransactorSession) TeeReportAttestations(attestationInfos []string) (*types.Transaction, error) {
	return _Contract.Contract.TeeReportAttestations(&_Contract.TransactOpts, attestationInfos)
}

// TeeSlash is a paid mutator transaction binding the contract method 0xe39a0beb.
//
// Solidity: function teeSlash(bytes32 attestationID) returns()
func (_Contract *ContractTransactor) TeeSlash(opts *bind.TransactOpts, attestationID [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "teeSlash", attestationID)
}

// TeeSlash is a paid mutator transaction binding the contract method 0xe39a0beb.
//
// Solidity: function teeSlash(bytes32 attestationID) returns()
func (_Contract *ContractSession) TeeSlash(attestationID [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.TeeSlash(&_Contract.TransactOpts, attestationID)
}

// TeeSlash is a paid mutator transaction binding the contract method 0xe39a0beb.
//
// Solidity: function teeSlash(bytes32 attestationID) returns()
func (_Contract *ContractTransactorSession) TeeSlash(attestationID [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.TeeSlash(&_Contract.TransactOpts, attestationID)
}

// TeeStake is a paid mutator transaction binding the contract method 0x9c509861.
//
// Solidity: function teeStake(uint256 amount) returns()
func (_Contract *ContractTransactor) TeeStake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "teeStake", amount)
}

// TeeStake is a paid mutator transaction binding the contract method 0x9c509861.
//
// Solidity: function teeStake(uint256 amount) returns()
func (_Contract *ContractSession) TeeStake(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TeeStake(&_Contract.TransactOpts, amount)
}

// TeeStake is a paid mutator transaction binding the contract method 0x9c509861.
//
// Solidity: function teeStake(uint256 amount) returns()
func (_Contract *ContractTransactorSession) TeeStake(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TeeStake(&_Contract.TransactOpts, amount)
}

// TeeUnstake is a paid mutator transaction binding the contract method 0x2d1af52e.
//
// Solidity: function teeUnstake() returns()
func (_Contract *ContractTransactor) TeeUnstake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "teeUnstake")
}

// TeeUnstake is a paid mutator transaction binding the contract method 0x2d1af52e.
//
// Solidity: function teeUnstake() returns()
func (_Contract *ContractSession) TeeUnstake() (*types.Transaction, error) {
	return _Contract.Contract.TeeUnstake(&_Contract.TransactOpts)
}

// TeeUnstake is a paid mutator transaction binding the contract method 0x2d1af52e.
//
// Solidity: function teeUnstake() returns()
func (_Contract *ContractTransactorSession) TeeUnstake() (*types.Transaction, error) {
	return _Contract.Contract.TeeUnstake(&_Contract.TransactOpts)
}

// Undelegate is a paid mutator transaction binding the contract method 0x6c68c0e1.
//
// Solidity: function undelegate(uint256 tokenID) returns()
func (_Contract *ContractTransactor) Undelegate(opts *bind.TransactOpts, tokenID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "undelegate", tokenID)
}

// Undelegate is a paid mutator transaction binding the contract method 0x6c68c0e1.
//
// Solidity: function undelegate(uint256 tokenID) returns()
func (_Contract *ContractSession) Undelegate(tokenID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Undelegate(&_Contract.TransactOpts, tokenID)
}

// Undelegate is a paid mutator transaction binding the contract method 0x6c68c0e1.
//
// Solidity: function undelegate(uint256 tokenID) returns()
func (_Contract *ContractTransactorSession) Undelegate(tokenID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Undelegate(&_Contract.TransactOpts, tokenID)
}

// UpdateSettingsAddress is a paid mutator transaction binding the contract method 0x23031abe.
//
// Solidity: function updateSettingsAddress(address settings_) returns()
func (_Contract *ContractTransactor) UpdateSettingsAddress(opts *bind.TransactOpts, settings_ common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateSettingsAddress", settings_)
}

// UpdateSettingsAddress is a paid mutator transaction binding the contract method 0x23031abe.
//
// Solidity: function updateSettingsAddress(address settings_) returns()
func (_Contract *ContractSession) UpdateSettingsAddress(settings_ common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateSettingsAddress(&_Contract.TransactOpts, settings_)
}

// UpdateSettingsAddress is a paid mutator transaction binding the contract method 0x23031abe.
//
// Solidity: function updateSettingsAddress(address settings_) returns()
func (_Contract *ContractTransactorSession) UpdateSettingsAddress(settings_ common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateSettingsAddress(&_Contract.TransactOpts, settings_)
}

// UpdateVrfAddress is a paid mutator transaction binding the contract method 0x4c9a41f9.
//
// Solidity: function updateVrfAddress(address carvVrf_) returns()
func (_Contract *ContractTransactor) UpdateVrfAddress(opts *bind.TransactOpts, carvVrf_ common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateVrfAddress", carvVrf_)
}

// UpdateVrfAddress is a paid mutator transaction binding the contract method 0x4c9a41f9.
//
// Solidity: function updateVrfAddress(address carvVrf_) returns()
func (_Contract *ContractSession) UpdateVrfAddress(carvVrf_ common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateVrfAddress(&_Contract.TransactOpts, carvVrf_)
}

// UpdateVrfAddress is a paid mutator transaction binding the contract method 0x4c9a41f9.
//
// Solidity: function updateVrfAddress(address carvVrf_) returns()
func (_Contract *ContractTransactorSession) UpdateVrfAddress(carvVrf_ common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateVrfAddress(&_Contract.TransactOpts, carvVrf_)
}

// ContractClaimMaliciousTeeRewardsIterator is returned from FilterClaimMaliciousTeeRewards and is used to iterate over the raw logs and unpacked data for ClaimMaliciousTeeRewards events raised by the Contract contract.
type ContractClaimMaliciousTeeRewardsIterator struct {
	Event *ContractClaimMaliciousTeeRewards // Event containing the contract specifics and raw log

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
func (it *ContractClaimMaliciousTeeRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractClaimMaliciousTeeRewards)
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
		it.Event = new(ContractClaimMaliciousTeeRewards)
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
func (it *ContractClaimMaliciousTeeRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractClaimMaliciousTeeRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractClaimMaliciousTeeRewards represents a ClaimMaliciousTeeRewards event raised by the Contract contract.
type ContractClaimMaliciousTeeRewards struct {
	Verifer common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimMaliciousTeeRewards is a free log retrieval operation binding the contract event 0xc1c7a9f718ce1c74f4b880bf361273760fb29077721becb28e8260c44cd6a9ea.
//
// Solidity: event ClaimMaliciousTeeRewards(address verifer, uint256 amount)
func (_Contract *ContractFilterer) FilterClaimMaliciousTeeRewards(opts *bind.FilterOpts) (*ContractClaimMaliciousTeeRewardsIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ClaimMaliciousTeeRewards")
	if err != nil {
		return nil, err
	}
	return &ContractClaimMaliciousTeeRewardsIterator{contract: _Contract.contract, event: "ClaimMaliciousTeeRewards", logs: logs, sub: sub}, nil
}

// WatchClaimMaliciousTeeRewards is a free log subscription operation binding the contract event 0xc1c7a9f718ce1c74f4b880bf361273760fb29077721becb28e8260c44cd6a9ea.
//
// Solidity: event ClaimMaliciousTeeRewards(address verifer, uint256 amount)
func (_Contract *ContractFilterer) WatchClaimMaliciousTeeRewards(opts *bind.WatchOpts, sink chan<- *ContractClaimMaliciousTeeRewards) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ClaimMaliciousTeeRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractClaimMaliciousTeeRewards)
				if err := _Contract.contract.UnpackLog(event, "ClaimMaliciousTeeRewards", log); err != nil {
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

// ParseClaimMaliciousTeeRewards is a log parse operation binding the contract event 0xc1c7a9f718ce1c74f4b880bf361273760fb29077721becb28e8260c44cd6a9ea.
//
// Solidity: event ClaimMaliciousTeeRewards(address verifer, uint256 amount)
func (_Contract *ContractFilterer) ParseClaimMaliciousTeeRewards(log types.Log) (*ContractClaimMaliciousTeeRewards, error) {
	event := new(ContractClaimMaliciousTeeRewards)
	if err := _Contract.contract.UnpackLog(event, "ClaimMaliciousTeeRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractClaimRewardsIterator is returned from FilterClaimRewards and is used to iterate over the raw logs and unpacked data for ClaimRewards events raised by the Contract contract.
type ContractClaimRewardsIterator struct {
	Event *ContractClaimRewards // Event containing the contract specifics and raw log

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
func (it *ContractClaimRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractClaimRewards)
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
		it.Event = new(ContractClaimRewards)
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
func (it *ContractClaimRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractClaimRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractClaimRewards represents a ClaimRewards event raised by the Contract contract.
type ContractClaimRewards struct {
	TokenID *big.Int
	Rewards *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimRewards is a free log retrieval operation binding the contract event 0x8b0944629ca33aba8dd5f33f7f8220efe77a2d5548a1651362c856c5ea586a65.
//
// Solidity: event ClaimRewards(uint256 tokenID, uint256 rewards)
func (_Contract *ContractFilterer) FilterClaimRewards(opts *bind.FilterOpts) (*ContractClaimRewardsIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ClaimRewards")
	if err != nil {
		return nil, err
	}
	return &ContractClaimRewardsIterator{contract: _Contract.contract, event: "ClaimRewards", logs: logs, sub: sub}, nil
}

// WatchClaimRewards is a free log subscription operation binding the contract event 0x8b0944629ca33aba8dd5f33f7f8220efe77a2d5548a1651362c856c5ea586a65.
//
// Solidity: event ClaimRewards(uint256 tokenID, uint256 rewards)
func (_Contract *ContractFilterer) WatchClaimRewards(opts *bind.WatchOpts, sink chan<- *ContractClaimRewards) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ClaimRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractClaimRewards)
				if err := _Contract.contract.UnpackLog(event, "ClaimRewards", log); err != nil {
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

// ParseClaimRewards is a log parse operation binding the contract event 0x8b0944629ca33aba8dd5f33f7f8220efe77a2d5548a1651362c856c5ea586a65.
//
// Solidity: event ClaimRewards(uint256 tokenID, uint256 rewards)
func (_Contract *ContractFilterer) ParseClaimRewards(log types.Log) (*ContractClaimRewards, error) {
	event := new(ContractClaimRewards)
	if err := _Contract.contract.UnpackLog(event, "ClaimRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractConfirmVrfNodesIterator is returned from FilterConfirmVrfNodes and is used to iterate over the raw logs and unpacked data for ConfirmVrfNodes events raised by the Contract contract.
type ContractConfirmVrfNodesIterator struct {
	Event *ContractConfirmVrfNodes // Event containing the contract specifics and raw log

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
func (it *ContractConfirmVrfNodesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractConfirmVrfNodes)
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
		it.Event = new(ContractConfirmVrfNodes)
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
func (it *ContractConfirmVrfNodesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractConfirmVrfNodesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractConfirmVrfNodes represents a ConfirmVrfNodes event raised by the Contract contract.
type ContractConfirmVrfNodes struct {
	RequestId *big.Int
	VrfChosen []uint32
	Deadline  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConfirmVrfNodes is a free log retrieval operation binding the contract event 0x455929120054502ca2ea8194b26e7bb3acb631d30177f6881ffa70581abd4a13.
//
// Solidity: event ConfirmVrfNodes(uint256 requestId, uint32[] vrfChosen, uint256 deadline)
func (_Contract *ContractFilterer) FilterConfirmVrfNodes(opts *bind.FilterOpts) (*ContractConfirmVrfNodesIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ConfirmVrfNodes")
	if err != nil {
		return nil, err
	}
	return &ContractConfirmVrfNodesIterator{contract: _Contract.contract, event: "ConfirmVrfNodes", logs: logs, sub: sub}, nil
}

// WatchConfirmVrfNodes is a free log subscription operation binding the contract event 0x455929120054502ca2ea8194b26e7bb3acb631d30177f6881ffa70581abd4a13.
//
// Solidity: event ConfirmVrfNodes(uint256 requestId, uint32[] vrfChosen, uint256 deadline)
func (_Contract *ContractFilterer) WatchConfirmVrfNodes(opts *bind.WatchOpts, sink chan<- *ContractConfirmVrfNodes) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ConfirmVrfNodes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractConfirmVrfNodes)
				if err := _Contract.contract.UnpackLog(event, "ConfirmVrfNodes", log); err != nil {
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

// ParseConfirmVrfNodes is a log parse operation binding the contract event 0x455929120054502ca2ea8194b26e7bb3acb631d30177f6881ffa70581abd4a13.
//
// Solidity: event ConfirmVrfNodes(uint256 requestId, uint32[] vrfChosen, uint256 deadline)
func (_Contract *ContractFilterer) ParseConfirmVrfNodes(log types.Log) (*ContractConfirmVrfNodes, error) {
	event := new(ContractConfirmVrfNodes)
	if err := _Contract.contract.UnpackLog(event, "ConfirmVrfNodes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegateIterator is returned from FilterDelegate and is used to iterate over the raw logs and unpacked data for Delegate events raised by the Contract contract.
type ContractDelegateIterator struct {
	Event *ContractDelegate // Event containing the contract specifics and raw log

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
func (it *ContractDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegate)
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
		it.Event = new(ContractDelegate)
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
func (it *ContractDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegate represents a Delegate event raised by the Contract contract.
type ContractDelegate struct {
	TokenID *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDelegate is a free log retrieval operation binding the contract event 0xb024739a6a34bded9b50497210d3eb50ccd1b03ddf8ab3738496175272c5bb65.
//
// Solidity: event Delegate(uint256 tokenID, address to)
func (_Contract *ContractFilterer) FilterDelegate(opts *bind.FilterOpts) (*ContractDelegateIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Delegate")
	if err != nil {
		return nil, err
	}
	return &ContractDelegateIterator{contract: _Contract.contract, event: "Delegate", logs: logs, sub: sub}, nil
}

// WatchDelegate is a free log subscription operation binding the contract event 0xb024739a6a34bded9b50497210d3eb50ccd1b03ddf8ab3738496175272c5bb65.
//
// Solidity: event Delegate(uint256 tokenID, address to)
func (_Contract *ContractFilterer) WatchDelegate(opts *bind.WatchOpts, sink chan<- *ContractDelegate) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Delegate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegate)
				if err := _Contract.contract.UnpackLog(event, "Delegate", log); err != nil {
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

// ParseDelegate is a log parse operation binding the contract event 0xb024739a6a34bded9b50497210d3eb50ccd1b03ddf8ab3738496175272c5bb65.
//
// Solidity: event Delegate(uint256 tokenID, address to)
func (_Contract *ContractFilterer) ParseDelegate(log types.Log) (*ContractDelegate, error) {
	event := new(ContractDelegate)
	if err := _Contract.contract.UnpackLog(event, "Delegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Contract contract.
type ContractInitializedIterator struct {
	Event *ContractInitialized // Event containing the contract specifics and raw log

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
func (it *ContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractInitialized)
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
		it.Event = new(ContractInitialized)
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
func (it *ContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractInitialized represents a Initialized event raised by the Contract contract.
type ContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractInitializedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractInitializedIterator{contract: _Contract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractInitialized) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractInitialized)
				if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) ParseInitialized(log types.Log) (*ContractInitialized, error) {
	event := new(ContractInitialized)
	if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNodeActivateIterator is returned from FilterNodeActivate and is used to iterate over the raw logs and unpacked data for NodeActivate events raised by the Contract contract.
type ContractNodeActivateIterator struct {
	Event *ContractNodeActivate // Event containing the contract specifics and raw log

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
func (it *ContractNodeActivateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNodeActivate)
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
		it.Event = new(ContractNodeActivate)
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
func (it *ContractNodeActivateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNodeActivateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNodeActivate represents a NodeActivate event raised by the Contract contract.
type ContractNodeActivate struct {
	Node common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNodeActivate is a free log retrieval operation binding the contract event 0x1b3699bd6aa5abdc4e768415f5023e2c35280c273edb4f358bad0c05f7b2701d.
//
// Solidity: event NodeActivate(address node)
func (_Contract *ContractFilterer) FilterNodeActivate(opts *bind.FilterOpts) (*ContractNodeActivateIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NodeActivate")
	if err != nil {
		return nil, err
	}
	return &ContractNodeActivateIterator{contract: _Contract.contract, event: "NodeActivate", logs: logs, sub: sub}, nil
}

// WatchNodeActivate is a free log subscription operation binding the contract event 0x1b3699bd6aa5abdc4e768415f5023e2c35280c273edb4f358bad0c05f7b2701d.
//
// Solidity: event NodeActivate(address node)
func (_Contract *ContractFilterer) WatchNodeActivate(opts *bind.WatchOpts, sink chan<- *ContractNodeActivate) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NodeActivate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNodeActivate)
				if err := _Contract.contract.UnpackLog(event, "NodeActivate", log); err != nil {
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

// ParseNodeActivate is a log parse operation binding the contract event 0x1b3699bd6aa5abdc4e768415f5023e2c35280c273edb4f358bad0c05f7b2701d.
//
// Solidity: event NodeActivate(address node)
func (_Contract *ContractFilterer) ParseNodeActivate(log types.Log) (*ContractNodeActivate, error) {
	event := new(ContractNodeActivate)
	if err := _Contract.contract.UnpackLog(event, "NodeActivate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNodeClaimIterator is returned from FilterNodeClaim and is used to iterate over the raw logs and unpacked data for NodeClaim events raised by the Contract contract.
type ContractNodeClaimIterator struct {
	Event *ContractNodeClaim // Event containing the contract specifics and raw log

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
func (it *ContractNodeClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNodeClaim)
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
		it.Event = new(ContractNodeClaim)
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
func (it *ContractNodeClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNodeClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNodeClaim represents a NodeClaim event raised by the Contract contract.
type ContractNodeClaim struct {
	Node    common.Address
	Claimer common.Address
	Rewards *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeClaim is a free log retrieval operation binding the contract event 0xea0babbef5f952932a55a6932b70dee7422d83300a5a59fae102bcf1e1136324.
//
// Solidity: event NodeClaim(address node, address claimer, uint256 rewards)
func (_Contract *ContractFilterer) FilterNodeClaim(opts *bind.FilterOpts) (*ContractNodeClaimIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NodeClaim")
	if err != nil {
		return nil, err
	}
	return &ContractNodeClaimIterator{contract: _Contract.contract, event: "NodeClaim", logs: logs, sub: sub}, nil
}

// WatchNodeClaim is a free log subscription operation binding the contract event 0xea0babbef5f952932a55a6932b70dee7422d83300a5a59fae102bcf1e1136324.
//
// Solidity: event NodeClaim(address node, address claimer, uint256 rewards)
func (_Contract *ContractFilterer) WatchNodeClaim(opts *bind.WatchOpts, sink chan<- *ContractNodeClaim) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NodeClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNodeClaim)
				if err := _Contract.contract.UnpackLog(event, "NodeClaim", log); err != nil {
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

// ParseNodeClaim is a log parse operation binding the contract event 0xea0babbef5f952932a55a6932b70dee7422d83300a5a59fae102bcf1e1136324.
//
// Solidity: event NodeClaim(address node, address claimer, uint256 rewards)
func (_Contract *ContractFilterer) ParseNodeClaim(log types.Log) (*ContractNodeClaim, error) {
	event := new(ContractNodeClaim)
	if err := _Contract.contract.UnpackLog(event, "NodeClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNodeClearIterator is returned from FilterNodeClear and is used to iterate over the raw logs and unpacked data for NodeClear events raised by the Contract contract.
type ContractNodeClearIterator struct {
	Event *ContractNodeClear // Event containing the contract specifics and raw log

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
func (it *ContractNodeClearIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNodeClear)
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
		it.Event = new(ContractNodeClear)
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
func (it *ContractNodeClearIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNodeClearIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNodeClear represents a NodeClear event raised by the Contract contract.
type ContractNodeClear struct {
	Node common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNodeClear is a free log retrieval operation binding the contract event 0x8a0859fa4a2e331800d512db6925d210facda82733207cb9fe49e7da954fc4aa.
//
// Solidity: event NodeClear(address node)
func (_Contract *ContractFilterer) FilterNodeClear(opts *bind.FilterOpts) (*ContractNodeClearIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NodeClear")
	if err != nil {
		return nil, err
	}
	return &ContractNodeClearIterator{contract: _Contract.contract, event: "NodeClear", logs: logs, sub: sub}, nil
}

// WatchNodeClear is a free log subscription operation binding the contract event 0x8a0859fa4a2e331800d512db6925d210facda82733207cb9fe49e7da954fc4aa.
//
// Solidity: event NodeClear(address node)
func (_Contract *ContractFilterer) WatchNodeClear(opts *bind.WatchOpts, sink chan<- *ContractNodeClear) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NodeClear")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNodeClear)
				if err := _Contract.contract.UnpackLog(event, "NodeClear", log); err != nil {
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

// ParseNodeClear is a log parse operation binding the contract event 0x8a0859fa4a2e331800d512db6925d210facda82733207cb9fe49e7da954fc4aa.
//
// Solidity: event NodeClear(address node)
func (_Contract *ContractFilterer) ParseNodeClear(log types.Log) (*ContractNodeClear, error) {
	event := new(ContractNodeClear)
	if err := _Contract.contract.UnpackLog(event, "NodeClear", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNodeModifyCommissionRateIterator is returned from FilterNodeModifyCommissionRate and is used to iterate over the raw logs and unpacked data for NodeModifyCommissionRate events raised by the Contract contract.
type ContractNodeModifyCommissionRateIterator struct {
	Event *ContractNodeModifyCommissionRate // Event containing the contract specifics and raw log

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
func (it *ContractNodeModifyCommissionRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNodeModifyCommissionRate)
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
		it.Event = new(ContractNodeModifyCommissionRate)
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
func (it *ContractNodeModifyCommissionRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNodeModifyCommissionRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNodeModifyCommissionRate represents a NodeModifyCommissionRate event raised by the Contract contract.
type ContractNodeModifyCommissionRate struct {
	Node           common.Address
	CommissionRate uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNodeModifyCommissionRate is a free log retrieval operation binding the contract event 0xb259504a8d6c681b25c45cd2a3825d7b986058852f5d01180dec6f9574f781f9.
//
// Solidity: event NodeModifyCommissionRate(address node, uint32 commissionRate)
func (_Contract *ContractFilterer) FilterNodeModifyCommissionRate(opts *bind.FilterOpts) (*ContractNodeModifyCommissionRateIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NodeModifyCommissionRate")
	if err != nil {
		return nil, err
	}
	return &ContractNodeModifyCommissionRateIterator{contract: _Contract.contract, event: "NodeModifyCommissionRate", logs: logs, sub: sub}, nil
}

// WatchNodeModifyCommissionRate is a free log subscription operation binding the contract event 0xb259504a8d6c681b25c45cd2a3825d7b986058852f5d01180dec6f9574f781f9.
//
// Solidity: event NodeModifyCommissionRate(address node, uint32 commissionRate)
func (_Contract *ContractFilterer) WatchNodeModifyCommissionRate(opts *bind.WatchOpts, sink chan<- *ContractNodeModifyCommissionRate) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NodeModifyCommissionRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNodeModifyCommissionRate)
				if err := _Contract.contract.UnpackLog(event, "NodeModifyCommissionRate", log); err != nil {
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

// ParseNodeModifyCommissionRate is a log parse operation binding the contract event 0xb259504a8d6c681b25c45cd2a3825d7b986058852f5d01180dec6f9574f781f9.
//
// Solidity: event NodeModifyCommissionRate(address node, uint32 commissionRate)
func (_Contract *ContractFilterer) ParseNodeModifyCommissionRate(log types.Log) (*ContractNodeModifyCommissionRate, error) {
	event := new(ContractNodeModifyCommissionRate)
	if err := _Contract.contract.UnpackLog(event, "NodeModifyCommissionRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNodeRegisterIterator is returned from FilterNodeRegister and is used to iterate over the raw logs and unpacked data for NodeRegister events raised by the Contract contract.
type ContractNodeRegisterIterator struct {
	Event *ContractNodeRegister // Event containing the contract specifics and raw log

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
func (it *ContractNodeRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNodeRegister)
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
		it.Event = new(ContractNodeRegister)
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
func (it *ContractNodeRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNodeRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNodeRegister represents a NodeRegister event raised by the Contract contract.
type ContractNodeRegister struct {
	Node common.Address
	Id   uint32
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNodeRegister is a free log retrieval operation binding the contract event 0xdcda586e1f19498517f947f53b7416526efbc24328864eb9071a9315ff8a86a9.
//
// Solidity: event NodeRegister(address node, uint32 id)
func (_Contract *ContractFilterer) FilterNodeRegister(opts *bind.FilterOpts) (*ContractNodeRegisterIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NodeRegister")
	if err != nil {
		return nil, err
	}
	return &ContractNodeRegisterIterator{contract: _Contract.contract, event: "NodeRegister", logs: logs, sub: sub}, nil
}

// WatchNodeRegister is a free log subscription operation binding the contract event 0xdcda586e1f19498517f947f53b7416526efbc24328864eb9071a9315ff8a86a9.
//
// Solidity: event NodeRegister(address node, uint32 id)
func (_Contract *ContractFilterer) WatchNodeRegister(opts *bind.WatchOpts, sink chan<- *ContractNodeRegister) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NodeRegister")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNodeRegister)
				if err := _Contract.contract.UnpackLog(event, "NodeRegister", log); err != nil {
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

// ParseNodeRegister is a log parse operation binding the contract event 0xdcda586e1f19498517f947f53b7416526efbc24328864eb9071a9315ff8a86a9.
//
// Solidity: event NodeRegister(address node, uint32 id)
func (_Contract *ContractFilterer) ParseNodeRegister(log types.Log) (*ContractNodeRegister, error) {
	event := new(ContractNodeRegister)
	if err := _Contract.contract.UnpackLog(event, "NodeRegister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNodeReportVerificationIterator is returned from FilterNodeReportVerification and is used to iterate over the raw logs and unpacked data for NodeReportVerification events raised by the Contract contract.
type ContractNodeReportVerificationIterator struct {
	Event *ContractNodeReportVerification // Event containing the contract specifics and raw log

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
func (it *ContractNodeReportVerificationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNodeReportVerification)
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
		it.Event = new(ContractNodeReportVerification)
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
func (it *ContractNodeReportVerificationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNodeReportVerificationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNodeReportVerification represents a NodeReportVerification event raised by the Contract contract.
type ContractNodeReportVerification struct {
	Node          common.Address
	AttestationID [32]byte
	Result        uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNodeReportVerification is a free log retrieval operation binding the contract event 0x52b7fb186080c2874dd6f580c1c5c01ddd95c9c2f414fd8c318a09ead3dc5536.
//
// Solidity: event NodeReportVerification(address node, bytes32 attestationID, uint8 result)
func (_Contract *ContractFilterer) FilterNodeReportVerification(opts *bind.FilterOpts) (*ContractNodeReportVerificationIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NodeReportVerification")
	if err != nil {
		return nil, err
	}
	return &ContractNodeReportVerificationIterator{contract: _Contract.contract, event: "NodeReportVerification", logs: logs, sub: sub}, nil
}

// WatchNodeReportVerification is a free log subscription operation binding the contract event 0x52b7fb186080c2874dd6f580c1c5c01ddd95c9c2f414fd8c318a09ead3dc5536.
//
// Solidity: event NodeReportVerification(address node, bytes32 attestationID, uint8 result)
func (_Contract *ContractFilterer) WatchNodeReportVerification(opts *bind.WatchOpts, sink chan<- *ContractNodeReportVerification) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NodeReportVerification")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNodeReportVerification)
				if err := _Contract.contract.UnpackLog(event, "NodeReportVerification", log); err != nil {
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

// ParseNodeReportVerification is a log parse operation binding the contract event 0x52b7fb186080c2874dd6f580c1c5c01ddd95c9c2f414fd8c318a09ead3dc5536.
//
// Solidity: event NodeReportVerification(address node, bytes32 attestationID, uint8 result)
func (_Contract *ContractFilterer) ParseNodeReportVerification(log types.Log) (*ContractNodeReportVerification, error) {
	event := new(ContractNodeReportVerification)
	if err := _Contract.contract.UnpackLog(event, "NodeReportVerification", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNodeReportVerificationBatchIterator is returned from FilterNodeReportVerificationBatch and is used to iterate over the raw logs and unpacked data for NodeReportVerificationBatch events raised by the Contract contract.
type ContractNodeReportVerificationBatchIterator struct {
	Event *ContractNodeReportVerificationBatch // Event containing the contract specifics and raw log

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
func (it *ContractNodeReportVerificationBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNodeReportVerificationBatch)
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
		it.Event = new(ContractNodeReportVerificationBatch)
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
func (it *ContractNodeReportVerificationBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNodeReportVerificationBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNodeReportVerificationBatch represents a NodeReportVerificationBatch event raised by the Contract contract.
type ContractNodeReportVerificationBatch struct {
	AttestationID [32]byte
	Infos         []IProtocolServiceVerificationInfo
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNodeReportVerificationBatch is a free log retrieval operation binding the contract event 0x16093eb01e7213cc4445469a1c8ded426fe3c5f18a72b6c4151b07dd4916cdf5.
//
// Solidity: event NodeReportVerificationBatch(bytes32 attestationID, (uint8,uint32,address,uint8,bytes32,bytes32)[] infos)
func (_Contract *ContractFilterer) FilterNodeReportVerificationBatch(opts *bind.FilterOpts) (*ContractNodeReportVerificationBatchIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NodeReportVerificationBatch")
	if err != nil {
		return nil, err
	}
	return &ContractNodeReportVerificationBatchIterator{contract: _Contract.contract, event: "NodeReportVerificationBatch", logs: logs, sub: sub}, nil
}

// WatchNodeReportVerificationBatch is a free log subscription operation binding the contract event 0x16093eb01e7213cc4445469a1c8ded426fe3c5f18a72b6c4151b07dd4916cdf5.
//
// Solidity: event NodeReportVerificationBatch(bytes32 attestationID, (uint8,uint32,address,uint8,bytes32,bytes32)[] infos)
func (_Contract *ContractFilterer) WatchNodeReportVerificationBatch(opts *bind.WatchOpts, sink chan<- *ContractNodeReportVerificationBatch) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NodeReportVerificationBatch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNodeReportVerificationBatch)
				if err := _Contract.contract.UnpackLog(event, "NodeReportVerificationBatch", log); err != nil {
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

// ParseNodeReportVerificationBatch is a log parse operation binding the contract event 0x16093eb01e7213cc4445469a1c8ded426fe3c5f18a72b6c4151b07dd4916cdf5.
//
// Solidity: event NodeReportVerificationBatch(bytes32 attestationID, (uint8,uint32,address,uint8,bytes32,bytes32)[] infos)
func (_Contract *ContractFilterer) ParseNodeReportVerificationBatch(log types.Log) (*ContractNodeReportVerificationBatch, error) {
	event := new(ContractNodeReportVerificationBatch)
	if err := _Contract.contract.UnpackLog(event, "NodeReportVerificationBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNodeSetClaimerIterator is returned from FilterNodeSetClaimer and is used to iterate over the raw logs and unpacked data for NodeSetClaimer events raised by the Contract contract.
type ContractNodeSetClaimerIterator struct {
	Event *ContractNodeSetClaimer // Event containing the contract specifics and raw log

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
func (it *ContractNodeSetClaimerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNodeSetClaimer)
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
		it.Event = new(ContractNodeSetClaimer)
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
func (it *ContractNodeSetClaimerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNodeSetClaimerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNodeSetClaimer represents a NodeSetClaimer event raised by the Contract contract.
type ContractNodeSetClaimer struct {
	Node    common.Address
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeSetClaimer is a free log retrieval operation binding the contract event 0x1454d6e227f2d31063b1c094100ff944573c2c21ddb0972ffc83313638583f8f.
//
// Solidity: event NodeSetClaimer(address node, address claimer)
func (_Contract *ContractFilterer) FilterNodeSetClaimer(opts *bind.FilterOpts) (*ContractNodeSetClaimerIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NodeSetClaimer")
	if err != nil {
		return nil, err
	}
	return &ContractNodeSetClaimerIterator{contract: _Contract.contract, event: "NodeSetClaimer", logs: logs, sub: sub}, nil
}

// WatchNodeSetClaimer is a free log subscription operation binding the contract event 0x1454d6e227f2d31063b1c094100ff944573c2c21ddb0972ffc83313638583f8f.
//
// Solidity: event NodeSetClaimer(address node, address claimer)
func (_Contract *ContractFilterer) WatchNodeSetClaimer(opts *bind.WatchOpts, sink chan<- *ContractNodeSetClaimer) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NodeSetClaimer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNodeSetClaimer)
				if err := _Contract.contract.UnpackLog(event, "NodeSetClaimer", log); err != nil {
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

// ParseNodeSetClaimer is a log parse operation binding the contract event 0x1454d6e227f2d31063b1c094100ff944573c2c21ddb0972ffc83313638583f8f.
//
// Solidity: event NodeSetClaimer(address node, address claimer)
func (_Contract *ContractFilterer) ParseNodeSetClaimer(log types.Log) (*ContractNodeSetClaimer, error) {
	event := new(ContractNodeSetClaimer)
	if err := _Contract.contract.UnpackLog(event, "NodeSetClaimer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNodeSlashIterator is returned from FilterNodeSlash and is used to iterate over the raw logs and unpacked data for NodeSlash events raised by the Contract contract.
type ContractNodeSlashIterator struct {
	Event *ContractNodeSlash // Event containing the contract specifics and raw log

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
func (it *ContractNodeSlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNodeSlash)
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
		it.Event = new(ContractNodeSlash)
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
func (it *ContractNodeSlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNodeSlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNodeSlash represents a NodeSlash event raised by the Contract contract.
type ContractNodeSlash struct {
	Node          common.Address
	AttestationID [32]byte
	Rewards       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNodeSlash is a free log retrieval operation binding the contract event 0x2ec1d25826f3279dd30b2af581bf1fb9fdca750ac9df9cc193d194704ea0d76d.
//
// Solidity: event NodeSlash(address node, bytes32 attestationID, uint256 rewards)
func (_Contract *ContractFilterer) FilterNodeSlash(opts *bind.FilterOpts) (*ContractNodeSlashIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NodeSlash")
	if err != nil {
		return nil, err
	}
	return &ContractNodeSlashIterator{contract: _Contract.contract, event: "NodeSlash", logs: logs, sub: sub}, nil
}

// WatchNodeSlash is a free log subscription operation binding the contract event 0x2ec1d25826f3279dd30b2af581bf1fb9fdca750ac9df9cc193d194704ea0d76d.
//
// Solidity: event NodeSlash(address node, bytes32 attestationID, uint256 rewards)
func (_Contract *ContractFilterer) WatchNodeSlash(opts *bind.WatchOpts, sink chan<- *ContractNodeSlash) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NodeSlash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNodeSlash)
				if err := _Contract.contract.UnpackLog(event, "NodeSlash", log); err != nil {
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

// ParseNodeSlash is a log parse operation binding the contract event 0x2ec1d25826f3279dd30b2af581bf1fb9fdca750ac9df9cc193d194704ea0d76d.
//
// Solidity: event NodeSlash(address node, bytes32 attestationID, uint256 rewards)
func (_Contract *ContractFilterer) ParseNodeSlash(log types.Log) (*ContractNodeSlash, error) {
	event := new(ContractNodeSlash)
	if err := _Contract.contract.UnpackLog(event, "NodeSlash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRedelegateIterator is returned from FilterRedelegate and is used to iterate over the raw logs and unpacked data for Redelegate events raised by the Contract contract.
type ContractRedelegateIterator struct {
	Event *ContractRedelegate // Event containing the contract specifics and raw log

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
func (it *ContractRedelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRedelegate)
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
		it.Event = new(ContractRedelegate)
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
func (it *ContractRedelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRedelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRedelegate represents a Redelegate event raised by the Contract contract.
type ContractRedelegate struct {
	TokenID *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRedelegate is a free log retrieval operation binding the contract event 0xa6494190270aff4668179671a838f4079a7b83070d0dc4aafc1252956fa93d13.
//
// Solidity: event Redelegate(uint256 tokenID, address to)
func (_Contract *ContractFilterer) FilterRedelegate(opts *bind.FilterOpts) (*ContractRedelegateIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Redelegate")
	if err != nil {
		return nil, err
	}
	return &ContractRedelegateIterator{contract: _Contract.contract, event: "Redelegate", logs: logs, sub: sub}, nil
}

// WatchRedelegate is a free log subscription operation binding the contract event 0xa6494190270aff4668179671a838f4079a7b83070d0dc4aafc1252956fa93d13.
//
// Solidity: event Redelegate(uint256 tokenID, address to)
func (_Contract *ContractFilterer) WatchRedelegate(opts *bind.WatchOpts, sink chan<- *ContractRedelegate) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Redelegate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRedelegate)
				if err := _Contract.contract.UnpackLog(event, "Redelegate", log); err != nil {
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

// ParseRedelegate is a log parse operation binding the contract event 0xa6494190270aff4668179671a838f4079a7b83070d0dc4aafc1252956fa93d13.
//
// Solidity: event Redelegate(uint256 tokenID, address to)
func (_Contract *ContractFilterer) ParseRedelegate(log types.Log) (*ContractRedelegate, error) {
	event := new(ContractRedelegate)
	if err := _Contract.contract.UnpackLog(event, "Redelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTeeReportAttestationsIterator is returned from FilterTeeReportAttestations and is used to iterate over the raw logs and unpacked data for TeeReportAttestations events raised by the Contract contract.
type ContractTeeReportAttestationsIterator struct {
	Event *ContractTeeReportAttestations // Event containing the contract specifics and raw log

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
func (it *ContractTeeReportAttestationsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTeeReportAttestations)
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
		it.Event = new(ContractTeeReportAttestations)
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
func (it *ContractTeeReportAttestationsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTeeReportAttestationsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTeeReportAttestations represents a TeeReportAttestations event raised by the Contract contract.
type ContractTeeReportAttestations struct {
	Tee              common.Address
	AttestationIDs   [][32]byte
	AttestationInfos []string
	RequestID        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTeeReportAttestations is a free log retrieval operation binding the contract event 0x89a3b784b99180438f3b2027aa89e97c3c3ed66e8dc78a555d7013b39caf1a89.
//
// Solidity: event TeeReportAttestations(address tee, bytes32[] attestationIDs, string[] attestationInfos, uint256 requestID)
func (_Contract *ContractFilterer) FilterTeeReportAttestations(opts *bind.FilterOpts) (*ContractTeeReportAttestationsIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TeeReportAttestations")
	if err != nil {
		return nil, err
	}
	return &ContractTeeReportAttestationsIterator{contract: _Contract.contract, event: "TeeReportAttestations", logs: logs, sub: sub}, nil
}

// WatchTeeReportAttestations is a free log subscription operation binding the contract event 0x89a3b784b99180438f3b2027aa89e97c3c3ed66e8dc78a555d7013b39caf1a89.
//
// Solidity: event TeeReportAttestations(address tee, bytes32[] attestationIDs, string[] attestationInfos, uint256 requestID)
func (_Contract *ContractFilterer) WatchTeeReportAttestations(opts *bind.WatchOpts, sink chan<- *ContractTeeReportAttestations) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TeeReportAttestations")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTeeReportAttestations)
				if err := _Contract.contract.UnpackLog(event, "TeeReportAttestations", log); err != nil {
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

// ParseTeeReportAttestations is a log parse operation binding the contract event 0x89a3b784b99180438f3b2027aa89e97c3c3ed66e8dc78a555d7013b39caf1a89.
//
// Solidity: event TeeReportAttestations(address tee, bytes32[] attestationIDs, string[] attestationInfos, uint256 requestID)
func (_Contract *ContractFilterer) ParseTeeReportAttestations(log types.Log) (*ContractTeeReportAttestations, error) {
	event := new(ContractTeeReportAttestations)
	if err := _Contract.contract.UnpackLog(event, "TeeReportAttestations", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTeeSlashIterator is returned from FilterTeeSlash and is used to iterate over the raw logs and unpacked data for TeeSlash events raised by the Contract contract.
type ContractTeeSlashIterator struct {
	Event *ContractTeeSlash // Event containing the contract specifics and raw log

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
func (it *ContractTeeSlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTeeSlash)
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
		it.Event = new(ContractTeeSlash)
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
func (it *ContractTeeSlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTeeSlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTeeSlash represents a TeeSlash event raised by the Contract contract.
type ContractTeeSlash struct {
	Tee           common.Address
	AttestationID [32]byte
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTeeSlash is a free log retrieval operation binding the contract event 0xcb2bfd0c073ab6c31355ed348da3a1b661a3a578be5a450eaa8d420a9323efaf.
//
// Solidity: event TeeSlash(address tee, bytes32 attestationID, uint256 amount)
func (_Contract *ContractFilterer) FilterTeeSlash(opts *bind.FilterOpts) (*ContractTeeSlashIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TeeSlash")
	if err != nil {
		return nil, err
	}
	return &ContractTeeSlashIterator{contract: _Contract.contract, event: "TeeSlash", logs: logs, sub: sub}, nil
}

// WatchTeeSlash is a free log subscription operation binding the contract event 0xcb2bfd0c073ab6c31355ed348da3a1b661a3a578be5a450eaa8d420a9323efaf.
//
// Solidity: event TeeSlash(address tee, bytes32 attestationID, uint256 amount)
func (_Contract *ContractFilterer) WatchTeeSlash(opts *bind.WatchOpts, sink chan<- *ContractTeeSlash) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TeeSlash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTeeSlash)
				if err := _Contract.contract.UnpackLog(event, "TeeSlash", log); err != nil {
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

// ParseTeeSlash is a log parse operation binding the contract event 0xcb2bfd0c073ab6c31355ed348da3a1b661a3a578be5a450eaa8d420a9323efaf.
//
// Solidity: event TeeSlash(address tee, bytes32 attestationID, uint256 amount)
func (_Contract *ContractFilterer) ParseTeeSlash(log types.Log) (*ContractTeeSlash, error) {
	event := new(ContractTeeSlash)
	if err := _Contract.contract.UnpackLog(event, "TeeSlash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTeeStakeIterator is returned from FilterTeeStake and is used to iterate over the raw logs and unpacked data for TeeStake events raised by the Contract contract.
type ContractTeeStakeIterator struct {
	Event *ContractTeeStake // Event containing the contract specifics and raw log

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
func (it *ContractTeeStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTeeStake)
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
		it.Event = new(ContractTeeStake)
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
func (it *ContractTeeStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTeeStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTeeStake represents a TeeStake event raised by the Contract contract.
type ContractTeeStake struct {
	Tee    common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTeeStake is a free log retrieval operation binding the contract event 0x1b457cf8cf55b14c9e9f6ca27cea4ca8eb08d8d1e032db6923681cf45093bc65.
//
// Solidity: event TeeStake(address tee, uint256 amount)
func (_Contract *ContractFilterer) FilterTeeStake(opts *bind.FilterOpts) (*ContractTeeStakeIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TeeStake")
	if err != nil {
		return nil, err
	}
	return &ContractTeeStakeIterator{contract: _Contract.contract, event: "TeeStake", logs: logs, sub: sub}, nil
}

// WatchTeeStake is a free log subscription operation binding the contract event 0x1b457cf8cf55b14c9e9f6ca27cea4ca8eb08d8d1e032db6923681cf45093bc65.
//
// Solidity: event TeeStake(address tee, uint256 amount)
func (_Contract *ContractFilterer) WatchTeeStake(opts *bind.WatchOpts, sink chan<- *ContractTeeStake) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TeeStake")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTeeStake)
				if err := _Contract.contract.UnpackLog(event, "TeeStake", log); err != nil {
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

// ParseTeeStake is a log parse operation binding the contract event 0x1b457cf8cf55b14c9e9f6ca27cea4ca8eb08d8d1e032db6923681cf45093bc65.
//
// Solidity: event TeeStake(address tee, uint256 amount)
func (_Contract *ContractFilterer) ParseTeeStake(log types.Log) (*ContractTeeStake, error) {
	event := new(ContractTeeStake)
	if err := _Contract.contract.UnpackLog(event, "TeeStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTeeUnstakeIterator is returned from FilterTeeUnstake and is used to iterate over the raw logs and unpacked data for TeeUnstake events raised by the Contract contract.
type ContractTeeUnstakeIterator struct {
	Event *ContractTeeUnstake // Event containing the contract specifics and raw log

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
func (it *ContractTeeUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTeeUnstake)
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
		it.Event = new(ContractTeeUnstake)
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
func (it *ContractTeeUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTeeUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTeeUnstake represents a TeeUnstake event raised by the Contract contract.
type ContractTeeUnstake struct {
	Tee    common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTeeUnstake is a free log retrieval operation binding the contract event 0x8348f568c1ea68de8db3504a38def7ab3a5c7d12d8f913fff88c0bf7e10794c9.
//
// Solidity: event TeeUnstake(address tee, uint256 amount)
func (_Contract *ContractFilterer) FilterTeeUnstake(opts *bind.FilterOpts) (*ContractTeeUnstakeIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TeeUnstake")
	if err != nil {
		return nil, err
	}
	return &ContractTeeUnstakeIterator{contract: _Contract.contract, event: "TeeUnstake", logs: logs, sub: sub}, nil
}

// WatchTeeUnstake is a free log subscription operation binding the contract event 0x8348f568c1ea68de8db3504a38def7ab3a5c7d12d8f913fff88c0bf7e10794c9.
//
// Solidity: event TeeUnstake(address tee, uint256 amount)
func (_Contract *ContractFilterer) WatchTeeUnstake(opts *bind.WatchOpts, sink chan<- *ContractTeeUnstake) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TeeUnstake")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTeeUnstake)
				if err := _Contract.contract.UnpackLog(event, "TeeUnstake", log); err != nil {
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

// ParseTeeUnstake is a log parse operation binding the contract event 0x8348f568c1ea68de8db3504a38def7ab3a5c7d12d8f913fff88c0bf7e10794c9.
//
// Solidity: event TeeUnstake(address tee, uint256 amount)
func (_Contract *ContractFilterer) ParseTeeUnstake(log types.Log) (*ContractTeeUnstake, error) {
	event := new(ContractTeeUnstake)
	if err := _Contract.contract.UnpackLog(event, "TeeUnstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUndelegateIterator is returned from FilterUndelegate and is used to iterate over the raw logs and unpacked data for Undelegate events raised by the Contract contract.
type ContractUndelegateIterator struct {
	Event *ContractUndelegate // Event containing the contract specifics and raw log

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
func (it *ContractUndelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUndelegate)
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
		it.Event = new(ContractUndelegate)
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
func (it *ContractUndelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUndelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUndelegate represents a Undelegate event raised by the Contract contract.
type ContractUndelegate struct {
	TokenID *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUndelegate is a free log retrieval operation binding the contract event 0xc29e8cd64b240d860ccd82cacb9453b926a03ba48a57753a663a6facde47105d.
//
// Solidity: event Undelegate(uint256 tokenID, address to)
func (_Contract *ContractFilterer) FilterUndelegate(opts *bind.FilterOpts) (*ContractUndelegateIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Undelegate")
	if err != nil {
		return nil, err
	}
	return &ContractUndelegateIterator{contract: _Contract.contract, event: "Undelegate", logs: logs, sub: sub}, nil
}

// WatchUndelegate is a free log subscription operation binding the contract event 0xc29e8cd64b240d860ccd82cacb9453b926a03ba48a57753a663a6facde47105d.
//
// Solidity: event Undelegate(uint256 tokenID, address to)
func (_Contract *ContractFilterer) WatchUndelegate(opts *bind.WatchOpts, sink chan<- *ContractUndelegate) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Undelegate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUndelegate)
				if err := _Contract.contract.UnpackLog(event, "Undelegate", log); err != nil {
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

// ParseUndelegate is a log parse operation binding the contract event 0xc29e8cd64b240d860ccd82cacb9453b926a03ba48a57753a663a6facde47105d.
//
// Solidity: event Undelegate(uint256 tokenID, address to)
func (_Contract *ContractFilterer) ParseUndelegate(log types.Log) (*ContractUndelegate, error) {
	event := new(ContractUndelegate)
	if err := _Contract.contract.UnpackLog(event, "Undelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUpdateSettingsAddressIterator is returned from FilterUpdateSettingsAddress and is used to iterate over the raw logs and unpacked data for UpdateSettingsAddress events raised by the Contract contract.
type ContractUpdateSettingsAddressIterator struct {
	Event *ContractUpdateSettingsAddress // Event containing the contract specifics and raw log

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
func (it *ContractUpdateSettingsAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdateSettingsAddress)
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
		it.Event = new(ContractUpdateSettingsAddress)
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
func (it *ContractUpdateSettingsAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdateSettingsAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdateSettingsAddress represents a UpdateSettingsAddress event raised by the Contract contract.
type ContractUpdateSettingsAddress struct {
	Settings common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateSettingsAddress is a free log retrieval operation binding the contract event 0x6bdb6edfe23f0a97be9dda487ad7ee5fc6590e2058caf57a9abd481269d4a2e3.
//
// Solidity: event UpdateSettingsAddress(address settings)
func (_Contract *ContractFilterer) FilterUpdateSettingsAddress(opts *bind.FilterOpts) (*ContractUpdateSettingsAddressIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdateSettingsAddress")
	if err != nil {
		return nil, err
	}
	return &ContractUpdateSettingsAddressIterator{contract: _Contract.contract, event: "UpdateSettingsAddress", logs: logs, sub: sub}, nil
}

// WatchUpdateSettingsAddress is a free log subscription operation binding the contract event 0x6bdb6edfe23f0a97be9dda487ad7ee5fc6590e2058caf57a9abd481269d4a2e3.
//
// Solidity: event UpdateSettingsAddress(address settings)
func (_Contract *ContractFilterer) WatchUpdateSettingsAddress(opts *bind.WatchOpts, sink chan<- *ContractUpdateSettingsAddress) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdateSettingsAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdateSettingsAddress)
				if err := _Contract.contract.UnpackLog(event, "UpdateSettingsAddress", log); err != nil {
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

// ParseUpdateSettingsAddress is a log parse operation binding the contract event 0x6bdb6edfe23f0a97be9dda487ad7ee5fc6590e2058caf57a9abd481269d4a2e3.
//
// Solidity: event UpdateSettingsAddress(address settings)
func (_Contract *ContractFilterer) ParseUpdateSettingsAddress(log types.Log) (*ContractUpdateSettingsAddress, error) {
	event := new(ContractUpdateSettingsAddress)
	if err := _Contract.contract.UnpackLog(event, "UpdateSettingsAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUpdateVrfAddressIterator is returned from FilterUpdateVrfAddress and is used to iterate over the raw logs and unpacked data for UpdateVrfAddress events raised by the Contract contract.
type ContractUpdateVrfAddressIterator struct {
	Event *ContractUpdateVrfAddress // Event containing the contract specifics and raw log

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
func (it *ContractUpdateVrfAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdateVrfAddress)
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
		it.Event = new(ContractUpdateVrfAddress)
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
func (it *ContractUpdateVrfAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdateVrfAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdateVrfAddress represents a UpdateVrfAddress event raised by the Contract contract.
type ContractUpdateVrfAddress struct {
	Vrf common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpdateVrfAddress is a free log retrieval operation binding the contract event 0x891cc243fd5e5ff1f52df8dae4e723c3a5ccb1ac26346ca41c55b82f7f7162eb.
//
// Solidity: event UpdateVrfAddress(address vrf)
func (_Contract *ContractFilterer) FilterUpdateVrfAddress(opts *bind.FilterOpts) (*ContractUpdateVrfAddressIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdateVrfAddress")
	if err != nil {
		return nil, err
	}
	return &ContractUpdateVrfAddressIterator{contract: _Contract.contract, event: "UpdateVrfAddress", logs: logs, sub: sub}, nil
}

// WatchUpdateVrfAddress is a free log subscription operation binding the contract event 0x891cc243fd5e5ff1f52df8dae4e723c3a5ccb1ac26346ca41c55b82f7f7162eb.
//
// Solidity: event UpdateVrfAddress(address vrf)
func (_Contract *ContractFilterer) WatchUpdateVrfAddress(opts *bind.WatchOpts, sink chan<- *ContractUpdateVrfAddress) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdateVrfAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdateVrfAddress)
				if err := _Contract.contract.UnpackLog(event, "UpdateVrfAddress", log); err != nil {
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

// ParseUpdateVrfAddress is a log parse operation binding the contract event 0x891cc243fd5e5ff1f52df8dae4e723c3a5ccb1ac26346ca41c55b82f7f7162eb.
//
// Solidity: event UpdateVrfAddress(address vrf)
func (_Contract *ContractFilterer) ParseUpdateVrfAddress(log types.Log) (*ContractUpdateVrfAddress, error) {
	event := new(ContractUpdateVrfAddress)
	if err := _Contract.contract.UnpackLog(event, "UpdateVrfAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

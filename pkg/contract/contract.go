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

// CarvProtocolServicecampaign is an auto generated low-level Go binding around an user-defined struct.
type CarvProtocolServicecampaign struct {
	CampaignId        string
	Url               string
	Creator           common.Address
	CampaignType      uint8
	RewardContract    common.Address
	RewardTotalAmount *big.Int
	RewardCount       *big.Int
	Status            uint8
	StartTime         *big.Int
	EndTime           *big.Int
	Requirements      string
}

// CarvProtocolServicereward is an auto generated low-level Go binding around an user-defined struct.
type CarvProtocolServicereward struct {
	CampaignId      string
	UserAddress     common.Address
	RewardAmount    *big.Int
	TotalNum        *big.Int
	ContractAddress common.Address
	ContractType    uint8
}

// CarvProtocolServiceuser is an auto generated low-level Go binding around an user-defined struct.
type CarvProtocolServiceuser struct {
	TokenId         *big.Int
	UserProfilePath string
	ProfileVersion  *big.Int
	Signature       []byte
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"PackedPtrLen__LenOverflow\",\"type\":\"error\",\"inputs\":[]},{\"name\":\"PackedPtrLen__PtrOverflow\",\"type\":\"error\",\"inputs\":[]},{\"name\":\"Slice__OutOfBounds\",\"type\":\"error\",\"inputs\":[]},{\"name\":\"Approval\",\"type\":\"event\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false,\"signature\":\"0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925\"},{\"name\":\"ApprovalForAll\",\"type\":\"event\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false,\"signature\":\"0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31\"},{\"name\":\"Initialized\",\"type\":\"event\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false,\"signature\":\"0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498\"},{\"name\":\"Minted\",\"type\":\"event\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token_id\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false,\"signature\":\"0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe\"},{\"name\":\"ProfixPayed\",\"type\":\"event\",\"inputs\":[{\"name\":\"erc20_address\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"from_address\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"to_address\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false,\"signature\":\"0x3832ba0f9bc9de56bd1f974e665b53f2643ed48a47b45c3d36976bc1ea314599\"},{\"name\":\"ReportTeeAttestation\",\"type\":\"event\",\"inputs\":[{\"name\":\"tee_address\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"campaign_id\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"attestation_id\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"attestation\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false,\"signature\":\"0x99a038e9d345d0b12130b3b1fb003bf8f2d3a5c27ce2a800bbb1608efff6c591\"},{\"name\":\"RewardPayed\",\"type\":\"event\",\"inputs\":[{\"name\":\"erc20_address\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"from_address\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"to_address\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false,\"signature\":\"0x704db66c978a7472842ddae0d478a8a7d123fb6ad449e4426b5e0f6a22081de7\"},{\"name\":\"RoleAdminChanged\",\"type\":\"event\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false,\"signature\":\"0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff\"},{\"name\":\"RoleGranted\",\"type\":\"event\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false,\"signature\":\"0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d\"},{\"name\":\"RoleRevoked\",\"type\":\"event\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false,\"signature\":\"0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b\"},{\"name\":\"SetIdentitiesRoot\",\"type\":\"event\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"identitiesRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false,\"signature\":\"0x21b5a040ca6f61b8bcacb8f423e25ce46e88932f887ce92e60343369c20ec06f\"},{\"name\":\"SubmitCampaign\",\"type\":\"event\",\"inputs\":[{\"name\":\"contract_address\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"campaign_id\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"requirements\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false,\"signature\":\"0xe8e692af5f416c4637c9307d7f08b7697fd8f4398f55103ce03ddd9388ccd84d\"},{\"name\":\"Transfer\",\"type\":\"event\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false,\"signature\":\"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\"},{\"name\":\"UserCampaignData\",\"type\":\"event\",\"inputs\":[{\"name\":\"carv_id\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"campaign_id\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"campaign_info\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false,\"signature\":\"0x83eda6edb918aac3f35c8d616a6198dbb75eeca3248e0efec05e66a30d511191\"},{\"name\":\"VerifyAttestation\",\"type\":\"event\",\"inputs\":[{\"name\":\"verifier_address\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"attestation_id\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"result\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false,\"signature\":\"0x937b4d1b5b461852fdbbb174b68adefb42155e668c05a1261fae11a2d22a240a\"},{\"name\":\"VerifyAttestationBatch\",\"type\":\"event\",\"inputs\":[{\"name\":\"verifier_address\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"attestation_ids\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"},{\"name\":\"results\",\"type\":\"bool[]\",\"indexed\":false,\"internalType\":\"bool[]\"}],\"anonymous\":false,\"signature\":\"0x04798cc2c828443d2d9ef127deaa1f73e782d8a9fffbc9db2bcca53e12765af5\"},{\"name\":\"DEFAULT_ADMIN_ROLE\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"value\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"internalType\":\"bytes32\"}],\"constant\":true,\"signature\":\"0xa217fddf\",\"stateMutability\":\"view\"},{\"name\":\"TEE_ROLE\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"value\":\"0xd16a327e6f5c32c69c8282ab355bc8a366432cf60ee1165bc5198414ca1b06c7\",\"internalType\":\"bytes32\"}],\"constant\":true,\"signature\":\"0x495eb828\",\"stateMutability\":\"view\"},{\"name\":\"VERIFIER_ROLE\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"value\":\"0x0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea09\",\"internalType\":\"bytes32\"}],\"constant\":true,\"signature\":\"0xe7705db6\",\"stateMutability\":\"view\"},{\"name\":\"__CarvProtocolService_init\",\"type\":\"function\",\"inputs\":[{\"name\":\"rewards_address\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"signature\":\"0xa28bfc43\",\"stateMutability\":\"nonpayable\"},{\"name\":\"add_tee_role\",\"type\":\"function\",\"inputs\":[{\"name\":\"tee_address\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"signature\":\"0xa35d20be\",\"stateMutability\":\"nonpayable\"},{\"name\":\"add_verifier_role\",\"type\":\"function\",\"inputs\":[{\"name\":\"verifier_address\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"signature\":\"0xad86d545\",\"stateMutability\":\"nonpayable\"},{\"name\":\"approve\",\"type\":\"function\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"signature\":\"0x095ea7b3\",\"stateMutability\":\"nonpayable\"},{\"name\":\"balanceOf\",\"type\":\"function\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"constant\":true,\"signature\":\"0x70a08231\",\"stateMutability\":\"view\"},{\"name\":\"getApproved\",\"type\":\"function\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"constant\":true,\"signature\":\"0x081812fc\",\"stateMutability\":\"view\"},{\"name\":\"getIdentitiesRoot\",\"type\":\"function\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"constant\":true,\"signature\":\"0x45692b82\",\"stateMutability\":\"view\"},{\"name\":\"getRoleAdmin\",\"type\":\"function\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"constant\":true,\"signature\":\"0x248a9ca3\",\"stateMutability\":\"view\"},{\"name\":\"get_attestation_result\",\"type\":\"function\",\"inputs\":[{\"name\":\"attestation_id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"constant\":true,\"signature\":\"0x49e9e834\",\"stateMutability\":\"view\"},{\"name\":\"get_pay_address\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"value\":\"0x0000000000000000000000000000000000000000\",\"internalType\":\"address\"}],\"constant\":true,\"signature\":\"0x1d17e1bd\",\"stateMutability\":\"view\"},{\"name\":\"get_proof_list\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"value\":[],\"internalType\":\"bytes32[]\"}],\"constant\":true,\"signature\":\"0x964b1f6e\",\"stateMutability\":\"view\"},{\"name\":\"get_user_by_address\",\"type\":\"function\",\"inputs\":[{\"name\":\"user_address\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"name\":\"token_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"user_profile_path\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"profile_version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"internalType\":\"structCarvProtocolService.user\"}],\"constant\":true,\"signature\":\"0xda7a2928\",\"stateMutability\":\"view\"},{\"name\":\"get_verifier_list\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"value\":[],\"internalType\":\"address[]\"}],\"constant\":true,\"signature\":\"0x70b3542d\",\"stateMutability\":\"view\"},{\"name\":\"grantRole\",\"type\":\"function\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"signature\":\"0x2f2ff15d\",\"stateMutability\":\"nonpayable\"},{\"name\":\"hasRole\",\"type\":\"function\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"constant\":true,\"signature\":\"0x91d14854\",\"stateMutability\":\"view\"},{\"name\":\"isApprovedForAll\",\"type\":\"function\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"constant\":true,\"signature\":\"0xe985e9c5\",\"stateMutability\":\"view\"},{\"name\":\"join_campaign\",\"type\":\"function\",\"inputs\":[{\"name\":\"carv_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"campaign_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"join_campaign_info\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"signature\":\"0x09a9c7f0\",\"stateMutability\":\"nonpayable\"},{\"name\":\"name\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"value\":\"\",\"internalType\":\"string\"}],\"constant\":true,\"signature\":\"0x06fdde03\",\"stateMutability\":\"view\"},{\"name\":\"ownerOf\",\"type\":\"function\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"constant\":true,\"signature\":\"0x6352211e\",\"stateMutability\":\"view\"},{\"name\":\"renounceRole\",\"type\":\"function\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"signature\":\"0x36568abe\",\"stateMutability\":\"nonpayable\"},{\"name\":\"report_tee_attestation\",\"type\":\"function\",\"inputs\":[{\"name\":\"campaign_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attestation\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"signature\":\"0xe97a9987\",\"stateMutability\":\"nonpayable\"},{\"name\":\"revokeRole\",\"type\":\"function\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"signature\":\"0xd547741f\",\"stateMutability\":\"nonpayable\"},{\"name\":\"safeTransferFrom\",\"type\":\"function\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"signature\":\"0x42842e0e\",\"stateMutability\":\"nonpayable\"},{\"name\":\"safeTransferFrom\",\"type\":\"function\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"signature\":\"0xb88d4fde\",\"stateMutability\":\"nonpayable\"},{\"name\":\"setApprovalForAll\",\"type\":\"function\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"signature\":\"0xa22cb465\",\"stateMutability\":\"nonpayable\"},{\"name\":\"setIdentitiesRootFunc\",\"type\":\"function\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiIdentitiesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"signature\":\"0xff30a118\",\"stateMutability\":\"nonpayable\"},{\"name\":\"set_identities_root\",\"type\":\"function\",\"inputs\":[{\"name\":\"user_address\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"user_profile_path\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"profile_version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiIdentitiesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"signature\":\"0x1a7f9955\",\"stateMutability\":\"nonpayable\"},{\"name\":\"set_pay_address\",\"type\":\"function\",\"inputs\":[{\"name\":\"pay_address\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"signature\":\"0x7e3461a3\",\"stateMutability\":\"nonpayable\"},{\"name\":\"submit_campaign\",\"type\":\"function\",\"inputs\":[{\"name\":\"reward_info\",\"type\":\"tuple\",\"components\":[{\"name\":\"campaign_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"user_address\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reward_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"total_num\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contract_address\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"contract_type\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"internalType\":\"structCarvProtocolService.reward\"},{\"name\":\"campaign_info\",\"type\":\"tuple\",\"components\":[{\"name\":\"campaign_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"url\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"campaign_type\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"reward_contract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reward_total_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reward_count\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"start_time\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"end_time\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requirements\",\"type\":\"string\",\"internalType\":\"string\"}],\"internalType\":\"structCarvProtocolService.campaign\"}],\"outputs\":[],\"payable\":true,\"signature\":\"0x3736c83f\",\"stateMutability\":\"payable\"},{\"name\":\"supportsInterface\",\"type\":\"function\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"constant\":true,\"signature\":\"0x01ffc9a7\",\"stateMutability\":\"view\"},{\"name\":\"symbol\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"value\":\"\",\"internalType\":\"string\"}],\"constant\":true,\"signature\":\"0x95d89b41\",\"stateMutability\":\"view\"},{\"name\":\"tokenURI\",\"type\":\"function\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"constant\":true,\"signature\":\"0xc87b56dd\",\"stateMutability\":\"view\"},{\"name\":\"transferFrom\",\"type\":\"function\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"signature\":\"0x23b872dd\",\"stateMutability\":\"nonpayable\"},{\"name\":\"verifyIdentitiesBinding\",\"type\":\"function\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nftOwnerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sigAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"userIDs\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"multiIdentitiesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"constant\":true,\"signature\":\"0x49d03037\",\"stateMutability\":\"view\"},{\"name\":\"verify_attestation\",\"type\":\"function\",\"inputs\":[{\"name\":\"attestation_id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"signature\":\"0x578907ca\",\"stateMutability\":\"nonpayable\"},{\"name\":\"verify_attestation_batch\",\"type\":\"function\",\"inputs\":[{\"name\":\"attestation_ids\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"results\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"signature\":\"0xc1f75848\",\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x608060405234801561001057600080fd5b50614200806100206000396000f3fe6080604052600436106101c05760003560e01c806301ffc9a7146101c557806306fdde03146101fa578063081812fc1461021c578063095ea7b31461025457806309a9c7f0146102765780631a7f9955146102965780631d17e1bd146102b657806323b872dd146102d4578063248a9ca3146102f45780632f2ff15d1461032257806336568abe146103425780633736c83f1461036257806342842e0e1461037557806345692b8214610395578063495eb828146103c257806349d03037146103e457806349e9e83414610404578063578907ca146104345780636352211e1461045457806370a082311461047457806370b3542d146104945780637e3461a3146104b657806391d14854146104d657806395d89b41146104f6578063964b1f6e1461050b578063a217fddf1461052d578063a22cb46514610542578063a28bfc4314610562578063a35d20be14610582578063ad86d545146105a2578063b88d4fde146105c2578063c1f75848146105e2578063c87b56dd14610602578063d547741f14610622578063da7a292814610642578063e7705db61461066f578063e97a998714610691578063e985e9c5146106b1578063ff30a118146106d1575b600080fd5b3480156101d157600080fd5b506101e56101e0366004612ffe565b6106f1565b60405190151581526020015b60405180910390f35b34801561020657600080fd5b5061020f610702565b6040516101f1919061306b565b34801561022857600080fd5b5061023c61023736600461307e565b610794565b6040516001600160a01b0390911681526020016101f1565b34801561026057600080fd5b5061027461026f3660046130ac565b6107bb565b005b34801561028257600080fd5b50610274610291366004613119565b6108d5565b3480156102a257600080fd5b506102746102b1366004613192565b61093c565b3480156102c257600080fd5b5060d3546001600160a01b031661023c565b3480156102e057600080fd5b506102746102ef366004613227565b6109bd565b34801561030057600080fd5b5061031461030f36600461307e565b6109ee565b6040519081526020016101f1565b34801561032e57600080fd5b5061027461033d366004613268565b610a03565b34801561034e57600080fd5b5061027461035d366004613268565b610a1f565b610274610370366004613298565b610a9d565b34801561038157600080fd5b50610274610390366004613227565b610b7b565b3480156103a157600080fd5b506103146103b036600461307e565b60009081526097602052604090205490565b3480156103ce57600080fd5b5061031460008051602061418b83398151915281565b3480156103f057600080fd5b506101e56103ff36600461339d565b610b96565b34801561041057600080fd5b506101e561041f36600461307e565b600090815260d9602052604090205460ff1690565b34801561044057600080fd5b5061027461044f3660046134f5565b610cd5565b34801561046057600080fd5b5061023c61046f36600461307e565b610dad565b34801561048057600080fd5b5061031461048f36600461351a565b610de1565b3480156104a057600080fd5b506104a9610e67565b6040516101f19190613537565b3480156104c257600080fd5b506102746104d136600461351a565b610ec8565b3480156104e257600080fd5b506101e56104f1366004613268565b610ef2565b34801561050257600080fd5b5061020f610f1d565b34801561051757600080fd5b50610520610f2c565b6040516101f19190613584565b34801561053957600080fd5b50610314600081565b34801561054e57600080fd5b5061027461055d3660046135bc565b610f83565b34801561056e57600080fd5b5061027461057d36600461351a565b610f8e565b34801561058e57600080fd5b5061027461059d36600461351a565b6110d2565b3480156105ae57600080fd5b506102746105bd36600461351a565b6110f5565b3480156105ce57600080fd5b506102746105dd3660046135ea565b611160565b3480156105ee57600080fd5b506102746105fd3660046136ad565b611198565b34801561060e57600080fd5b5061020f61061d36600461307e565b61139e565b34801561062e57600080fd5b5061027461063d366004613268565b611412565b34801561064e57600080fd5b5061066261065d36600461351a565b61142e565b6040516101f19190613718565b34801561067b57600080fd5b506103146000805160206141ab83398151915281565b34801561069d57600080fd5b506102746106ac36600461376e565b6115bb565b3480156106bd57600080fd5b506101e56106cc3660046137cd565b611695565b3480156106dd57600080fd5b506102746106ec3660046137fb565b6116c3565b60006106fc8261170b565b92915050565b6060606580546107119061381d565b80601f016020809104026020016040519081016040528092919081815260200182805461073d9061381d565b801561078a5780601f1061075f5761010080835404028352916020019161078a565b820191906000526020600020905b81548152906001019060200180831161076d57829003601f168201915b5050505050905090565b600061079f82611730565b506000908152606960205260409020546001600160a01b031690565b60006107c682610dad565b9050806001600160a01b0316836001600160a01b0316036108385760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b60648201526084015b60405180910390fd5b336001600160a01b038216148061085457506108548133611695565b6108c65760405162461bcd60e51b815260206004820152603d60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206f7220617070726f76656420666f7220616c6c000000606482015260840161082f565b6108d08383611755565b505050565b60d085905560d16108e78486836138b4565b5060d26108f58284836138b4565b507f83eda6edb918aac3f35c8d616a6198dbb75eeca3248e0efec05e66a30d51119160d05460d160d260405161092d939291906139ea565b60405180910390a15050505050565b6001600160a01b038716600090815260d7602052604090206001016109628688836138b4565b506001600160a01b038716600090815260d760205260409020600281018590556003016109908284836138b4565b506001600160a01b038716600090815260d760205260409020546109b490846116c3565b50505050505050565b6109c733826117c3565b6109e35760405162461bcd60e51b815260040161082f90613a15565b6108d0838383611822565b60009081526099602052604090206001015490565b610a0c826109ee565b610a1581611986565b6108d08383611990565b6001600160a01b0381163314610a8f5760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b606482015260840161082f565b610a998282611a16565b5050565b610aa78233611a7d565b8060d5610ab48280613a62565b604051610ac2929190613aa8565b908152604051908190036020019020610adb8282613af8565b5082905060d4610aeb8280613a62565b604051610af9929190613aa8565b908152604051908190036020019020610b128282613bf5565b507fe8e692af5f416c4637c9307d7f08b7697fd8f4398f55103ce03ddd9388ccd84d9050610b4660a084016080850161351a565b610b508380613a62565b610b5e610140860186613a62565b604051610b6f959493929190613d2d565b60405180910390a15050565b6108d083838360405180602001604052806000815250611160565b6000866001600160a01b0316610bab89610dad565b6001600160a01b031614610bfc5760405162461bcd60e51b81526020600482015260186024820152771b999d081bdddb995c881a5cc81b9bdd0818dbdc9c9958dd60421b604482015260640161082f565b845180610c445760405162461bcd60e51b81526020600482015260166024820152757573657249442063616e6e6f7420626520656d70747960501b604482015260640161082f565b60005b81811015610c8357610c71878281518110610c6457610c64613d71565b6020026020010151611b12565b80610c7b81613d9d565b915050610c47565b506000610cc7888787878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611d7e92505050565b9a9950505050505050505050565b610cdd611ddf565b610ce682611e4f565b610d025760405162461bcd60e51b815260040161082f90613db6565b600082815260da602090815260408220805460018101825590835291200180546001600160a01b03191633179055610d3982611ea5565b8015610d425750805b15610d6b57610d5082611ed1565b600082815260d960205260409020805460ff19168215151790555b6040805133815260208101849052821515918101919091527f937b4d1b5b461852fdbbb174b68adefb42155e668c05a1261fae11a2d22a240a90606001610b6f565b600080610db983611f9c565b90506001600160a01b0381166106fc5760405162461bcd60e51b815260040161082f90613de8565b60006001600160a01b038216610e4b5760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b606482015260840161082f565b506001600160a01b031660009081526068602052604090205490565b606060dd80548060200260200160405190810160405280929190818152602001828054801561078a57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610ea1575050505050905090565b610ed0611fb7565b60d380546001600160a01b0319166001600160a01b0392909216919091179055565b60009182526099602090815260408084206001600160a01b0393909316845291905290205460ff1690565b6060606680546107119061381d565b606060dc80548060200260200160405190810160405280929190818152602001828054801561078a57602002820191906000526020600020905b815481526020019060010190808311610f66575050505050905090565b610a9933838361200e565b600054610100900460ff1615808015610fae5750600054600160ff909116105b80610fcf5750610fbd306120d8565b158015610fcf575060005460ff166001145b6110325760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161082f565b6000805460ff191660011790558015611055576000805461ff0019166101001790555b60cc8054336001600160a01b0319918216811790925560cb80549091166001600160a01b038516179055600160cd55611090906000906120e7565b8015610a99576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602001610b6f565b6110da611fb7565b6110f260008051602061418b833981519152826120e7565b50565b6110fd611fb7565b60dd80546001810182556000919091527fac507b9f8bf86ad8bb770f71cd2b1992902ae0314d93fc0f2bb011d70e7962260180546001600160a01b0319166001600160a01b0383161790556110f26000805160206141ab833981519152826120e7565b61116a33836117c3565b6111865760405162461bcd60e51b815260040161082f90613a15565b611192848484846120f1565b50505050565b6111a0611ddf565b8260005b81811015611366576111cd8686838181106111c1576111c1613d71565b90506020020135611e4f565b6111e95760405162461bcd60e51b815260040161082f90613db6565b8383828181106111fb576111fb613d71565b90506020020160208101906112109190613e1a565b60d9600088888581811061122657611226613d71565b90506020020135815260200190815260200160002060006101000a81548160ff02191690831515021790555060da600087878481811061126857611268613d71565b6020908102929092013583525081810192909252604001600090812080546001810182559082529190200180546001600160a01b031916331790556112c48686838181106112b8576112b8613d71565b90506020020135611ea5565b15611354576112ea8686838181106112de576112de613d71565b90506020020135611ed1565b8383828181106112fc576112fc613d71565b90506020020160208101906113119190613e1a565b60d9600088888581811061132757611327613d71565b90506020020135815260200190815260200160002060006101000a81548160ff0219169083151502179055505b8061135e81613d9d565b9150506111a4565b507f04798cc2c828443d2d9ef127deaa1f73e782d8a9fffbc9db2bcca53e12765af5338686868660405161092d959493929190613e37565b60606113a982611730565b60006113c060408051602081019091526000815290565b905060008151116113e0576040518060200160405280600081525061140b565b806113ea84612124565b6040516020016113fb929190613ebb565b6040516020818303038152906040525b9392505050565b61141b826109ee565b61142481611986565b6108d08383611a16565b6114596040518060800160405280600081526020016060815260200160008152602001606081525090565b6001600160a01b038216600090815260d760209081526040918290208251608081019093528054835260018101805491928401916114969061381d565b80601f01602080910402602001604051908101604052809291908181526020018280546114c29061381d565b801561150f5780601f106114e45761010080835404028352916020019161150f565b820191906000526020600020905b8154815290600101906020018083116114f257829003601f168201915b50505050508152602001600282015481526020016003820180546115329061381d565b80601f016020809104026020016040519081016040528092919081815260200182805461155e9061381d565b80156115ab5780601f10611580576101008083540402835291602001916115ab565b820191906000526020600020905b81548152906001019060200180831161158e57829003601f168201915b5050505050815250509050919050565b6115c36121b6565b600082826040516115d5929190613aa8565b60405190819003812060dc805460018101825560009182527f3162b0988d4210bff484413ed451d170a03887272177efc0b7d000f10abe9edf018290559092509060d4906116269088908890613aa8565b90815260408051918290036020908101832060020154600086815260db90925291902081905591507f99a038e9d345d0b12130b3b1fb003bf8f2d3a5c27ce2a800bbb1608efff6c591906116859033908990899087908a908a90613eea565b60405180910390a1505050505050565b6001600160a01b039182166000908152606a6020908152604080832093909416825291909152205460ff1690565b60008281526097602090815260409182902083905581518481529081018390527f21b5a040ca6f61b8bcacb8f423e25ce46e88932f887ce92e60343369c20ec06f9101610b6f565b60006001600160e01b03198216637965db0b60e01b14806106fc57506106fc82612219565b6117398161223e565b6110f25760405162461bcd60e51b815260040161082f90613de8565b600081815260696020526040902080546001600160a01b0319166001600160a01b038416908117909155819061178a82610dad565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6000806117cf83610dad565b9050806001600160a01b0316846001600160a01b031614806117f657506117f68185611695565b8061181a5750836001600160a01b031661180f84610794565b6001600160a01b0316145b949350505050565b826001600160a01b031661183582610dad565b6001600160a01b03161461185b5760405162461bcd60e51b815260040161082f90613f35565b6001600160a01b0382166118bd5760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b606482015260840161082f565b826001600160a01b03166118d082610dad565b6001600160a01b0316146118f65760405162461bcd60e51b815260040161082f90613f35565b600081815260696020908152604080832080546001600160a01b03199081169091556001600160a01b0387811680865260688552838620805460001901905590871680865283862080546001019055868652606790945282852080549092168417909155905184937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6110f2813361225b565b61199a8282610ef2565b610a995760008281526099602090815260408083206001600160a01b03851684529091529020805460ff191660011790556119d23390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b611a208282610ef2565b15610a995760008281526099602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b6000611a8f60a084016080850161351a565b60cc54909150611ab090829084906001600160a01b031660408701356122b4565b7f704db66c978a7472842ddae0d478a8a7d123fb6ad449e4426b5e0f6a22081de7611ae160a085016080860161351a565b60cc5460408051611b05939287926001600160a01b03909116919089013590613f7a565b60405180910390a1505050565b6000815111611b5d5760405162461bcd60e51b81526020600482015260176024820152767573657249442063616e206e6f7420626520656d70747960481b604482015260640161082f565b6040805180820190915260018152601d60f91b602082015260008080611b828561230e565b9050611b97611b908561230e565b8290612319565b9194509250905082611bfb5760405162461bcd60e51b815260206004820152602760248201527f74686520666972737420706172742064656c696d6974657220646f6573206e6f6044820152661d08195e1a5cdd60ca1b606482015260840161082f565b6000611c0683612368565b5111611c545760405162461bcd60e51b815260206004820152601d60248201527f746865206669727374207061727420646f6573206e6f74206578697374000000604482015260640161082f565b611c60611b908561230e565b9194509250905082611cc55760405162461bcd60e51b815260206004820152602860248201527f746865207365636f6e6420706172742064656c696d6974657220646f6573206e6044820152671bdd08195e1a5cdd60c21b606482015260840161082f565b6000611cd083612368565b5111611d1e5760405162461bcd60e51b815260206004820152601e60248201527f746865207365636f6e64207061727420646f6573206e6f742065786973740000604482015260640161082f565b611d2781612368565b51604014611d775760405162461bcd60e51b815260206004820152601d60248201527f69642068617368206c656e677468206973206e6f7420636f7272656374000000604482015260640161082f565b5050505050565b6000806000611d8d8585612373565b90925090506000816004811115611da657611da6613fa4565b148015611dc45750856001600160a01b0316826001600160a01b0316145b80611dd55750611dd58686866123b8565b9695505050505050565b611df76000805160206141ab83398151915233610ef2565b611e4d5760405162461bcd60e51b815260206004820152602160248201527f73656e64657220646f65736e2774206861766520766572696669657220726f6c6044820152606560f81b606482015260840161082f565b565b6000805b60dc54811015611e9c578260dc8281548110611e7157611e71613d71565b906000526020600020015403611e8a5750600192915050565b80611e9481613d9d565b915050611e53565b50600092915050565b600081815260da602052604081205460dd54611ec2826002613fba565b1115611e9c5750600192915050565b60d35460408051808201909152601b81527a3c3c361039b7b6103830bcafb83937b334bc101d1d1d101717171760291b60208201526001600160a01b0390911690611f1b906124a4565b60cc54600083815260db6020526040902054611f469183916001600160a01b039091169033906122b4565b60d35460cc54600084815260db6020526040908190205490517f3832ba0f9bc9de56bd1f974e665b53f2643ed48a47b45c3d36976bc1ea31459993610b6f936001600160a01b0391821693911691339190613f7a565b6000908152606760205260409020546001600160a01b031690565b611fc2600033610ef2565b611e4d5760405162461bcd60e51b815260206004820152601e60248201527f73656e64657220646f65736e277420686176652061646d696e20726f6c650000604482015260640161082f565b816001600160a01b0316836001600160a01b03160361206b5760405162461bcd60e51b815260206004820152601960248201527822a9219b99189d1030b8383937bb32903a379031b0b63632b960391b604482015260640161082f565b6001600160a01b038381166000818152606a6020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6001600160a01b03163b151590565b610a998282611990565b6120fc848484611822565b612108848484846124e7565b6111925760405162461bcd60e51b815260040161082f90613fd1565b60606000612131836125ef565b60010190506000816001600160401b0381111561215057612150613300565b6040519080825280601f01601f19166020018201604052801561217a576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a850494508461218457509392505050565b6121ce60008051602061418b83398151915233610ef2565b611e4d5760405162461bcd60e51b815260206004820152601c60248201527b73656e64657220646f65736e277420686176652074656520726f6c6560201b604482015260640161082f565b60006001600160e01b0319821663f389baad60e01b14806106fc57506106fc826126c5565b60008061224a83611f9c565b6001600160a01b0316141592915050565b6122658282610ef2565b610a995761227281612715565b61227d836020612727565b60405160200161228e929190614023565b60408051601f198184030181529082905262461bcd60e51b825261082f9160040161306b565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b1790526111929085906128c2565b60006106fc82612a14565b60008080806123288686612a28565b9050600019810361234457600086600093509350935050612361565b6123598161235187612b3f565b889190612b4b565b935093509350505b9250925092565b60606106fc82612bba565b60008082516041036123a95760208301516040840151606085015160001a61239d87828585612c2f565b945094505050506123b1565b506000905060025b9250929050565b6000806000856001600160a01b0316631626ba7e60e01b86866040516024016123e2929190614092565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b031990941693909317909252905161242091906140ab565b600060405180830381855afa9150503d806000811461245b576040519150601f19603f3d011682016040523d82523d6000602084013e612460565b606091505b509150915081801561247457506020815110155b8015611dd557508051630b135d3f60e11b9061249990830160209081019084016140c7565b149695505050505050565b6110f2816040516024016124b8919061306b565b60408051601f198184030181529190526020810180516001600160e01b031663104c13eb60e21b179052612ce9565b60006124fb846001600160a01b03166120d8565b156125e457604051630a85bd0160e11b81526001600160a01b0385169063150b7a02906125329033908990889088906004016140e0565b6020604051808303816000875af192505050801561256d575060408051601f3d908101601f1916820190925261256a91810190614113565b60015b6125ca573d80801561259b576040519150601f19603f3d011682016040523d82523d6000602084013e6125a0565b606091505b5080516000036125c25760405162461bcd60e51b815260040161082f90613fd1565b805181602001fd5b6001600160e01b031916630a85bd0160e11b14905061181a565b506001949350505050565b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b831061262e5772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6904ee2d6d415b85acef8160201b8310612658576904ee2d6d415b85acef8160201b830492506020015b662386f26fc10000831061267657662386f26fc10000830492506010015b6305f5e100831061268e576305f5e100830492506008015b61271083106126a257612710830492506004015b606483106126b4576064830492506002015b600a83106106fc5760010192915050565b60006001600160e01b031982166380ac58cd60e01b14806126f657506001600160e01b03198216635b5e139f60e01b145b806106fc57506301ffc9a760e01b6001600160e01b03198316146106fc565b60606106fc6001600160a01b03831660145b60606000612736836002613fba565b612741906002614130565b6001600160401b0381111561275857612758613300565b6040519080825280601f01601f191660200182016040528015612782576020820181803683370190505b509050600360fc1b8160008151811061279d5761279d613d71565b60200101906001600160f81b031916908160001a905350600f60fb1b816001815181106127cc576127cc613d71565b60200101906001600160f81b031916908160001a90535060006127f0846002613fba565b6127fb906001614130565b90505b6001811115612873576f181899199a1a9b1b9c1cb0b131b232b360811b85600f166010811061282f5761282f613d71565b1a60f81b82828151811061284557612845613d71565b60200101906001600160f81b031916908160001a90535060049490941c9361286c81614143565b90506127fe565b50831561140b5760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e74604482015260640161082f565b813b806129095760405162461bcd60e51b8152602060048201526015602482015274115490cc8c0e881b9bdd08184818dbdb9d1c9858dd605a1b604482015260640161082f565b600080846001600160a01b03168460405161292491906140ab565b6000604051808303816000865af19150503d8060008114612961576040519150601f19603f3d011682016040523d82523d6000602084013e612966565b606091505b5091509150816129ad5760405162461bcd60e51b8152602060048201526012602482015271115490cc8c0e8818d85b1b0819985a5b195960721b604482015260640161082f565b805115611d7757808060200190518101906129c8919061415a565b611d775760405162461bcd60e51b815260206004820181905260248201527f45524332303a206f7065726174696f6e20646964206e6f742073756363656564604482015260640161082f565b60008060208301905061140b818451612cf2565b600080612a3484612b3f565b90506000612a4184612b3f565b905080600003612a56576000925050506106fc565b811580612a6257508181115b15612a7357600019925050506106fc565b6000612a7f8660801c90565b90506000612a8d8660801c90565b90506000612a9c825160001a90565b90505b6000612aac848784612cfe565b90506000198103612ac75760001996505050505050506106fc565b94859003949283019285851115612ae85760001996505050505050506106fc565b84832085852003612b1457612afd8960801c90565b612b079085614177565b96505050505050506106fc565b85600103612b2c5760001996505050505050506106fc565b6000199095019460019093019250612a9f565b6001600160801b031690565b600080600080612b5b8760801c90565b90506000612b6888612b3f565b905086860181811115612b8e576040516365f4e9df60e01b815260040160405180910390fd5b612b988389612dbe565b9450612ba8818401828403612dbe565b93506001955050505093509350939050565b6060612bc582612b3f565b6001600160401b03811115612bdc57612bdc613300565b6040519080825280601f01601f191660200182016040528015612c06576020820181803683370190505b50905060208101612c2981612c1b8560801c90565b612c2486612b3f565b612dd1565b50919050565b6000806fa2a8918ca85bafe22016d0b997e4df60600160ff1b03831115612c5c5750600090506003612ce0565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015612cb0573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116612cd957600060019250925050612ce0565b9150600090505b94509492505050565b6110f281612ddf565b600061140b8383612e00565b600060208311612d1a57612d13848484612e0c565b905061140b565b83601f841680850382016000612d2f86612f0f565b90505b81881015612d745787518118612d4781612f37565b15612d6857848903612d5b8a60208a612e0c565b019550505050505061140b565b60208901985050612d32565b82600003612d8a5760001994505050505061140b565b612d95828488612e0c565b94506000198503612dae5760001994505050505061140b565b838203850194505050505061140b565b6001600160801b031660809190911b1790565b8083828460045afa50505050565b60006a636f6e736f6c652e6c6f679050600080835160208501845afa505050565b600061140b8383612f82565b8251600090816020851115612e2057602094505b60128510612ea3576000612e3385612f0f565b83189050612e496001600160801b038217612f37565b600003612e7f5760109150601a8610612e7a57612e6e6001600160401b038217612f37565b600003612e7a57601891505b612e9d565b612e916001600160c01b038217612f37565b600003612e9d57600891505b50612eda565b600a8510612eda576000612eb685612f0f565b83189050612ecc6001600160c01b038217612f37565b600003612ed857600891505b505b84811015612f025781811a60ff85168103612ef95750915061140b9050565b50600101612eda565b5060001995945050505050565b60ff167f01010101010101010101010101010101010101010101010101010101010101010290565b7ffefefefefefefefefefefefefefefefefefefefefefefefefefefefefefefeff81019019167f80808080808080808080808080808080808080808080808080808080808080801690565b60006001600160801b03831115612fac5760405163fee7506f60e01b815260040160405180910390fd5b6001600160801b03821115612fd457604051633b6b098d60e01b815260040160405180910390fd5b506001600160801b031660809190911b1790565b6001600160e01b0319811681146110f257600080fd5b60006020828403121561301057600080fd5b813561140b81612fe8565b60005b8381101561303657818101518382015260200161301e565b50506000910152565b6000815180845261305781602086016020860161301b565b601f01601f19169290920160200192915050565b60208152600061140b602083018461303f565b60006020828403121561309057600080fd5b5035919050565b6001600160a01b03811681146110f257600080fd5b600080604083850312156130bf57600080fd5b82356130ca81613097565b946020939093013593505050565b60008083601f8401126130ea57600080fd5b5081356001600160401b0381111561310157600080fd5b6020830191508360208285010111156123b157600080fd5b60008060008060006060868803121561313157600080fd5b8535945060208601356001600160401b038082111561314f57600080fd5b61315b89838a016130d8565b9096509450604088013591508082111561317457600080fd5b50613181888289016130d8565b969995985093965092949392505050565b600080600080600080600060a0888a0312156131ad57600080fd5b87356131b881613097565b965060208801356001600160401b03808211156131d457600080fd5b6131e08b838c016130d8565b909850965060408a0135955060608a0135945060808a013591508082111561320757600080fd5b506132148a828b016130d8565b989b979a50959850939692959293505050565b60008060006060848603121561323c57600080fd5b833561324781613097565b9250602084013561325781613097565b929592945050506040919091013590565b6000806040838503121561327b57600080fd5b82359150602083013561328d81613097565b809150509250929050565b600080604083850312156132ab57600080fd5b82356001600160401b03808211156132c257600080fd5b9084019060c082870312156132d657600080fd5b909250602084013590808211156132ec57600080fd5b508301610160818603121561328d57600080fd5b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b038111828210171561333e5761333e613300565b604052919050565b60006001600160401b0383111561335f5761335f613300565b613372601f8401601f1916602001613316565b905082815283838301111561338657600080fd5b828260208301376000602084830101529392505050565b600080600080600080600060c0888a0312156133b857600080fd5b873596506133c96020890135613097565b602088013595506133dd6040890135613097565b604088013594506001600160401b0360608901358110156133fd57600080fd5b606089013589018a601f82011261341357600080fd5b818135111561342457613424613300565b6134346020823560051b01613316565b81358082526020808301929160051b8401018d81111561345357600080fd5b602084015b818110156134a457858135111561346e57600080fd5b8e603f82358701011261348057600080fd5b6134968f82358701602081013590604001613346565b845260209384019301613458565b5050809750505050608089013593508060a08a013511156134c457600080fd5b506134d58960a08a01358a016130d8565b979a9699509497509295919491925050565b80151581146110f257600080fd5b6000806040838503121561350857600080fd5b82359150602083013561328d816134e7565b60006020828403121561352c57600080fd5b813561140b81613097565b6020808252825182820181905260009190848201906040850190845b818110156135785783516001600160a01b031683529284019291840191600101613553565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b81811015613578578351835292840192918401916001016135a0565b600080604083850312156135cf57600080fd5b82356135da81613097565b9150602083013561328d816134e7565b6000806000806080858703121561360057600080fd5b843561360b81613097565b9350602085013561361b81613097565b92506040850135915060608501356001600160401b0381111561363d57600080fd5b8501601f8101871361364e57600080fd5b61365d87823560208401613346565b91505092959194509250565b60008083601f84011261367b57600080fd5b5081356001600160401b0381111561369257600080fd5b6020830191508360208260051b85010111156123b157600080fd5b600080600080604085870312156136c357600080fd5b84356001600160401b03808211156136da57600080fd5b6136e688838901613669565b909650945060208701359150808211156136ff57600080fd5b5061370c87828801613669565b95989497509550505050565b6020815281516020820152600060208301516080604084015261373e60a084018261303f565b9050604084015160608401526060840151601f19848303016080850152613765828261303f565b95945050505050565b6000806000806040858703121561378457600080fd5b84356001600160401b038082111561379b57600080fd5b6137a7888389016130d8565b909650945060208701359150808211156137c057600080fd5b5061370c878288016130d8565b600080604083850312156137e057600080fd5b82356137eb81613097565b9150602083013561328d81613097565b6000806040838503121561380e57600080fd5b50508035926020909101359150565b600181811c9082168061383157607f821691505b602082108103612c2957634e487b7160e01b600052602260045260246000fd5b601f8211156108d057600081815260208120601f850160051c810160208610156138785750805b601f850160051c820191505b8181101561389757828155600101613884565b505050505050565b600019600383901b1c191660019190911b1790565b6001600160401b038311156138cb576138cb613300565b6138df836138d9835461381d565b83613851565b6000601f84116001811461390d57600085156138fb5750838201355b613905868261389f565b845550611d77565b600083815260209020601f19861690835b8281101561393e578685013582556020948501946001909201910161391e565b508682101561395b5760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b6000815461397a8161381d565b80855260206001838116801561399757600181146139b1576139df565b60ff1985168884015283151560051b8801830195506139df565b866000528260002060005b858110156139d75781548a82018601529083019084016139bc565b890184019650505b505050505092915050565b838152606060208201526000613a03606083018561396d565b8281036040840152611dd5818561396d565b6020808252602d908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526c1c881bdc88185c1c1c9bdd9959609a1b606082015260800190565b6000808335601e19843603018112613a7957600080fd5b8301803591506001600160401b03821115613a9357600080fd5b6020019150368190038213156123b157600080fd5b8183823760009101908152919050565b600081356106fc81613097565b80546001600160a01b0319166001600160a01b0392909216919091179055565b6000813560ff811681146106fc57600080fd5b613b028283613a62565b613b0d8183856138b4565b5050613b1c6020830183613a62565b613b2a8183600186016138b4565b505060028101613b45613b3f60408501613ab8565b82613ac5565b613b72613b5460608501613ae5565b82805460ff60a01b191660a09290921b60ff60a01b16919091179055565b50613b8b613b8260808401613ab8565b60038301613ac5565b60a0820135600482015560c08201356005820155613bc3613bae60e08401613ae5565b6006830160ff821660ff198254161781555050565b61010082013560078201556101208201356008820155613be7610140830183613a62565b6111928183600986016138b4565b613bff8283613a62565b6001600160401b03811115613c1657613c16613300565b613c2a81613c24855461381d565b85613851565b6000601f821160018114613c585760008315613c465750838201355b613c50848261389f565b865550613cb2565b600085815260209020601f19841690835b82811015613c895786850135825560209485019460019092019101613c69565b5084821015613ca65760001960f88660031b161c19848701351681555b505060018360011b0185555b50505050613cce613cc560208401613ab8565b60018301613ac5565b604082013560028201556060820135600382015560048101613cf5613b3f60808501613ab8565b6108d0613b5460a08501613ae5565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b0386168152606060208201819052600090613d529083018688613d04565b8281036040840152613d65818587613d04565b98975050505050505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201613daf57613daf613d87565b5060010190565b602080825260189082015277185d1d195cdd185d1a5bdb881a5cc81b9bdd08195e1a5cdd60421b604082015260600190565b602080825260189082015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b604082015260600190565b600060208284031215613e2c57600080fd5b813561140b816134e7565b6001600160a01b0386168152606060208083018290529082018590526000906001600160fb1b03861115613e6a57600080fd5b8560051b808860808601378301838103608090810160408601528101859052859060009060a0015b86821015610cc7578235613ea5816134e7565b1515815291830191600191909101908301613e92565b60008351613ecd81846020880161301b565b835190830190613ee181836020880161301b565b01949350505050565b6001600160a01b0387168152608060208201819052600090613f0f9083018789613d04565b8560408401528281036060840152613f28818587613d04565b9998505050505050505050565b60208082526025908201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060408201526437bbb732b960d91b606082015260800190565b6001600160a01b039485168152928416602084015292166040820152606081019190915260800190565b634e487b7160e01b600052602160045260246000fd5b80820281158282048414176106fc576106fc613d87565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b76020b1b1b2b9b9a1b7b73a3937b61d1030b1b1b7bab73a1604d1b81526000835161405581601785016020880161301b565b7001034b99036b4b9b9b4b733903937b6329607d1b601791840191820152835161408681602884016020880161301b565b01602801949350505050565b82815260406020820152600061181a604083018461303f565b600082516140bd81846020870161301b565b9190910192915050565b6000602082840312156140d957600080fd5b5051919050565b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090611dd59083018461303f565b60006020828403121561412557600080fd5b815161140b81612fe8565b808201808211156106fc576106fc613d87565b60008161415257614152613d87565b506000190190565b60006020828403121561416c57600080fd5b815161140b816134e7565b818103818111156106fc576106fc613d8756fed16a327e6f5c32c69c8282ab355bc8a366432cf60ee1165bc5198414ca1b06c70ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea09a2646970667358221220a3bad617b92cc38b98507db9778d316a40ab70849f937dc1528fbd524596614c64736f6c63430008110033",
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

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Contract *ContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Contract *ContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Contract.Contract.DEFAULTADMINROLE(&_Contract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Contract *ContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Contract.Contract.DEFAULTADMINROLE(&_Contract.CallOpts)
}

// TEEROLE is a free data retrieval call binding the contract method 0x495eb828.
//
// Solidity: function TEE_ROLE() view returns(bytes32)
func (_Contract *ContractCaller) TEEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "TEE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TEEROLE is a free data retrieval call binding the contract method 0x495eb828.
//
// Solidity: function TEE_ROLE() view returns(bytes32)
func (_Contract *ContractSession) TEEROLE() ([32]byte, error) {
	return _Contract.Contract.TEEROLE(&_Contract.CallOpts)
}

// TEEROLE is a free data retrieval call binding the contract method 0x495eb828.
//
// Solidity: function TEE_ROLE() view returns(bytes32)
func (_Contract *ContractCallerSession) TEEROLE() ([32]byte, error) {
	return _Contract.Contract.TEEROLE(&_Contract.CallOpts)
}

// VERIFIERROLE is a free data retrieval call binding the contract method 0xe7705db6.
//
// Solidity: function VERIFIER_ROLE() view returns(bytes32)
func (_Contract *ContractCaller) VERIFIERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "VERIFIER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VERIFIERROLE is a free data retrieval call binding the contract method 0xe7705db6.
//
// Solidity: function VERIFIER_ROLE() view returns(bytes32)
func (_Contract *ContractSession) VERIFIERROLE() ([32]byte, error) {
	return _Contract.Contract.VERIFIERROLE(&_Contract.CallOpts)
}

// VERIFIERROLE is a free data retrieval call binding the contract method 0xe7705db6.
//
// Solidity: function VERIFIER_ROLE() view returns(bytes32)
func (_Contract *ContractCallerSession) VERIFIERROLE() ([32]byte, error) {
	return _Contract.Contract.VERIFIERROLE(&_Contract.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Contract *ContractCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Contract *ContractSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Contract *ContractCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Contract *ContractCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Contract *ContractSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetApproved(&_Contract.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Contract *ContractCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetApproved(&_Contract.CallOpts, tokenId)
}

// GetIdentitiesRoot is a free data retrieval call binding the contract method 0x45692b82.
//
// Solidity: function getIdentitiesRoot(uint256 id) view returns(bytes32)
func (_Contract *ContractCaller) GetIdentitiesRoot(opts *bind.CallOpts, id *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getIdentitiesRoot", id)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetIdentitiesRoot is a free data retrieval call binding the contract method 0x45692b82.
//
// Solidity: function getIdentitiesRoot(uint256 id) view returns(bytes32)
func (_Contract *ContractSession) GetIdentitiesRoot(id *big.Int) ([32]byte, error) {
	return _Contract.Contract.GetIdentitiesRoot(&_Contract.CallOpts, id)
}

// GetIdentitiesRoot is a free data retrieval call binding the contract method 0x45692b82.
//
// Solidity: function getIdentitiesRoot(uint256 id) view returns(bytes32)
func (_Contract *ContractCallerSession) GetIdentitiesRoot(id *big.Int) ([32]byte, error) {
	return _Contract.Contract.GetIdentitiesRoot(&_Contract.CallOpts, id)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Contract *ContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Contract *ContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Contract.Contract.GetRoleAdmin(&_Contract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Contract *ContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Contract.Contract.GetRoleAdmin(&_Contract.CallOpts, role)
}

// GetAttestationResult is a free data retrieval call binding the contract method 0x49e9e834.
//
// Solidity: function get_attestation_result(bytes32 attestation_id) view returns(bool)
func (_Contract *ContractCaller) GetAttestationResult(opts *bind.CallOpts, attestation_id [32]byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "get_attestation_result", attestation_id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetAttestationResult is a free data retrieval call binding the contract method 0x49e9e834.
//
// Solidity: function get_attestation_result(bytes32 attestation_id) view returns(bool)
func (_Contract *ContractSession) GetAttestationResult(attestation_id [32]byte) (bool, error) {
	return _Contract.Contract.GetAttestationResult(&_Contract.CallOpts, attestation_id)
}

// GetAttestationResult is a free data retrieval call binding the contract method 0x49e9e834.
//
// Solidity: function get_attestation_result(bytes32 attestation_id) view returns(bool)
func (_Contract *ContractCallerSession) GetAttestationResult(attestation_id [32]byte) (bool, error) {
	return _Contract.Contract.GetAttestationResult(&_Contract.CallOpts, attestation_id)
}

// GetPayAddress is a free data retrieval call binding the contract method 0x1d17e1bd.
//
// Solidity: function get_pay_address() view returns(address)
func (_Contract *ContractCaller) GetPayAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "get_pay_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPayAddress is a free data retrieval call binding the contract method 0x1d17e1bd.
//
// Solidity: function get_pay_address() view returns(address)
func (_Contract *ContractSession) GetPayAddress() (common.Address, error) {
	return _Contract.Contract.GetPayAddress(&_Contract.CallOpts)
}

// GetPayAddress is a free data retrieval call binding the contract method 0x1d17e1bd.
//
// Solidity: function get_pay_address() view returns(address)
func (_Contract *ContractCallerSession) GetPayAddress() (common.Address, error) {
	return _Contract.Contract.GetPayAddress(&_Contract.CallOpts)
}

// GetProofList is a free data retrieval call binding the contract method 0x964b1f6e.
//
// Solidity: function get_proof_list() view returns(bytes32[])
func (_Contract *ContractCaller) GetProofList(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "get_proof_list")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetProofList is a free data retrieval call binding the contract method 0x964b1f6e.
//
// Solidity: function get_proof_list() view returns(bytes32[])
func (_Contract *ContractSession) GetProofList() ([][32]byte, error) {
	return _Contract.Contract.GetProofList(&_Contract.CallOpts)
}

// GetProofList is a free data retrieval call binding the contract method 0x964b1f6e.
//
// Solidity: function get_proof_list() view returns(bytes32[])
func (_Contract *ContractCallerSession) GetProofList() ([][32]byte, error) {
	return _Contract.Contract.GetProofList(&_Contract.CallOpts)
}

// GetUserByAddress is a free data retrieval call binding the contract method 0xda7a2928.
//
// Solidity: function get_user_by_address(address user_address) view returns((uint256,string,uint256,bytes))
func (_Contract *ContractCaller) GetUserByAddress(opts *bind.CallOpts, user_address common.Address) (CarvProtocolServiceuser, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "get_user_by_address", user_address)

	if err != nil {
		return *new(CarvProtocolServiceuser), err
	}

	out0 := *abi.ConvertType(out[0], new(CarvProtocolServiceuser)).(*CarvProtocolServiceuser)

	return out0, err

}

// GetUserByAddress is a free data retrieval call binding the contract method 0xda7a2928.
//
// Solidity: function get_user_by_address(address user_address) view returns((uint256,string,uint256,bytes))
func (_Contract *ContractSession) GetUserByAddress(user_address common.Address) (CarvProtocolServiceuser, error) {
	return _Contract.Contract.GetUserByAddress(&_Contract.CallOpts, user_address)
}

// GetUserByAddress is a free data retrieval call binding the contract method 0xda7a2928.
//
// Solidity: function get_user_by_address(address user_address) view returns((uint256,string,uint256,bytes))
func (_Contract *ContractCallerSession) GetUserByAddress(user_address common.Address) (CarvProtocolServiceuser, error) {
	return _Contract.Contract.GetUserByAddress(&_Contract.CallOpts, user_address)
}

// GetVerifierList is a free data retrieval call binding the contract method 0x70b3542d.
//
// Solidity: function get_verifier_list() view returns(address[])
func (_Contract *ContractCaller) GetVerifierList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "get_verifier_list")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetVerifierList is a free data retrieval call binding the contract method 0x70b3542d.
//
// Solidity: function get_verifier_list() view returns(address[])
func (_Contract *ContractSession) GetVerifierList() ([]common.Address, error) {
	return _Contract.Contract.GetVerifierList(&_Contract.CallOpts)
}

// GetVerifierList is a free data retrieval call binding the contract method 0x70b3542d.
//
// Solidity: function get_verifier_list() view returns(address[])
func (_Contract *ContractCallerSession) GetVerifierList() ([]common.Address, error) {
	return _Contract.Contract.GetVerifierList(&_Contract.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Contract *ContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Contract *ContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Contract.Contract.HasRole(&_Contract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Contract *ContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Contract.Contract.HasRole(&_Contract.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Contract *ContractCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Contract *ContractSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Contract.Contract.IsApprovedForAll(&_Contract.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Contract *ContractCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Contract.Contract.IsApprovedForAll(&_Contract.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCallerSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Contract *ContractCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Contract *ContractSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.OwnerOf(&_Contract.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Contract *ContractCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.OwnerOf(&_Contract.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contract *ContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contract *ContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contract.Contract.SupportsInterface(&_Contract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contract *ContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contract.Contract.SupportsInterface(&_Contract.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractSession) Symbol() (string, error) {
	return _Contract.Contract.Symbol(&_Contract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractCallerSession) Symbol() (string, error) {
	return _Contract.Contract.Symbol(&_Contract.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Contract *ContractCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Contract *ContractSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Contract.Contract.TokenURI(&_Contract.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Contract *ContractCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Contract.Contract.TokenURI(&_Contract.CallOpts, tokenId)
}

// VerifyIdentitiesBinding is a free data retrieval call binding the contract method 0x49d03037.
//
// Solidity: function verifyIdentitiesBinding(uint256 id, address nftOwnerAddress, address sigAddress, string[] userIDs, bytes32 multiIdentitiesRoot, bytes signature) view returns(bool)
func (_Contract *ContractCaller) VerifyIdentitiesBinding(opts *bind.CallOpts, id *big.Int, nftOwnerAddress common.Address, sigAddress common.Address, userIDs []string, multiIdentitiesRoot [32]byte, signature []byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "verifyIdentitiesBinding", id, nftOwnerAddress, sigAddress, userIDs, multiIdentitiesRoot, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyIdentitiesBinding is a free data retrieval call binding the contract method 0x49d03037.
//
// Solidity: function verifyIdentitiesBinding(uint256 id, address nftOwnerAddress, address sigAddress, string[] userIDs, bytes32 multiIdentitiesRoot, bytes signature) view returns(bool)
func (_Contract *ContractSession) VerifyIdentitiesBinding(id *big.Int, nftOwnerAddress common.Address, sigAddress common.Address, userIDs []string, multiIdentitiesRoot [32]byte, signature []byte) (bool, error) {
	return _Contract.Contract.VerifyIdentitiesBinding(&_Contract.CallOpts, id, nftOwnerAddress, sigAddress, userIDs, multiIdentitiesRoot, signature)
}

// VerifyIdentitiesBinding is a free data retrieval call binding the contract method 0x49d03037.
//
// Solidity: function verifyIdentitiesBinding(uint256 id, address nftOwnerAddress, address sigAddress, string[] userIDs, bytes32 multiIdentitiesRoot, bytes signature) view returns(bool)
func (_Contract *ContractCallerSession) VerifyIdentitiesBinding(id *big.Int, nftOwnerAddress common.Address, sigAddress common.Address, userIDs []string, multiIdentitiesRoot [32]byte, signature []byte) (bool, error) {
	return _Contract.Contract.VerifyIdentitiesBinding(&_Contract.CallOpts, id, nftOwnerAddress, sigAddress, userIDs, multiIdentitiesRoot, signature)
}

// CarvProtocolServiceInit is a paid mutator transaction binding the contract method 0xa28bfc43.
//
// Solidity: function __CarvProtocolService_init(address rewards_address) returns()
func (_Contract *ContractTransactor) CarvProtocolServiceInit(opts *bind.TransactOpts, rewards_address common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "__CarvProtocolService_init", rewards_address)
}

// CarvProtocolServiceInit is a paid mutator transaction binding the contract method 0xa28bfc43.
//
// Solidity: function __CarvProtocolService_init(address rewards_address) returns()
func (_Contract *ContractSession) CarvProtocolServiceInit(rewards_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.CarvProtocolServiceInit(&_Contract.TransactOpts, rewards_address)
}

// CarvProtocolServiceInit is a paid mutator transaction binding the contract method 0xa28bfc43.
//
// Solidity: function __CarvProtocolService_init(address rewards_address) returns()
func (_Contract *ContractTransactorSession) CarvProtocolServiceInit(rewards_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.CarvProtocolServiceInit(&_Contract.TransactOpts, rewards_address)
}

// AddTeeRole is a paid mutator transaction binding the contract method 0xa35d20be.
//
// Solidity: function add_tee_role(address tee_address) returns()
func (_Contract *ContractTransactor) AddTeeRole(opts *bind.TransactOpts, tee_address common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "add_tee_role", tee_address)
}

// AddTeeRole is a paid mutator transaction binding the contract method 0xa35d20be.
//
// Solidity: function add_tee_role(address tee_address) returns()
func (_Contract *ContractSession) AddTeeRole(tee_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddTeeRole(&_Contract.TransactOpts, tee_address)
}

// AddTeeRole is a paid mutator transaction binding the contract method 0xa35d20be.
//
// Solidity: function add_tee_role(address tee_address) returns()
func (_Contract *ContractTransactorSession) AddTeeRole(tee_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddTeeRole(&_Contract.TransactOpts, tee_address)
}

// AddVerifierRole is a paid mutator transaction binding the contract method 0xad86d545.
//
// Solidity: function add_verifier_role(address verifier_address) returns()
func (_Contract *ContractTransactor) AddVerifierRole(opts *bind.TransactOpts, verifier_address common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "add_verifier_role", verifier_address)
}

// AddVerifierRole is a paid mutator transaction binding the contract method 0xad86d545.
//
// Solidity: function add_verifier_role(address verifier_address) returns()
func (_Contract *ContractSession) AddVerifierRole(verifier_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddVerifierRole(&_Contract.TransactOpts, verifier_address)
}

// AddVerifierRole is a paid mutator transaction binding the contract method 0xad86d545.
//
// Solidity: function add_verifier_role(address verifier_address) returns()
func (_Contract *ContractTransactorSession) AddVerifierRole(verifier_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddVerifierRole(&_Contract.TransactOpts, verifier_address)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Contract *ContractTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Contract *ContractSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, to, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Contract *ContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Contract *ContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contract.Contract.GrantRole(&_Contract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Contract *ContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contract.Contract.GrantRole(&_Contract.TransactOpts, role, account)
}

// JoinCampaign is a paid mutator transaction binding the contract method 0x09a9c7f0.
//
// Solidity: function join_campaign(uint256 carv_id, string campaign_id, string join_campaign_info) returns()
func (_Contract *ContractTransactor) JoinCampaign(opts *bind.TransactOpts, carv_id *big.Int, campaign_id string, join_campaign_info string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "join_campaign", carv_id, campaign_id, join_campaign_info)
}

// JoinCampaign is a paid mutator transaction binding the contract method 0x09a9c7f0.
//
// Solidity: function join_campaign(uint256 carv_id, string campaign_id, string join_campaign_info) returns()
func (_Contract *ContractSession) JoinCampaign(carv_id *big.Int, campaign_id string, join_campaign_info string) (*types.Transaction, error) {
	return _Contract.Contract.JoinCampaign(&_Contract.TransactOpts, carv_id, campaign_id, join_campaign_info)
}

// JoinCampaign is a paid mutator transaction binding the contract method 0x09a9c7f0.
//
// Solidity: function join_campaign(uint256 carv_id, string campaign_id, string join_campaign_info) returns()
func (_Contract *ContractTransactorSession) JoinCampaign(carv_id *big.Int, campaign_id string, join_campaign_info string) (*types.Transaction, error) {
	return _Contract.Contract.JoinCampaign(&_Contract.TransactOpts, carv_id, campaign_id, join_campaign_info)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Contract *ContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Contract *ContractSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RenounceRole(&_Contract.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Contract *ContractTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RenounceRole(&_Contract.TransactOpts, role, account)
}

// ReportTeeAttestation is a paid mutator transaction binding the contract method 0xe97a9987.
//
// Solidity: function report_tee_attestation(string campaign_id, string attestation) returns()
func (_Contract *ContractTransactor) ReportTeeAttestation(opts *bind.TransactOpts, campaign_id string, attestation string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "report_tee_attestation", campaign_id, attestation)
}

// ReportTeeAttestation is a paid mutator transaction binding the contract method 0xe97a9987.
//
// Solidity: function report_tee_attestation(string campaign_id, string attestation) returns()
func (_Contract *ContractSession) ReportTeeAttestation(campaign_id string, attestation string) (*types.Transaction, error) {
	return _Contract.Contract.ReportTeeAttestation(&_Contract.TransactOpts, campaign_id, attestation)
}

// ReportTeeAttestation is a paid mutator transaction binding the contract method 0xe97a9987.
//
// Solidity: function report_tee_attestation(string campaign_id, string attestation) returns()
func (_Contract *ContractTransactorSession) ReportTeeAttestation(campaign_id string, attestation string) (*types.Transaction, error) {
	return _Contract.Contract.ReportTeeAttestation(&_Contract.TransactOpts, campaign_id, attestation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Contract *ContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Contract *ContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RevokeRole(&_Contract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Contract *ContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RevokeRole(&_Contract.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFrom(&_Contract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFrom(&_Contract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Contract *ContractTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Contract *ContractSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFrom0(&_Contract.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Contract *ContractTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFrom0(&_Contract.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Contract *ContractTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Contract *ContractSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.Contract.SetApprovalForAll(&_Contract.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Contract *ContractTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.Contract.SetApprovalForAll(&_Contract.TransactOpts, operator, approved)
}

// SetIdentitiesRootFunc is a paid mutator transaction binding the contract method 0x6f6df14c.
//
// Solidity: function setIdentitiesRootFunc(uint256 id, bytes32 multiIdentitiesRoot) returns()
func (_Contract *ContractTransactor) SetIdentitiesRootFunc(opts *bind.TransactOpts, id *big.Int, multiIdentitiesRoot [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setIdentitiesRootFunc", id, multiIdentitiesRoot)
}

// SetIdentitiesRootFunc is a paid mutator transaction binding the contract method 0x6f6df14c.
//
// Solidity: function setIdentitiesRootFunc(uint256 id, bytes32 multiIdentitiesRoot) returns()
func (_Contract *ContractSession) SetIdentitiesRootFunc(id *big.Int, multiIdentitiesRoot [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SetIdentitiesRootFunc(&_Contract.TransactOpts, id, multiIdentitiesRoot)
}

// SetIdentitiesRootFunc is a paid mutator transaction binding the contract method 0x6f6df14c.
//
// Solidity: function setIdentitiesRootFunc(uint256 id, bytes32 multiIdentitiesRoot) returns()
func (_Contract *ContractTransactorSession) SetIdentitiesRootFunc(id *big.Int, multiIdentitiesRoot [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SetIdentitiesRootFunc(&_Contract.TransactOpts, id, multiIdentitiesRoot)
}

// SetIdentitiesRoot is a paid mutator transaction binding the contract method 0x1a7f9955.
//
// Solidity: function set_identities_root(address user_address, string user_profile_path, uint256 profile_version, bytes32 multiIdentitiesRoot, bytes signature) returns()
func (_Contract *ContractTransactor) SetIdentitiesRoot(opts *bind.TransactOpts, user_address common.Address, user_profile_path string, profile_version *big.Int, multiIdentitiesRoot [32]byte, signature []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "set_identities_root", user_address, user_profile_path, profile_version, multiIdentitiesRoot, signature)
}

// SetIdentitiesRoot is a paid mutator transaction binding the contract method 0x1a7f9955.
//
// Solidity: function set_identities_root(address user_address, string user_profile_path, uint256 profile_version, bytes32 multiIdentitiesRoot, bytes signature) returns()
func (_Contract *ContractSession) SetIdentitiesRoot(user_address common.Address, user_profile_path string, profile_version *big.Int, multiIdentitiesRoot [32]byte, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetIdentitiesRoot(&_Contract.TransactOpts, user_address, user_profile_path, profile_version, multiIdentitiesRoot, signature)
}

// SetIdentitiesRoot is a paid mutator transaction binding the contract method 0x1a7f9955.
//
// Solidity: function set_identities_root(address user_address, string user_profile_path, uint256 profile_version, bytes32 multiIdentitiesRoot, bytes signature) returns()
func (_Contract *ContractTransactorSession) SetIdentitiesRoot(user_address common.Address, user_profile_path string, profile_version *big.Int, multiIdentitiesRoot [32]byte, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetIdentitiesRoot(&_Contract.TransactOpts, user_address, user_profile_path, profile_version, multiIdentitiesRoot, signature)
}

// SetPayAddress is a paid mutator transaction binding the contract method 0x7e3461a3.
//
// Solidity: function set_pay_address(address pay_address) returns()
func (_Contract *ContractTransactor) SetPayAddress(opts *bind.TransactOpts, pay_address common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "set_pay_address", pay_address)
}

// SetPayAddress is a paid mutator transaction binding the contract method 0x7e3461a3.
//
// Solidity: function set_pay_address(address pay_address) returns()
func (_Contract *ContractSession) SetPayAddress(pay_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetPayAddress(&_Contract.TransactOpts, pay_address)
}

// SetPayAddress is a paid mutator transaction binding the contract method 0x7e3461a3.
//
// Solidity: function set_pay_address(address pay_address) returns()
func (_Contract *ContractTransactorSession) SetPayAddress(pay_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetPayAddress(&_Contract.TransactOpts, pay_address)
}

// SubmitCampaign is a paid mutator transaction binding the contract method 0x3736c83f.
//
// Solidity: function submit_campaign((string,address,uint256,uint256,address,uint8) reward_info, (string,string,address,uint8,address,uint256,uint256,uint8,uint256,uint256,string) campaign_info) payable returns()
func (_Contract *ContractTransactor) SubmitCampaign(opts *bind.TransactOpts, reward_info CarvProtocolServicereward, campaign_info CarvProtocolServicecampaign) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "submit_campaign", reward_info, campaign_info)
}

// SubmitCampaign is a paid mutator transaction binding the contract method 0x3736c83f.
//
// Solidity: function submit_campaign((string,address,uint256,uint256,address,uint8) reward_info, (string,string,address,uint8,address,uint256,uint256,uint8,uint256,uint256,string) campaign_info) payable returns()
func (_Contract *ContractSession) SubmitCampaign(reward_info CarvProtocolServicereward, campaign_info CarvProtocolServicecampaign) (*types.Transaction, error) {
	return _Contract.Contract.SubmitCampaign(&_Contract.TransactOpts, reward_info, campaign_info)
}

// SubmitCampaign is a paid mutator transaction binding the contract method 0x3736c83f.
//
// Solidity: function submit_campaign((string,address,uint256,uint256,address,uint8) reward_info, (string,string,address,uint8,address,uint256,uint256,uint8,uint256,uint256,string) campaign_info) payable returns()
func (_Contract *ContractTransactorSession) SubmitCampaign(reward_info CarvProtocolServicereward, campaign_info CarvProtocolServicecampaign) (*types.Transaction, error) {
	return _Contract.Contract.SubmitCampaign(&_Contract.TransactOpts, reward_info, campaign_info)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom(&_Contract.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom(&_Contract.TransactOpts, from, to, tokenId)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x578907ca.
//
// Solidity: function verify_attestation(bytes32 attestation_id, bool result) returns()
func (_Contract *ContractTransactor) VerifyAttestation(opts *bind.TransactOpts, attestation_id [32]byte, result bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "verify_attestation", attestation_id, result)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x578907ca.
//
// Solidity: function verify_attestation(bytes32 attestation_id, bool result) returns()
func (_Contract *ContractSession) VerifyAttestation(attestation_id [32]byte, result bool) (*types.Transaction, error) {
	return _Contract.Contract.VerifyAttestation(&_Contract.TransactOpts, attestation_id, result)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x578907ca.
//
// Solidity: function verify_attestation(bytes32 attestation_id, bool result) returns()
func (_Contract *ContractTransactorSession) VerifyAttestation(attestation_id [32]byte, result bool) (*types.Transaction, error) {
	return _Contract.Contract.VerifyAttestation(&_Contract.TransactOpts, attestation_id, result)
}

// VerifyAttestationBatch is a paid mutator transaction binding the contract method 0xc1f75848.
//
// Solidity: function verify_attestation_batch(bytes32[] attestation_ids, bool[] results) returns()
func (_Contract *ContractTransactor) VerifyAttestationBatch(opts *bind.TransactOpts, attestation_ids [][32]byte, results []bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "verify_attestation_batch", attestation_ids, results)
}

// VerifyAttestationBatch is a paid mutator transaction binding the contract method 0xc1f75848.
//
// Solidity: function verify_attestation_batch(bytes32[] attestation_ids, bool[] results) returns()
func (_Contract *ContractSession) VerifyAttestationBatch(attestation_ids [][32]byte, results []bool) (*types.Transaction, error) {
	return _Contract.Contract.VerifyAttestationBatch(&_Contract.TransactOpts, attestation_ids, results)
}

// VerifyAttestationBatch is a paid mutator transaction binding the contract method 0xc1f75848.
//
// Solidity: function verify_attestation_batch(bytes32[] attestation_ids, bool[] results) returns()
func (_Contract *ContractTransactorSession) VerifyAttestationBatch(attestation_ids [][32]byte, results []bool) (*types.Transaction, error) {
	return _Contract.Contract.VerifyAttestationBatch(&_Contract.TransactOpts, attestation_ids, results)
}

// ContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Contract contract.
type ContractApprovalIterator struct {
	Event *ContractApproval // Event containing the contract specifics and raw log

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
func (it *ContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractApproval)
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
		it.Event = new(ContractApproval)
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
func (it *ContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractApproval represents a Approval event raised by the Contract contract.
type ContractApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Contract *ContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractApprovalIterator{contract: _Contract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Contract *ContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractApproval)
				if err := _Contract.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Contract *ContractFilterer) ParseApproval(log types.Log) (*ContractApproval, error) {
	event := new(ContractApproval)
	if err := _Contract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Contract contract.
type ContractApprovalForAllIterator struct {
	Event *ContractApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ContractApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractApprovalForAll)
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
		it.Event = new(ContractApprovalForAll)
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
func (it *ContractApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractApprovalForAll represents a ApprovalForAll event raised by the Contract contract.
type ContractApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Contract *ContractFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ContractApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractApprovalForAllIterator{contract: _Contract.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Contract *ContractFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ContractApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractApprovalForAll)
				if err := _Contract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Contract *ContractFilterer) ParseApprovalForAll(log types.Log) (*ContractApprovalForAll, error) {
	event := new(ContractApprovalForAll)
	if err := _Contract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ContractMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the Contract contract.
type ContractMintedIterator struct {
	Event *ContractMinted // Event containing the contract specifics and raw log

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
func (it *ContractMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMinted)
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
		it.Event = new(ContractMinted)
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
func (it *ContractMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMinted represents a Minted event raised by the Contract contract.
type ContractMinted struct {
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address to, uint256 token_id)
func (_Contract *ContractFilterer) FilterMinted(opts *bind.FilterOpts) (*ContractMintedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return &ContractMintedIterator{contract: _Contract.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address to, uint256 token_id)
func (_Contract *ContractFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *ContractMinted) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMinted)
				if err := _Contract.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address to, uint256 token_id)
func (_Contract *ContractFilterer) ParseMinted(log types.Log) (*ContractMinted, error) {
	event := new(ContractMinted)
	if err := _Contract.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractProfixPayedIterator is returned from FilterProfixPayed and is used to iterate over the raw logs and unpacked data for ProfixPayed events raised by the Contract contract.
type ContractProfixPayedIterator struct {
	Event *ContractProfixPayed // Event containing the contract specifics and raw log

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
func (it *ContractProfixPayedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractProfixPayed)
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
		it.Event = new(ContractProfixPayed)
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
func (it *ContractProfixPayedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractProfixPayedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractProfixPayed represents a ProfixPayed event raised by the Contract contract.
type ContractProfixPayed struct {
	Erc20Address common.Address
	FromAddress  common.Address
	ToAddress    common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProfixPayed is a free log retrieval operation binding the contract event 0x3832ba0f9bc9de56bd1f974e665b53f2643ed48a47b45c3d36976bc1ea314599.
//
// Solidity: event ProfixPayed(address erc20_address, address from_address, address to_address, uint256 amount)
func (_Contract *ContractFilterer) FilterProfixPayed(opts *bind.FilterOpts) (*ContractProfixPayedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ProfixPayed")
	if err != nil {
		return nil, err
	}
	return &ContractProfixPayedIterator{contract: _Contract.contract, event: "ProfixPayed", logs: logs, sub: sub}, nil
}

// WatchProfixPayed is a free log subscription operation binding the contract event 0x3832ba0f9bc9de56bd1f974e665b53f2643ed48a47b45c3d36976bc1ea314599.
//
// Solidity: event ProfixPayed(address erc20_address, address from_address, address to_address, uint256 amount)
func (_Contract *ContractFilterer) WatchProfixPayed(opts *bind.WatchOpts, sink chan<- *ContractProfixPayed) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ProfixPayed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractProfixPayed)
				if err := _Contract.contract.UnpackLog(event, "ProfixPayed", log); err != nil {
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

// ParseProfixPayed is a log parse operation binding the contract event 0x3832ba0f9bc9de56bd1f974e665b53f2643ed48a47b45c3d36976bc1ea314599.
//
// Solidity: event ProfixPayed(address erc20_address, address from_address, address to_address, uint256 amount)
func (_Contract *ContractFilterer) ParseProfixPayed(log types.Log) (*ContractProfixPayed, error) {
	event := new(ContractProfixPayed)
	if err := _Contract.contract.UnpackLog(event, "ProfixPayed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractReportTeeAttestationIterator is returned from FilterReportTeeAttestation and is used to iterate over the raw logs and unpacked data for ReportTeeAttestation events raised by the Contract contract.
type ContractReportTeeAttestationIterator struct {
	Event *ContractReportTeeAttestation // Event containing the contract specifics and raw log

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
func (it *ContractReportTeeAttestationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractReportTeeAttestation)
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
		it.Event = new(ContractReportTeeAttestation)
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
func (it *ContractReportTeeAttestationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractReportTeeAttestationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractReportTeeAttestation represents a ReportTeeAttestation event raised by the Contract contract.
type ContractReportTeeAttestation struct {
	TeeAddress    common.Address
	CampaignId    string
	AttestationId [32]byte
	Attestation   string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReportTeeAttestation is a free log retrieval operation binding the contract event 0x99a038e9d345d0b12130b3b1fb003bf8f2d3a5c27ce2a800bbb1608efff6c591.
//
// Solidity: event ReportTeeAttestation(address tee_address, string campaign_id, bytes32 attestation_id, string attestation)
func (_Contract *ContractFilterer) FilterReportTeeAttestation(opts *bind.FilterOpts) (*ContractReportTeeAttestationIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ReportTeeAttestation")
	if err != nil {
		return nil, err
	}
	return &ContractReportTeeAttestationIterator{contract: _Contract.contract, event: "ReportTeeAttestation", logs: logs, sub: sub}, nil
}

// WatchReportTeeAttestation is a free log subscription operation binding the contract event 0x99a038e9d345d0b12130b3b1fb003bf8f2d3a5c27ce2a800bbb1608efff6c591.
//
// Solidity: event ReportTeeAttestation(address tee_address, string campaign_id, bytes32 attestation_id, string attestation)
func (_Contract *ContractFilterer) WatchReportTeeAttestation(opts *bind.WatchOpts, sink chan<- *ContractReportTeeAttestation) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ReportTeeAttestation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractReportTeeAttestation)
				if err := _Contract.contract.UnpackLog(event, "ReportTeeAttestation", log); err != nil {
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

// ParseReportTeeAttestation is a log parse operation binding the contract event 0x99a038e9d345d0b12130b3b1fb003bf8f2d3a5c27ce2a800bbb1608efff6c591.
//
// Solidity: event ReportTeeAttestation(address tee_address, string campaign_id, bytes32 attestation_id, string attestation)
func (_Contract *ContractFilterer) ParseReportTeeAttestation(log types.Log) (*ContractReportTeeAttestation, error) {
	event := new(ContractReportTeeAttestation)
	if err := _Contract.contract.UnpackLog(event, "ReportTeeAttestation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardPayedIterator is returned from FilterRewardPayed and is used to iterate over the raw logs and unpacked data for RewardPayed events raised by the Contract contract.
type ContractRewardPayedIterator struct {
	Event *ContractRewardPayed // Event containing the contract specifics and raw log

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
func (it *ContractRewardPayedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardPayed)
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
		it.Event = new(ContractRewardPayed)
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
func (it *ContractRewardPayedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardPayedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardPayed represents a RewardPayed event raised by the Contract contract.
type ContractRewardPayed struct {
	Erc20Address common.Address
	FromAddress  common.Address
	ToAddress    common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRewardPayed is a free log retrieval operation binding the contract event 0x704db66c978a7472842ddae0d478a8a7d123fb6ad449e4426b5e0f6a22081de7.
//
// Solidity: event RewardPayed(address erc20_address, address from_address, address to_address, uint256 amount)
func (_Contract *ContractFilterer) FilterRewardPayed(opts *bind.FilterOpts) (*ContractRewardPayedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RewardPayed")
	if err != nil {
		return nil, err
	}
	return &ContractRewardPayedIterator{contract: _Contract.contract, event: "RewardPayed", logs: logs, sub: sub}, nil
}

// WatchRewardPayed is a free log subscription operation binding the contract event 0x704db66c978a7472842ddae0d478a8a7d123fb6ad449e4426b5e0f6a22081de7.
//
// Solidity: event RewardPayed(address erc20_address, address from_address, address to_address, uint256 amount)
func (_Contract *ContractFilterer) WatchRewardPayed(opts *bind.WatchOpts, sink chan<- *ContractRewardPayed) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RewardPayed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardPayed)
				if err := _Contract.contract.UnpackLog(event, "RewardPayed", log); err != nil {
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

// ParseRewardPayed is a log parse operation binding the contract event 0x704db66c978a7472842ddae0d478a8a7d123fb6ad449e4426b5e0f6a22081de7.
//
// Solidity: event RewardPayed(address erc20_address, address from_address, address to_address, uint256 amount)
func (_Contract *ContractFilterer) ParseRewardPayed(log types.Log) (*ContractRewardPayed, error) {
	event := new(ContractRewardPayed)
	if err := _Contract.contract.UnpackLog(event, "RewardPayed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Contract contract.
type ContractRoleAdminChangedIterator struct {
	Event *ContractRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ContractRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRoleAdminChanged)
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
		it.Event = new(ContractRoleAdminChanged)
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
func (it *ContractRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRoleAdminChanged represents a RoleAdminChanged event raised by the Contract contract.
type ContractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Contract *ContractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ContractRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ContractRoleAdminChangedIterator{contract: _Contract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Contract *ContractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ContractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRoleAdminChanged)
				if err := _Contract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Contract *ContractFilterer) ParseRoleAdminChanged(log types.Log) (*ContractRoleAdminChanged, error) {
	event := new(ContractRoleAdminChanged)
	if err := _Contract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Contract contract.
type ContractRoleGrantedIterator struct {
	Event *ContractRoleGranted // Event containing the contract specifics and raw log

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
func (it *ContractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRoleGranted)
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
		it.Event = new(ContractRoleGranted)
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
func (it *ContractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRoleGranted represents a RoleGranted event raised by the Contract contract.
type ContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Contract *ContractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractRoleGrantedIterator, error) {

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

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractRoleGrantedIterator{contract: _Contract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Contract *ContractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ContractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRoleGranted)
				if err := _Contract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Contract *ContractFilterer) ParseRoleGranted(log types.Log) (*ContractRoleGranted, error) {
	event := new(ContractRoleGranted)
	if err := _Contract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Contract contract.
type ContractRoleRevokedIterator struct {
	Event *ContractRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ContractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRoleRevoked)
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
		it.Event = new(ContractRoleRevoked)
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
func (it *ContractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRoleRevoked represents a RoleRevoked event raised by the Contract contract.
type ContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Contract *ContractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractRoleRevokedIterator, error) {

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

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractRoleRevokedIterator{contract: _Contract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Contract *ContractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ContractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRoleRevoked)
				if err := _Contract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Contract *ContractFilterer) ParseRoleRevoked(log types.Log) (*ContractRoleRevoked, error) {
	event := new(ContractRoleRevoked)
	if err := _Contract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSetIdentitiesRootIterator is returned from FilterSetIdentitiesRoot and is used to iterate over the raw logs and unpacked data for SetIdentitiesRoot events raised by the Contract contract.
type ContractSetIdentitiesRootIterator struct {
	Event *ContractSetIdentitiesRoot // Event containing the contract specifics and raw log

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
func (it *ContractSetIdentitiesRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetIdentitiesRoot)
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
		it.Event = new(ContractSetIdentitiesRoot)
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
func (it *ContractSetIdentitiesRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetIdentitiesRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetIdentitiesRoot represents a SetIdentitiesRoot event raised by the Contract contract.
type ContractSetIdentitiesRoot struct {
	Id             *big.Int
	IdentitiesRoot [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetIdentitiesRoot is a free log retrieval operation binding the contract event 0x21b5a040ca6f61b8bcacb8f423e25ce46e88932f887ce92e60343369c20ec06f.
//
// Solidity: event SetIdentitiesRoot(uint256 id, bytes32 identitiesRoot)
func (_Contract *ContractFilterer) FilterSetIdentitiesRoot(opts *bind.FilterOpts) (*ContractSetIdentitiesRootIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetIdentitiesRoot")
	if err != nil {
		return nil, err
	}
	return &ContractSetIdentitiesRootIterator{contract: _Contract.contract, event: "SetIdentitiesRoot", logs: logs, sub: sub}, nil
}

// WatchSetIdentitiesRoot is a free log subscription operation binding the contract event 0x21b5a040ca6f61b8bcacb8f423e25ce46e88932f887ce92e60343369c20ec06f.
//
// Solidity: event SetIdentitiesRoot(uint256 id, bytes32 identitiesRoot)
func (_Contract *ContractFilterer) WatchSetIdentitiesRoot(opts *bind.WatchOpts, sink chan<- *ContractSetIdentitiesRoot) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetIdentitiesRoot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetIdentitiesRoot)
				if err := _Contract.contract.UnpackLog(event, "SetIdentitiesRoot", log); err != nil {
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

// ParseSetIdentitiesRoot is a log parse operation binding the contract event 0x21b5a040ca6f61b8bcacb8f423e25ce46e88932f887ce92e60343369c20ec06f.
//
// Solidity: event SetIdentitiesRoot(uint256 id, bytes32 identitiesRoot)
func (_Contract *ContractFilterer) ParseSetIdentitiesRoot(log types.Log) (*ContractSetIdentitiesRoot, error) {
	event := new(ContractSetIdentitiesRoot)
	if err := _Contract.contract.UnpackLog(event, "SetIdentitiesRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSubmitCampaignIterator is returned from FilterSubmitCampaign and is used to iterate over the raw logs and unpacked data for SubmitCampaign events raised by the Contract contract.
type ContractSubmitCampaignIterator struct {
	Event *ContractSubmitCampaign // Event containing the contract specifics and raw log

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
func (it *ContractSubmitCampaignIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSubmitCampaign)
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
		it.Event = new(ContractSubmitCampaign)
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
func (it *ContractSubmitCampaignIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSubmitCampaignIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSubmitCampaign represents a SubmitCampaign event raised by the Contract contract.
type ContractSubmitCampaign struct {
	ContractAddress common.Address
	CampaignId      string
	Requirements    string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSubmitCampaign is a free log retrieval operation binding the contract event 0xe8e692af5f416c4637c9307d7f08b7697fd8f4398f55103ce03ddd9388ccd84d.
//
// Solidity: event SubmitCampaign(address contract_address, string campaign_id, string requirements)
func (_Contract *ContractFilterer) FilterSubmitCampaign(opts *bind.FilterOpts) (*ContractSubmitCampaignIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SubmitCampaign")
	if err != nil {
		return nil, err
	}
	return &ContractSubmitCampaignIterator{contract: _Contract.contract, event: "SubmitCampaign", logs: logs, sub: sub}, nil
}

// WatchSubmitCampaign is a free log subscription operation binding the contract event 0xe8e692af5f416c4637c9307d7f08b7697fd8f4398f55103ce03ddd9388ccd84d.
//
// Solidity: event SubmitCampaign(address contract_address, string campaign_id, string requirements)
func (_Contract *ContractFilterer) WatchSubmitCampaign(opts *bind.WatchOpts, sink chan<- *ContractSubmitCampaign) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SubmitCampaign")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSubmitCampaign)
				if err := _Contract.contract.UnpackLog(event, "SubmitCampaign", log); err != nil {
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

// ParseSubmitCampaign is a log parse operation binding the contract event 0xe8e692af5f416c4637c9307d7f08b7697fd8f4398f55103ce03ddd9388ccd84d.
//
// Solidity: event SubmitCampaign(address contract_address, string campaign_id, string requirements)
func (_Contract *ContractFilterer) ParseSubmitCampaign(log types.Log) (*ContractSubmitCampaign, error) {
	event := new(ContractSubmitCampaign)
	if err := _Contract.contract.UnpackLog(event, "SubmitCampaign", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Contract contract.
type ContractTransferIterator struct {
	Event *ContractTransfer // Event containing the contract specifics and raw log

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
func (it *ContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTransfer)
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
		it.Event = new(ContractTransfer)
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
func (it *ContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTransfer represents a Transfer event raised by the Contract contract.
type ContractTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Contract *ContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractTransferIterator{contract: _Contract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Contract *ContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTransfer)
				if err := _Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Contract *ContractFilterer) ParseTransfer(log types.Log) (*ContractTransfer, error) {
	event := new(ContractTransfer)
	if err := _Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUserCampaignDataIterator is returned from FilterUserCampaignData and is used to iterate over the raw logs and unpacked data for UserCampaignData events raised by the Contract contract.
type ContractUserCampaignDataIterator struct {
	Event *ContractUserCampaignData // Event containing the contract specifics and raw log

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
func (it *ContractUserCampaignDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUserCampaignData)
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
		it.Event = new(ContractUserCampaignData)
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
func (it *ContractUserCampaignDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUserCampaignDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUserCampaignData represents a UserCampaignData event raised by the Contract contract.
type ContractUserCampaignData struct {
	CarvId       *big.Int
	CampaignId   string
	CampaignInfo string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUserCampaignData is a free log retrieval operation binding the contract event 0x83eda6edb918aac3f35c8d616a6198dbb75eeca3248e0efec05e66a30d511191.
//
// Solidity: event UserCampaignData(uint256 carv_id, string campaign_id, string campaign_info)
func (_Contract *ContractFilterer) FilterUserCampaignData(opts *bind.FilterOpts) (*ContractUserCampaignDataIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UserCampaignData")
	if err != nil {
		return nil, err
	}
	return &ContractUserCampaignDataIterator{contract: _Contract.contract, event: "UserCampaignData", logs: logs, sub: sub}, nil
}

// WatchUserCampaignData is a free log subscription operation binding the contract event 0x83eda6edb918aac3f35c8d616a6198dbb75eeca3248e0efec05e66a30d511191.
//
// Solidity: event UserCampaignData(uint256 carv_id, string campaign_id, string campaign_info)
func (_Contract *ContractFilterer) WatchUserCampaignData(opts *bind.WatchOpts, sink chan<- *ContractUserCampaignData) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UserCampaignData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUserCampaignData)
				if err := _Contract.contract.UnpackLog(event, "UserCampaignData", log); err != nil {
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

// ParseUserCampaignData is a log parse operation binding the contract event 0x83eda6edb918aac3f35c8d616a6198dbb75eeca3248e0efec05e66a30d511191.
//
// Solidity: event UserCampaignData(uint256 carv_id, string campaign_id, string campaign_info)
func (_Contract *ContractFilterer) ParseUserCampaignData(log types.Log) (*ContractUserCampaignData, error) {
	event := new(ContractUserCampaignData)
	if err := _Contract.contract.UnpackLog(event, "UserCampaignData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractVerifyAttestationIterator is returned from FilterVerifyAttestation and is used to iterate over the raw logs and unpacked data for VerifyAttestation events raised by the Contract contract.
type ContractVerifyAttestationIterator struct {
	Event *ContractVerifyAttestation // Event containing the contract specifics and raw log

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
func (it *ContractVerifyAttestationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractVerifyAttestation)
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
		it.Event = new(ContractVerifyAttestation)
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
func (it *ContractVerifyAttestationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractVerifyAttestationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractVerifyAttestation represents a VerifyAttestation event raised by the Contract contract.
type ContractVerifyAttestation struct {
	VerifierAddress common.Address
	AttestationId   [32]byte
	Result          bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVerifyAttestation is a free log retrieval operation binding the contract event 0x937b4d1b5b461852fdbbb174b68adefb42155e668c05a1261fae11a2d22a240a.
//
// Solidity: event VerifyAttestation(address verifier_address, bytes32 attestation_id, bool result)
func (_Contract *ContractFilterer) FilterVerifyAttestation(opts *bind.FilterOpts) (*ContractVerifyAttestationIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "VerifyAttestation")
	if err != nil {
		return nil, err
	}
	return &ContractVerifyAttestationIterator{contract: _Contract.contract, event: "VerifyAttestation", logs: logs, sub: sub}, nil
}

// WatchVerifyAttestation is a free log subscription operation binding the contract event 0x937b4d1b5b461852fdbbb174b68adefb42155e668c05a1261fae11a2d22a240a.
//
// Solidity: event VerifyAttestation(address verifier_address, bytes32 attestation_id, bool result)
func (_Contract *ContractFilterer) WatchVerifyAttestation(opts *bind.WatchOpts, sink chan<- *ContractVerifyAttestation) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "VerifyAttestation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractVerifyAttestation)
				if err := _Contract.contract.UnpackLog(event, "VerifyAttestation", log); err != nil {
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

// ParseVerifyAttestation is a log parse operation binding the contract event 0x937b4d1b5b461852fdbbb174b68adefb42155e668c05a1261fae11a2d22a240a.
//
// Solidity: event VerifyAttestation(address verifier_address, bytes32 attestation_id, bool result)
func (_Contract *ContractFilterer) ParseVerifyAttestation(log types.Log) (*ContractVerifyAttestation, error) {
	event := new(ContractVerifyAttestation)
	if err := _Contract.contract.UnpackLog(event, "VerifyAttestation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractVerifyAttestationBatchIterator is returned from FilterVerifyAttestationBatch and is used to iterate over the raw logs and unpacked data for VerifyAttestationBatch events raised by the Contract contract.
type ContractVerifyAttestationBatchIterator struct {
	Event *ContractVerifyAttestationBatch // Event containing the contract specifics and raw log

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
func (it *ContractVerifyAttestationBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractVerifyAttestationBatch)
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
		it.Event = new(ContractVerifyAttestationBatch)
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
func (it *ContractVerifyAttestationBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractVerifyAttestationBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractVerifyAttestationBatch represents a VerifyAttestationBatch event raised by the Contract contract.
type ContractVerifyAttestationBatch struct {
	VerifierAddress common.Address
	AttestationIds  [][32]byte
	Results         []bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVerifyAttestationBatch is a free log retrieval operation binding the contract event 0x04798cc2c828443d2d9ef127deaa1f73e782d8a9fffbc9db2bcca53e12765af5.
//
// Solidity: event VerifyAttestationBatch(address verifier_address, bytes32[] attestation_ids, bool[] results)
func (_Contract *ContractFilterer) FilterVerifyAttestationBatch(opts *bind.FilterOpts) (*ContractVerifyAttestationBatchIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "VerifyAttestationBatch")
	if err != nil {
		return nil, err
	}
	return &ContractVerifyAttestationBatchIterator{contract: _Contract.contract, event: "VerifyAttestationBatch", logs: logs, sub: sub}, nil
}

// WatchVerifyAttestationBatch is a free log subscription operation binding the contract event 0x04798cc2c828443d2d9ef127deaa1f73e782d8a9fffbc9db2bcca53e12765af5.
//
// Solidity: event VerifyAttestationBatch(address verifier_address, bytes32[] attestation_ids, bool[] results)
func (_Contract *ContractFilterer) WatchVerifyAttestationBatch(opts *bind.WatchOpts, sink chan<- *ContractVerifyAttestationBatch) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "VerifyAttestationBatch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractVerifyAttestationBatch)
				if err := _Contract.contract.UnpackLog(event, "VerifyAttestationBatch", log); err != nil {
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

// ParseVerifyAttestationBatch is a log parse operation binding the contract event 0x04798cc2c828443d2d9ef127deaa1f73e782d8a9fffbc9db2bcca53e12765af5.
//
// Solidity: event VerifyAttestationBatch(address verifier_address, bytes32[] attestation_ids, bool[] results)
func (_Contract *ContractFilterer) ParseVerifyAttestationBatch(log types.Log) (*ContractVerifyAttestationBatch, error) {
	event := new(ContractVerifyAttestationBatch)
	if err := _Contract.contract.UnpackLog(event, "VerifyAttestationBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

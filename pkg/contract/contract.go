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

// CarvProtocolServicedelegateData is an auto generated low-level Go binding around an user-defined struct.
type CarvProtocolServicedelegateData struct {
	From      common.Address
	To        common.Address
	TokenId   *big.Int
	Nonce     *big.Int
	Timestamp *big.Int
}

// CarvProtocolServiceopenVerifierNodeData is an auto generated low-level Go binding around an user-defined struct.
type CarvProtocolServiceopenVerifierNodeData struct {
	Account   common.Address
	TokenId   *big.Int
	Timestamp *big.Int
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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"PackedPtrLen__LenOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PackedPtrLen__PtrOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Slice__OutOfBounds\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token_id\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"OpenVerifierNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ProfixPayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"eip712DomainHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hashStruct\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"RecoverParams\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tee_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"campaign_id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"attestation_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"attestation\",\"type\":\"string\"}],\"name\":\"ReportTeeAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardPayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"identitiesRoot\",\"type\":\"bytes32\"}],\"name\":\"SetIdentitiesRoot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contract_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"campaign_id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"requirements\",\"type\":\"string\"}],\"name\":\"SubmitCampaign\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"carv_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"campaign_id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"campaign_info\",\"type\":\"string\"}],\"name\":\"UserCampaignData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"VerifierWeightChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"attestation_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"name\":\"VerifyAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"attestation_ids\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"results\",\"type\":\"bool[]\"}],\"name\":\"VerifyAttestationBatch\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712_DOMAIN_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FORWARDER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PLATFORM_MINETR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TEE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERIFIER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_admin_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_campaign_id\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_campaign_info\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_carv_id\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_cur_token_id\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addressTeeInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"mrEnclave\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"address_user_map\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"token_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"user_profile_path\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"profile_version\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"address_vote_weight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"attestation_id_list\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"attestation_id_map\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"attestation_id_result_map\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"campain_reward_map\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"campaign_id\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"user_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reward_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total_num\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"contract_address\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"contract_type\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"token_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structCarvProtocolService.openVerifierNodeData\",\"name\":\"verifier_data\",\"type\":\"tuple\"}],\"name\":\"close_verifier_node\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getIdentitiesRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endId\",\"type\":\"uint256\"}],\"name\":\"get_proof_list\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"id_campaign_map\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"campaign_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"campaign_type\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"reward_contract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reward_total_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward_count\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"start_time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end_time\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"requirements\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rewards_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nft_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"carv_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"campaign_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"join_campaign_info\",\"type\":\"string\"}],\"name\":\"join_campaign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nft_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"token_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structCarvProtocolService.openVerifierNodeData\",\"name\":\"verifier_data\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"open_verifier_node\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"campaign_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"attestation\",\"type\":\"string\"}],\"name\":\"report_tee_attestation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"multiIdentitiesRoot\",\"type\":\"bytes32\"}],\"name\":\"setIdentitiesRootFunc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"mrEnclave\",\"type\":\"string\"}],\"name\":\"setTeeInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_address\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"user_profile_path\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"profile_version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"multiIdentitiesRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"set_identities_root\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nft_address\",\"type\":\"address\"}],\"name\":\"set_nft_address\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault_address\",\"type\":\"address\"}],\"name\":\"set_vault_address\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_verifier_pass_threshold\",\"type\":\"uint256\"}],\"name\":\"set_verifier_pass_threshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vrf_address\",\"type\":\"address\"}],\"name\":\"set_vrf_address\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"campaign_id\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"user_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reward_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total_num\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"contract_address\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"contract_type\",\"type\":\"uint8\"}],\"internalType\":\"structCarvProtocolService.reward\",\"name\":\"reward_info\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"campaign_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"campaign_type\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"reward_contract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reward_total_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward_count\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"start_time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end_time\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"requirements\",\"type\":\"string\"}],\"internalType\":\"structCarvProtocolService.campaign\",\"name\":\"campaign_info\",\"type\":\"tuple\"}],\"name\":\"submit_campaign\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"verifier_block\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structCarvProtocolService.delegateData\",\"name\":\"delegate_data\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verifier_delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"verifier_list\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier_pass_threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structCarvProtocolService.delegateData\",\"name\":\"delegate_data\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verifier_redelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structCarvProtocolService.delegateData\",\"name\":\"delegate_data\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verifier_undelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftOwnerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sigAddress\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"userIDs\",\"type\":\"string[]\"},{\"internalType\":\"bytes32\",\"name\":\"multiIdentitiesRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verifyIdentitiesBinding\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"attestation_id\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"name\":\"verify_attestation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"attestation_ids\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"results\",\"type\":\"bool[]\"}],\"name\":\"verify_attestation_batch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vrf_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50615c1980620000216000396000f3fe6080604052600436106102f35760003560e01c806208f8f9146102f857806301ffc9a71461032e57806306fdde031461035e578063079d6a9314610380578063081812fc14610395578063095ea7b3146103c257806309a9c7f0146103e45780630d44ad76146104045780630ded29e9146104465780631794bb3c1461047857806318a14be9146104985780631a514dad146104b85780631a7f9955146104e857806323b872dd14610508578063248a9ca31461052857806327c4a224146105485780632c8ebcde146105685780632f2ff15d1461059557806335f6fb87146105b557806336568abe146105d75780633736c83f146105f75780633ed9c6da1461060a57806342842e0e1461061f57806345692b821461063f5780634679d4901461066c57806347be46e4146106a3578063495eb828146106d057806349d03037146106f25780634f9d5a65146107125780635656a85214610732578063578907ca146107525780635d1fe591146107725780636352211e146107925780636a1aec4a146107b25780636a627842146107d257806370824ab1146107f257806370a0823114610812578063733f23991461083257806385910c251461085257806385e1f4d01461087257806389fb5f4a1461088857806391582484146108b857806391d14854146108d857806395d89b41146108f857806397baf3541461090d5780639bb5bb381461093d5780639ff098ac1461095b578063a217fddf14610971578063a22cb46514610986578063a4bf0813146109a6578063affed0e0146109bc578063b88d4fde146109d2578063ba030180146109f2578063c1f7584814610a12578063c1fb555e14610a32578063c78de70914610a60578063c87b56dd14610a80578063d547741f14610aa0578063d5c6329b14610ac0578063e306f77914610ae0578063e5180e4e14610af6578063e664025214610b16578063e7705db614610b2c578063e97a998714610b60578063e985e9c514610b80578063ff30a11814610ba0578063ffd2f5cc14610bc0575b600080fd5b34801561030457600080fd5b50610318610313366004614685565b610be0565b60405161032591906146a7565b60405180910390f35b34801561033a57600080fd5b5061034e610349366004614701565b610c9d565b6040519015158152602001610325565b34801561036a57600080fd5b50610373610ca8565b604051610325919061476e565b34801561038c57600080fd5b50610373610d3a565b3480156103a157600080fd5b506103b56103b0366004614781565b610dc8565b604051610325919061479a565b3480156103ce57600080fd5b506103e26103dd3660046147c3565b610def565b005b3480156103f057600080fd5b506103e26103ff366004614830565b610f09565b34801561041057600080fd5b506104387fc9c8c5e0314ffb46567ae41ff0c1c5ea54b0c785c8519fcde1714a9ec337f7d981565b604051908152602001610325565b34801561045257600080fd5b5061046661046136600461495e565b610f70565b60405161032596959493929190614992565b34801561048457600080fd5b506103e26104933660046149e0565b61104a565b3480156104a457600080fd5b5060d3546103b5906001600160a01b031681565b3480156104c457600080fd5b5061034e6104d3366004614781565b60df6020526000908152604090205460ff1681565b3480156104f457600080fd5b506103e2610503366004614a21565b611285565b34801561051457600080fd5b506103e26105233660046149e0565b611306565b34801561053457600080fd5b50610438610543366004614781565b611337565b34801561055457600080fd5b506103e2610563366004614ab6565b61134c565b34801561057457600080fd5b50610438610583366004614ab6565b60e36020526000908152604090205481565b3480156105a157600080fd5b506103e26105b0366004614ad3565b61137a565b3480156105c157600080fd5b50610438600080516020615bc483398151915281565b3480156105e357600080fd5b506103e26105f2366004614ad3565b611396565b6103e2610605366004614b03565b611414565b34801561061657600080fd5b506103736114f2565b34801561062b57600080fd5b506103e261063a3660046149e0565b6114ff565b34801561064b57600080fd5b5061043861065a366004614781565b60009081526097602052604090205490565b34801561067857600080fd5b5061068c61068736600461495e565b61151a565b6040516103259b9a99989796959493929190614b6b565b3480156106af57600080fd5b506104386106be366004614ab6565b60e56020526000908152604090205481565b3480156106dc57600080fd5b50610438600080516020615b8483398151915281565b3480156106fe57600080fd5b5061034e61070d366004614c01565b61172c565b34801561071e57600080fd5b5060cc546103b5906001600160a01b031681565b34801561073e57600080fd5b506103b561074d366004614781565b61186b565b34801561075e57600080fd5b506103e261076d366004614d18565b611895565b34801561077e57600080fd5b506103e261078d366004614ab6565b611a0b565b34801561079e57600080fd5b506103b56107ad366004614781565b611a39565b3480156107be57600080fd5b506103e26107cd366004614d3d565b611a6d565b3480156107de57600080fd5b506103e26107ed366004614ab6565b611ac5565b3480156107fe57600080fd5b506103e261080d366004614ab6565b611ba7565b34801561081e57600080fd5b5061043861082d366004614ab6565b611bd5565b34801561083e57600080fd5b5060cb546103b5906001600160a01b031681565b34801561085e57600080fd5b506103e261086d366004614da8565b611c5b565b34801561087e57600080fd5b5061043860e75481565b34801561089457600080fd5b506108a86108a3366004614ab6565b611f01565b6040516103259493929190614dfd565b3480156108c457600080fd5b506103e26108d3366004614781565b61203a565b3480156108e457600080fd5b5061034e6108f3366004614ad3565b61204b565b34801561090457600080fd5b50610373612076565b34801561091957600080fd5b5061034e610928366004614781565b60d16020526000908152604090205460ff1681565b34801561094957600080fd5b506103e2610958366004614e4b565b50565b34801561096757600080fd5b5061043860cd5481565b34801561097d57600080fd5b50610438600081565b34801561099257600080fd5b506103e26109a1366004614e67565b612085565b3480156109b257600080fd5b5061043860d85481565b3480156109c857600080fd5b5061043860d95481565b3480156109de57600080fd5b506103e26109ed366004614e95565b612090565b3480156109fe57600080fd5b506103e2610a0d366004614da8565b6120c2565b348015610a1e57600080fd5b506103e2610a2d366004614f44565b6122d4565b348015610a3e57600080fd5b50610a52610a4d366004614ab6565b6124a0565b604051610325929190614fa3565b348015610a6c57600080fd5b5060d6546103b5906001600160a01b031681565b348015610a8c57600080fd5b50610373610a9b366004614781565b6125cc565b348015610aac57600080fd5b506103e2610abb366004614ad3565b612640565b348015610acc57600080fd5b50610438610adb366004614781565b61265c565b348015610aec57600080fd5b5061043860e85481565b348015610b0257600080fd5b506103e2610b11366004614da8565b61267d565b348015610b2257600080fd5b5061043860d75481565b348015610b3857600080fd5b506104387f0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea0981565b348015610b6c57600080fd5b506103e2610b7b366004614d3d565b612894565b348015610b8c57600080fd5b5061034e610b9b366004614fd1565b612994565b348015610bac57600080fd5b506103e2610bbb366004614685565b6129c2565b348015610bcc57600080fd5b506103e2610bdb366004614fff565b612a0a565b60606000610bee848461504d565b6001600160401b03811115610c0557610c056148a9565b604051908082528060200260200182016040528015610c2e578160200160208202803683370190505b509050835b83811015610c935760d08181548110610c4e57610c4e615060565b9060005260206000200154828683610c66919061504d565b81518110610c7657610c76615060565b602090810291909101015280610c8b81615076565b915050610c33565b5090505b92915050565b6000610c9782612ba2565b606060658054610cb79061508f565b80601f0160208091040260200160405190810160405280929190818152602001828054610ce39061508f565b8015610d305780601f10610d0557610100808354040283529160200191610d30565b820191906000526020600020905b815481529060010190602001808311610d1357829003601f168201915b5050505050905090565b60ce8054610d479061508f565b80601f0160208091040260200160405190810160405280929190818152602001828054610d739061508f565b8015610dc05780601f10610d9557610100808354040283529160200191610dc0565b820191906000526020600020905b815481529060010190602001808311610da357829003601f168201915b505050505081565b6000610dd382612bc7565b506000908152606960205260409020546001600160a01b031690565b6000610dfa82611a39565b9050806001600160a01b0316836001600160a01b031603610e6c5760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b60648201526084015b60405180910390fd5b336001600160a01b0382161480610e885750610e888133612994565b610efa5760405162461bcd60e51b815260206004820152603d60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206f7220617070726f76656420666f7220616c6c0000006064820152608401610e63565b610f048383612bec565b505050565b60cd85905560ce610f1b84868361511e565b5060cf610f2982848361511e565b507f83eda6edb918aac3f35c8d616a6198dbb75eeca3248e0efec05e66a30d51119160cd5460ce60cf604051610f6193929190615254565b60405180910390a15050505050565b805160208183018101805160da82529282019190930120915280548190610f969061508f565b80601f0160208091040260200160405190810160405280929190818152602001828054610fc29061508f565b801561100f5780601f10610fe45761010080835404028352916020019161100f565b820191906000526020600020905b815481529060010190602001808311610ff257829003601f168201915b5050505060018301546002840154600385015460049095015493946001600160a01b039283169491935091811690600160a01b900460ff1686565b600054610100900460ff161580801561106a5750600054600160ff909116105b8061108b575061107930612c5a565b15801561108b575060005460ff166001145b6110ee5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610e63565b6000805460ff191660011790558015611111576000805461ff0019166101001790555b60cc8054336001600160a01b03199182161790915560cb805482166001600160a01b038781169190911790915560d68054909216908516179055600160d881905560e783905560408051808201825260138152724361727650726f746f636f6c5365727669636560681b60209182015281518083018352928352603160f81b928101929092525161120a917fc2f8787176b8ac6bf7215b4adcc1e069bf4ab82d9ab1df05a57a91d425935b6e917ffacd796373658d15e1306676417a685404c80da96526d405dcce52fda8803ac9917fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc69187910161527f565b60408051601f19818403018152919052805160209091012060e85561122d612c69565b611238600033612cd6565b801561127f576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15b50505050565b6001600160a01b038716600090815260dd602052604090206001016112ab86888361511e565b506001600160a01b038716600090815260dd60205260409020600281018590556003016112d982848361511e565b506001600160a01b038716600090815260dd60205260409020546112fd90846129c2565b50505050505050565b6113103382612d5c565b61132c5760405162461bcd60e51b8152600401610e639061529a565b610f04838383612dbb565b60009081526099602052604090206001015490565b600061135781612f1f565b5060d680546001600160a01b0319166001600160a01b0392909216919091179055565b61138382611337565b61138c81612f1f565b610f048383612cd6565b6001600160a01b03811633146114065760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b6064820152608401610e63565b6114108282612f29565b5050565b61141e8233612f90565b8060db61142b82806152e7565b60405161143992919061532d565b9081526040519081900360200190206114528282615386565b5082905060da61146282806152e7565b60405161147092919061532d565b9081526040519081900360200190206114898282615483565b507fe8e692af5f416c4637c9307d7f08b7697fd8f4398f55103ce03ddd9388ccd84d90506114bd60a0840160808501614ab6565b6114c783806152e7565b6114d56101408601866152e7565b6040516114e69594939291906155bb565b60405180910390a15050565b60cf8054610d479061508f565b610f0483838360405180602001604052806000815250612090565b805160208183018101805160db825292820191909301209152805481906115409061508f565b80601f016020809104026020016040519081016040528092919081815260200182805461156c9061508f565b80156115b95780601f1061158e576101008083540402835291602001916115b9565b820191906000526020600020905b81548152906001019060200180831161159c57829003601f168201915b5050505050908060010180546115ce9061508f565b80601f01602080910402602001604051908101604052809291908181526020018280546115fa9061508f565b80156116475780601f1061161c57610100808354040283529160200191611647565b820191906000526020600020905b81548152906001019060200180831161162a57829003601f168201915b50505060028401546003850154600486015460058701546006880154600789015460088a015460098b0180549a9b6001600160a01b03808a169c60ff600160a01b909b048b169c5098169950959794969390941694919390926116a99061508f565b80601f01602080910402602001604051908101604052809291908181526020018280546116d59061508f565b80156117225780601f106116f757610100808354040283529160200191611722565b820191906000526020600020905b81548152906001019060200180831161170557829003601f168201915b505050505090508b565b6000866001600160a01b031661174189611a39565b6001600160a01b0316146117925760405162461bcd60e51b81526020600482015260186024820152771b999d081bdddb995c881a5cc81b9bdd0818dbdc9c9958dd60421b6044820152606401610e63565b8451806117da5760405162461bcd60e51b81526020600482015260166024820152757573657249442063616e6e6f7420626520656d70747960501b6044820152606401610e63565b60005b81811015611819576118078782815181106117fa576117fa615060565b60200260200101516130ef565b8061181181615076565b9150506117dd565b50600061185d888787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061335b92505050565b9a9950505050505050505050565b60d2818154811061187b57600080fd5b6000918252602090912001546001600160a01b0316905081565b33600090815260e560205260409020546118c15760405162461bcd60e51b8152600401610e63906155ff565b600082815260d1602052604090205460ff166118ef5760405162461bcd60e51b8152600401610e6390615642565b6118f8826133bc565b156119515760405162461bcd60e51b815260206004820152602360248201527f6174746573746174696f6e2063616e206e6f742062792076657269667920616760448201526230b4b760e91b6064820152608401610e63565b600082815260e0602090815260408220805460018101825590835291200180546001600160a01b0319163317905561198882613489565b80156119915750805b156119ba5761199f336134b3565b600082815260df60205260409020805460ff19168215151790555b33600081815260e3602090815260409182902043905581519283528201849052821515908201527f937b4d1b5b461852fdbbb174b68adefb42155e668c05a1261fae11a2d22a240a906060016114e6565b6000611a1681612f1f565b5060d380546001600160a01b0319166001600160a01b0392909216919091179055565b600080611a45836135d2565b90506001600160a01b038116610c975760405162461bcd60e51b8152600401610e6390615674565b600080516020615b84833981519152611a8581612f1f565b33600090815260e260205260409020611a9f85878361511e565b5033600090815260e260205260409020600101611abd83858361511e565b505050505050565b60d65460d8546040516340c10f1960e01b81526001600160a01b03909216916340c10f1991611af9918591906004016156a6565b600060405180830381600087803b158015611b1357600080fd5b505af1158015611b27573d6000803e3d6000fd5b505050506001600160a01b038116600090815260e560205260408120805491611b4f83615076565b91905055507f30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe8160d854604051611b879291906156a6565b60405180910390a160d88054906000611b9f83615076565b919050555050565b6000611bb281612f1f565b5060cb80546001600160a01b0319166001600160a01b0392909216919091179055565b60006001600160a01b038216611c3f5760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b6064820152608401610e63565b506001600160a01b031660009081526068602052604090205490565b600080516020615bc4833981519152611c7381612f1f565b611c806020840184614ab6565b60d654604080516331a9108f60e11b81529086013560048201526001600160a01b039283169290911690636352211e90602401602060405180830381865afa158015611cd0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cf491906156bf565b6001600160a01b031614611d1a5760405162461bcd60e51b8152600401610e63906156dc565b60d954836060013511611d3f5760405162461bcd60e51b8152600401610e6390615725565b6000611d4b84846135ed565b9050611d5a6020850185614ab6565b6001600160a01b0316816001600160a01b031614611d8a5760405162461bcd60e51b8152600401610e639061576a565b600060e481611d9c6020880188614ab6565b6001600160a01b039081168252602080830193909352604091820160009081208984013582529093529120541614611e2a5760405162461bcd60e51b815260206004820152602b60248201527f4361727650726f746f636f6c536572766963653a20616c72656164792062656560448201526a1b8819195b1959d85d195960aa1b6064820152608401610e63565b611e4f611e3a6020860186614ab6565b611e4a6040870160208801614ab6565b6135fc565b611e5f6040850160208601614ab6565b60e46000611e706020880188614ab6565b6001600160a01b0390811682526020808301939093526040918201600090812089840135825290935290822080546001600160a01b031916939091169290921790915560d9805491611ec183615076565b90915550600080516020615ba48339815191529050611ee36020860186614ab6565b611ef36040870160208801614ab6565b6040516112769291906157ba565b60dd6020526000908152604090208054600182018054919291611f239061508f565b80601f0160208091040260200160405190810160405280929190818152602001828054611f4f9061508f565b8015611f9c5780601f10611f7157610100808354040283529160200191611f9c565b820191906000526020600020905b815481529060010190602001808311611f7f57829003601f168201915b505050505090806002015490806003018054611fb79061508f565b80601f0160208091040260200160405190810160405280929190818152602001828054611fe39061508f565b80156120305780601f1061200557610100808354040283529160200191612030565b820191906000526020600020905b81548152906001019060200180831161201357829003601f168201915b5050505050905084565b600061204581612f1f565b5060d755565b60009182526099602090815260408084206001600160a01b0393909316845291905290205460ff1690565b606060668054610cb79061508f565b6114103383836136d0565b61209a3383612d5c565b6120b65760405162461bcd60e51b8152600401610e639061529a565b61127f8484848461379a565b600080516020615bc48339815191526120da81612f1f565b6120e76020840184614ab6565b60d654604080516331a9108f60e11b81529086013560048201526001600160a01b039283169290911690636352211e90602401602060405180830381865afa158015612137573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061215b91906156bf565b6001600160a01b0316146121815760405162461bcd60e51b8152600401610e63906156dc565b60d9548360600135116121a65760405162461bcd60e51b8152600401610e6390615725565b60006121b284846135ed565b90506121c16020850185614ab6565b6001600160a01b0316816001600160a01b0316146121f15760405162461bcd60e51b8152600401610e63906157d4565b600060e4816122036020880188614ab6565b6001600160a01b0390811682526020808301939093526040918201600090812089840135825284529190912054169150612245908290611e4a90880188614ab6565b600060e4816122576020890189614ab6565b6001600160a01b039081168252602080830193909352604091820160009081208a840135825284529190912080546001600160a01b0319169390911692909217909155600080516020615ba4833981519152906122b690870187614ab6565b6122c66040880160208901614ab6565b604051610f619291906157ba565b33600090815260e560205260409020546123005760405162461bcd60e51b8152600401610e63906155ff565b8260005b818110156124575760d1600087878481811061232257612322615060565b602090810292909201358352508101919091526040016000205460ff1661235b5760405162461bcd60e51b8152600401610e6390615642565b60e0600087878481811061237157612371615060565b6020908102929092013583525081810192909252604001600090812080546001810182559082529190200180546001600160a01b031916331790556123cd8686838181106123c1576123c1615060565b90506020020135613489565b15612445576123db336134b3565b8383828181106123ed576123ed615060565b90506020020160208101906124029190615809565b60df600088888581811061241857612418615060565b90506020020135815260200190815260200160002060006101000a81548160ff0219169083151502179055505b8061244f81615076565b915050612304565b5033600081815260e3602052604090819020439055517f04798cc2c828443d2d9ef127deaa1f73e782d8a9fffbc9db2bcca53e12765af591610f61918890889088908890615826565b60e2602052600090815260409020805481906124bb9061508f565b80601f01602080910402602001604051908101604052809291908181526020018280546124e79061508f565b80156125345780601f1061250957610100808354040283529160200191612534565b820191906000526020600020905b81548152906001019060200180831161251757829003601f168201915b5050505050908060010180546125499061508f565b80601f01602080910402602001604051908101604052809291908181526020018280546125759061508f565b80156125c25780601f10612597576101008083540402835291602001916125c2565b820191906000526020600020905b8154815290600101906020018083116125a557829003601f168201915b5050505050905082565b60606125d782612bc7565b60006125ee60408051602081019091526000815290565b9050600081511161260e5760405180602001604052806000815250612639565b80612618846137cd565b6040516020016126299291906158aa565b6040516020818303038152906040525b9392505050565b61264982611337565b61265281612f1f565b610f048383612f29565b60d0818154811061266c57600080fd5b600091825260209091200154905081565b600080516020615bc483398151915261269581612f1f565b6126a26020840184614ab6565b60d654604080516331a9108f60e11b81529086013560048201526001600160a01b039283169290911690636352211e90602401602060405180830381865afa1580156126f2573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061271691906156bf565b6001600160a01b03161461273c5760405162461bcd60e51b8152600401610e63906156dc565b60d9548360600135116127615760405162461bcd60e51b8152600401610e6390615725565b600061276d84846135ed565b905061277c6020850185614ab6565b6001600160a01b0316816001600160a01b0316146127ac5760405162461bcd60e51b8152600401610e639061576a565b600060e4816127be6020880188614ab6565b6001600160a01b039081168252602080830193909352604091820160009081208984018035835290855292902054169250612800918391611e4a918901614ab6565b6128106040860160208701614ab6565b60e460006128216020890189614ab6565b6001600160a01b039081168252602080830193909352604091820160009081208a840135825290935290822080546001600160a01b031916939091169290921790915560d980549161287283615076565b90915550600080516020615ba483398151915290506122b66020870187614ab6565b600080516020615b848339815191526128ac81612f1f565b600083836040516128be92919061532d565b6040805191829003822060d0805460018181019092557fe89d44c8fd6a9bac8af33ce47f56337617d449bf7ff3956b618c646de829cbcb01829055600082815260d1602052928320805460ff1916909117905592509060da90612924908990899061532d565b90815260408051918290036020908101832060020154600086815260e190925291902081905591507f99a038e9d345d0b12130b3b1fb003bf8f2d3a5c27ce2a800bbb1608efff6c591906129839033908a908a9087908b908b906158d9565b60405180910390a150505050505050565b6001600160a01b039182166000908152606a6020908152604080832093909416825291909152205460ff1690565b60008281526097602090815260409182902083905581518481529081018390527f21b5a040ca6f61b8bcacb8f423e25ce46e88932f887ce92e60343369c20ec06f91016114e6565b6000612a16838361385f565b60d6546040516331a9108f60e11b8152602086013560048201529192506000916001600160a01b0390911690636352211e90602401602060405180830381865afa158015612a68573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a8c91906156bf565b9050612a9b6020850185614ab6565b6001600160a01b0316826001600160a01b031614612acb5760405162461bcd60e51b8152600401610e63906157d4565b806001600160a01b0316826001600160a01b031614612b2a5760405162461bcd60e51b815260206004820152601b60248201527a1d995c9a599e481d1a19481b999d081bdddb995c8819985a5b1959602a1b6044820152606401610e63565b60208085018035600090815260e99092526040909120805460ff191660011790557f832607f249995a2ed86c13bde705f2d0576d5eee1780a83084d8cee8992641c490612b779086614ab6565b604080516001600160a01b039092168252602080880135908301528087013590820152606001611276565b60006001600160e01b03198216637965db0b60e01b1480610c975750610c97826138e3565b612bd081613908565b6109585760405162461bcd60e51b8152600401610e6390615674565b600081815260696020526040902080546001600160a01b0319166001600160a01b0384169081179091558190612c2182611a39565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6001600160a01b03163b151590565b600054610100900460ff16612cd45760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608401610e63565b565b612ce0828261204b565b6114105760008281526099602090815260408083206001600160a01b03851684529091529020805460ff19166001179055612d183390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b600080612d6883611a39565b9050806001600160a01b0316846001600160a01b03161480612d8f5750612d8f8185612994565b80612db35750836001600160a01b0316612da884610dc8565b6001600160a01b0316145b949350505050565b826001600160a01b0316612dce82611a39565b6001600160a01b031614612df45760405162461bcd60e51b8152600401610e6390615924565b6001600160a01b038216612e565760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b6064820152608401610e63565b826001600160a01b0316612e6982611a39565b6001600160a01b031614612e8f5760405162461bcd60e51b8152600401610e6390615924565b600081815260696020908152604080832080546001600160a01b03199081169091556001600160a01b0387811680865260688552838620805460001901905590871680865283862080546001019055868652606790945282852080549092168417909155905184937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6109588133613925565b612f33828261204b565b156114105760008281526099602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b612fa060c0830160a08401615969565b60ff16600003612fed5760cc54604080516001600160a01b03909216919084013580156108fc02916000818181858888f19350505050158015612fe7573d6000803e3d6000fd5b5061309a565b612ffd60c0830160a08401615969565b60ff1660010361309a5761301760a0830160808401614ab6565b60cc54604080516323b872dd60e01b81526001600160a01b03858116600483015292831660248201529085013560448201529116906323b872dd906064016020604051808303816000875af1158015613074573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130989190615986565b505b7f704db66c978a7472842ddae0d478a8a7d123fb6ad449e4426b5e0f6a22081de76130cb60a0840160808501614ab6565b60cc54604080516114e6939286926001600160a01b039091169190880135906159a3565b600081511161313a5760405162461bcd60e51b81526020600482015260176024820152767573657249442063616e206e6f7420626520656d70747960481b6044820152606401610e63565b6040805180820190915260018152601d60f91b60208201526000808061315f8561397e565b905061317461316d8561397e565b8290613989565b91945092509050826131d85760405162461bcd60e51b815260206004820152602760248201527f74686520666972737420706172742064656c696d6974657220646f6573206e6f6044820152661d08195e1a5cdd60ca1b6064820152608401610e63565b60006131e3836139d8565b51116132315760405162461bcd60e51b815260206004820152601d60248201527f746865206669727374207061727420646f6573206e6f742065786973740000006044820152606401610e63565b61323d61316d8561397e565b91945092509050826132a25760405162461bcd60e51b815260206004820152602860248201527f746865207365636f6e6420706172742064656c696d6974657220646f6573206e6044820152671bdd08195e1a5cdd60c21b6064820152608401610e63565b60006132ad836139d8565b51116132fb5760405162461bcd60e51b815260206004820152601e60248201527f746865207365636f6e64207061727420646f6573206e6f7420657869737400006044820152606401610e63565b613304816139d8565b516040146133545760405162461bcd60e51b815260206004820152601d60248201527f69642068617368206c656e677468206973206e6f7420636f72726563740000006044820152606401610e63565b5050505050565b600080600061336a85856139e3565b90925090506000816004811115613383576133836159cd565b1480156133a15750856001600160a01b0316826001600160a01b0316145b806133b257506133b2868686613a28565b9695505050505050565b600081815260e0602090815260408083208054825181850281018501909352808352849383018282801561341957602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116133fb575b505083519394506000925050505b8181101561347e57336001600160a01b031683828151811061344b5761344b615060565b60200260200101516001600160a01b03160361346c57506001949350505050565b8061347681615076565b915050613427565b506000949350505050565b600081815260e0602052604081205460d75481106134aa5750600192915050565b50600092915050565b60cb5460405163732f496f60e01b81526000916001600160a01b03169063732f496f906134e490309060040161479a565b6020604051808303816000875af1158015613503573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061352791906159e3565b60cb546040516324e2624160e01b81529192506001600160a01b0316906324e262419061355890859060040161479a565b600060405180830381600087803b15801561357257600080fd5b505af1158015613586573d6000803e3d6000fd5b505060cb5460cc546040517f3832ba0f9bc9de56bd1f974e665b53f2643ed48a47b45c3d36976bc1ea31459994506114e693506001600160a01b039283169290911690339086906159a3565b6000908152606760205260409020546001600160a01b031690565b60006126396020840184614ab6565b6001600160a01b038216600090815260e5602052604090205461367a5760405162461bcd60e51b815260206004820152603060248201527f4361727650726f746f636f6c536572766963653a2066726f6d20766f7465207760448201526f195a59da1d081a5cc81a5b9d985b1a5960821b6064820152608401610e63565b6001600160a01b038216600090815260e56020526040812080549161369e836159fc565b90915550506001600160a01b038116600090815260e5602052604081208054916136c783615076565b91905055505050565b816001600160a01b0316836001600160a01b03160361372d5760405162461bcd60e51b815260206004820152601960248201527822a9219b99189d1030b8383937bb32903a379031b0b63632b960391b6044820152606401610e63565b6001600160a01b038381166000818152606a6020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6137a5848484612dbb565b6137b184848484613b14565b61127f5760405162461bcd60e51b8152600401610e6390615a13565b606060006137da83613c1c565b60010190506000816001600160401b038111156137f9576137f96148a9565b6040519080825280601f01601f191660200182016040528015613823576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a850494508461382d57509392505050565b6000807f86045fb4d5a003fb9c37d1c87ce076fd4fc05a38a81da79d19f6bc700f66dc4f6138906020860186614ab6565b604080516020818101949094526001600160a01b039092168282015291860135606082015290850135608082015260a001604051602081830303815290604052805190602001209050612db38184613cf2565b60006001600160e01b0319821663f389baad60e01b1480610c975750610c9782613d7d565b600080613914836135d2565b6001600160a01b0316141592915050565b61392f828261204b565b6114105761393c81613dcd565b613947836020613ddf565b604051602001613958929190615a65565b60408051601f198184030181529082905262461bcd60e51b8252610e639160040161476e565b6000610c9782613f7a565b60008080806139988686613f8e565b905060001981036139b4576000866000935093509350506139d1565b6139c9816139c1876140a5565b8891906140b1565b935093509350505b9250925092565b6060610c9782614120565b6000808251604103613a195760208301516040840151606085015160001a613a0d87828585614195565b94509450505050613a21565b506000905060025b9250929050565b6000806000856001600160a01b0316631626ba7e60e01b8686604051602401613a52929190615ad4565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051613a909190615aed565b600060405180830381855afa9150503d8060008114613acb576040519150601f19603f3d011682016040523d82523d6000602084013e613ad0565b606091505b5091509150818015613ae457506020815110155b80156133b257508051630b135d3f60e11b90613b0990830160209081019084016159e3565b149695505050505050565b6000613b28846001600160a01b0316612c5a565b15613c1157604051630a85bd0160e11b81526001600160a01b0385169063150b7a0290613b5f903390899088908890600401615b09565b6020604051808303816000875af1925050508015613b9a575060408051601f3d908101601f19168201909252613b9791810190615b3c565b60015b613bf7573d808015613bc8576040519150601f19603f3d011682016040523d82523d6000602084013e613bcd565b606091505b508051600003613bef5760405162461bcd60e51b8152600401610e6390615a13565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050612db3565b506001949350505050565b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b8310613c5b5772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6904ee2d6d415b85acef8160201b8310613c85576904ee2d6d415b85acef8160201b830492506020015b662386f26fc100008310613ca357662386f26fc10000830492506010015b6305f5e1008310613cbb576305f5e100830492506008015b6127108310613ccf57612710830492506004015b60648310613ce1576064830492506002015b600a8310610c975760010192915050565b60e85460405161190160f01b602082015260228101919091526042810183905260009081906062016040516020818303038152906040528051906020012090507fa50615fb8919bf8c8a6350a26c7921f53c0434b16d11c80df59de6b281a83bf160e75460e8548684604051613d6b949392919061527f565b60405180910390a1612db3818461424f565b60006001600160e01b031982166380ac58cd60e01b1480613dae57506001600160e01b03198216635b5e139f60e01b145b80610c9757506301ffc9a760e01b6001600160e01b0319831614610c97565b6060610c976001600160a01b03831660145b60606000613dee836002615b59565b613df9906002615b70565b6001600160401b03811115613e1057613e106148a9565b6040519080825280601f01601f191660200182016040528015613e3a576020820181803683370190505b509050600360fc1b81600081518110613e5557613e55615060565b60200101906001600160f81b031916908160001a905350600f60fb1b81600181518110613e8457613e84615060565b60200101906001600160f81b031916908160001a9053506000613ea8846002615b59565b613eb3906001615b70565b90505b6001811115613f2b576f181899199a1a9b1b9c1cb0b131b232b360811b85600f1660108110613ee757613ee7615060565b1a60f81b828281518110613efd57613efd615060565b60200101906001600160f81b031916908160001a90535060049490941c93613f24816159fc565b9050613eb6565b5083156126395760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401610e63565b60008060208301905061263981845161426b565b600080613f9a846140a5565b90506000613fa7846140a5565b905080600003613fbc57600092505050610c97565b811580613fc857508181115b15613fd95760001992505050610c97565b6000613fe58660801c90565b90506000613ff38660801c90565b90506000614002825160001a90565b90505b6000614012848784614277565b9050600019810361402d576000199650505050505050610c97565b9485900394928301928585111561404e576000199650505050505050610c97565b8483208585200361407a576140638960801c90565b61406d908561504d565b9650505050505050610c97565b85600103614092576000199650505050505050610c97565b6000199095019460019093019250614005565b6001600160801b031690565b6000806000806140c18760801c90565b905060006140ce886140a5565b9050868601818111156140f4576040516365f4e9df60e01b815260040160405180910390fd5b6140fe8389614337565b945061410e818401828403614337565b93506001955050505093509350939050565b606061412b826140a5565b6001600160401b03811115614142576141426148a9565b6040519080825280601f01601f19166020018201604052801561416c576020820181803683370190505b5090506020810161418f816141818560801c90565b61418a866140a5565b61434a565b50919050565b6000806fa2a8918ca85bafe22016d0b997e4df60600160ff1b038311156141c25750600090506003614246565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015614216573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661423f57600060019250925050614246565b9150600090505b94509492505050565b600080600061425e85856139e3565b91509150610c9381614358565b6000612639838361449d565b6000602083116142935761428c8484846144a9565b9050612639565b83601f8416808503820160006142a8866145ac565b90505b818810156142ed57875181186142c0816145d4565b156142e1578489036142d48a60208a6144a9565b0195505050505050612639565b602089019850506142ab565b8260000361430357600019945050505050612639565b61430e8284886144a9565b9450600019850361432757600019945050505050612639565b8382038501945050505050612639565b6001600160801b031660809190911b1790565b8083828460045afa50505050565b600081600481111561436c5761436c6159cd565b036143745750565b6001816004811115614388576143886159cd565b036143d05760405162461bcd60e51b815260206004820152601860248201527745434453413a20696e76616c6964207369676e617475726560401b6044820152606401610e63565b60028160048111156143e4576143e46159cd565b036144315760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610e63565b6003816004811115614445576144456159cd565b036109585760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610e63565b6000612639838361461f565b82516000908160208511156144bd57602094505b601285106145405760006144d0856145ac565b831890506144e66001600160801b0382176145d4565b60000361451c5760109150601a86106145175761450b6001600160401b0382176145d4565b60000361451757601891505b61453a565b61452e6001600160c01b0382176145d4565b60000361453a57600891505b50614577565b600a8510614577576000614553856145ac565b831890506145696001600160c01b0382176145d4565b60000361457557600891505b505b8481101561459f5781811a60ff85168103614596575091506126399050565b50600101614577565b5060001995945050505050565b60ff167f01010101010101010101010101010101010101010101010101010101010101010290565b7ffefefefefefefefefefefefefefefefefefefefefefefefefefefefefefefeff81019019167f80808080808080808080808080808080808080808080808080808080808080801690565b60006001600160801b038311156146495760405163fee7506f60e01b815260040160405180910390fd5b6001600160801b0382111561467157604051633b6b098d60e01b815260040160405180910390fd5b506001600160801b031660809190911b1790565b6000806040838503121561469857600080fd5b50508035926020909101359150565b6020808252825182820181905260009190848201906040850190845b818110156146df578351835292840192918401916001016146c3565b50909695505050505050565b6001600160e01b03198116811461095857600080fd5b60006020828403121561471357600080fd5b8135612639816146eb565b60005b83811015614739578181015183820152602001614721565b50506000910152565b6000815180845261475a81602086016020860161471e565b601f01601f19169290920160200192915050565b6020815260006126396020830184614742565b60006020828403121561479357600080fd5b5035919050565b6001600160a01b0391909116815260200190565b6001600160a01b038116811461095857600080fd5b600080604083850312156147d657600080fd5b82356147e1816147ae565b946020939093013593505050565b60008083601f84011261480157600080fd5b5081356001600160401b0381111561481857600080fd5b602083019150836020828501011115613a2157600080fd5b60008060008060006060868803121561484857600080fd5b8535945060208601356001600160401b038082111561486657600080fd5b61487289838a016147ef565b9096509450604088013591508082111561488b57600080fd5b50614898888289016147ef565b969995985093965092949392505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b03811182821017156148e7576148e76148a9565b604052919050565b600082601f83011261490057600080fd5b81356001600160401b03811115614919576149196148a9565b61492c601f8201601f19166020016148bf565b81815284602083860101111561494157600080fd5b816020850160208301376000918101602001919091529392505050565b60006020828403121561497057600080fd5b81356001600160401b0381111561498657600080fd5b612db3848285016148ef565b60c0815260006149a560c0830189614742565b6001600160a01b03978816602084015260408301969096525060608101939093529316608082015260ff90921660a090920191909152919050565b6000806000606084860312156149f557600080fd5b8335614a00816147ae565b92506020840135614a10816147ae565b929592945050506040919091013590565b600080600080600080600060a0888a031215614a3c57600080fd5b8735614a47816147ae565b965060208801356001600160401b0380821115614a6357600080fd5b614a6f8b838c016147ef565b909850965060408a0135955060608a0135945060808a0135915080821115614a9657600080fd5b50614aa38a828b016147ef565b989b979a50959850939692959293505050565b600060208284031215614ac857600080fd5b8135612639816147ae565b60008060408385031215614ae657600080fd5b823591506020830135614af8816147ae565b809150509250929050565b60008060408385031215614b1657600080fd5b82356001600160401b0380821115614b2d57600080fd5b9084019060c08287031215614b4157600080fd5b90925060208401359080821115614b5757600080fd5b5083016101608186031215614af857600080fd5b6000610160808352614b7f8184018f614742565b90508281036020840152614b93818e614742565b6001600160a01b038d8116604086015260ff8d81166060870152908c16608086015260a085018b905260c085018a9052881660e0850152610100840187905261012084018690528381036101408501529050614bef8185614742565b9e9d5050505050505050505050505050565b600080600080600080600060c0888a031215614c1c57600080fd5b873596506020880135614c2e816147ae565b95506040880135614c3e816147ae565b945060608801356001600160401b0380821115614c5a57600080fd5b818a0191508a601f830112614c6e57600080fd5b813581811115614c8057614c806148a9565b8060051b614c90602082016148bf565b9182526020818501810192908101908e841115614cac57600080fd5b6020860192505b83831015614cea578483351115614cc957600080fd5b614cd98f602085358901016148ef565b825260209283019290910190614cb3565b985050505060808a0135945060a08a0135915080821115614a9657600080fd5b801515811461095857600080fd5b60008060408385031215614d2b57600080fd5b823591506020830135614af881614d0a565b60008060008060408587031215614d5357600080fd5b84356001600160401b0380821115614d6a57600080fd5b614d76888389016147ef565b90965094506020870135915080821115614d8f57600080fd5b50614d9c878288016147ef565b95989497509550505050565b60008082840360c0811215614dbc57600080fd5b60a0811215614dca57600080fd5b5082915060a08201356001600160401b03811115614de757600080fd5b614df3858286016148ef565b9150509250929050565b848152608060208201526000614e166080830186614742565b8460408401528281036060840152614e2e8185614742565b979650505050505050565b60006060828403121561418f57600080fd5b600060608284031215614e5d57600080fd5b6126398383614e39565b60008060408385031215614e7a57600080fd5b8235614e85816147ae565b91506020830135614af881614d0a565b60008060008060808587031215614eab57600080fd5b8435614eb6816147ae565b93506020850135614ec6816147ae565b92506040850135915060608501356001600160401b03811115614ee857600080fd5b614ef4878288016148ef565b91505092959194509250565b60008083601f840112614f1257600080fd5b5081356001600160401b03811115614f2957600080fd5b6020830191508360208260051b8501011115613a2157600080fd5b60008060008060408587031215614f5a57600080fd5b84356001600160401b0380821115614f7157600080fd5b614f7d88838901614f00565b90965094506020870135915080821115614f9657600080fd5b50614d9c87828801614f00565b604081526000614fb66040830185614742565b8281036020840152614fc88185614742565b95945050505050565b60008060408385031215614fe457600080fd5b8235614fef816147ae565b91506020830135614af8816147ae565b6000806080838503121561501257600080fd5b61501c8484614e39565b915060608301356001600160401b03811115614de757600080fd5b634e487b7160e01b600052601160045260246000fd5b81810381811115610c9757610c97615037565b634e487b7160e01b600052603260045260246000fd5b60006001820161508857615088615037565b5060010190565b600181811c908216806150a357607f821691505b60208210810361418f57634e487b7160e01b600052602260045260246000fd5b601f821115610f0457600081815260208120601f850160051c810160208610156150ea5750805b601f850160051c820191505b81811015611abd578281556001016150f6565b600019600383901b1c191660019190911b1790565b6001600160401b03831115615135576151356148a9565b61514983615143835461508f565b836150c3565b6000601f84116001811461517757600085156151655750838201355b61516f8682615109565b845550613354565b600083815260209020601f19861690835b828110156151a85786850135825560209485019460019092019101615188565b50868210156151c55760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b600081546151e48161508f565b808552602060018381168015615201576001811461521b57615249565b60ff1985168884015283151560051b880183019550615249565b866000528260002060005b858110156152415781548a8201860152908301908401615226565b890184019650505b505050505092915050565b83815260606020820152600061526d60608301856151d7565b82810360408401526133b281856151d7565b93845260208401929092526040830152606082015260800190565b6020808252602d908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526c1c881bdc88185c1c1c9bdd9959609a1b606082015260800190565b6000808335601e198436030181126152fe57600080fd5b8301803591506001600160401b0382111561531857600080fd5b602001915036819003821315613a2157600080fd5b8183823760009101908152919050565b60008135610c97816147ae565b80546001600160a01b0319166001600160a01b0392909216919091179055565b60ff8116811461095857600080fd5b60008135610c978161536a565b61539082836152e7565b61539b81838561511e565b50506153aa60208301836152e7565b6153b881836001860161511e565b5050600281016153d36153cd6040850161533d565b8261534a565b6154006153e260608501615379565b82805460ff60a01b191660a09290921b60ff60a01b16919091179055565b506154196154106080840161533d565b6003830161534a565b60a0820135600482015560c0820135600582015561545161543c60e08401615379565b6006830160ff821660ff198254161781555050565b610100820135600782015561012082013560088201556154756101408301836152e7565b61127f81836009860161511e565b61548d82836152e7565b6001600160401b038111156154a4576154a46148a9565b6154b8816154b2855461508f565b856150c3565b6000601f8211600181146154e657600083156154d45750838201355b6154de8482615109565b865550615540565b600085815260209020601f19841690835b8281101561551757868501358255602094850194600190920191016154f7565b50848210156155345760001960f88660031b161c19848701351681555b505060018360011b0185555b5050505061555c6155536020840161533d565b6001830161534a565b6040820135600282015560608201356003820155600481016155836153cd6080850161533d565b610f046153e260a08501615379565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b03861681526060602082018190526000906155e09083018688615592565b82810360408401526155f3818587615592565b98975050505050505050565b60208082526023908201527f4361727650726f746f636f6c536572766963653a206e6f20766f74652077656960408201526219da1d60ea1b606082015260800190565b602080825260189082015277185d1d195cdd185d1a5bdb881a5cc81b9bdd08195e1a5cdd60421b604082015260600190565b602080825260189082015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b604082015260600190565b6001600160a01b03929092168252602082015260400190565b6000602082840312156156d157600080fd5b8151612639816147ae565b60208082526029908201527f4361727650726f746f636f6c536572766963653a206e6674206f776e657220696040820152681cc81a5b9d985b1a5960ba1b606082015260800190565b60208082526025908201527f4361727650726f746f636f6c536572766963653a206e6f6e636520697320696e6040820152641d985b1a5960da1b606082015260800190565b60208082526030908201527f4361727650726f746f636f6c536572766963653a20766572696679207468652060408201526f1cda59db985d1d5c994819985a5b195960821b606082015260800190565b6001600160a01b0392831681529116602082015260400190565b6020808252601b908201527a1d995c9a599e481d1a19481cda59db985d1d5c994819985a5b1959602a1b604082015260600190565b60006020828403121561581b57600080fd5b813561263981614d0a565b6001600160a01b0386168152606060208083018290529082018590526000906001600160fb1b0386111561585957600080fd5b8560051b808860808601378301838103608090810160408601528101859052859060009060a0015b8682101561185d57823561589481614d0a565b1515815291830191600191909101908301615881565b600083516158bc81846020880161471e565b8351908301906158d081836020880161471e565b01949350505050565b6001600160a01b03871681526080602082018190526000906158fe9083018789615592565b8560408401528281036060840152615917818587615592565b9998505050505050505050565b60208082526025908201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060408201526437bbb732b960d91b606082015260800190565b60006020828403121561597b57600080fd5b81356126398161536a565b60006020828403121561599857600080fd5b815161263981614d0a565b6001600160a01b039485168152928416602084015292166040820152606081019190915260800190565b634e487b7160e01b600052602160045260246000fd5b6000602082840312156159f557600080fd5b5051919050565b600081615a0b57615a0b615037565b506000190190565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b76020b1b1b2b9b9a1b7b73a3937b61d1030b1b1b7bab73a1604d1b815260008351615a9781601785016020880161471e565b7001034b99036b4b9b9b4b733903937b6329607d1b6017918401918201528351615ac881602884016020880161471e565b01602801949350505050565b828152604060208201526000612db36040830184614742565b60008251615aff81846020870161471e565b9190910192915050565b6001600160a01b03858116825284166020820152604081018390526080606082018190526000906133b290830184614742565b600060208284031215615b4e57600080fd5b8151612639816146eb565b8082028115828204841417610c9757610c97615037565b80820180821115610c9757610c9761503756fed16a327e6f5c32c69c8282ab355bc8a366432cf60ee1165bc5198414ca1b06c7379ea341c2e88a5e16cbe91a8487b8b783789df95f99c134d2c5b73b4c5ecf223fb90a982568460bdf5505b984928e3c942db3525e60c25e39051cacec08b60fa264697066735822122064567dfdb024f9be15bd35bb97875303bb70b2a7607d87c238f4414f062c1c1364736f6c63430008110033",
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

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(uint256)
func (_Contract *ContractCaller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(uint256)
func (_Contract *ContractSession) CHAINID() (*big.Int, error) {
	return _Contract.Contract.CHAINID(&_Contract.CallOpts)
}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(uint256)
func (_Contract *ContractCallerSession) CHAINID() (*big.Int, error) {
	return _Contract.Contract.CHAINID(&_Contract.CallOpts)
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

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xe306f779.
//
// Solidity: function EIP712_DOMAIN_HASH() view returns(bytes32)
func (_Contract *ContractCaller) EIP712DOMAINHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "EIP712_DOMAIN_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xe306f779.
//
// Solidity: function EIP712_DOMAIN_HASH() view returns(bytes32)
func (_Contract *ContractSession) EIP712DOMAINHASH() ([32]byte, error) {
	return _Contract.Contract.EIP712DOMAINHASH(&_Contract.CallOpts)
}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xe306f779.
//
// Solidity: function EIP712_DOMAIN_HASH() view returns(bytes32)
func (_Contract *ContractCallerSession) EIP712DOMAINHASH() ([32]byte, error) {
	return _Contract.Contract.EIP712DOMAINHASH(&_Contract.CallOpts)
}

// FORWARDERROLE is a free data retrieval call binding the contract method 0x35f6fb87.
//
// Solidity: function FORWARDER_ROLE() view returns(bytes32)
func (_Contract *ContractCaller) FORWARDERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "FORWARDER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FORWARDERROLE is a free data retrieval call binding the contract method 0x35f6fb87.
//
// Solidity: function FORWARDER_ROLE() view returns(bytes32)
func (_Contract *ContractSession) FORWARDERROLE() ([32]byte, error) {
	return _Contract.Contract.FORWARDERROLE(&_Contract.CallOpts)
}

// FORWARDERROLE is a free data retrieval call binding the contract method 0x35f6fb87.
//
// Solidity: function FORWARDER_ROLE() view returns(bytes32)
func (_Contract *ContractCallerSession) FORWARDERROLE() ([32]byte, error) {
	return _Contract.Contract.FORWARDERROLE(&_Contract.CallOpts)
}

// PLATFORMMINETRROLE is a free data retrieval call binding the contract method 0x0d44ad76.
//
// Solidity: function PLATFORM_MINETR_ROLE() view returns(bytes32)
func (_Contract *ContractCaller) PLATFORMMINETRROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "PLATFORM_MINETR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PLATFORMMINETRROLE is a free data retrieval call binding the contract method 0x0d44ad76.
//
// Solidity: function PLATFORM_MINETR_ROLE() view returns(bytes32)
func (_Contract *ContractSession) PLATFORMMINETRROLE() ([32]byte, error) {
	return _Contract.Contract.PLATFORMMINETRROLE(&_Contract.CallOpts)
}

// PLATFORMMINETRROLE is a free data retrieval call binding the contract method 0x0d44ad76.
//
// Solidity: function PLATFORM_MINETR_ROLE() view returns(bytes32)
func (_Contract *ContractCallerSession) PLATFORMMINETRROLE() ([32]byte, error) {
	return _Contract.Contract.PLATFORMMINETRROLE(&_Contract.CallOpts)
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

// AdminAddress is a free data retrieval call binding the contract method 0x4f9d5a65.
//
// Solidity: function _admin_address() view returns(address)
func (_Contract *ContractCaller) AdminAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_admin_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminAddress is a free data retrieval call binding the contract method 0x4f9d5a65.
//
// Solidity: function _admin_address() view returns(address)
func (_Contract *ContractSession) AdminAddress() (common.Address, error) {
	return _Contract.Contract.AdminAddress(&_Contract.CallOpts)
}

// AdminAddress is a free data retrieval call binding the contract method 0x4f9d5a65.
//
// Solidity: function _admin_address() view returns(address)
func (_Contract *ContractCallerSession) AdminAddress() (common.Address, error) {
	return _Contract.Contract.AdminAddress(&_Contract.CallOpts)
}

// CampaignId is a free data retrieval call binding the contract method 0x079d6a93.
//
// Solidity: function _campaign_id() view returns(string)
func (_Contract *ContractCaller) CampaignId(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_campaign_id")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CampaignId is a free data retrieval call binding the contract method 0x079d6a93.
//
// Solidity: function _campaign_id() view returns(string)
func (_Contract *ContractSession) CampaignId() (string, error) {
	return _Contract.Contract.CampaignId(&_Contract.CallOpts)
}

// CampaignId is a free data retrieval call binding the contract method 0x079d6a93.
//
// Solidity: function _campaign_id() view returns(string)
func (_Contract *ContractCallerSession) CampaignId() (string, error) {
	return _Contract.Contract.CampaignId(&_Contract.CallOpts)
}

// CampaignInfo is a free data retrieval call binding the contract method 0x3ed9c6da.
//
// Solidity: function _campaign_info() view returns(string)
func (_Contract *ContractCaller) CampaignInfo(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_campaign_info")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CampaignInfo is a free data retrieval call binding the contract method 0x3ed9c6da.
//
// Solidity: function _campaign_info() view returns(string)
func (_Contract *ContractSession) CampaignInfo() (string, error) {
	return _Contract.Contract.CampaignInfo(&_Contract.CallOpts)
}

// CampaignInfo is a free data retrieval call binding the contract method 0x3ed9c6da.
//
// Solidity: function _campaign_info() view returns(string)
func (_Contract *ContractCallerSession) CampaignInfo() (string, error) {
	return _Contract.Contract.CampaignInfo(&_Contract.CallOpts)
}

// CarvId is a free data retrieval call binding the contract method 0x9ff098ac.
//
// Solidity: function _carv_id() view returns(uint256)
func (_Contract *ContractCaller) CarvId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_carv_id")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CarvId is a free data retrieval call binding the contract method 0x9ff098ac.
//
// Solidity: function _carv_id() view returns(uint256)
func (_Contract *ContractSession) CarvId() (*big.Int, error) {
	return _Contract.Contract.CarvId(&_Contract.CallOpts)
}

// CarvId is a free data retrieval call binding the contract method 0x9ff098ac.
//
// Solidity: function _carv_id() view returns(uint256)
func (_Contract *ContractCallerSession) CarvId() (*big.Int, error) {
	return _Contract.Contract.CarvId(&_Contract.CallOpts)
}

// CurTokenId is a free data retrieval call binding the contract method 0xa4bf0813.
//
// Solidity: function _cur_token_id() view returns(uint256)
func (_Contract *ContractCaller) CurTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_cur_token_id")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurTokenId is a free data retrieval call binding the contract method 0xa4bf0813.
//
// Solidity: function _cur_token_id() view returns(uint256)
func (_Contract *ContractSession) CurTokenId() (*big.Int, error) {
	return _Contract.Contract.CurTokenId(&_Contract.CallOpts)
}

// CurTokenId is a free data retrieval call binding the contract method 0xa4bf0813.
//
// Solidity: function _cur_token_id() view returns(uint256)
func (_Contract *ContractCallerSession) CurTokenId() (*big.Int, error) {
	return _Contract.Contract.CurTokenId(&_Contract.CallOpts)
}

// AddressTeeInfo is a free data retrieval call binding the contract method 0xc1fb555e.
//
// Solidity: function addressTeeInfo(address ) view returns(string publicKey, string mrEnclave)
func (_Contract *ContractCaller) AddressTeeInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	PublicKey string
	MrEnclave string
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "addressTeeInfo", arg0)

	outstruct := new(struct {
		PublicKey string
		MrEnclave string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PublicKey = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.MrEnclave = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// AddressTeeInfo is a free data retrieval call binding the contract method 0xc1fb555e.
//
// Solidity: function addressTeeInfo(address ) view returns(string publicKey, string mrEnclave)
func (_Contract *ContractSession) AddressTeeInfo(arg0 common.Address) (struct {
	PublicKey string
	MrEnclave string
}, error) {
	return _Contract.Contract.AddressTeeInfo(&_Contract.CallOpts, arg0)
}

// AddressTeeInfo is a free data retrieval call binding the contract method 0xc1fb555e.
//
// Solidity: function addressTeeInfo(address ) view returns(string publicKey, string mrEnclave)
func (_Contract *ContractCallerSession) AddressTeeInfo(arg0 common.Address) (struct {
	PublicKey string
	MrEnclave string
}, error) {
	return _Contract.Contract.AddressTeeInfo(&_Contract.CallOpts, arg0)
}

// AddressUserMap is a free data retrieval call binding the contract method 0x89fb5f4a.
//
// Solidity: function address_user_map(address ) view returns(uint256 token_id, string user_profile_path, uint256 profile_version, bytes signature)
func (_Contract *ContractCaller) AddressUserMap(opts *bind.CallOpts, arg0 common.Address) (struct {
	TokenId         *big.Int
	UserProfilePath string
	ProfileVersion  *big.Int
	Signature       []byte
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "address_user_map", arg0)

	outstruct := new(struct {
		TokenId         *big.Int
		UserProfilePath string
		ProfileVersion  *big.Int
		Signature       []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UserProfilePath = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.ProfileVersion = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Signature = *abi.ConvertType(out[3], new([]byte)).(*[]byte)

	return *outstruct, err

}

// AddressUserMap is a free data retrieval call binding the contract method 0x89fb5f4a.
//
// Solidity: function address_user_map(address ) view returns(uint256 token_id, string user_profile_path, uint256 profile_version, bytes signature)
func (_Contract *ContractSession) AddressUserMap(arg0 common.Address) (struct {
	TokenId         *big.Int
	UserProfilePath string
	ProfileVersion  *big.Int
	Signature       []byte
}, error) {
	return _Contract.Contract.AddressUserMap(&_Contract.CallOpts, arg0)
}

// AddressUserMap is a free data retrieval call binding the contract method 0x89fb5f4a.
//
// Solidity: function address_user_map(address ) view returns(uint256 token_id, string user_profile_path, uint256 profile_version, bytes signature)
func (_Contract *ContractCallerSession) AddressUserMap(arg0 common.Address) (struct {
	TokenId         *big.Int
	UserProfilePath string
	ProfileVersion  *big.Int
	Signature       []byte
}, error) {
	return _Contract.Contract.AddressUserMap(&_Contract.CallOpts, arg0)
}

// AddressVoteWeight is a free data retrieval call binding the contract method 0x47be46e4.
//
// Solidity: function address_vote_weight(address ) view returns(uint256)
func (_Contract *ContractCaller) AddressVoteWeight(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "address_vote_weight", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressVoteWeight is a free data retrieval call binding the contract method 0x47be46e4.
//
// Solidity: function address_vote_weight(address ) view returns(uint256)
func (_Contract *ContractSession) AddressVoteWeight(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.AddressVoteWeight(&_Contract.CallOpts, arg0)
}

// AddressVoteWeight is a free data retrieval call binding the contract method 0x47be46e4.
//
// Solidity: function address_vote_weight(address ) view returns(uint256)
func (_Contract *ContractCallerSession) AddressVoteWeight(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.AddressVoteWeight(&_Contract.CallOpts, arg0)
}

// AttestationIdList is a free data retrieval call binding the contract method 0xd5c6329b.
//
// Solidity: function attestation_id_list(uint256 ) view returns(bytes32)
func (_Contract *ContractCaller) AttestationIdList(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "attestation_id_list", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AttestationIdList is a free data retrieval call binding the contract method 0xd5c6329b.
//
// Solidity: function attestation_id_list(uint256 ) view returns(bytes32)
func (_Contract *ContractSession) AttestationIdList(arg0 *big.Int) ([32]byte, error) {
	return _Contract.Contract.AttestationIdList(&_Contract.CallOpts, arg0)
}

// AttestationIdList is a free data retrieval call binding the contract method 0xd5c6329b.
//
// Solidity: function attestation_id_list(uint256 ) view returns(bytes32)
func (_Contract *ContractCallerSession) AttestationIdList(arg0 *big.Int) ([32]byte, error) {
	return _Contract.Contract.AttestationIdList(&_Contract.CallOpts, arg0)
}

// AttestationIdMap is a free data retrieval call binding the contract method 0x97baf354.
//
// Solidity: function attestation_id_map(bytes32 ) view returns(bool)
func (_Contract *ContractCaller) AttestationIdMap(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "attestation_id_map", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AttestationIdMap is a free data retrieval call binding the contract method 0x97baf354.
//
// Solidity: function attestation_id_map(bytes32 ) view returns(bool)
func (_Contract *ContractSession) AttestationIdMap(arg0 [32]byte) (bool, error) {
	return _Contract.Contract.AttestationIdMap(&_Contract.CallOpts, arg0)
}

// AttestationIdMap is a free data retrieval call binding the contract method 0x97baf354.
//
// Solidity: function attestation_id_map(bytes32 ) view returns(bool)
func (_Contract *ContractCallerSession) AttestationIdMap(arg0 [32]byte) (bool, error) {
	return _Contract.Contract.AttestationIdMap(&_Contract.CallOpts, arg0)
}

// AttestationIdResultMap is a free data retrieval call binding the contract method 0x1a514dad.
//
// Solidity: function attestation_id_result_map(bytes32 ) view returns(bool)
func (_Contract *ContractCaller) AttestationIdResultMap(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "attestation_id_result_map", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AttestationIdResultMap is a free data retrieval call binding the contract method 0x1a514dad.
//
// Solidity: function attestation_id_result_map(bytes32 ) view returns(bool)
func (_Contract *ContractSession) AttestationIdResultMap(arg0 [32]byte) (bool, error) {
	return _Contract.Contract.AttestationIdResultMap(&_Contract.CallOpts, arg0)
}

// AttestationIdResultMap is a free data retrieval call binding the contract method 0x1a514dad.
//
// Solidity: function attestation_id_result_map(bytes32 ) view returns(bool)
func (_Contract *ContractCallerSession) AttestationIdResultMap(arg0 [32]byte) (bool, error) {
	return _Contract.Contract.AttestationIdResultMap(&_Contract.CallOpts, arg0)
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

// CampainRewardMap is a free data retrieval call binding the contract method 0x0ded29e9.
//
// Solidity: function campain_reward_map(string ) view returns(string campaign_id, address user_address, uint256 reward_amount, uint256 total_num, address contract_address, uint8 contract_type)
func (_Contract *ContractCaller) CampainRewardMap(opts *bind.CallOpts, arg0 string) (struct {
	CampaignId      string
	UserAddress     common.Address
	RewardAmount    *big.Int
	TotalNum        *big.Int
	ContractAddress common.Address
	ContractType    uint8
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "campain_reward_map", arg0)

	outstruct := new(struct {
		CampaignId      string
		UserAddress     common.Address
		RewardAmount    *big.Int
		TotalNum        *big.Int
		ContractAddress common.Address
		ContractType    uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CampaignId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.UserAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.RewardAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalNum = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ContractAddress = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.ContractType = *abi.ConvertType(out[5], new(uint8)).(*uint8)

	return *outstruct, err

}

// CampainRewardMap is a free data retrieval call binding the contract method 0x0ded29e9.
//
// Solidity: function campain_reward_map(string ) view returns(string campaign_id, address user_address, uint256 reward_amount, uint256 total_num, address contract_address, uint8 contract_type)
func (_Contract *ContractSession) CampainRewardMap(arg0 string) (struct {
	CampaignId      string
	UserAddress     common.Address
	RewardAmount    *big.Int
	TotalNum        *big.Int
	ContractAddress common.Address
	ContractType    uint8
}, error) {
	return _Contract.Contract.CampainRewardMap(&_Contract.CallOpts, arg0)
}

// CampainRewardMap is a free data retrieval call binding the contract method 0x0ded29e9.
//
// Solidity: function campain_reward_map(string ) view returns(string campaign_id, address user_address, uint256 reward_amount, uint256 total_num, address contract_address, uint8 contract_type)
func (_Contract *ContractCallerSession) CampainRewardMap(arg0 string) (struct {
	CampaignId      string
	UserAddress     common.Address
	RewardAmount    *big.Int
	TotalNum        *big.Int
	ContractAddress common.Address
	ContractType    uint8
}, error) {
	return _Contract.Contract.CampainRewardMap(&_Contract.CallOpts, arg0)
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

// GetProofList is a free data retrieval call binding the contract method 0x0008f8f9.
//
// Solidity: function get_proof_list(uint256 startId, uint256 endId) view returns(bytes32[])
func (_Contract *ContractCaller) GetProofList(opts *bind.CallOpts, startId *big.Int, endId *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "get_proof_list", startId, endId)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetProofList is a free data retrieval call binding the contract method 0x0008f8f9.
//
// Solidity: function get_proof_list(uint256 startId, uint256 endId) view returns(bytes32[])
func (_Contract *ContractSession) GetProofList(startId *big.Int, endId *big.Int) ([][32]byte, error) {
	return _Contract.Contract.GetProofList(&_Contract.CallOpts, startId, endId)
}

// GetProofList is a free data retrieval call binding the contract method 0x0008f8f9.
//
// Solidity: function get_proof_list(uint256 startId, uint256 endId) view returns(bytes32[])
func (_Contract *ContractCallerSession) GetProofList(startId *big.Int, endId *big.Int) ([][32]byte, error) {
	return _Contract.Contract.GetProofList(&_Contract.CallOpts, startId, endId)
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

// IdCampaignMap is a free data retrieval call binding the contract method 0x4679d490.
//
// Solidity: function id_campaign_map(string ) view returns(string campaign_id, string url, address creator, uint8 campaign_type, address reward_contract, uint256 reward_total_amount, uint256 reward_count, uint8 status, uint256 start_time, uint256 end_time, string requirements)
func (_Contract *ContractCaller) IdCampaignMap(opts *bind.CallOpts, arg0 string) (struct {
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
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "id_campaign_map", arg0)

	outstruct := new(struct {
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
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CampaignId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Url = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Creator = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.CampaignType = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.RewardContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.RewardTotalAmount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.RewardCount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[7], new(uint8)).(*uint8)
	outstruct.StartTime = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.Requirements = *abi.ConvertType(out[10], new(string)).(*string)

	return *outstruct, err

}

// IdCampaignMap is a free data retrieval call binding the contract method 0x4679d490.
//
// Solidity: function id_campaign_map(string ) view returns(string campaign_id, string url, address creator, uint8 campaign_type, address reward_contract, uint256 reward_total_amount, uint256 reward_count, uint8 status, uint256 start_time, uint256 end_time, string requirements)
func (_Contract *ContractSession) IdCampaignMap(arg0 string) (struct {
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
}, error) {
	return _Contract.Contract.IdCampaignMap(&_Contract.CallOpts, arg0)
}

// IdCampaignMap is a free data retrieval call binding the contract method 0x4679d490.
//
// Solidity: function id_campaign_map(string ) view returns(string campaign_id, string url, address creator, uint8 campaign_type, address reward_contract, uint256 reward_total_amount, uint256 reward_count, uint8 status, uint256 start_time, uint256 end_time, string requirements)
func (_Contract *ContractCallerSession) IdCampaignMap(arg0 string) (struct {
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
}, error) {
	return _Contract.Contract.IdCampaignMap(&_Contract.CallOpts, arg0)
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

// NftAddress is a free data retrieval call binding the contract method 0xc78de709.
//
// Solidity: function nft_address() view returns(address)
func (_Contract *ContractCaller) NftAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "nft_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftAddress is a free data retrieval call binding the contract method 0xc78de709.
//
// Solidity: function nft_address() view returns(address)
func (_Contract *ContractSession) NftAddress() (common.Address, error) {
	return _Contract.Contract.NftAddress(&_Contract.CallOpts)
}

// NftAddress is a free data retrieval call binding the contract method 0xc78de709.
//
// Solidity: function nft_address() view returns(address)
func (_Contract *ContractCallerSession) NftAddress() (common.Address, error) {
	return _Contract.Contract.NftAddress(&_Contract.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Contract *ContractCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Contract *ContractSession) Nonce() (*big.Int, error) {
	return _Contract.Contract.Nonce(&_Contract.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Contract *ContractCallerSession) Nonce() (*big.Int, error) {
	return _Contract.Contract.Nonce(&_Contract.CallOpts)
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

// VaultAddress is a free data retrieval call binding the contract method 0x733f2399.
//
// Solidity: function vault_address() view returns(address)
func (_Contract *ContractCaller) VaultAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "vault_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultAddress is a free data retrieval call binding the contract method 0x733f2399.
//
// Solidity: function vault_address() view returns(address)
func (_Contract *ContractSession) VaultAddress() (common.Address, error) {
	return _Contract.Contract.VaultAddress(&_Contract.CallOpts)
}

// VaultAddress is a free data retrieval call binding the contract method 0x733f2399.
//
// Solidity: function vault_address() view returns(address)
func (_Contract *ContractCallerSession) VaultAddress() (common.Address, error) {
	return _Contract.Contract.VaultAddress(&_Contract.CallOpts)
}

// VerifierBlock is a free data retrieval call binding the contract method 0x2c8ebcde.
//
// Solidity: function verifier_block(address ) view returns(uint256)
func (_Contract *ContractCaller) VerifierBlock(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "verifier_block", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VerifierBlock is a free data retrieval call binding the contract method 0x2c8ebcde.
//
// Solidity: function verifier_block(address ) view returns(uint256)
func (_Contract *ContractSession) VerifierBlock(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.VerifierBlock(&_Contract.CallOpts, arg0)
}

// VerifierBlock is a free data retrieval call binding the contract method 0x2c8ebcde.
//
// Solidity: function verifier_block(address ) view returns(uint256)
func (_Contract *ContractCallerSession) VerifierBlock(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.VerifierBlock(&_Contract.CallOpts, arg0)
}

// VerifierList is a free data retrieval call binding the contract method 0x5656a852.
//
// Solidity: function verifier_list(uint256 ) view returns(address)
func (_Contract *ContractCaller) VerifierList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "verifier_list", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VerifierList is a free data retrieval call binding the contract method 0x5656a852.
//
// Solidity: function verifier_list(uint256 ) view returns(address)
func (_Contract *ContractSession) VerifierList(arg0 *big.Int) (common.Address, error) {
	return _Contract.Contract.VerifierList(&_Contract.CallOpts, arg0)
}

// VerifierList is a free data retrieval call binding the contract method 0x5656a852.
//
// Solidity: function verifier_list(uint256 ) view returns(address)
func (_Contract *ContractCallerSession) VerifierList(arg0 *big.Int) (common.Address, error) {
	return _Contract.Contract.VerifierList(&_Contract.CallOpts, arg0)
}

// VerifierPassThreshold is a free data retrieval call binding the contract method 0xe6640252.
//
// Solidity: function verifier_pass_threshold() view returns(uint256)
func (_Contract *ContractCaller) VerifierPassThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "verifier_pass_threshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VerifierPassThreshold is a free data retrieval call binding the contract method 0xe6640252.
//
// Solidity: function verifier_pass_threshold() view returns(uint256)
func (_Contract *ContractSession) VerifierPassThreshold() (*big.Int, error) {
	return _Contract.Contract.VerifierPassThreshold(&_Contract.CallOpts)
}

// VerifierPassThreshold is a free data retrieval call binding the contract method 0xe6640252.
//
// Solidity: function verifier_pass_threshold() view returns(uint256)
func (_Contract *ContractCallerSession) VerifierPassThreshold() (*big.Int, error) {
	return _Contract.Contract.VerifierPassThreshold(&_Contract.CallOpts)
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

// VrfAddress is a free data retrieval call binding the contract method 0x18a14be9.
//
// Solidity: function vrf_address() view returns(address)
func (_Contract *ContractCaller) VrfAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "vrf_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VrfAddress is a free data retrieval call binding the contract method 0x18a14be9.
//
// Solidity: function vrf_address() view returns(address)
func (_Contract *ContractSession) VrfAddress() (common.Address, error) {
	return _Contract.Contract.VrfAddress(&_Contract.CallOpts)
}

// VrfAddress is a free data retrieval call binding the contract method 0x18a14be9.
//
// Solidity: function vrf_address() view returns(address)
func (_Contract *ContractCallerSession) VrfAddress() (common.Address, error) {
	return _Contract.Contract.VrfAddress(&_Contract.CallOpts)
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

// CloseVerifierNode is a paid mutator transaction binding the contract method 0x9bb5bb38.
//
// Solidity: function close_verifier_node((address,uint256,uint256) verifier_data) returns()
func (_Contract *ContractTransactor) CloseVerifierNode(opts *bind.TransactOpts, verifier_data CarvProtocolServiceopenVerifierNodeData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "close_verifier_node", verifier_data)
}

// CloseVerifierNode is a paid mutator transaction binding the contract method 0x9bb5bb38.
//
// Solidity: function close_verifier_node((address,uint256,uint256) verifier_data) returns()
func (_Contract *ContractSession) CloseVerifierNode(verifier_data CarvProtocolServiceopenVerifierNodeData) (*types.Transaction, error) {
	return _Contract.Contract.CloseVerifierNode(&_Contract.TransactOpts, verifier_data)
}

// CloseVerifierNode is a paid mutator transaction binding the contract method 0x9bb5bb38.
//
// Solidity: function close_verifier_node((address,uint256,uint256) verifier_data) returns()
func (_Contract *ContractTransactorSession) CloseVerifierNode(verifier_data CarvProtocolServiceopenVerifierNodeData) (*types.Transaction, error) {
	return _Contract.Contract.CloseVerifierNode(&_Contract.TransactOpts, verifier_data)
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

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address rewards_address, address _nft_address, uint256 _chainId) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, rewards_address common.Address, _nft_address common.Address, _chainId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", rewards_address, _nft_address, _chainId)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address rewards_address, address _nft_address, uint256 _chainId) returns()
func (_Contract *ContractSession) Initialize(rewards_address common.Address, _nft_address common.Address, _chainId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, rewards_address, _nft_address, _chainId)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address rewards_address, address _nft_address, uint256 _chainId) returns()
func (_Contract *ContractTransactorSession) Initialize(rewards_address common.Address, _nft_address common.Address, _chainId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, rewards_address, _nft_address, _chainId)
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

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address _to) returns()
func (_Contract *ContractTransactor) Mint(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "mint", _to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address _to) returns()
func (_Contract *ContractSession) Mint(_to common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Mint(&_Contract.TransactOpts, _to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address _to) returns()
func (_Contract *ContractTransactorSession) Mint(_to common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Mint(&_Contract.TransactOpts, _to)
}

// OpenVerifierNode is a paid mutator transaction binding the contract method 0xffd2f5cc.
//
// Solidity: function open_verifier_node((address,uint256,uint256) verifier_data, bytes signature) returns()
func (_Contract *ContractTransactor) OpenVerifierNode(opts *bind.TransactOpts, verifier_data CarvProtocolServiceopenVerifierNodeData, signature []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "open_verifier_node", verifier_data, signature)
}

// OpenVerifierNode is a paid mutator transaction binding the contract method 0xffd2f5cc.
//
// Solidity: function open_verifier_node((address,uint256,uint256) verifier_data, bytes signature) returns()
func (_Contract *ContractSession) OpenVerifierNode(verifier_data CarvProtocolServiceopenVerifierNodeData, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.OpenVerifierNode(&_Contract.TransactOpts, verifier_data, signature)
}

// OpenVerifierNode is a paid mutator transaction binding the contract method 0xffd2f5cc.
//
// Solidity: function open_verifier_node((address,uint256,uint256) verifier_data, bytes signature) returns()
func (_Contract *ContractTransactorSession) OpenVerifierNode(verifier_data CarvProtocolServiceopenVerifierNodeData, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.OpenVerifierNode(&_Contract.TransactOpts, verifier_data, signature)
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

// SetTeeInfo is a paid mutator transaction binding the contract method 0x6a1aec4a.
//
// Solidity: function setTeeInfo(string publicKey, string mrEnclave) returns()
func (_Contract *ContractTransactor) SetTeeInfo(opts *bind.TransactOpts, publicKey string, mrEnclave string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setTeeInfo", publicKey, mrEnclave)
}

// SetTeeInfo is a paid mutator transaction binding the contract method 0x6a1aec4a.
//
// Solidity: function setTeeInfo(string publicKey, string mrEnclave) returns()
func (_Contract *ContractSession) SetTeeInfo(publicKey string, mrEnclave string) (*types.Transaction, error) {
	return _Contract.Contract.SetTeeInfo(&_Contract.TransactOpts, publicKey, mrEnclave)
}

// SetTeeInfo is a paid mutator transaction binding the contract method 0x6a1aec4a.
//
// Solidity: function setTeeInfo(string publicKey, string mrEnclave) returns()
func (_Contract *ContractTransactorSession) SetTeeInfo(publicKey string, mrEnclave string) (*types.Transaction, error) {
	return _Contract.Contract.SetTeeInfo(&_Contract.TransactOpts, publicKey, mrEnclave)
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

// SetNftAddress is a paid mutator transaction binding the contract method 0x27c4a224.
//
// Solidity: function set_nft_address(address _nft_address) returns()
func (_Contract *ContractTransactor) SetNftAddress(opts *bind.TransactOpts, _nft_address common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "set_nft_address", _nft_address)
}

// SetNftAddress is a paid mutator transaction binding the contract method 0x27c4a224.
//
// Solidity: function set_nft_address(address _nft_address) returns()
func (_Contract *ContractSession) SetNftAddress(_nft_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetNftAddress(&_Contract.TransactOpts, _nft_address)
}

// SetNftAddress is a paid mutator transaction binding the contract method 0x27c4a224.
//
// Solidity: function set_nft_address(address _nft_address) returns()
func (_Contract *ContractTransactorSession) SetNftAddress(_nft_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetNftAddress(&_Contract.TransactOpts, _nft_address)
}

// SetVaultAddress is a paid mutator transaction binding the contract method 0x70824ab1.
//
// Solidity: function set_vault_address(address _vault_address) returns()
func (_Contract *ContractTransactor) SetVaultAddress(opts *bind.TransactOpts, _vault_address common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "set_vault_address", _vault_address)
}

// SetVaultAddress is a paid mutator transaction binding the contract method 0x70824ab1.
//
// Solidity: function set_vault_address(address _vault_address) returns()
func (_Contract *ContractSession) SetVaultAddress(_vault_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetVaultAddress(&_Contract.TransactOpts, _vault_address)
}

// SetVaultAddress is a paid mutator transaction binding the contract method 0x70824ab1.
//
// Solidity: function set_vault_address(address _vault_address) returns()
func (_Contract *ContractTransactorSession) SetVaultAddress(_vault_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetVaultAddress(&_Contract.TransactOpts, _vault_address)
}

// SetVerifierPassThreshold is a paid mutator transaction binding the contract method 0x91582484.
//
// Solidity: function set_verifier_pass_threshold(uint256 _verifier_pass_threshold) returns()
func (_Contract *ContractTransactor) SetVerifierPassThreshold(opts *bind.TransactOpts, _verifier_pass_threshold *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "set_verifier_pass_threshold", _verifier_pass_threshold)
}

// SetVerifierPassThreshold is a paid mutator transaction binding the contract method 0x91582484.
//
// Solidity: function set_verifier_pass_threshold(uint256 _verifier_pass_threshold) returns()
func (_Contract *ContractSession) SetVerifierPassThreshold(_verifier_pass_threshold *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetVerifierPassThreshold(&_Contract.TransactOpts, _verifier_pass_threshold)
}

// SetVerifierPassThreshold is a paid mutator transaction binding the contract method 0x91582484.
//
// Solidity: function set_verifier_pass_threshold(uint256 _verifier_pass_threshold) returns()
func (_Contract *ContractTransactorSession) SetVerifierPassThreshold(_verifier_pass_threshold *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetVerifierPassThreshold(&_Contract.TransactOpts, _verifier_pass_threshold)
}

// SetVrfAddress is a paid mutator transaction binding the contract method 0x5d1fe591.
//
// Solidity: function set_vrf_address(address _vrf_address) returns()
func (_Contract *ContractTransactor) SetVrfAddress(opts *bind.TransactOpts, _vrf_address common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "set_vrf_address", _vrf_address)
}

// SetVrfAddress is a paid mutator transaction binding the contract method 0x5d1fe591.
//
// Solidity: function set_vrf_address(address _vrf_address) returns()
func (_Contract *ContractSession) SetVrfAddress(_vrf_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetVrfAddress(&_Contract.TransactOpts, _vrf_address)
}

// SetVrfAddress is a paid mutator transaction binding the contract method 0x5d1fe591.
//
// Solidity: function set_vrf_address(address _vrf_address) returns()
func (_Contract *ContractTransactorSession) SetVrfAddress(_vrf_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetVrfAddress(&_Contract.TransactOpts, _vrf_address)
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

// VerifierDelegate is a paid mutator transaction binding the contract method 0x85910c25.
//
// Solidity: function verifier_delegate((address,address,uint256,uint256,uint256) delegate_data, bytes signature) returns()
func (_Contract *ContractTransactor) VerifierDelegate(opts *bind.TransactOpts, delegate_data CarvProtocolServicedelegateData, signature []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "verifier_delegate", delegate_data, signature)
}

// VerifierDelegate is a paid mutator transaction binding the contract method 0x85910c25.
//
// Solidity: function verifier_delegate((address,address,uint256,uint256,uint256) delegate_data, bytes signature) returns()
func (_Contract *ContractSession) VerifierDelegate(delegate_data CarvProtocolServicedelegateData, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.VerifierDelegate(&_Contract.TransactOpts, delegate_data, signature)
}

// VerifierDelegate is a paid mutator transaction binding the contract method 0x85910c25.
//
// Solidity: function verifier_delegate((address,address,uint256,uint256,uint256) delegate_data, bytes signature) returns()
func (_Contract *ContractTransactorSession) VerifierDelegate(delegate_data CarvProtocolServicedelegateData, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.VerifierDelegate(&_Contract.TransactOpts, delegate_data, signature)
}

// VerifierRedelegate is a paid mutator transaction binding the contract method 0xe5180e4e.
//
// Solidity: function verifier_redelegate((address,address,uint256,uint256,uint256) delegate_data, bytes signature) returns()
func (_Contract *ContractTransactor) VerifierRedelegate(opts *bind.TransactOpts, delegate_data CarvProtocolServicedelegateData, signature []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "verifier_redelegate", delegate_data, signature)
}

// VerifierRedelegate is a paid mutator transaction binding the contract method 0xe5180e4e.
//
// Solidity: function verifier_redelegate((address,address,uint256,uint256,uint256) delegate_data, bytes signature) returns()
func (_Contract *ContractSession) VerifierRedelegate(delegate_data CarvProtocolServicedelegateData, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.VerifierRedelegate(&_Contract.TransactOpts, delegate_data, signature)
}

// VerifierRedelegate is a paid mutator transaction binding the contract method 0xe5180e4e.
//
// Solidity: function verifier_redelegate((address,address,uint256,uint256,uint256) delegate_data, bytes signature) returns()
func (_Contract *ContractTransactorSession) VerifierRedelegate(delegate_data CarvProtocolServicedelegateData, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.VerifierRedelegate(&_Contract.TransactOpts, delegate_data, signature)
}

// VerifierUndelegate is a paid mutator transaction binding the contract method 0xba030180.
//
// Solidity: function verifier_undelegate((address,address,uint256,uint256,uint256) delegate_data, bytes signature) returns()
func (_Contract *ContractTransactor) VerifierUndelegate(opts *bind.TransactOpts, delegate_data CarvProtocolServicedelegateData, signature []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "verifier_undelegate", delegate_data, signature)
}

// VerifierUndelegate is a paid mutator transaction binding the contract method 0xba030180.
//
// Solidity: function verifier_undelegate((address,address,uint256,uint256,uint256) delegate_data, bytes signature) returns()
func (_Contract *ContractSession) VerifierUndelegate(delegate_data CarvProtocolServicedelegateData, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.VerifierUndelegate(&_Contract.TransactOpts, delegate_data, signature)
}

// VerifierUndelegate is a paid mutator transaction binding the contract method 0xba030180.
//
// Solidity: function verifier_undelegate((address,address,uint256,uint256,uint256) delegate_data, bytes signature) returns()
func (_Contract *ContractTransactorSession) VerifierUndelegate(delegate_data CarvProtocolServicedelegateData, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.VerifierUndelegate(&_Contract.TransactOpts, delegate_data, signature)
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

// ContractOpenVerifierNodeIterator is returned from FilterOpenVerifierNode and is used to iterate over the raw logs and unpacked data for OpenVerifierNode events raised by the Contract contract.
type ContractOpenVerifierNodeIterator struct {
	Event *ContractOpenVerifierNode // Event containing the contract specifics and raw log

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
func (it *ContractOpenVerifierNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenVerifierNode)
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
		it.Event = new(ContractOpenVerifierNode)
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
func (it *ContractOpenVerifierNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenVerifierNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenVerifierNode represents a OpenVerifierNode event raised by the Contract contract.
type ContractOpenVerifierNode struct {
	Account   common.Address
	TokenId   *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOpenVerifierNode is a free log retrieval operation binding the contract event 0x832607f249995a2ed86c13bde705f2d0576d5eee1780a83084d8cee8992641c4.
//
// Solidity: event OpenVerifierNode(address account, uint256 token_id, uint256 timestamp)
func (_Contract *ContractFilterer) FilterOpenVerifierNode(opts *bind.FilterOpts) (*ContractOpenVerifierNodeIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OpenVerifierNode")
	if err != nil {
		return nil, err
	}
	return &ContractOpenVerifierNodeIterator{contract: _Contract.contract, event: "OpenVerifierNode", logs: logs, sub: sub}, nil
}

// WatchOpenVerifierNode is a free log subscription operation binding the contract event 0x832607f249995a2ed86c13bde705f2d0576d5eee1780a83084d8cee8992641c4.
//
// Solidity: event OpenVerifierNode(address account, uint256 token_id, uint256 timestamp)
func (_Contract *ContractFilterer) WatchOpenVerifierNode(opts *bind.WatchOpts, sink chan<- *ContractOpenVerifierNode) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OpenVerifierNode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenVerifierNode)
				if err := _Contract.contract.UnpackLog(event, "OpenVerifierNode", log); err != nil {
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

// ParseOpenVerifierNode is a log parse operation binding the contract event 0x832607f249995a2ed86c13bde705f2d0576d5eee1780a83084d8cee8992641c4.
//
// Solidity: event OpenVerifierNode(address account, uint256 token_id, uint256 timestamp)
func (_Contract *ContractFilterer) ParseOpenVerifierNode(log types.Log) (*ContractOpenVerifierNode, error) {
	event := new(ContractOpenVerifierNode)
	if err := _Contract.contract.UnpackLog(event, "OpenVerifierNode", log); err != nil {
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

// ContractRecoverParamsIterator is returned from FilterRecoverParams and is used to iterate over the raw logs and unpacked data for RecoverParams events raised by the Contract contract.
type ContractRecoverParamsIterator struct {
	Event *ContractRecoverParams // Event containing the contract specifics and raw log

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
func (it *ContractRecoverParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRecoverParams)
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
		it.Event = new(ContractRecoverParams)
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
func (it *ContractRecoverParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRecoverParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRecoverParams represents a RecoverParams event raised by the Contract contract.
type ContractRecoverParams struct {
	ChainId          *big.Int
	Eip712DomainHash [32]byte
	HashStruct       [32]byte
	Digest           [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRecoverParams is a free log retrieval operation binding the contract event 0xa50615fb8919bf8c8a6350a26c7921f53c0434b16d11c80df59de6b281a83bf1.
//
// Solidity: event RecoverParams(uint256 chainId, bytes32 eip712DomainHash, bytes32 hashStruct, bytes32 digest)
func (_Contract *ContractFilterer) FilterRecoverParams(opts *bind.FilterOpts) (*ContractRecoverParamsIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RecoverParams")
	if err != nil {
		return nil, err
	}
	return &ContractRecoverParamsIterator{contract: _Contract.contract, event: "RecoverParams", logs: logs, sub: sub}, nil
}

// WatchRecoverParams is a free log subscription operation binding the contract event 0xa50615fb8919bf8c8a6350a26c7921f53c0434b16d11c80df59de6b281a83bf1.
//
// Solidity: event RecoverParams(uint256 chainId, bytes32 eip712DomainHash, bytes32 hashStruct, bytes32 digest)
func (_Contract *ContractFilterer) WatchRecoverParams(opts *bind.WatchOpts, sink chan<- *ContractRecoverParams) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RecoverParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRecoverParams)
				if err := _Contract.contract.UnpackLog(event, "RecoverParams", log); err != nil {
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

// ParseRecoverParams is a log parse operation binding the contract event 0xa50615fb8919bf8c8a6350a26c7921f53c0434b16d11c80df59de6b281a83bf1.
//
// Solidity: event RecoverParams(uint256 chainId, bytes32 eip712DomainHash, bytes32 hashStruct, bytes32 digest)
func (_Contract *ContractFilterer) ParseRecoverParams(log types.Log) (*ContractRecoverParams, error) {
	event := new(ContractRecoverParams)
	if err := _Contract.contract.UnpackLog(event, "RecoverParams", log); err != nil {
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

// ContractVerifierWeightChangedIterator is returned from FilterVerifierWeightChanged and is used to iterate over the raw logs and unpacked data for VerifierWeightChanged events raised by the Contract contract.
type ContractVerifierWeightChangedIterator struct {
	Event *ContractVerifierWeightChanged // Event containing the contract specifics and raw log

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
func (it *ContractVerifierWeightChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractVerifierWeightChanged)
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
		it.Event = new(ContractVerifierWeightChanged)
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
func (it *ContractVerifierWeightChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractVerifierWeightChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractVerifierWeightChanged represents a VerifierWeightChanged event raised by the Contract contract.
type ContractVerifierWeightChanged struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterVerifierWeightChanged is a free log retrieval operation binding the contract event 0x379ea341c2e88a5e16cbe91a8487b8b783789df95f99c134d2c5b73b4c5ecf22.
//
// Solidity: event VerifierWeightChanged(address from, address to)
func (_Contract *ContractFilterer) FilterVerifierWeightChanged(opts *bind.FilterOpts) (*ContractVerifierWeightChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "VerifierWeightChanged")
	if err != nil {
		return nil, err
	}
	return &ContractVerifierWeightChangedIterator{contract: _Contract.contract, event: "VerifierWeightChanged", logs: logs, sub: sub}, nil
}

// WatchVerifierWeightChanged is a free log subscription operation binding the contract event 0x379ea341c2e88a5e16cbe91a8487b8b783789df95f99c134d2c5b73b4c5ecf22.
//
// Solidity: event VerifierWeightChanged(address from, address to)
func (_Contract *ContractFilterer) WatchVerifierWeightChanged(opts *bind.WatchOpts, sink chan<- *ContractVerifierWeightChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "VerifierWeightChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractVerifierWeightChanged)
				if err := _Contract.contract.UnpackLog(event, "VerifierWeightChanged", log); err != nil {
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

// ParseVerifierWeightChanged is a log parse operation binding the contract event 0x379ea341c2e88a5e16cbe91a8487b8b783789df95f99c134d2c5b73b4c5ecf22.
//
// Solidity: event VerifierWeightChanged(address from, address to)
func (_Contract *ContractFilterer) ParseVerifierWeightChanged(log types.Log) (*ContractVerifierWeightChanged, error) {
	event := new(ContractVerifierWeightChanged)
	if err := _Contract.contract.UnpackLog(event, "VerifierWeightChanged", log); err != nil {
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

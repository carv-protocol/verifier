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
	ABI: "[{\"inputs\":[],\"name\":\"PackedPtrLen__LenOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PackedPtrLen__PtrOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Slice__OutOfBounds\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token_id\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ProfixPayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tee_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"campaign_id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"attestation_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"attestation\",\"type\":\"string\"}],\"name\":\"ReportTeeAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardPayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"identitiesRoot\",\"type\":\"bytes32\"}],\"name\":\"SetIdentitiesRoot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contract_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"campaign_id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"requirements\",\"type\":\"string\"}],\"name\":\"SubmitCampaign\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"carv_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"campaign_id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"campaign_info\",\"type\":\"string\"}],\"name\":\"UserCampaignData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"attestation_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"name\":\"VerifyAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"attestation_ids\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"results\",\"type\":\"bool[]\"}],\"name\":\"VerifyAttestationBatch\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TEE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERIFIER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rewards_address\",\"type\":\"address\"}],\"name\":\"__CarvProtocolService_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tee_address\",\"type\":\"address\"}],\"name\":\"add_tee_role\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier_address\",\"type\":\"address\"}],\"name\":\"add_verifier_role\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getIdentitiesRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teeAddress\",\"type\":\"address\"}],\"name\":\"getTeeInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"attestation_id\",\"type\":\"bytes32\"}],\"name\":\"get_attestation_result\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_pay_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_proof_list\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_address\",\"type\":\"address\"}],\"name\":\"get_user_by_address\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"token_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"user_profile_path\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"profile_version\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structCarvProtocolService.user\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_verifier_list\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_vrf_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"carv_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"campaign_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"join_campaign_info\",\"type\":\"string\"}],\"name\":\"join_campaign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"campaign_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"attestation\",\"type\":\"string\"}],\"name\":\"report_tee_attestation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"multiIdentitiesRoot\",\"type\":\"bytes32\"}],\"name\":\"setIdentitiesRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"mrEnclave\",\"type\":\"string\"}],\"name\":\"setTeeInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_address\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"user_profile_path\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"profile_version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"multiIdentitiesRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"set_identities_root1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pay_address\",\"type\":\"address\"}],\"name\":\"set_pay_address\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vrf_address\",\"type\":\"address\"}],\"name\":\"set_vrf_address\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"campaign_id\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"user_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reward_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total_num\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"contract_address\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"contract_type\",\"type\":\"uint8\"}],\"internalType\":\"structCarvProtocolService.reward\",\"name\":\"reward_info\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"campaign_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"campaign_type\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"reward_contract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reward_total_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward_count\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"start_time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end_time\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"requirements\",\"type\":\"string\"}],\"internalType\":\"structCarvProtocolService.campaign\",\"name\":\"campaign_info\",\"type\":\"tuple\"}],\"name\":\"submit_campaign\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftOwnerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sigAddress\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"userIDs\",\"type\":\"string[]\"},{\"internalType\":\"bytes32\",\"name\":\"multiIdentitiesRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verifyIdentitiesBinding\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"attestation_id\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"name\":\"verify_attestation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"attestation_ids\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"results\",\"type\":\"bool[]\"}],\"name\":\"verify_attestation_batch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506145ad806100206000396000f3fe6080604052600436106101ec5760003560e01c806301ffc9a7146101f157806306fdde0314610226578063081812fc14610248578063095ea7b31461028057806309a9c7f0146102a257806310131a5c146102c25780631a7f9955146102e05780631d17e1bd1461030057806323b872dd1461031e578063248a9ca31461033e5780632f2ff15d1461036c57806336568abe1461038c5780633736c83f146103ac57806342842e0e146103bf57806345692b82146103df578063495eb8281461040c57806349d030371461042e57806349e9e8341461044e578063578907ca1461047e5780635d1fe5911461049e5780636352211e146104db5780636a1aec4a146104fb57806370a082311461051b57806370b3542d1461053b5780637e3461a31461055d57806391d148541461057d57806395d89b411461059d578063964b1f6e146105b2578063a217fddf146105d4578063a22cb465146105e9578063a28bfc4314610609578063a35d20be14610629578063ad86d54514610649578063b41080b014610669578063b88d4fde14610697578063c1f75848146106b7578063c87b56dd146106d7578063d547741f146106f7578063da7a292814610717578063e7705db614610744578063e97a998714610766578063e985e9c514610786578063ff30a118146107a6575b600080fd5b3480156101fd57600080fd5b5061021161020c366004613386565b6107c6565b60405190151581526020015b60405180910390f35b34801561023257600080fd5b5061023b6107d7565b60405161021d91906133f3565b34801561025457600080fd5b50610268610263366004613406565b610869565b6040516001600160a01b03909116815260200161021d565b34801561028c57600080fd5b506102a061029b366004613434565b610890565b005b3480156102ae57600080fd5b506102a06102bd3660046134a1565b6109aa565b3480156102ce57600080fd5b5060de546001600160a01b0316610268565b3480156102ec57600080fd5b506102a06102fb36600461351a565b610a11565b34801561030c57600080fd5b5060d3546001600160a01b0316610268565b34801561032a57600080fd5b506102a06103393660046135af565b610a92565b34801561034a57600080fd5b5061035e610359366004613406565b610ac3565b60405190815260200161021d565b34801561037857600080fd5b506102a06103873660046135f0565b610ad8565b34801561039857600080fd5b506102a06103a73660046135f0565b610af4565b6102a06103ba366004613620565b610b72565b3480156103cb57600080fd5b506102a06103da3660046135af565b610c50565b3480156103eb57600080fd5b5061035e6103fa366004613406565b60009081526097602052604090205490565b34801561041857600080fd5b5061035e60008051602061453883398151915281565b34801561043a57600080fd5b50610211610449366004613725565b610c6b565b34801561045a57600080fd5b50610211610469366004613406565b600090815260d9602052604090205460ff1690565b34801561048a57600080fd5b506102a061049936600461387d565b610daa565b3480156104aa57600080fd5b506102a06104b93660046138a2565b60de80546001600160a01b0319166001600160a01b0392909216919091179055565b3480156104e757600080fd5b506102686104f6366004613406565b610ee4565b34801561050757600080fd5b506102a06105163660046138bf565b610f18565b34801561052757600080fd5b5061035e6105363660046138a2565b610f5f565b34801561054757600080fd5b50610550610fe5565b60405161021d919061392a565b34801561056957600080fd5b506102a06105783660046138a2565b611046565b34801561058957600080fd5b506102116105983660046135f0565b611070565b3480156105a957600080fd5b5061023b61109b565b3480156105be57600080fd5b506105c76110aa565b60405161021d9190613977565b3480156105e057600080fd5b5061035e600081565b3480156105f557600080fd5b506102a06106043660046139af565b611101565b34801561061557600080fd5b506102a06106243660046138a2565b61110c565b34801561063557600080fd5b506102a06106443660046138a2565b611250565b34801561065557600080fd5b506102a06106643660046138a2565b611273565b34801561067557600080fd5b506106896106843660046138a2565b6112de565b60405161021d9291906139dd565b3480156106a357600080fd5b506102a06106b2366004613a0b565b611422565b3480156106c357600080fd5b506102a06106d2366004613ace565b61145a565b3480156106e357600080fd5b5061023b6106f2366004613406565b611660565b34801561070357600080fd5b506102a06107123660046135f0565b6116d4565b34801561072357600080fd5b506107376107323660046138a2565b6116f0565b60405161021d9190613b2d565b34801561075057600080fd5b5061035e60008051602061455883398151915281565b34801561077257600080fd5b506102a06107813660046138bf565b61187d565b34801561079257600080fd5b506102116107a1366004613b7a565b611957565b3480156107b257600080fd5b506102a06107c1366004613ba8565b611985565b60006107d1826119cd565b92915050565b6060606580546107e690613bca565b80601f016020809104026020016040519081016040528092919081815260200182805461081290613bca565b801561085f5780601f106108345761010080835404028352916020019161085f565b820191906000526020600020905b81548152906001019060200180831161084257829003601f168201915b5050505050905090565b6000610874826119f2565b506000908152606960205260409020546001600160a01b031690565b600061089b82610ee4565b9050806001600160a01b0316836001600160a01b03160361090d5760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b60648201526084015b60405180910390fd5b336001600160a01b038216148061092957506109298133611957565b61099b5760405162461bcd60e51b815260206004820152603d60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206f7220617070726f76656420666f7220616c6c0000006064820152608401610904565b6109a58383611a17565b505050565b60d085905560d16109bc848683613c61565b5060d26109ca828483613c61565b507f83eda6edb918aac3f35c8d616a6198dbb75eeca3248e0efec05e66a30d51119160d05460d160d2604051610a0293929190613d97565b60405180910390a15050505050565b6001600160a01b038716600090815260d760205260409020600101610a37868883613c61565b506001600160a01b038716600090815260d76020526040902060028101859055600301610a65828483613c61565b506001600160a01b038716600090815260d76020526040902054610a899084611985565b50505050505050565b610a9c3382611a85565b610ab85760405162461bcd60e51b815260040161090490613dc2565b6109a5838383611ae4565b60009081526099602052604090206001015490565b610ae182610ac3565b610aea81611c48565b6109a58383611c52565b6001600160a01b0381163314610b645760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b6064820152608401610904565b610b6e8282611cd8565b5050565b610b7c8233611d3f565b8060d5610b898280613e0f565b604051610b97929190613e55565b908152604051908190036020019020610bb08282613ea5565b5082905060d4610bc08280613e0f565b604051610bce929190613e55565b908152604051908190036020019020610be78282613fa2565b507fe8e692af5f416c4637c9307d7f08b7697fd8f4398f55103ce03ddd9388ccd84d9050610c1b60a08401608085016138a2565b610c258380613e0f565b610c33610140860186613e0f565b604051610c449594939291906140da565b60405180910390a15050565b6109a583838360405180602001604052806000815250611422565b6000866001600160a01b0316610c8089610ee4565b6001600160a01b031614610cd15760405162461bcd60e51b81526020600482015260186024820152771b999d081bdddb995c881a5cc81b9bdd0818dbdc9c9958dd60421b6044820152606401610904565b845180610d195760405162461bcd60e51b81526020600482015260166024820152757573657249442063616e6e6f7420626520656d70747960501b6044820152606401610904565b60005b81811015610d5857610d46878281518110610d3957610d3961411e565b6020026020010151611dd4565b80610d508161414a565b915050610d1c565b506000610d9c888787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061203992505050565b9a9950505050505050505050565b610db261209a565b610dbb8261210a565b610dd75760405162461bcd60e51b815260040161090490614163565b610de082612160565b15610e395760405162461bcd60e51b815260206004820152602360248201527f6174746573746174696f6e2063616e206e6f742062792076657269667920616760448201526230b4b760e91b6064820152608401610904565b600082815260da602090815260408220805460018101825590835291200180546001600160a01b03191633179055610e708261222d565b8015610e795750805b15610ea257610e8782612259565b600082815260d960205260409020805460ff19168215151790555b6040805133815260208101849052821515918101919091527f937b4d1b5b461852fdbbb174b68adefb42155e668c05a1261fae11a2d22a240a90606001610c44565b600080610ef083612324565b90506001600160a01b0381166107d15760405162461bcd60e51b815260040161090490614195565b610f2061233f565b33600090815260df60205260409020610f3a848683613c61565b5033600090815260df60205260409020600101610f58828483613c61565b5050505050565b60006001600160a01b038216610fc95760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b6064820152608401610904565b506001600160a01b031660009081526068602052604090205490565b606060dd80548060200260200160405190810160405280929190818152602001828054801561085f57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161101f575050505050905090565b61104e6123a2565b60d380546001600160a01b0319166001600160a01b0392909216919091179055565b60009182526099602090815260408084206001600160a01b0393909316845291905290205460ff1690565b6060606680546107e690613bca565b606060dc80548060200260200160405190810160405280929190818152602001828054801561085f57602002820191906000526020600020905b8154815260200190600101908083116110e4575050505050905090565b610b6e3383836123f9565b600054610100900460ff161580801561112c5750600054600160ff909116105b8061114d575061113b306124c3565b15801561114d575060005460ff166001145b6111b05760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610904565b6000805460ff1916600117905580156111d3576000805461ff0019166101001790555b60cc8054336001600160a01b0319918216811790925560cb80549091166001600160a01b038516179055600160cd5561120e906000906124d2565b8015610b6e576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602001610c44565b6112586123a2565b611270600080516020614538833981519152826124d2565b50565b61127b6123a2565b60dd80546001810182556000919091527fac507b9f8bf86ad8bb770f71cd2b1992902ae0314d93fc0f2bb011d70e7962260180546001600160a01b0319166001600160a01b038316179055611270600080516020614558833981519152826124d2565b6001600160a01b038116600090815260df60205260409020805460609182916001820190829061130d90613bca565b80601f016020809104026020016040519081016040528092919081815260200182805461133990613bca565b80156113865780601f1061135b57610100808354040283529160200191611386565b820191906000526020600020905b81548152906001019060200180831161136957829003601f168201915b5050505050915080805461139990613bca565b80601f01602080910402602001604051908101604052809291908181526020018280546113c590613bca565b80156114125780601f106113e757610100808354040283529160200191611412565b820191906000526020600020905b8154815290600101906020018083116113f557829003601f168201915b5050505050905091509150915091565b61142c3383611a85565b6114485760405162461bcd60e51b815260040161090490613dc2565b611454848484846124dc565b50505050565b61146261209a565b8260005b818110156116285761148f8686838181106114835761148361411e565b9050602002013561210a565b6114ab5760405162461bcd60e51b815260040161090490614163565b8383828181106114bd576114bd61411e565b90506020020160208101906114d291906141c7565b60d960008888858181106114e8576114e861411e565b90506020020135815260200190815260200160002060006101000a81548160ff02191690831515021790555060da600087878481811061152a5761152a61411e565b6020908102929092013583525081810192909252604001600090812080546001810182559082529190200180546001600160a01b0319163317905561158686868381811061157a5761157a61411e565b9050602002013561222d565b15611616576115ac8686838181106115a0576115a061411e565b90506020020135612259565b8383828181106115be576115be61411e565b90506020020160208101906115d391906141c7565b60d960008888858181106115e9576115e961411e565b90506020020135815260200190815260200160002060006101000a81548160ff0219169083151502179055505b806116208161414a565b915050611466565b507f04798cc2c828443d2d9ef127deaa1f73e782d8a9fffbc9db2bcca53e12765af53386868686604051610a029594939291906141e4565b606061166b826119f2565b600061168260408051602081019091526000815290565b905060008151116116a257604051806020016040528060008152506116cd565b806116ac8461250f565b6040516020016116bd929190614268565b6040516020818303038152906040525b9392505050565b6116dd82610ac3565b6116e681611c48565b6109a58383611cd8565b61171b6040518060800160405280600081526020016060815260200160008152602001606081525090565b6001600160a01b038216600090815260d7602090815260409182902082516080810190935280548352600181018054919284019161175890613bca565b80601f016020809104026020016040519081016040528092919081815260200182805461178490613bca565b80156117d15780601f106117a6576101008083540402835291602001916117d1565b820191906000526020600020905b8154815290600101906020018083116117b457829003601f168201915b50505050508152602001600282015481526020016003820180546117f490613bca565b80601f016020809104026020016040519081016040528092919081815260200182805461182090613bca565b801561186d5780601f106118425761010080835404028352916020019161186d565b820191906000526020600020905b81548152906001019060200180831161185057829003601f168201915b5050505050815250509050919050565b61188561233f565b60008282604051611897929190613e55565b60405190819003812060dc805460018101825560009182527f3162b0988d4210bff484413ed451d170a03887272177efc0b7d000f10abe9edf018290559092509060d4906118e89088908890613e55565b90815260408051918290036020908101832060020154600086815260db90925291902081905591507f99a038e9d345d0b12130b3b1fb003bf8f2d3a5c27ce2a800bbb1608efff6c591906119479033908990899087908a908a90614297565b60405180910390a1505050505050565b6001600160a01b039182166000908152606a6020908152604080832093909416825291909152205460ff1690565b60008281526097602090815260409182902083905581518481529081018390527f21b5a040ca6f61b8bcacb8f423e25ce46e88932f887ce92e60343369c20ec06f9101610c44565b60006001600160e01b03198216637965db0b60e01b14806107d157506107d1826125a1565b6119fb816125c6565b6112705760405162461bcd60e51b815260040161090490614195565b600081815260696020526040902080546001600160a01b0319166001600160a01b0384169081179091558190611a4c82610ee4565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b600080611a9183610ee4565b9050806001600160a01b0316846001600160a01b03161480611ab85750611ab88185611957565b80611adc5750836001600160a01b0316611ad184610869565b6001600160a01b0316145b949350505050565b826001600160a01b0316611af782610ee4565b6001600160a01b031614611b1d5760405162461bcd60e51b8152600401610904906142e2565b6001600160a01b038216611b7f5760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b6064820152608401610904565b826001600160a01b0316611b9282610ee4565b6001600160a01b031614611bb85760405162461bcd60e51b8152600401610904906142e2565b600081815260696020908152604080832080546001600160a01b03199081169091556001600160a01b0387811680865260688552838620805460001901905590871680865283862080546001019055868652606790945282852080549092168417909155905184937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b61127081336125e3565b611c5c8282611070565b610b6e5760008281526099602090815260408083206001600160a01b03851684529091529020805460ff19166001179055611c943390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b611ce28282611070565b15610b6e5760008281526099602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b6000611d5160a08401608085016138a2565b60cc54909150611d7290829084906001600160a01b0316604087013561263c565b7f704db66c978a7472842ddae0d478a8a7d123fb6ad449e4426b5e0f6a22081de7611da360a08501608086016138a2565b60cc5460408051611dc7939287926001600160a01b03909116919089013590614327565b60405180910390a1505050565b6000815111611e1f5760405162461bcd60e51b81526020600482015260176024820152767573657249442063616e206e6f7420626520656d70747960481b6044820152606401610904565b6040805180820190915260018152601d60f91b602082015260008080611e4485612696565b9050611e59611e5285612696565b82906126a1565b9194509250905082611ebd5760405162461bcd60e51b815260206004820152602760248201527f74686520666972737420706172742064656c696d6974657220646f6573206e6f6044820152661d08195e1a5cdd60ca1b6064820152608401610904565b6000611ec8836126f0565b5111611f165760405162461bcd60e51b815260206004820152601d60248201527f746865206669727374207061727420646f6573206e6f742065786973740000006044820152606401610904565b611f22611e5285612696565b9194509250905082611f875760405162461bcd60e51b815260206004820152602860248201527f746865207365636f6e6420706172742064656c696d6974657220646f6573206e6044820152671bdd08195e1a5cdd60c21b6064820152608401610904565b6000611f92836126f0565b5111611fe05760405162461bcd60e51b815260206004820152601e60248201527f746865207365636f6e64207061727420646f6573206e6f7420657869737400006044820152606401610904565b611fe9816126f0565b51604014610f585760405162461bcd60e51b815260206004820152601d60248201527f69642068617368206c656e677468206973206e6f7420636f72726563740000006044820152606401610904565b600080600061204885856126fb565b9092509050600081600481111561206157612061614351565b14801561207f5750856001600160a01b0316826001600160a01b0316145b806120905750612090868686612740565b9695505050505050565b6120b260008051602061455883398151915233611070565b6121085760405162461bcd60e51b815260206004820152602160248201527f73656e64657220646f65736e2774206861766520766572696669657220726f6c6044820152606560f81b6064820152608401610904565b565b6000805b60dc54811015612157578260dc828154811061212c5761212c61411e565b9060005260206000200154036121455750600192915050565b8061214f8161414a565b91505061210e565b50600092915050565b600081815260da60209081526040808320805482518185028101850190935280835284938301828280156121bd57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161219f575b505083519394506000925050505b8181101561222257336001600160a01b03168382815181106121ef576121ef61411e565b60200260200101516001600160a01b03160361221057506001949350505050565b8061221a8161414a565b9150506121cb565b506000949350505050565b600081815260da602052604081205460dd5461224a826002614367565b11156121575750600192915050565b60d35460408051808201909152601b81527a3c3c361039b7b6103830bcafb83937b334bc101d1d1d101717171760291b60208201526001600160a01b03909116906122a39061282c565b60cc54600083815260db60205260409020546122ce9183916001600160a01b0390911690339061263c565b60d35460cc54600084815260db6020526040908190205490517f3832ba0f9bc9de56bd1f974e665b53f2643ed48a47b45c3d36976bc1ea31459993610c44936001600160a01b0391821693911691339190614327565b6000908152606760205260409020546001600160a01b031690565b61235760008051602061453883398151915233611070565b6121085760405162461bcd60e51b815260206004820152601c60248201527b73656e64657220646f65736e277420686176652074656520726f6c6560201b6044820152606401610904565b6123ad600033611070565b6121085760405162461bcd60e51b815260206004820152601e60248201527f73656e64657220646f65736e277420686176652061646d696e20726f6c6500006044820152606401610904565b816001600160a01b0316836001600160a01b0316036124565760405162461bcd60e51b815260206004820152601960248201527822a9219b99189d1030b8383937bb32903a379031b0b63632b960391b6044820152606401610904565b6001600160a01b038381166000818152606a6020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6001600160a01b03163b151590565b610b6e8282611c52565b6124e7848484611ae4565b6124f38484848461286f565b6114545760405162461bcd60e51b81526004016109049061437e565b6060600061251c83612977565b60010190506000816001600160401b0381111561253b5761253b613688565b6040519080825280601f01601f191660200182016040528015612565576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a850494508461256f57509392505050565b60006001600160e01b0319821663f389baad60e01b14806107d157506107d182612a4d565b6000806125d283612324565b6001600160a01b0316141592915050565b6125ed8282611070565b610b6e576125fa81612a9d565b612605836020612aaf565b6040516020016126169291906143d0565b60408051601f198184030181529082905262461bcd60e51b8252610904916004016133f3565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052611454908590612c4a565b60006107d182612d9c565b60008080806126b08686612db0565b905060001981036126cc576000866000935093509350506126e9565b6126e1816126d987612ec7565b889190612ed3565b935093509350505b9250925092565b60606107d182612f42565b60008082516041036127315760208301516040840151606085015160001a61272587828585612fb7565b94509450505050612739565b506000905060025b9250929050565b6000806000856001600160a01b0316631626ba7e60e01b868660405160240161276a92919061443f565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516127a89190614458565b600060405180830381855afa9150503d80600081146127e3576040519150601f19603f3d011682016040523d82523d6000602084013e6127e8565b606091505b50915091508180156127fc57506020815110155b801561209057508051630b135d3f60e11b906128219083016020908101908401614474565b149695505050505050565b6112708160405160240161284091906133f3565b60408051601f198184030181529190526020810180516001600160e01b031663104c13eb60e21b179052613071565b6000612883846001600160a01b03166124c3565b1561296c57604051630a85bd0160e11b81526001600160a01b0385169063150b7a02906128ba90339089908890889060040161448d565b6020604051808303816000875af19250505080156128f5575060408051601f3d908101601f191682019092526128f2918101906144c0565b60015b612952573d808015612923576040519150601f19603f3d011682016040523d82523d6000602084013e612928565b606091505b50805160000361294a5760405162461bcd60e51b81526004016109049061437e565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050611adc565b506001949350505050565b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b83106129b65772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6904ee2d6d415b85acef8160201b83106129e0576904ee2d6d415b85acef8160201b830492506020015b662386f26fc1000083106129fe57662386f26fc10000830492506010015b6305f5e1008310612a16576305f5e100830492506008015b6127108310612a2a57612710830492506004015b60648310612a3c576064830492506002015b600a83106107d15760010192915050565b60006001600160e01b031982166380ac58cd60e01b1480612a7e57506001600160e01b03198216635b5e139f60e01b145b806107d157506301ffc9a760e01b6001600160e01b03198316146107d1565b60606107d16001600160a01b03831660145b60606000612abe836002614367565b612ac99060026144dd565b6001600160401b03811115612ae057612ae0613688565b6040519080825280601f01601f191660200182016040528015612b0a576020820181803683370190505b509050600360fc1b81600081518110612b2557612b2561411e565b60200101906001600160f81b031916908160001a905350600f60fb1b81600181518110612b5457612b5461411e565b60200101906001600160f81b031916908160001a9053506000612b78846002614367565b612b839060016144dd565b90505b6001811115612bfb576f181899199a1a9b1b9c1cb0b131b232b360811b85600f1660108110612bb757612bb761411e565b1a60f81b828281518110612bcd57612bcd61411e565b60200101906001600160f81b031916908160001a90535060049490941c93612bf4816144f0565b9050612b86565b5083156116cd5760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401610904565b813b80612c915760405162461bcd60e51b8152602060048201526015602482015274115490cc8c0e881b9bdd08184818dbdb9d1c9858dd605a1b6044820152606401610904565b600080846001600160a01b031684604051612cac9190614458565b6000604051808303816000865af19150503d8060008114612ce9576040519150601f19603f3d011682016040523d82523d6000602084013e612cee565b606091505b509150915081612d355760405162461bcd60e51b8152602060048201526012602482015271115490cc8c0e8818d85b1b0819985a5b195960721b6044820152606401610904565b805115610f585780806020019051810190612d509190614507565b610f585760405162461bcd60e51b815260206004820181905260248201527f45524332303a206f7065726174696f6e20646964206e6f7420737563636565646044820152606401610904565b6000806020830190506116cd81845161307a565b600080612dbc84612ec7565b90506000612dc984612ec7565b905080600003612dde576000925050506107d1565b811580612dea57508181115b15612dfb57600019925050506107d1565b6000612e078660801c90565b90506000612e158660801c90565b90506000612e24825160001a90565b90505b6000612e34848784613086565b90506000198103612e4f5760001996505050505050506107d1565b94859003949283019285851115612e705760001996505050505050506107d1565b84832085852003612e9c57612e858960801c90565b612e8f9085614524565b96505050505050506107d1565b85600103612eb45760001996505050505050506107d1565b6000199095019460019093019250612e27565b6001600160801b031690565b600080600080612ee38760801c90565b90506000612ef088612ec7565b905086860181811115612f16576040516365f4e9df60e01b815260040160405180910390fd5b612f208389613146565b9450612f30818401828403613146565b93506001955050505093509350939050565b6060612f4d82612ec7565b6001600160401b03811115612f6457612f64613688565b6040519080825280601f01601f191660200182016040528015612f8e576020820181803683370190505b50905060208101612fb181612fa38560801c90565b612fac86612ec7565b613159565b50919050565b6000806fa2a8918ca85bafe22016d0b997e4df60600160ff1b03831115612fe45750600090506003613068565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015613038573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661306157600060019250925050613068565b9150600090505b94509492505050565b61127081613167565b60006116cd8383613188565b6000602083116130a25761309b848484613194565b90506116cd565b83601f8416808503820160006130b786613297565b90505b818810156130fc57875181186130cf816132bf565b156130f0578489036130e38a60208a613194565b01955050505050506116cd565b602089019850506130ba565b82600003613112576000199450505050506116cd565b61311d828488613194565b94506000198503613136576000199450505050506116cd565b83820385019450505050506116cd565b6001600160801b031660809190911b1790565b8083828460045afa50505050565b60006a636f6e736f6c652e6c6f679050600080835160208501845afa505050565b60006116cd838361330a565b82516000908160208511156131a857602094505b6012851061322b5760006131bb85613297565b831890506131d16001600160801b0382176132bf565b6000036132075760109150601a8610613202576131f66001600160401b0382176132bf565b60000361320257601891505b613225565b6132196001600160c01b0382176132bf565b60000361322557600891505b50613262565b600a851061326257600061323e85613297565b831890506132546001600160c01b0382176132bf565b60000361326057600891505b505b8481101561328a5781811a60ff85168103613281575091506116cd9050565b50600101613262565b5060001995945050505050565b60ff167f01010101010101010101010101010101010101010101010101010101010101010290565b7ffefefefefefefefefefefefefefefefefefefefefefefefefefefefefefefeff81019019167f80808080808080808080808080808080808080808080808080808080808080801690565b60006001600160801b038311156133345760405163fee7506f60e01b815260040160405180910390fd5b6001600160801b0382111561335c57604051633b6b098d60e01b815260040160405180910390fd5b506001600160801b031660809190911b1790565b6001600160e01b03198116811461127057600080fd5b60006020828403121561339857600080fd5b81356116cd81613370565b60005b838110156133be5781810151838201526020016133a6565b50506000910152565b600081518084526133df8160208601602086016133a3565b601f01601f19169290920160200192915050565b6020815260006116cd60208301846133c7565b60006020828403121561341857600080fd5b5035919050565b6001600160a01b038116811461127057600080fd5b6000806040838503121561344757600080fd5b82356134528161341f565b946020939093013593505050565b60008083601f84011261347257600080fd5b5081356001600160401b0381111561348957600080fd5b60208301915083602082850101111561273957600080fd5b6000806000806000606086880312156134b957600080fd5b8535945060208601356001600160401b03808211156134d757600080fd5b6134e389838a01613460565b909650945060408801359150808211156134fc57600080fd5b5061350988828901613460565b969995985093965092949392505050565b600080600080600080600060a0888a03121561353557600080fd5b87356135408161341f565b965060208801356001600160401b038082111561355c57600080fd5b6135688b838c01613460565b909850965060408a0135955060608a0135945060808a013591508082111561358f57600080fd5b5061359c8a828b01613460565b989b979a50959850939692959293505050565b6000806000606084860312156135c457600080fd5b83356135cf8161341f565b925060208401356135df8161341f565b929592945050506040919091013590565b6000806040838503121561360357600080fd5b8235915060208301356136158161341f565b809150509250929050565b6000806040838503121561363357600080fd5b82356001600160401b038082111561364a57600080fd5b9084019060c0828703121561365e57600080fd5b9092506020840135908082111561367457600080fd5b508301610160818603121561361557600080fd5b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b03811182821017156136c6576136c6613688565b604052919050565b60006001600160401b038311156136e7576136e7613688565b6136fa601f8401601f191660200161369e565b905082815283838301111561370e57600080fd5b828260208301376000602084830101529392505050565b600080600080600080600060c0888a03121561374057600080fd5b87359650613751602089013561341f565b60208801359550613765604089013561341f565b604088013594506001600160401b03606089013581101561378557600080fd5b606089013589018a601f82011261379b57600080fd5b81813511156137ac576137ac613688565b6137bc6020823560051b0161369e565b81358082526020808301929160051b8401018d8111156137db57600080fd5b602084015b8181101561382c5785813511156137f657600080fd5b8e603f82358701011261380857600080fd5b61381e8f823587016020810135906040016136ce565b8452602093840193016137e0565b5050809750505050608089013593508060a08a0135111561384c57600080fd5b5061385d8960a08a01358a01613460565b979a9699509497509295919491925050565b801515811461127057600080fd5b6000806040838503121561389057600080fd5b8235915060208301356136158161386f565b6000602082840312156138b457600080fd5b81356116cd8161341f565b600080600080604085870312156138d557600080fd5b84356001600160401b03808211156138ec57600080fd5b6138f888838901613460565b9096509450602087013591508082111561391157600080fd5b5061391e87828801613460565b95989497509550505050565b6020808252825182820181905260009190848201906040850190845b8181101561396b5783516001600160a01b031683529284019291840191600101613946565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b8181101561396b57835183529284019291840191600101613993565b600080604083850312156139c257600080fd5b82356139cd8161341f565b915060208301356136158161386f565b6040815260006139f060408301856133c7565b8281036020840152613a0281856133c7565b95945050505050565b60008060008060808587031215613a2157600080fd5b8435613a2c8161341f565b93506020850135613a3c8161341f565b92506040850135915060608501356001600160401b03811115613a5e57600080fd5b8501601f81018713613a6f57600080fd5b613a7e878235602084016136ce565b91505092959194509250565b60008083601f840112613a9c57600080fd5b5081356001600160401b03811115613ab357600080fd5b6020830191508360208260051b850101111561273957600080fd5b60008060008060408587031215613ae457600080fd5b84356001600160401b0380821115613afb57600080fd5b613b0788838901613a8a565b90965094506020870135915080821115613b2057600080fd5b5061391e87828801613a8a565b60208152815160208201526000602083015160806040840152613b5360a08401826133c7565b9050604084015160608401526060840151601f19848303016080850152613a0282826133c7565b60008060408385031215613b8d57600080fd5b8235613b988161341f565b915060208301356136158161341f565b60008060408385031215613bbb57600080fd5b50508035926020909101359150565b600181811c90821680613bde57607f821691505b602082108103612fb157634e487b7160e01b600052602260045260246000fd5b601f8211156109a557600081815260208120601f850160051c81016020861015613c255750805b601f850160051c820191505b81811015613c4457828155600101613c31565b505050505050565b600019600383901b1c191660019190911b1790565b6001600160401b03831115613c7857613c78613688565b613c8c83613c868354613bca565b83613bfe565b6000601f841160018114613cba5760008515613ca85750838201355b613cb28682613c4c565b845550610f58565b600083815260209020601f19861690835b82811015613ceb5786850135825560209485019460019092019101613ccb565b5086821015613d085760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b60008154613d2781613bca565b808552602060018381168015613d445760018114613d5e57613d8c565b60ff1985168884015283151560051b880183019550613d8c565b866000528260002060005b85811015613d845781548a8201860152908301908401613d69565b890184019650505b505050505092915050565b838152606060208201526000613db06060830185613d1a565b82810360408401526120908185613d1a565b6020808252602d908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526c1c881bdc88185c1c1c9bdd9959609a1b606082015260800190565b6000808335601e19843603018112613e2657600080fd5b8301803591506001600160401b03821115613e4057600080fd5b60200191503681900382131561273957600080fd5b8183823760009101908152919050565b600081356107d18161341f565b80546001600160a01b0319166001600160a01b0392909216919091179055565b6000813560ff811681146107d157600080fd5b613eaf8283613e0f565b613eba818385613c61565b5050613ec96020830183613e0f565b613ed7818360018601613c61565b505060028101613ef2613eec60408501613e65565b82613e72565b613f1f613f0160608501613e92565b82805460ff60a01b191660a09290921b60ff60a01b16919091179055565b50613f38613f2f60808401613e65565b60038301613e72565b60a0820135600482015560c08201356005820155613f70613f5b60e08401613e92565b6006830160ff821660ff198254161781555050565b61010082013560078201556101208201356008820155613f94610140830183613e0f565b611454818360098601613c61565b613fac8283613e0f565b6001600160401b03811115613fc357613fc3613688565b613fd781613fd18554613bca565b85613bfe565b6000601f8211600181146140055760008315613ff35750838201355b613ffd8482613c4c565b86555061405f565b600085815260209020601f19841690835b828110156140365786850135825560209485019460019092019101614016565b50848210156140535760001960f88660031b161c19848701351681555b505060018360011b0185555b5050505061407b61407260208401613e65565b60018301613e72565b6040820135600282015560608201356003820155600481016140a2613eec60808501613e65565b6109a5613f0160a08501613e92565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b03861681526060602082018190526000906140ff90830186886140b1565b82810360408401526141128185876140b1565b98975050505050505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006001820161415c5761415c614134565b5060010190565b602080825260189082015277185d1d195cdd185d1a5bdb881a5cc81b9bdd08195e1a5cdd60421b604082015260600190565b602080825260189082015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b604082015260600190565b6000602082840312156141d957600080fd5b81356116cd8161386f565b6001600160a01b0386168152606060208083018290529082018590526000906001600160fb1b0386111561421757600080fd5b8560051b808860808601378301838103608090810160408601528101859052859060009060a0015b86821015610d9c5782356142528161386f565b151581529183019160019190910190830161423f565b6000835161427a8184602088016133a3565b83519083019061428e8183602088016133a3565b01949350505050565b6001600160a01b03871681526080602082018190526000906142bc90830187896140b1565b85604084015282810360608401526142d58185876140b1565b9998505050505050505050565b60208082526025908201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060408201526437bbb732b960d91b606082015260800190565b6001600160a01b039485168152928416602084015292166040820152606081019190915260800190565b634e487b7160e01b600052602160045260246000fd5b80820281158282048414176107d1576107d1614134565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b76020b1b1b2b9b9a1b7b73a3937b61d1030b1b1b7bab73a1604d1b8152600083516144028160178501602088016133a3565b7001034b99036b4b9b9b4b733903937b6329607d1b60179184019182015283516144338160288401602088016133a3565b01602801949350505050565b828152604060208201526000611adc60408301846133c7565b6000825161446a8184602087016133a3565b9190910192915050565b60006020828403121561448657600080fd5b5051919050565b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090612090908301846133c7565b6000602082840312156144d257600080fd5b81516116cd81613370565b808201808211156107d1576107d1614134565b6000816144ff576144ff614134565b506000190190565b60006020828403121561451957600080fd5b81516116cd8161386f565b818103818111156107d1576107d161413456fed16a327e6f5c32c69c8282ab355bc8a366432cf60ee1165bc5198414ca1b06c70ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea09a26469706673582212209baadaa41ca8ae766071ba5842fb053178f987a4875d76286edf6c4c695f529564736f6c63430008110033",
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

// GetTeeInfo is a free data retrieval call binding the contract method 0xb41080b0.
//
// Solidity: function getTeeInfo(address teeAddress) view returns(string, string)
func (_Contract *ContractCaller) GetTeeInfo(opts *bind.CallOpts, teeAddress common.Address) (string, string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getTeeInfo", teeAddress)

	if err != nil {
		return *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// GetTeeInfo is a free data retrieval call binding the contract method 0xb41080b0.
//
// Solidity: function getTeeInfo(address teeAddress) view returns(string, string)
func (_Contract *ContractSession) GetTeeInfo(teeAddress common.Address) (string, string, error) {
	return _Contract.Contract.GetTeeInfo(&_Contract.CallOpts, teeAddress)
}

// GetTeeInfo is a free data retrieval call binding the contract method 0xb41080b0.
//
// Solidity: function getTeeInfo(address teeAddress) view returns(string, string)
func (_Contract *ContractCallerSession) GetTeeInfo(teeAddress common.Address) (string, string, error) {
	return _Contract.Contract.GetTeeInfo(&_Contract.CallOpts, teeAddress)
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

// GetVrfAddress is a free data retrieval call binding the contract method 0x10131a5c.
//
// Solidity: function get_vrf_address() view returns(address)
func (_Contract *ContractCaller) GetVrfAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "get_vrf_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVrfAddress is a free data retrieval call binding the contract method 0x10131a5c.
//
// Solidity: function get_vrf_address() view returns(address)
func (_Contract *ContractSession) GetVrfAddress() (common.Address, error) {
	return _Contract.Contract.GetVrfAddress(&_Contract.CallOpts)
}

// GetVrfAddress is a free data retrieval call binding the contract method 0x10131a5c.
//
// Solidity: function get_vrf_address() view returns(address)
func (_Contract *ContractCallerSession) GetVrfAddress() (common.Address, error) {
	return _Contract.Contract.GetVrfAddress(&_Contract.CallOpts)
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

// SetIdentitiesRoot is a paid mutator transaction binding the contract method 0xff30a118.
//
// Solidity: function setIdentitiesRoot(uint256 id, bytes32 multiIdentitiesRoot) returns()
func (_Contract *ContractTransactor) SetIdentitiesRoot(opts *bind.TransactOpts, id *big.Int, multiIdentitiesRoot [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setIdentitiesRoot", id, multiIdentitiesRoot)
}

// SetIdentitiesRoot is a paid mutator transaction binding the contract method 0xff30a118.
//
// Solidity: function setIdentitiesRoot(uint256 id, bytes32 multiIdentitiesRoot) returns()
func (_Contract *ContractSession) SetIdentitiesRoot(id *big.Int, multiIdentitiesRoot [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SetIdentitiesRoot(&_Contract.TransactOpts, id, multiIdentitiesRoot)
}

// SetIdentitiesRoot is a paid mutator transaction binding the contract method 0xff30a118.
//
// Solidity: function setIdentitiesRoot(uint256 id, bytes32 multiIdentitiesRoot) returns()
func (_Contract *ContractTransactorSession) SetIdentitiesRoot(id *big.Int, multiIdentitiesRoot [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SetIdentitiesRoot(&_Contract.TransactOpts, id, multiIdentitiesRoot)
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

// SetIdentitiesRoot1 is a paid mutator transaction binding the contract method 0xc678daa1.
//
// Solidity: function set_identities_root1(address user_address, string user_profile_path, uint256 profile_version, bytes32 multiIdentitiesRoot, bytes signature) returns()
func (_Contract *ContractTransactor) SetIdentitiesRoot1(opts *bind.TransactOpts, user_address common.Address, user_profile_path string, profile_version *big.Int, multiIdentitiesRoot [32]byte, signature []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "set_identities_root1", user_address, user_profile_path, profile_version, multiIdentitiesRoot, signature)
}

// SetIdentitiesRoot1 is a paid mutator transaction binding the contract method 0xc678daa1.
//
// Solidity: function set_identities_root1(address user_address, string user_profile_path, uint256 profile_version, bytes32 multiIdentitiesRoot, bytes signature) returns()
func (_Contract *ContractSession) SetIdentitiesRoot1(user_address common.Address, user_profile_path string, profile_version *big.Int, multiIdentitiesRoot [32]byte, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetIdentitiesRoot1(&_Contract.TransactOpts, user_address, user_profile_path, profile_version, multiIdentitiesRoot, signature)
}

// SetIdentitiesRoot1 is a paid mutator transaction binding the contract method 0xc678daa1.
//
// Solidity: function set_identities_root1(address user_address, string user_profile_path, uint256 profile_version, bytes32 multiIdentitiesRoot, bytes signature) returns()
func (_Contract *ContractTransactorSession) SetIdentitiesRoot1(user_address common.Address, user_profile_path string, profile_version *big.Int, multiIdentitiesRoot [32]byte, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetIdentitiesRoot1(&_Contract.TransactOpts, user_address, user_profile_path, profile_version, multiIdentitiesRoot, signature)
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

// SetVrfAddress is a paid mutator transaction binding the contract method 0x5d1fe591.
//
// Solidity: function set_vrf_address(address vrf_address) returns()
func (_Contract *ContractTransactor) SetVrfAddress(opts *bind.TransactOpts, vrf_address common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "set_vrf_address", vrf_address)
}

// SetVrfAddress is a paid mutator transaction binding the contract method 0x5d1fe591.
//
// Solidity: function set_vrf_address(address vrf_address) returns()
func (_Contract *ContractSession) SetVrfAddress(vrf_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetVrfAddress(&_Contract.TransactOpts, vrf_address)
}

// SetVrfAddress is a paid mutator transaction binding the contract method 0x5d1fe591.
//
// Solidity: function set_vrf_address(address vrf_address) returns()
func (_Contract *ContractTransactorSession) SetVrfAddress(vrf_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetVrfAddress(&_Contract.TransactOpts, vrf_address)
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

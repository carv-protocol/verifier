package worker

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/carv-protocol/verifier/api/gasless"
	common2 "github.com/carv-protocol/verifier/internal/common"
	"github.com/carv-protocol/verifier/pkg/tools"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/go-kratos/kratos/v2/log"
	"math/big"
	"strconv"
	"time"
)

// Update node commission rate by service
func UpdateNodeCommissionRateByGaslessService(ctx context.Context, c *Chain, commissionRate uint32, expiredAt *big.Int) (bool, error) {
	_, found := c.cache.Get(common2.UPDATE_NODE_COMMISSION_RATE_BY_GASLESS_SERVICE)
	if found {
		c.logger.WithContext(ctx).Errorf("UpdateNodeCommissionRateByGaslessService is in progress")
		return false, fmt.Errorf("UpdateNodeCommissionRateByGaslessService is in progress")
	}
	c.logger.WithContext(ctx).Infof("UpdateNodeCommissionRateByGaslessService start........")
	c.cache.Set(common2.UPDATE_NODE_COMMISSION_RATE_BY_GASLESS_SERVICE, true, common2.CACHE_EXPIRED_TIME*time.Second)

	// Call Gasless serivice
	typedData := apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": {
				{
					Name: "name",
					Type: "string",
				},
				{
					Name: "version",
					Type: "string",
				},
				{
					Name: "chainId",
					Type: "uint256",
				},
			},
			"NodeModifyCommissionRateData": {
				{
					Name: "commissionRate",
					Type: "uint32",
				},
				{
					Name: "expiredAt",
					Type: "uint256",
				},
			},
		},
		PrimaryType: "NodeModifyCommissionRateData",
		Domain: apitypes.TypedDataDomain{
			Name:    c.cf.Signature.DomainName,
			Version: c.cf.Signature.DomainVersion,
			ChainId: (*math.HexOrDecimal256)(big.NewInt(c.cf.Chain.ChainId)),
		},
		Message: apitypes.TypedDataMessage{
			"commissionRate": strconv.Itoa(int(commissionRate)),
			"expiredAt":      expiredAt,
		},
	}
	v, r, s, err := tools.SignTypedDataAndSplit(typedData, c.verifierPrivKey)
	if err != nil {
		log.Errorf("SignTypedData error: %s", err.Error())
		return false, err
	}
	// Send signature
	setRes, err := c.gaslessClient.ExplorerSendTxModifyCommissionRate(ctx, &gasless.ExplorerSendTxModifyCommissionRateRequest{
		Signer:         c.verifierAddress.String(),
		CommissionRate: commissionRate,
		ExpiredAt:      expiredAt.Uint64(),
		V:              uint32(v),
		R:              hex.EncodeToString(r[:]),
		S:              hex.EncodeToString(s[:]),
	})
	if err != nil {
		return false, err
	}
	if setRes.Code != 0 {
		return false, fmt.Errorf("UpdateNodeCommissionRateByGaslessService failed %s", setRes.Msg)
	}
	return true, nil
}

// Update node reward claimer by service
func UpdateNodeRewardClaimerByGaslessService(ctx context.Context, c *Chain, rewardClaimer common.Address, expiredAt *big.Int) (bool, error) {
	_, found := c.cache.Get(common2.UPDATE_NODE_REWARD_CLAIMER_BY_GASLESS_SERVICE)
	if found {
		c.logger.WithContext(ctx).Errorf("UpdateNodeRewardClaimerByGaslessService is in progress")
		return false, fmt.Errorf("UpdateNodeRewardClaimerByGaslessService is in progress")
	}
	c.logger.WithContext(ctx).Infof("UpdateNodeRewardClaimerByGaslessService start........")
	c.cache.Set(common2.UPDATE_NODE_REWARD_CLAIMER_BY_GASLESS_SERVICE, true, common2.CACHE_EXPIRED_TIME*time.Second)
	// TODO Call Gasless serivice
	typedData := apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": {
				{
					Name: "name",
					Type: "string",
				},
				{
					Name: "version",
					Type: "string",
				},
				{
					Name: "chainId",
					Type: "uint256",
				},
			},
			"NodeSetRewardClaimerData": {
				{
					Name: "claimer",
					Type: "address",
				},
				{
					Name: "expiredAt",
					Type: "uint256",
				},
			},
		},
		PrimaryType: "NodeSetRewardClaimerData",
		Domain: apitypes.TypedDataDomain{
			Name:    c.cf.Signature.DomainName,
			Version: c.cf.Signature.DomainVersion,
			ChainId: (*math.HexOrDecimal256)(big.NewInt(c.cf.Chain.ChainId)),
		},
		Message: apitypes.TypedDataMessage{
			"claimer":   rewardClaimer.String(),
			"expiredAt": expiredAt,
		},
	}

	v, r, s, err := tools.SignTypedDataAndSplit(typedData, c.verifierPrivKey)
	if err != nil {
		return false, err
	}
	explorerSendTxSetRewardClaimerRequest := &gasless.ExplorerSendTxSetRewardClaimerRequest{
		Signer:    c.verifierAddress.String(),
		Claimer:   rewardClaimer.String(),
		ExpiredAt: expiredAt.Uint64(),
		V:         uint32(v),
		R:         hex.EncodeToString(r[:]),
		S:         hex.EncodeToString(s[:]),
	}
	// Send signature
	setRes, err := c.gaslessClient.ExplorerSendTxSetRewardClaimer(ctx, explorerSendTxSetRewardClaimerRequest)
	if err != nil {
		return false, err
	}
	if setRes.Code != 0 {
		return false, fmt.Errorf("UpdateNodeRewardClaimerByGaslessService failed %s", setRes.Msg)
	}
	return true, nil
}

// Node Exit by service
func NodeExitByGaslessService(ctx context.Context, c *Chain, expiredAt *big.Int) {
	_, found := c.cache.Get(common2.NODE_EXIT_BY_GASLESS_SERVICE)
	if found {
		c.logger.WithContext(ctx).Errorf("NodeExitByGaslessService is in progress")
		return
	}
	c.logger.WithContext(ctx).Infof("NodeExitByGaslessService start........")
	c.cache.Set(common2.NODE_EXIT_BY_GASLESS_SERVICE, true, common2.CACHE_EXPIRED_TIME*time.Second)
	// Call Gasless serivice
	typedData := apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": {
				{
					Name: "name",
					Type: "string",
				},
				{
					Name: "version",
					Type: "string",
				},
				{
					Name: "chainId",
					Type: "uint256",
				},
			},
			"NodeExitData": {
				{
					Name: "expiredAt",
					Type: "uint256",
				},
			},
		},
		PrimaryType: "NodeExitData",
		Domain: apitypes.TypedDataDomain{
			Name:    c.cf.Signature.DomainName,
			Version: c.cf.Signature.DomainVersion,
			ChainId: (*math.HexOrDecimal256)(big.NewInt(c.cf.Chain.ChainId)),
		},
		Message: apitypes.TypedDataMessage{
			"expiredAt": expiredAt,
		},
	}

	v, r, s, err := tools.SignTypedDataAndSplit(typedData, c.verifierPrivKey)
	if err != nil {
		return
	}
	// Send signature
	fmt.Println(v, r, s)

}

// Node Enter by service
func NodeEnterByGaslessService(ctx context.Context, c *Chain, replacedNode common.Address, expiredAt *big.Int) (bool, error) {
	_, found := c.cache.Get(common2.NODE_ENTRY_BY_GASLESS_SERVICE)
	if found {
		c.logger.WithContext(ctx).Errorf("NodeEnterByGaslessService is in progress")
		return false, fmt.Errorf("NodeEnterByGaslessService is in progress")
	}
	c.logger.WithContext(ctx).Infof("NodeEnterByGaslessService start........")
	c.cache.Set(common2.NODE_ENTRY_BY_GASLESS_SERVICE, true, common2.CACHE_EXPIRED_TIME*time.Second)

	// Call Gasless serivice
	typedData := apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": {
				{
					Name: "name",
					Type: "string",
				},
				{
					Name: "version",
					Type: "string",
				},
				{
					Name: "chainId",
					Type: "uint256",
				},
			},
			"NodeEnterData": {
				{
					Name: "replacedNode",
					Type: "address",
				},
				{
					Name: "expiredAt",
					Type: "uint256",
				},
			},
		},
		PrimaryType: "NodeEnterData",
		Domain: apitypes.TypedDataDomain{
			Name:    c.cf.Signature.DomainName,
			Version: c.cf.Signature.DomainVersion,
			ChainId: (*math.HexOrDecimal256)(big.NewInt(c.cf.Chain.ChainId)),
		},
		Message: apitypes.TypedDataMessage{
			"replacedNode": replacedNode.String(),
			"expiredAt":    expiredAt,
		},
	}

	v, r, s, err := tools.SignTypedDataAndSplit(typedData, c.verifierPrivKey)
	if err != nil {
		return false, err
	}
	nodeEnterRequest := &gasless.ExplorerSendTxNodeEnterRequest{
		Signer:       c.verifierAddress.String(),
		ReplacedNode: replacedNode.String(),
		ExpiredAt:    expiredAt.Uint64(),
		V:            uint32(v),
		R:            hex.EncodeToString(r[:]),
		S:            hex.EncodeToString(s[:]),
	}
	enter, err := c.gaslessClient.ExplorerSendTxNodeEnter(ctx, nodeEnterRequest)
	if err != nil {
		return false, err
	}
	c.logger.WithContext(ctx).Infof("NodeEnterByGaslessService success %s", enter)
	if enter.Code != 0 {
		return false, fmt.Errorf("NodeEnterByGaslessService failed %s", enter.Msg)
	}
	return true, err

}

// Node Report Verification Batch by service
func NodeReportVerificationBatchByGaslessService(ctx context.Context, c *Chain, attestationId [32]byte, result uint8, index uint32) (bool, error) {
	_, found := c.cache.Get(common2.NODE_REPORT_VERIFICATION_BATCH_BY_GASLESS_SERVICE)
	if found {
		c.logger.WithContext(ctx).Errorf("NodeReportVerificationBatchByGaslessService is in progress")
		return false, fmt.Errorf("NodeReportVerificationBatchByGaslessService is in progress")

	}
	c.logger.WithContext(ctx).Infof("NodeReportVerificationBatchByGaslessService start........")
	c.cache.Set(common2.NODE_REPORT_VERIFICATION_BATCH_BY_GASLESS_SERVICE, true, common2.CACHE_EXPIRED_TIME*time.Second)
	// Call Gasless serivice
	typedData := apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": {
				{
					Name: "name",
					Type: "string",
				},
				{
					Name: "version",
					Type: "string",
				},
				{
					Name: "chainId",
					Type: "uint256",
				},
			},
			"VerificationData": {
				{
					Name: "attestationID",
					Type: "bytes32",
				},
				{
					Name: "result",
					Type: "uint8",
				},
				{
					Name: "index",
					Type: "uint32",
				},
			},
		},
		PrimaryType: "VerificationData",
		Domain: apitypes.TypedDataDomain{
			Name:    c.cf.Signature.DomainName,
			Version: c.cf.Signature.DomainVersion,
			ChainId: (*math.HexOrDecimal256)(big.NewInt(c.cf.Chain.ChainId)),
		},
		Message: apitypes.TypedDataMessage{
			"attestationID": attestationId[:],
			"result":        strconv.Itoa(int(result)),
			"index":         strconv.Itoa(int(index)),
		},
	}

	v, r, s, err := tools.SignTypedDataAndSplit(typedData, c.verifierPrivKey)
	if err != nil {
		return false, err
	}
	// Send signature
	reportVerificationRequest := &gasless.ExplorerSendTxNodeReportVerificationRequest{
		Signer:        c.verifierAddress.String(),
		AttestationId: hex.EncodeToString(attestationId[:]),
		Result:        uint32(result),
		Index:         index,
		V:             uint32(v),
		R:             hex.EncodeToString(r[:]),
		S:             hex.EncodeToString(s[:]),
	}
	reportVerificationRes, err := c.gaslessClient.ExplorerSendTxNodeReportVerification(ctx, reportVerificationRequest)
	if err != nil {
		return false, err
	}
	if reportVerificationRes.Code != 0 {
		return false, fmt.Errorf("NodeReportVerificationBatchByGaslessService failed %s", reportVerificationRes.Msg)
	}
	return true, nil

}

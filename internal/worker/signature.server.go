package worker

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/carv-protocol/verifier/api/gasless"
	common2 "github.com/carv-protocol/verifier/internal/common"
	"github.com/carv-protocol/verifier/pkg/tools"
)

// Node Exit by service
func NodeExitByGaslessService(ctx context.Context, c *Chain, expiredAt *big.Int) (bool, error) {
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
		return false, err
	}
	nodeExitRequest := &gasless.ExplorerSendTxNodeExitRequest{
		Signer:    c.verifierAddress.String(),
		ExpiredAt: expiredAt.Uint64(),
		V:         uint32(v),
		R:         hex.EncodeToString(r[:]),
		S:         hex.EncodeToString(s[:]),
	}
	if !c.cf.Chain.IsGasless {
		// get Auth
		auth, err := c.GetTransactionConfig(ctx)
		if err != nil {
			c.logger.WithContext(ctx).Errorf("Get Auth failed %s", err.Error())
			return false, err
		}
		// Call NodeExit
		exit, err := c.protocolServiceContractObj.NodeExit(auth)
		if err != nil {
			c.logger.WithContext(ctx).Errorf("NodeExitByGaslessService failed %s", err.Error())
			return false, err
		}
		c.logger.WithContext(ctx).Infof("NodeExitByGaslessService success %s", exit.Hash())
		//TODO Check Transaction result from chain by hash.
		return true, nil
	}
	var exit *gasless.Response
	var exitErr error
	for i := 0; i < 3; i++ {
		c.logger.WithContext(ctx).Infof("NodeExitByGaslessService try %d", i+1)
		exit, exitErr = c.gaslessClient.ExplorerSendTxNodeExit(ctx, nodeExitRequest)
		if exitErr == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if exitErr != nil {
		c.logger.WithContext(ctx).Errorf("NodeExitByGaslessService failed %s", exitErr.Error())
		return false, err
	}
	if exit == nil {
		return false, fmt.Errorf("NodeExitByGaslessService failed")
	}
	if exit.Code != 0 {
		c.logger.WithContext(ctx).Errorf("NodeExitByGaslessService failed %s", exit.Msg)
		os.Exit(0)
	}
	c.logger.WithContext(ctx).Infof("NodeExitByGaslessService success %s", exit)
	return true, err
}

// Node Enter by service
func NodeEnterByGaslessService(ctx context.Context, c *Chain, replacedNode common.Address, expiredAt *big.Int) (bool, error) {

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
	if !c.cf.Chain.IsGasless {
		auth, err := c.GetTransactionConfig(ctx)
		if err != nil {
			c.logger.WithContext(ctx).Errorf("Get Auth failed %s", err.Error())
			return false, err
		}
		enter, err := c.protocolServiceContractObj.NodeEnter(auth, replacedNode)
		if err != nil {
			c.logger.WithContext(ctx).Errorf("NodeEnterByGaslessService failed %s", err.Error())
			os.Exit(0)
		}
		c.logger.WithContext(ctx).Infof("NodeEnterByGaslessService success %s", enter.Hash())
		//TODO Check Transaction result from chain by hash.
		return true, nil
	}

	nodeEnterRequest := &gasless.ExplorerSendTxNodeEnterRequest{
		Signer:       c.verifierAddress.String(),
		ReplacedNode: replacedNode.String(),
		ExpiredAt:    expiredAt.Uint64(),
		V:            uint32(v),
		R:            hex.EncodeToString(r[:]),
		S:            hex.EncodeToString(s[:]),
	}

	var enter *gasless.Response
	var enterErr error
	for i := 0; i < 3; i++ {
		c.logger.WithContext(ctx).Infof("NodeEnterByGaslessService try %d", i+1)
		enter, enterErr = c.gaslessClient.ExplorerSendTxNodeEnter(ctx, nodeEnterRequest)
		if enterErr == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}

	if enterErr != nil {
		c.logger.WithContext(ctx).Errorf("NodeEnterByGaslessService failed %s", enterErr.Error())
		return false, err
	}
	if enter == nil {
		return false, fmt.Errorf("NodeEnterByGaslessService failed")
	}
	if enter.Code != 0 {
		c.logger.WithContext(ctx).Errorf("NodeEnterByGaslessService failed %s", enter.Msg)
		os.Exit(0)
	}
	c.logger.WithContext(ctx).Infof("NodeEnterByGaslessService success %s", enter)

	return true, err

}

// Update node commission rate by service
func UpdateNodeCommissionRateByGaslessService(ctx context.Context, c *Chain, commissionRate uint32, expiredAt *big.Int) (bool, error) {
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
	if !c.cf.Chain.IsGasless {
		auth, err := c.GetTransactionConfig(ctx)
		if err != nil {
			c.logger.WithContext(ctx).Errorf("Get Auth failed %s", err.Error())
			return false, err
		}
		update, err := c.protocolServiceContractObj.NodeModifyCommissionRate(auth, commissionRate)
		if err != nil {
			c.logger.WithContext(ctx).Errorf("UpdateNodeCommissionRateByGaslessService failed %s", err.Error())
			os.Exit(0)
		}
		c.logger.WithContext(ctx).Infof("UpdateNodeCommissionRateByGaslessService success %s", update.Hash())
		//TODO Check Transaction result from chain by hash.
		return true, nil
	}
	var setRes *gasless.Response
	var setErr error
	for i := 0; i < 3; i++ {
		c.logger.WithContext(ctx).Infof("UpdateNodeCommissionRateByGaslessService try %d", i+1)
		setRes, setErr = c.gaslessClient.ExplorerSendTxModifyCommissionRate(ctx, &gasless.ExplorerSendTxModifyCommissionRateRequest{
			Signer:         c.verifierAddress.String(),
			CommissionRate: commissionRate,
			ExpiredAt:      expiredAt.Uint64(),
			V:              uint32(v),
			R:              hex.EncodeToString(r[:]),
			S:              hex.EncodeToString(s[:]),
		})
		if setErr == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}

	if setErr != nil {
		c.logger.WithContext(ctx).Errorf("UpdateNodeCommissionRateByGaslessService failed %s", setErr.Error())
		return false, err
	}
	if setRes == nil {
		return false, fmt.Errorf("UpdateNodeCommissionRateByGaslessService failed")
	}
	if setRes.Code != 0 {
		c.logger.WithContext(ctx).Errorf("UpdateNodeCommissionRateByGaslessService failed %s", setRes.Msg)
		os.Exit(0)
	}
	c.logger.WithContext(ctx).Infof("UpdateNodeCommissionRateByGaslessService success %s", setRes)

	return true, nil
}

// Update node reward claimer by service
func UpdateNodeRewardClaimerByGaslessService(ctx context.Context, c *Chain, rewardClaimer common.Address, expiredAt *big.Int) (bool, error) {
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
	if !c.cf.Chain.IsGasless {
		auth, err := c.GetTransactionConfig(ctx)
		if err != nil {
			c.logger.WithContext(ctx).Errorf("Get Auth failed %s", err.Error())
			return false, err
		}
		update, err := c.protocolServiceContractObj.NodeSetRewardClaimerWithSignature(auth, rewardClaimer, expiredAt, c.verifierAddress, v, r, s)
		if err != nil {
			c.logger.WithContext(ctx).Errorf("UpdateNodeRewardClaimerByGaslessService failed %s", err.Error())
			os.Exit(0)
		}
		c.logger.WithContext(ctx).Infof("UpdateNodeRewardClaimerByGaslessService success %s", update.Hash())
		//TODO Check Transaction result from chain by hash.
		return true, nil
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
	var setRes *gasless.Response
	var setErr error
	for i := 0; i < 3; i++ {
		c.logger.WithContext(ctx).Infof("UpdateNodeRewardClaimerByGaslessService try %d", i+1)
		setRes, setErr = c.gaslessClient.ExplorerSendTxSetRewardClaimer(ctx, explorerSendTxSetRewardClaimerRequest)
		if setErr == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if setErr != nil {
		c.logger.WithContext(ctx).Errorf("UpdateNodeRewardClaimerByGaslessService failed %s", setErr.Error())
		return false, err
	}
	if setRes == nil {
		return false, fmt.Errorf("UpdateNodeRewardClaimerByGaslessService failed")
	}
	if setRes.Code != 0 {
		c.logger.WithContext(ctx).Errorf("UpdateNodeRewardClaimerByGaslessService failed %s", setRes.Msg)
		os.Exit(0)
	}
	c.logger.WithContext(ctx).Infof("UpdateNodeRewardClaimerByGaslessService success %s", setRes)
	return true, nil
}

// Node Report Verification Batch by service
func NodeReportVerificationBatchByGaslessService(ctx context.Context, c *Chain, attestationId [32]byte, result uint8, index uint32) (bool, error) {
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
	if !c.cf.Chain.IsGasless {
		auth, err := c.GetTransactionConfig(ctx)
		if err != nil {
			c.logger.WithContext(ctx).Errorf("Get Auth failed %s", err.Error())
			return false, err
		}
		update, err := c.protocolServiceContractObj.NodeReportVerification(auth, attestationId, uint32(result), uint8(index))
		if err != nil {
			c.logger.WithContext(ctx).Errorf("NodeReportVerificationBatchByGaslessService failed %s", err.Error())
			os.Exit(0)
		}
		c.logger.WithContext(ctx).Infof("NodeReportVerificationBatchByGaslessService success %s", update.Hash())
		//TODO Check Transaction result from chain by hash.
		return true, nil
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
	var reportVerificationRes *gasless.Response
	var reportVerificationErr error
	for i := 0; i < 3; i++ {
		c.logger.WithContext(ctx).Infof("NodeReportVerificationBatchByGaslessService try %d", i+1)
		reportVerificationRes, reportVerificationErr = c.gaslessClient.ExplorerSendTxNodeReportVerification(ctx, reportVerificationRequest)
		if reportVerificationErr == nil {
			c.logger.WithContext(ctx).Infof("NodeReportVerificationBatchByGaslessService failed %s", reportVerificationRes)
			break
		}
		time.Sleep(2 * time.Second)
	}
	if reportVerificationErr != nil {
		c.logger.WithContext(ctx).Errorf("NodeReportVerificationBatchByGaslessService failed %s", reportVerificationErr.Error())
		return false, err
	}
	if reportVerificationRes == nil {
		return false, fmt.Errorf("NodeReportVerificationBatchByGaslessService failed")
	}
	if reportVerificationRes.Code != 0 {
		return false, fmt.Errorf("NodeReportVerificationBatchByGaslessService failed %s", reportVerificationRes.Msg)
	}
	c.logger.WithContext(ctx).Infof("NodeReportVerificationBatchByGaslessService success %s", reportVerificationRes)
	return true, nil

}

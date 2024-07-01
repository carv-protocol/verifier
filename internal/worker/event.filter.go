package worker

import (
	"context"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/carv-protocol/verifier/pkg/dcap"
	"github.com/carv-protocol/verifier/pkg/tools"
)

type LogFilter struct {
}

func (l *LogFilter) NodeReportVerificationBatchLogFilter(ctx context.Context, c *Chain, cLog types.Log) error {
	unpackedData, err := c.protocolServiceContractObj.ParseTeeReportAttestations(cLog)
	if err != nil {
		return errors.Wrap(err, "contract ParseReportTeeAttestation error")
	}

	// Verify attestation
	for i := 0; i < len(unpackedData.AttestationIDs); i++ {
		attestationID := unpackedData.AttestationIDs[i]
		attestationInfo := unpackedData.AttestationInfos[i]
		// Verify attestation
		identityJson, tcbJson, trustedJson, err := tools.GetDcapConfigJsonString()
		if err != nil {
			c.logger.WithContext(ctx).Error("get dcap config json string error: %s", err.Error())
			return err
		}
		result, err := dcap.VerifyAttestation(attestationInfo, identityJson, tcbJson, trustedJson)
		if err != nil {
			// If attestation is unable to be parsed and verified, this attestation should be ignored by all verifiers
			c.logger.WithContext(ctx).Error(
				"verify failed, attestation id: %s, error: %s",
				attestationID,
				err.Error(),
			)
		}

		c.verifyResultList = append(c.verifyResultList, verifyResult{
			BlockNumber:     cLog.BlockNumber,
			ContractAddress: cLog.Address,
			TxHash:          cLog.TxHash,
			TxIndex:         unpackedData.Raw.TxIndex,
			requestId:       unpackedData.RequestID,
			attestationID:   attestationID,
			result:          result,
			isReported:      false,
		})
		c.logger.WithContext(ctx).Infof("verify %v result: %v", attestationID, result)
	}

	return nil
}

func (l *LogFilter) ConfirmVrfNodesLogFilter(ctx context.Context, c *Chain, cLog types.Log) error {
	unpackedData, err := c.protocolServiceContractObj.ParseConfirmVrfNodes(cLog)
	if err != nil {
		return errors.Wrap(err, "contract ParseConfirmVrfNodes error")
	}
	if len(unpackedData.VrfChosen) == 0 {
		c.logger.WithContext(ctx).Infof("no vrf node chosen")
		return nil
	}
	logInfo := ConfirmVrfNodesEvent{
		BlockNumber:     cLog.BlockNumber,
		ContractAddress: cLog.Address,
		TxHash:          cLog.TxHash,
		TxIndex:         unpackedData.Raw.TxIndex,
		RequestId:       unpackedData.RequestId,
		VrfChosen:       unpackedData.VrfChosen,
		Deadline:        unpackedData.Deadline,
	}
	searchNodeIndex := -1
	for i := 0; i < len(unpackedData.VrfChosen); i++ {
		if unpackedData.VrfChosen[i] == c.nodeInf.nodeId {
			searchNodeIndex = i
			break
		}
	}
	c.logger.WithContext(ctx).Infof("searchResBool: %v", searchNodeIndex)
	if searchNodeIndex == -1 {
		return nil
	}

	c.confirmVrfNodeChan <- confirmVrfNodesInfo{
		nodeId:       c.nodeInf.nodeId,
		vrfNodeIndex: uint32(searchNodeIndex),
		requestId:    unpackedData.RequestId,
	}

	c.logger.WithContext(ctx).Infof("logInfo: %+v", logInfo)
	return nil
}

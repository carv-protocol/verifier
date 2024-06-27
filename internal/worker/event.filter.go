package worker

import (
	"context"
	"github.com/pkg/errors"
	"sort"

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

	logInfo := TeeReportAttestationEvent{
		BlockNumber:       cLog.BlockNumber,
		ContractAddress:   cLog.Address,
		TxHash:            cLog.TxHash,
		TxIndex:           unpackedData.Raw.TxIndex,
		AttestationIDs:    unpackedData.AttestationIDs,
		VerificationInfos: unpackedData.AttestationInfos,
		RequestId:         unpackedData.RequestID,
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
		c.verifyResultMap.Set(unpackedData.RequestID.String(), verifyResult{
			attestationID: attestationID,
			result:        result,
		})

		c.logger.WithContext(ctx).Infof("logInfo: %+v", logInfo)
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
	// check whether the vrf node is in the vrf node list
	// vrfChosen is node id
	sort.Slice(unpackedData.VrfChosen, func(i, j int) bool {
		return unpackedData.VrfChosen[i] < (unpackedData.VrfChosen[j])
	})
	//
	searchResBool := tools.BinarySearch(unpackedData.VrfChosen, c.nodeInf.nodeId, func(u uint32, u2 uint32) bool {
		return u < u2
	})
	c.logger.WithContext(ctx).Error("searchResBool: %v", searchResBool)
	if searchResBool == -1 {
		return nil
	}

	c.confirmVrfNodeChan <- confirmVrfNodesInfo{
		nodeId:    c.nodeInf.nodeId,
		requestId: unpackedData.RequestId,
	}

	c.logger.WithContext(ctx).Infof("logInfo: %+v", logInfo)
	return nil
}

package worker

import (
	"context"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

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
			c.logger.WithContext(ctx).Errorf(
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
	}

	return nil
}

func (l *LogFilter) ConfirmVrfNodesLogFilter(ctx context.Context, c *Chain, cLog types.Log) error {
	unpackedData, err := c.protocolServiceContractObj.ParseConfirmVrfNodes(cLog)
	if err != nil {
		return errors.Wrap(err, "contract ParseConfirmVrfNodes error")
	}
	if len(unpackedData.VrfChosen) == 0 {
		c.logger.WithContext(ctx).Info("no vrf node chosen")
		return nil
	}
	searchNodeIndex := -1
	for i := 0; i < len(unpackedData.VrfChosen); i++ {
		if unpackedData.VrfChosen[i] == c.nodeInf.nodeId {
			searchNodeIndex = i
			break
		}
	}
	if searchNodeIndex == -1 {
		return nil
	}

	c.confirmVrfNodeChan <- confirmVrfNodesInfo{
		nodeId:       c.nodeInf.nodeId,
		vrfNodeIndex: uint32(searchNodeIndex),
		requestId:    unpackedData.RequestId,
		deadline:     unpackedData.Deadline,
	}

	return nil
}

func (l *LogFilter) NodeClearLogFilter(ctx context.Context, c *Chain, log types.Log) error {
	unpackedData, err := c.protocolServiceContractObj.ParseNodeClear(log)
	if err != nil {
		c.logger.WithContext(ctx).Errorf("contract ParseNodeClear error: %s", err.Error())
		return errors.Wrap(err, "contract ParseNodeClear error")
	}
	// add now active check
	nodeInfos, err := c.protocolServiceContractObj.NodeInfos(&bind.CallOpts{}, unpackedData.Node)
	if err != nil {
		return err
	}
	if nodeInfos.Active {
		c.logger.WithContext(ctx).Infof("node %s is active, ignore node clear event", unpackedData.Node.String())
		return nil
	}
	if strings.ToLower(unpackedData.Node.String()) == strings.ToLower(c.verifierAddress.String()) {
		c.exitNodeChan <- struct{}{}
	}
	return nil
}

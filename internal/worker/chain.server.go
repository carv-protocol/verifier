package worker

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func NodeEnter(ctx context.Context, c *Chain, auth *bind.TransactOpts, replaceNodeAddr common.Address) (common.Hash, error) {

	enter, err := c.protocolServiceContractObj.NodeEnter(auth, replaceNodeAddr)
	if err != nil {
		return [32]byte{}, err
	}
	return enter.Hash(), nil
}

func NodeExit(ctx context.Context, c *Chain, auth *bind.TransactOpts) (common.Hash, error) {

	exit, err := c.protocolServiceContractObj.NodeExit(auth)
	if err != nil {
		return [32]byte{}, err
	}
	return exit.Hash(), nil
}

func UpdateNodeCommissionRate(ctx context.Context, c *Chain, auth *bind.TransactOpts, commissionRate uint32) (common.Hash, error) {

	update, err := c.protocolServiceContractObj.NodeModifyCommissionRate(auth, commissionRate)
	if err != nil {
		return [32]byte{}, err
	}
	return update.Hash(), nil
}

func NodeReportVerification(ctx context.Context, c *Chain, auth *bind.TransactOpts, attestationId [32]byte, index uint32, result uint8) (common.Hash, error) {

	report, err := c.protocolServiceContractObj.NodeReportVerification(auth, attestationId, index, result)
	if err != nil {
		return [32]byte{}, err
	}
	return report.Hash(), nil

}

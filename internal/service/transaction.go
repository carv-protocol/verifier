package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	verifierContract "github.com/carv-protocol/verifier/pkg/contract"
)

type TransactionService struct {
}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (s *TransactionService) CreateTransaction() {
	client, err := ethclient.Dial("")

	if err != nil {
		log.Fatal(err)
	}

	// get Private
	privateKey, err := crypto.HexToECDSA("")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	address := common.HexToAddress("")

	instance, err := verifierContract.NewContract(address, client)

	if err != nil {
		log.Fatal(err)
	}

	tx, err := instance.ReportTeeAttestation(auth, "", "")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx send: %s ", tx.Hash().Hex())
}

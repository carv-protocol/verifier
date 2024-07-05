package tools

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common/math"
	"math/big"
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func TestSignTypedData(t *testing.T) {
	attestationIdArray := "0xf3f8a84bf2661f5b1ded6e23f820085ec3610274547dfd5639883e58dc0f19ba"
	// Step 1: Initialize Test Data
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
			Name:    "Test",
			Version: "1.0.0",
			ChainId: (*math.HexOrDecimal256)(big.NewInt(1)),
		},
		Message: apitypes.TypedDataMessage{
			"attestationID": attestationIdArray,
			"result":        strconv.Itoa(1),
			"index":         strconv.Itoa(1),
		},
	}

	// Step 2: Generate a Private Key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	// Step 3: Call SignTypedData
	sig, err := SignTypedData(typedData, privateKey)
	if err != nil {
		t.Fatalf("SignTypedData failed: %v", err)
	}

	// Step 4: Assert the Outcome
	if len(sig) == 0 {
		t.Errorf("Expected a signature, got none")
	}

	// public key
	publicKey := privateKey.Public()
	publicKeyECDA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Fatalf("Failed to convert public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDA)

	hash, err := EncodeForSigning(typedData)
	if err != nil {
		t.Fatalf("EncodeForSigning failed: %v", err)
	}
	if !VerifySig(address.String(), "0x"+hex.EncodeToString(sig), hash.Bytes()) {
		t.Errorf("VerifySig failed")
	}
}

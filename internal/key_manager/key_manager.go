package key_manager

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/carv-protocol/verifier/internal/conf"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"os"
)

var keyManager *KeyManager

type KeyManager struct {
	priKey *ecdsa.PrivateKey
}

func GenerateKeystore(keystorePath string) {
	if len(keystorePath) == 0 {
		fmt.Printf("keystore path is needed")
		os.Exit(1)
	}
	password := utils.GetPassPhrase("Your new account is locked with a password. Please give a password. Do not forget this password.", true)
	ks := keystore.NewKeyStore(keystorePath, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		fmt.Printf("keystore NewAccount err: %v", err)
		os.Exit(1)
	}
	fmt.Printf("keystore NewAccount: %v", account)
}

func NewKeyManager(walletConf *conf.Wallet, envPrivateKey, envKeystorePath, envKeystorePassword string) (err error) {
	keyManager, err = newKeyManager(walletConf, envPrivateKey, envKeystorePath, envKeystorePassword)
	return err
}

func newKeyManager(walletConf *conf.Wallet, envPrivateKey, envKeystorePath, envKeystorePassword string) (*KeyManager, error) {
	switch walletConf.Mode {
	case conf.Wallet_ModeEnvVar:
		if len(envPrivateKey) == 0 && len(envKeystorePath) == 0 {
			return nil, fmt.Errorf("env private key or keystore path is needed")
		}
		if len(envPrivateKey) != 0 {
			return fromPrivateKey(envPrivateKey)
		}
		return fromKeystore(envKeystorePath, envKeystorePassword)

	case conf.Wallet_ModePrivateKey:
		return fromPrivateKey(walletConf.PrivateKey)

	case conf.Wallet_ModeKeyStore:
		return fromKeystore(walletConf.KeystorePath, walletConf.KeystorePassword)

	default:
		return nil, errors.New("unknown wallet mode")
	}
}

func fromPrivateKey(privateKey string) (*KeyManager, error) {
	priKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "HexToECDSA")
	}
	return &KeyManager{priKey: priKey}, nil
}

func fromKeystore(keystorePath, keystorePassword string) (*KeyManager, error) {
	keystoreFile, err := os.ReadFile(keystorePath)
	if err != nil {
		return nil, errors.Wrap(err, "read keystore file")
	}
	key, err := keystore.DecryptKey(keystoreFile, keystorePassword)
	if err != nil {
		return nil, errors.Wrap(err, "decrypt keystore file")
	}
	return &KeyManager{key.PrivateKey}, nil
}

func (km KeyManager) PrivateKey() *ecdsa.PrivateKey {
	return km.priKey
}

func Inst() *KeyManager {
	return keyManager
}

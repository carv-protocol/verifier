package worker

import (
	"bytes"
	"crypto/cipher"
	"encoding/hex"

	"github.com/tjfoc/gmsm/sm4"

	"github.com/carv-protocol/verifier/internal/conf"
)

func Sm4Encrypt(key, iv, plainText []byte) ([]byte, error) {
	block, err := sm4.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData := pkcs5Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cryted := make([]byte, len(origData))
	blockMode.CryptBlocks(cryted, origData)

	return cryted, nil
}

func Sm4Decrypt(cf *conf.Bootstrap) ([]byte, error) {
	key := []byte(cf.Wallet.Key)
	iv := []byte(cf.Wallet.Iv)
	private_encode := cf.Wallet.PrivateEncode

	cipherText, err := hex.DecodeString(private_encode)
	if err != nil {
		return nil, err
	}

	block, err := sm4.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(origData, cipherText)
	origData = pkcs5UnPadding(origData)

	return origData, nil
}

// pkcs5 padding
func pkcs5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(src, padtext...)
}

func pkcs5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])

	return src[:(length - unpadding)]
}

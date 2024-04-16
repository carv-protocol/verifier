package sm4

import (
	"bytes"
	"crypto/cipher"
	"github.com/tjfoc/gmsm/sm4"
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

func Sm4Decrypt(key, iv, cipherText []byte) ([]byte, error) {
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

//func main() {
//	// 128 bit key
//	key := []byte("1234567890abcdef")
//	// 128 bit iv
//	iv := make([]byte, sm4.BlockSize)
//	data := []byte("hello SM4")
//	fmt.Println("SM4 encode text ：", string(data))
//	ciphertxt, err := sm4Encrypt(key, iv, data)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("SM4 encode result: %x\n", ciphertxt)
//	fmt.Println("SM4 encode result :\n", base64.StdEncoding.EncodeToString(ciphertxt))
//
//	res, err := sm4Decrypt(key, iv, ciphertxt)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("SM4 decode text ：", string(res))
//
//}

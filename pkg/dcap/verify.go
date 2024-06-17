package dcap

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/andybalholm/brotli"
	"github.com/carv-protocol/verifier/internal/common"
	"github.com/carv-protocol/verifier/internal/conf"
	"io"
	"io/ioutil"
	"net/http"
)

func VerifyAttestation(data string, cf *conf.Bootstrap) (bool, error) {

	b64Data, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return false, err
	}
	quoteByte, err := decompressDataWithBrotli(b64Data)
	if err != nil {
		return false, err
	}

	var quote = Quote{}
	var byteReader = bytes.NewReader(quoteByte)
	err = binary.Read(byteReader, binary.BigEndian, &quote)
	if err != nil {
		return false, err
	}
	quoteAuth, err := GetQuoteV3Auth(quoteByte)
	if err != nil {
		return false, err
	}

	result, err := TrustedLoadFromUrl(common.TRUSTED_PATH)
	if err != nil {
		return false, err
	}
	err = quote.VerifyQuote(quoteByte, &result, &quoteAuth, cf)
	if err != nil {
		return false, err
	}
	return true, nil

}

type TrusTEEInfo struct {
	TrustedEnclaves map[string][]string
}

func TrustedLoadFromUrl(url string) (TrusTEEInfo, error) {
	resp, err := http.Get(url)
	if err != nil {
		return TrusTEEInfo{}, fmt.Errorf("error getting JSON file: %w", err)
	}
	defer resp.Body.Close()
	byteValue, err := io.ReadAll(resp.Body)
	if err != nil {
		return TrusTEEInfo{}, fmt.Errorf("error reading JSON file: %w", err)
	}
	var info TrusTEEInfo
	if err := json.Unmarshal(byteValue, &info); err != nil {
		return TrusTEEInfo{}, fmt.Errorf("error decoding JSON: %w", err)
	}

	return info, nil
}

func decompressDataWithBrotli(compressedData []byte) ([]byte, error) {
	reader := brotli.NewReader(bytes.NewReader(compressedData))
	decompressedData, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return decompressedData, nil
}

package dcap

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/carv-protocol/verifier/internal/common"
	"github.com/carv-protocol/verifier/internal/conf"
	"io"
	"net/http"
	"os"
)

func VerifyAttestation(data string, cf *conf.Bootstrap) (bool, error) {
	b64Data, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return false, err
	}

	var quote = Quote{}
	var byteReader = bytes.NewReader(b64Data)
	err = binary.Read(byteReader, binary.BigEndian, &quote)
	if err != nil {
		return false, err
	}
	quoteAuth, err := GetQuoteV3Auth(b64Data)
	if err != nil {
		return false, err
	}

	result, err := TrustedLoadFromUrl(common.BASE_URl + common.TRUSTED_PATH)
	if err != nil {
		return false, err
	}
	err = quote.VerifyQuote(b64Data, &result, &quoteAuth, cf)
	if err != nil {
		return false, err
	}
	return true, nil
}

type TrusTEEInfo struct {
	TrustedEnclaves map[string][]string
}

func TrustedLoad(path string) (TrusTEEInfo, error) {
	file, err := os.Open(path)
	if err != nil {
		return TrusTEEInfo{}, fmt.Errorf("error opening JSON file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return TrusTEEInfo{}, fmt.Errorf("error reading JSON file: %w", err)
	}

	var info TrusTEEInfo
	if err := json.Unmarshal(byteValue, &info); err != nil {
		return TrusTEEInfo{}, fmt.Errorf("error decoding JSON: %w", err)
	}

	return info, nil
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

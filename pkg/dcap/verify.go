package dcap

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func VerifyAttestation(data string) (bool, error) {
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

	err = quote.VerifyQuote(b64Data, nil, &quoteAuth)
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

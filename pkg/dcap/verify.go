package dcap

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
)

func VerifyAttestation(data string, identity, tcb, trusted string) (bool, error) {
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

	result, err := TrustedLoad(trusted)
	if err != nil {
		return false, err
	}
	err = quote.VerifyQuote(b64Data, &result, &quoteAuth, identity, tcb)
	if err != nil {
		return false, err
	}
	return true, nil
}

type TrusTEEInfo struct {
	TrustedEnclaves map[string][]string
}

func TrustedLoad(trusted string) (TrusTEEInfo, error) {
	var info TrusTEEInfo
	if err := json.Unmarshal([]byte(trusted), &info); err != nil {
		return TrusTEEInfo{}, fmt.Errorf("error decoding JSON: %w", err)
	}

	return info, nil
}

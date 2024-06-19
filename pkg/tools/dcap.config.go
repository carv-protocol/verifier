package tools

import (
	"fmt"
	"github.com/carv-protocol/verifier/internal/common"
	"io"
	"net/http"
)

// loadEnclaveId loads and parses the enclave ID from a JSON file.
// get iden
func LoadEnclaveIdFromUrl(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func TrustedLoadFromUrl(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error getting JSON file: %w", err)
	}
	defer resp.Body.Close()
	byteValue, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading JSON file: %w", err)
	}
	return string(byteValue), nil
}

func LoadTcbInfoFromUrl(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func GetDcapConfigJsonString() (string, string, string, error) {
	identityJsonStr, err := LoadEnclaveIdFromUrl(common.BASE_URl + common.IDENTITY_JSON)
	if err != nil {
		return "", "", "", err
	}
	tcbJsonStr, err := LoadTcbInfoFromUrl(common.BASE_URl + common.TCB_JSON)
	if err != nil {
		return "", "", "", err
	}
	trustedJsonStr, err := TrustedLoadFromUrl(common.BASE_URl + common.TRUSTED_PATH)
	if err != nil {
		return "", "", "", err
	}
	return identityJsonStr, tcbJsonStr, trustedJsonStr, nil
}

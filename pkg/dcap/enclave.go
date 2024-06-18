package dcap

import (
	"encoding/hex"
	"encoding/json"
	"strings"
	"sync"
)

type EnclaveIdStatus string

const (
	OK                                EnclaveIdStatus = "OK"
	SGX_ENCLAVE_REPORT_ISVSVN_REVOKED EnclaveIdStatus = "SGX_ENCLAVE_REPORT_ISVSVN_REVOKED"
)

type TcbLevel struct {
	Isvsvn    uint16
	TcbStatus EnclaveIdStatus
}

type EnclaveId struct {
	Miscselect     uint32
	MiscselectMask uint32
	Isvprodid      uint16
	Attributes     [16]byte
	AttributesMask [16]byte
	Mrsigner       [32]byte
	TcbLevels      []TcbLevel
}

var readOnly struct {
	sync.Once
	instance *EnclaveId
}

// parseHex parses a hex string into a byte slice.
func parseHex(s string) ([]byte, error) {
	cleaned := strings.Trim(s, `"`)
	return hex.DecodeString(cleaned)
}

func (e *EnclaveId) GetEnclaveID(identity string) *EnclaveId {
	readOnly.Do(func() {
		e, err := loadEnclaveId(identity)
		if err != nil {
			panic(err) // or handle error gracefully
		}
		readOnly.instance = e
	})
	return readOnly.instance
}

// loadEnclaveId loads and parses the enclave ID from a JSON file.
func loadEnclaveId(identityJson string) (*EnclaveId, error) {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal([]byte(identityJson), &raw); err != nil {
		return nil, err
	}

	var enclave EnclaveId
	var identity json.RawMessage
	if err := json.Unmarshal(raw["enclaveIdentity"], &identity); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(identity, &raw); err != nil {
		return nil, err
	}

	// Parse miscselect and miscselectMask
	if miscselect, err := parseHex(string(raw["miscselect"])); err == nil && len(miscselect) == 4 {
		enclave.Miscselect = uint32(miscselect[3]) | uint32(miscselect[2])<<8 | uint32(miscselect[1])<<16 | uint32(miscselect[0])<<24
	}
	if miscselectMask, err := parseHex(string(raw["miscselectMask"])); err == nil && len(miscselectMask) == 4 {
		enclave.MiscselectMask = uint32(miscselectMask[3]) | uint32(miscselectMask[2])<<8 | uint32(miscselectMask[1])<<16 | uint32(miscselectMask[0])<<24
	}
	// Parse Attributes and AttributesMask
	if attributes, err := parseHex(string(raw["attributes"])); err == nil && len(attributes) == 16 {
		copy(enclave.Attributes[:], attributes)
	}
	if attributesMask, err := parseHex(string(raw["attributesMask"])); err == nil && len(attributesMask) == 16 {
		copy(enclave.AttributesMask[:], attributesMask)
	}

	// Parse mrsigner
	if mrsigner, err := parseHex(string(raw["mrsigner"])); err == nil && len(mrsigner) == 32 {
		copy(enclave.Mrsigner[:], mrsigner)
	}

	// Parse isvprodid
	var isvprodid uint64
	if err := json.Unmarshal(raw["isvprodid"], &isvprodid); err == nil {
		enclave.Isvprodid = uint16(isvprodid)
	}

	// Parse TcbLevels
	var levels []map[string]json.RawMessage
	if err := json.Unmarshal(raw["tcbLevels"], &levels); err == nil {
		for _, level := range levels {
			var tcb map[string]uint64
			if err := json.Unmarshal(level["tcb"], &tcb); err != nil {
				continue
			}
			var status string
			err := json.Unmarshal(level["tcbStatus"], &status)
			if err != nil {
				return nil, err
			}
			if status == "UpToDate" {

				status = string(OK)
			} else {
				status = string(SGX_ENCLAVE_REPORT_ISVSVN_REVOKED)
			}
			tcbLevel := TcbLevel{
				Isvsvn:    uint16(tcb["isvsvn"]),
				TcbStatus: EnclaveIdStatus(status),
			}
			enclave.TcbLevels = append(enclave.TcbLevels, tcbLevel)
		}
	}

	return &enclave, nil
}

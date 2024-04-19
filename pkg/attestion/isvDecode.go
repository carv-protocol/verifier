package attestion

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
)

type ECDSAQuoteV3AuthData struct {
	ECDSA256BitSignature [64]byte `json:"ecdsa256BitSignature"`
	ECDSAAttestationKey  [64]byte `json:"ecdsaAttestationKey"`
	QEReport             QEReport `json:"QEReport"`
	QEReportSignature    [64]byte `json:"qeReportSignature"`
	QEAuthData           []byte   `json:"qeAuthData"`
	QECert               QECert   `json:"qeCert"`
}

type QECert struct {
	CertType uint16 `json:"certType"`
	Data     []byte `json:"Data"`
}

type QEReport struct {
	CPUSVN     [16]byte `json:"cpusvn"`
	MiscSelect uint32   `json:"miscSelect"`
	Reserved1  [28]byte `json:"reserved1"`
	Attributes [16]byte `json:"attributes"`
	MREnclave  [32]byte `json:"mrEnclave"`
	Reserved2  [32]byte `json:"reserved2"`
	MRSigner   [32]byte `json:"mrSigner"`
	Reserved3  [96]byte `json:"reserved3"`
	ISVProdID  uint16   `json:"isvProdID"`
	ISVSVN     uint16   `json:"isvSVN"`
	Reserved4  [60]byte `json:"reserved4"`
	ReportData [64]byte `json:"reportData"`
}

func DecodeFromBytes(bytesStr string) (ECDSAQuoteV3AuthData, error) {
	// base64 decode
	bytes, err := base64.StdEncoding.DecodeString(bytesStr)
	fmt.Println(hex.EncodeToString(bytes))
	if err != nil {
		return ECDSAQuoteV3AuthData{}, err
	}
	if len(bytes) < 576 {
		return ECDSAQuoteV3AuthData{}, errors.New("Invalid length for ECDSA Quote V3 Auth Data")
	}

	var ecdsaQuote ECDSAQuoteV3AuthData

	ecdsaQuote.QEReport.MREnclave = *(*[32]byte)(bytes[112:144])

	bytes = bytes[576:]
	quAuthDataSize := binary.LittleEndian.Uint16(bytes[0:2])

	bytes = bytes[2:]
	if len(bytes) < int(quAuthDataSize) {
		return ECDSAQuoteV3AuthData{}, errors.New("Invalid QE auth data size")
	}
	ecdsaQuote.QEAuthData = bytes[:quAuthDataSize]

	return ecdsaQuote, nil
}

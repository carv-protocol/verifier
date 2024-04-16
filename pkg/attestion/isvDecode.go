package attestion

import (
	"encoding/binary"
	"github.com/pkg/errors"
)

type ECDSAQuoteV3AuthData struct {
	ECDSA256BitSignature [64]byte  `json:"ecdsa256BitSignature"`
	ECDSAAttestationKey  [64]byte  `json:"ecdsaAttestationKey"`
	QEReport             [384]byte `json:"QEReport"`
	QEReportSignature    [64]byte  `json:"qeReportSignature"`
	QEAuthData           []byte    `json:"qeAuthData"`
	QECert               QECert    `json:"qeCert"`
}

type QECert struct {
	CertType uint16 `json:"certType"`
	Data     []byte `json:"Data"`
}

func DecodeFromBytes(bytes []byte) (ECDSAQuoteV3AuthData, error) {
	if len(bytes) < 576 {
		return ECDSAQuoteV3AuthData{}, errors.New("Invalid length for ECDSA Quote V3 Auth Data")
	}

	var ecdsaQuote ECDSAQuoteV3AuthData

	ecdsaQuote.ECDSA256BitSignature = *(*[64]byte)(bytes[0:64])
	ecdsaQuote.ECDSAAttestationKey = *(*[64]byte)(bytes[64:128])
	ecdsaQuote.QEReport = *(*[384]byte)(bytes[128:512])
	ecdsaQuote.QEReportSignature = *(*[64]byte)(bytes[512:576])

	bytes = bytes[576:]
	quAuthDataSize := binary.LittleEndian.Uint16(bytes[0:2])

	bytes = bytes[2:]
	if len(bytes) < int(quAuthDataSize) {
		return ECDSAQuoteV3AuthData{}, errors.New("Invalid QE auth data size")
	}
	ecdsaQuote.QEAuthData = bytes[:quAuthDataSize]

	// Assuming QECertData has a similar fromBytes function for parsing its data
	// qeCert, err := QECertDataFromBytes(bytes[quAuthDataSize:])
	// if err != nil {
	//     return ECDSAQuoteV3AuthData{}, err
	// }
	// ecdsaQuote.QECert = qeCert

	return ecdsaQuote, nil
}

package worker

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
	Reserved1  uint8    `json:"reserved1"`
	ISVProdID  uint16   `json:"isvProdID"`
	Attributes [16]byte `json:"attributes"`
	MREnclave  [32]byte `json:"mrEnclave"`
	Reserved2  [32]byte `json:"reserved2"`
	MRSigner   [32]byte `json:"mrSigner"`
	Reserved3  [96]byte `json:"reserved3"`
	ISVSVN     uint16   `json:"isvSVN"`
	Reserved4  [60]byte `json:"reserved4"`
	ReportData [64]byte `json:"reportData"`
}

func verifyAttestation(c *Chain, bytesStr string) (bool, error) {
	// base64 decode
	bytes, err := base64.StdEncoding.DecodeString(bytesStr)
	if err != nil {
		return false, errors.Wrapf(err, "Failed to decode base64 string")
	}
	if len(bytes) < 576 {
		return false, errors.New("Invalid length for ECDSA Quote V3 Auth Data")
	}

	var ecdsaQuote ECDSAQuoteV3AuthData
	//ecdsaQuote.QEReport.CPUSVN = *(*[16]byte)(bytes[12:28])
	//ecdsaQuote.QEReport.MiscSelect = binary.LittleEndian.Uint32(bytes[28:32])
	//ecdsaQuote.QEReport.Reserved1 = bytes[32]
	//ecdsaQuote.QEReport.ISVProdID = binary.LittleEndian.Uint16(bytes[32:34])
	//ecdsaQuote.QEReport.Attributes = *(*[16]byte)(bytes[96:112])
	ecdsaQuote.QEReport.MREnclave = *(*[32]byte)(bytes[112:144])
	//ecdsaQuote.QEReport.Reserved2 = *(*[32]byte)(bytes[144:176])
	//ecdsaQuote.QEReport.MRSigner = *(*[32]byte)(bytes[176:208])
	//ecdsaQuote.QEReport.Reserved3 = *(*[96]byte)(bytes[208:304])
	//ecdsaQuote.QEReport.ISVSVN = binary.LittleEndian.Uint16(bytes[304:306])
	//ecdsaQuote.QEReport.Reserved4 = *(*[60]byte)(bytes[306:366])
	//ecdsaQuote.QEReport.ReportData = *(*[64]byte)(bytes[366:430])

	bytes = bytes[576:]
	quAuthDataSize := binary.LittleEndian.Uint16(bytes[0:2])

	bytes = bytes[2:]
	if len(bytes) < int(quAuthDataSize) {
		return false, errors.New("Invalid QE auth data size")
	}
	ecdsaQuote.QEAuthData = bytes[:quAuthDataSize]

	// mrenclave
	_, mrEnclaveFromContract, err := c.contractObj.GetTeeInfo(&bind.CallOpts{}, common.HexToAddress(c.cf.Contract.TeeAddr))
	if err != nil {
		return false, errors.Wrap(err, "contract GetTeeInfo error")
	}
	mrEnclave := hex.EncodeToString(ecdsaQuote.QEReport.MREnclave[:])
	if mrEnclave != mrEnclaveFromContract {
		return false, errors.New("mrenclave not match")
	}

	return true, nil

}

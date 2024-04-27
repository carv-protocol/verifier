package dcap

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/carv-protocol/verifier/internal/conf"
	"math/big"
)

type Quote struct {
	Base   Base   `json:"base"`
	Report Report `json:"report_body"`
}

// VerifyQuote validates both the header and the enclave report of a quote
func (q *Quote) VerifyQuote(quoteByte []byte, load *TrusTEEInfo, q3Atuth *QuoteV3Auth, cf *conf.Bootstrap) error {
	if err := q.Base.Verify(); err != nil {
		return err
	}
	// load is nil when the enclave remote
	if err := q.Report.Verify(load); err != nil {
		return err
	}
	if err := q3Atuth.Verify(cf); err != nil {
		return err
	}
	// Last: Verify local attestation signature
	if quoteByte != nil {
		err := VerifySignature(q3Atuth.AttestationKey[:], q3Atuth.Signature[:], quoteByte[:432], true)
		if err != nil {
			return err
		}
	}
	return nil
}

// EnclaveReport mirrors the Rust EnclaveReport struct
type Report struct {
	Cpusvn     [16]byte `json:"cpusvn"`
	Miscselect [4]byte  `json:"miscselect"`
	Reserved1  [28]byte `json:"reserved1"`
	Attributes [16]byte `json:"attributes"`
	MrEnclave  [32]byte `json:"mrenclave"`
	Reserved2  [32]byte `json:"reserved2"`
	MrSigner   [32]byte `json:"mrsigner"`
	Reserved3  [96]byte `json:"reserved3"`
	IsvProdid  [2]byte  `json:"isv_prodid"`
	IsvSvn     [2]byte  `json:"isv_svn"`
	Reserved4  [60]byte `json:"reserved4"`
	ReportData [64]byte `json:"reportdata"`
}

// FromBytes populates an EnclaveReport from a byte slice
func (e *Report) Verify(info *TrusTEEInfo) error {
	if info == nil {
		return errors.New("TrusTEEInfo is nil")
	}
	k, v := info.TrustedEnclaves[hex.EncodeToString(e.MrSigner[:])]
	if !v {
		return errors.New("report not from a trusted Singner")
	}

	for _, enclave := range k {
		if enclave == hex.EncodeToString(e.MrEnclave[:]) {
			return nil
		}
	}
	return fmt.Errorf("report not from a trusted enclave")
}

// Header mirrors the Rust Header struct
type Base struct {
	Version    [2]byte  `json:"version"`
	KeyType    [2]byte  `json:"key_type"`
	Reserved   [4]byte  `json:"reserved"`
	QeSvn      [2]byte  `json:"qe_svn"`
	PceSvn     [2]byte  `json:"pce_svn"`
	QeVendorId [16]byte `json:"qe_vendor_id"`
	UserData   [20]byte `json:"user_data"`
}

const supportedVersion uint16 = 3
const supportedAttestationKeyType uint16 = 2
const supportedTeeType uint32 = 0

var validQeVendorID = [16]byte{0x93, 0x9a, 0x72, 0x33, 0xf7, 0x9c, 0x4c, 0xa9, 0x94, 0x0a, 0x0d, 0xb3, 0x95, 0x7f, 0x06, 0x07}

func (b *Base) Verify() error {

	if binary.LittleEndian.Uint16(b.Version[:]) != supportedVersion {
		return errors.New("unsupported SGX quote version")
	}
	if binary.LittleEndian.Uint16(b.KeyType[:]) != supportedAttestationKeyType {
		return errors.New("unsupported attestation key type")
	}
	if binary.LittleEndian.Uint32(b.Reserved[:]) != supportedTeeType {
		return errors.New("unsupported TEE type")
	}
	if b.QeVendorId != validQeVendorID {
		return errors.New("invalid QE vendor ID")
	}
	return nil
}

type QuoteV3Auth struct {
	Signature         [64]byte  `json:"signature"`
	AttestationKey    [64]byte  `json:"attestation_key"`
	RawQEReport       [384]byte `json:"raw_qe_report"`
	QEReportSignature [64]byte  `json:"qe_report_signature"`
	QEAuthData        []byte    `json:"qe_auth_data"`
	QECert            CertData  `json:"qe_cert"`
}

// FromBytes deserializes bytes into ECDSAQuoteV3AuthData
func (e *QuoteV3Auth) FromBytes(data []byte) error {
	if len(data) < 576 {
		return errors.New("invalid length for ECDSA Quote V3 Auth Data")
	}

	copy(e.Signature[:], data[:64])
	copy(e.AttestationKey[:], data[64:128])
	copy(e.RawQEReport[:], data[128:512])
	copy(e.QEReportSignature[:], data[512:576])

	remainingData := data[576:]
	if len(remainingData) < 2 {
		return errors.New("missing QE auth data size")
	}

	qeAuthDataSize := binary.LittleEndian.Uint16(remainingData[:2])
	remainingData = remainingData[2:]

	if int(qeAuthDataSize) > len(remainingData) {
		return errors.New("QE auth data size exceeds remaining data length")
	}

	e.QEAuthData = make([]byte, qeAuthDataSize)
	copy(e.QEAuthData, remainingData[:qeAuthDataSize])
	remainingData = remainingData[qeAuthDataSize:]

	if err := e.QECert.FromBytes(remainingData); err != nil {
		return err
	}

	return nil
}

// ToBytes serializes ECDSAQuoteV3AuthData into bytes
func (e *QuoteV3Auth) ToBytes() ([]byte, error) {
	buf := bytes.NewBuffer(nil)

	buf.Write(e.Signature[:])
	buf.Write(e.AttestationKey[:])
	buf.Write(e.RawQEReport[:])
	buf.Write(e.QEReportSignature[:])

	err := binary.Write(buf, binary.LittleEndian, uint16(len(e.QEAuthData)))
	if err != nil {
		return nil, err
	}
	buf.Write(e.QEAuthData)

	qeCertBytes, err := e.QECert.ToBytes()
	if err != nil {
		return nil, err
	}
	buf.Write(qeCertBytes)

	return buf.Bytes(), nil
}
func (q *QuoteV3Auth) Verify(cf *conf.Bootstrap) error {
	enclaveID := new(EnclaveId).GetEnclaveID(cf.Dacp.IdentityPath)
	qeReport, err := q.GetQEReport()
	if err != nil {
		return err
	}

	miscselectMatched := (binary.LittleEndian.Uint32(qeReport.Miscselect[:]) & enclaveID.MiscselectMask) == uint32(enclaveID.Miscselect)
	attributesMatched := false
	for i, attr := range qeReport.Attributes {
		if (attr & enclaveID.AttributesMask[i]) == enclaveID.Attributes[i] {
			attributesMatched = true
			break
		}
	}

	mrsignerMatched := qeReport.MrSigner == enclaveID.Mrsigner

	tcbFound := false
	for _, tcb := range enclaveID.TcbLevels {
		if binary.BigEndian.Uint16(qeReport.IsvSvn[:]) >= tcb.Isvsvn && tcb.TcbStatus == "OK" {
			tcbFound = true
			break
		}
	}

	if !(miscselectMatched && attributesMatched && mrsignerMatched && tcbFound) {
		return errors.New("enclave ID does not match")
	}

	// Assuming certificate parsing is already handled elsewhere
	root, ca, pck, err := ParseCertificates(q.QECert.CertData)
	if err != nil {
		return err
	}
	//  PCK check
	tcbInfo := new(TcbInfo).GetTcbInfo(cf.Dacp.TcbPath)
	if !bytes.Equal(tcbInfo.Fmspc, pck.Fmspc[:]) || !bytes.Equal(tcbInfo.PceID, pck.PceID[:]) {
		return errors.New("unmatched TCB: fmspc or pceId does not match")
	}

	// TCB check
	for _, tcbLevel := range tcbInfo.TcbLevels {
		pceSvnCheck := pck.TCB.PceSVN >= tcbLevel.Tcb.Pcesvn
		cpuSvnCheck := true
		for i, compSvn := range pck.TCB.CompSVNArray {
			if compSvn < tcbLevel.Tcb.SgxtcbcompSvn[i] {
				cpuSvnCheck = false
				break
			}
		}
		if pceSvnCheck && cpuSvnCheck {
			if tcbLevel.TcbStatus == "Revoked" {
				return errors.New("TCB is revoked")
			}
			break
		}
	}

	// Assuming root, ca, pck are loaded x509.Certificate and the signatures are loaded
	rootPubKeyBytes, err := getPublicKeyFromCert(root.Raw)
	if err != nil {
		return err
	}
	caPubKeyBytes, err := getPublicKeyFromCert(ca.Raw)
	if err != nil {
		return err
	}
	pckPubKeyBytes, err := getPublicKeyFromCert(pck.Inner.Raw)
	if err != nil {
		return err
	}

	// Verify the certificate chain
	err = VerifySignature(rootPubKeyBytes, ca.Signature, ca.RawTBSCertificate, false)
	if err != nil {
		return err
	}
	err = VerifySignature(caPubKeyBytes, pck.Inner.Signature, pck.Inner.RawTBSCertificate, false)
	if err != nil {
		return err
	}
	// Verify the QE report signature
	err = VerifySignature(pckPubKeyBytes, q.QEReportSignature[:], q.RawQEReport[:], false)
	if err != nil {
		return err
	}

	// Verify the QE Report's Hash
	hash := sha256.New()
	hash.Write(q.AttestationKey[:])
	hash.Write(q.QEAuthData)
	actualHash := hash.Sum(nil)
	if !bytes.Equal(actualHash, qeReport.ReportData[:32]) {
		return errors.New("QE Report's Hash is Invalid")
	}

	return nil
}
func (q *QuoteV3Auth) GetQEReport() (Report, error) {
	var report Report
	var byteReader = bytes.NewReader(q.RawQEReport[:])
	err := binary.Read(byteReader, binary.LittleEndian, &report)
	if err != nil {
		return Report{}, err
	}
	return report, nil
}

func GetQuoteV3Auth(quote []byte) (QuoteV3Auth, error) {
	if len(quote) < 576 {
		return QuoteV3Auth{}, errors.New("quote too short")
	}
	signatureLen := binary.LittleEndian.Uint32(quote[432:436])
	fmt.Println(signatureLen)
	var quoteAuth QuoteV3Auth
	b := quote[436 : 436+signatureLen]
	err := quoteAuth.FromBytes(b)
	if err != nil {
		return QuoteV3Auth{}, err
	}
	return quoteAuth, nil
}

type VerifyingKey struct {
	*ecdsa.PublicKey
}

// fromPublicKeyBytes initializes a VerifyingKey from a public key in uncompressed SEC1 format.
func fromPublicKeyBytes(pubKeyBytes []byte) (*VerifyingKey, error) {
	pubKey, err := x509.ParsePKIXPublicKey(pubKeyBytes)
	if err != nil {
		return nil, errors.New("invalid public key: " + err.Error())
	}
	ecdsaPubKey, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("public key is not of type ECDSA")
	}
	return &VerifyingKey{ecdsaPubKey}, nil
}

// VerifySignature verifies a signature using the ECDSA public key, SHA256 hashing, and ASN.1 DER-encoded signature.
func VerifySignature(pubKeyBytes, signatureBytes, message []byte, untagged bool) error {
	var vk *VerifyingKey
	var err error
	if untagged {
		vk, err = fromUntaggedBytes(pubKeyBytes)
		if err != nil {
			return err
		}
	} else {
		vk, err = fromPublicKeyBytes(pubKeyBytes)
		if err != nil {
			return err
		}
	}

	hash := sha256.New()
	hash.Write(message)
	hashed := hash.Sum(nil)

	var sig struct {
		R, S *big.Int
	}
	_, err = asn1.Unmarshal(signatureBytes, &sig)
	if err != nil {
		r := big.NewInt(0).SetBytes(signatureBytes[:32])
		s := big.NewInt(0).SetBytes(signatureBytes[32:])
		sig.R = r
		sig.S = s
	}

	if !ecdsa.Verify(vk.PublicKey, hashed, sig.R, sig.S) {
		return errors.New("signature verification failed")
	}
	return nil
}

// Extract the public key info from a certificate in DER format
func getPublicKeyFromCert(certBytes []byte) ([]byte, error) {
	cert, err := x509.ParseCertificate(certBytes)
	if err != nil {
		return nil, err
	}
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(cert.PublicKey)
	if err != nil {
		return nil, err
	}
	return pubKeyBytes, nil
}

// Convert untagged bytes to an ECDSA public key
func fromUntaggedBytes(bytes []byte) (*VerifyingKey, error) {
	if len(bytes) != 64 {
		return nil, fmt.Errorf("expect 64 bytes but found %d", len(bytes))
	}

	x := big.NewInt(0).SetBytes(bytes[:32])
	y := big.NewInt(0).SetBytes(bytes[32:])

	pubKey := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	return &VerifyingKey{pubKey}, nil
}

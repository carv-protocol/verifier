package dcap

import (
	"bytes"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/binary"
	"encoding/pem"
	"errors"
	"fmt"
)

var (
	OIDTCBCompSVN = []asn1.ObjectIdentifier{
		{1, 2, 840, 113741, 1, 13, 1, 2, 1},
		{1, 2, 840, 113741, 1, 13, 1, 2, 2},
		{1, 2, 840, 113741, 1, 13, 1, 2, 3},
		{1, 2, 840, 113741, 1, 13, 1, 2, 4},
		{1, 2, 840, 113741, 1, 13, 1, 2, 5},
		{1, 2, 840, 113741, 1, 13, 1, 2, 6},
		{1, 2, 840, 113741, 1, 13, 1, 2, 7},
		{1, 2, 840, 113741, 1, 13, 1, 2, 8},
		{1, 2, 840, 113741, 1, 13, 1, 2, 9},
		{1, 2, 840, 113741, 1, 13, 1, 2, 10},
		{1, 2, 840, 113741, 1, 13, 1, 2, 11},
		{1, 2, 840, 113741, 1, 13, 1, 2, 12},
		{1, 2, 840, 113741, 1, 13, 1, 2, 13},
		{1, 2, 840, 113741, 1, 13, 1, 2, 14},
		{1, 2, 840, 113741, 1, 13, 1, 2, 15},
		{1, 2, 840, 113741, 1, 13, 1, 2, 16},
	}
	OIDTCBPCESVN    = asn1.ObjectIdentifier{1, 2, 840, 113741, 1, 13, 1, 2, 17}
	OIDPCEID        = asn1.ObjectIdentifier{1, 2, 840, 113741, 1, 13, 1, 3}
	OIDFMSPC        = asn1.ObjectIdentifier{1, 2, 840, 113741, 1, 13, 1, 4}
	OIDSGXExtension = asn1.ObjectIdentifier{1, 2, 840, 113741, 1, 13, 1}
	OIDTCB          = asn1.ObjectIdentifier{1, 2, 840, 113741, 1, 13, 1, 2}
)

// ASN.1
type Ext struct {
	Key   asn1.ObjectIdentifier
	Value asn1.RawValue
}

type TCB struct {
	PceSVN       uint16
	CompSVNArray [16]byte
}

type PCK struct {
	Inner *x509.Certificate
	Fmspc [6]byte
	PceID [2]byte
	TCB   TCB
}

func NewPCK(cert *x509.Certificate) (*PCK, error) {
	var sgxExt pkix.Extension
	for _, ext := range cert.Extensions {
		if ext.Id.Equal(OIDSGXExtension) {
			sgxExt = ext
			break
		}
	}
	if sgxExt.Id == nil {
		return nil, errors.New("SGX extension not found in certificate")
	}

	var sgxExts []Ext
	_, err := asn1.Unmarshal(sgxExt.Value, &sgxExts)
	if err != nil {
		return nil, errors.New("failed to unmarshal SGX extension")
	}

	sgxExtMap := make(map[string]asn1.RawValue)
	for _, ext := range sgxExts {
		sgxExtMap[fmt.Sprintf("%v", ext.Key)] = ext.Value
	}
	fmspc, err := parseOctetString(sgxExtMap[fmt.Sprintf("%v", OIDFMSPC)])
	if err != nil {
		return nil, err
	}
	pceID, err := parseOctetStringID(sgxExtMap[fmt.Sprintf("%v", OIDPCEID)])
	if err != nil {
		return nil, err
	}
	tcb, err := parseTCB(sgxExtMap, OIDTCBCompSVN, OIDTCB)
	if err != nil {
		return nil, err
	}

	return &PCK{
		Inner: cert,
		Fmspc: fmspc,
		PceID: pceID,
		TCB:   tcb,
	}, nil
}

func parseOctetString(raw asn1.RawValue) ([6]byte, error) {
	var result [6]byte
	if len(raw.Bytes) != 6 {
		return result, errors.New("incorrect length for octet string")
	}
	copy(result[:], raw.Bytes)
	return result, nil
}
func parseOctetStringID(raw asn1.RawValue) ([2]byte, error) {
	var result [2]byte
	if len(raw.Bytes) != 2 {
		return result, errors.New("incorrect length for octet string")
	}
	copy(result[:], raw.Bytes)
	return result, nil

}

// parseTCB
func parseTCB(sgxExtMap map[string]asn1.RawValue, compSVNOIDs []asn1.ObjectIdentifier, pceSVNOID asn1.ObjectIdentifier) (TCB, error) {
	var tcb TCB
	pceSVNRaw, exists := sgxExtMap[fmt.Sprintf("%v", pceSVNOID)]
	if !exists {
		return tcb, errors.New(fmt.Sprintf("PCE SVN OID not found: %v", pceSVNOID))
	}
	var sequence []Ext
	_, err := asn1.Unmarshal(pceSVNRaw.FullBytes, &sequence)
	if err != nil {
		return tcb, errors.New(fmt.Sprintf("failed to unmarshal ASN.1 sequence: %w", err))
	}
	pceExtMap := make(map[string]asn1.RawValue)
	for i, ext := range sequence {
		pceExtMap[fmt.Sprintf("%v", ext.Key)] = sequence[i].Value
	}
	pceSVN, err := parseUint16ASN1(sequence)
	if err != nil {
		return tcb, errors.New(fmt.Sprintf("error parsing PCE SVN: %w", err))
	}
	tcb.PceSVN = pceSVN
	var compSVNArray []byte
	for _, oid := range compSVNOIDs {
		compSVNRaw, exists := pceExtMap[fmt.Sprintf("%v", oid)]
		if !exists {
			continue // Some OIDs might not be present, depending on the certificate
		}
		compSVN := compSVNRaw.Bytes

		compSVNArray = append(compSVNArray, compSVN[:]...)
	}
	j := 0
	for i := 0; i < len(compSVNArray); i++ {
		if j > len(tcb.CompSVNArray) {
			return tcb, errors.New(fmt.Sprintf("error parsing Comp SVN: %w", err))
		}
		if compSVNArray[i] > 0 {
			tcb.CompSVNArray[j] = compSVNArray[i]
			j++
		}
	}

	return tcb, nil
}

// parseUint16ASN1 parses a uint16 from ASN.1 encoded data
func parseUint16ASN1(sequence []Ext) (uint16, error) {

	for _, ext := range sequence {
		if ext.Key.Equal(OIDTCBPCESVN) {
			// Convert num to []byte
			if len(ext.Value.Bytes) == 1 {
				return uint16(ext.Value.Bytes[0]), nil
			}
			num := binary.BigEndian.Uint16(ext.Value.Bytes)
			return num, nil
		}
	}
	return 0, errors.New("PCE SVN not found in ASN.1 sequence")

}

type CertData struct {
	CertType uint16 `json:"cert_type"`
	CertData []byte `json:"cert_data"`
}

// FromBytes deserializes bytes into QECertData
func (q *CertData) FromBytes(data []byte) error {
	if len(data) < 6 {
		return errors.New("invalid length for QE Cert Data")
	}
	q.CertType = binary.LittleEndian.Uint16(data[:2])
	certSize := binary.LittleEndian.Uint32(data[2:6])

	if len(data[6:]) != int(certSize) {
		return errors.New("mismatched cert size")
	}

	q.CertData = data[6 : 6+certSize]
	return nil
}

// ToBytes serializes QECertData into bytes
func (q *CertData) ToBytes() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, q.CertType)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, uint32(len(q.CertData)))
	if err != nil {
		return nil, err
	}
	buf.Write(q.CertData)
	return buf.Bytes(), nil
}

// ParseCertificates parses a PEM-encoded certificate chain and returns the root CA certificate, the intermediate CA certificate, and the PCK certificate.
func ParseCertificates(certData []byte) (rootCert, caCert *x509.Certificate, pck *PCK, err error) {
	var certs []*x509.Certificate

	for {
		var block *pem.Block
		block, certData = pem.Decode(certData)
		if block == nil {
			break
		}

		if block.Type == "CERTIFICATE" {
			cert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				return nil, nil, nil, err
			}
			certs = append(certs, cert)
		}
	}

	if len(certs) < 3 {
		return nil, nil, nil, errors.New("expected at least three certificates")
	}
	pck, err = NewPCK(certs[0])
	if err != nil {
		return nil, nil, nil, err
	}
	return certs[2], certs[1], pck, nil
}

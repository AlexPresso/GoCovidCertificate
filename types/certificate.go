package types

const (
	DCC     CertificateType = "DCC"
	TWODDOC CertificateType = "2DDOC"
)

type CertificateType string

type Certificate struct {
	Type CertificateType
	Data interface{}
}

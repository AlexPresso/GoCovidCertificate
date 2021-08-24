package types

const (
	DCC     CertificateName = "DCC"
	TWODDOC CertificateName = "2DDOC"
)

type CertificateName string

type Certificate struct {
	Type CertificateName
	Data interface{}
}

package types

type TwoDDoc struct {
	Header    TwoDDocHeader
	Message   TwoDDocMessage
	Signature TwoDDocSignature
}

type TwoDDocHeader struct {
	IDFlag         string
	Version        int
	Issuer         string
	CertID         string
	DocumentDate   string
	SignatureDate  string
	DocumentTypeID string
	PerimeterID    string
	CountryID      string
}

type TwoDDocMessage struct {
	Data map[string]string
}

type TwoDDocSignature struct {
	Raw string
}

package types

type TwoDDoc struct {
	Header    *TwoDDocHeader
	Message   *TwoDDocMessage
	Signature string
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
	Lastname          string `json:"L0"`
	Firstname         string `json:"L1"`
	DateOfBirth       string `json:"L2"`
	TargetedAgent     string `json:"L3"`
	VaccineATC        string `json:"L4"`
	Dose1Manufacturer string `json:"L5"`
	Dose2Manufacturer string `json:"L6"`
	Dose              int    `json:"L7,string"`
	RequiredDoses     int    `json:"L8,string"`
	Date              int    `json:"L9,string"`
}

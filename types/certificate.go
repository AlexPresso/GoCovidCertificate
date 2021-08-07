package types

// DccRoot see https://github.com/ehn-dcc-development/ehn-dcc-schema
type DccRoot struct {
	DCC covidCertificate `cbor:"1,keyasint"`
}

type covidCertificate struct {
	Version			string			`cbor:"ver" json:"version"`
	DateOfBirth		string			`cbor:"dob" json:"dateOfBirth"`
	Name			name			`cbor:"nam" json:"name"`
	Vaccines		[]vaccineEntry	`cbor:"v" json:"vaccines"`
	Tests			[]testEntry 	`cbor:"t" json:"tests"`
	Recoveries		[]recoveryEntry	`cbor:"r" json:"recoveries"`
}

type name struct {
	Surname			string	`cbor:"fn" json:"surname"`
	StdSurname		string	`cbor:"fnt" json:"stdSurname"`
	Forename		string	`cbor:"gn" json:"forename"`
	StdForename		string	`cbor:"gnt" json:"stdForename"`
}

type vaccineEntry struct {
	TargetedAgent	string	`cbor:"tg" json:"targetedAgent"`
	Vaccine			string  `cbor:"vp" json:"vaccine"`
	Product			string  `cbor:"mp" json:"product"`
	Manufacturer	string  `cbor:"ma" json:"manufacturer"`
	Dose			int64	`cbor:"dn" json:"doseNumber"`
	DoseSeries		int64	`cbor:"sd" json:"doseSeries"`
	Date			string	`cbor:"dt" json:"date"`
	Country			string	`cbor:"co" json:"country"`
	Issuer			string	`cbor:"is" json:"issuer"`
	CertificateID	string	`cbor:"ci" json:"certificateID"`
}

type testEntry struct {
	TargetedAgent	string	`cbor:"tg" json:"targetedAgent"`
	TestType 		string	`cbor:"tt" json:"testType"`
	Name 			string	`cbor:"nm" json:"name"`
	Manufacturer   	string	`cbor:"ma" json:"manufacturer"`
	SampleDatetime 	string	`cbor:"sc" json:"sampleDatetime"`
	TestResult     	string	`cbor:"tr" json:"testResult"`
	TestingCentre  	string	`cbor:"tc" json:"testingCentre"`
	Country       	string	`cbor:"co" json:"country"`
	Issuer        	string	`cbor:"is" json:"issuer"`
	CertificateID 	string	`cbor:"ci" json:"certificateID"`
}

type recoveryEntry struct {
	TargetedAgent	string	`cbor:"tg" json:"targetedAgent"`
	Country 		string	`cbor:"co" json:"country"`
	Issuer 			string	`cbor:"is" json:"issuer"`
	ValidFrom		string	`cbor:"df" json:"validFrom"`
	ValidUntil		string	`cbor:"du" json:"validUntil"`
	CertificateID 	string	`cbor:"ci" json:"certificateID"`
}

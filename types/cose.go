package types

// Signed see https://datatracker.ietf.org/doc/html/rfc8152#section-2
type Signed struct {
	_			struct{} `cbor:",toarray"`
	Protected	[]byte
	Unprotected	map[interface{}]interface{}
	Content		[]byte
	Signature	[]byte
}

// Header see https://datatracker.ietf.org/doc/html/rfc8152#section-3.1 & https://www.iana.org/assignments/cose/cose.xhtml
type Header struct {
	Alg		int 	`cbor:"1,keyasint,omitempty"`
	Kid		[]byte 	`cbor:"4,keyasint,omitempty"`
	IV		[]byte	`cbor:"5,keyasint,omitempty"`
}

// Claims see https://datatracker.ietf.org/doc/html/draft-ietf-ace-cbor-web-token-08#section-3.1
type Claims struct {
	Iss		string	`cbor:"1,keyasint"`
	Sub		string	`cbor:"2,keyasint"`
	Aud		string	`cbor:"3,keyasint"`
	Exp		int64	`cbor:"4,keyasint"`
	Nbf		int	`cbor:"5,keyasint"`
	Iat		int64	`cbor:"6,keyasint"`
	Cti		[]byte  `cbor:"7,keyasint"`
	HCData		DccRoot `cbor:"-260,keyasint"`
	LCData		DccRoot `cbor:"-250,keyasint"`
}

type COSE struct {
	Signed	Signed `json:"-"`
	Header	Header
	Claims	Claims
}

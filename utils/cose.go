package utils

import (
	"github.com/alexpresso/gocovidcertificate/types"
	"github.com/fxamacker/cbor/v2"
	"log"
)


func DecodeCOSE(bytes []byte) *types.COSE {
	var signed types.Signed
	if err := cbor.Unmarshal(bytes, &signed); err != nil {
		log.Fatal(err)
	}

	var header types.Header
	if len(signed.Protected) > 0 {
		if err := cbor.Unmarshal(signed.Protected, &header); err != nil {
			log.Fatal(err)
		}
	}

	var claims types.Claims
	if err := cbor.Unmarshal(signed.Content, &claims); err != nil {
		log.Fatal(err)
	}

	return &types.COSE{
		Signed: signed,
		Header: header,
		Claims: claims,
	}
}
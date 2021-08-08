package utils

import (
	"errors"
	"github.com/alexpresso/gocovidcertificate/types"
	"github.com/fxamacker/cbor/v2"
)

func DecodeCOSE(bytes []byte) (*types.COSE, error) {
	var signed types.Signed
	if err := cbor.Unmarshal(bytes, &signed); err != nil {
		return nil, errors.New("unable to decode cose")
	}

	var header types.Header
	if len(signed.Protected) > 0 {
		if err := cbor.Unmarshal(signed.Protected, &header); err != nil {
			return nil, errors.New("unable to compute signed data")
		}
	}

	var claims types.Claims
	if err := cbor.Unmarshal(signed.Content, &claims); err != nil {
		return nil, errors.New("")
	}

	return &types.COSE{
		Signed: signed,
		Header: header,
		Claims: claims,
	}, nil
}

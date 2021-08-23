package decoders

import (
	"github.com/alexpresso/gocovidcertificate/types"
	"github.com/alexpresso/gocovidcertificate/utils"
	"strconv"
)

var TwoDPrefix = "DC"

func twoDDocDecode(input string) (data *types.TwoDDoc, err error) {
	header, remaining, err := decodeHeader(input)
	if err != nil {
		return nil, err
	}

	message, remaining, err := decodeMessage(remaining)
	if err != nil {
		return nil, err
	}

	signature, err := decodeSignature(remaining)
	if err != nil {
		return nil, err
	}

	return &types.TwoDDoc{
		Header:    *header,
		Message:   *message,
		Signature: *signature,
	}, nil
}

func decodeHeader(input string) (header *types.TwoDDocHeader, remaining string, err error) {
	length := 22
	version, err := strconv.Atoi(utils.Substring(input, 2, 4))
	if err != nil {
		return nil, "", err
	}

	header = &types.TwoDDocHeader{
		IDFlag:         utils.Substring(input, 0, 2),
		Version:        version,
		Issuer:         utils.Substring(input, 4, 8),
		CertID:         utils.Substring(input, 8, 12),
		DocumentDate:   utils.Substring(input, 12, 16),
		SignatureDate:  utils.Substring(input, 16, 20),
		DocumentTypeID: utils.Substring(input, 20, 22),
	}

	switch version {
	case 3:
		header.PerimeterID = utils.Substring(input, 22, 24)
		length = 24
	case 4:
		header.PerimeterID = utils.Substring(input, 22, 24)
		header.CountryID = utils.Substring(input, 24, 26)
		length = 26
	}

	return header, utils.Substring(input, length, len(input)), err
}

func decodeMessage(input string) (message *types.TwoDDocMessage, remaining string, err error) {
	remaining = input
	data := make(map[string]string)

	message = &types.TwoDDocMessage{
		Data: data,
	}

	return message, remaining, nil
}

func decodeField(input string) (key string, value string, remaining string, err error) {
	key = utils.Substring(input, 0, 2)
	value = utils.Substring(input, 2, 4)
	remaining = utils.Substring(input, 4, len(input))

	return
}

func decodeSignature(input string) (signature *types.TwoDDocSignature, err error) {
	return nil, nil
}

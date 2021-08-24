package decoders

import (
	"encoding/json"
	"github.com/alexpresso/gocovidcertificate/types"
	"github.com/alexpresso/gocovidcertificate/utils"
	"strconv"
	"strings"
)

var TwoDPrefix = "DC"

func twoDDocDecode(input string) (certificate *types.Certificate, err error) {
	header, remaining, err := decodeHeader(input)
	if err != nil {
		return nil, err
	}

	message, signature, err := decodeData(remaining)
	if err != nil {
		return nil, err
	}

	return &types.Certificate{
		Type: types.TWODDOC,
		Data: &types.TwoDDoc{
			Header:    header,
			Message:   message,
			Signature: signature,
		},
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

func decodeData(input string) (message *types.TwoDDocMessage, signature string, err error) {
	data := make(map[string]interface{})

	units := strings.Split(input, string(rune(31)))
	groups := strings.Split(units[0], string(rune(29)))
	signature = units[1]

	for _, group := range groups {
		key, value := decodeField(group)
		data[key] = value
	}

	jsonString, err := json.Marshal(data)
	err = json.Unmarshal(jsonString, &message)

	return
}

func decodeField(input string) (key string, value string) {
	key = utils.Substring(input, 0, 2)
	value = utils.Substring(input, 2, len(input))
	return
}

package decoders

import (
	"github.com/alexpresso/gocovidcertificate/types"
	"github.com/alexpresso/gocovidcertificate/utils"
	"strings"
)

var DCCPrefixes = map[string]bool{"HC1": true, "LT1": true}

func dccDecode(input string) (*types.COSE, error) {
	var err error
	var bytes []byte

	for prefix := range DCCPrefixes {
		input = strings.TrimPrefix(input, prefix+":")
	}

	bytes, err = utils.Base45decode(input)
	if err != nil {
		return nil, err
	}

	if bytes[0] == 0x78 {
		bytes, err = utils.ZlibDecompress(bytes)
		if err != nil {
			return nil, err
		}
	}

	return utils.DecodeCOSE(bytes)
}

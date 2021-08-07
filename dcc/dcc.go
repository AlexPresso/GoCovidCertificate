package dcc

import (
	"errors"
	"github.com/alexpresso/gocovidcertificate/types"
	"github.com/alexpresso/gocovidcertificate/utils"
	"strings"
)


var prefixes = map[string]bool{"HC1":true, "LT1":true}

func Decode(input string) (*types.COSE, error) {
	if split := strings.Split(input, ":")[0]; !prefixes[split] {
		return nil, errors.New("input not starting with a DCC prefix")
	}

	for prefix := range prefixes {
		input = strings.TrimPrefix(input, prefix + ":")
	}

	bytes := utils.Base45decode(input)

	if bytes[0] == 0x78 {
		bytes = utils.ZlibDecompress(bytes)
	}

	return utils.DecodeCOSE(bytes), nil
}

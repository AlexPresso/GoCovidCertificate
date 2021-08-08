package decoders

import (
	"errors"
	"strings"
)

func Decode(input string) (interface{}, error) {
	if split := strings.Split(input, ":")[0]; DCCPrefixes[split] {
		return dccDecode(input)
	} else if strings.HasPrefix(input, TwoDPrefix) {
		return twoDDocDecode(input)
	}

	return nil, errors.New("unsupported code format")
}

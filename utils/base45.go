package utils

import "bytes"


var charset = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./:")

func Base45decode(input string) []byte {
	length := len(input)
	buffer := make([]int, length)

	var res []byte

	for i := 0; i < length; i++ {
		buffer[i] = bytes.IndexByte(charset, input[i])
	}

	for i := 0; i < length; i += 3 {
		if length - i >= 3 {
			var x = buffer[i] + buffer[i + 1] * 45 + buffer[i + 2] * 45 * 45
			quotient, remainder := divmod(x, 256)

			res = append(res, byte(quotient), byte(remainder))
		} else {
			res = append(res, byte(buffer[i] + buffer[i + 1] * 45))
		}
	}

	return res
}

func divmod(numerator, denominator int) (quotient, remainder int) {
	return numerator / denominator, numerator % denominator
}
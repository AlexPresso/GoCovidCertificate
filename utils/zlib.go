package utils

import (
	"bytes"
	"compress/zlib"
	"errors"
	"io/ioutil"
)

func ZlibDecompress(b []byte) ([]byte, error) {
	var err error

	z, err := zlib.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, errors.New("error while decompressing data, zlib error")
	}

	defer func() {
		err = z.Close()
	}()
	if err != nil {
		return nil, errors.New("error while decompressing data, zlib reader refused to close")
	}

	out, _ := ioutil.ReadAll(z)

	return out, nil
}

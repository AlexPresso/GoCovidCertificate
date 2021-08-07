package utils

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"
	"log"
)


func ZlibDecompress(b []byte) []byte {
	var err error

	z, err := zlib.NewReader(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = z.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	out, _ := ioutil.ReadAll(z)

	return out
}

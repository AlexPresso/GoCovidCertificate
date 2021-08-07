package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/alexpresso/gocovidcertificate/dcc"
	"log"
)


func main() {
	var (
		err error
		mustPrint = flag.Bool("print", true, "Prints content of the decoded certificate")
		code = flag.String("code", "", "The DCC code you want to process")
	)

	flag.Parse()

	if len([]rune(*code)) == 0 {
		log.Fatal("Missing -code flag")
	}

	cose, err := dcc.Decode(*code)
	if err != nil {
		log.Fatal(err)
	}

	if *mustPrint {
		indent, err := json.MarshalIndent(cose, "", "   ")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", string(indent))
	}
}

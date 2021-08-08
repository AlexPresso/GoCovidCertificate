# GoCovidCertificate

Golang implementation of the covid certificates. At the moment it only includes DCC signed data decoding but I've planned to add a lot more of features related to certificates processing.

## Features
- Decode signed DCC (European QRCodes) data ✅
- Decode 2D-DOC data ❌ (planned)
- Pretty-print decoded data as JSON ✅
- Download public-keys from european gateway ❌ (planned)
- Verify signature ❌ (planned)

## Installation
`go install github.com/alexpresso/gocovidcertificate@latest`

## Usage
`gocovidcertificate <flags>`

| Flag   | Type   | Description                                               | Required | Default value |
| ------ | ------ | --------------------------------------------------------- | -------- | ------------- |
| -code  | string | QRCode data to decode (put it between double-quotes `""`) | true     | none          |
| -print | bool   | Prints the QRCode data to console                         | false    | true          |

Example:  
`gocovidcertificate -code "HC1:..." -print`

## Use as dependency/library
(API docs comming soon)  
`go get github.com/alexpresso/gocovidcertificate`

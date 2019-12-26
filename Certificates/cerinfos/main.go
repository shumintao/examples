package main

import (
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/grantae/certinfo"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <cer file>\n", os.Args[0])
		return
	}

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}

	cert, err := x509.ParseCertificate(data)
	if err != nil {
		fmt.Println("Failed to parse data:", err)
		return
	}

	text, err := certinfo.CertificateText(cert)
	if err != nil {
		fmt.Println("Failed to parse certificate:", err)
		return
	}

	fmt.Println("After parse: ", text)
}

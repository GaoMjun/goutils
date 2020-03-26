package goutils

import (
	"crypto/rsa"
	"crypto/x509"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestCA(t *testing.T) {
	var (
		err error
		ca  *x509.Certificate
		pk  *rsa.PrivateKey
	)
	defer func() {
		if err != nil {
			log.Println(err)
		}
	}()

	if ca, pk, err = NewAuthority("htun"); err != nil {
		return
	}

	log.Println(ca, pk)

	ioutil.WriteFile("htun.cer", ca.Raw, os.ModePerm)
	ioutil.WriteFile("htun.key", x509.MarshalPKCS1PrivateKey(pk), os.ModePerm)
}

func TestCertLoad(t *testing.T) {
	var (
		err          error
		caRaw, pkRaw []byte
		ca           *x509.Certificate
		pk           *rsa.PrivateKey
	)
	defer func() {
		if err != nil {
			log.Println(err)
		}
	}()

	if caRaw, err = ioutil.ReadFile("htun.cer"); err != nil {
		return
	}

	if pkRaw, err = ioutil.ReadFile("htun.key"); err != nil {
		return
	}

	if ca, err = x509.ParseCertificate(caRaw); err != nil {
		return
	}

	if pk, err = x509.ParsePKCS1PrivateKey(pkRaw); err != nil {
		return
	}

	log.Println(ca, pk)
}

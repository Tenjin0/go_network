package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"time"
)

func GenX509Cert() {

	random := rand.Reader

	var key rsa.PrivateKey
	loadKey("private.key", &key)

	now := time.Now()
	then := now.Add(60 * 60 * 24 * 365 * 1000 * 1000 * 1000)

	template := x509.Certificate{
		SerialNumber: big.NewInt(2019),
		Subject: pkix.Name{
			Organization:  []string{"Company, INC."},
			Country:       []string{"FR"},
			Province:      []string{""},
			Locality:      []string{"PARIS"},
			StreetAddress: []string{"10 rue du sapin vert"},
			PostalCode:    []string{"75014"},
		},
		NotBefore:    now,
		NotAfter:     then,
		SubjectKeyId: []byte{1, 2, 3, 4},
		KeyUsage: x509.KeyUsageCertSign |
			x509.KeyUsageKeyEncipherment |
			x509.KeyUsageDigitalSignature,
		IsCA:                  true,
		BasicConstraintsValid: true,
		DNSNames: []string{
			"patrice.petit.name",
			"localhost",
		},
	}

	derBytes, err := x509.CreateCertificate(random, &template, &template, &key.PublicKey, &key)
	checkError(err)

	certCerFile, err := os.Create("certs/patrice.petit.name.cer")
	checkError(err)
	certCerFile.Write(derBytes)
	certCerFile.Close()

	certPEMFile, err := os.Create("certs/patrice.petit.name.pem")
	checkError(err)
	pem.Encode(certPEMFile, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: derBytes,
	})
	certPEMFile.Close()

	keyPEMFile, err := os.Create("certs/private.pem")
	checkError(err)
	pem.Encode(keyPEMFile, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(&key),
	})
	keyPEMFile.Close()
}

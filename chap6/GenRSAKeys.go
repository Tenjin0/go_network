package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"os"
)

func saveGobKey(fileName string, key interface{}) {

	outFile, err := os.Create("certs/" + fileName)

	checkError(err)
	defer outFile.Close()

	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)

}

func savePEMKey(fileName string, key *rsa.PrivateKey) {

	outFile, err := os.Create("certs/" + fileName)
	checkError(err)

	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	pem.Encode(outFile, privateKey)
}

func loadKey(fileName string, key interface{}) {

	inFile, err := os.Open("certs/" + fileName)
	checkError(err)
	defer inFile.Close()

	decoder := gob.NewDecoder(inFile)
	err = decoder.Decode(key)
	checkError(err)
}

func LoadRSAKeys() {

	var key rsa.PrivateKey
	loadKey("private.key", &key)

	fmt.Println("Private key primes",
		key.Primes[0].String(),
		key.Primes[1].String())
	fmt.Println("Private key exponent",
		key.D.String())

	var publicKey rsa.PublicKey
	loadKey("public.key", &publicKey)

	fmt.Println("Public key modulus",
		publicKey.N.String())
	fmt.Println("Public key exponent", publicKey.E)
}

func GenRSAKeys() {

	reader := rand.Reader
	bitSize := 512
	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)

	fmt.Println("Private key primes",
		key.Primes[0].String(),
		key.Primes[1].String())
	fmt.Println("Private key exponent",
		key.D.String())

	publicKey := key.PublicKey
	fmt.Println("Public key modulus",
		publicKey.N.String())
	fmt.Println("Public key exponent",
		publicKey.E)

	saveGobKey("private.key", key)
	saveGobKey("public.key", key.PublicKey)

	savePEMKey("private.pem", key)
}

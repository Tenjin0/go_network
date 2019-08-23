package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"os"
)

func Encrypt(keystring string, plainstring string) (string, error) {

	key := []byte(keystring)

	plainbytes := []byte(plainstring)
	ciphertext := make([]byte, aes.BlockSize+len(plainstring))

	cipherBlock := newCipher(key)
	err := initIV(ciphertext[:aes.BlockSize])

	if err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(cipherBlock, ciphertext[:aes.BlockSize])

	stream.XORKeyStream(ciphertext[aes.BlockSize:], plainbytes)
	return string(ciphertext), nil
}

func Decrypt(keystring string, cipherstring string) (string, error) {

	ciphertext := []byte(cipherstring)

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	key := []byte(keystring)
	cipherBlock, err := aes.NewCipher(key)

	if err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(cipherBlock, iv)

	plainbytes := make([]byte, len(ciphertext))

	stream.XORKeyStream(plainbytes, ciphertext)

	return string(plainbytes), nil
}

func initIV(iv []byte) error {

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err

	}
	return nil
}

func newCipher(key []byte) cipher.Block {

	block, err := aes.NewCipher(key)

	if err != nil {
		panic(err)
	}

	return block
}

func AesOneBlock() {

	key := []byte("my key, len 16 b")
	cipher, err := aes.NewCipher(key)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	src := []byte("hello 16 b block")

	var enc [16]byte

	cipher.Encrypt(enc[0:], src)

	fmt.Println(string(enc[0:]))

	var decrypt [16]byte
	cipher.Decrypt(decrypt[0:], enc[0:])
	result := bytes.NewBuffer(nil)

	result.Write(decrypt[0:])
	fmt.Println(string(result.Bytes()), string(decrypt[0:]))
}

package main

import (
	"crypto/hmac"
	"crypto/md5"
	"hash"
)

// MD5Hash return a MD5 hash of the string
func MD5Hash(str string) hash.Hash {

	hash := hmac.New(md5.New, []byte("secret"))
	bytes := []byte(str)

	hash.Write(bytes)

	return hash
}

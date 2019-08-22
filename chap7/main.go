package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error:", err.Error())
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) < 2 {

		fmt.Fprintf(os.Stderr, "Usage: %s \"function\" ...args\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	} else {

		name := os.Args[1]
		args := os.Args[2:]

		switch name {
		case "MD5Hash":
			MD5Hash(args[0])
		case "RSAKeys":
			GenRSAKeys()
		case "X509Cert":
			GenX509Cert()
		case "TLSGetHead":
			TLSGetHead(args[0])
		case "TLSEchoServer":
			TLSEchoServer()
		case "TLSEchoClient":
			TLSEchoClient(args[0])
		case "LoadRSAKeys":
			LoadRSAKeys()
		case "LoadX509Cert":
			LoadX509Cert()
		case "Aes":
			key := "example key 1234"

			cryptoText, err := Encrypt(key, args[0])
			fmt.Println("Encrypt", cryptoText)

			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			// encrypt base64 crypto to original value
			text, err := Decrypt(key, cryptoText)

			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			fmt.Println("Decrypt", string(text))
		default:
			fmt.Fprintln(os.Stderr, "No", name, "available")
		}
	}
}

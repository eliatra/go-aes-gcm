package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"

	"github.com/eliatra/go-aes-gcm/gcm"
)

func main() {

	key, set := os.LookupEnv("AES_SECRET_KEY")
	pass, pset := os.LookupEnv("AES_SECRET_PASSWORD")

	if !set && !pset {
		fmt.Println("AES_SECRET_KEY and AES_SECRET_PASSWORD not set")
		os.Exit(-1)
		return
	}

	if set && pset {
		fmt.Println("AES_SECRET_KEY and AES_SECRET_PASSWORD set. Only one should be defined.")
		os.Exit(-1)
		return
	}

	var keyBytes []byte

	if key != "" {
		keyBytes = []byte(key)
	} else {
		keyBytes0, err := gcm.CreateKeyFromPassword(pass)
		if err != nil {
			fmt.Println("ERROR: " + err.Error())
			os.Exit(-1)
			return
		}

		keyBytes = keyBytes0
	}

	if len(os.Args) <= 1 {
		fmt.Println("No arguments")
		os.Exit(-1)
		return
	}

	if len(os.Args) == 4 {
		if os.Args[1] == "encrypt" {
			err := gcm.EncryptFile(os.Args[2], os.Args[3], keyBytes)
			if err != nil {
				fmt.Println("ERROR: " + err.Error())
				os.Exit(-1)
				return
			}
			os.Exit(0)
			return

		} else {
			err := gcm.DecryptFile(os.Args[2], os.Args[3], keyBytes)
			if err != nil {
				fmt.Println("ERROR: " + err.Error())
				os.Exit(-1)
				return
			}
			os.Exit(0)
			return
		}
	}

	if len(os.Args) == 2 {
		if os.Args[1] == "encrypt" {
			stdin, err := io.ReadAll(os.Stdin)
			if err != nil {
				fmt.Println("ERROR: " + err.Error())
				os.Exit(-1)
				return
			}

			cipherText, err := gcm.Encrypt(stdin, keyBytes, nil)
			if err != nil {
				fmt.Println("ERROR: " + err.Error())
				os.Exit(-1)
				return
			}

			fmt.Print(base64.StdEncoding.EncodeToString(cipherText))

			os.Exit(0)
			return

		} else {
			stdin, err := io.ReadAll(os.Stdin)
			if err != nil {
				fmt.Println("ERROR: " + err.Error())
				os.Exit(-1)
				return
			}

			base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(stdin)))
			l, err := base64.StdEncoding.Decode(base64Text, stdin)

			if err != nil {
				fmt.Println("ERROR: " + err.Error())
				os.Exit(-1)
				return
			}

			plainText, err := gcm.Decrypt(base64Text[:l], keyBytes, nil)
			if err != nil {
				fmt.Println("ERROR: " + err.Error())
				os.Exit(-1)
				return
			}

			fmt.Print(string(plainText))

			os.Exit(0)
			return
		}

	}

	fmt.Println("Wrong number of arguments")
	os.Exit(-1)
}

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/blaskovicz/go-cryptkeeper"
)

var (
	operation string
	key       string
	value     string
)

func init() {
	flag.StringVar(&operation, "o", "", "operation: encrypt or decrypt")
	flag.StringVar(&key, "k", "", "encryption key")
	flag.StringVar(&value, "v", "", "target value of the operation")

	flag.Parse()
}

func main() {
	if operation == "" || key == "" || value == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if err := cryptkeeper.SetCryptKey([]byte(key)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var result string
	var err error
	switch operation {
	case "encrypt":
		result, err = cryptkeeper.Encrypt(value)
	case "decrypt":
		result, err = cryptkeeper.Decrypt(value)
	default:
		fmt.Println("Invalid operation: ", operation)
		os.Exit(1)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(result)
}

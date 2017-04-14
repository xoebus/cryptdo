package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/xoebus/cryptdo"
)

var passphrase = flag.String("passphrase", "", "passphrase for file encryption")

func main() {
	flag.Parse()

	if *passphrase == "" {
		fmt.Println("passphrase must not be empty")
		os.Exit(1)
	}

	filenames := flag.Args()

	if len(filenames) == 0 {
		usage()
	}

	for _, filename := range filenames {
		plaintext, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatalln(err)
		}

		ciphertext, err := cryptdo.Encrypt(plaintext, *passphrase)
		if err != nil {
			log.Fatalln(err)
		}

		encPath := filename + ".enc"
		err = ioutil.WriteFile(encPath, ciphertext, 0400)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(filename, "has been encrypted and placed at", encPath)
	}

	fmt.Println()
	fmt.Println("Please make sure you have the encryption key stored safely and then")
	fmt.Println("delete the original files.")
}

func usage() {
	fmt.Println("usage: cryptdo-bootstrap -passphrase PASSPHRASE FILES...")
	os.Exit(1)
}

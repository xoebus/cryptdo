package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	flags "github.com/jessevdk/go-flags"

	"github.com/xoebus/cryptdo"
)

var opts struct {
	Passphrase string `short:"p" long:"passphrase" description:"passphrase for file encryption" required:"true"`

	Positional struct {
		Files []string `positional-arg-name:"FILE" required:"true"`
	} `positional-args:"yes"`
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	for _, filename := range opts.Positional.Files {
		plaintext, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatalln(err)
		}

		ciphertext, err := cryptdo.Encrypt(plaintext, opts.Passphrase)
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

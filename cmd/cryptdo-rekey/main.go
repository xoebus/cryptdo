package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/xoebus/cryptdo"
)

var old = flag.String("old", "", "old passphrase for file encryption")
var new = flag.String("new", "", "new passphrase for file encryption")

func main() {
	flag.Parse()

	if *old == "" {
		fmt.Println("old passphrase must not be empty")
		os.Exit(1)
	}

	if *new == "" {
		fmt.Println("new passphrase must not be empty")
		os.Exit(1)
	}

	filenames := flag.Args()

	if len(filenames) == 0 {
		usage()
	}

	for _, filename := range filenames {
		oldText, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatalln(err)
		}

		plaintext, err := cryptdo.Decrypt(oldText, *old)
		if err != nil {
			log.Fatalln(err)
		}

		newText, err := cryptdo.Encrypt(plaintext, *new)
		if err != nil {
			log.Fatalln(err)
		}

		newPath := filename + ".new"
		err = ioutil.WriteFile(newPath, newText, 0400)
		if err != nil {
			log.Fatalln(err)
		}

		if err := os.Rename(newPath, filename); err != nil {
			log.Fatalln(err)
		}

		fmt.Println(filename, "has been re-keyed")
	}
}

func usage() {
	fmt.Println("usage: cryptdo-rekey -old OLD-PASSPHRASE -new NEW-PASSPHRASE FILES...")
	os.Exit(1)
}

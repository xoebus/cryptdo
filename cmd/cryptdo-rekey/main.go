package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	flags "github.com/jessevdk/go-flags"

	"github.com/xoebus/cryptdo"
)

var opts struct {
	Old string `short:"o" long:"old-passphrase" description:"old passphrase for file encryption" required:"true"`
	New string `short:"n" long:"new-passphrase" description:"new passphrase for file encryption" required:"true"`
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	encryptedFiles, _ := filepath.Glob("*.enc")

	for _, filename := range encryptedFiles {
		oldText, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatalln(err)
		}

		plaintext, err := cryptdo.Decrypt(oldText, opts.Old)
		if err != nil {
			log.Fatalln(err)
		}

		newText, err := cryptdo.Encrypt(plaintext, opts.New)
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

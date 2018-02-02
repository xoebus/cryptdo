package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	flags "github.com/jessevdk/go-flags"

	"code.xoeb.us/cryptdo/cryptdo"
	"code.xoeb.us/cryptdo/cryptdo/internal/flag"
)

var opts struct {
	Old string `short:"o" long:"old-passphrase" description:"old passphrase for file encryption"`
	New string `short:"n" long:"new-passphrase" description:"new passphrase for file encryption" `

	Extension flag.Ext `short:"e" long:"extension" description:"extension to use for encrypted files" default:".enc"`
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	oldPassphrase := oldPass()
	newPassphrase := newPass()

	encryptedFiles, _ := filepath.Glob("*" + string(opts.Extension))

	for _, filename := range encryptedFiles {
		oldText, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatalln(err)
		}

		plaintext, err := cryptdo.Decrypt(oldText, oldPassphrase)
		if err != nil {
			log.Fatalln(err)
		}

		newText, err := cryptdo.Encrypt(plaintext, newPassphrase)
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

func oldPass() string {
	if opts.Old != "" {
		return opts.Old
	}

	pass, err := cryptdo.ReadPassphrase("old")
	if err != nil {
		log.Fatalln(err)
	}

	return pass
}

func newPass() string {
	if opts.New != "" {
		return opts.New
	}

	pass, err := cryptdo.ReadPassphrase("new")
	if err != nil {
		log.Fatalln(err)
	}

	return pass
}

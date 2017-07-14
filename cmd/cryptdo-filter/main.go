package main

import (
	"io/ioutil"
	"log"
	"os"

	flags "github.com/jessevdk/go-flags"

	"code.xoeb.us/cryptdo"
)

var opts struct {
	Passphrase string `short:"p" long:"passphrase" description:"passphrase for file encryption"`
	Direction  string `short:"d" long:"direction" description:"one of 'clean' or 'smudge'"`
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	pass := passphrase()

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	var output []byte

	switch opts.Direction {
	case "clean":
		output, err = cryptdo.Decrypt(input, pass)
		if err != nil {
			log.Fatalln(err)
		}
	case "smudge":
		output, err = cryptdo.Encrypt(input, pass)
		if err != nil {
			log.Fatalln(err)
		}
	default:
		log.Fatalln("direction must be either 'clean' or 'smudge'")
	}

	_, err = os.Stdout.Write(output)
	if err != nil {
		log.Fatalln(err)
	}
}

func passphrase() string {
	if opts.Passphrase != "" {
		return opts.Passphrase
	}

	pass, err := cryptdo.ReadPassphrase("passphrase")
	if err != nil {
		log.Fatalln(err)
	}

	return pass
}

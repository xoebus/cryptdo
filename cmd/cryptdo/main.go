package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/xoebus/cryptdo"
)

var passphrase = flag.String("passphrase", "", "passphrase for file encryption")

func main() {
	flag.Parse()

	if *passphrase == "" {
		fmt.Println("passphrase must not be empty")
		os.Exit(1)
	}

	command := flag.Args()
	if len(command) < 1 {
		usage()
	}

	encryptedFiles, _ := filepath.Glob("*.enc")

	for _, file := range encryptedFiles {
		ciphertext, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatalln(err)
		}

		plaintext, err := cryptdo.Decrypt(ciphertext, *passphrase)
		if err != nil {
			log.Fatalln(err)
		}

		err = ioutil.WriteFile(decryptedName(file), plaintext, 0400)
		if err != nil {
			log.Fatalln(err)
		}
	}

	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}

	for _, file := range encryptedFiles {
		plaintext, err := ioutil.ReadFile(decryptedName(file))
		if err != nil {
			log.Fatalln(err)
		}

		ciphertext, err := cryptdo.Encrypt(plaintext, *passphrase)
		if err != nil {
			log.Fatalln(err)
		}

		newPath := file + ".new"
		err = ioutil.WriteFile(newPath, ciphertext, 0400)
		if err != nil {
			log.Fatalln(err)
		}

		if err = os.Rename(newPath, file); err != nil {
			log.Fatalln(err)
		}

		if err = os.Remove(decryptedName(file)); err != nil {
			log.Fatalln(err)
		}
	}
}

func decryptedName(file string) string {
	return file[:strings.LastIndex(file, ".enc")]
}

func usage() {
	fmt.Println("usage: cryptdo -passphrase PASSPHRASE -- command ...")
	os.Exit(1)
}

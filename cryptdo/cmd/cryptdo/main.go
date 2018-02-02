package main

import (
	"crypto/sha256"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	flags "github.com/jessevdk/go-flags"

	"code.xoeb.us/cryptdo/cryptdo"
	"code.xoeb.us/cryptdo/cryptdo/internal/flag"
)

var opts struct {
	Passphrase string   `short:"p" long:"passphrase" description:"passphrase for file encryption"`
	Extension  flag.Ext `short:"e" long:"extension" description:"extension to use for encrypted files" default:".enc"`

	Positional struct {
		Command string   `positional-arg-name:"COMMAND" required:"true"`
		Args    []string `positional-arg-name:"ARG"`
	} `positional-args:"yes"`
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	pass := passphrase()
	ext := string(opts.Extension)

	encryptedFiles, _ := filepath.Glob("*" + ext)

	fingerprints := make(map[string][32]byte)

	for _, file := range encryptedFiles {
		ciphertext, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatalln(err)
		}

		plaintext, err := cryptdo.Decrypt(ciphertext, pass)
		if err != nil {
			log.Fatalln(err)
		}

		fingerprints[file] = fingerprint(plaintext)

		err = ioutil.WriteFile(decryptedName(file, ext), plaintext, 0600)
		if err != nil {
			log.Fatalln(err)
		}
	}

	cmd := exec.Command(opts.Positional.Command, opts.Positional.Args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	exitStatus := run(cmd)

	for _, file := range encryptedFiles {
		plaintext, err := ioutil.ReadFile(decryptedName(file, ext))
		if err != nil {
			log.Fatalln(err)
		}

		needsEncrypt := false
		if fingerprint(plaintext) != fingerprints[file] {
			needsEncrypt = true
		}

		if needsEncrypt {
			ciphertext, err := cryptdo.Encrypt(plaintext, pass)
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
		}

		if err = os.Remove(decryptedName(file, ext)); err != nil {
			log.Fatalln(err)
		}
	}

	os.Exit(exitStatus)
}

func decryptedName(file, ext string) string {
	return file[:strings.LastIndex(file, ext)]
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

func run(cmd *exec.Cmd) (status int) {
	if err := cmd.Run(); err != nil {
		log.Println("cryptdo: command failed:", err)

		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				return status.ExitStatus()
			}
		}

		return 1
	}

	return 0
}

func fingerprint(contents []byte) [32]byte {
	return sha256.Sum256(contents)
}

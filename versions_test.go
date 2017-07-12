package cryptdo_test

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"

	"code.xoeb.us/cryptdo"
)

const expectedPath = "testdata/encrypt-this.txt"

var versions = []struct {
	version    string
	file       string
	passphrase string
}{
	{"1.0", "1_0.txt.enc", "password"},
}

func TestOldVersions(t *testing.T) {
	expected, err := ioutil.ReadFile(expectedPath)
	if err != nil {
		t.Errorf("failed to open expected data file: %s", err)
	}

	for _, version := range versions {
		file := filepath.Join("testdata", version.file)
		ciphertext, err := ioutil.ReadFile(file)
		if err != nil {
			t.Errorf("failed to open encrypted data file (version %s): %s", version.version, err)
		}

		actual, err := cryptdo.Decrypt(ciphertext, version.passphrase)
		if err != nil {
			t.Errorf("failed to decrypt (version %s): %s", version.version, err)
		}

		if !bytes.Equal(actual, expected) {
			// I couldn't find a good way to show this failure so you're just going
			// to have to add it when this fails.
			t.Errorf("version %s decrypted data did not match expected data!", version.version, actual, expected)
		}
	}
}

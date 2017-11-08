package cryptdo

import (
	"bytes"
	"testing"
	"testing/quick"
)

var config = &quick.Config{
	MaxCount: 5, // The encrypt/decrypt cycle is fairly slow.
}

func TestRoundtrip(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping roundtrip test in short mode")
	}

	roundtrip := func(plaintext []byte, passphrase string) bool {
		ciphertext, err := Encrypt(plaintext, passphrase)
		if err != nil {
			return false
		}

		plain, err := Decrypt(ciphertext, passphrase)
		if err != nil {
			return false
		}

		return bytes.Equal(plaintext, plain)
	}

	if err := quick.Check(roundtrip, config); err != nil {
		t.Error(err)
	}
}

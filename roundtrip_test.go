package cryptdo

import (
	"bytes"
	"testing"
	"testing/quick"
)

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

	if err := quick.Check(roundtrip, nil); err != nil {
		t.Error(err)
	}
}

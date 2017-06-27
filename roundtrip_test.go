package cryptdo

import (
	"bytes"
	"testing"
	"testing/quick"

	"code.xoeb.us/cryptdo"
)

func TestRoundtrip(t *testing.T) {
	roundtrip := func(plaintext []byte, passphrase string) bool {
		ciphertext, err := cryptdo.Encrypt(plaintext, passphrase)
		if err != nil {
			return false
		}

		plain, err := cryptdo.Decrypt(ciphertext, passphrase)
		if err != nil {
			return false
		}

		return bytes.Equal(plaintext, plain)
	}

	if err := quick.Check(roundtrip, nil); err != nil {
		t.Error(err)
	}
}

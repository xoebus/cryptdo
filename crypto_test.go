package cryptdo_test

import (
	"testing"

	"github.com/golang/protobuf/proto"

	"code.xoeb.us/cryptdo"
	cryptdopb "code.xoeb.us/cryptdo/proto"
)

func TestCurrentCrypto(t *testing.T) {
	passphrase := "hunter2"
	plaintext := []byte("something secret")

	ciphertext, err := cryptdo.Encrypt(plaintext, passphrase)
	if err != nil {
		t.Error("got error while encrypting:", err)
	}

	message := &cryptdopb.Message{}
	if err := proto.Unmarshal(ciphertext, message); err != nil {
		t.Error("got error while unmarshaling", err)
	}

	if iterations := message.GetIterations(); iterations != 100000 {
		t.Errorf("iterations was incorrect, got: %d, want: %d", iterations, 100000)
	}

	if salt := message.GetSalt(); len(salt) != 48 {
		t.Errorf("salt length was incorrect, got: %d, want: %d", len(salt), 48)
	}
}

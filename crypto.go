package cryptdo

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"

	"github.com/golang/protobuf/proto"

	"code.xoeb.us/cryptdo/cryptdopb"
)

const currentVersion = 1

// ErrEmptyMessage is caused by trying to decrypt and empty function.
var ErrEmptyMessage = errors.New("cryptdo: empty message")

// Encrypt takes a plaintext blob and encrypts it with a key based on the
// passphrase provided. The byte slice returned is suitable for passing in to
// the Decrypt function with the same passphrase in order to retrieve the data.
//
// The output format and internal details of the cryptography performed is
// documented in the associated protocol buffers file.
//
// Due to the use of authenticated encryption we need to read the entire
// plaintext into memory. Therefore it is recommended to only use this for
// smaller plaintexts.
func Encrypt(plaintext []byte, passphrase string) ([]byte, error) {
	// Encryption always uses the current cryptography version.
	v, err := lookup(currentVersion)
	if err != nil {
		return nil, err
	}

	message, err := v.encrypt(plaintext, passphrase)
	if err != nil {
		return nil, err
	}

	return proto.Marshal(message)
}

// Decrypt decrypts a piece of ciphertext which was encrypted with the Encrypt
// function. The original plaintext is returned if no error occured during the
// decryption. It supports passing in ciphertext which was made with previous
// versions of the library.
//
// The output format and internal details of the cryptography performed is
// documented in the associated protocol buffers file.
//
// Due to the use of authenticated encryption we need to read the entire
// ciphertext into memory. Therefore it is recommended to only use this for
// smaller plaintexts.
func Decrypt(ciphertext []byte, passphrase string) ([]byte, error) {
	if len(ciphertext) == 0 {
		return nil, ErrEmptyMessage
	}

	var message cryptdopb.Message
	if err := proto.Unmarshal(ciphertext, &message); err != nil {
		return nil, err
	}

	v, err := lookup(message.GetVersion())
	if err != nil {
		return nil, err
	}

	return v.decrypt(&message, passphrase)
}

// randomBytes returns a byte slice (of the requested length) filled with
// cryptographically secure random bytes.
func randomBytes(count int) ([]byte, error) {
	bs := make([]byte, count)

	if _, err := io.ReadFull(rand.Reader, bs); err != nil {
		return nil, err
	}

	return bs, nil
}

// InvalidNonceError is caused by a mismatch between the expected nonce length
// from the encryption algorithm and the actual nonce length provided in the
// message.
type InvalidNonceError struct {
	expected int
	actual   int
}

func (i *InvalidNonceError) Error() string {
	return fmt.Sprintf("cryptdo: message nonce has incorrect size (expected %d, got: %d)", i.expected, i.actual)
}

// UnknownVersionError is caused when a message is too new for the executed
// version to understand. This will often be caused by a mismatch between two
// cryptography versions operating on the same data.
type UnknownVersionError struct {
	version int
}

func (u *UnknownVersionError) Error() string {
	return fmt.Sprintf("cryptdo: message has incompatible version (expected %d or below, got: %d)", currentVersion, u.version)
}

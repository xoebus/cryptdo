package cryptdo

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"errors"
	"fmt"
	"io"

	"github.com/golang/protobuf/proto"
	"golang.org/x/crypto/pbkdf2"

	"code.xoeb.us/cryptdo/cryptdopb"
)

const (
	currentVersion = 1

	// Key Derivation
	iterations = 100000

	// Encryption
	keySize   = 32
	nonceSize = 12
)

var (
	hashAlg  = sha512.New384
	saltSize = hashAlg().Size()
)

var ErrEmptyMessage = errors.New("cryptdo: empty message")

func Encrypt(plaintext []byte, passphrase string) ([]byte, error) {
	salt, err := randomBytes(saltSize)
	if err != nil {
		return nil, err
	}

	key := derivedKey(passphrase, salt, iterations)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	nonce, err := randomBytes(nonceSize)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	message := &cryptdopb.Message{
		Version:    currentVersion,
		Salt:       salt,
		Nonce:      nonce,
		Ciphertext: ciphertext,
	}

	return proto.Marshal(message)
}

func Decrypt(ciphertext []byte, passphrase string) ([]byte, error) {
	if len(ciphertext) == 0 {
		return nil, ErrEmptyMessage
	}

	message := &cryptdopb.Message{}
	if err := proto.Unmarshal(ciphertext, message); err != nil {
		return nil, err
	}

	key := derivedKey(passphrase, message.Salt, iterations)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if aesgcm.NonceSize() != len(message.Nonce) {
		return nil, &InvalidNonceError{
			expected: aesgcm.NonceSize(),
			actual:   len(message.Nonce),
		}
	}

	return aesgcm.Open(nil, message.Nonce, message.Ciphertext, nil)
}

func derivedKey(passphrase string, salt []byte, iters int) []byte {
	return pbkdf2.Key([]byte(passphrase), salt, iters, keySize, hashAlg)
}

func randomBytes(count int) ([]byte, error) {
	bs := make([]byte, count)

	if _, err := io.ReadFull(rand.Reader, bs); err != nil {
		return nil, err
	}

	return bs, nil
}

type InvalidNonceError struct {
	expected int
	actual   int
}

func (i *InvalidNonceError) Error() string {
	return fmt.Sprintf("cryptdo: message nonce has incorrect size (expected %d, got: %d)", i.expected, i.actual)
}

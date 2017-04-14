package cryptdo

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"errors"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

const (
	// Key Derivation
	iterations = 100000
	saltSize   = 32

	// Encryption
	keySize   = 32
	nonceSize = 12
)

var ErrShortCiphertext = errors.New("cryptdo: ciphertext too short")

func Encrypt(plaintext []byte, passphrase string) ([]byte, error) {
	salt, err := randomBytes(saltSize)
	if err != nil {
		return nil, err
	}

	key := derivedKey(passphrase, salt)
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

	return join(salt, nonce, ciphertext), nil
}

func Decrypt(ciphertext []byte, passphrase string) ([]byte, error) {
	if len(ciphertext) < (nonceSize + saltSize) {
		return nil, ErrShortCiphertext
	}

	salt := ciphertext[:saltSize]
	nonce := ciphertext[saltSize : saltSize+nonceSize]
	pure := ciphertext[saltSize+nonceSize:]

	key := derivedKey(passphrase, salt)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return aesgcm.Open(nil, nonce, pure, nil)
}

func derivedKey(passphrase string, salt []byte) []byte {
	return pbkdf2.Key([]byte(passphrase), salt, iterations, keySize, sha512.New384)
}

func join(slices ...[]byte) []byte {
	return bytes.Join(slices, []byte{})
}

func randomBytes(count int) ([]byte, error) {
	bs := make([]byte, count)

	if _, err := io.ReadFull(rand.Reader, bs); err != nil {
		return nil, err
	}

	return bs, nil
}

package cryptdo

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"io"

	"github.com/golang/protobuf/proto"
	"golang.org/x/crypto/pbkdf2"

	"github.com/xoebus/cryptdo/cryptdopb"
)

const (
	// Key Derivation
	iterations = 100000
	saltSize   = 32

	// Encryption
	keySize   = 32
	nonceSize = 12
)

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
		Iterations: iterations,
		Salt:       salt,
		Nonce:      nonce,
		Ciphertext: ciphertext,
	}

	return proto.Marshal(message)
}

func Decrypt(ciphertext []byte, passphrase string) ([]byte, error) {
	message := &cryptdopb.Message{}
	if err := proto.Unmarshal(ciphertext, message); err != nil {
		return nil, err
	}

	key := derivedKey(passphrase, message.Salt, int(message.Iterations))
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return aesgcm.Open(nil, message.Nonce, message.Ciphertext, nil)
}

func derivedKey(passphrase string, salt []byte, iters int) []byte {
	return pbkdf2.Key([]byte(passphrase), salt, iters, keySize, sha512.New384)
}

func randomBytes(count int) ([]byte, error) {
	bs := make([]byte, count)

	if _, err := io.ReadFull(rand.Reader, bs); err != nil {
		return nil, err
	}

	return bs, nil
}

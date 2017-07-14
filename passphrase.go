package cryptdo

import (
	"errors"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// ErrEmptyPassphase is returned when the user does not type a passphrase
// before pressing return.
var ErrEmptyPassphase = errors.New("cryptdo: passphrase must not be empty")

// ReadPassphrase presents a prompt to the user before allowing them to type a
// passphrase without echoing the characters they type back to the terminal. It
// returns ErrEmptyPassphase if the user does not enter a passphrase.
func ReadPassphrase(prompt string) (string, error) {
	if _, err := fmt.Fprintf(os.Stderr, "%s: ", prompt); err != nil {
		return "", err
	}
	defer fmt.Println()

	input, err := terminal.ReadPassword(syscall.Stdin)
	if err != nil {
		return "", err
	}

	if len(input) == 0 {
		return "", ErrEmptyPassphase
	}

	return string(input), nil
}

// +build gofuzz

package cryptdo

func Fuzz(data []byte) int {
	if _, err := Decrypt(data, "fuzzing"); err != nil {
		return 0
	}

	return 1
}

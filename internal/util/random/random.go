package random

import (
	"crypto/rand"
)

const (
	letters    = `0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz:;.,/?\|-@£#$%^&*`
	lettersHex = `0123456789abcdef`
)

func Bytes(n int) []byte {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return nil
	}
	return b
}

func String(n int) string {
	bytes := Bytes(n)
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes)
}

func HexString(n int) string {
	bytes := Bytes(n)
	for i, b := range bytes {
		bytes[i] = lettersHex[b%byte(len(lettersHex))]
	}
	return string(bytes)
}

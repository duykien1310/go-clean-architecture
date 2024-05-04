package util

import (
	"crypto/rand"
	"io"
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

func GenerateOTPCode(otpLength int) (string, error) {
	b := make([]byte, otpLength)
	n, err := io.ReadAtLeast(rand.Reader, b, otpLength)
	if err != nil {
		return "", err
	}

	if n != otpLength {
		return "", err
	}

	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}

	return string(b), nil
}

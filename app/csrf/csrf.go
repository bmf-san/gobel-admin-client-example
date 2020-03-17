package csrf

import (
	"crypto/rand"
	"fmt"
)

const tokenLength = 16

// GenerateToken returns a unique token.
func GenerateToken() (string, error) {
	b := make([]byte, tokenLength)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

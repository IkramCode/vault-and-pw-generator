package crypto

import (
	"crypto/rand"
	"crypto/sha256"

	"crypto/pbkdf2"
)

func MasterKey(password string, salt []byte) ([]byte, error) {
	const (
		saltSize   = 16
		iterations = 4096
		keylength  = 32
	)
	key, err := pbkdf2.Key(sha256.New, password, salt, iterations, keylength)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}

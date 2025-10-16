package gen

import (
	"crypto/rand"
	"math/big"
)

func GenPassword(length int, includeUpper bool, includeNums bool, includeSymbols bool) (string, error) {
	charset := "abcdefghijklmnopqrstuvwxyz"
	if includeUpper {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if includeNums {
		charset += "1234567890"
	}
	if includeSymbols {
		charset += "!@#$%^&*()-_=+[]{}<>?/|"

	}
	password := make([]byte, length)
	for i := range password {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[n.Int64()]
	}
	return string(password), nil
}

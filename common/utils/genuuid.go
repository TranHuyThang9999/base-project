package utils

import (
	"crypto/rand"
	"math/big"
	"time"
)

func GenUUID() int64 {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	randomBits, _ := rand.Int(rand.Reader, big.NewInt(1<<20)) // 20 bits
	return (timestamp << 20) | randomBits.Int64()
}

func GenTime() time.Time {
	return time.Now().UTC()
}

func GenPassWord() int64 {
	max := big.NewInt(99999999 - 10000000 + 1)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0
	}

	return n.Int64() + 10000000
}

func GenPasswordString(length int) string {
	if length < 8 {
		length = 8
	}

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+"
	password := make([]byte, length)

	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return ""
		}
		password[i] = charset[index.Int64()]
	}

	return string(password)
}

func GenOTP(length int) int64 {
	if length < 4 {
		length = 4
	} else if length > 10 {
		length = 10
	}

	max := big.NewInt(1)
	for i := 0; i < length; i++ {
		max.Mul(max, big.NewInt(10))
	}

	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0
	}

	return n.Int64()
}

func ValidatePassword(password string) bool {
	var (
		hasUpper  bool
		hasLower  bool
		hasNumber bool
		hasSymbol bool
	)

	for _, char := range password {
		switch {
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= 'a' && char <= 'z':
			hasLower = true
		case char >= '0' && char <= '9':
			hasNumber = true
		case char >= '!' && char <= '/':
			hasSymbol = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSymbol
}

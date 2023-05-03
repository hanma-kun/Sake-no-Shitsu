package utils

import (
	"fmt"

	"gitlab.com/avarf/getenvs"
	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) string {
	salt := []byte(getenvs.GetEnvString("SALTVALUE", "dTVHI5XRuLL1YAzBysgy4MYyKE1GBRqd"))
	time := uint32(2)
	memory := uint32(64 * 1024)
	threads := uint8(4)
	keyLength := uint32(32)
	hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLength)

	return fmt.Sprintf("%x", hash)
}

func VerifyPassword(password, hashedPassword string) bool {
	return hashedPassword == HashPassword(password)
}

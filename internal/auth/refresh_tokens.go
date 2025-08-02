package auth

import (
	"crypto/rand"
	"encoding/hex"
)

func MakeRefreshToken() string {
	dat := make([]byte, 32)
	rand.Read(dat)
	token := hex.EncodeToString(dat)
	return token
}

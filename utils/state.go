package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomState() string {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		panic("Failed to generate random state")
	}
	return base64.URLEncoding.EncodeToString(bytes)
}

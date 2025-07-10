package utils

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"strings"
)

func GenerateRandToken(len int16) string {
	tokenBytes := make([]byte, len)
	_, err := rand.Read(tokenBytes)

	if err != nil {
		panic(err)
	}

	token := base64.StdEncoding.EncodeToString(tokenBytes)

	return token
}

func GenerateHexToken(parts int) string {
	partLength := 4
	totalLength := parts * partLength

	bytes := make([]byte, (totalLength+1)/2)
	_, err := rand.Read(bytes)

	if err != nil {
		panic(err)
	}

	hexStr := hex.EncodeToString(bytes)[:totalLength]

	var result []string

	for i := 0; i < totalLength; i += partLength {
		result = append(result, hexStr[i:i+partLength])
	}

	return strings.Join(result, "-")
}

package helper

import (
	"math/rand"
)

type Str struct{}

// Get a non-cryptographically secure random string.
func (str *Str) Random(length int) string {
	const characters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)
	for i := range result {
		result[i] = characters[rand.Intn(len(characters))]
	}
	return string(result)
}

func NewStr() *Str {
	return &Str{}
}

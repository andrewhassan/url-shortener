package utils

import (
	"math/rand"
	"time"
)

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

// CreateSlug creates a random slug of length `length` using alphanumeric characters
func CreateSlug(length int) string {
	rand.Seed(time.Now().UnixNano())
	result := make([]rune, length)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

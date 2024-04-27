package utils

import "math/rand"

const (
	chars = "abcdefghijklmnopqrstuvwxyz0123456789"
)

func GenerateID(length int) string {
	id := make([]byte, length)
	for i := range id {
		id[i] = chars[rand.Intn(len(chars))]
	}
	return string(id)
}

func isCharValid(char rune) bool {
	return char >= 'a' && char <= 'z' || char >= '0' && char <= '9'
}

func IsIDValid(id string) bool {
	if len(id) != 7 {
		return false
	}
	for _, char := range id {
		if !isCharValid(char) {
			return false
		}
	}
	return true
}

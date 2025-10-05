package utils

import (
	"errors"
	"strings"
)

var (
	ErrInvalidChar = errors.New("invalid character in base62 string")
)

const (
	base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base        = 62
)

// Encode converts a number to a base62 string
func Encode(num int64) string {
	if num == 0 {
		return "0"
	}

	var result strings.Builder
	for num > 0 {
		remainder := num % base
		result.WriteByte(base62Chars[remainder])
		num = num / base
	}

	// Reverse the string since we built it backwards
	encoded := result.String()
	runes := []rune(encoded)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

// Decode converts a base62 string back to a number
func Decode(str string) (int64, error) {
	var result int64
	var power int64 = 1

	// Process from right to left
	for i := len(str) - 1; i >= 0; i-- {
		char := str[i]
		value := int64(strings.IndexByte(base62Chars, char))
		if value == -1 {
			return 0, ErrInvalidChar
		}
		result += value * power
		power *= base
	}

	return result, nil
}

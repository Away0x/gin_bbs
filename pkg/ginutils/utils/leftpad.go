package utils

import (
	"fmt"
	"strings"
)

// LeftPad - LeftPad("a", 4, '0) => 000a
func LeftPad(s string, n int, r rune) (string, error) {
	if n < 0 {
		return "", fmt.Errorf("invalid length %d", n)
	}
	if len(s) > n {
		return s, nil
	}
	return strings.Repeat(string(r), n-len(s)) + s, nil
}

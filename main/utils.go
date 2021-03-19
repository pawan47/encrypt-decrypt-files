// This file contains util functions

package main

import (
	"bufio"
	"os"
	"strings"
)

// padPassKey pads pass key to make 16 bytes (AES-128), 24 bytes (AES-192)
// or 32 bytes (AES-256)
func padPassKey(key string) string {
	l := len(key)
	if l == 0 {
		return ""
	}
	if l <= 16 {
		return strings.Repeat(key, 16)[:16]
	} else if l > 16 && l <= 24 {
		return strings.Repeat(key, 2)[:24]
	} else if l < 32 {
		return strings.Repeat(key, 2)[:32]
	}
	return key[:32]
}

// scanStr scans string from std
func scanStr() string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

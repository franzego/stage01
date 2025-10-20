package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"time"
	"unicode"
)

// func to check for length
func Length(s string) int {
	lowerS := strings.ToLower(s)
	return len(lowerS)
}

// func to check for palindrome
func IsPalindrome(s string) bool {
	var newString strings.Builder
	// check for alphanumeric characters
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			newString.WriteRune(unicode.ToLower(r))
		}

	}
	st := newString.String()
	runes := []rune(st)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

// func to check fo unique xter
func UniqueChar(s string) int {
	set := make(map[rune]bool)
	for _, char := range s {
		set[char] = true
	}
	return len(set)
}

// func to check number of words separated by white spaces
func WordCount(s string) int {

	colStrings := strings.Split(s, " ")
	return len(colStrings)

}

// func to generate sha256 from a string
func Sha256Encoding(s string) string {
	data := s

	hash := sha256.New()
	hash.Write([]byte(data))
	encodedHash := hash.Sum(nil)
	return hex.EncodeToString(encodedHash) // use hex encoding
}

// func to check for a character frequency
func CharFrequency(s string) map[string]int {
	hashMap := make(map[string]int)
	for _, r := range s {
		char := string(r)
		hashMap[char]++
	}
	// for r, count := range hashMap {
	// 	fmt.Printf("'%c': %d\n", r, count)
	// }
	return hashMap

}

// func to generate current time
func Timestamp() string {
	return time.Now().UTC().Format(time.RFC3339)
}

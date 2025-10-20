package internal

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/franzego/stage01/dto"
	"github.com/franzego/stage01/utils"
	"github.com/gin-gonic/gin"
)

func PostString(c *gin.Context) {
	// decode
	var req dto.Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	// validate
	if strings.TrimSpace(req.Value) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	// check if the value exits in my slice
	for _, u := range HardcodedStrings {
		if u.Value == req.Value {
			c.JSON(http.StatusConflict, gin.H{"error": "String already exists in the system"})
			c.Abort()
			return
		}
	}
	// status 409
	// var resp dto.Resp
	resp := &dto.Resp{
		Id:    utils.Sha256Encoding(req.Value),
		Value: req.Value,
		Props: dto.Props{
			Length:                utils.Length(req.Value),
			IsPalindrome:          utils.IsPalindrome(req.Value),
			UniqueCharacters:      utils.UniqueChar(req.Value),
			WordCount:             utils.WordCount(req.Value),
			Sha256Hash:            utils.Sha256Encoding(req.Value),
			CharacterFrequencyMap: utils.CharFrequency(req.Value),
		},
		CreatedAt: utils.Timestamp(),
	}
	c.JSON(http.StatusCreated, resp)
	HardcodedStrings = append(HardcodedStrings, *resp)

}

// func to get a particular string
func GetString(c *gin.Context) {
	str := c.Param("string_value")
	for _, s := range HardcodedStrings {
		if s.Value == str {
			c.JSON(http.StatusOK, s)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Requested string not found"})
}

// func to execute queries
func GetQueries(c *gin.Context) {
	var filtered []dto.Resp

	// Get query params
	isPalindromeQuery := c.Query("is_palindrome")
	minLengthQuery := c.Query("min_length")
	maxLengthQuery := c.Query("max_length")
	wordCountQuery := c.Query("word_count")
	containsChar := c.Query("contains_character")

	// Parsed values
	var (
		isPalindrome                                              bool
		minLength, maxLength, wordCount                           int
		hasIsPalindrome, hasMinLength, hasMaxLength, hasWordCount bool
	)

	if isPalindromeQuery != "" {
		val, err := strconv.ParseBool(isPalindromeQuery)
		if err == nil {
			isPalindrome = val
			hasIsPalindrome = true
		}
	}
	if minLengthQuery != "" {
		val, err := strconv.Atoi(minLengthQuery)
		if err == nil {
			minLength = val
			hasMinLength = true
		}
	}
	if maxLengthQuery != "" {
		val, err := strconv.Atoi(maxLengthQuery)
		if err == nil {
			maxLength = val
			hasMaxLength = true
		}
	}
	if wordCountQuery != "" {
		val, err := strconv.Atoi(wordCountQuery)
		if err == nil {
			wordCount = val
			hasWordCount = true
		}
	}

	// Filter logic
	for _, s := range HardcodedStrings {
		if hasIsPalindrome && s.Props.IsPalindrome != isPalindrome {
			continue
		}
		if hasMinLength && s.Props.Length < minLength {
			continue
		}
		if hasMaxLength && s.Props.Length > maxLength {
			continue
		}
		if hasWordCount && s.Props.WordCount != wordCount {
			continue
		}
		if containsChar != "" && !strings.Contains(s.Value, containsChar) {
			continue
		}

		filtered = append(filtered, s)
	}

	if len(filtered) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No strings matched your filters"})
		return
	}
	response := gin.H{
		"data":  filtered,
		"count": len(filtered),
		"filters_applied": gin.H{
			"is_palindrome":      isPalindrome,
			"min_length":         minLength,
			"max_length":         maxLength,
			"word_count":         wordCount,
			"contains_character": containsChar,
		},
	}

	c.JSON(http.StatusOK, response)
}

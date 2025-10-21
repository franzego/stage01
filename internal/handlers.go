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

	// Parsed values from the string queries to their respective types. Also doing it safely by checking if they are empty
	// created variables to hold these values
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
		c.JSON(http.StatusBadRequest, gin.H{"message": "No strings matched your filters"})
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

// func for natural language processing
func GetByNaturalLanguage(c *gin.Context) {
	qString := c.Query("query")

	if strings.TrimSpace(qString) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'query' is required"})
		return
	}
	f, err := ParseNatrualLanguage(qString)
	if err != nil {
		c.JSON(400, gin.H{"error": "string not found"})
		return
	}

	var filtered []dto.Resp

	for _, s := range HardcodedStrings {
		// Match Palindrome
		if f.IsPalindrome && !s.Props.IsPalindrome {
			continue
		}

		// Match min length
		if f.MinLength > 0 && s.Props.Length < f.MinLength {
			continue
		}

		// Match max length
		if f.MaxLength > 0 && s.Props.Length > f.MaxLength {
			continue
		}

		// Match word count
		if f.WordCount > 0 && s.Props.WordCount != f.WordCount {
			continue
		}

		// Match contains character
		if f.ContainsChar != "" && !strings.Contains(s.Value, f.ContainsChar) {
			continue
		}

		filtered = append(filtered, s)
	}
	if f.MinLength > 0 && f.MaxLength > 0 && f.MinLength > f.MaxLength {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Query parsed but resulted in conflicting filters",
		})
		return
	}

	if len(filtered) == 0 {
		c.JSON(404, gin.H{"message": "No strings matched your filters"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":  filtered,
		"count": len(filtered),
		"interpreted_query": gin.H{
			"original":       qString,
			"parsed_filters": f,
		},
	})

}

// func to delete a particular string
func DeleteString(c *gin.Context) {
	str := c.Param("string_value")
	indextoRemove := -1
	for i, s := range HardcodedStrings {
		if s.Value == str {
			indextoRemove = i
			break
		}
	}
	if indextoRemove == -1 {
		c.JSON(http.StatusNotFound, gin.H{"Message": "string does not exist in the system"})
		return
	}
	// Safe to remove now
	HardcodedStrings = append(HardcodedStrings[:indextoRemove], HardcodedStrings[indextoRemove+1:]...)
	c.JSON(http.StatusNoContent, gin.H{})

}

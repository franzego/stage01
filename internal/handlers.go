package internal

import (
	"net/http"
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

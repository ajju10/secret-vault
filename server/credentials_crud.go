package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = map[string]string{
	"api_key":  "abcdefghijk",
	"api_key1": "abcdefghijk",
	"api_key2": "abcdefghijk",
	"api_key3": "abcdefghijk",
	"api_key4": "abcdefghijk",
}

func getCredentials(c *gin.Context) {
	c.JSON(http.StatusOK, db)
}

func setCredentials(c *gin.Context) {
	data := &Credentials{}
	err := c.BindJSON(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not parse request body",
		})
		return
	}

	db[data.Key] = data.Value
	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully added credentials",
	})
}

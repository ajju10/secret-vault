package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getCredentials(c *gin.Context) {
	username := c.Param("username")
	key := c.Param("key")
	user, err := findUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User does not exist",
		})
		return
	}

	res, err := getCredentialsFromDb(user.UID, key)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func setCredentials(c *gin.Context) {
	username := c.Param("username")
	data := &Credentials{}
	err := c.BindJSON(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not parse request body",
		})
		return
	}

	user, err := findUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User does not exist",
		})
		return
	}

	err = insertCredentialsInDb(user.UID, data.Key, data.Value)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Credentials saved successfully",
	})
}

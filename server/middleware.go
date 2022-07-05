package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, "+
				"accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func verifyTokenMiddleware(c *gin.Context) {
	token, ok := getToken(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}

	id, username, err := validateToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}

	c.Set("id", id)
	c.Set("username", username)
	c.Writer.Header().Set("Authorization", "Bearer "+token)
	c.Next()
}

func getSession(c *gin.Context) (string, string, bool) {
	id, ok := c.Get("id")
	if !ok {
		return "", "", false
	}

	username, ok := c.Get("username")
	if !ok {
		return "", "", false
	}

	return id.(string), username.(string), true
}

func getToken(c *gin.Context) (string, bool) {
	authValue := c.GetHeader("Authorization")
	arr := strings.Split(authValue, " ")
	if len(arr) != 2 {
		return "", false
	}

	authType := strings.Trim(arr[0], "\n\r\t")
	if strings.ToLower(authType) != strings.ToLower("Bearer") {
		return "", false
	}

	return strings.Trim(arr[1], "\n\t\r"), true
}

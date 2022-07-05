package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	apiRouter := router.Group("/api/v1")

	apiRouter.POST("/signup", signup)

	apiRouter.POST("/login", login)

	apiRouter.Use(verifyTokenMiddleware)

	apiRouter.GET("/credentials/get/:username/:key", getCredentials)

	apiRouter.POST("/credentials/add/:username", setCredentials)

	router.Run()
}

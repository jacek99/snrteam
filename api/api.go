package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(router *gin.Engine) {
	// REST API
	rest := router.Group("/snrteam")
	{
		rest.GET("/api/users", getAllUsers)
	}

	// HTML APIs
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	// admin tasks
	router.GET("/healthcheck", adminHealthCheck)
}

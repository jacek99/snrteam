package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const DEFAULT_LANGUAGE = "en-us"

var SUPPORTED_LANGES = map[string]interface{}{
	"en-US": nil,
	"es": nil,
}
const HEADER_ACCEPT_LANGUAGE = "Accept-Language"


type RestError struct {
	Message string
	EntityType string
	EntityField string
	EntityId interface{}
}

func InitRouter(router *gin.Engine) {
	// REST API
	rest := router.Group("/snrteam")
	{
		rest.GET("/api/users", getAllUsers)
		rest.GET("/api/users/:UserId", getUser)
	}

	// HTML APIs
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	// admin tasks
	router.GET("/healthcheck", adminHealthCheck)
}

// parses
func getRequestLanguage(c *gin.Context) string {
	lang := c.Request.Header.Get(HEADER_ACCEPT_LANGUAGE)
	if lang == "" {
		return DEFAULT_LANGUAGE
	} else {
		_, ok := SUPPORTED_LANGES[lang]
		if ok {
			return lang
		} else {
			return DEFAULT_LANGUAGE
		}
	}
}
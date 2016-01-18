package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/jacek99/snrteam/common"
	"log"
	"github.com/nicksnyder/go-i18n/i18n"
)

const DEFAULT_LANGUAGE = "en-us"

var SUPPORTED_LANGUAGES = map[string]interface{}{
	"en-US": nil,
	"es": nil,
}
const HEADER_ACCEPT_LANGUAGE = "Accept-Language"

func InitRouter(router *gin.Engine) {
	// REST API
	rest := router.Group("/snrteam")
	{
		rest.GET("/api/users", getAllUsers)
		rest.GET("/api/users/:UserName", getUser)

		rest.POST("/api/users", saveUser)
	}

	// HTML APIs
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	// admin tasks
	router.GET("/healthcheck", adminHealthCheck)

	admin := router.Group("/admin")
	{
		admin.POST("/tasks/test/truncate", truncate)
	}
}

// parses
func getRequestLanguage(c *gin.Context) string {
	lang := c.Request.Header.Get(HEADER_ACCEPT_LANGUAGE)
	if lang == "" {
		return DEFAULT_LANGUAGE
	} else {
		_, ok := SUPPORTED_LANGUAGES[lang]
		if ok {
			return lang
		} else {
			return DEFAULT_LANGUAGE
		}
	}
}

// returns the correct i18n TranslateFunc for the current request
func getI18n(c *gin.Context) i18n.TranslateFunc {
	T, _ := i18n.Tfunc(getRequestLanguage(c))
	return T
}

// standard error handler
func handleError(c *gin.Context, err error) {
	log.Println(err)
	switch t := err.(type) {
	case common.NotFoundError:
		c.JSON(http.StatusNotFound, err)
	case common.ConflictError:
		c.JSON(http.StatusConflict, err)
	case common.GenericError:
		c.JSON(http.StatusInternalServerError, err)
	default:
		_ = t
		c.JSON(http.StatusInternalServerError, "Server error. Please consult support")
	}
}

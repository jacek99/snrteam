package server


import (
	"net/http"
	"time"
	"os"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jacek99/snrteam/common"
)


var Server *http.Server
var Router *gin.Engine

func init() {
	Router = gin.Default()

	port := os.Getenv(common.ENV_HTTP_PORT)
	if port == "" {
		port = "8080"
	}

	Server = &http.Server{
		Addr:           fmt.Sprintf(":%s",port),
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	Router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

}



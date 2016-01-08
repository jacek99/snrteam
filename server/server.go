package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jacek99/snrteam/common"
	"net/http"
	"os"
	"time"
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
		Addr:           fmt.Sprintf(":%s", port),
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

}

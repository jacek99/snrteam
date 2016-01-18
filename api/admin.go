package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/jacek99/snrteam/database"
)

type HealthCheck struct {
	Message string
}

// TODO
func adminHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, []HealthCheck{})
}

func truncate(c *gin.Context) {
	database.Truncate()
	c.JSON(http.StatusOK,nil)
}


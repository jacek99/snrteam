package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthCheck struct {
	Message string
}

// TODO
func adminHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, []HealthCheck{})
}


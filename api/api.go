package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/jacek99/snrteam/database"
	"log"
)

func InitRouter(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	router.GET("/api/users", func(c *gin.Context) {

		users, err := database.GetAllUsers()
		if err != nil {
			log.Println(err)
			c.Error(err)
		} else {
			c.JSON(http.StatusOK, users)
		}

	})
}

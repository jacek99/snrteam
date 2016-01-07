package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jacek99/snrteam/database"
	"log"
	"net/http"
)

func getAllUsers(c *gin.Context) {
	users, err := database.GetAllUsers()
	if err != nil {
		log.Println(err)
		c.Error(err)
	} else {
		c.JSON(http.StatusOK, users)
	}
}
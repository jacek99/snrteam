package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/i18n"
	"github.com/jacek99/snrteam/database"
	"log"
	"net/http"
)

const USER_ID = "UserId"

func getAllUsers(c *gin.Context) {
	users, err := database.GetAllUsers()
	if err != nil {
		log.Println(err)
		c.Error(err)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func getUser(c *gin.Context) {
	userId := c.Param(USER_ID)
	user, err := database.GetUser(userId)
	if err != nil {
		log.Println(err)
		c.Error(err)
	} else if user == nil {
		T, _ := i18n.Tfunc(getRequestLanguage(c))
		params := map[string]interface{}{USER_ID:userId}
		c.JSON(http.StatusNotFound, RestError{T("user_not_found",params),T("user"),USER_ID, userId})
	} else {
		c.JSON(http.StatusOK, user)
	}
}





package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/i18n"
	"github.com/jacek99/snrteam/database"
	"log"
	"net/http"
	"github.com/jacek99/snrteam/common"
)

const USER_NAME = "UserName"


func getAllUsers(c *gin.Context) {
	if users, err := database.GetAllUsers(); err != nil {
		log.Println(err)
		c.Error(err)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func getUser(c *gin.Context) {
	userName := c.Param(USER_NAME)
	if user, err := database.GetUserByName(userName); err != nil {
		if err == common.RECORD_NOT_FOUND_ERROR {
			T, _ := i18n.Tfunc(getRequestLanguage(c))
			params := map[string]interface{}{USER_NAME:userName}
			c.JSON(http.StatusNotFound, RestError{T("user_not_found",params),T("user"),USER_NAME, userName})
		} else {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, RestSystemError{"System error occurred."})
		}
	} else {
		c.JSON(http.StatusOK, user)
	}
}





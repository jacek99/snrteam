package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jacek99/snrteam/database"
	"net/http"
	"github.com/jacek99/snrteam/model"
)

func getAllUsers(c *gin.Context) {
	if users, err := database.GetAllUsers(); err != nil {
		handleError(c, err)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func getUser(c *gin.Context) {
	T := getI18n(c)
	if entity, err := database.GetUserByName(c.Param(database.USER_NAME),T); err != nil {
		handleError(c, err)
	} else {
		c.JSON(http.StatusOK, entity)
	}
}

func saveUser(c *gin.Context) {
	var user model.User
	T := getI18n(c)

	if err := c.BindJSON(&user); err == nil {
		if err = database.SaveUser(&user,T); err == nil {
			c.JSON(http.StatusCreated,nil)
		} else {
			handleError(c,err)
		}
	} else {
		handleError(c, err)
	}

}





package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jacek99/snrteam/database"
	"net/http"
	"github.com/jacek99/snrteam/model"
)

// special structs for user to allow specifying / hiding password field
type InputUser struct {
	model.User
	Password     string
}

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
	var input InputUser
	T := getI18n(c)

	if err := c.Bind(&input); err == nil {
		if err = database.SaveUser(&input.User,input.Password, T); err == nil {
			c.JSON(http.StatusCreated, input)
		} else {
			handleError(c,err)
		}
	} else {
		handleError(c, err)
	}

}





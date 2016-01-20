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
		handleError(c, "DB get all users",err)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func getUser(c *gin.Context) {
	T := getI18n(c)
	if entity, err := database.GetUserByName(c.Param(database.USER_NAME),T); err != nil {
		handleError(c, "DB get single user",err)
	} else {
		c.JSON(http.StatusOK, entity)
	}
}

func saveUser(c *gin.Context) {
	var input InputUser
	T := getI18n(c)

	if err := c.BindJSON(&input); err == nil {

//		t, _ := json.Marshal(input)
//		log.Printf("User: %s", t)

		if err = database.SaveUser(&input.User,input.Password, T); err == nil {
			if user, err2 := database.GetUserByName(input.UserName,T); err2 != nil {
				handleError(c,"DB retrieve of saved user",err2)
			} else {
				c.JSON(http.StatusCreated, user)
			}
		} else {
			handleError(c,"DB save of user",err)
		}
	} else {
		handleError(c, "JSON binding to user",err)
	}

}





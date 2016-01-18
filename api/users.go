package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jacek99/snrteam/database"
	"net/http"
	"github.com/jacek99/snrteam/model"
)

// special structs for user to allow specifying / hiding password field
type InputUser struct {
	UserId       int64  `thrift:"user_id,1,required"`
	UserName     string `thrift:"user_name,2,required"`
	EmailAddress string `thrift:"email_address,3,required"`
	FirstName    string `thrift:"first_name,4,required"`
	LastName     string `thrift:"last_name,5,required"`
	BirthDate    *model.Date  `thrift:"birth_date,6,required"`
	CreationDate *model.Date  `thrift:"creation_date,7,required"`
	Photo        []byte `thrift:"photo,8"`
	Password     string `thrift:"pwd_hash,9"`
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
	var input model.User
	T := getI18n(c)

	if err := c.Bind(&input); err == nil {
		if err = database.SaveUser(&input,T); err == nil {
			c.JSON(http.StatusCreated, input)
		} else {
			handleError(c,err)
		}
	} else {
		handleError(c, err)
	}

}





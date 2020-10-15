package users

import (
	"github.com/dypiconpaba/bookstore_users-api/bookstore_users-api/domain/users"
	"github.com/dypiconpaba/bookstore_users-api/bookstore_users-api/services"
	"github.com/dypiconpaba/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
userId, userError := strconv.ParseInt(c.Param("user_id"), 10, 64)

if userError != nil{
	err:= errors.NewBadRequestError("invalid user id should be a number")
	c.JSON(err.Status, err)
	return
}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")

		c.JSON(restErr.Status, restErr)
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

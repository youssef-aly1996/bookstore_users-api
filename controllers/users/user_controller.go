package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/youssef-aly1996/bookstore_users-api/domain/userdomain"
	"github.com/youssef-aly1996/bookstore_users-api/services"
	"github.com/youssef-aly1996/bookstore_users-api/utils/errors"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		err := errors.NewBadRequest("user id must be a positive number")
		c.JSON(err.StatusCode, err)
	}
	user, restError := services.GetUserById(userId)
	if restError != nil {
		c.JSON(restError.StatusCode, restError)
		return
	}
	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {
	var user userdomain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.RestError{
			Message:    "Not a Valid json",
			StatusCode: http.StatusBadRequest,
			Error:      "bad request",
		}
		c.JSON(restError.StatusCode, restError)
		return
	}
	createdUser, createError := services.CreateUser(user)
	if createError != nil {
		c.JSON(createError.StatusCode, createError)
		return
	}
	c.JSON(http.StatusCreated, createdUser)
}

func GetAllusers(c *gin.Context) {
	var user userdomain.User
	allUsers, err := services.GetAllusers(user)
	if err != nil {
		c.JSON(400, "bad requset")
		return
	}
	for i := range allUsers {
		c.JSON(200, &i)
	}
}

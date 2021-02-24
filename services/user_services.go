package services

import (
	"github.com/youssef-aly1996/bookstore_users-api/domain/userdomain"
	"github.com/youssef-aly1996/bookstore_users-api/utils/errors"
)

func CreateUser(user userdomain.User) (*userdomain.User, *errors.RestError) {

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetAllusers(user userdomain.User) ([]*userdomain.User, *errors.RestError) {
	allUsers := user.GetAll()
	return allUsers, nil
}

func GetUserById(id int64) (*userdomain.User, *errors.RestError) {
	fetchedUser := userdomain.User{Id: id}
	if err := fetchedUser.Get(); err != nil {
		return nil, err
	}
	return &fetchedUser, nil
}

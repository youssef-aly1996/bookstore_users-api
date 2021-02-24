package userdomain

import (
	"fmt"

	"github.com/youssef-aly1996/bookstore_users-api/utils/errors"
)

var (
	userDb = make(map[int64]*User)
)

func (user *User) Get() *errors.RestError {
	fetchedUser := userDb[user.Id]
	if fetchedUser == nil {
		restError := errors.NotFoundRequest(fmt.Sprintf("user with id %d not found", user.Id))
		return restError
	}
	user.Id = fetchedUser.Id
	user.FirstName = fetchedUser.FirstName
	user.LastName = fetchedUser.LastName
	user.Email = fetchedUser.Email
	user.CreatedAt = fetchedUser.CreatedAt
	return nil
}

func (user User) Save() *errors.RestError {
	userToBeSaved := userDb[user.Id]
	if userToBeSaved != nil && userToBeSaved.Email == user.Email {
		err := errors.NewBadRequest(fmt.Sprintf("user id %d or email %s already exist", userToBeSaved.Id, userToBeSaved.Email))
		return err
	}
	userDb[user.Id] = &user
	fmt.Println(userDb)
	return nil
}

func (user User) GetAll() []*User {
	var userSlice []*User
	for _, v := range userDb {
		userSlice = append(userSlice, v)
	}
	fmt.Println(&userSlice)
	return userSlice
}

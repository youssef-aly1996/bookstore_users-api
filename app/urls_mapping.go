package app

import "github.com/youssef-aly1996/bookstore_users-api/controllers/users"

func mapUrl() {
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users", users.GetAllusers)
	router.POST("/users", users.CreateUser)
}

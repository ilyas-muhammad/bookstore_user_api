package app

import (
	"petaniweb.com/rest/v1/bookstore_user_api/controller"
	"petaniweb.com/rest/v1/bookstore_user_api/controller/users"
)

func mapUrls() {
	router.GET("/ping", controller.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
}

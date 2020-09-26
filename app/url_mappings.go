package app

import (
	"github.com/ilyas-muhammad/bookstore_user_api/controller"
	"github.com/ilyas-muhammad/bookstore_user_api/controller/users"
)

func mapUrls() {
	router.GET("/ping", controller.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.PUT("/users/:user_id", users.UpdateUser)
}

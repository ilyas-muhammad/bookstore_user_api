package app

import (
	"github.com/ilyas-muhammad/bookstore_user_api/controller"
	"github.com/ilyas-muhammad/bookstore_user_api/controller/users"
)

func mapUrls() {
	router.GET("/ping", controller.Ping)

	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("internal/users/search", users.Search)
}

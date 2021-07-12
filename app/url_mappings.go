package app

import (
	"github.com/gopal/bookstore_user-api/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}

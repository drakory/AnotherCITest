package main

import (
	"bookapi/config"
	"bookapi/controller"
	"github.com/gin-gonic/gin"
)

func Addition(a int , b int) int {
	c := a + b
	return c
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		book := v1.Group("/book")
		{
			book.POST("/", controller.InsertBook)
			book.GET("/", controller.GetAllBooks)
			book.GET("/:id", controller.GetBook)
			book.PUT("/:id", controller.UpdateBook)
			book.DELETE("/:id", controller.DeleteBook)
		}

		user := v1.Group("/user")
		{
			user.GET("/", controller.GetAllUsers)
			user.POST("/", controller.Register)
			user.GET("/:id", controller.Profile)
			user.PUT("/:id", controller.UpdateProfile)
			user.DELETE("/:id", controller.DeleteAccount)
		}

	}
	return router
} 

func main() {
	config.ConnectDB()

	defer config.CloseDb()
	router := setupRouter()
	err := router.Run(":3000")
	if err != nil {
		panic("Error router.run")
	}
}

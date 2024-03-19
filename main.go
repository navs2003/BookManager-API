package main

import (
	"crudApi/controllers"
	"crudApi/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	r := gin.Default()
	r.POST("/books", controllers.BookCreate)
	r.Run() // listen and serve on 0.0.0.0:8080
}

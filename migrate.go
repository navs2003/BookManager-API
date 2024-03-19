package main

import (
	"crudApi/initializers"
	"crudApi/models"
)

func init() {
	initializers.DatabaseConnection()
	initializers.LoadEnvVariables()
}

func main() {
	initializers.DB.AutoMigrate(&models.Book{})
}

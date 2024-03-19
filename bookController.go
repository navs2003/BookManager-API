package controllers

import (
	"crudApi/initializers"
	"crudApi/models"

	"github.com/gin-gonic/gin"
)

func BookCreate(c *gin.Context) {
	// get data

	//create book
	book := models.Book{Title: "Harry Potter", Content: "Wizard"}
	result := initializers.DB.Create(&book)
	//return it
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

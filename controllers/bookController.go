package controllers

import (
	"fmt"
	"strconv"

	"github.com/brenogmrs/go-gin-api/database"
	"github.com/brenogmrs/go-gin-api/models"
	"github.com/gin-gonic/gin"
)

func ShowBook(c *gin.Context) {

	id := c.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be an integer",
		})
		return
	}

	db := database.GetDatabase()

	var foundBook models.Book

	err = db.First(&foundBook, newId).Error

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(200, foundBook)

}

func CreateBook(c *gin.Context) {

	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Something went wrong" + err.Error(),
		})

		return
	}

	fmt.Println(book)

	err = db.Create(&book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot Create book" + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func ShowBooks(c *gin.Context) {

	db := database.GetDatabase()

	var books []models.Book

	err := db.Find(&books).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot List book" + err.Error(),
		})

		return
	}

	c.JSON(200, books)

}

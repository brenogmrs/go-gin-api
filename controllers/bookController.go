package controllers

import (
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
			"error": "Book not found" + err.Error(),
		})

		return
	}

	c.JSON(200, foundBook)

}

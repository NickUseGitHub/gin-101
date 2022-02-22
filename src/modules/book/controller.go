package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func FindBook(c *gin.Context) { // Get model if exist
	var book Book

	db, ok := c.Keys["DB"].(*gorm.DB)
	if !ok {
		panic("DB is not called on *gin.Context")
	}

	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func FindBooks(c *gin.Context) {
	var books []Book

	db, ok := c.Keys["DB"].(*gorm.DB)
	if !ok {
		panic("DB is not called on *gin.Context")
	}

	db.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(c *gin.Context) {
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, ok := c.Keys["DB"].(*gorm.DB)
	if !ok {
		panic("DB is not called on *gin.Context")
	}

	// Create book
	book := Book{Title: input.Title, Author: input.Author}
	db.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {
	// Get model if exist
	var book Book

	db, ok := c.Keys["DB"].(*gorm.DB)
	if !ok {
		panic("DB is not called on *gin.Context")
	}

	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	// Get model if exist
	var book Book

	db, ok := c.Keys["DB"].(*gorm.DB)
	if !ok {
		panic("DB is not called on *gin.Context")
	}

	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

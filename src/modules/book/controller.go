package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func FindBooks(c *gin.Context) {
	var books []Book

	db, ok := c.Keys["DB"].(*gorm.DB)
	if !ok {
		panic("DB is not called on *gin.Context")
	}

	db.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

package main

import (
	"github.com/NickUseGitHub/gin-101/src/modules"
	"github.com/NickUseGitHub/gin-101/src/modules/book"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func useDatabase(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(useDatabase(modules.ConnectDatabase()))

	book.RegisterController(r)

	r.Run()
}

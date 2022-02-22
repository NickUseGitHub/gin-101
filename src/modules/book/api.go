package book

import "github.com/gin-gonic/gin"

func RegisterController(r *gin.Engine) {
	r.GET("/books", FindBooks)
	r.GET("/books/:id", FindBook)
	r.POST("/books", CreateBook)
	r.PATCH("/books/:id", UpdateBook)
	r.DELETE("/books/:id", DeleteBook)
}

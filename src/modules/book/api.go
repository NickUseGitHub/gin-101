package book

import "github.com/gin-gonic/gin"

func RegisterController(r *gin.Engine) {
	r.GET("/books", FindBooks)
}

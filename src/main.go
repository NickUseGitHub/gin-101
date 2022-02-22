package main

import (
	"github.com/NickUseGitHub/gin-101/src/modules"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	modules.ConnectDatabase()

	r.Run()
}

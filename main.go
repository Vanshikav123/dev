package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//go mod init <name>
//github for go/gin https://github.com/gin-gonic/gin
//what are frameworks
//an essential supporting structure of a building applications
//types of frameworks in golang
//Gin/Gin-Gonic
//Beego
//Echo
//go mod tidy to install all the packages needed in the go.mod

func main() {
	r := gin.Default()
	r.GET("/ping", test)
	r.Run(":9090")
}
func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//api= application programming interface

type api struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var data api

func main() {
	r := gin.Default()
	r.GET("/get", getValues) //path and controller it containes all the execution parameters
	r.POST("/post", postValues)
	r.PUT("/put", putValues)
	r.DELETE("/delete", deleteValues)
	r.Run(":9092")
}

/*func getValues(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}*/

func postValues(c *gin.Context) {
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "something wrong",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func putValues(c *gin.Context) {
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "something wrong",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func deleteValues(c *gin.Context) {
	data = api{}
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

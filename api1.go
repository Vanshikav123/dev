package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/get", getValues)
	r.Run(":9093")
}

var url = "http://date.jsontest.com/"

func getValues(c *gin.Context) {
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}
	//close response body
	defer resp.Body.Close()

	//read all response body
	data, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}

	var target map[string]interface{}
	err = json.Unmarshal(data, &target)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": target,
	})

}

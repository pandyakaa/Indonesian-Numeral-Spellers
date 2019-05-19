package main

import (
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func read(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var query = c.Query("text")
	var answer = readUtils(query)
	c.JSON(200, gin.H{
		"number": answer,
	})
}

func spell(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var query = c.Query("number")
	var num, err = strconv.Atoi(query)
	var result = spellUtils(num)
	if err != nil {
		c.JSON(200, gin.H{
			"text": "error",
		})
	} else {
		c.JSON(200, gin.H{
			"text": result,
		})
	}
}

func main() {
	var router = gin.Default()
	var config = cors.DefaultConfig()
	config.AddAllowHeaders("*")
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET"}
	router.Use(cors.New(config))
	router.GET("/spell", spell)
	router.POST("/read", read)
	router.Run()
}
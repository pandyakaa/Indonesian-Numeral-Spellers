package main

import (
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func read(c *gin.Context) {
	c.Header("Content-Type", "application/x-www-form-urlencoded")
	var query = c.PostForm("text")
	if isValidText(query) {
		var answer = readUtils(query)
		if answer <= 2000000000 {
			c.JSON(200, gin.H{
				"number": answer,
			})
		} else {
			c.JSON(200, gin.H{
				"number": "error",
			})
		}
	} else {
		c.JSON(200, gin.H{
			"number": "error",
		})
	}
}

func spell(c *gin.Context) {
	var query = c.Query("number")
	var num, err = strconv.Atoi(query)
	var result = ""
	if num > 0 {
		result = spellUtils(num)
	} else {
		result = "negatif " + spellUtils(num*-1)
	}
	if err != nil {
		c.JSON(200, gin.H{
			"text": "error",
		})
	} else {
		if isValidNumber(num) {
			c.JSON(200, gin.H{
				"text": result,
			})
		} else {
			c.JSON(200, gin.H{
				"text": "error",
			})
		}
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
	router.Run(":47000")
}

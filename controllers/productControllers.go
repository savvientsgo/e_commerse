package controllers

import (
	"github.com/gin-gonic/gin"
)

func probuct(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "200",
	})
}

// func main() {
// 	fmt.Println("Hello World")

// 	r := gin.Default()
// 	r.GET("/", probuct)

// 	r.Run()
// }

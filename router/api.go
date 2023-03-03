package router

import (
	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "[]",
	})
}

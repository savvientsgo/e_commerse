package main

import (
	"E_COMMERSE/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", router.User)

	r.Run("localhost:8080")
}

package main

import (
	"fmt"
	"interview-app/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	lib.DBConn = lib.ConnectionDB()
	fmt.Println("DB Connected")
	route := gin.Default()
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	route.Run()
}

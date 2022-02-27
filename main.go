package main

import (
	"interview-app/lib"

	"github.com/gin-gonic/gin"
)

func main() {
    lib.DBConn = lib.ConnectionDB()
	route := gin.Default()
    route.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    route.Run()
}

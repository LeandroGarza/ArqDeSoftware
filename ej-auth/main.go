package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	engine.POST("/login")
	engine.Run(":8080")
}

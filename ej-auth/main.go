package main

import (
	"github.com/gin-gonic/gin"
	"https://github.com/emikohmann/arq-software/ej-auth/domain"
)

func main() {
	engine := gin.New()
	engine.POST("/login")
	engine.Run(":8080")
}

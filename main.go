package main

import (
	"github.com/dawnnguyen-vn/learn-go/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.String(200, "ok")
	})

	r.GET("/", func(c *gin.Context) {
		c.String(200, utils.GetValue())
	})

	r.Run()
}

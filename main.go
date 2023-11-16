package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dawnnguyen-vn/learn-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if APP_ENV := os.Getenv("APP_ENV"); APP_ENV == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	if APP_ENV := os.Getenv("APP_ENV"); APP_ENV != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()

	r.GET("/health", func(c *gin.Context) {
		c.String(200, "okk")
	})

	r.GET("/", func(c *gin.Context) {
		c.String(200, utils.GetValue())
	})

	if APP_ENV := os.Getenv("APP_ENV"); APP_ENV != "dev" {
		fmt.Print("Listening and serving HTTP on :" + os.Getenv("PORT"))
	}
	r.Run()
}

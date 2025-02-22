package main

import (
	"example.com/short_link/db"
	"example.com/short_link/handlers"
	"example.com/short_link/utils"
	"github.com/gin-gonic/gin"
)

func init() {
	db.ConnectDB()
	utils.LoadEnvs()
}
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	router.GET("/", handlers.HomePageController)
	router.POST("/shorten-url", handlers.ShortenController)
	router.GET("/short/:url", handlers.GetUrl)
	router.Run(":8080")
}

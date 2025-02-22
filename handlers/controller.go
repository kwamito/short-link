package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomePageController(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"name": "Kwame",
	})
}

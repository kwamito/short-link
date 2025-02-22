package handlers

import (
	"html/template"
	"net/http"

	"example.com/short_link/utils"
	"github.com/gin-gonic/gin"
)

func ShortenController(c *gin.Context) {
	url := c.PostForm("url")
	shortenedUrl, err := utils.ShortenURL(url)
	if err != nil {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"error": "eror",
		})
	}
	c.HTML(http.StatusOK, "home.html", gin.H{
		"originalURL": url,
		// "shortendURL": "short/" + shortenedUrl,
		"shortenedURL": template.URL(shortenedUrl),
	})
}

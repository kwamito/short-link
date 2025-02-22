package handlers

import (
	"fmt"
	"net/http"

	"example.com/short_link/db"
	"example.com/short_link/models"
	"github.com/gin-gonic/gin"
)

func GetUrl(c *gin.Context) {
	var foundLink models.ShortURL
	shortUrl := c.Param("url")
	fmt.Println(shortUrl, "short uuid here")
	db.DB.Where("url_id=?", shortUrl).Find(&foundLink)
	if foundLink.ID == 0 {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"error": "error",
		})
	}
	fmt.Println("redirecting link")
	fmt.Println(foundLink.OriginalUrl, "This is the original link")
	c.Redirect(http.StatusFound, foundLink.OriginalUrl)

}

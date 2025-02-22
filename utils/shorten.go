package utils

import (
	"fmt"
	"os"

	"example.com/short_link/db"
	"example.com/short_link/models"
	"github.com/lithammer/shortuuid/v4"
)

func ShortenURL(url string) (string, error) {
	// Shorten URL
	// db.DB.Create(&models.ShortURL{OriginalUrl: url})

	var foundLink models.ShortURL
	db.ConnectDB()
	fmt.Println(db.DB, "database value, this time again")
	db.DB.Where("original_url=?", url).Find(&foundLink)

	if foundLink.ID == 0 {
		shortUuid := shortuuid.New()
		generatedShortUrl := "http://" + os.Getenv("BASE_URL") + "/" + "short" + "/" + shortUuid
		url := models.ShortURL{OriginalUrl: url, ShortUrl: generatedShortUrl, UrlId: shortUuid}
		db.DB.Create(&url)
		return generatedShortUrl, nil
	} else {
		return foundLink.ShortUrl, nil
		// return "", errors.New("no ")
	}

}

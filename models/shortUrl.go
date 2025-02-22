package models

type ShortURL struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	OriginalUrl string `json:"original_url" gorm:"unique"`
	ShortUrl    string `json:"short_url" gorm:"unique"`
	UrlId       string `json:"url_id" gorm:"unique"`
}

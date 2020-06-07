package models

import (
	"fmt"
	"url-shortener/utils"

	"github.com/astaxie/beego/orm"
)

// ShortenedURL represents a shortened URL
type ShortenedURL struct {
	Slug        string `orm:"pk" json:"slug"`
	OriginalURL string `orm:"column(original_url)" json:"originalUrl"`
}

// TableName redefines the ShortenedURL table to "shortened_url"
// The default name is "shortened_u_r_l"
func (u *ShortenedURL) TableName() string {
	return "shortened_url"
}

var maxTries = 3

// CreateShortenedURL creates and persists a shortened URL record with a unique slug
func CreateShortenedURL(url string) (ShortenedURL, error) {
	o := orm.NewOrm()
	err := o.Begin()

	if err != nil {
		o.Rollback()
		return ShortenedURL{}, err
	}

	for i := 0; i < maxTries; i++ {
		slug := utils.CreateSlug(6)
		shortenedURL := ShortenedURL{Slug: slug}
		err := o.Read(&shortenedURL)

		if err == orm.ErrNoRows {
			newURL := ShortenedURL{Slug: slug, OriginalURL: url}
			_, insertErr := o.Insert(&newURL)

			if insertErr == nil {
				commitErr := o.Commit()

				if commitErr == nil {
					return newURL, nil
				}
			}

			// TODO: refactor this to return the actual error
			break
		}
	}

	o.Rollback()
	return ShortenedURL{}, fmt.Errorf("Could not create unique shortened URL")
}

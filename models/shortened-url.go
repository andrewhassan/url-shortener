package models

// ShortenedURL represents a shortened URL
type ShortenedURL struct {
	Slug        string `orm:"pk"`
	OriginalURL string `orm:"column(original_url)"`
}

// TableName redefines the ShortenedURL table to "shortened_url"
// The default name is "shortened_u_r_l"
func (u *ShortenedURL) TableName() string {
	return "shortened_url"
}

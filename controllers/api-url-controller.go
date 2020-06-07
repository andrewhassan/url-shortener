package controllers

import (
	"fmt"
	"net/url"
	"url-shortener/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// APIURLController is an API controller for URLs
type APIURLController struct {
	beego.Controller
}

// Get handles GET /api/url/:slug
func (c *APIURLController) Get() {
	slug := c.Ctx.Input.Param(":slug")

	o := orm.NewOrm()
	shortenedURL := models.ShortenedURL{Slug: slug}
	err := o.Read(&shortenedURL)

	if err == orm.ErrNoRows {
		c.Abort("404")
	} else if err != nil {
		c.Abort("500")
	}

	c.Data["json"] = &shortenedURL
	c.ServeJSON()
}

// Post handles POST /api/url
func (c *APIURLController) Post() {
	parsed, err := url.ParseRequestURI(c.GetString("url"))

	if err != nil {
		c.Abort("400")
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		c.Abort("400")
	}

	shortenedURL, err := models.CreateShortenedURL(parsed.String())

	if err != nil {
		c.Abort("500")
	}

	fmt.Println(shortenedURL)

	c.Data["json"] = &shortenedURL
	c.ServeJSON()
}

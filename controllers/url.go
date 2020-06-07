package controllers

import (
	"fmt"
	"net/url"
	"url-shortener/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// URLController handles redirecting
type URLController struct {
	beego.Controller
}

// Get handles GET /:id
func (c *URLController) Get() {
	slug := c.Ctx.Input.Param(":slug")

	o := orm.NewOrm()
	shortenedURL := models.ShortenedURL{Slug: slug}
	err := o.Read(&shortenedURL)

	fmt.Println(err)

	if err == orm.ErrNoRows {
		// TODO: Figure out why it can't find the template
		c.TplName = "NotFound"
		return
	} else if err != nil {
		// TODO: Figure out why it can't find the template
		c.TplName = "InternalServerError"
		return
	}

	c.Redirect(shortenedURL.OriginalURL, 302)
}

// APIURLController is an API controller for URLs
type APIURLController struct {
	beego.Controller
}

// Get handles GET /api/url/:slug
func (c *APIURLController) Get() {
	c.Data["JSON"] = "{\"test\": 1234}"
	c.ServeJSON()
}

// Post handles POST /api/url
func (c *APIURLController) Post() {
	parsed, err := url.ParseRequestURI(c.GetString("url"))

	if err != nil {
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

package controllers

import (
	"fmt"
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

	if err != nil {
		fmt.Println("Failed to read ShortenedURL:")
		fmt.Println(err)
		// TODO: Figure out why it can't find the template
		c.TplName = "NotFound"
		return
	}

	c.Redirect(shortenedURL.OriginalURL, 302)
}

// APIURLController is an API controller for URLs
type APIURLController struct {
	beego.Controller
}

// Get handles GET /api/url
func (c *APIURLController) Get() {
	c.Data["JSON"] = "{\"test\": 1234}"
	c.ServeJSON()
}

// Post handles POST /api/url
func (c *APIURLController) Post() {
	c.Data["JSON"] = "{\"test\": 1234}"
	c.ServeJSON()
}

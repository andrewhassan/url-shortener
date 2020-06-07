package controllers

import (
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

package controllers

import (
	"github.com/astaxie/beego"
)

// URLController handles redirecting
type URLController struct {
	beego.Controller
}

// Get handles GET /:id
func (c *URLController) Get() {
	id := c.Ctx.Input.Param(":id")
	c.Data["json"] = map[string]interface{}{"id": id}
	c.ServeJSON()
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

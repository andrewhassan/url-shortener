package controllers

import (
	"github.com/astaxie/beego"
)

// MainController is the root beego controller for this app
type MainController struct {
	beego.Controller
}

// Get is the GET request handler for MainController
func (c *MainController) Get() {
	c.TplName = "index.tpl"
}

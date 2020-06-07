package routers

import (
	"url-shortener/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/:id", &controllers.URLController{})

	api := beego.NewNamespace("/api",
		beego.NSNamespace("/url",
			beego.NSRouter("/:id", &controllers.APIURLController{}),
		),
	)

	beego.AddNamespace(api)
}

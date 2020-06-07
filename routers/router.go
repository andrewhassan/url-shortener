package routers

import (
	"url-shortener/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/:slug", &controllers.URLController{})

	api := beego.NewNamespace("/api",
		beego.NSNamespace("/url",
			beego.NSRouter("/", &controllers.APIURLController{}, "post:Post"),
			beego.NSRouter("/:slug", &controllers.APIURLController{}, "get:Get"),
		),
	)

	beego.AddNamespace(api)
}

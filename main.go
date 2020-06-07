package main

import (
	"url-shortener/models"
	_ "url-shortener/routers"

	"github.com/astaxie/beego"
)

func main() {
	models.SyncDatabase()
	beego.Run()
}

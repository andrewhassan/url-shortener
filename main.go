package main

import (
	_ "url-shortener/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

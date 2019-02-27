package main

import (
	_ "fullStack/routers"
	"github.com/astaxie/beego"
)

func main() {

	//StaticDir["/static"] = "static"
	beego.SetStaticPath("/down", "download") 
	beego.Run()
}


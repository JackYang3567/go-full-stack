package controllers

import (
	//"fullStack/models"
	//"encoding/json"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {	
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	
	v := c.GetSession("asta")
    if v == nil {
        c.SetSession("asta", int(1))
        c.Data["num"] = 0
    } else {
        c.SetSession("asta", v.(int)+1)
        c.Data["num"] = v.(int)
    }

	c.Layout = "admin/layout.html" //使用布局
	c.TplName = "index.tpl"        //渲染到视图

	//c.Ctx.WriteString("hello")  //输出字符串

    // log := models.LogPV()
	 //c.Data["json"] = log //c.Data
	//c.ServeJSON() //输出JSON字符

	
}



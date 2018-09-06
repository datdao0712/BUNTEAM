package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Profile() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	main.TplName = "viet561995/viet"
}

func (main *MainController) HelloSitepoint() {
    main.Data["ten"] = "Ten: Hoang"
    main.Data["team"] = "Bun Team"
    main.Data["dob"] = "ngay sinh:02/10/1995"
    main.TplName = "hoang0210/anhbun.tpl"
}

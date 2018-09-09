package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
type SangController struct {
	beego.Controller
}
type HoangController struct {
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
	c.TplName = "viet561995/viet.tpl"
}

func (main *MainController) HelloSitepoint() {
	main.Data["ten"] = "Ten: Hoang"
	main.Data["team"] = "Bun Team"
	main.Data["dob"] = "ngay sinh:02/10/1995"
	main.TplName = "hoang0210/anhbun.tpl"
}
func (sa *SangController) Get() {
	sa.TplName = "sang/sang.html"
}
func (login	 *HoangController) Login() {
	login.TplName = "hoang0210/login.html"
}

func (show *HoangController) Show(){
	user := show.GetString("username")
	pass := show.GetString("pass")
	show.Data["username"] = user
	show.Data["pass"] = pass
	show.TplName = "hoang0210/show.html"
} 


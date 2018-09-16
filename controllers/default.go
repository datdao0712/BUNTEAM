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
type Sang2Controller struct {
	beego.Controller
}

type HoangController struct {
	beego.Controller
}

type VietController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (Study *VietController) Study() {
	Study.TplName = "viet/hoc.tpl"
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

func (sub *SangController) Post() {
	subfirstname := sub.GetString("firstname")
	sublastname := sub.GetString("lastname")
	sub.Data["firstname"] = subfirstname
	sub.Data["lastname"] = sublastname
	sub.TplName = "sang/sang2.html"
}

func (login *HoangController) Login() {
	login.TplName = "hoang0210/login.html"
}

func (show *HoangController) Show() {
	user := show.GetString("username")
	pass := show.GetString("pass")
	show.Data["username"] = user
	show.Data["pass"] = pass
	show.TplName = "hoang0210/show.html"
}

func (signup *HoangController) Signup() {
	signup.TplName = "hoang0210/signup.html"
}

func (signup *VietController) Signup() {
	signup.TplName = "viet/signup.html"
}
func (signup *VietController) Submit() {
	firstname := signup.GetString("firstname")
	lastname := signup.GetString("lastname")
	signup.Data["firstname"] = firstname
	signup.Data["lastname"] = lastname
	signup.TplName = "viet/ketqua.html"
}

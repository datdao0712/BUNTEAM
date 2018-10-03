package controllers

import (
	"github.com/astaxie/beego"
)

type HoangController struct {
	beego.Controller
}

func (hoang *HoangController) Hoang() {
	hoang.TplName = "hoang/index.tpl"
}

func (home *HoangController) Home() {
	home.TplName = "hoang/home.html"
}

func (homepage1 *HoangController) Homepage1() {
	homepage1.TplName = "hoang/homepage1.html"
}

package controllers

import (
	"github.com/astaxie/beego"
)

type VietController struct {
	beego.Controller
}

func (dangnhap *VietController) Get() {
	dangnhap.TplName = "viet/trangchu.html"
}

func (dangnhap *VietController) Home() {
	dangnhap.TplName = "viet/home.html"
}

func (dangnhap *VietController) About() {
	dangnhap.TplName = "viet/about.html"
}

func (dangnhap *VietController) Login() {
	dangnhap.TplName = "viet/login.html"
}

package controllers

import (
	"github.com/astaxie/beego"
)

type SangController struct {
	beego.Controller
}

func (s *SangController) Get() {
	s.Data["Website"] = "beego.me"
	s.Data["Email"] = "astaxie@gmail.com"
	s.TplName = "sang/sang.html"
}
func (s *SangController) Home() {
	s.TplName = "sang/home.html"
}
func (s *SangController) Contact() {
	s.TplName = "sang/contact.html"
}

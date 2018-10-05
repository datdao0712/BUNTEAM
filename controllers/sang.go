package controllers

import (
	"github.com/astaxie/beego"
)

type SangController struct {
	beego.Controller
}

func (s *SangController) Get() {
	s.TplName = "sang/homework.html"
}
func (s *SangController) Home() {
	s.TplName = "sang/home.html"
}
func (s *SangController) Contact() {
	s.TplName = "sang/contact.html"
}

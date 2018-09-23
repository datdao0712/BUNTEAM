package routers

import (
	"github.com/astaxie/beego"
	"github.com/datdao0712/BUNTEAM/controllers"
)

func init() {
	beego.Router("/sang", &controllers.SangController{})
	beego.Router("/sang/home", &controllers.SangController{}, "get:Home")
	beego.Router("/sang/contact", &controllers.SangController{}, "get:Contact")
}

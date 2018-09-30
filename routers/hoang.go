package routers

import (
	"github.com/astaxie/beego"
	"github.com/datdao0712/BUNTEAM/controllers"
)

func init() {
	beego.Router("/hoang", &controllers.HoangController{}, "get:Hoang")
	beego.Router("/hoang/home", &controllers.HoangController{}, "get:Home")
}

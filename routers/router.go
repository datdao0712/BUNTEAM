package routers

import (
	"github.com/astaxie/beego"
	"github.com/datdao0712/BUNTEAM/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hoang", &controllers.MainController{}, "get:HelloSitepoint")
	beego.Router("/viet", &controllers.MainController{}, "get:Profile")

	//hoang
	beego.Router("/hoang/login", &controllers.HoangController{}, "get:Login")
	beego.Router("/hoang/show", &controllers.HoangController{}, "post:Show")
	beego.Router("/signup", &controllers.HoangController{}, "get:Signup")
	beego.Router("/signup/privacy", &controllers.HoangController{}, "get:Privacy")
}

package routers

import (
	"github.com/astaxie/beego"
	"github.com/datdao0712/BUNTEAM/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hoang", &controllers.MainController{}, "get:HelloSitepoint")

	beego.Router("/sang", &controllers.SangController{})
	beego.Router("/sang2", &controllers.Sang2Controller{})
	//hoang
	beego.Router("/hoang/login", &controllers.HoangController{}, "get:Login")
	beego.Router("/hoang/show", &controllers.HoangController{}, "post:Show")
	beego.Router("/signup", &controllers.HoangController{}, "get:Signup")

	//viet
	beego.Router("/viet", &controllers.VietController{}, "get:Study")

	beego.Router("/viet/signup", &controllers.VietController{}, "get:Signup")
	beego.Router("/viet/signup", &controllers.VietController{}, "post:Submit")

}

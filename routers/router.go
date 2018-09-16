package routers

import (
	"github.com/astaxie/beego"
	"github.com/datdao0712/BUNTEAM/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hoang", &controllers.MainController{}, "get:HelloSitepoint")
	beego.Router("/viet", &controllers.MainController{}, "get:Profile")
	beego.Router("/sang", &controllers.SangController{})
<<<<<<< HEAD
	beego.Router("/sang2", &controllers.Sang2Controller{})
=======
	//hoang
	beego.Router("/hoang/login", &controllers.HoangController{}, "get:Login")
	beego.Router("/hoang/show", &controllers.HoangController{}, "post:Show")
	beego.Router("/signup", &controllers.HoangController{}, "get:Signup")
>>>>>>> 37ad8b86eb3ecf0726eb014b354a4775177ef154
}

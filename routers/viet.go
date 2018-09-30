package routers

import (
	"github.com/astaxie/beego"
	"github.com/datdao0712/BUNTEAM/controllers"
)

func init() {
	beego.Router("/viet/trangchu", &controllers.VietController{})
	beego.Router("/viet/home", &controllers.VietController{}, "get:Home")
	beego.Router("/viet/about", &controllers.VietController{}, "get:About")
	beego.Router("/viet/login", &controllers.VietController{}, "get:Login")

}

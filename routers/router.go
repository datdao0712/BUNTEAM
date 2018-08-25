package routers

import (
	"github.com/datdao0712/BUNTEAM/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}

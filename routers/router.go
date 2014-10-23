package routers

import (
	"proxy/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.RESTRouter("/Gateway/InterfaceI", &controllers.FaceIController{})
    beego.RESTRouter("/Gateway/InterfaceII", &controllers.FaceIIController{})
	beego.Router("/proxy/*", &controllers.CommonController{})
}

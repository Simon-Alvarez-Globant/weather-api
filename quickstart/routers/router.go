package routers

import (
	"bapi/quickstart/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/weather", &controllers.WeatherController{}) // <- the controller of this endpoint
}

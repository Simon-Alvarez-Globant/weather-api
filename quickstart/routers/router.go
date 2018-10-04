package routers

import (
	controller "bapi/quickstart/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/weather", &controller.WeatherController{}) // <- the controller of this endpoint
	beego.Router("/scheduler/weather", &controller.SchedulerController{})
}

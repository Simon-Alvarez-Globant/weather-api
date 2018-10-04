// @APIVersion 1.0.0
// @Title Weather API
// @Description This API can get the weather info of more than 200.000 cities!
// @Contact simon.alvarez@globant.com
package routers

import (
	controller "bapi/quickstart/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/weather", &controller.WeatherController{}) // <- the controller of this endpoint
	beego.Router("/scheduler/weather", &controller.SchedulerController{})
}

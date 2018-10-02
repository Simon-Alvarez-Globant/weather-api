package controller

import (
	"bapi/quickstart/lib/http"
	"bapi/quickstart/lib/scheduler"

	"github.com/astaxie/beego"
)

type SchedulerController struct {
	beego.Controller
}

func (c SchedulerController) Put() {
	var param http.HttpParams
	param.City = c.GetString("city")
	param.Country = c.GetString("country")

	status := scheduler.Schedule(param)

	c.Ctx.Output.SetStatus(status)
}

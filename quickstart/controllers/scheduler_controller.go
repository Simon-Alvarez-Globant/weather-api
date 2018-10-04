package controller

import (
	"bapi/quickstart/lib/scheduler"
	"fmt"

	"github.com/astaxie/beego"
)

type SchedulerController struct {
	beego.Controller
}

type Request struct {
	City    string
	Country string
}

// @Title putScheduler
// @Summary putscheduler
// @Description add a new scheduled worker to fetch data each hour
// @Param city query string true
// @Param country query string true
// @Success 200
// @Failure 400 Bad request
// @Failure 404 Not found
// @Accept json
// @router /scheduler/weather [put]
func (c SchedulerController) Put() {
	var param Request
	param.City = c.GetString("city")
	param.Country = c.GetString("country")
	var status int
	err := scheduler.Schedule(param.City, param.Country)
	if err != nil {
		fmt.Println("ERROR:::", err)
		status = 400
	} else {
		fmt.Printf("Scheduler %v, added with success \n", param.City)
		status = 202
	}

	c.Ctx.Output.SetStatus(status)
}

package controller

import (
	"bapi/quickstart/lib"
	"bapi/quickstart/lib/http"

	"github.com/astaxie/beego"
)

type WeatherController struct {
	beego.Controller
}

func (c WeatherController) Get() {
	var param http.HttpParams
	param.City = c.GetString("city")
	param.Country = c.GetString("country")

	json := lib.GetData(param)

	c.Ctx.Output.Header("Content-Type", "application/json")
	c.Ctx.Output.JSON(json, true, true)
}

package controller

import (
	"bapi/quickstart/lib"
	"fmt"

	"github.com/astaxie/beego"
)

type WeatherController struct {
	beego.Controller
}

// @Title getWeather
// @Summary getWeather
// @Description get the weather data from the database or an external API
// @Param   city string, country string
// @Success 200
// @Failure 400 Bad request
// @Failure 404 Not found
// @router /weather?city=$City&country=$Country [get]
func (c WeatherController) Get() {
	city := c.GetString("city")
	country := c.GetString("country")

	json, err := lib.GetData(city, country)
	if err != nil {
		fmt.Println("ERROR:::::::", err)
	}

	c.Ctx.Output.Header("Content-Type", "application/json")
	c.Ctx.Output.JSON(json, true, true)
}

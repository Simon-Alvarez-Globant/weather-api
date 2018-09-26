package controllers

import (
	"bapi/quickstart/utils"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type WeatherController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *WeatherController) Get() {
	city := c.GetString("city")
	country := c.GetString("country")
	id := beego.AppConfig.String("appid")
	urlExt := beego.AppConfig.String("externalapi")

	url := urlExt + city + "," + country + "&appid=" + id

	json := utils.GetData(url)
	// json := map[string]string{
	// 	"location_name": city + ", " + country,
	// }
	c.Ctx.Output.Header("Content-Type", "application/json")
	c.Ctx.Output.JSON(json, true, true)
}

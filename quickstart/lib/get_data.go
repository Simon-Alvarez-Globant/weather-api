package lib

import (
	"bapi/quickstart/lib/http"
	"bapi/quickstart/models"
	"bapi/quickstart/utils"
	"strings"
)

func GetData(params http.HttpParams) (jsons http.ResponseStruct) {
	if params.City == "" || params.Country == "" {
		return
	}

	configs := utils.GetConfigs()

	var url strings.Builder
	url.WriteString(configs.Externalapi + params.City + "," + params.Country + "&appid=" + configs.Appid)
	jsons = models.GetWeather(url.String())
	return
}

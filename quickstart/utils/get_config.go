package utils

import (
	"github.com/astaxie/beego"
)

func GetConfigs(key string) string {
	return beego.AppConfig.String(key)
}

package utils

import (
	"github.com/astaxie/beego"
)

type Config struct {
	Appname         string
	Httpport        int
	Runmode         string
	Appid           string
	Externalapi     string
	Copyrequestbody bool
	Queryraw        bool
}

func GetConfigs() (c Config) {
	c.Appname = beego.AppConfig.String("appname")
	c.Appid = beego.AppConfig.String("appid")
	// c.Httpport, _ = beego.AppConfig.Int("httpport")
	// c.Runmode = beego.AppConfig.String("runmode")
	c.Externalapi = beego.AppConfig.String("externalapi")
	// c.Copyrequestbody, _ = beego.AppConfig.Bool("copyrequestbody")
	c.Queryraw, _ = beego.AppConfig.Bool("queryraw")
	return
}

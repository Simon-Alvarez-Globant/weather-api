package models

import (
	"bapi/quickstart/lib/http"
	"bapi/quickstart/utils"
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Register struct {
	Id        int       `orm:"column(id)"` // nolint
	Url       string    `orm:"column(url)"`
	Response  string    `orm:"column(response)"`
	Timestamp time.Time `orm:"column(timestamp);type(datetime)"`
}

const (
	TableName = "weather_api"
)

func init() {
	orm.RegisterModel(new(Register))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@(localhost:3307)/weather_requests_db")
	orm.DefaultTimeLoc = time.Local
}

// TableName returns the name of the table
func (u *Register) TableName() string {
	return TableName
}

func GetWeather(url string) (jsons http.ResponseStruct) {
	configs := utils.GetConfigs()

	if configs.Queryraw {
		return getRaw(url)
	}
	return getOrm(url)

}

func Create(url string) (jsons http.ResponseStruct) {
	o := orm.NewOrm()
	jsons = http.GetDataApi(url)
	response, err := json.Marshal(jsons)
	if err != nil {
		fmt.Println(err)
	}
	var register Register
	register.Response = string(response)
	register.Url = url
	register.Timestamp = time.Now().Local()

	id, err := o.Insert(&register)
	if err == nil {
		fmt.Println("New register id:", id)
	} else {
		fmt.Println(err)
	}
	return
}

func getRaw(url string) (jsons http.ResponseStruct) {
	return
}

func getOrm(url string) (jsons http.ResponseStruct) {
	o := orm.NewOrm()
	qs := o.QueryTable("weather_api")

	var request []orm.Params

	exists := qs.Filter("url", url).Exist()

	if exists {
		num, err := qs.Filter("url", url).OrderBy("-timestamp").Limit(1).Values(&request)
		if err != nil {
			fmt.Println(err)
			return
		}
		jsons = createOrRead(request, num, url)
	} else {
		fmt.Println("Creating register...")
		jsons = Create(url)
	}
	return
}

func timelapse(t int64) bool {
	now := time.Now().Add(time.Hour * 5).Unix()
	duration := now - t
	if duration <= 300 {
		return true
	}
	return false
}

func createOrRead(request []orm.Params, num int64, url string) (jsons http.ResponseStruct) {
	if num != 0 {
		fmt.Println("Getting Register...")
		timestamp := request[0]["Timestamp"].(time.Time).Unix()
		if timelapse(timestamp) {
			err := json.Unmarshal([]byte(request[0]["Response"].(string)), &jsons)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Register is too old to fetch")
			fmt.Println("Creating new register...")
			jsons = Create(url)
		}
	}
	return
}

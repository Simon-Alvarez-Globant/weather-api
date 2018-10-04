package models

import (
	"bapi/quickstart/utils"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Register struct {
	Id        int       `orm:"column(id)"` // nolint
	City      string    `orm:"column(city)"`
	Response  string    `orm:"column(response)"`
	Timestamp time.Time `orm:"column(timestamp);type(datetime)"`
}

type Getter interface {
	GetReq(string) (string, time.Time)
}

const (
	TableName = "weather"
)

func init() {
	driver := utils.GetConfigs("drivername")
	alias := utils.GetConfigs("aliasname")
	dataSource := utils.GetConfigs("datasource")
	orm.RegisterModel(new(Register))
	orm.RegisterDriver(driver, orm.DRMySQL)
	orm.RegisterDataBase(alias, driver, dataSource)
	orm.DefaultTimeLoc = time.Local
}

// TableName returns the name of the table
func (u *Register) TableName() string {
	return TableName
}

func Get(g Getter, city string) (response string, timestamp time.Time) {
	return g.GetReq(city)
}

func Create(city string, response string) {
	o := orm.NewOrm()

	var register Register
	register.Response = response
	register.City = city
	register.Timestamp = time.Now().Local()

	id, err := o.Insert(&register)
	if err == nil {
		fmt.Println("New register id:", id)
	} else {
		fmt.Println(err)
	}
}

func getResponse(num int64, r []orm.Params) (response string, timestamp time.Time) {
	if num != 0 {
		response = r[0]["Response"].(string)
		timestamp = r[0]["Timestamp"].(time.Time)
	}
	return

}

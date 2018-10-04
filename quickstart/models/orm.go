package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Orm struct {
	str string
}

func (r Orm) GetReq(city string) (response string, timestamp time.Time) {

	o := orm.NewOrm()
	qs := o.QueryTable("weather")

	var request []orm.Params
	num, err := qs.Filter("city", city).OrderBy("-timestamp").Limit(1).Values(&request)
	if err != nil {
		fmt.Println(err)
		return
	}

	response, timestamp = getResponse(num, request)
	return
}

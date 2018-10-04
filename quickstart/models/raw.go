package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Raw struct {
	str string
}

func (r Raw) GetReq(city string) (response string, timestamp time.Time) {
	o := orm.NewOrm()

	var request []orm.Params
	num, err := o.Raw("SELECT * FROM weather WHERE city=? ORDER BY id DESC LIMIT 1", city).Values(&request)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(request[0]["response"])
	response, ts := getResponser(num, request)
	t, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	timestamp = time.Unix(t, 0)
	return
}
func getResponser(num int64, r []orm.Params) (response string, timestamp string) {
	if num != 0 {
		response = r[0]["response"].(string)
		timestamp = r[0]["timestamp"].(string)
	}
	return

}

package main

import (
	_ "bapi/quickstart/routers"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

}
func main() {
	beego.Run()
}

package test

import (
	_ "bapi/quickstart/routers"
	"bapi/quickstart/utils"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestBeego is a sample to run an endpoint test
func TestUtils(t *testing.T) {
	// var ResponseStruct struct {
	// 	LocationName   string
	// 	temperature    string
	// 	Wind           string
	// 	Cloudines      string
	// 	Pressure       string
	// 	Humidity       string
	// 	Sunrise        string
	// 	Sunset         string
	// 	GeoCoordinates []float64
	// 	RequestedTime  string
	// }
	url := "http://localhost:8080/weather?city=Bogota&country=co"
	Convey("Given a propper url, GetData", t, func() {
		// Convey("Returns a ResponseStruct type", func() {
		// 	So(utils.GetData(url), ShouldResemble, ResponseStruct)
		// })
		Convey("Should not panic", func() {
			So(utils.GetData(url), ShouldNotPanic)
		})
	})
}

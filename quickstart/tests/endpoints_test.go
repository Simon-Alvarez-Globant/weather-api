package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

type TestController struct {
	beego.Controller
}

func (tc *TestController) Get() {
	tc.Data["Username"] = "astaxie"
	tc.Ctx.Output.Body([]byte("ok"))
}

func (tc *TestController) Put() {
	tc.Data["Username"] = "astaxie"
	tc.Ctx.Output.Body([]byte("ok"))

}

func TestGetWeather(t *testing.T) {
	r, _ := http.NewRequest("GET", "/weather", nil)
	w := httptest.NewRecorder()
	handler := beego.NewControllerRegister()
	handler.Add("/weather", &TestController{}, "*:Get")
	handler.ServeHTTP(w, r)

	Convey("Test Get weather Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})

	})
}

func TestScheduler(t *testing.T) {
	r, _ := http.NewRequest("PUT", "/scheduler/weather", nil)
	w := httptest.NewRecorder()
	handler := beego.NewControllerRegister()
	handler.Add("/scheduler/weather", &TestController{}, "*:Put")
	handler.ServeHTTP(w, r)

	Convey("Test scheduler Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})

	})
}

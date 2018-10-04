package scheduler

import (
	"bapi/quickstart/lib"
	"bapi/quickstart/lib/http"

	"github.com/astaxie/beego/toolbox"
)

func Schedule(params http.HttpParams) (status int) {
	if params.City == "" || params.Country == "" {
		return 400
	}

	task := toolbox.NewTask(params.City, "0 0 */1 * * *", func() error {
		lib.GetData(params)
		return nil
	})

	toolbox.AddTask(params.City, task)
	toolbox.StopTask()
	toolbox.StartTask()

	return 202
}

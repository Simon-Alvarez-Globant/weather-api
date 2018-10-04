package scheduler

import (
	"bapi/quickstart/lib"

	"github.com/astaxie/beego/toolbox"
)

type RequestError struct {
	Message string
}

func (e RequestError) Error() string {
	return e.Message
}

type Request struct {
	City    string
	Country string
}

func Schedule(city, country string) error { // add error
	if city == "" || country == "" {
		return RequestError{
			"No city or country inside the request",
		}
	}

	task := toolbox.NewTask(city, "0 0 */1 * * *", func() error {
		lib.GetData(city, country)
		return nil
	})

	toolbox.AddTask(city, task)
	toolbox.StopTask()
	toolbox.StartTask()

	return nil
}

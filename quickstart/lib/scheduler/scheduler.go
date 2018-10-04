package scheduler

import (
	"bapi/quickstart/lib"
	"fmt"

	"github.com/astaxie/beego/toolbox"
)

type RequestError struct {
	Message string
}

func (e RequestError) Error() string {
	return e.Message
}

func Schedule(city, country string) error { // add error
	if city == "" || country == "" {
		return RequestError{
			"No city or country inside the request",
		}
	}

	task := toolbox.NewTask(city, "0 0 */1 * * *", func() error {
		_, err := lib.GetData(city, country)
		if err != nil {
			fmt.Println("ERROR::::::", err)
		}
		return nil
	})

	toolbox.AddTask(city, task)
	toolbox.StopTask()
	toolbox.StartTask()

	return nil
}

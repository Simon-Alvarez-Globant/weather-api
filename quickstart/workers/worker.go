package worker

import (
	"fmt"

	"github.com/astaxie/beego/httplib"
)

func worker(jobChan <-chan string, results chan<- string) {
	for job := range jobChan {
		req := httplib.Get(job)
		data, errReq := req.String()
		if errReq != nil {
			fmt.Println(errReq)
		}
		results <- data
	}
}

var (
	jobChan = make(chan string, 5)
	results = make(chan string, 5)
)

func init() {

	go worker(jobChan, results)
}

func enqueue(job string) bool {
	select {
	case jobChan <- job:
		return true
	default:
		return false
	}
}

func RequestOpenWeather(url string) string {
	enqueue(url)
	return <-results

}

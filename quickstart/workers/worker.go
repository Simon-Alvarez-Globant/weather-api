// package worker

// import "bapi/quickstart/utils"

// func worker(jobChan <-chan Job, resultChan chan<- results) {
// 	for job := range jobChan {
// 		process(job)
// 	}
// }

// func TryEnqueue(job Job, jobChan <-chan Job) bool {
// 	select {
// 	case jobChan <- job:
// 		return true
// 	default:
// 		return false
// 	}
// }

// func main() {
// 	jobChan := make(chan utils.GetData(), 5)
// 	resultChan : make(chan results,5)
// 	// start the worker
// 	go worker(jobChan)

// 	// enqueue a job
// 	jobChan <- job
// }

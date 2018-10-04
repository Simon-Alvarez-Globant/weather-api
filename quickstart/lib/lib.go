package lib

import (
	"bapi/quickstart/lib/http"
	"bapi/quickstart/models"
	"encoding/json"
	"fmt"
	"time"
)

type ResponseStruct struct {
	LocationName   string
	Temperature    string
	Wind           string
	Cloudines      string
	Pressure       string
	Humidity       string
	Sunrise        string
	Sunset         string
	GeoCoordinates []float64
	RequestedTime  string
}
type RequestError struct {
	Message string
}

func (e RequestError) Error() string {
	return e.Message
}
func GetData(city string, country string) (jsons interface{}, err error) {
	if city == "" || country == "" {
		return "", RequestError{
			"No city or country inside the request",
		}
	}
	implementation := models.SetImplementation()
	jsons, err = createOrRead(implementation, city, country)
	return
}
func createOrRead(implementation models.Getter, city, country string) (interface{}, error) {
	// Read Db
	data, timestamp := models.Get(implementation, city)

	// If city is in db
	if data != "" {
		// If time is less than 300 seconds
		if timelapse(timestamp.Unix()) {
			fmt.Println("Getting register...")
			return unmarshal(data)
		} else {
			fmt.Println("Register is too old to fetch!")
		}
	}

	// If city is not in db or time is more than 300 seconds
	// get data from the external api
	fmt.Println("Fetching new data...")
	response := http.Get(city, country)
	// create new register in db
	fmt.Println("Creating new register...")
	models.Create(city, response)
	return unmarshal(response)
}

func timelapse(t int64) bool {
	now := time.Now().Add(time.Hour * 5).Unix()
	duration := now - t
	if duration <= 300 {
		return true
	}
	return false
}

func unmarshal(s string) (jsons ResponseStruct, err error) {
	err = json.Unmarshal([]byte(s), &jsons)
	if err != nil {
		fmt.Println(err)
	}
	return
}

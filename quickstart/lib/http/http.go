package http

import (
	"bapi/quickstart/utils"
	worker "bapi/quickstart/workers"
	"encoding/json"
	"fmt"
	"strings"

	"strconv"
	"time"
)

type HttpParams struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

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

type raw struct {
	Coord      *coord        `json: "coord"`
	Weather    []interface{} `json: "weather"`
	Base       string        `json:"base"`
	Main       *Main         `json:"main"`
	Visibility int           `json: "visibility"`
	Wind       *wind         `json: "wind"`
	Clouds     *clouds       `json: "clouds"`
	Dt         int           `json: "dt"`
	Sys        *sys          `json: "sys"`
	Id         int           `json: "id"`
	Name       string        `json: "name"`
	Cod        int           `json: "cod"`
}

type coord struct {
	Lon float64 `json: "lon"`
	Lat float64 `json: "lat"`
}

type weather struct{}

type Main struct {
	Temp     float64 `json: "temp"`
	Pressure int     `json: "pressure"`
	Humidity int     `json: "humidity"`
	TempMin  float32 `json: "temp_min"`
	TempMax  float32 `json: "temp_max"`
}

type wind struct {
	Speed float64 `json: "speed"`
	Deg   float64 `json: "deg"`
}

type clouds struct {
	All int `json: "all"`
}

type sys struct {
	Typee   int     `json: "type"`
	Id      int     `json: "id"`
	Message float64 `json: "message"`
	Country string  `json: "country"`
	Sunrise int64   `json: "sunrise"`
	Sunset  int64   `json: "sunset"`
}

func Get(city, country string) (response string) {
	url := createUrl(city, country)
	jsons := getDataApi(url)
	resp, err := json.Marshal(jsons)
	if err != nil {
		fmt.Println(err)
	}
	response = string(resp)
	return
}

func getDataApi(url string) (jsons ResponseStruct) {
	data := worker.RequestOpenWeather(url)
	var raw *raw

	err := json.Unmarshal([]byte(data), &raw)
	if err != nil {
		fmt.Print(err)
	}

	jsons.prepData(raw)

	return

}

func createUrl(city, country string) string {
	urlInit := utils.GetConfigs("externalapi")
	appid := utils.GetConfigs("appid")

	return stringify(urlInit + city + "," + country + "&appid=" + appid)

}

func stringify(s string) string {
	var x strings.Builder
	x.WriteString(s)
	return x.String()
}

func (r *ResponseStruct) prepData(raw *raw) {
	t := raw.Main.Temp - 273.15
	temp := strconv.FormatFloat(t, 'f', 2, 64)
	r.Temperature = stringify(temp + " ºC")

	s := raw.Wind.Speed
	windCondition := windCondition(s)

	d := raw.Wind.Deg
	windDirection := windDirection(d)

	r.Wind = stringify(windCondition + ", " + strconv.FormatFloat(s, 'f', 2, 64) + " m/s, " + windDirection)

	r.Cloudines = raw.Weather[0].(map[string]interface{})["description"].(string)

	r.Pressure = stringify(strconv.Itoa(raw.Main.Pressure) + " hpa")

	r.Humidity = stringify(strconv.Itoa(raw.Main.Humidity) + "%")

	r.Sunset = time.Unix(raw.Sys.Sunset, 0).Format("15:04")

	r.Sunrise = time.Unix(raw.Sys.Sunrise, 0).Format("15:04")

	r.GeoCoordinates = []float64{raw.Coord.Lat, raw.Coord.Lon}

	r.RequestedTime = time.Now().Format("2006-01-02 15:04:05")

	r.LocationName = stringify(raw.Name + ", " + raw.Sys.Country)
}

func windCondition(s float64) string {
	switch {
	case s < 4:
		return "Gentle breeze"
	case s < 10:
		return "Windy"
	default:
		return "Be careful, crazy wind outside!"
	}
}

func windDirection(d float64) string {
	switch {
	case d < 11.25:
		return "north"
	case d <= 33.75:
		return "north-northeast"
	case d <= 56.25:
		return "northeast"
	case d <= 78.75:
		return "east-northeast"
	case d <= 101.25:
		return "east"
	case d <= 123.75:
		return "east-southeast"
	case d <= 146.25:
		return "southeast"
	case d <= 168.75:
		return "south-southeast"
	case d <= 191.25:
		return "south"
	case d <= 213.75:
		return "south-southwest"
	case d <= 236.25:
		return "southwest"
	case d <= 258.75:
		return "west-southwest"
	case d <= 281.25:
		return "west"
	case d <= 303.75:
		return "west-northwest"
	case d <= 326.25:
		return "northwest"
	case d <= 348.75:
		return "north-northweast"
	default:
		return "north"
	}
}

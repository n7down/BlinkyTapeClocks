package usgsapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	usgsApiFeedUrl = "https://earthquakes.usgs.gov/feeds/v1.0"
)

func GetAllEarthquakesPastHour() (AllEarthquakesPastHour, error) {
	url := "/summary/all_hour.json"
	apiUrl := usgsApiFeedUrl + url
	response, err := http.Get(apiUrl)
	if err != nil {
		return AllEarthquakesPastHour{}, err
	}

	temp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return AllEarthquakesPastHour{}, err
	}

	n := AllEarthquakesPastHour{}
	err = json.Unmarshal(temp, &n)
	if err != nil {
		return AllEarthquakesPastHour{}, nil
	}

	return n, nil
}

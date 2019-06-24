package spacexapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	spaceXApiUrl = "https://api.spacexdata.com/v3/"
)

func GetNextLaunch() (NextLaunch, error) {
	url := "launches/next"
	apiUrl := spaceXApiUrl + url
	response, err := http.Get(apiUrl)
	if err != nil {
		return NextLaunch{}, err
	}

	temp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return NextLaunch{}, err
	}

	n := NextLaunch{}
	err = json.Unmarshal(temp, &n)
	if err != nil {
		return NextLaunch{}, nil
	}
	return n, nil
}

func GetAPIInfo() (ApiInfo, error) {
	response, err := http.Get(spaceXApiUrl)
	if err != nil {
		return ApiInfo{}, err
	}

	temp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ApiInfo{}, err
	}

	n := ApiInfo{}
	err = json.Unmarshal(temp, &n)
	if err != nil {
		return ApiInfo{}, nil
	}
	return n, nil
}

func GetRoadster() (Roadster, error) {
	url := "roadster"
	apiUrl := spaceXApiUrl + url
	response, err := http.Get(apiUrl)
	if err != nil {
		return Roadster{}, err
	}

	temp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Roadster{}, err
	}

	n := Roadster{}
	err = json.Unmarshal(temp, &n)
	if err != nil {
		return Roadster{}, nil
	}
	return n, nil
}

func GetRocket(rocketID string) (Rocket, error) {
	url := "rockets/" + rocketID
	apiUrl := spaceXApiUrl + url
	response, err := http.Get(apiUrl)
	if err != nil {
		return Rocket{}, err
	}

	temp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Rocket{}, err
	}

	r := Rocket{}
	err = json.Unmarshal(temp, &r)
	if err != nil {
		return Rocket{}, nil
	}
	return r, nil
}

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

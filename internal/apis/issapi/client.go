package issapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	openNotifyApi = "http://api.open-notify.org/"
)

func GetIssNow() (IssNow, error) {
	url := "iss-now.json"
	apiUrl := openNotifyApi + url

	client := &http.Client{}
	req, err := http.NewRequest("GET", apiUrl, nil)

	res, err := client.Do(req)
	if err != nil {
		return IssNow{}, err
	}

	temp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return IssNow{}, err
	}

	i := IssNow{}
	err = json.Unmarshal(temp, &i)
	if err != nil {
		return IssNow{}, nil
	}
	return i, nil
}

func GetAstronauts() (Astronauts, error) {
	url := "astros.json"
	apiUrl := openNotifyApi + url

	client := &http.Client{}
	req, err := http.NewRequest("GET", apiUrl, nil)

	res, err := client.Do(req)
	if err != nil {
		return Astronauts{}, err
	}

	temp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Astronauts{}, err
	}

	a := Astronauts{}
	err = json.Unmarshal(temp, &a)
	if err != nil {
		return Astronauts{}, nil
	}
	return a, nil
}

func GetIssPassTimes(lat string, lon string) (IssPassTimes, error) {
	url := fmt.Sprintf("iss-pass.json?%s=LAT&lon=%s", lat, lon)
	apiUrl := openNotifyApi + url

	client := &http.Client{}
	req, err := http.NewRequest("GET", apiUrl, nil)

	res, err := client.Do(req)
	if err != nil {
		return IssPassTimes{}, err
	}

	temp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return IssPassTimes{}, err
	}

	a := IssPassTimes{}
	err = json.Unmarshal(temp, &a)
	if err != nil {
		return IssPassTimes{}, nil
	}
	return a, nil
}

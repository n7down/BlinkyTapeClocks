package githubapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	githubApiUrl = "https://api.github.com/"
)

func GetReleases(owner string, repo string, githubToken string) (Releases, error) {
	url := fmt.Sprintf("repos/%s/%s/releases", owner, repo)
	apiUrl := githubApiUrl + url

	token := fmt.Sprintf("token %s", githubToken)

	client := &http.Client{}
	req, err := http.NewRequest("GET", apiUrl, nil)
	req.Header.Set("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		return Releases{}, err
	}

	temp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Releases{}, err
	}

	t := Releases{}
	err = json.Unmarshal(temp, &t)
	if err != nil {
		return Releases{}, nil
	}
	return t, nil
}

func GetTags(owner string, repo string, githubToken string) (Tags, error) {
	url := fmt.Sprintf("repos/%s/%s/tags", owner, repo)
	apiUrl := githubApiUrl + url

	token := fmt.Sprintf("token %s", githubToken)

	client := &http.Client{}
	req, err := http.NewRequest("GET", apiUrl, nil)
	req.Header.Set("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		return Tags{}, err
	}

	temp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Tags{}, err
	}

	t := Tags{}
	err = json.Unmarshal(temp, &t)
	if err != nil {
		return Tags{}, nil
	}
	return t, nil
}

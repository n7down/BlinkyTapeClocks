package utils

import (
	"encoding/json"
	"errors"
	"github.com/n7down/Displays/internal/spacexapi"
	"io/ioutil"
	"os"
	"time"
)

type DisplayDataJson struct {
	jsonFileName string
	LastUpdate   time.Time            `json:"last_update"`
	NextLaunch   spacexapi.NextLaunch `json:"next_launch"`
	Rocket       spacexapi.Rocket     `json:"rocket"`
}

func exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err != nil, err
}

type DisplayData interface {
	Read() (interface{}, error)
	Create(string) error
	Write() error
}

func (d DisplayDataJson) Read() error {
	fileExists, err := exists(d.jsonFileName)
	if err != nil {
		return err
	}

	if fileExists {
		jsonFile, err := os.Open(d.jsonFileName)
		if err != nil {
			return err
		}
		defer jsonFile.Close()

		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			return err
		}
		json.Unmarshal(byteValue, &d)

	} else {
		return errors.New("file doesn't exists")
	}
	return nil
}

// FIXME: overwrite the file if it exists
func (d DisplayDataJson) Create() error {
	nextLaunch, err := spacexapi.GetNextLaunch()
	if err != nil {
		return err
	}

	rocket, err := spacexapi.GetRocket(d.NextLaunch.Rocket.RocketID)
	if err != nil {
		return err
	}

	d.NextLaunch = nextLaunch
	d.Rocket = rocket
	d.LastUpdate = time.Now()

	err = d.Write()
	if err != nil {
		return err
	}
	return nil
}

// FIXME: overwrite the file if it exists
func (d DisplayDataJson) Write() error {
	spaceXDataJson, err := json.Marshal(d)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(d.jsonFileName, spaceXDataJson, 0644)
	if err != nil {
		return err
	}
	return nil
}

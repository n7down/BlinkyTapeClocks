package utils

import (
	"encoding/json"
	"github.com/n7down/Displays/internal/spacexapi"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"time"
)

const (
	jsonFileName = "spacex.json"
)

var (
	spaceXData SpaceXData
)

type SpaceXData struct {
	LastUpdate time.Time            `json:"last_update"`
	NextLaunch spacexapi.NextLaunch `json:"next_launch"`
	Rocket     spacexapi.Rocket     `json:"rocket"`
}

func exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err != nil, err
}

// FIXME: turn this into an interface
func SaveToJson() {

	// FIXME: how do i make it so that i dont have to update this every time
	// get the data and store it if it does not exist
	fileExists, err := exists(jsonFileName)
	if err != nil {
		log.Error(err)
	}

	if fileExists {

		// if the file exists - read it in
		// Open our jsonFile
		jsonFile, err := os.Open(jsonFileName)
		// if we os.Open returns an error then handle it
		if err != nil {
			log.Error(err)
		}
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()

		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			log.Error(err)
		}

		// we unmarshal our byteArray which contains our
		// jsonFile's content into 'users' which we defined above
		json.Unmarshal(byteValue, &spaceXData)

	} else {
		// file does not exist - create the spacexdata.json file
		nextLaunch, err := spacexapi.GetNextLaunch()
		if err != nil {
			log.Error(err)
		}

		rocket, err := spacexapi.GetRocket(nextLaunch.Rocket.RocketID)
		if err != nil {
			log.Error(err)
		}

		spaceXData := SpaceXData{
			LastUpdate: time.Now(),
			NextLaunch: nextLaunch,
			Rocket:     rocket,
		}

		spaceXDataJson, err := json.Marshal(spaceXData)
		if err != nil {
			log.Error(err)
		}

		err = ioutil.WriteFile(jsonFileName, spaceXDataJson, 0644)
		if err != nil {
			log.Error(err)
		}
	}
}

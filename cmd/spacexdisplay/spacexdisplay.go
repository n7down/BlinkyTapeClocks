package main

import (
	"fmt"
	aurora "github.com/logrusorgru/aurora"
	"github.com/n7down/Displays/internal/spacexapi"
	"github.com/n7down/Displays/internal/util"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

const (
	spaceXApiVersion   = "3"
	spaceXClockVersion = "1.0.0"
)

var (
	spaceXData SpaceXData
)

type SpaceXData struct {
	LastUpdate time.Time            `json:"last_update"`
	NextLaunch spacexapi.NextLaunch `json:"next_launch"`
	Rocket     spacexapi.Rocket     `json:"rocket"`
}

func main() {

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

	rocketTypeCamelCase := spaceXData.Rocket.Engines.Type
	rocketTypeCamelCase = strings.ToUpper(string(rocketTypeCamelCase[0])) + rocketTypeCamelCase[1:]

	ipAddress, err := util.GetLocalIP()
	if err != nil {
		log.Error(err)
	}

	timeNow := time.Now().Format("Mon Jan _2, 2006 15:04:05")
	timeNowUTC := time.Now().UTC().Format("Mon Jan _2, 2006 15:04:05")
	nextLaunchTimeUtc := spaceXData.NextLaunch.LaunchDateUtc
	nextLaunchTimeUtcFormated := nextLaunchTimeUtc.Format("Mon Jan _2, 2006 15:04:05 ")
	elapsedTime := util.NewElapsedTime(time.Until(nextLaunchTimeUtc))

	fmt.Println(aurora.Cyan("____ ___  ____ ____ ____ _  _"))
	fmt.Println(aurora.Cyan("[__  |__] |__| |    |___  \\/"))  //   \tSpaceX API: \t[v%s]\n", spaceXApiVersion)
	fmt.Println(aurora.Cyan("___] |    |  | |___ |___ _/\\_")) //  \tVersion: \t[v%s]\n", spaceXClockVersion)
	fmt.Println()
	fmt.Println("SYSTEM ========================================================")
	fmt.Printf(" IPv4: \t\t\t\t%s\n", ipAddress)
	fmt.Printf(" Time: \t\t\t\t%s\n", timeNow)
	fmt.Printf(" Time UTC: \t\t\t%s\n", timeNowUTC)
	fmt.Println("LAUNCH ========================================================")
	fmt.Printf(" Mission Name: \t\t\t%s\n", spaceXData.NextLaunch.MissionName)
	fmt.Printf(" Flight Number: \t\t%d\n", spaceXData.NextLaunch.FlightNumber)
	fmt.Printf(" Launch Site: \t\t\t%s\n", spaceXData.NextLaunch.LaunchSite.SiteName)
	fmt.Printf(" Launch Time UTC: \t\t%s\n", nextLaunchTimeUtcFormated)
	fmt.Printf(" Elapsed Time: \t\t\t%s\n", elapsedTime.String())
	fmt.Printf(" [%s]\n", elapsedTime.PrintBar())
	fmt.Println("ROCKET ========================================================")
	fmt.Printf(" Name: \t\t\t\t%s\n", spaceXData.NextLaunch.Rocket.RocketName)
	fmt.Printf(" Engines: \t\t\t%d x %s %s\n", spaceXData.Rocket.Engines.Number, rocketTypeCamelCase, spaceXData.Rocket.Engines.Version)
}

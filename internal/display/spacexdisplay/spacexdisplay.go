package spacexdisplay

import (
	"bytes"
	"fmt"
	"github.com/n7down/timelord/internal/spacexapi"
	"github.com/n7down/timelord/internal/utils"

	"github.com/spf13/viper"
	"strings"
	"time"
)

const (
	spaceXApiVersion = "3"
)

type SpaceXDisplay struct {
	LastUpdate  time.Time            `json:"last_update"`
	NextLaunch  spacexapi.NextLaunch `json:"next_launch"`
	Rocket      spacexapi.Rocket     `json:"rocket"`
	config      *viper.Viper
	refreshTime time.Time
}

func refreshDisplay(config *viper.Viper) (*SpaceXDisplay, error) {
	nextLaunch, err := spacexapi.GetNextLaunch()
	if err != nil {
		return nil, err
	}

	rocket, err := spacexapi.GetRocket(nextLaunch.Rocket.RocketID)
	if err != nil {
		return nil, err
	}

	refreshTime := nextLaunch.LaunchDateUtc.AddDate(0, 0, 7)

	d := SpaceXDisplay{
		LastUpdate:  time.Now(),
		NextLaunch:  nextLaunch,
		Rocket:      rocket,
		config:      config,
		refreshTime: refreshTime,
	}

	return &d, nil
}

func NewSpaceXDisplay(config *viper.Viper) (*SpaceXDisplay, error) {
	d, err := refreshDisplay(config)
	if err != nil {
		return &SpaceXDisplay{}, err
	}
	return d, nil
}

// Get new data when needed
func (s SpaceXDisplay) Refresh() error {
	elapsed := time.Since(s.refreshTime)
	if elapsed > 0 {
		// refresh the data
		d, err := refreshDisplay(s.config)
		if err != nil {
			return err
		}
		s = *d
	}
	return nil
}

// Render the display
func (s SpaceXDisplay) Render() string {
	var buffer bytes.Buffer

	rocketType := strings.Title(s.Rocket.Engines.Type)
	//propellant1 := strings.Title(s.Rocket.Engines.Propellant1)
	//propellant2 := strings.Title(s.Rocket.Engines.Propellant2)
	timeNow := time.Now().Format("Mon Jan _2, 2006 15:04:05")
	timeNowUTC := time.Now().UTC().Format("Mon Jan _2, 2006 15:04:05")
	nextLaunchTimeUtc := s.NextLaunch.LaunchDateUtc
	nextLaunchTimeUtcFormated := nextLaunchTimeUtc.Format("Mon Jan _2, 2006 15:04:05 ")
	elapsedTime := utils.NewElapsedTime(time.Until(nextLaunchTimeUtc))
	timelordVersion := s.config.GetString("version")

	buffer.WriteString(fmt.Sprintf(" SPACEX\t[v%s][%s]\n", spaceXApiVersion, timelordVersion))
	buffer.WriteString(fmt.Sprintf("\n"))
	buffer.WriteString(" MISSION --------------------------------------------------\n")
	buffer.WriteString(fmt.Sprintf("  Name: %s \t\t\tFlight Number: %d\n", s.NextLaunch.MissionName, s.NextLaunch.FlightNumber))
	buffer.WriteString(" ROCKET ---------------------------------------------------\n")
	buffer.WriteString(fmt.Sprintf("  Name: %s \t\tEngines: %d x %s %s\n", s.NextLaunch.Rocket.RocketName, s.Rocket.Engines.Number, rocketType, s.Rocket.Engines.Version))
	//buffer.WriteString(fmt.Sprintf("  Propellant: \t\t\t%s/%s\n", propellant1, propellant2))
	buffer.WriteString("  Thrust\n")
	buffer.WriteString(fmt.Sprintf("  - Weight: \t\t\t%v\n", s.Rocket.Engines.ThrustToWeight))
	buffer.WriteString(fmt.Sprintf("  - Sea Level (kN/lbf): \t%v/%v\n", s.Rocket.Engines.ThrustSeaLevel.KN, s.Rocket.Engines.ThrustSeaLevel.Lbf))
	buffer.WriteString(fmt.Sprintf("  - Vacuum (kN/lbf): \t\t%v/%v\n", s.Rocket.Engines.ThrustVacuum.KN, s.Rocket.Engines.ThrustVacuum.Lbf))
	buffer.WriteString(" LAUNCH ---------------------------------------------------\n")
	buffer.WriteString(fmt.Sprintf("  Launch Site: \t\t\t%s\n", s.NextLaunch.LaunchSite.SiteName))
	buffer.WriteString(fmt.Sprintf("  Time: \t\t\t%s\n", timeNow))
	buffer.WriteString(fmt.Sprintf("  Time UTC: \t\t\t%s\n", timeNowUTC))
	buffer.WriteString(fmt.Sprintf("  Launch Time UTC: \t\t%s\n", nextLaunchTimeUtcFormated))
	buffer.WriteString(fmt.Sprintf("  Elapsed Time: \t\t%s\n", elapsedTime.String()))
	buffer.WriteString(fmt.Sprintf("  [%s]\n", elapsedTime.PrintBar()))

	return buffer.String()
}

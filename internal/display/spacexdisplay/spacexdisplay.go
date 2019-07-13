package spacexdisplay

import (
	"bytes"
	"fmt"
	"github.com/n7down/timelord/internal/spacexapi"
	"github.com/n7down/timelord/internal/utils"

	//aurora "github.com/logrusorgru/aurora"
	log "github.com/sirupsen/logrus"
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

func NewSpaceXDisplay(config *viper.Viper) (*SpaceXDisplay, error) {
	nextLaunch, err := spacexapi.GetNextLaunch()
	if err != nil {
		return nil, err
	}

	rocket, err := spacexapi.GetRocket(nextLaunch.Rocket.RocketID)
	if err != nil {
		return nil, err
	}

	// TODO: set the refresh time to the next rocket launch time +1 day

	d := SpaceXDisplay{
		LastUpdate: time.Now(),
		NextLaunch: nextLaunch,
		Rocket:     rocket,
		config:     config,
	}

	return &d, nil
}

// Returns true if this object should be recreated
func (s SpaceXDisplay) Refresh() bool {
	return false
}

// Render the display
func (s SpaceXDisplay) Render() string {
	var buffer bytes.Buffer

	rocketType := strings.Title(s.Rocket.Engines.Type)

	ipAddress, err := utils.GetLocalIP()
	if err != nil {
		log.Error(err)
	}

	hostname, err := utils.GetHostName()
	if err != nil {
		log.Error(err)
	}

	timeNow := time.Now().Format("Mon Jan _2, 2006 15:04:05")
	timeNowUTC := time.Now().UTC().Format("Mon Jan _2, 2006 15:04:05")
	nextLaunchTimeUtc := s.NextLaunch.LaunchDateUtc
	nextLaunchTimeUtcFormated := nextLaunchTimeUtc.Format("Mon Jan _2, 2006 15:04:05 ")
	elapsedTime := utils.NewElapsedTime(time.Until(nextLaunchTimeUtc))

	// FIXME: get the pitftversion from the config
	//piTftVersion := ""
	piTftVersion := s.config.GetString("version")

	//fmt.Println("")
	//fmt.Println("     ███████╗██████╗  █████╗  ██████╗███████╗██╗  ██╗")
	//fmt.Println("     ██╔════╝██╔══██╗██╔══██╗██╔════╝██╔════╝╚██╗██╔╝")
	//fmt.Println("     ███████╗██████╔╝███████║██║     █████╗   ╚███╔╝ ")
	//fmt.Println("     ╚════██║██╔═══╝ ██╔══██║██║     ██╔══╝   ██╔██╗")
	//fmt.Println("     ███████║██║     ██║  ██║╚██████╗███████╗██╔╝ ██╗")
	//fmt.Println("     ╚══════╝╚═╝     ╚═╝  ╚═╝ ╚═════╝╚══════╝╚═╝  ╚═╝")

	buffer.WriteString("\n")
	buffer.WriteString(fmt.Sprintf(" SPACEX: [v%s]\n", spaceXApiVersion))
	buffer.WriteString(fmt.Sprintf(" PITFTDisplay: [%s]\n", piTftVersion))
	buffer.WriteString(fmt.Sprintf(" Hostname: [%s] \t\tIPv4: [%s/24]\n", hostname, ipAddress))
	buffer.WriteString("\n")
	buffer.WriteString(" MISSION --------------------------------------------------\n")
	buffer.WriteString(fmt.Sprintf("  Name: %s \t\t\tFlight Number: %d\n", s.NextLaunch.MissionName, s.NextLaunch.FlightNumber))
	buffer.WriteString(" LAUNCH ---------------------------------------------------\n")
	buffer.WriteString(fmt.Sprintf("  Launch Site: \t\t\t%s\n", s.NextLaunch.LaunchSite.SiteName))
	buffer.WriteString(fmt.Sprintf("  Time: \t\t\t%s\n", timeNow))
	buffer.WriteString(fmt.Sprintf("  Time UTC: \t\t\t%s\n", timeNowUTC))
	buffer.WriteString(fmt.Sprintf("  Launch Time UTC: \t\t%s\n", nextLaunchTimeUtcFormated))
	buffer.WriteString(fmt.Sprintf("  Elapsed Time: \t\t%s\n", elapsedTime.String()))
	buffer.WriteString(fmt.Sprintf("  [%s]\n", elapsedTime.PrintBar()))
	buffer.WriteString(" ROCKET ---------------------------------------------------\n")
	buffer.WriteString(fmt.Sprintf("  Name: %s \t\tEngines: %d x %s %s\n", s.NextLaunch.Rocket.RocketName, s.Rocket.Engines.Number, rocketType, s.Rocket.Engines.Version))
	return buffer.String()
}

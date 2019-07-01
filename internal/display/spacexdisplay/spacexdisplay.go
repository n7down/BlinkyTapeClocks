package spacexdisplay

import (
	"fmt"
	"github.com/n7down/PITFTDisplays/internal/spacexapi"
	"github.com/n7down/PITFTDisplays/internal/utils"

	//aurora "github.com/logrusorgru/aurora"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

const (
	spaceXApiVersion   = "3"
	spaceXClockVersion = "1.0.0"
)

type SpaceXDisplay struct {
	LastUpdate time.Time            `json:"last_update"`
	NextLaunch spacexapi.NextLaunch `json:"next_launch"`
	Rocket     spacexapi.Rocket     `json:"rocket"`
}

func NewSpaceXDisplay() *SpaceXDisplay {
	nextLaunch, err := spacexapi.GetNextLaunch()
	if err != nil {
		log.Error(err)
	}

	rocket, err := spacexapi.GetRocket(nextLaunch.Rocket.RocketID)
	if err != nil {
		log.Error(err)
	}

	d := SpaceXDisplay{
		LastUpdate: time.Now(),
		NextLaunch: nextLaunch,
		Rocket:     rocket,
	}

	return &d
}

// Returns true if this object should be recreated
func (s SpaceXDisplay) Refresh() bool {
	return false
}

// Render the display
func (s SpaceXDisplay) Render() {
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

	fmt.Println("")
	fmt.Println("     ███████╗██████╗  █████╗  ██████╗███████╗██╗  ██╗")
	fmt.Println("     ██╔════╝██╔══██╗██╔══██╗██╔════╝██╔════╝╚██╗██╔╝")
	fmt.Println("     ███████╗██████╔╝███████║██║     █████╗   ╚███╔╝ ")
	fmt.Println("     ╚════██║██╔═══╝ ██╔══██║██║     ██╔══╝   ██╔██╗")
	fmt.Println("     ███████║██║     ██║  ██║╚██████╗███████╗██╔╝ ██╗")
	fmt.Println("     ╚══════╝╚═╝     ╚═╝  ╚═╝ ╚═════╝╚══════╝╚═╝  ╚═╝")
	fmt.Printf("\tSpaceX API: [v%s] \t Version: [v%s]\n", spaceXApiVersion, spaceXClockVersion)
	fmt.Println(" MISSION ======================================================")
	fmt.Printf("  Name: %s \t\t\tFlight Number: %d\n", s.NextLaunch.MissionName, s.NextLaunch.FlightNumber)
	fmt.Println(" LAUNCH ========================================================")
	fmt.Printf("  Launch Site: \t\t\t%s\n", s.NextLaunch.LaunchSite.SiteName)
	fmt.Printf("  Time: \t\t\t%s\n", timeNow)
	fmt.Printf("  Time UTC: \t\t\t%s\n", timeNowUTC)
	fmt.Printf("  Launch Time UTC: \t\t%s\n", nextLaunchTimeUtcFormated)
	fmt.Printf("  Elapsed Time: \t\t%s\n", elapsedTime.String())
	fmt.Printf("  [%s]\n", elapsedTime.PrintBar())
	fmt.Println(" ROCKET ========================================================")
	fmt.Printf("  Name: %s \t\tEngines: %d x %s %s\n", s.NextLaunch.Rocket.RocketName, s.Rocket.Engines.Number, rocketType, s.Rocket.Engines.Version)
	fmt.Println(" SYSTEM ========================================================")
	fmt.Printf("  Name: %s \t\tIPv4: %s/24\n", hostname, ipAddress)
}

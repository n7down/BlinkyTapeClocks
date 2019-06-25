package main

import (
	"errors"
	"fmt"
	"github.com/n7down/Displays/internal/spacexapi"
	log "github.com/sirupsen/logrus"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	spaceXApiVersion   = "3"
	spaceXClockVersion = "1.0.0"
)

// FIXME: this does not work
func fmtDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02dh %02dm %02ds", h, m, s)
}

func getLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("Unable to get ip address")
}

func main() {

	// FIXME: how do i make it so that i dont have to update this every time
	// get the data and store it if it does not exist
	nextLaunch, err := spacexapi.GetNextLaunch()
	if err != nil {
		log.Error(err)
	}

	rocket, err := spacexapi.GetRocket(nextLaunch.Rocket.RocketID)
	if err != nil {
		log.Error(err)
	}

	rocketTypeCamelCase := rocket.Engines.Type
	rocketTypeCamelCase = strings.ToUpper(string(rocketTypeCamelCase[0])) + rocketTypeCamelCase[1:]

	ipAddress, err := getLocalIP()
	if err != nil {
		log.Error(err)
	}

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	timeNow := time.Now().Format("Mon Jan _2, 2006 15:04:05")
	timeNowUTC := time.Now().UTC().Format("Mon Jan _2, 2006 15:04:05")
	nextLaunchTimeUtc := nextLaunch.LaunchDateUtc
	nextLaunchTimeUtcFormated := nextLaunchTimeUtc.Format("Mon Jan _2, 2006 15:04:05 ")
	elapsedTime := time.Until(nextLaunchTimeUtc)

	fmt.Println("____ ___  ____ ____ ____ _  _")
	fmt.Printf("[__  |__] |__| |    |___  \\/   \tSpaceX API: \t[v%s]\n", spaceXApiVersion)
	fmt.Printf("___] |    |  | |___ |___ _/\\_  \tVersion: \t[v%s]\n", spaceXClockVersion)
	fmt.Println()
	fmt.Println("SYSTEM ========================================================")
	fmt.Printf(" IPv4: \t\t\t\t%s\n", ipAddress)
	fmt.Printf(" Time: \t\t\t\t%s\n", timeNow)
	fmt.Printf(" Time UTC: \t\t\t%s\n", timeNowUTC)
	fmt.Println("LAUNCH ======================================================")
	fmt.Printf(" Mission Name: \t\t\t%s\n", nextLaunch.MissionName)
	fmt.Printf(" Flight Number: \t\t%d\n", nextLaunch.FlightNumber)
	fmt.Printf(" Launch Site: \t\t\t%s\n", nextLaunch.LaunchSite.SiteName)
	fmt.Printf(" Launch Time UTC: \t\t%s\n", nextLaunchTimeUtcFormated)
	fmt.Printf(" Elapsed Time: \t\t\t%s\n", elapsedTime)

	// TODO: show graph of elapsed time - show elapsed time after elapsed time is < 24 hours
	fmt.Print(" [\t\t\t\t\t\t\t]\n")
	fmt.Println("ROCKET =====================================================")
	fmt.Printf(" Name: \t\t\t\t%s\n", nextLaunch.Rocket.RocketName)
	fmt.Printf(" Engines: \t\t\t%d x %s %s\n", rocket.Engines.Number, rocketTypeCamelCase, rocket.Engines.Version)
}

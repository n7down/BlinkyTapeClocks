// MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWWWNNNNNNWWWWMMMMM
// MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWWNNNXXXXNNNNNWWWWMMMMMMMMM
// MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWNNXKKKKKXXNNWWMMMMMMMMMMMMMMMMMM
// MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWNXXK0000KXNNWMMMMMMMMMMMMMMMMMMMMMMMMM
// MWXKKKKKKKKKKK00KKKXWMMMMMMMNK0000000KKKK00000KKXWMMMMMMMMMMWNNMMMMMMMMMMMMMMMMWNXKK00KKKKKKK00KKKXNWWMMMMMMNKK0KKKKKKKKKKKKKKKKKXWMMMMWK000000XWMMMMMMMMMMMWWXK0OOOOKXNWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
// 0l;,,;;;;;;;;;;;;;;:oKWMMMMWk;,,,,;;;;;;;;;;;,,,:l0WMMMMMMMXd:oKWMMMMMMMMMMMMMNx:;,,,;;;;;;;;;;;;;:cokNMMMMWO:,,,,,,,,,,,,,,,,,,,oNMMMMNOl;,,,,;oOXWMMMMWNXK0OkkOKXNWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
// o'''oO000000000000000XWMMMMWx,'',d00000000000d;'''dNMMMMMMWO;'':kNMMMMMMMMMMMM0:'''ckOOOOOOOO00OOO0OOKWMMMMMXOkkkkkkkkkkkkkkkkOOOKWMMMMMMN0d:,''.,cONWXK0kkkO0XNWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
// o'',o0KKKKKKKKKKKKKXNWMMMMMWx,'';kWWWWWWWWWWW0:'''oNMMMMMMMWO:'',oKWMMMMMMMMMMO;.''dWMMMMMMMMMMMMMMMMMMMMMMMKxdddddddddddxONMMMMMMMMMMMMMMMWNOoclxkO0OkkkO0XWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
// k;'',,,,,,,,,,,,,,;;cxNMMMMWx,''':lllllllllll:,'',xWMMMMMMMMWKl,.':ONMMMMMMMMMO;'''dWMMMMMMMMMMMMMMMMMMMMMMWk,''';ccccccccxNMMMMMMMMMMMMMMMMMMWXK0OkkkO0OOXWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
// WKOkkxxxxxxxxxxxxo;'';OMMMMWx,'''cooooooooooooodx0NMWKkxxxxxxxo,''',dXWMMMMMMMO;'''dWMMMMMMMMMMMMMMMMMMMMMMMk,'',xNWWWWWWWWMMMMMMMMMMMMMMMMMNX0Okkkk00kl;,:d0NMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
// WWWWWWWWWWWWWWWWW0c'';OMMMMWx,'';OWMMMMMMMMMMMMMMMMW0o:::::::::;;,'.'cOWMMMMMMO;'''dNWWWWWWWWWWWWWWWWWMMMMMMk,'',xWMMMMMMMMMMMMMMMMMMMMMMWNKOkkkkOKNWXx:'''',cxKWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
// ollllllllllllllll:,''c0MMMMWx,'.;OMMMMMMMMMMMMMMMMMWNXXXXXXXXXXXX0o,'';dXWMMMMKc''';cllllllllllllllllkNMMMMWk,''':ooooooooooooood0WMMMMWX0kkkkkOKNWMMMWXkl;'''';lkXWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
// d:,''''''''''''''',;cOWMMMMWx,'':OMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNd,'',dXMMMMWOc;,'''''''''''''',;ld0WMMMMMk;'''''''''''''''''',dWMMMMXOkkkkk0NMMMMMMMMMNOc,'''',lKMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM

package main

import (
	"fmt"
	"github.com/n7down/Displays/internal/spacexapi"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

// TODO: build
// 1. every 1 min - get the current time in utc
// 2. query the spacex api to get the next launch in utc:
// - curl --location --request GET "https://api.spacexdata.com/v3/launches/next" | jq '.launch_date_utc'
// 3. translate the difference of the year, month, day, hour, minute, second and millisecond? into rgb
// - where each rgb value is between 0 and 255
// - for each new row get the rgb value below it and divide it by 2
// - so if there are 4 rows then the first row is the original rgb the 2nd row is the 1st row divided by 2
// - then the 3rd row is the 2nd row divided by 2 and so on
// 4. render to the blinky tape

const (
	spaceXApiVersion   = "3"
	spaceXClockVersion = "1.0.0"
)

func main() {

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

	timeNow := time.Now().UTC().Format("Mon Jan _2, 2006 15:04:05")
	nextLaunchTimeUtc := nextLaunch.LaunchDateUtc
	nextLaunchTimeUtcFormated := nextLaunchTimeUtc.Format("Mon Jan _2, 2006 15:04:05 ")

	fmt.Println("____ ___  ____ ____ ____ _  _")
	fmt.Printf("[__  |__] |__| |    |___  \\/   \tSpaceX API: \t[v%s]\n", spaceXApiVersion)
	fmt.Printf("___] |    |  | |___ |___ _/\\_  \tVersion: \t[v%s]\n", spaceXClockVersion)
	fmt.Println()
	fmt.Println("TIME ========================================================")
	fmt.Printf(" Time Now UTC: \t\t\t%s\n", timeNow)
	fmt.Println("LAUNCH ======================================================")
	fmt.Printf(" Mission Name: \t\t\t%s\n", nextLaunch.MissionName)
	fmt.Printf(" Flight Number: \t\t%d\n", nextLaunch.FlightNumber)
	fmt.Printf(" Launch Site: \t\t\t%s\n", nextLaunch.LaunchSite.SiteName)
	fmt.Printf(" Launch Time UTC: \t\t%s\n", nextLaunchTimeUtcFormated)
	fmt.Printf(" Elapsed Time: \t\t\t%s\n", time.Since(nextLaunchTimeUtc))

	// TODO: show graph of elapsed time - show elapsed time after elapsed time is < 24 hours
	fmt.Print(" [\t\t\t\t\t\t\t]\n")
	fmt.Println("ROCKET =====================================================")
	fmt.Printf(" Rocket Name: \t\t\t%s\n", nextLaunch.Rocket.RocketName)
	fmt.Printf(" Engines: \t\t\t%d\n", rocket.Engines.Number)
	fmt.Printf(" Name: \t\t\t\t%s\n", rocketTypeCamelCase)
	fmt.Printf(" Version: \t\t\t%s\n", rocket.Engines.Version)
}

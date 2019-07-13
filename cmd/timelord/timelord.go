package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"

	"github.com/n7down/timelord/internal/config"
	"github.com/n7down/timelord/internal/display"
	"github.com/n7down/timelord/internal/display/spacexdisplay"
	"github.com/n7down/timelord/internal/utils"
	//"github.com/n7down/timelord/internal/display/githubdisplay"
	//"github.com/n7down/timelord/internal/display/usgsdisplay"
)

var (
	Version = "v0.1"
	Build   string
)

func main() {
	versionPtr := flag.Bool("v", false, "show version and build")
	flag.Parse()
	if *versionPtr {
		fmt.Printf("pitftdisplay version %s build %s", Version, Build)
	} else {
		c, err := config.Config()
		if err != nil {
			log.Error(err)
		}

		c.Set("version", Version)

		spaceXDisplay, err := spacexdisplay.NewSpaceXDisplay(c)
		if err != nil {
			log.Error(err)
			return
		}

		//usgsDisplay, err := usgsdisplay.NewUsgsDisplay(c)
		//if err != nil {
		//log.Error(err)
		//return
		//}

		displayManager := display.NewDisplayManager()
		displayManager.AddDisplay(spaceXDisplay)
		//displayManager.AddDisplay(usgsDisplay)
		//displayManager.AddDisplay(githubdisplay.NewGithubReleasesDisplay(c))
		// TODO: render every second
		for {
			time.Sleep(time.Second)
			utils.ClearScreen()
			displayManager.Render()
		}
	}
}

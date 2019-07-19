package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"

	"github.com/n7down/timelord/internal/config"
	"github.com/n7down/timelord/internal/displays"
	"github.com/n7down/timelord/internal/displays/issdisplay"
	"github.com/n7down/timelord/internal/displays/spacexdisplay"
	//"github.com/n7down/timelord/internal/utils"
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

		displayManager := display.NewDisplayManager(time.Second * 5)

		issDisplay := issdisplay.NewIssDisplay(c)

		displayManager.AddDisplay(issDisplay)

		spaceXDisplay, err := spacexdisplay.NewSpaceXDisplay(c)
		if err != nil {
			log.Error(err)
			return
		}

		displayManager.AddDisplay(spaceXDisplay)

		for {
			time.Sleep(time.Second)
			//utils.ClearScreen()
			err := displayManager.Refresh()
			if err != nil {
				log.Error(err.Error)
				return
			}
			displayManager.Render()
		}
	}
}

package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	//"time"

	"github.com/n7down/pitftdisplays/internal/config"
	"github.com/n7down/pitftdisplays/internal/display"
	"github.com/n7down/pitftdisplays/internal/display/spacexdisplay"
	"github.com/n7down/pitftdisplays/internal/utils"
	//"github.com/n7down/pitftdisplays/internal/display/githubdisplay"
)

var (
	Version string
	Build   string
)

func main() {
	versionPtr := flag.Bool("v", false, "show version and build")
	flag.Parse()
	if *versionPtr {
		fmt.Printf("pitftdisplay version %s build %s", Version, Build)
	} else {
		_, err := config.Config()
		if err != nil {
			log.Error(err)
		}

		displayManager := display.NewDisplayManager()
		displayManager.AddDisplay(spacexdisplay.NewSpaceXDisplay())
		//displayManager.AddDisplay(githubdisplay.NewGithubReleasesDisplay(c))
		// TODO: render every second
		//for {
		//time.Sleep(time.Second)
		utils.ClearScreen()
		displayManager.Render()
		//}
	}
}

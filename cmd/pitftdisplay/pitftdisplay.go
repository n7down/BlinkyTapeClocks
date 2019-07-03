package main

import (
	"github.com/n7down/pitftdisplays/internal/config"
	"github.com/n7down/pitftdisplays/internal/display"
	"github.com/n7down/pitftdisplays/internal/display/spacexdisplay"
	"time"
	//"github.com/n7down/pitftdisplays/internal/display/githubdisplay"
	"flag"
	"fmt"
	"github.com/n7down/pitftdisplays/internal/utils"
	log "github.com/sirupsen/logrus"
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
		//displayManager.AddDisplay(githubdisplay.NewGithubDisplay(c))
		// TODO: render every second
		for {
			time.Sleep(time.Second)
			utils.ClearScreen()
			displayManager.Render()
		}
	}
}

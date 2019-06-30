package main

import (
	"github.com/n7down/PITFTDisplays/internal/config"
	"github.com/n7down/PITFTDisplays/internal/display"
	"github.com/n7down/PITFTDisplays/internal/display/spacexdisplay"
	log "github.com/sirupsen/logrus"
)

func main() {
	c, err := config.Config()
	if err != nil {
		log.Error(err)
	}

	displayManager := display.NewDisplayManager(c)
	displayManager.AddDisplay(spacexdisplay.NewSpaceXDisplay())
	displayManager.Render()
}

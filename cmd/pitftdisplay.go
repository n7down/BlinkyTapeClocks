package main

import (
	"github.com/n7down/PITFTDisplays/internal/config"
	//"github.com/n7down/PITFTDisplays/internal/display/spacexdisplay"
	log "github.com/sirupsen/logrus"
	//"github.com/spf13/viper"
)

func main() {
	//spacexdisplay.NewSpaceXDisplay().Render()
	c, err := config.Config()
	if err != nil {
		log.Error(err)
	}
	log.Info(c.Get("github"))
}

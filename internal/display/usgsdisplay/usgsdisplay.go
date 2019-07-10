package usgsdisplay

import (
	"bytes"
	"github.com/n7down/pitftdisplays/internal/usgsapi"
	"github.com/n7down/pitftdisplays/internal/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type UsgsDisplay struct {
	config              *viper.Viper
	EarthquakesPastHour usgsapi.AllEarthquakesPastHour
}

func getEarthquakesPastHour(e usgsapi.AllEarthquakesPastHour, count int) usgsapi.AllEarthquakesPastHour {
	r := e
	r.Features = e.Features[len(e.Features)-count : len(e.Features)]
	return r
}

func NewUsgsDisplay(config *viper.Viper) (*UsgsDisplay, error) {
	e, err := usgsapi.GetAllEarthquakesPastHour()
	if err != nil {
		return nil, err
	}

	d := UsgsDisplay{
		config:              config,
		EarthquakesPastHour: e,
	}

	return &d, nil
}

func (u UsgsDisplay) Refresh() bool {
	return false
}

func (u UsgsDisplay) Render() string {
	var buffer bytes.Buffer
	out, err := utils.ExecCommand("spark")
	if err != nil {
		log.Error(err)
	}
	buffer.WriteString(out)
	return buffer.String()
}

package issdisplay

import (
	"bytes"
	//"github.com/n7down/timelord/internal/apis/issapi"
	//log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type IssDisplay struct {
	config *viper.Viper
}

func NewIssDisplay(config *viper.Viper) *IssDisplay {
	i := IssDisplay{
		config: config,
	}

	return &i
}

func (i IssDisplay) Refresh() error {
	return nil
}

func (i IssDisplay) Render() string {
	var buffer bytes.Buffer
	buffer.WriteString("ISS")
	return buffer.String()
}

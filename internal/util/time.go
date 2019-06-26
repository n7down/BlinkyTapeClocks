package util

import (
	"bytes"
	"fmt"
	aurora "github.com/logrusorgru/aurora"
	//log "github.com/sirupsen/logrus"
	//"math"
	"time"
)

const (
	weekInHours = 168
)

type ElapsedTime struct {
	Hour   time.Duration
	Minute time.Duration
	Second time.Duration
}

func NewElapsedTime(d time.Duration) *ElapsedTime {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	e := ElapsedTime{
		Hour:   h,
		Minute: m,
		Second: s,
	}
	return &e
}

func (e ElapsedTime) String() string {
	return fmt.Sprintf("%02dh %02dm %02ds", e.Hour, e.Minute, e.Second)
}

func (e ElapsedTime) PrintBar() string {
	var buffer bytes.Buffer
	numberOfBars := 54
	bar := "■"

	// if greater then a week
	if e.Hour > weekInHours {

		// show all red
		for i := 0; i < numberOfBars; i++ {
			coloredBar := aurora.Sprintf(aurora.Red(bar))
			buffer.WriteString(coloredBar)
		}
	} else {

		// % increase = increase / original * 100

		// FIXME: not sure what is wrong here
		//percentage := fmt.Sprintf("%02d", e.Hour) / weekInHours
		//log.Info(fmt.Sprintf("%:%s", percentage))
		//numberOfRedBars := int(math.Round(percentage * float64(numberOfBars)))
		//log.Info(fmt.Sprintf("# red bars: %d", numberOfRedBars))
		//numberOfGreenBars := numberOfBars - numberOfRedBars
		//log.Info(fmt.Sprintf("# green bars: %d", numberOfGreenBars))
		//for i := 0; i <= numberOfGreenBars; i++ {
		//coloredBar := aurora.Sprintf(aurora.Green(bar))
		//buffer.WriteString(coloredBar)
		//}
		//for i := 0; i <= numberOfRedBars; i++ {
		//coloredBar := aurora.Sprintf(aurora.Red(bar))
		//buffer.WriteString(coloredBar)
		//}
	}
	return buffer.String()
}

package util

import (
	"bytes"
	"fmt"
	"math"
	"time"

	aurora "github.com/logrusorgru/aurora"
)

const (
	weekInHours = 168
)

type ElapsedTime struct {
	Hour   int
	Minute int
	Second int
}

func NewElapsedTime(d time.Duration) *ElapsedTime {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	e := ElapsedTime{
		Hour:   int(h),
		Minute: int(m),
		Second: int(s),
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
	if e.Hour > weekInHours {
		for i := 0; i < numberOfBars; i++ {
			coloredBar := aurora.Sprintf(aurora.Red(bar))
			buffer.WriteString(coloredBar)
		}
	} else if e.Hour <= weekInHours && e.Hour > 0 {
		percentage := float64(e.Hour) / float64(weekInHours)
		numberOfRedBars := int(math.Round(percentage * float64(numberOfBars)))
		numberOfGreenBars := numberOfBars - numberOfRedBars
		for i := 0; i <= numberOfGreenBars; i++ {
			coloredBar := aurora.Sprintf(aurora.Green(bar))
			buffer.WriteString(coloredBar)
		}
		for i := 0; i < numberOfRedBars; i++ {
			coloredBar := aurora.Sprintf(aurora.Red(bar))
			buffer.WriteString(coloredBar)
		}
	} else {
		for i := 0; i < numberOfBars; i++ {
			coloredBar := aurora.Sprintf(aurora.Green(bar))
			buffer.WriteString(coloredBar)
		}
	}
	return buffer.String()
}

package util

import (
	"bytes"
	"fmt"
	aurora "github.com/logrusorgru/aurora"
	"time"
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
	//numberOfBars := 10
	bar := "■"
	coloredBar := aurora.Sprintf(aurora.Cyan(bar))
	buffer.WriteString(coloredBar)
	return buffer.String()
}

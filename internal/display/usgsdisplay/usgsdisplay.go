package usgsdisplay

import (
	"bytes"
)

type UsgsDisplay struct{}

func NewUsgsDisplay() *UsgsDisplay {
	d := UsgsDisplay{}

	return &d
}

func Refresh() bool {
	return false
}

func Render() string {
	var buffer bytes.Buffer
	return buffer.String()
}

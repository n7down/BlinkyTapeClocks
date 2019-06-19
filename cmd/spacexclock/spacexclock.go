package main

import (
	//"fmt"
	//"github.com/n7down/BlinkyTapeClocks/internal/spacexapi"
	blinky "github.com/wI2L/blinkygo"
	"log"
)

// TODO: build
// 1. every 1 min - get the current time in utc
// 2. query the spacex api to get the next launch in utc:
// - curl --location --request GET "https://api.spacexdata.com/v3/launches/next" | jq '.launch_date_utc'
// 3. translate the difference of the year, month, day, hour, minute, second and millisecond? into rgb
// - where each rgb value is between 0 and 255
// - for each new row get the rgb value below it and divide it by 2
// - so if there are 4 rows then the first row is the original rgb the 2nd row is the 1st row divided by 2
// - then the 3rd row is the 2nd row divided by 2 and so on
// 4. render to the blinky tape

func main() {
	// render to blinky tape
	bt, err := blinky.NewBlinkyTape("/dev/ttyACM0", 60)
	if err != nil {
		log.Fatal(err)
	}
	defer bt.Close()

	pixel := blinky.Pixel{Color: blinky.NewRGBColor(255, 0, 0)}
	err = bt.SetPixelAt(&pixel, 0)
	if err != nil {
		log.Fatal(err)
	}
	err = bt.Render()
	if err != nil {
		log.Fatal(err)
	}

	//nextLaunch, err := spacexapi.GetNextLaunch()
	//if err != nil {
	//log.Fatal(err)
	//}

	//fmt.Printf("%s", nextLaunch.LaunchDateUtc)
}

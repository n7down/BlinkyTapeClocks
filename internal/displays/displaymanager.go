package display

import (
	"container/list"
	"fmt"
	"time"
)

// A display manager, manages a list of displays and figures out when to refresh its stats and when to render it
type DisplayManager struct {
	displayList       *list.List
	switchDisplayTime time.Duration
	startTime         time.Time
}

func NewDisplayManager(switchDisplayTime time.Duration) *DisplayManager {
	dm := DisplayManager{
		displayList:       list.New(),
		switchDisplayTime: switchDisplayTime,
		startTime:         time.Now(),
	}
	return &dm
}

func (dm DisplayManager) AddDisplay(d Display) {
	dm.displayList.PushBack(d)
	fmt.Println(fmt.Sprintf("%v", dm.displayList))
}

func (dm DisplayManager) Refresh() error {
	display := dm.displayList.Front().Value.(Display)

	err := display.Refresh()
	if err != nil {
		return err
	}
	return nil
}

func (dm DisplayManager) Render() {
	elapsedTime := time.Since(dm.startTime)
	fmt.Println(fmt.Sprintf("%v", elapsedTime))

	//if dm.displayList.Len() > 1 && elapsedTime > dm.switchDisplayTime {
	if elapsedTime > dm.switchDisplayTime {
		fmt.Println("switch display")
		fmt.Println(fmt.Sprintf("%v", dm.displayList))
		display := dm.displayList.Front()

		fmt.Println(fmt.Sprintf("%v", dm.displayList))
		dm.displayList.MoveToBack(display)
		fmt.Println(fmt.Sprintf("%v", dm.displayList))
		dm.startTime = time.Now()
	}

	display := dm.displayList.Front()
	fmt.Println(display.Value.(Display).Render())
}

package display

import (
	"fmt"
	"github.com/n7down/timelord/pkg/queue"
	"time"
)

// A display manager, manages a list of displays and figures out when to refresh its stats and when to render it
type DisplayManager struct {
	displayQueue      *queue.Queue
	switchDisplayTime time.Duration
	startTime         time.Time
}

func NewDisplayManager(switchDisplayTime time.Duration) *DisplayManager {
	dm := DisplayManager{
		displayQueue:      queue.NewQueue(),
		switchDisplayTime: switchDisplayTime,
		startTime:         time.Now(),
	}
	return &dm
}

func (dm DisplayManager) AddDisplay(d Display) {
	dm.displayQueue.Put(d)
	fmt.Println(fmt.Sprintf("%v", dm.displayQueue))
}

func (dm DisplayManager) Refresh() error {
	display, err := dm.displayQueue.Peek()
	if err != nil {
	}

	err = display.(Display).Refresh()
	if err != nil {
		return err
	}
	return nil
}

func (dm DisplayManager) Render() {
	elapsedTime := time.Since(dm.startTime)
	fmt.Println(fmt.Sprintf("%v", elapsedTime))
	if dm.displayQueue.Len() > 1 && elapsedTime > dm.switchDisplayTime {
		fmt.Println("switch display")
		fmt.Println(fmt.Sprintf("%v", dm.displayQueue))
		display, err := dm.displayQueue.Dequeue()
		if err != nil {
		}

		fmt.Println(fmt.Sprintf("%v", dm.displayQueue))
		dm.displayQueue.Put(display)
		fmt.Println(fmt.Sprintf("%v", dm.displayQueue))
		dm.startTime = time.Now()
	}
	display, err := dm.displayQueue.Peek()
	if err != nil {
	}

	fmt.Println(display.(Display).Render())
}

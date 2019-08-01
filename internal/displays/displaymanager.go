package display

import (
	"fmt"
	"github.com/golang-collections/go-datastructures/queue"
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
		displayQueue:      queue.New(2),
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
	display := dm.displayQueue.Peek().(Display)
	err := display.Refresh()
	if err != nil {
		return err
	}
	return nil
}

func (dm DisplayManager) Render() {

	// FIXME: this is not working - not switching the time
	// FIXME: may need to implement stack - https://stackoverflow.com/questions/28541609/looking-for-reasonable-stack-implementation-in-golang
	elapsedTime := time.Since(dm.startTime)
	fmt.Println(fmt.Sprintf("%v", elapsedTime))
	if dm.displayQueue.Len() > 1 && elapsedTime > dm.switchDisplayTime {
		fmt.Println("switch display")
		fmt.Println(fmt.Sprintf("%v", dm.displayStack))
		display := dm.displayStack.Pop()
		fmt.Println(fmt.Sprintf("%v", dm.displayStack))
		dm.displayStack.Push(display)
		fmt.Println(fmt.Sprintf("%v", dm.displayStack))
		dm.startTime = time.Now()
	}
	display := dm.displayStack.Peek().(Display)
	fmt.Println(display.Render())
}

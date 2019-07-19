package display

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
	"time"
)

// A display manager, manages a list of displays and figures out when to refresh its stats and when to render it
type DisplayManager struct {
	displayStack      *stack.Stack
	switchDisplayTime time.Duration
	startTime         *time.Time
}

func NewDisplayManager(switchDisplayTime time.Duration) *DisplayManager {
	startTime := time.Now()
	dm := DisplayManager{
		displayStack:      stack.New(),
		switchDisplayTime: switchDisplayTime,
		startTime:         &startTime,
	}
	return &dm
}

func (dm DisplayManager) AddDisplay(d Display) {
	dm.displayStack.Push(d)
	fmt.Println(fmt.Sprintf("%v", dm.displayStack))
}

func (dm DisplayManager) Refresh() error {
	display := dm.displayStack.Peek().(Display)
	err := display.Refresh()
	if err != nil {
		return err
	}
	return nil
}

func (dm DisplayManager) Render() {

	// FIXME: this is not working - not switching the time
	elapsedTime := time.Since(*dm.startTime)
	fmt.Println(fmt.Sprintf("%v", elapsedTime))
	if elapsedTime > dm.switchDisplayTime {
		fmt.Println("switch display")
		display := dm.displayStack.Pop()
		dm.displayStack.Push(display)
		newTime := time.Now()
		dm.startTime = &newTime
	}
	display := dm.displayStack.Peek().(Display)
	fmt.Println(display.Render())
}

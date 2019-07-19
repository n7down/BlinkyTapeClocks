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
	startTime         time.Time
}

func NewDisplayManager(switchDisplayTime time.Duration) *DisplayManager {
	dm := DisplayManager{
		displayStack:      stack.New(),
		switchDisplayTime: switchDisplayTime,
		startTime:         time.Now(),
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
	elapsedTime := time.Since(dm.startTime)
	fmt.Println(fmt.Sprintf("%v", elapsedTime))
	if elapsedTime > dm.switchDisplayTime {
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

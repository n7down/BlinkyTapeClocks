package display

import (
	"fmt"

	// FIXME: may have to change this to a stack
	"container/list"
)

// A display manager, manages a list of displays and figures out when to refresh its stats and when to render it
type DisplayManager struct {
	displayList *list.List
}

func NewDisplayManager() *DisplayManager {
	dm := DisplayManager{
		displayList: list.New(),
	}
	return &dm
}

func (dm DisplayManager) AddDisplay(d Display) {
	dm.displayList.PushBack(d)
}

func (dm DisplayManager) Render() {
	display := dm.displayList.Front().Value.(Display).Render()
	fmt.Println(display)
}

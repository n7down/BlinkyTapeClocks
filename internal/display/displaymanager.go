package display

import (
	"container/list"
	"github.com/spf13/viper"
)

// A display manager, manages a list of displays and figures out when to refresh its stats and when to render it
type DisplayManager struct {
	displayList *list.List
	config      *viper.Viper
}

func NewDisplayManager(config *viper.Viper) *DisplayManager {
	dm := DisplayManager{
		displayList: list.New(),
		config:      config,
	}
	return &dm
}

func (dm DisplayManager) AddDisplay(d Display) {
	dm.displayList.PushBack(d)
}

func (dm DisplayManager) Render() {
	dm.displayList.Front().Value.(Display).Render()
}

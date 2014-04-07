package addons

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/jonas747/fortia/gameserver/netengine"
	"github.com/robertkrimen/otto"
)

type AddonManager struct {
	Addons          []*Addon
	OttoInstance    *otto.Otto
	NetEngine       *netengine.Engine
	Log             *logrus.Logger
	Players         map[int]*Player
	BlockTypes      []BlockType
	ScriptEventChan chan Event

	idChan chan int
}

func LoadAddons(names []string, folder string) (*AddonManager, error) {
	ottoInstance := otto.New()

	addons := make([]*Addon, 0)
	for i := 0; i < len(names); i++ {
		path := names[i]
		addon, err := LoadAddon(folder + "/" + path)
		if err != nil {
			return nil, errors.New("Error loading addon " + path + ", " + err.Error())
		}
		addons = append(addons, addon)
	}

	idChan := make(chan int)
	idGen(idChan)

	return &AddonManager{
		addons,
		ottoInstance,
		nil,
		nil,
		make(map[int]*Player),
		make([]BlockType, 0),
		make(chan Event),
		idChan,
	}, nil
}

func (a *AddonManager) EnableAddon(name string) {
	// TODO re run this addon
}

func (a *AddonManager) DisableAddon(name string) {
	// TODO remove all event listeners from this addon
}

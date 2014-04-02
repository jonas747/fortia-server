package addons

import (
	"errors"
	"github.com/jonas747/fortia/gameserver/netengine"
	"github.com/robertkrimen/otto"
)

type Logger interface {
	Info(info ...string)
	Error(err ...string)
}

type AddonManager struct {
	Addons       []*Addon
	OttoInstance *otto.Otto
	NetEngine    *netengine.Engine
	Log          Logger
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
	return &AddonManager{addons, ottoInstance, nil, nil}, nil
}

// Runs all addons
func (a *AddonManager) RunAddons() {
	for i := 0; i < len(a.Addons); i++ {
		addon := a.Addons[i]
		errs := addon.Run(a.OttoInstance)
		for _, err := range errs {
			a.Log.Error(err.Error())
		}
	}
}

func (a *AddonManager) EnableAddon(name string) {
	// TODO re run this addon
}

func (a *AddonManager) DisableAddon(name string) {
	// TODO remove all event listeners from this addon
}

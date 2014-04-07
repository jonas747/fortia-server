package addons

import (
	"github.com/robertkrimen/otto"
)

// Keeps track of all the scripts event and such
func (a *AddonManager) RunScripts() {
	a.NetEngine.EmitConnOnClose = true // make sure the netengine emits the closed connections on the connclose channel
	// Run all scripts once
	for i := 0; i < len(a.Addons); i++ {
		addon := a.Addons[i]
		a.Log.Info("Running addon ", addon.Info.Name)
		errs := addon.Run(a.OttoInstance)
		for _, err := range errs {
			a.Log.Error(err.Error())
		}
		a.Log.Info("Done running addon ", addon.Info.Name)
	}

	for {
		select {
		case evt := <-a.ScriptEventChan:
			err := a.EmitEvent(evt)
			if err != nil {
				a.Log.Error("Error emiting event: ", err.Error())
			}
		}
	}
}

type Event struct {
	Name string
	Data interface{}
}

func (a *AddonManager) EmitEvent(evt Event) error {
	evtVal, _ := a.OttoInstance.ToValue(evt.Name)
	dataVal, _ := a.OttoInstance.ToValue(evt.Data)
	fortia, err := a.OttoInstance.Get("Fortia")
	if err != nil {
		return err
	}

	function, err := fortia.Object().Get("emit")
	if err != nil {
		return err
	}
	_, err = function.Call(otto.NullValue(), evtVal, dataVal)
	return err
}

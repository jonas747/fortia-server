package addons

import (
	"github.com/robertkrimen/otto"
	"strings"
)

// Keeps track of all the scripts event and such
func (a *AddonManager) RunScripts() {
	a.NetEngine.EmitConnOnClose = true // make sure the netengine emits the closed connections on the connclose channel

	// Were on the server
	a.OttoInstance.Set("_fortiaclient", false)

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
			path := ""
			switch evt.(type) {
			case GeneralEvent:
				path = "Fortia.emit"
			case NetEvent:
				path = "Fortia.Net.emit"
			default:
				a.Log.Error("Unknown event recieved in scripeventchan")
				continue
			}
			emitter, err := a.GetEmitter(path)
			if err != nil {
				a.Log.Error(err)
				continue
			}
			err = a.EmitEvent(evt, emitter)
			if err != nil {
				a.Log.Error(err)
				continue
			}
		}
	}
}

type Event interface {
	GetName() string
	GetData() []interface{}
}

type GeneralEvent struct {
	Name string
	Data interface{}
}

func (g GeneralEvent) GetName() string {
	return g.Name
}

func (g GeneralEvent) GetData() []interface{} {
	slice := []interface{}{g.Data}
	return slice
}

type NetEvent struct {
	Name   string
	Data   string
	Sender otto.Value
}

func (n NetEvent) GetName() string {
	return n.Name
}

func (n NetEvent) GetData() []interface{} {
	return []interface{}{n.Data, n.Sender}
}

func (a *AddonManager) EmitEvent(evt Event, emitter otto.Value) error {
	dataVals := make([]interface{}, 0)
	dataVals = append(dataVals, evt.GetName())
	dataVals = append(dataVals, evt.GetData()...)

	_, err := emitter.Call(otto.NullValue(), dataVals...)
	return err
}

func (a *AddonManager) GetEmitter(path string) (otto.Value, error) {
	split := strings.Split(path, ".")
	lastObject, err := a.OttoInstance.Get(split[0])
	if err != nil {
		return otto.NullValue(), err
	}
	for i := 1; i < len(split); i++ {
		p := split[i]
		lastObject, err = lastObject.Object().Get(p)
		if err != nil {
			return otto.NullValue(), err
		}
	}
	return lastObject, nil
}

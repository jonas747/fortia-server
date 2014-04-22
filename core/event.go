package core

import (
	"encoding/json"
	"github.com/idada/v8.go"
)

type Event interface {
	GetName() string        // The event name
	GetData() []interface{} // Event callback arguments
	GetObject() *v8.Object  // Optional: call on an object (for example an entity)
}

type GeneralEvent struct {
	Name string
	Data []interface{}
}

func NewGeneralEvent(name string, data ...interface{}) GeneralEvent {
	return GeneralEvent{name, data}
}

func (g GeneralEvent) GetName() string {
	return g.Name
}

func (g GeneralEvent) GetData() []interface{} {
	return g.Data
}

func (g GeneralEvent) GetObject() *v8.Object {
	return nil
}

type NetEvent struct {
	Name   string
	Data   string
	Sender int
	Obj    *v8.Object
}

func (e *Engine) NewNetEvent(name, data string, sender int) NetEvent {
	// Get the netevent object
	// TODO: Instead of getting this each time theres a new netevent, cache it somewhere?
	evt := NetEvent{
		name, data, sender, nil,
	}
	e.JsContext.Scope(func(cs v8.ContextScope) {
		global := cs.Global()
		fortiaVal := global.GetProperty("Fortia")
		fortiaObj := fortiaVal.ToObject()
		NetObj := fortiaObj.GetProperty("Net").ToObject()
		evt.Obj = NetObj
	})
	return evt
}

func (n NetEvent) GetName() string {
	return n.Name
}

func (n NetEvent) GetData() []interface{} {
	return []interface{}{n.Data, n.Sender}
}

func (n NetEvent) GetObject() *v8.Object {
	return n.Obj
}

func (e *Engine) EmitEvent(evt Event) {
	dataVals := make([]interface{}, 0)
	dataVals = append(dataVals, evt.GetName())
	dataVals = append(dataVals, evt.GetData()...)

	dataJsVals := make([]*v8.Value, 0)

	e.JsContext.Scope(func(cs v8.ContextScope) {
		for i := 0; i < len(dataVals); i++ {
			val := dataVals[i]

			// Make sure we only convert the values we need to convert
			switch t := val.(type) {
			case *v8.Value:
				dataJsVals = append(dataJsVals, t)
			default:
				jsVal := e.ToJsValue(val, cs)
				dataJsVals = append(dataJsVals, jsVal)
			}
		}

		if evt.GetObject() == nil {
			global := cs.Global()
			fortiaVal := global.GetProperty("Fortia")
			fortiaObj := fortiaVal.ToObject()
			emitVal := fortiaObj.GetProperty("emit")
			emitFunc := emitVal.ToFunction()
			emitFunc.Call(dataJsVals...)
		} else {
			obj := evt.GetObject()
			emitVal := obj.GetProperty("emit")
			emitFunc := emitVal.ToFunction()
			emitFunc.Call(dataJsVals...)
		}
	})

}

// v8's govaluetojsvalue dosent handle maps very well so this is a slow workaround
func (e *Engine) ToJsValue(value interface{}, cs v8.ContextScope) *v8.Value {
	encoded, err := json.Marshal(value)
	if err != nil {
		e.Log.Error("Error encoding value ", err)
		return nil
	}
	return cs.ParseJSON(string(encoded))
}

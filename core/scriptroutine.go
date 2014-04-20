package core

import ()

// Keeps track of all the scripts event and such
func (e *Engine) RunScripts() {

	// make sure the netengine emits the closed connections on the connclose channel
	e.NetEngine.EmitConnOnClose = true

	// Goroutine for receiving disconnected player
	// TODO: avoid race conditions
	go func() {
		for {
			disconnected := <-e.NetEngine.ConnCloseChan
			sessionData := disconnected.GetSessionData()
			id, exists := sessionData.Get("id")
			if !exists {
				e.Log.Error("Connection not authorized closed...")
				continue
			}

			idInt := id.(int)
			ply := e.Players[idInt]
			evt := NewGeneralEvent("playerleave", ply.Map())

			e.EmitEvent(evt)

			delete(e.Players, idInt)
		}
	}()

	for i := 0; i < len(e.Addons); i++ {
		addon := e.Addons[i]
		e.Log.Info("Running addon ", addon.Info.Name)
		addon.RunScripts(e.JsContext)
	}
	e.Log.Info("All addons running")
}

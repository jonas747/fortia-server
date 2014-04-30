package core

import (
	"os"
	"os/signal"
	"syscall"
)

func (e *Engine) ListenForSignal() {
	sigChan := make(chan os.Signal)

	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	sig := <-sigChan
	e.Log.Printf("Received %s signal, shutting server down... :(", sig.String())

	// Emit the shutdown event and shutdown
	evt := NewGeneralEvent("shutdown")
	e.EmitEvent(evt)
	os.Exit(1)

}

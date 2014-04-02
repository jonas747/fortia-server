package netengine

import (
	"bytes"
	"code.google.com/p/goprotobuf/proto"
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"
)

// The networking engine. Holds togheter all the connections and handlers
type Engine struct {
	registerConn   chan Connection // Channel for registering new connections
	unregisterConn chan Connection // Channel for unregistering connections
	broadcastChan  chan []byte     // Channel for broadcasting messages to all connections

	listeners   []Listener          // Slice Containing all listeners
	handlers    map[int32]Handler   // Map with all the event handlers, their id's as keys
	connections map[Connection]bool // Map containing all conncetions
}

func NewEngine() *Engine {
	return &Engine{
		registerConn:   make(chan Connection),
		unregisterConn: make(chan Connection),
		broadcastChan:  make(chan []byte),
		listeners:      make([]Listener, 0),
		handlers:       make(map[int32]Handler),
		connections:    make(map[Connection]bool),
	}
}

// Adds a listener and make it start listening for incoming connections
func (e *Engine) AddListener(listener Listener) {
	e.listeners = append(e.listeners, listener)
	listener.Listen()
}

// Handles connections
func (e *Engine) handleConn(conn Connection) {
	conn.Run()
	e.registerConn <- conn
	fmt.Println("New connection!")
	for {
		data, err := conn.Read()
		if err != nil {
			break
		}
		err = e.handleMessage(data, conn)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	fmt.Println("Disconnected from a connection :(")
	e.unregisterConn <- conn
}

// Retrieves the event id, decodes the data and calls the callback
func (e *Engine) handleMessage(data []byte, conn Connection) error {
	// Needs to atleast contain an event id
	if len(data) < 4 {
		return errors.New("Data is less than 4 bytes")
	}

	// The first 4 bytes is a signed 32 bit integer representing the event id
	evtIdRaw := data[:4]
	data = data[4:] // Resize the data slice so it dosen't include the event id
	buf := bytes.NewReader(evtIdRaw)
	var evtId int32
	err := binary.Read(buf, binary.LittleEndian, &evtId)
	if err != nil {
		// There seems to be a problem here?
		return errors.New("error: " + fmt.Sprintf("%d", evtId) + "; " + err.Error())
	}

	handler, found := e.handlers[evtId]
	if !found {
		return errors.New(fmt.Sprintf("No handler found for %d", evtId))
	}

	decoded := reflect.New(handler.DataType).Interface().(proto.Message) // We use reflect to unmarshal the data into the appropiate type
	err = proto.Unmarshal(data, decoded)
	if err != nil {
		return err
	}
	// ready the function
	funcVal := reflect.ValueOf(handler.CallBack)
	decVal := reflect.Indirect(reflect.ValueOf(decoded)) // decoded is a pointer, so we get the value it points to
	sessionDVal := reflect.ValueOf(conn.GetSessionData())
	resp := funcVal.Call([]reflect.Value{decVal, sessionDVal}) // Call it

	if len(resp) == 0 {
		return nil
	}

	returnVal := resp[0]
	if returnVal.Kind() == reflect.Slice {
		inter := returnVal.Interface()
		responseRaw := inter.([]byte)
		conn.Send(responseRaw)
	}

	return nil
}

// Adds a handler
func (e *Engine) AddHandler(handler Handler) {
	e.handlers[handler.Event] = handler
}

func (e *Engine) ListenChannels() {
	for {
		select {
		case d := <-e.registerConn: //Register a connection
			e.connections[d] = true
		case d := <-e.unregisterConn: //Unregister a connection
			delete(e.connections, d)
		case msg := <-e.broadcastChan: //Broadcast a message to all connections
			for conn := range e.connections {
				conn.Send(msg)
			}
		}
	}
}

func (e *Engine) NumClients() int {
	return len(e.connections)
}

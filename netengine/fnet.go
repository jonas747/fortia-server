package netengine

import (
	"bytes"
	"code.google.com/p/go.net/websocket"
	"code.google.com/p/goprotobuf/proto"
	"encoding/binary"
	"errors"
	"net/http"
	"reflect"
)

// Struct which represents an event
type Event struct {
	Id   int32
	Data reflect.Value
}

// Struct which represents a event handler
// TODO: example
type Handler struct {
	CallBack interface{}
	Event    int32
	DataType reflect.Type
}

func NewHandler(callback interface{}, evt int32) (Handler, error) {
	err := validateCallback(callback)
	if err != nil {
		return Handler{}, err
	}

	dType := reflect.TypeOf(callback).In(0) // Get the type of the first parameter for later use

	return Handler{
		CallBack: callback,
		Event:    evt,
		DataType: dType,
	}, nil
}

func validateCallback(callback interface{}) error {
	t := reflect.TypeOf(callback)
	if t.Kind() != reflect.Func {
		return errors.New("Callback not a function")
	}
	if t.NumIn() != 2 {
		return errors.New("Callback does not have 2 parameters")
	}
	return nil
}

// Listener is a interface for listening for incoming connections
type Listener interface {
	Listen() // Listens for incoming connections
}

type WebsocketListener struct {
	Engine *Engine
	Addr   string
}

// Implements Listener.Listen()
func (w *WebsocketListener) Listen() {
	handler := func(ws *websocket.Conn) {
		conn := NewWebsocketConn(ws)
		w.Engine.handleConn(conn)
	}

	server := websocket.Server{Handler: handler}

	http.Handle("/", server)

	err := http.ListenAndServe(w.Addr, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func EncodeMessage(msg proto.Message, evtId int32) ([]byte, error) {
	// Encode the protobuf message itself
	encoded, err := proto.Marshal(msg)
	if err != nil {
		return make([]byte, 0), err
	}

	// Create a new buffer, stuff the event id and the encoded message in it
	buffer := new(bytes.Buffer)
	err = binary.Write(buffer, binary.LittleEndian, evtId)
	if err != nil {
		return make([]byte, 0), err
	}

	_, err = buffer.Write(encoded)
	if err != nil {
		return make([]byte, 0), err
	}

	unread := buffer.Bytes()
	return unread, nil
}

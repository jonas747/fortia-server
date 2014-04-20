package netengine

import (
	"code.google.com/p/go.net/websocket"
	"errors"
	"time"
)

type Connection interface {
	Send(data []byte) error        // Sends data
	Read() ([]byte, error)         // Blocks untill data is received, then returns []byte
	Kind() string                  // What kind of connection is it (websocket, tcp etc..)
	Close()                        // Closes the connections ands stops all goroutines associated with it
	GetSessionData() *SessionStore // Gets the session datastore
	Run()                          // Starts the writer and reader goroutines
	Open() bool                    // Wether this connection is open ot not
}

type SessionStore struct {
	data map[string]interface{}
}

func (s *SessionStore) Set(key string, val interface{}) {
	s.data[key] = val
}

func (s *SessionStore) Get(key string) (val interface{}, exists bool) {
	value, exists := s.data[key]
	return value, exists
}

type WebsocketConn struct {
	sessionStore *SessionStore
	conn         *websocket.Conn

	readChan    chan []byte
	readErrChan chan error
	writeChan   chan []byte

	stopWriting chan bool
	isOpen      bool
}

func NewWebsocketConn(c *websocket.Conn) Connection {
	store := &SessionStore{make(map[string]interface{})}
	conn := WebsocketConn{
		sessionStore: store,
		conn:         c,
		readChan:     make(chan []byte),
		writeChan:    make(chan []byte),
		stopWriting:  make(chan bool),
		readErrChan:  make(chan error),
		isOpen:       true,
	}
	return &conn
}

// Implements Connection.Send([]byte)
func (w *WebsocketConn) Send(b []byte) error {
	if !w.isOpen {
		return errors.New("Cannot call WebsocketConn.Send() on a closed connection")
	}
	after := time.After(time.Duration(5) * time.Second) // Time out
	select {
	case w.writeChan <- b:
		return nil
	case <-after:
		w.isOpen = false
		w.Close()
		return errors.New("Timed out sending payload to writechan")
	}
}

// Implements Connection.Read()([]byte, error)
func (w *WebsocketConn) Read() ([]byte, error) {
	select {
	case data := <-w.readChan:
		return data, nil
	case err := <-w.readErrChan:
		return make([]byte, 0), err
	}
}

// Implements Connection.Kind() string
func (w *WebsocketConn) Kind() string {
	return "websocket"
}

// Implements Connection.Close()
func (w *WebsocketConn) Close() {
	w.isOpen = false
	w.conn.Close()
}

func (w *WebsocketConn) Open() bool {
	return w.isOpen
}

func (w *WebsocketConn) GetSessionData() *SessionStore {
	return w.sessionStore
}

// Implements Connection.Run()
func (w *WebsocketConn) Run() {
	// Launch the write and read goroutines
	go w.writer()
	go w.reader()
}

// Writes messages from WebsocketConn.writeChan, Which is used by WebsocketConn.Write([]byte)
func (w *WebsocketConn) writer() {
	for {
		select {
		case m := <-w.writeChan:
			err := websocket.Message.Send(w.conn, m)
			if err != nil {
				break
			}
		case <-w.stopWriting:
			return
		}
	}
}

// Reads messages from the websocket connection into WebsocketConn.readChan, which is used by WebcosketConn.Read()([]byte, error)
func (w *WebsocketConn) reader() {
	for {
		var msg []byte
		err := websocket.Message.Receive(w.conn, &msg)
		if err != nil {
			w.Close()
			go func() {
				after := time.After(time.Duration(5) * time.Second)
				select {
				case w.stopWriting <- true:
					return
				case <-after:
					return
				}
			}()
			w.readErrChan <- err
			w.isOpen = false
			return
		}
		w.readChan <- msg
	}
}

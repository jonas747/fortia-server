package core

import (
	"errors"
	"github.com/Sirupsen/logrus"
	v8 "github.com/idada/v8.go"
	"github.com/jonas747/fortia-server/netengine"
	"sync"
)

var log *logrus.Logger

type Engine struct {
	Addons      []*Addon
	JsEngine    *v8.Engine
	JsContext   *v8.Context
	NetEngine   *netengine.Engine
	Log         *logrus.Logger
	Players     map[int]*Player
	PlayersLock *sync.Mutex
	BlockTypes  []BlockType

	currentlyLoading *Addon

	idChan chan int
}

func NewEngine(logger *logrus.Logger) *Engine {
	log = logger
	jEngine := v8.NewEngine()
	e := &Engine{
		Addons:      make([]*Addon, 0),
		JsEngine:    jEngine,
		Log:         logger,
		Players:     make(map[int]*Player),
		PlayersLock: new(sync.Mutex),
		BlockTypes:  make([]BlockType, 0),
	}
	return e
}

// Loads all addons, and starts the netengine
// TODO: Be able to stop player id generation
func (e *Engine) Start(listenAddr string, addons []string, folder string) {

	// Initialize and start the networking engine
	e.Log.Debug("Initializing networking engine")
	net := netengine.NewEngine()
	go net.ListenChannels()
	e.NetEngine = net

	e.AddNetHandlers()

	listener := netengine.WebsocketListener{
		net,
		listenAddr,
	}
	go net.AddListener(&listener)

	e.Log.Debug("Running player id generation goroutine")
	// The player id channel
	idChan := make(chan int)
	idGen(idChan)
	e.idChan = idChan

	e.Log.Debug("Adding javascript extensions")
	// Adds the javascript extensions, sets e.JsContext
	e.AddJsExtensions()

	e.Log.Debug("Loading, compiling and running addons")
	// Load all the addons, compile the scripts etc..
	err := e.LoadAddons(addons, folder)
	if err != nil {
		e.Log.Error("Error loading addons: ", err)
		return
	}
	go e.ListenForClosedConnection()
}

// Loads and compiles addons into memory
func (e *Engine) LoadAddons(names []string, folder string) error {
	addons := make([]*Addon, 0)

	for i := 0; i < len(names); i++ {
		name := names[i]
		e.Log.Debug("Loading addon ", name)

		// Load the addon
		addon, err := e.LoadAddon(folder + "/" + name)
		if err != nil {
			return errors.New("Error loading addon " + names[i] + ", " + err.Error())
		}

		addons = append(addons, addon)
	}
	e.Addons = addons
	return nil
}

func (a *Engine) EnableAddon(name string) {
	// TODO re run this addon
}

func (a *Engine) DisableAddon(name string) {
	// TODO remove all event listeners from this addon
}

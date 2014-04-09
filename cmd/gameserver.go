package main

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/jonas747/fortia/gameserver/addons"
	"github.com/jonas747/fortia/gameserver/netengine"
	"io/ioutil"
	"runtime"
	"time"
)

var (
	config Config
	log    *logrus.Logger
)

type Config struct {
	ConVars map[string]string
	Addons  []string
	Debug   bool
}

func LoadConfig() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic("Error reading configuration file\n" + err.Error())
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		panic("Error decoding configuration file\n" + err.Error())
	}
}

func main() {
	log = logrus.New()
	// Set appropiate loglevel
	log.Level = logrus.Info
	log.Formatter = new(TextFormatter)

	log.Info("Starting fortia gamsesever")
	LoadConfig()
	if config.Debug {
		log.Level = logrus.Debug
	}
	log.Debug("Debug level enabled")

	log.Debug("Loading addons fortia gamsesever")
	am, err := addons.LoadAddons(config.Addons, "addons")
	if err != nil {
		panic(err)
	}
	am.Log = log

	log.Debug("Creating networking engine and starting it")
	engine := netengine.NewEngine()
	go engine.ListenChannels()

	am.NetEngine = engine

	log.Debug("Adding netengine handlers and javascript functions")
	am.AddHandlers()
	am.AddExtensions()

	log.Debug("Adding websocket listener")
	listener := netengine.WebsocketListener{
		engine,
		":8093",
	}
	go engine.AddListener(&listener)

	if config.Debug {
		go func() {
			for {
				after := time.After(time.Duration(10) * time.Second)
				log.Debugf("Number of goroutines: %d", runtime.NumGoroutine())
				<-after
			}
		}()
	}

	log.Info("Server Sucessfully started")
	log.Debug("Final step running script go routine")
	am.RunScripts()
}

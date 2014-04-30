package main

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/jonas747/fortia-server/core"
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
	runtime.GOMAXPROCS(runtime.NumCPU())
	log = logrus.New()
	// Set appropiate loglevel
	log.Level = logrus.Info
	log.Formatter = new(TextFormatter)

	log.Info("Starting fortia server")
	LoadConfig()
	if config.Debug {
		log.Level = logrus.Debug
		log.Debug("Debug level enabled")
	}

	engine := core.NewEngine(log)

	engine.Start(":"+config.ConVars["port"], config.Addons, "addons")
	log.Info("Server Sucessfully started on " + config.ConVars["port"])

	if config.Debug {
		func() {
			for {
				after := time.After(time.Duration(10) * time.Second)
				log.Debugf("Number of goroutines: %d Online players: %d", runtime.NumGoroutine(), len(engine.Players))
				<-after
			}
		}()
	}
}

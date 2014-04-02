package main

import (
	"encoding/json"
	"github.com/jonas747/fortia/gameserver/addons"
	"io/ioutil"
	"log"
)

var config Config

type Config struct {
	ConVars map[string]string
	Addons  []string
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
	LoadConfig()
	log.Println(config)
	am, err := addons.LoadAddons(config.Addons, "addons")
	if err != nil {
		panic(err)
	}
	am.Log = new(Logger)
	am.RunAddons()
}

type Logger struct{}

func (l Logger) Info(info ...string) {
	log.Println(info)
}

func (l Logger) Error(err ...string) {
	log.Println("ERROR:", err)
}

package core

import (
	"encoding/json"
	v8 "github.com/idada/v8.go"
	"io/ioutil"
	"os"
)

// Struct representing the addon.json file
type AddonInfo struct {
	Name    string
	Version string
	License string
	Type    string
	Author  string
}

// Struct representing an addon
type Addon struct {
	Info          AddonInfo
	ClientScripts []*Script
	ServerScripts []*Script
	SharedScripts []*Script
	Resources     []*Resource
}

func LoadAddon(path string) (*Addon, error) {

	file, err := ioutil.ReadFile(path + "/addon.json")
	if err != nil {
		return nil, err
	}

	var info AddonInfo
	err = json.Unmarshal(file, &info)
	if err != nil {
		return nil, err
	}

	cdirExists, _ := exists(path + "/scripts/client")
	sdirExists, _ := exists(path + "/scripts/server")
	shdirExists, _ := exists(path + "/scripts/shared")

	clientScripts := make([]*Script, 0)
	if cdirExists {
		clientScripts, err = LoadScripts(path+"/scripts/client", SCRIPT_CLIENT)
		if err != nil {
			return nil, err
		}
	}

	serverScripts := make([]*Script, 0)
	if sdirExists {
		serverScripts, err = LoadScripts(path+"/scripts/server", SCRIPT_SERVER)
		if err != nil {
			return nil, err
		}
	}

	sharedScripts := make([]*Script, 0)
	if shdirExists {
		sharedScripts, err = LoadScripts(path+"/scripts/shared", SCRIPT_SHARED)
		if err != nil {
			return nil, err
		}
	}

	resources, err := LoadResources(path + "/resources")
	if err != nil {
		return nil, err
	}

	return &Addon{
		Info:          info,
		ClientScripts: clientScripts,
		ServerScripts: serverScripts,
		SharedScripts: sharedScripts,
		Resources:     resources,
	}, nil
}

// Compiles, and runs the shared and server scripts on the server
func (a *Addon) RunScripts(ctx *v8.Context) {
	for _, sharedScript := range a.SharedScripts {
		sharedScript.Run(ctx)
	}
	for _, serverScript := range a.ServerScripts {
		serverScript.Run(ctx)
	}
}

// Compiles all server scripts
func (a *Addon) CompileScripts(js *v8.Engine) {
	for _, sharedScript := range a.SharedScripts {
		sharedScript.Compile(js)
	}
	for _, serverScript := range a.ServerScripts {
		serverScript.Compile(js)
	}
}

// Scans a directory recursively returning all paths
func scanDir(path string) ([]string, error) {
	dirInfo, err := ioutil.ReadDir(path)
	if err != nil {
		return make([]string, 0), err
	}

	out := make([]string, 0)
	for _, v := range dirInfo {
		completePath := path + "/" + v.Name()
		if string(path[len(path)-1]) == "/" {
			completePath = path + v.Name()
		}
		if v.IsDir() {
			nOut, err := scanDir(completePath)
			if err != nil {
				return make([]string, 0), err
			}
			out = append(out, nOut...)
		} else {
			out = append(out, completePath)
		}
	}

	return out, nil
}

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

package core

import (
	"encoding/json"
	//"github.com/idada/v8.go"
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
	Resources     []*Resource
	Path          string
}

func (e *Engine) LoadAddon(path string) (*Addon, error) {

	file, err := ioutil.ReadFile(path + "/addon.json")
	if err != nil {
		return nil, err
	}

	var info AddonInfo
	err = json.Unmarshal(file, &info)
	if err != nil {
		return nil, err
	}

	initScript, err := LoadScript(path+"/scripts/init.js", SCRIPT_SERVER)
	if err != nil {
		return nil, err
	}

	resources, err := LoadResources(path + "/resources")
	if err != nil {
		return nil, err
	}

	addon := &Addon{
		Info:      info,
		Resources: resources,
		Path:      path,
	}

	e.currentlyLoading = addon
	//compile and Execute the init script
	initScript.Compile(e.JsEngine)
	initScript.Run(e.JsContext)

	return addon, nil
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

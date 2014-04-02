package addons

import (
	"encoding/json"
	"io/ioutil"
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

//Scans a directory recursively returning all paths
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

	clientScripts, err := LoadScripts(path+"/scripts/client", SCRIPT_CLIENT)
	if err != nil {
		return nil, err
	}

	serverScripts, err := LoadScripts(path+"/scripts/server", SCRIPT_SERVER)
	if err != nil {
		return nil, err
	}

	sharedScripts, err := LoadScripts(path+"/scripts/shared", SCRIPT_SHARED)
	if err != nil {
		return nil, err
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

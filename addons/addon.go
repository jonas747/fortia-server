package addons

import (
	//"encoding/json"
	"io/ioutil"
)

type Addon struct {
	Path          string
	Name          string
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

	// Get the locations of all the files
	/*
		clientScriptPaths, err := scanDir(path + "/scripts/client")
		if err != nil {
			return nil, err
		}

		serverScriptPaths, err := scanDir(path + "/scripts/server")
		if err != nil {
			return nil, err
		}

		sharedScriptPaths, err := scanDir(path + "/scripts/shared")
		if err != nil {
			return nil, err
		}

		resourcePaths, err := scanDir(path + "/resources")
		if err != nil {
			return nil, err
		}
	*/

	return nil, nil
}

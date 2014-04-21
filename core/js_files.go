/*
Javascript file library

Functions:
    Append - Appends data to the end of the file
    Delete - Deletes a file
	ReadDir - Returns all the contents of a directory
    IsDir - Returns wether path is a directory or not
    Size - Returns the size of a file
    Time - Returns the time a file was modified

Currently implemented:
    CreateDir 	- Creates a directory
    Write		- Writes to a file, erasing previous contents and creating it if it dosent exist
    Read		- Reads the contends of a file
    Exists		- Returns wtherer this file exsists or not

TODO:
Throw exceptions somehow if things go wrong (dont have permission to view a file and so on) instead of just logging it
Be able to write with different permission bits than 770
*/

package core

import (
	"github.com/idada/v8.go"
	"io/ioutil"
	"os"
)

// Returns wether a file exsits or not
func jsFileExists(path string) bool {
	exists, err := fileExists(path)
	if err != nil {
		log.Error("Error checking if [", path, "] exists: ", err)
		return false
	}
	return exists
}

// Returns the contents of a file
func jsFileRead(path string) string {
	exists, err := fileExists(path)
	if err != nil {
		log.Error("Error reading [", path, "]: ", err)
		return ""
	}

	if !exists {
		log.Error("Error readin [", path, "]: File does not exist")
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error("Error reading [", path, "]: ", err)
		return ""
	}

	return string(file)
}

// writes data to path
func jsFileWrite(path, data string) {
	err := ioutil.WriteFile(path, []byte(data), os.FileMode(0770))
	if err != nil {
		log.Error("Error writing [", path, "] ", err)
	}
}

// creates directory path
func jsfileCreateDir(path string) {
	err := os.MkdirAll(path, os.FileMode(0770|os.ModeDir))
	if err != nil {
		log.Error("Error creating direcotiry(ies) [", path, "]: ", err)
	}
}

// Adds the functions to the global object
func jsFileApi(template *v8.ObjectTemplate) {
	template.Bind("fileExists", jsFileExists)
	template.Bind("fileRead", jsFileRead)
	template.Bind("fileWrite", jsFileWrite)
	template.Bind("fileCreateDir", jsfileCreateDir)
}

// exists returns whether the given file or directory exists or not
func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

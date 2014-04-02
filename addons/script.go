package addons

import (
	"github.com/robertkrimen/otto"
	"io/ioutil"
)

type ScriptType uint8

const (
	SCRIPT_SERVER ScriptType = iota
	SCRIPT_CLIENT
	SCRIPT_SHARED
)

type Script struct {
	Source   []byte
	Path     string
	Type     ScriptType
	OttoInst *otto.Otto
}

func LoadScript(path string, kind ScriptType) (*Script, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &Script{
		Source:   contents,
		Path:     path,
		Type:     kind,
		OttoInst: nil,
	}, nil
}

func (s *Script) Run(ottoInstance *otto.Otto) error {
	s.OttoInst = ottoInstance
	_, err := ottoInstance.Run(string(s.Source))
	return err
}

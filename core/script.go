package core

import (
	v8 "github.com/idada/v8.go"
	"io/ioutil"
)

type ScriptType uint8

const (
	SCRIPT_SERVER ScriptType = iota
	SCRIPT_CLIENT
	SCRIPT_SHARED
)

type Script struct {
	Source     []byte
	Path       string
	Type       ScriptType
	V8Script   *v8.Script
	ShouldExec bool
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
		V8Script: nil,
	}, nil
}

func LoadScripts(path string, kind ScriptType) ([]*Script, error) {
	paths, err := scanDir(path)
	if err != nil {
		return nil, err
	}

	scripts := make([]*Script, 0)
	for _, v := range paths {
		script, err := LoadScript(v, kind)
		if err != nil {
			return scripts, err
		}
		scripts = append(scripts, script)
	}
	return scripts, nil
}

func (s *Script) Compile(js *v8.Engine) {
	s.V8Script = js.Compile(s.Source, js.NewScriptOrigin(s.Path, 0, 0))
}

// Executes the script
func (s *Script) Run(ctx *v8.Context) {
	if s.V8Script == nil {
		log.Error("Tried running script that wasnt compiled")
		return
	}
	ctx.Scope(func(callback v8.ContextScope) {
		callback.Run(s.V8Script)
	})
}

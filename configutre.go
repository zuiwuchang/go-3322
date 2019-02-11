package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	jsonnet "github.com/google/go-jsonnet"
)

// Configure 配置 檔案
type Configure struct {
	Name  string
	Host  string
	A     string
	Timer string
}

var _Configure Configure

func initConfigure() (*Configure, error) {
	path, e := exec.LookPath(os.Args[0])
	if e != nil {
		return nil, e
	}
	path, e = filepath.Abs(path)
	if e != nil {
		return nil, e
	}
	path = filepath.Dir(path) + "/" + ConfigureFile

	b, e := ioutil.ReadFile(path)
	if e != nil {
		return nil, e
	}

	vm := jsonnet.MakeVM()
	jsonStr, e := vm.EvaluateSnippet("", string(b))
	if e != nil {
		return nil, e
	}
	b = []byte(jsonStr)

	cnf := getConfigure()
	e = json.Unmarshal(b, cnf)
	if e != nil {
		return nil, e
	}

	return cnf, nil
}
func getConfigure() *Configure {
	return &_Configure
}

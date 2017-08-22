package main

import (
	"encoding/json"
	"io/ioutil"
)

type Configure struct {
	Name  string
	Host  string
	A     string
	Timer string
}

var g_Configure Configure

func initConfigure() (*Configure, error) {
	b, e := ioutil.ReadFile(ConfigureFile)
	if e != nil {
		return nil, e
	}

	cnf := getConfigure()
	e = json.Unmarshal(b, cnf)
	if e != nil {
		return nil, e
	}

	return cnf, nil
}
func getConfigure() *Configure {
	return &g_Configure
}

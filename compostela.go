package main

import (
	"container/list"
	"encoding/json"
	"io/ioutil"
)

type compostelaState struct {
	pilgrimEncrypted bool
}

type compostelaData struct {
	Pilgrim []byte
	Stamps  list.List
}

type Compostela struct {
	Data  compostelaData
	State compostelaState
}

/*
Obtiene un objeto Compostela desde un archivo.
*/
func GetCompostelaFromFile(path string) (Compostela, error) {
	comp := Compostela{}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return comp, err
	}

	comp.State.pilgrimEncrypted = true

	err = json.Unmarshal(file, &comp.Data)
	if err != nil {
		return comp, err
	}

	return comp, nil
}

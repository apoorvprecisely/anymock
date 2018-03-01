package conf

import (
	"encoding/json"
	"os"
)

const confFile = "mock-api.json"

type Configuration struct {
	Object ObjectType
}

var Conf Configuration

type ObjectType struct {
	ApiData []ApiData
	Port    string
}

type ApiData struct {
	Url      string
	Response string
}

func Load() {
	file, err := os.Open(confFile)
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	Conf = Configuration{}
	err = decoder.Decode(&Conf)
	if err != nil {
		panic(err)
	}
}

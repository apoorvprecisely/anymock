package main

import (
	"net/http"
	"fmt"
	"os"
	"encoding/json"
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

func main() {
	Load()
	for _, data := range Conf.Object.ApiData {
		http.HandleFunc(data.Url, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, data.Response)
		})
	}
	http.ListenAndServe(Conf.Object.Port, nil)
}
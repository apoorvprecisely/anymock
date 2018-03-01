package main

import (
	"net/http"
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
func handle(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(format))
	}
	return http.HandlerFunc(fn)
}
func main() {
	Load()
	for _, data := range Conf.Object.ApiData {
		http.Handle(data.Url, handle(data.Response))
	}
	http.ListenAndServe(Conf.Object.Port, nil)
}
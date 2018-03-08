package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

const defaultConfFile = "mock-api.json"

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

func Load(fileName *string) {
	file, err := os.Open(*fileName)
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
func myUsage() {
	fmt.Printf("Usage: %s [OPTIONS] argument ...\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Print("{\n" +
		"  \"object\": {\n" +
		"    \"apiData\": [\n" +
		"      {\n" +
		"        \"url\": \"/abc/xyz\",\n" +
		"        \"response\": \"Success\"\n" +
		"      }\n" +
		"    ],\n" +
		"    \"port\": \"1234\"\n" +
		"  }\n" +
		"}")
}
func main() {
	file := flag.String("config", "mock-api.json", "Config File to pick up requests and responses")
	flag.Usage = myUsage
	flag.Parse()
	Load(file)
	for _, data := range Conf.Object.ApiData {
		http.Handle(data.Url, handle(data.Response))
	}
	http.ListenAndServe(Conf.Object.Port, nil)
}
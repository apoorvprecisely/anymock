package main

import (
	"conf"
	"net/http"
	"fmt"
)

func main() {
	conf.Load()
	for _, data := range conf.Conf.Object.ApiData {
		http.HandleFunc(data.Url, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, data.Response)
		})
	}
	http.ListenAndServe(conf.Conf.Object.Port, nil)
}
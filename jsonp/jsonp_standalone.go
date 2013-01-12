// +build !appengine

package main

import (
	"flag"
	"net/http"
)

var client *http.Client
var port *string

func getHttpClient(r *http.Request) *http.Client {
	if client == nil {
		client = &http.Client{}
	}
	return client
}

func main() {
	// if you want to run it standalone, you have the option
	port = flag.String("port", ":8989", "port to run on")
	flag.Parse()
	err := http.ListenAndServe(*port, nil)
	panic(err)
}

package jsonp

import (
	"appengine"
	"appengine/datastore"
	"appengine/urlfetch"
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/jsonp/", jsonpHandler)
	http.Handle("/", http.FileServer(http.Dir("web/")))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello static")
}

func jsonpHandler(w http.ResponseWriter, r *http.Request) {
	r.Header["Content-Type"] = ["text/javascript"]
	
}

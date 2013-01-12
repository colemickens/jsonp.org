package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func init() {
	http.HandleFunc("/api", jsonpHandler)
	http.HandleFunc("/testdoc.json", testDocHandler)
	http.Handle("/", http.FileServer(http.Dir("www/")))
}

func testDocHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{ "test": "test", "test2": "test2" }`))
}

func jsonpHandler(w http.ResponseWriter, r *http.Request) {
	client := getHttpClient(r)

	_method := r.URL.Query().Get("method")
	_url := r.URL.Query().Get("url")
	_cb := r.URL.Query().Get("cb")

	if _cb == "" {
		_cb = "callback"
	}

	goodUrl, err := url.Parse(_url)
	if err != nil {
		w.WriteHeader(504)
		w.Write([]byte("No url specified." + err.Error()))
		return
	}

	resp, err := client.Get(goodUrl.String())
	if err != nil {
		w.WriteHeader(504)
		w.Write([]byte("Couldn't retrieve url." + err.Error()))
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	// check to ensure it's valid json. probably not the most efficient method
	// also, this is (correctly, but unnecessarily) picky
	junk := make(map[string]interface{})
	err = json.Unmarshal(body, &junk)
	if err != nil {
		w.WriteHeader(504)
		w.Write([]byte("Url returned bad JSON." + err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/javascript")

	if _method == "rewrap" {
		w.WriteHeader(501) // not implemented yet
	} else if _method == "proxy" || _method == "" {
		doc := _cb + "(" + string(body) + ")"
		w.Write([]byte(doc))
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

//тестирование http сервера

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}

	name := query.Get("name")
	if len(name) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You must supply a name"))
		return
	}
	w.WriteHeader(http.StatusOK)
	res := fmt.Sprintf("Hello %s", name)
	w.Write([]byte(res))
}

func main() {
	http.HandleFunc("/greet", RequestHandler)
	log.Fatal(http.ListenAndServe(":3030", nil))
}

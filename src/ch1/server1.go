package main

import (
	"fmt"
	"log"
	"net/http"

)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/lissajous", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	//print header
	for key, value := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", key, value)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	for key, value := range r.Form {
		fmt.Fprintf(w, "From[%q] = %q\n", key, value)
	}
}








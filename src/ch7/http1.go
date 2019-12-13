package main

import "net/http"

import "fmt"

import "log"

type dollars float64

type database map[string]dollars

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

func (db database) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	switch request.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(response, "%s: $%s\n", item, price)
		}
		break
	case "/price":
		item := request.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			response.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(response, "no such item : %s\n", item)
			return
		}
		fmt.Fprintf(response, "%s: %s\n", item, price)
		break
	default:
		response.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(response, "no such page\n")
	}

}

func (db database) list(response http.ResponseWriter, request *http.Request) {
	for item, price := range db {
		fmt.Fprintf(response, "%s: $%s\n", item, price)
	}
}

func (db database) price(response http.ResponseWriter, request *http.Request) {
	item := request.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		response.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(response, "no such item : %s\n", item)
		return
	}
	fmt.Fprintf(response, "%s: %s\n", item, price)
}

func TestHttp1() {
	db := database{
		"nike shoe": 66.9,
		"cs book":   19.9,
		"ice":       9.9,
	}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

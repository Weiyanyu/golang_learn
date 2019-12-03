package main

import "encoding/json"

import "log"

import "fmt"

type Moive struct {
	Title  string
	Year   int  `json:"release"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func MovieJson() {
	var moive = []Moive{
		{Title: "zhanlang", Year: 2019, Color: false, Actors: []string{"wujing"}},
		{Title: "haizeiwang", Year: 2019, Color: true, Actors: []string{"lufei", "suolong"}},
	}

	data, err := json.MarshalIndent(moive, "", "	")
	if err != nil {
		log.Fatalf("Json marshal error %s\n", err)
	}
	fmt.Printf("%s\n", data)

	var titles []struct{ Title string }
	json.Unmarshal(data, &titles)
	fmt.Println(titles)
}

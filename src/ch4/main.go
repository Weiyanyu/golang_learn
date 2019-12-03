package main

import "os"

import "log"

import "fmt"

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues ", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%‐5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

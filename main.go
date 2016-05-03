// +build !appengine
// Above is a special build command: https://blog.golang.org/the-app-engine-sdk-and-workspaces-gopath

package main

import (
	"net/http"
	"log"
)

func main() {

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	print(err)
	print("All good! go check localhost:8080")
}
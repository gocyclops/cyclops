package main

import (
	"log"

	"github.com/gocyclops/cyclops/cmd/create"
)

func main() {
	if err := create.Execute(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"

	"events-fork/internals/server"
)

func main() {
	app := server.New()

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

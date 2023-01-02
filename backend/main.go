package main

import (
	"events-fork/internals/server"
	_ "events-fork/migrations"
	"log"
)

func main() {
	app := server.New()

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"

	"github.com/holmanskih/operations-research/internal/ui"
)

func main() {
	app, err := ui.NewApp()
	if err != nil {
		log.Fatalf("gfailed to start the app: %s", err)
	}
	app.LoadUI("Dichotomy")
}

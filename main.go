package main

import (
	"log"

	"github.com/holmanskih/goplotter/ui"
)

const (
	appWindowName = "GoPlotter"
)

func main() {
	app, err := ui.NewApp()
	if err != nil {
		log.Fatalf("failed to start the app: %s", err)
	}
	app.LoadUI(appWindowName)
}

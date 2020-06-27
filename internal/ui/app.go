package ui

import (
	"fmt"
	"log"
	"math"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"

	"github.com/holmanskih/operations-research/internal/plotter"
)

const (
	plotName        = "Dichotomy"
	plotExportImage = "dichotomy.png"
)

func f(x float64) float64 {
	return math.Pow(math.Sin(x), 2) / (2 + math.Sin(x))
}

// Represents the UI fyne.io application struct
type App struct {
	F func(x float64) float64

	Plotter *plotter.Plotter

	ui fyne.App

	image  *canvas.Image
	widget fyne.Widget
	window fyne.Window
}

func NewApp() (*App, error) {
	a := &App{}

	p, err := plotter.NewPlotter(plotName)
	if err != nil {
		return nil, fmt.Errorf("failed to create new Plotter: %s", err)
	}
	p.AddFunc(f, "f(x)")
	a.Plotter = p

	a.ui = app.New()
	a.ui.Settings().SetTheme(theme.LightTheme())

	a.image = &canvas.Image{FillMode: canvas.ImageFillOriginal}

	// widget init
	xMinEntry := widget.NewEntry()
	xMinEntry.SetPlaceHolder("xMin")

	xMaxEntry := widget.NewEntry()
	xMaxEntry.SetPlaceHolder("xMax")

	preciousEntry := widget.NewEntry()
	preciousEntry.SetPlaceHolder("Precious")
	a.widget = &widget.Form{
		Items: []*widget.FormItem{
			{Text: "xMin", Widget: xMinEntry},
			{Text: "xMax", Widget: xMaxEntry},
			{Text: "Precious", Widget: preciousEntry},
		},
		OnSubmit: func() {
			// plot the function
			// save function to a file
			err = a.Plotter.Save(plotExportImage)
			if err != nil {
				log.Println("failed to save plot image")
			}

			// read the file by fyne
			a.setImage(plotExportImage)
		},
	}
	return a, nil
}

func (a *App) LoadUI(windowName string) {
	a.window = a.ui.NewWindow(windowName)
	a.window.SetContent(fyne.NewContainerWithLayout(
		layout.NewGridLayout(2),
		fyne.NewContainerWithLayout(layout.NewGridLayout(3),
			a.widget,
		),
		fyne.NewContainerWithLayout(layout.NewGridLayout(1),
			a.image,
		),
	))
	a.window.ShowAndRun()
}

func (a *App) setImage(path string) {
	a.image.File = path
	canvas.Refresh(a.image)
}

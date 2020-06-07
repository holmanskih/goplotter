package ui

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// Represents the UI fyne.io application struct
type App struct {
	ui fyne.App

	image *canvas.Image

	//output *widget.Label

	buttons map[string]*widget.Button
	window  fyne.Window
}

func NewApp() *App {
	a := &App{}
	a.ui = app.New()
	a.ui.Settings().SetTheme(theme.LightTheme())

	a.image = &canvas.Image{FillMode: canvas.ImageFillOriginal}
	//a.functions = make(map[string]func())
	a.buttons = make(map[string]*widget.Button)
	return a
}

//func (a *App) addButton(text string, action func()) *widget.Button {
//	button := widget.NewButton(text, action)
//	a.buttons[text] = button
//	return button
//}

func (a *App) LoadUI(windowName string) {
	a.window = a.ui.NewWindow(windowName)

	a.setImage("f1.jpg")

	a.window.SetContent(fyne.NewContainerWithLayout(
		layout.NewGridLayout(2),
		fyne.NewContainerWithLayout(layout.NewGridLayout(3),
			makeFormTab(),
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

func makeFormTab() fyne.Widget {
	xMinEntry := widget.NewEntry()
	xMinEntry.SetPlaceHolder("xMin")

	xMaxEntry := widget.NewEntry()
	xMaxEntry.SetPlaceHolder("xMax")

	preciousEntry := widget.NewEntry()
	preciousEntry.SetPlaceHolder("Precious")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "xMin", Widget: xMinEntry},
			{Text: "xMax", Widget: xMaxEntry},
			{Text: "Precious", Widget: preciousEntry},
		},
		OnSubmit: func() {
			fmt.Println("Form submitted")
			//fyne.CurrentApp().SendNotification(&fyne.Notification{
			//	Title:   "Form for: " + name.Text,
			//	Content: largeText.Text,
			//})
		},
	}
	return form
}

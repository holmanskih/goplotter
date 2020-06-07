package ui

import (
	"fmt"
	"log"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/fyne-io/examples/img/icon"

	"github.com/Knetic/govaluate"
)

// Represents the UI fyne.io application struct
//type App struct {
//	ui fyne.App
//
//
//
//	buttons map[string]*widget.Button
//	window  fyne.Window
//}

type App struct {
	ui fyne.App

	equation  string
	functions map[string]func()

	output  *widget.Label
	buttons map[string]*widget.Button
	window  fyne.Window
}

func NewApp() *App {
	a := &App{}
	a.functions = make(map[string]func())
	a.buttons = make(map[string]*widget.Button)
	a.ui = app.New()
	return a
}

func (c *App) display(newtext string) {
	c.equation = newtext
	c.output.SetText(newtext)
}

func (c *App) character(char rune) {
	c.display(c.equation + string(char))
}

func (c *App) digit(d int) {
	r := rune(d)
	r += '0'
	c.character(r)
}

func (c *App) clear() {
	c.display("")
}

func (c *App) evaluate() {
	expression, err := govaluate.NewEvaluableExpression(c.output.Text)
	if err == nil {
		result, err := expression.Evaluate(nil)
		if err == nil {
			c.display(strconv.FormatFloat(result.(float64), 'f', -1, 64))
		}
	}

	if err != nil {
		log.Println("Error in calculation", err)
		c.display("error")
	}

	c.equation = ""
}

func (c *App) addButton(text string, action func()) *widget.Button {
	button := widget.NewButton(text, action)
	c.buttons[text] = button

	return button
}

func (c *App) digitButton(number int) *widget.Button {
	str := fmt.Sprintf("%d", number)
	action := func() {
		c.digit(number)
	}
	c.functions[str] = action

	return c.addButton(str, action)
}

func (c *App) charButton(char rune) *widget.Button {
	action := func() {
		c.character(char)
	}
	c.functions[string(char)] = action

	return c.addButton(string(char), action)
}

func (c *App) typedRune(r rune) {
	if r == '=' {
		c.evaluate()
		return
	} else if r == 'c' {
		c.clear()
		return
	}

	action := c.functions[string(r)]
	if action != nil {
		action()
	}
}

func (c *App) typedKey(ev *fyne.KeyEvent) {
	if ev.Name == fyne.KeyReturn || ev.Name == fyne.KeyEnter {
		c.evaluate()
		return
	}
}

func (c *App) LoadUI() {
	c.output = widget.NewLabel("")
	c.output.Alignment = fyne.TextAlignTrailing
	c.output.TextStyle.Monospace = true
	equals := c.addButton("=", func() {
		c.evaluate()
	})
	equals.Style = widget.PrimaryButton

	c.window = c.ui.NewWindow("App")
	c.window.SetIcon(icon.CalculatorBitmap)
	c.window.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(1),
		c.output,
		fyne.NewContainerWithLayout(layout.NewGridLayout(4),
			c.charButton('+'),
			c.charButton('-'),
			c.charButton('*'),
			c.charButton('/')),
		fyne.NewContainerWithLayout(layout.NewGridLayout(4),
			c.digitButton(7),
			c.digitButton(8),
			c.digitButton(9),
			c.addButton("C", func() {
				c.clear()
			})),
		fyne.NewContainerWithLayout(layout.NewGridLayout(4),
			c.digitButton(4),
			c.digitButton(5),
			c.digitButton(6),
			c.charButton('(')),
		fyne.NewContainerWithLayout(layout.NewGridLayout(4),
			c.digitButton(1),
			c.digitButton(2),
			c.digitButton(3),
			c.charButton(')')),
		fyne.NewContainerWithLayout(layout.NewGridLayout(2),
			fyne.NewContainerWithLayout(layout.NewGridLayout(2),
				c.digitButton(0),
				c.charButton('.')),
			equals)),
	)

	c.window.Canvas().SetOnTypedRune(c.typedRune)
	c.window.Canvas().SetOnTypedKey(c.typedKey)
	c.window.ShowAndRun()
}

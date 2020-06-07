package plotter

import (
	"gonum.org/v1/plot"
)

type Plotter struct {
	Plot *plot.Plot
}

func NewPlotter(plotName string) (*Plotter, error) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = plotName
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	return &Plotter{Plot: p}, err
}

// Save the plot to a PNG file.
//if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
//panic(err)
//}

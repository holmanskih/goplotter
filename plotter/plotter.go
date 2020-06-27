package plotter

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const (
	PlotMaxSizeX = 4 * vg.Inch
	PlotMaxSizeY = 4 * vg.Inch
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

func (p *Plotter) SetBounds(xMin, xMax, yMin, yMax float64) {
	p.Plot.X.Min = xMin
	p.Plot.X.Max = xMax
	p.Plot.Y.Min = yMin
	p.Plot.Y.Max = yMax

}

func (p *Plotter) AddFunc(f func(x float64) float64, legendName string) {
	pFunc := plotter.NewFunction(f)
	pFunc.Width = vg.Points(2)
	pFunc.Color = color.RGBA{G: 255, A: 255}

	p.Plot.Add(pFunc)
	p.Plot.Legend.Add(legendName, pFunc)
}

func (p *Plotter) Save(path string) error {
	return p.Plot.Save(PlotMaxSizeX, PlotMaxSizeY, path)
}

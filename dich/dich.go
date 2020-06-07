package dich

import (
	"fmt"

	"github.com/holmanskih/operations-research/internal/plotter"
)

type DichotomyService struct {
	F func(x float64) float64

	XMin, XMax float64
	plotter    *plotter.Plotter
}

func NewDichotomyService(f func(x float64) float64, xMin, xMax float64) (*DichotomyService, error) {
	p, err := plotter.NewPlotter("dich")
	if err != nil {
		return nil, fmt.Errorf("failed to create new plotter: %s", err)
	}

	return &DichotomyService{
		F:       f,
		XMin:    xMin,
		XMax:    xMax,
		plotter: p,
	}, nil
}

func (d *DichotomyService) Dichotomy() (float64, error) {
	if d.F(d.XMin)*d.F(d.XMax) >= 0 {
		return 0, fmt.Errorf("wrong interal")
	}
	c := d.XMin

	for (d.XMax - d.XMin) >= 0.01 {
		c = (d.XMin + d.XMax) / 2
		if d.F(c) == 0 {
			break
		}

		if d.F(c)*d.F(d.XMin) < 0 {
			d.XMax = c
		} else {
			d.XMin = c
		}
	}

	return c, nil
}

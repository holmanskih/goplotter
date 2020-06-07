package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/holmanskih/operations-research/dich"
)

func f(x float64) float64 {
	return x*x*x - x*x + 2
}

func main() {
	xMin := flag.Float64("xMin", -200, "xMin dichotomy")
	xMax := flag.Float64("xMax", 300, "xMax dichotomy")
	flag.Parse()

	dich, err := dich.NewDichotomyService(f, *xMin, *xMax)
	if err != nil {
		log.Fatalf("failed to run dichotomy service")
	}
	dichRes, err := dich.Dichotomy()
	if err != nil {
		log.Fatalf("failed to do dechotomy")
	}
	log.Println(dichRes)

	fmt.Println("xMin:", *xMin)
	fmt.Println("xMax:", *xMax)
}

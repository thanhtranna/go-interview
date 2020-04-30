package main

import (
	"image/color"
	"log"

	"go-interview/algorithms/gridder"
)

func main() {
	imageConfig := gridder.ImageConfig{
		Width:  500,
		Height: 500,
		Name:   "example2.png",
	}
	gridConfig := gridder.GridConfig{
		Rows:              4,
		Columns:           4,
		MarginWidth:       20,
		LineDashes:        10,
		LineStrokeWidth:   2,
		BorderDashes:      20,
		BorderStrokeWidth: 4,
		LineColor:         color.RGBA{R: 255 / 2, A: 255},
		BorderColor:       color.RGBA{B: 255 / 2, A: 255},
		BackgroundColor:   color.RGBA{G: 255 / 2, A: 255},
	}

	grid, err := gridder.New(imageConfig, gridConfig)
	if err != nil {
		log.Fatal(err)
	}

	grid.SavePNG()
}

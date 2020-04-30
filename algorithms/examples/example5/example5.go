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
		Name:   "example5.png",
	}
	gridConfig := gridder.GridConfig{
		Rows:              4,
		Columns:           4,
		LineStrokeWidth:   2,
		BorderStrokeWidth: 4,
	}
	grid, err := gridder.New(imageConfig, gridConfig)
	if err != nil {
		log.Fatal(err)
	}

	grid.DrawCircle(0, 0, gridder.CircleConfig{Radius: 30, Color: color.Black, Stroke: true})
	grid.DrawCircle(0, 3, gridder.CircleConfig{Radius: 30, Color: color.Black, Stroke: true, StrokeWidth: 25})
	grid.DrawCircle(2, 1, gridder.CircleConfig{Radius: 90, Color: color.RGBA{R: 255 / 2, A: 255 / 2}})
	grid.DrawCircle(3, 3, gridder.CircleConfig{Radius: 30, Color: color.Black, Stroke: false})
	grid.SavePNG()
}

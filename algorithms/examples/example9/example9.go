package main

import (
	"image/color"
	"log"

	"go-interview/algorithms/gridder"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

func main() {
	imageConfig := gridder.ImageConfig{
		Width:  2000,
		Height: 1000,
		Name:   "example9.png",
	}
	gridConfig := gridder.GridConfig{
		Rows:              4,
		Columns:           8,
		MarginWidth:       32,
		LineStrokeWidth:   2,
		BorderStrokeWidth: 20,
	}

	grid, err := gridder.New(imageConfig, gridConfig)
	if err != nil {
		log.Fatal(err)
	}

	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	fontFace := truetype.NewFace(font, &truetype.Options{Size: 48})

	lineConfig := gridder.PathConfig{Dashes: 10}
	circleConfig := gridder.CircleConfig{Color: color.Gray{}, Radius: 10}

	grid.PaintCell(1, 2, color.NRGBA{R: 0, G: 0, B: 0, A: 255 / 2})
	grid.DrawString(1, 2, "Block", fontFace)

	grid.DrawCircle(0, 0, gridder.CircleConfig{Color: color.NRGBA{R: 255 / 2, G: 0, B: 0, A: 255 / 2}, Radius: 60})
	grid.DrawPath(0, 0, 1, 1, lineConfig)
	grid.DrawCircle(1, 1, circleConfig)
	grid.DrawPath(1, 1, 2, 2, lineConfig)
	grid.DrawCircle(2, 2, circleConfig)
	grid.DrawPath(2, 2, 2, 3, lineConfig)
	grid.DrawCircle(2, 3, circleConfig)
	grid.DrawPath(2, 3, 2, 4, lineConfig)
	grid.DrawCircle(2, 4, circleConfig)
	grid.DrawPath(2, 4, 2, 5, lineConfig)
	grid.DrawCircle(2, 5, circleConfig)
	grid.DrawPath(2, 5, 2, 6, lineConfig)
	grid.DrawCircle(2, 6, circleConfig)
	grid.DrawPath(2, 6, 3, 7, lineConfig)
	grid.DrawCircle(3, 7, gridder.CircleConfig{Color: color.NRGBA{R: 0, G: 255 / 2, B: 0, A: 255 / 2}, Radius: 60})

	grid.SavePNG()
}

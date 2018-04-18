package main

import (
	"flag"
	"image"
	"image/color"
	"log"
)

var (
	WIDTH     = flag.Int("width", 50, "set image width")
	HEIGHT    = flag.Int("height", 512, "set image height")
	LANDSCAPE = flag.Bool("landscape", false, "set landscape mode")
	INVERSE   = flag.Bool("inverse", false, "inverse order")
	OUT       = flag.String("o", "out.png", "set file output")
)

func main() {
	flag.Parse()

	bounds := image.Rect(0, 0, *WIDTH, *HEIGHT)
	out := image.NewRGBA(bounds)

	gr := New()
	gr.Append(color.RGBA{R: 0, G: 0, B: 0, A: 255})
	gr.Append(color.RGBA{R: 56, G: 15, B: 109, A: 255})
	gr.Append(color.RGBA{R: 182, G: 54, B: 121, A: 255})
	gr.Append(color.RGBA{R: 253, G: 154, B: 105, A: 255})
	gr.Append(color.RGBA{R: 252, G: 246, B: 184, A: 255})

	if *LANDSCAPE {
		for x := 0; x < bounds.Dx(); x += 1 {
			c := gr.ColorAt(float64(x) / float64(bounds.Dx()))

			for y := 0; y < bounds.Dy(); y += 1 {
				if *INVERSE {
					out.SetRGBA(bounds.Dx()-x, y, c)
				} else {
					out.SetRGBA(x, y, c)
				}
			}
		}
	} else {
		for y := 0; y < bounds.Dy(); y += 1 {
			c := gr.ColorAt(float64(y) / float64(bounds.Dy()))

			for x := 0; x < bounds.Dx(); x += 1 {
				if *INVERSE {
					out.SetRGBA(x, bounds.Dy()-y, c)
				} else {
					out.SetRGBA(x, y, c)
				}
			}
		}
	}

	err := savePng(out, *OUT)
	if err != nil {
		log.Fatalf("savePng failed: %v", err)
	}
}

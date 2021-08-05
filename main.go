package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	data := []int{10, 33, 73, 64}
	h, w := 100, len(data) * 60 + 10
	r := image.Rect(0, 0, w, h)
	img:= image.NewRGBA(r)
	f, err := os.Create("image.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	//set background - white
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.Set(x, y, color.RGBA{
				R: 255,
				G: 255,
				B: 255,
				A: 255,
			})
		}
	}
	
	//create bar charts
	for i, dp := range data  {
		for x := 10+60*i; x <= 60*(i+1); x++ {
			for y := h; y >= h-dp; y-- {
				img.Set(x, y, color.RGBA{
					R: 200,
					G: 100,
					B: 100,
					A: 255,
				})
			}
		}

	}

	if err = png.Encode(f, img); err != nil {
		log.Fatalln(err)
	}
}


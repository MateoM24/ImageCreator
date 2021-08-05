package main

import (
	"image"
	"image/png"
	"log"
	"os"
)

func main() {
	r := image.Rect(0, 0, 500, 150)
	img:= image.NewRGBA(r)
	f, err := os.Create("image.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	if err = png.Encode(f, img); err != nil {
		log.Fatalln(err)
	}


}


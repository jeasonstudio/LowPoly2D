package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
)

func bbb() {
	// ddd := 10

	file, err := os.Open("cat.jpg")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file1, err := os.Create("tag.jpg")

	if err != nil {
		fmt.Println(err)
	}
	defer file1.Close()
	img, _ := jpeg.Decode(file)

	xWidth := img.Bounds().Dx()
	yHeight := img.Bounds().Dy()

	jpg := image.NewRGBA64(img.Bounds())

	// for (i := 1,p := 1); (i / ddd) <= (xWidth / ddd); (i++,p=i) {
	// 	for (j := 1,q := 1); (j / ddd) < (yHeight / ddd); (j++,q=j) {
	// 		thisR, thisG, thisB, thisA := img.At(i, j).RGBA()
	// 		jpg.SetRGBA64(i, j)
	// 	}
	// }

	draw.Draw(jpg, img.Bounds().Add(image.Pt(xWidth, yHeight)), img, img.Bounds().Min, draw.Src)
	jpeg.Encode(file1, jpg, nil)

}

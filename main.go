package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"math/rand"
	"os"
	"time"
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

func main() {
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

	for i := 1; i < xWidth-1; i++ {
		for j := 1; j < yHeight-1; j++ {
			o := rand.New(rand.NewSource(time.Now().UnixNano()))
			fmt.Println(o.Int()%3 == 0)
			if o.Int()%2 == 0 && o.Int()%3 == 0 && o.Int()%5 == 0 && o.Int()%7 == 0 {
				var newColor color.RGBA64
				newColor.R = 65535
				newColor.G = 65535
				newColor.B = 65535
				newColor.A = 65535
				jpg.SetRGBA64(i, j, newColor)
			} else {
				thisR, thisG, thisB, thisA := img.At(i, j).RGBA()
				var newColor color.RGBA64
				newColor.R = uint16(thisR)
				newColor.G = uint16(thisG)
				newColor.B = uint16(thisB)
				newColor.A = uint16(thisA)
				jpg.SetRGBA64(i, j, newColor)
			}
		}
	}

	draw.Draw(jpg, img.Bounds().Add(image.Pt(xWidth, yHeight)), img, img.Bounds().Min, draw.Src)
	jpeg.Encode(file1, jpg, nil)
}

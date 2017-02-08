package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
)

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

	jpg := image.NewGray16(img.Bounds())

	for i := 0; i < xWidth; i++ {
		for j := 0; j < yHeight; j++ {

			// 四方向索贝尔算子
			// GX := SumGray(RGBAToGray(img.At(i+1, j-1)), RGBAToGray(img.At(i+1, j)), RGBAToGray(img.At(i+1, j+1)), RGBAToGray(img.At(i-1, j-1)), RGBAToGray(img.At(i-1, j)), RGBAToGray(img.At(i-1, j+1)))
			// GY := SumGray(RGBAToGray(img.At(i-1, j-1)), RGBAToGray(img.At(i, j-1)), RGBAToGray(img.At(i+1, j-1)), RGBAToGray(img.At(i-1, j+1)), RGBAToGray(img.At(i, j+1)), RGBAToGray(img.At(i+1, j+1)))
			// G := GX + GY

			// 八方向索贝尔算子
			G := SumEightGray(img, i, j)
			if G > 63000 {
				fmt.Println("(", i, ",", j, ")")
				var m color.Gray16
				m.Y = G
				jpg.SetGray16(i, j, m)
			}
			// fmt.Println(G)
			// var m color.Gray16
			// m.Y = G
			// jpg.SetGray16(i, j, m)
		}
	}

	// gg := jpg.Pix[7]

	// fmt.Println(gg)

	draw.Draw(jpg, img.Bounds().Add(image.Pt(xWidth, yHeight)), img, img.Bounds().Min, draw.Src)
	jpeg.Encode(file1, jpg, nil)

}

// RGBAToGray change RGB to Gray
func RGBAToGray(color color.Color) uint16 {
	thisR, thisG, thisB, _ := color.RGBA()
	return uint16((thisR*299 + thisG*587 + thisB*114 + 500) / 1000)
}

// SumGray sum tag Gx 四方向索贝尔算子
func SumGray(a, b, c, d, e, f uint16) uint16 {
	return (a + 2*b + c) - (d + 2*e + f)
}

// SumEightGray 八方向索贝尔算子
func SumEightGray(img image.Image, i, j int) uint16 {

	G1 := (RGBAToGray(img.At(i-2, j+1)) + 2*RGBAToGray(img.At(i-1, j+1)) + 4*RGBAToGray(img.At(i, j+1)) + 2*RGBAToGray(img.At(i+1, j+1)) + RGBAToGray(img.At(i+2, j+1))) - (RGBAToGray(img.At(i-2, j-1)) + 2*RGBAToGray(img.At(i-1, j-1)) + 4*RGBAToGray(img.At(i, j-1)) + 2*RGBAToGray(img.At(i+1, j-1)) + RGBAToGray(img.At(i+2, j-1)))

	G2 := (RGBAToGray(img.At(i+2, j)) + 2*RGBAToGray(img.At(i+1, j+1)) + 4*RGBAToGray(img.At(i, j+1)) + 2*RGBAToGray(img.At(i-1, j+1)) + 4*RGBAToGray(img.At(i+1, j))) - (RGBAToGray(img.At(i-2, j)) + 2*RGBAToGray(img.At(i-1, j-1)) + 4*RGBAToGray(img.At(i, j-1)) + 2*RGBAToGray(img.At(i+1, j-1)) + 4*RGBAToGray(img.At(i-1, j)))

	G3 := (RGBAToGray(img.At(i+2, j-1)) + 2*RGBAToGray(img.At(i+1, j+1)) + 4*RGBAToGray(img.At(i, j+1)) + RGBAToGray(img.At(i-1, j+2)) + 4*RGBAToGray(img.At(i+1, j))) - (RGBAToGray(img.At(i-2, j+1)) + 2*RGBAToGray(img.At(i-1, j-1)) + 4*RGBAToGray(img.At(i, j-1)) + RGBAToGray(img.At(i+1, j-2)) + 4*RGBAToGray(img.At(i-1, j)))

	G4 := (2*RGBAToGray(img.At(i+1, j-1)) + 2*RGBAToGray(img.At(i+1, j+1)) + 4*RGBAToGray(img.At(i, j+1)) + RGBAToGray(img.At(i, j+2)) + 4*RGBAToGray(img.At(i+1, j))) - (2*RGBAToGray(img.At(i-1, j+1)) + 2*RGBAToGray(img.At(i-1, j-1)) + 4*RGBAToGray(img.At(i, j-1)) + RGBAToGray(img.At(i, j-2)) + 4*RGBAToGray(img.At(i-1, j)))

	G5 := (RGBAToGray(img.At(i+1, j-2)) + 2*RGBAToGray(img.At(i+1, j-1)) + 4*RGBAToGray(img.At(i+1, j)) + 2*RGBAToGray(img.At(i+1, j+1)) + RGBAToGray(img.At(i+1, j+2))) - (RGBAToGray(img.At(i-1, j-2)) + 2*RGBAToGray(img.At(i-1, j-1)) + 4*RGBAToGray(img.At(i-1, j)) + 2*RGBAToGray(img.At(i-1, j+1)) + RGBAToGray(img.At(i-1, j+2)))

	G6 := (RGBAToGray(img.At(i, j+2)) + 2*RGBAToGray(img.At(i+1, j+1)) + 4*RGBAToGray(img.At(i+1, j)) + 2*RGBAToGray(img.At(i+1, j-1)) + 4*RGBAToGray(img.At(i, j+1))) - (RGBAToGray(img.At(i, j-2)) + 2*RGBAToGray(img.At(i-1, j-1)) + 4*RGBAToGray(img.At(i-1, j)) + 2*RGBAToGray(img.At(i-1, j+1)) + 4*RGBAToGray(img.At(i, j-1)))

	G7 := (RGBAToGray(img.At(i-1, j+2)) + 2*RGBAToGray(img.At(i+1, j+1)) + 4*RGBAToGray(img.At(i+1, j)) + RGBAToGray(img.At(i+2, j-1)) + 4*RGBAToGray(img.At(i, j+1))) - (RGBAToGray(img.At(i+1, j-2)) + 2*RGBAToGray(img.At(i-1, j-1)) + 4*RGBAToGray(img.At(i-1, j)) + RGBAToGray(img.At(i-2, j+1)) + 4*RGBAToGray(img.At(i, j-1)))

	G8 := (2*RGBAToGray(img.At(i-1, j+1)) + 2*RGBAToGray(img.At(i+1, j+1)) + 4*RGBAToGray(img.At(i+1, j)) + RGBAToGray(img.At(i+2, j)) + 4*RGBAToGray(img.At(i, j+1))) - (2*RGBAToGray(img.At(i+1, j-1)) + 2*RGBAToGray(img.At(i-1, j-1)) + 4*RGBAToGray(img.At(i-1, j)) + RGBAToGray(img.At(i-2, j)) + 4*RGBAToGray(img.At(i, j-1)))

	return G1 + G2 + G3 + G4 + G5 + G6 + G7 + G8
}

// Gx = [  f(x+1,y-1) + 2*f(x+1,y) + f(x+1,y+1)  ]  -  [  f(x-1,y-1) + 2*f(x-1,y) + f(x-1,y+1)  ]
// Gy = [  f(x-1,y-1) + 2*f(x,y-1) + f(x+1,y-1)  ]  -  [  f(x-1,y+1) + 2*f(x,y+1) + f(x+1,y+1)  ]
// 224676

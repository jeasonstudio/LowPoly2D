package SobelEdge

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"math"
	"os"
)

// SobelEdge 索贝尔算子处理图片边缘
func SobelEdge(sourceImg, tagImg string, YUDATA uint16) {
	file, err := os.Open(sourceImg)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file1, err := os.Create(tagImg)

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
			// GYX, GXY := SumGray(img, i, j)
			// fmt.Println(GYX, GXY)
			// fmt.Println(G - RGBAToGray(img.At(i, j)))

			// 四方向索贝尔算子
			// G := SumFourGray(img, i, j)

			// 四方向索贝尔算子
			G := SumGrayNo(img, i, j)

			// 八方向索贝尔算子
			// G := SumEightGray(img, i, j)
			if G > YUDATA {
				fmt.Println("(", i, ",", j, ")", G)
				var m color.Gray16
				m.Y = 65535
				jpg.SetGray16(i, j, m)
			}
			// fmt.Println(G)
			// var m color.Gray16
			// m.Y = G
			// jpg.SetGray16(i, j, m)
		}
	}

	draw.Draw(jpg, img.Bounds().Add(image.Pt(xWidth, yHeight)), img, img.Bounds().Min, draw.Src)
	jpeg.Encode(file1, jpg, nil)

}

// RGBAToGray change RGB to Gray
func RGBAToGray(color color.Color) uint16 {
	thisR, thisG, thisB, _ := color.RGBA()
	return uint16((thisR*299 + thisG*587 + thisB*114 + 500) / 1000)
}

// SumGray sum tag Gx 四方向索贝尔算子
func SumGray(img image.Image, i, j int) (float64, float64) {

	var atyx, atxy float64

	GX := (RGBAToGray(img.At(i+1, j-1)) + 2*RGBAToGray(img.At(i+1, j)) + RGBAToGray(img.At(i+1, j+1))) - (RGBAToGray(img.At(i-1, j-1)) + 2*RGBAToGray(img.At(i-1, j)) + RGBAToGray(img.At(i-1, j+1)))
	GY := (RGBAToGray(img.At(i-1, j-1)) + 2*RGBAToGray(img.At(i, j-1)) + RGBAToGray(img.At(i+1, j-1))) - (RGBAToGray(img.At(i-1, j+1)) + 2*RGBAToGray(img.At(i, j+1)) + RGBAToGray(img.At(i+1, j+1)))
	if (GY > 0) && (GX > 0) {
		atyx = math.Atan(float64(GY / GX))
		atxy = math.Atan(float64(GX / GY))
	} else if GX <= 0 && GY > 0 {
		atxy = math.Atan(float64(GX / GY))
		atyx = math.Pi / 2
	} else if GY <= 0 && GX > 0 {
		atxy = math.Pi / 2
		atyx = math.Atan(float64(GY / GX))
	}
	return atyx, atxy
}

// SumGrayNo sum tag Gx 四方向索贝尔算子
func SumGrayNo(img image.Image, i, j int) uint16 {

	GX := (RGBAToGray(img.At(i+1, j-1)) + 2*RGBAToGray(img.At(i+1, j)) + RGBAToGray(img.At(i+1, j+1))) - (RGBAToGray(img.At(i-1, j-1)) + 2*RGBAToGray(img.At(i-1, j)) + RGBAToGray(img.At(i-1, j+1)))
	GY := (RGBAToGray(img.At(i-1, j-1)) + 2*RGBAToGray(img.At(i, j-1)) + RGBAToGray(img.At(i+1, j-1))) - (RGBAToGray(img.At(i-1, j+1)) + 2*RGBAToGray(img.At(i, j+1)) + RGBAToGray(img.At(i+1, j+1)))

	return GX + GY
}

// SumFourGray 八方向索贝尔算子
func SumFourGray(img image.Image, i, j int) uint16 {

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

// SumEightGray 四方向索贝尔算子
func SumEightGray(img image.Image, i, j int) uint16 {

	GX := (RGBAToGray(img.At(i-2, j+1)) + 2*RGBAToGray(img.At(i-1, j+1)) + 4*RGBAToGray(img.At(i, j+1)) + 2*RGBAToGray(img.At(i+1, j+1)) + RGBAToGray(img.At(i+2, j+1))) - (RGBAToGray(img.At(i-2, j-1)) + 2*RGBAToGray(img.At(i-1, j-1)) + 4*RGBAToGray(img.At(i, j-1)) + 2*RGBAToGray(img.At(i+1, j-1)) + RGBAToGray(img.At(i+2, j-1)))

	GY := (RGBAToGray(img.At(i+1, j-2)) + 2*RGBAToGray(img.At(i+1, j-1)) + 4*RGBAToGray(img.At(i+1, j)) + 2*RGBAToGray(img.At(i+1, j+1)) + RGBAToGray(img.At(i+1, j+2))) - (RGBAToGray(img.At(i-1, j-2)) + 2*RGBAToGray(img.At(i-1, j-1)) + 4*RGBAToGray(img.At(i-1, j)) + 2*RGBAToGray(img.At(i-1, j+1)) + RGBAToGray(img.At(i-1, j+2)))

	return GX + GY
}

// Gx = [  f(x+1,y-1) + 2*f(x+1,y) + f(x+1,y+1)  ]  -  [  f(x-1,y-1) + 2*f(x-1,y) + f(x-1,y+1)  ]
// Gy = [  f(x-1,y-1) + 2*f(x,y-1) + f(x+1,y-1)  ]  -  [  f(x-1,y+1) + 2*f(x,y+1) + f(x+1,y+1)  ]
// 224676

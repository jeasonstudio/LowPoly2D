package main

import (
	"fmt"
	"os"
	"image/draw"
	"image/jpeg"
	"image/color"
	"image"
	"math"
)

const(
	resetColor = ""
)

func main() {
	fmt.Println("Hello World")
	file, err := os.Open("./testFlower.JPG")
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

	// jpg2 := image.NewGray16(img.Bounds())
	// for i := 0; i < xWidth; i++ {
	// 	for j := 0; j < yHeight; j++ {
	// 		thisR, _, _, _ := img.At(i,j).RGBA()
	// 		// var n color.RGBA64
	// 		var m color.Gray16
	// 		if thisR < 40000 {
	// 			// fmt.Println(thisR)
	// 			// n.R = strongColor(thisR, 50)
	// 			// n.G = uint16(thisG)
	// 			// n.B = uint16(thisB)
	// 			// n.A = uint16(thisA)
	// 			// jpg.SetRGBA64(i,j,n)
	// 			m.Y = RGBAToGray(img.At(i,j))
	// 			jpg2.SetGray16(i, j, m)
	// 		}
	// 	}
	// }
	// draw.Draw(jpg2, img.Bounds().Add(image.Pt(xWidth, yHeight)), img, img.Bounds().Min, draw.Src)
	// jpeg.Encode(file1, jpg2, nil)


	jpg := image.NewRGBA64(img.Bounds())
	for i := 0; i < xWidth; i++ {
		for j := 0; j < yHeight; j++ {
			thisR, thisG, thisB, thisA := img.At(i,j).RGBA()
			var n color.RGBA64
			fmt.Println(RGBAToHSV(img.At(i,j)))
			if thisR > 30000 {
				n.R = strongColor(thisR, 50)
				n.G = uint16(thisG)
				n.B = uint16(thisB)
				n.A = uint16(thisA)
			} else {
				n.R = uint16(thisR)
				n.G = uint16(thisG)
				n.B = uint16(thisB)
				n.A = uint16(thisA)
			}
			jpg.SetRGBA64(i,j,n)
		}
	}
	draw.Draw(jpg, img.Bounds().Add(image.Pt(xWidth, yHeight)), img, img.Bounds().Min, draw.Src)
	jpeg.Encode(file1, jpg, nil)

}

func RGBAToGray(color color.Color) uint16  {
	thisR, thisG, thisB, _ := color.RGBA()
	return uint16((thisR*299 + thisG*587 + thisB*114 + 500) / 1000)
}

func strongColor(t uint32, how int) uint16  {
	if uint16(int(t) * (how/100 + 1)) > uint16(65535) {
		return uint16(65535) 
	} else {
	 	return uint16(int(t) * (how/100 + 1))
	}
}

// func wakeColor(t uint32, how int) uint16  {
// 	return nil
// }

func RGBAToHSV(c color.Color) (int,int,int)  {
	R, G, B, _ := c.RGBA()
	thisR := int(R/math.MaxUint32)
	thisG := int(G/math.MaxUint32)
	thisB := int(B/math.MaxUint32)

	max := getMax(thisR, thisG, thisB)
	min := getMin(thisR, thisG, thisB)
	var H int

	if thisR == max {
		H = int((thisG - thisB)/(max - min))
	} else if thisG == max {
		H = 2 + int((thisB - thisR)/(max - min))
	} else if thisB == max {
		H = 4 + int((thisR - thisG)/(max - min))
	} 
	H = H * 60
	if H < 0 { H = H + 360}

	S := (max - min)/max
	V := getMax(thisR, thisG, thisB)

	return H, S, V
}

func getMax(a, b, c int) int  {
	if a > b && a > c {
		return a
	} else if b > a && b > c {
		return b
	} else if c > a && c > b {
		return c
	} else {return 0}
}

func getMin(a, b, c int) int  {
	if a < b && a < c {
		return a
	} else if b < a && b < c {
		return b
	} else if c < a && c < b {
		return c
	} else {return 0}
}
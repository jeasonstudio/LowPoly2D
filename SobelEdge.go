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

// Putpixel 划线函数
type Putpixel func(x, y int)

// SobelEdge 索贝尔算子处理图片边缘
func SobelEdge(sourceImg, tagImg string, lowSigema, highSigema uint16, p, q int) {
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
	// oldX := 0
	// oldY := 0

	for i := 1; i < xWidth-1; i++ {
		for j := 1; j < yHeight-1; j++ {

			// 四方向索贝尔算子
			// GYX, GXY := SumGray(img, i, j)
			// fmt.Println(GYX, GXY)
			// fmt.Println(G - RGBAToGray(img.At(i, j)))

			// 四方向索贝尔算子
			// G := SumFourGray(img, i, j)

			// 八方向索贝尔算子
			G := SumEightGray(img, i, j)
			// fmt.Println(GX + GY - RGBAToGray(img.At(i, j)))

			// 拉布拉斯算子
			// G := LaplaceGray(img, i, j)
			// fmt.Println(G, YUDATA)
			// fmt.Println(GX, GY)
			// if GX+GY <= 1000 {
			// }
			// var G uint16

			// 四方向索贝尔算子
			// GB := SumGrayNo(img, i, j)

			if G > lowSigema && G < highSigema {
				// fmt.Println("(", i, ",", j, ")", G)
				var m color.Gray16
				m.Y = G
				jpg.SetGray16(i, j, m)
			}
			// var m color.Gray16
			// o := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10000)

			// // fmt.Println(o)
			// if GX+GY > lowSigema && GX+GY < highSigema && o <= p {
			// 	// fmt.Println("(", i, ",", j, ")", G)
			// 	m.Y = 65535
			// 	jpg.SetGray16(i, j, m)
			// 	// drawline(oldX, oldY, i, j, func(x, y int) {
			// 	// 	jpg.Set(x, y, color.RGBA64{65535, 65535, 65535, 65535})
			// 	// })
			// 	// oldX = i
			// 	// oldY = j
			// } else if o <= q {
			// 	m.Y = 50000
			// 	jpg.SetGray16(i, j, m)
			// }
		}
	}

	// drawline(5, 5, xWidth-8, yHeight-10, func(x, y int) {
	// 	jpg.Set(x, y, color.RGBA64{65535, 65535, 65535, 65535})
	// })

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

	if GX > GY {
		return GX
	} else {
		return GY
	}
}

// SumEightGray 八方向索贝尔算子
func SumEightGray(img image.Image, i, j int) uint16 {

	var G []uint16

	G = append(G, (RGBAToGray(img.At(i-2, j+1))+2*RGBAToGray(img.At(i-1, j+1))+4*RGBAToGray(img.At(i, j+1))+2*RGBAToGray(img.At(i+1, j+1))+RGBAToGray(img.At(i+2, j+1)))-(RGBAToGray(img.At(i-2, j-1))+2*RGBAToGray(img.At(i-1, j-1))+4*RGBAToGray(img.At(i, j-1))+2*RGBAToGray(img.At(i+1, j-1))+RGBAToGray(img.At(i+2, j-1))))

	G = append(G, (RGBAToGray(img.At(i+2, j))+2*RGBAToGray(img.At(i+1, j+1))+4*RGBAToGray(img.At(i, j+1))+2*RGBAToGray(img.At(i-1, j+1))+4*RGBAToGray(img.At(i+1, j)))-(RGBAToGray(img.At(i-2, j))+2*RGBAToGray(img.At(i-1, j-1))+4*RGBAToGray(img.At(i, j-1))+2*RGBAToGray(img.At(i+1, j-1))+4*RGBAToGray(img.At(i-1, j))))

	G = append(G, (RGBAToGray(img.At(i+2, j-1))+2*RGBAToGray(img.At(i+1, j+1))+4*RGBAToGray(img.At(i, j+1))+RGBAToGray(img.At(i-1, j+2))+4*RGBAToGray(img.At(i+1, j)))-(RGBAToGray(img.At(i-2, j+1))+2*RGBAToGray(img.At(i-1, j-1))+4*RGBAToGray(img.At(i, j-1))+RGBAToGray(img.At(i+1, j-2))+4*RGBAToGray(img.At(i-1, j))))

	G = append(G, (2*RGBAToGray(img.At(i+1, j-1))+2*RGBAToGray(img.At(i+1, j+1))+4*RGBAToGray(img.At(i, j+1))+RGBAToGray(img.At(i, j+2))+4*RGBAToGray(img.At(i+1, j)))-(2*RGBAToGray(img.At(i-1, j+1))+2*RGBAToGray(img.At(i-1, j-1))+4*RGBAToGray(img.At(i, j-1))+RGBAToGray(img.At(i, j-2))+4*RGBAToGray(img.At(i-1, j))))

	G = append(G, (RGBAToGray(img.At(i+1, j-2))+2*RGBAToGray(img.At(i+1, j-1))+4*RGBAToGray(img.At(i+1, j))+2*RGBAToGray(img.At(i+1, j+1))+RGBAToGray(img.At(i+1, j+2)))-(RGBAToGray(img.At(i-1, j-2))+2*RGBAToGray(img.At(i-1, j-1))+4*RGBAToGray(img.At(i-1, j))+2*RGBAToGray(img.At(i-1, j+1))+RGBAToGray(img.At(i-1, j+2))))

	G = append(G, (RGBAToGray(img.At(i, j+2))+2*RGBAToGray(img.At(i+1, j+1))+4*RGBAToGray(img.At(i+1, j))+2*RGBAToGray(img.At(i+1, j-1))+4*RGBAToGray(img.At(i, j+1)))-(RGBAToGray(img.At(i, j-2))+2*RGBAToGray(img.At(i-1, j-1))+4*RGBAToGray(img.At(i-1, j))+2*RGBAToGray(img.At(i-1, j+1))+4*RGBAToGray(img.At(i, j-1))))

	G = append(G, (RGBAToGray(img.At(i-1, j+2))+2*RGBAToGray(img.At(i+1, j+1))+4*RGBAToGray(img.At(i+1, j))+RGBAToGray(img.At(i+2, j-1))+4*RGBAToGray(img.At(i, j+1)))-(RGBAToGray(img.At(i+1, j-2))+2*RGBAToGray(img.At(i-1, j-1))+4*RGBAToGray(img.At(i-1, j))+RGBAToGray(img.At(i-2, j+1))+4*RGBAToGray(img.At(i, j-1))))

	G = append(G, (2*RGBAToGray(img.At(i-1, j+1))+2*RGBAToGray(img.At(i+1, j+1))+4*RGBAToGray(img.At(i+1, j))+RGBAToGray(img.At(i+2, j))+4*RGBAToGray(img.At(i, j+1)))-(2*RGBAToGray(img.At(i+1, j-1))+2*RGBAToGray(img.At(i-1, j-1))+4*RGBAToGray(img.At(i-1, j))+RGBAToGray(img.At(i-2, j))+4*RGBAToGray(img.At(i, j-1))))

	maxUint := G[0]

	for i := 1; i < 8; i++ {
		if G[i] > maxUint {
			maxUint = G[i]
		}
	}
	return maxUint
}

// SumFourGray 四方向索贝尔算子
func SumFourGray(img image.Image, i, j int) uint16 {

	GX := (RGBAToGray(img.At(i-2, j+1)) + 2*RGBAToGray(img.At(i-1, j+1)) + 4*RGBAToGray(img.At(i, j+1)) + 2*RGBAToGray(img.At(i+1, j+1)) + RGBAToGray(img.At(i+2, j+1))) - (RGBAToGray(img.At(i-2, j-1)) + 2*RGBAToGray(img.At(i-1, j-1)) + 4*RGBAToGray(img.At(i, j-1)) + 2*RGBAToGray(img.At(i+1, j-1)) + RGBAToGray(img.At(i+2, j-1)))

	GY := (RGBAToGray(img.At(i+1, j-2)) + 2*RGBAToGray(img.At(i+1, j-1)) + 4*RGBAToGray(img.At(i+1, j)) + 2*RGBAToGray(img.At(i+1, j+1)) + RGBAToGray(img.At(i+1, j+2))) - (RGBAToGray(img.At(i-1, j-2)) + 2*RGBAToGray(img.At(i-1, j-1)) + 4*RGBAToGray(img.At(i-1, j)) + 2*RGBAToGray(img.At(i-1, j+1)) + RGBAToGray(img.At(i-1, j+2)))

	return GX + GY
}

// LaplaceGray 拉布拉斯算子
func LaplaceGray(img image.Image, i, j int) uint16 {
	return RGBAToGray(img.At(i-1, j-1)) + RGBAToGray(img.At(i-1, j)) + RGBAToGray(img.At(i-1, j+1)) + RGBAToGray(img.At(i, j-1)) + RGBAToGray(img.At(i, j+1)) + RGBAToGray(img.At(i+1, j-1)) + RGBAToGray(img.At(i+1, j+1)) - 8*RGBAToGray(img.At(i, j))
}

// Gx = [  f(x+1,y-1) + 2*f(x+1,y) + f(x+1,y+1)  ]  -  [  f(x-1,y-1) + 2*f(x-1,y) + f(x-1,y+1)  ]
// Gy = [  f(x-1,y-1) + 2*f(x,y-1) + f(x+1,y-1)  ]  -  [  f(x-1,y+1) + 2*f(x,y+1) + f(x+1,y+1)  ]
// laplace ={ f(x+1,y)+f(x-1,y)+2f(x,y) } +  { f(x,y-1) + f(x, y+1) - 2f(x, y)}= \\\\\f(x-1,y) + f(x+1,y) + f(x,y-1) + f(x, y+1) - 4f(x, y);
// 224676

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func drawline(x0, y0, x1, y1 int, brush Putpixel) {
	dx := abs(x1 - x0)
	dy := abs(y1 - y0)
	sx, sy := 1, 1
	if x0 >= x1 {
		sx = -1
	}
	if y0 >= y1 {
		sy = -1
	}
	err := dx - dy
	for {
		brush(x0, y0)
		if x0 == x1 && y0 == y1 {
			return
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

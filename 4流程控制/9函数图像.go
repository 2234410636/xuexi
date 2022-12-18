package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func main() {
	const size = 300 //图像大小
	pic := image.NewGray(image.Rect(0, 0, size, size))
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			pic.SetGray(x, y, color.Gray{255})

		}
	}
	for x := 0; x < size; x++ {
		s := float64(x) * 2 * math.Pi / size
		y := size/2 - math.Cos(s)*size/2
		pic.SetGray(x, int(y), color.Gray{00})
	}
	file, _ := os.Create("Cos.png")
	png.Encode(file, pic)
	file.Close()

}

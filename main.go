package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
)

func main() {

	myApp := app.New()
	w := myApp.NewWindow("Mandelbrot Set")

	raster := canvas.NewRasterWithPixels(
		func(x, y, w, h int) color.Color {
			return getColorFromPos(x, y, w, h)
		})
	w.SetContent(raster)

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		if k.Name == "G" {
			raster.Refresh()
		}
		fmt.Println(k.Name)
	})

	w.Resize(fyne.NewSize(600, 600))
	w.ShowAndRun()
}

func getColorFromPos(xPos, yPos, xSize, ySize int) color.Color {
	manX := ((float64(xPos) / float64(xSize)) * 2.5) - 2
	manY := ((float64(yPos) / float64(ySize)) * 2.5) - 1.25

	c := complex(manX, manY)

	return colorFromComplexNum(c)
}

func colorFromComplexNum(c complex128) color.Color {
	iVal := calcIterations(c)

	// black and white version
	// return color.RGBA{255 - uint8(iVal), 255 - uint8(iVal), 255 - uint8(iVal), 255}

	var colorList [11]color.Color
	var threshold [11]int

	colorList[0] = color.RGBA{249, 65, 68, 255}
	threshold[0] = 2

	colorList[1] = color.RGBA{243, 114, 44, 255}
	threshold[1] = 3

	colorList[2] = color.RGBA{249, 132, 74, 255}
	threshold[2] = 4

	colorList[3] = color.RGBA{248, 150, 30, 255}
	threshold[3] = 5

	colorList[4] = color.RGBA{249, 199, 79, 255}
	threshold[4] = 8

	colorList[5] = color.RGBA{144, 190, 109, 255}
	threshold[5] = 16

	colorList[6] = color.RGBA{67, 170, 139, 255}
	threshold[6] = 32

	colorList[7] = color.RGBA{77, 144, 142, 255}
	threshold[7] = 64

	colorList[8] = color.RGBA{87, 117, 144, 255}
	threshold[8] = 100

	colorList[9] = color.RGBA{39, 125, 161, 255}
	threshold[9] = 175

	colorList[10] = color.RGBA{39, 79, 160, 255}
	threshold[10] = 254

	for i := 0; i < len(threshold); i++ {
		if threshold[i] > iVal {
			return colorList[i]
		}
	}
	return color.RGBA{0, 0, 0, 255}

}

func calcIterations(c complex128) int {
	z := 0 + 0i
	i := 0

	for i < 255 && real(z)*real(z)+imag(z)*imag(z) < 4 {
		z *= z
		z += c
		i++
	}

	return i
}

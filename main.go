package main

import (
	"fmt"
	"image/color"
	"math"

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
	iVal, maxI := calcIterations(c)

	H := math.Mod(float64(iVal*10), 360)
	S := float64(0.8)
	L := float64(0.5)

	if iVal == maxI {
		return color.RGBA{0, 0, 0, 255}
	}

	return hslToRGBA(H, L, S)

}

func calcIterations(c complex128) (int, int) {
	max := 1000
	z := 0 + 0i
	i := 0

	for i < max && real(z)*real(z)+imag(z)*imag(z) < 4 {
		z *= z
		z += c
		i++
	}

	return i, max
}

func hslToRGBA(H, L, S float64) color.Color {

	C := (1 - math.Abs((2*L)-1)) * S
	X := C * (1 - math.Abs(math.Mod((H/60), 2)-1))
	M := L - (C / 2)

	var R, G, B float64

	if H >= 0 && H < 60 {
		R = C
		G = X
		B = 0
	} else if H >= 60 && H < 120 {
		R = X
		G = C
		B = 0
	} else if H >= 120 && H < 180 {
		R = 0
		G = C
		B = X
	} else if H >= 180 && H < 240 {
		R = 0
		G = X
		B = C
	} else if H >= 240 && H < 300 {
		R = X
		G = 0
		B = C
	} else {
		R = C
		G = 0
		B = X
	}

	return color.RGBA{uint8((R + M) * 255), uint8((G + M) * 255), uint8((B + M) * 255), 255}
}

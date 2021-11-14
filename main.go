package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Raster")

	blue := uint8(0)
	raster := canvas.NewRasterWithPixels(
		func(x, y, w, h int) color.Color {
			return getColorFromPos(x, y, w, h, blue)
		})
	w.SetContent(raster)

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		if k.Name == "G" {
			blue += 10
			raster.Refresh()
		}
		fmt.Println(k.Name)
	})

	w.Resize(fyne.NewSize(600, 600))
	w.ShowAndRun()
}

func getColorFromPos(xPos, yPos, xSize, ySize int, blue uint8) color.Color {
	red := uint8(255 * float32(xPos) / float32(xSize))
	green := uint8(255 * float32(yPos) / float32(ySize))

	return color.RGBA{red, green, blue, 255}
}

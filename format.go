package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"image/jpeg"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func formatLabel() {
	// Create blank label with specified dimensions
	width := 991
	height := 306
	canvas := image.NewRGBA(image.Rect(0, 0, width, height))
	
	// Fill with white background
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
	
	// Overlay an image at specific coordinates
	overlayImage(canvas, "assets/monash_automation_logo.png", 100, 150)
	
	// Add text at specific coordinates
	addText(canvas, 100, 200, "Hello, World!", color.RGBA{0, 0, 0, 255})
	
	// Save the result
	outFile, err := os.Create("temp/label.jpg")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	
	jpeg.Encode(outFile, canvas, &jpeg.Options{Quality: 95})
}

func overlayImage(canvas *image.RGBA, imagePath string, x, y int) error {
	file, err := os.Open(imagePath)
	if err != nil {
		return err
	}
	defer file.Close()
	
	img, err := png.Decode(file)
	if err != nil {
		return err
	}
	
	offset := image.Pt(x, y)
	draw.Draw(canvas, img.Bounds().Add(offset), img, image.Point{}, draw.Over)
	return nil
}

func addText(img *image.RGBA, x, y int, label string, col color.Color) {
	point := fixed.Point26_6{
		X: fixed.Int26_6(x * 64),
		Y: fixed.Int26_6(y * 64),
	}
	
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}

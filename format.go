package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
	"golang.org/x/image/font/gofont/gomonobold"
	"golang.org/x/image/font/gofont/goregular"
)

const (
	WIDTH      = 991
	HEIGHT     = 306
	MARGINS    = 30
	QR_CODE_LW = 241
)

func formatLabel(itemId, serial, name string) error {
	// Create blank white label
	canvas := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))
	draw.Draw(canvas, canvas.Bounds(), image.White, image.Point{}, draw.Src)

	// Add the text
	normalFont := "fonts/roboto-font/RobotoRegular.ttf"
	boldFont  := "fonts/roboto-font/RobotoBlack-Powx.ttf"

	addTextWithFont(canvas, MARGINS+30+25, MARGINS, "Monash Automation", 30, normalFont, false)
	addTextWithFont(canvas, MARGINS, HEIGHT/2-50, serial, 100, boldFont, true)
	addTextWithFont(canvas, MARGINS, HEIGHT-MARGINS-40, name, 40, normalFont, false)

	// Add the MA logo
	if err := overlayImage(canvas, "assets/monash_automation_logo.png", MARGINS, MARGINS, 40, 40); err != nil {
		return err
	}

	createQR(canvas, itemId, QR_CODE_LW)

	// Save the result
	outFile, err := os.Create("temp/label.png")
	if err != nil {
		return err
	}
	defer outFile.Close()

	return png.Encode(outFile, canvas)
}

func createQR(canvas *image.RGBA, itemId string, length int) error {
	qr, err := qrcode.New(itemId, qrcode.High)
	if err != nil {
		return err
	}

	qr.DisableBorder = true
	qrImg := qr.Image(length)

	// Overlay on canvas
	offset := image.Pt(WIDTH-length-MARGINS, MARGINS)
	draw.Draw(canvas, qrImg.Bounds().Add(offset), qrImg, image.Point{}, draw.Over)
	return nil
}

func overlayImage(canvas *image.RGBA, imagePath string, x, y, size_x, size_y int) error {
	file, err := os.Open(imagePath)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	img = resize.Resize(uint(size_x), uint(size_y), img, resize.Lanczos3)
	if err != nil {
		return err
	}

	offset := image.Pt(x, y)
	draw.Draw(canvas, img.Bounds().Add(offset), img, image.Point{}, draw.Over)
	return nil
}

func addTextWithFont(img *image.RGBA, x, y int, text string, fontSize float64, fontPath string, bold bool) {
	col := color.RGBA{0, 0, 0, 255}
	
	var fontData []byte
	if fontPath != "" {
		data, err := os.ReadFile(fontPath)
		if err == nil {
			fontData = data
		}
	}
	
	// Fallback to embedded fonts
	if fontData == nil {
		if bold {
			fontData = gomonobold.TTF
		} else {
			fontData = goregular.TTF
		}
	}
	
	f, _ := truetype.Parse(fontData)
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(f)
	c.SetFontSize(fontSize)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.NewUniform(col))
	pt := freetype.Pt(x, y+int(c.PointToFixed(fontSize)>>6))
	c.DrawString(text, pt)
}

package main

import (
    "fmt"
    "image"
    "image/png"
    "os"
)

func main() {
    // Load the PNG image
    file, err := os.Open("frog.png")
    if err != nil {
        fmt.Println("Error opening image:", err)
        return
    }
    defer file.Close()

    img, err := png.Decode(file)
    if err != nil {
        fmt.Println("Error decoding PNG:", err)
        return
    }

    // Open printer device
    device, err := os.OpenFile("/dev/usb/lp0", os.O_RDWR, 0)
    if err != nil {
        fmt.Println("Error opening device:", err)
        return
    }
    defer device.Close()

    // Send Brother QL-700 commands
    err = printImageQL700(device, img)
    if err != nil {
        fmt.Println("Error printing:", err)
        return
    }

    fmt.Println("Print job sent successfully!")
}

func printImageQL700(device *os.File, img image.Image) error {
    bounds := img.Bounds()
    width := bounds.Dx()
    height := bounds.Dy()

    // Initialize printer
    initCmd := []byte{
        0x1B, 0x40, // ESC @ - Initialize
        0x1B, 0x69, 0x61, 0x01, // Select automatic status mode
        0x1B, 0x69, 0x7A, 0x84, 0x00, // Set media & quality (62mm continuous)
        0x4D, 0x02, // Mode setting
    }
    device.Write(initCmd)

    // Convert image to raster data
    // Brother QL-700 expects monochrome data
    bytesPerLine := (width + 7) / 8
    
    for y := 0; y < height; y++ {
        lineData := make([]byte, bytesPerLine)
        
        for x := 0; x < width; x++ {
            r, g, b, _ := img.At(x+bounds.Min.X, y+bounds.Min.Y).RGBA()
            // Convert to grayscale and threshold
            gray := (r + g + b) / 3
            if gray < 32768 { // If darker than middle gray
                byteIndex := x / 8
                bitIndex := 7 - (x % 8)
                lineData[byteIndex] |= (1 << bitIndex)
            }
        }
        
        // Send raster line
        rasterCmd := []byte{0x67, 0x00, byte(bytesPerLine)}
        device.Write(rasterCmd)
        device.Write(lineData)
    }

    // Print command
    printCmd := []byte{0x1A} // Print with feeding
    device.Write(printCmd)

    return nil
}

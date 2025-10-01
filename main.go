package main

import (
    "fmt"
    "os"
)

// This attempts to write directly to the USB device
func printDirectToUSB(imageFile string) error {
    // On macOS, USB printers often appear in /dev
    // Try to find the device
    possiblePaths := []string{
        "/dev/usb/lp0",
        "/dev/lp0",
        "/dev/usblp0",
    }
    
    var devicePath string
    for _, path := range possiblePaths {
        if _, err := os.Stat(path); err == nil {
            devicePath = path
            break
        }
    }
    
    if devicePath == "" {
        return fmt.Errorf("could not find USB printer device")
    }
    
    fmt.Println("Found device at:", devicePath)
    
    // Read image
    imageData, err := os.ReadFile(imageFile)
    if err != nil {
        return fmt.Errorf("failed to read image: %v", err)
    }
    
    // Write directly to device
    err = os.WriteFile(devicePath, imageData, 0644)
    if err != nil {
        return fmt.Errorf("failed to write to device: %v", err)
    }
    
    return nil
}

func main() {
    err := printDirectToUSB("frog.png")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Data sent to printer!")
    }
}
